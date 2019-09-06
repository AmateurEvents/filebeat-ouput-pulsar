// +build example

package main

import (
	"context"
	"fmt"
	"os"
	"sort"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

// Lists all encrypted objects owned by an account. The `accounts` string
// contains a list of profiles to use.
//
// Usage:
// listObjectsConcurrentlv
func main() {
	accounts := []string{"default", "default2", "otherprofile"}

	// Spin off a worker for each account to retrieve that account's
	bucketCh := make(chan Bucket, 5)
	var wg sync.WaitGroup
	for _, acc := range accounts {
		wg.Add(1)
		go func(acc string) {
			defer wg.Done()

			cfg, err := external.LoadDefaultAWSConfig(
				external.WithSharedConfigProfile(acc),
			)
			if err != nil {
				exitErrorf("failed to load config for account, %s, %v\n", acc, err)
			}
			if err = getAccountBuckets(cfg, bucketCh, acc); err != nil {
				exitErrorf("failed to get account %s's bucket info, %v\n", acc, err)
			}
		}(acc)
	}
	// Spin off a goroutine which will wait until all account buckets have been collected and
	// added to the bucketCh. Close the bucketCh so the for range below will exit once all
	// bucket info is printed.
	go func() {
		wg.Wait()
		close(bucketCh)
	}()

	// Receive from the bucket channel printing the information for each bucket to the console
	// when the bucketCh channel is drained.
	buckets := []Bucket{}
	for b := range bucketCh {
		buckets = append(buckets, b)
	}

	sortBuckets(buckets)
	for _, b := range buckets {
		if b.Error != nil {
			fmt.Fprintf(os.Stderr, "Bucket %s, owned by: %s, failed: %v\n", b.Name, b.Owner, b.Error)
			continue
		}

		encObjs := b.encryptedObjects()
		fmt.Printf("Bucket: %s, owned by: %s, total objects: %d, failed objects: %d, encrypted objects: %d\n",
			b.Name, b.Owner, len(b.Objects), len(b.ErrObjects), len(encObjs))
		if len(encObjs) > 0 {
			for _, encObj := range encObjs {
				fmt.Printf("\t%s %s:%s/%s\n", encObj.EncryptionType, b.Region, b.Name, encObj.Key)
			}
		}
	}
}

func sortBuckets(buckets []Bucket) {
	s := sortalbeBuckets(buckets)
	sort.Sort(s)
}

type sortalbeBuckets []Bucket

func (s sortalbeBuckets) Len() int      { return len(s) }
func (s sortalbeBuckets) Swap(a, b int) { s[a], s[b] = s[b], s[a] }
func (s sortalbeBuckets) Less(a, b int) bool {
	if s[a].Owner == s[b].Owner && s[a].Name < s[b].Name {
		return true
	}

	if s[a].Owner < s[b].Owner {
		return true
	}

	return false
}

func getAccountBuckets(cfg aws.Config, bucketCh chan<- Bucket, owner string) error {
	svc := s3.New(cfg)
	buckets, err := listBuckets(svc)
	if err != nil {
		return fmt.Errorf("failed to list buckets, %v", err)
	}
	for _, bucket := range buckets {
		bucket.Owner = owner
		if bucket.Error != nil {
			continue
		}

		cfgCp := cfg.Copy()
		cfgCp.Region = bucket.Region

		bckSvc := s3.New(cfgCp)
		bucketDetails(bckSvc, bucket)
		bucketCh <- bucket
	}

	return nil
}

func bucketDetails(svc *s3.Client, bucket Bucket) {
	objs, errObjs, err := listBucketObjects(svc, bucket.Name)
	if err != nil {
		bucket.Error = err
	} else {
		bucket.Objects = objs
		bucket.ErrObjects = errObjs
	}
}

// A Object provides details of an S3 object
type Object struct {
	Bucket         string
	Key            string
	Encrypted      bool
	EncryptionType string
}

// An ErrObject provides details of the error occurred retrieving
// an object's status.
type ErrObject struct {
	Bucket string
	Key    string
	Error  error
}

// A Bucket provides details about a bucket and its objects
type Bucket struct {
	Owner        string
	Name         string
	CreationDate time.Time
	Region       string
	Objects      []Object
	Error        error
	ErrObjects   []ErrObject
}

func (b *Bucket) encryptedObjects() []Object {
	encObjs := []Object{}
	for _, obj := range b.Objects {
		if obj.Encrypted {
			encObjs = append(encObjs, obj)
		}
	}
	return encObjs
}

func listBuckets(svc *s3.Client) ([]Bucket, error) {
	listReq := svc.ListBucketsRequest(&s3.ListBucketsInput{})
	res, err := listReq.Send(context.Background())
	if err != nil {
		return nil, err
	}

	buckets := make([]Bucket, len(res.Buckets))
	for i, b := range res.Buckets {
		buckets[i] = Bucket{
			Name:         *b.Name,
			CreationDate: *b.CreationDate,
		}

		getReq := svc.GetBucketLocationRequest(&s3.GetBucketLocationInput{
			Bucket: b.Name,
		})
		locRes, err := getReq.Send(context.Background())
		if err != nil {
			buckets[i].Error = err
			continue
		}

		if len(locRes.LocationConstraint) == 0 {
			buckets[i].Region = "us-east-1"
		} else {
			buckets[i].Region = string(locRes.LocationConstraint)
		}
	}

	return buckets, nil
}

func listBucketObjects(svc *s3.Client, bucket string) ([]Object, []ErrObject, error) {
	listReq := svc.ListObjectsRequest(&s3.ListObjectsInput{
		Bucket: &bucket,
	})
	listRes, err := listReq.Send(context.Background())
	if err != nil {
		return nil, nil, err
	}

	objs := make([]Object, 0, len(listRes.Contents))
	errObjs := []ErrObject{}
	for _, listObj := range listRes.Contents {
		objReq := svc.HeadObjectRequest(&s3.HeadObjectInput{
			Bucket: &bucket,
			Key:    listObj.Key,
		})
		objData, err := objReq.Send(context.Background())
		if err != nil {
			errObjs = append(errObjs, ErrObject{Bucket: bucket, Key: *listObj.Key, Error: err})
			continue
		}

		obj := Object{Bucket: bucket, Key: *listObj.Key}
		if len(objData.ServerSideEncryption) > 0 {
			obj.Encrypted = true
			obj.EncryptionType = string(objData.ServerSideEncryption)
		}

		objs = append(objs, obj)
	}

	return objs, errObjs, nil
}
