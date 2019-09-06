// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudhsmv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudhsmv2-2017-04-28/DescribeBackupsRequest
type DescribeBackupsInput struct {
	_ struct{} `type:"structure"`

	// One or more filters to limit the items returned in the response.
	//
	// Use the backupIds filter to return only the specified backups. Specify backups
	// by their backup identifier (ID).
	//
	// Use the sourceBackupIds filter to return only the backups created from a
	// source backup. The sourceBackupID of a source backup is returned by the CopyBackupToRegion
	// operation.
	//
	// Use the clusterIds filter to return only the backups for the specified clusters.
	// Specify clusters by their cluster identifier (ID).
	//
	// Use the states filter to return only backups that match the specified state.
	Filters map[string][]string `type:"map"`

	// The maximum number of backups to return in the response. When there are more
	// backups than the number you specify, the response contains a NextToken value.
	MaxResults *int64 `min:"1" type:"integer"`

	// The NextToken value that you received in the previous response. Use this
	// value to get more backups.
	NextToken *string `type:"string"`

	SortAscending *bool `type:"boolean"`
}

// String returns the string representation
func (s DescribeBackupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeBackupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeBackupsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudhsmv2-2017-04-28/DescribeBackupsResponse
type DescribeBackupsOutput struct {
	_ struct{} `type:"structure"`

	// A list of backups.
	Backups []Backup `type:"list"`

	// An opaque string that indicates that the response contains only a subset
	// of backups. Use this value in a subsequent DescribeBackups request to get
	// more backups.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeBackupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeBackups = "DescribeBackups"

// DescribeBackupsRequest returns a request value for making API operation for
// AWS CloudHSM V2.
//
// Gets information about backups of AWS CloudHSM clusters.
//
// This is a paginated operation, which means that each response might contain
// only a subset of all the backups. When the response contains only a subset
// of backups, it includes a NextToken value. Use this value in a subsequent
// DescribeBackups request to get more backups. When you receive a response
// with no NextToken (or an empty or null value), that means there are no more
// backups to get.
//
//    // Example sending a request using DescribeBackupsRequest.
//    req := client.DescribeBackupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudhsmv2-2017-04-28/DescribeBackups
func (c *Client) DescribeBackupsRequest(input *DescribeBackupsInput) DescribeBackupsRequest {
	op := &aws.Operation{
		Name:       opDescribeBackups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeBackupsInput{}
	}

	req := c.newRequest(op, input, &DescribeBackupsOutput{})
	return DescribeBackupsRequest{Request: req, Input: input, Copy: c.DescribeBackupsRequest}
}

// DescribeBackupsRequest is the request type for the
// DescribeBackups API operation.
type DescribeBackupsRequest struct {
	*aws.Request
	Input *DescribeBackupsInput
	Copy  func(*DescribeBackupsInput) DescribeBackupsRequest
}

// Send marshals and sends the DescribeBackups API request.
func (r DescribeBackupsRequest) Send(ctx context.Context) (*DescribeBackupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeBackupsResponse{
		DescribeBackupsOutput: r.Request.Data.(*DescribeBackupsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeBackupsRequestPaginator returns a paginator for DescribeBackups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeBackupsRequest(input)
//   p := cloudhsmv2.NewDescribeBackupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeBackupsPaginator(req DescribeBackupsRequest) DescribeBackupsPaginator {
	return DescribeBackupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeBackupsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeBackupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeBackupsPaginator struct {
	aws.Pager
}

func (p *DescribeBackupsPaginator) CurrentPage() *DescribeBackupsOutput {
	return p.Pager.CurrentPage().(*DescribeBackupsOutput)
}

// DescribeBackupsResponse is the response type for the
// DescribeBackups API operation.
type DescribeBackupsResponse struct {
	*DescribeBackupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeBackups request.
func (r *DescribeBackupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
