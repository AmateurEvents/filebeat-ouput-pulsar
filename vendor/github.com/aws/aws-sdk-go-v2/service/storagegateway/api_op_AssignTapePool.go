// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/AssignTapePoolInput
type AssignTapePoolInput struct {
	_ struct{} `type:"structure"`

	// The ID of the pool that you want to add your tape to for archiving. The tape
	// in this pool is archived in the S3 storage class that is associated with
	// the pool. When you use your backup application to eject the tape, the tape
	// is archived directly into the storage class (Glacier or Deep Archive) that
	// corresponds to the pool.
	//
	// Valid values: "GLACIER", "DEEP_ARCHIVE"
	//
	// PoolId is a required field
	PoolId *string `min:"1" type:"string" required:"true"`

	// The unique Amazon Resource Name (ARN) of the virtual tape that you want to
	// add to the tape pool.
	//
	// TapeARN is a required field
	TapeARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s AssignTapePoolInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssignTapePoolInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssignTapePoolInput"}

	if s.PoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PoolId"))
	}
	if s.PoolId != nil && len(*s.PoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PoolId", 1))
	}

	if s.TapeARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("TapeARN"))
	}
	if s.TapeARN != nil && len(*s.TapeARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("TapeARN", 50))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/AssignTapePoolOutput
type AssignTapePoolOutput struct {
	_ struct{} `type:"structure"`

	// The unique Amazon Resource Names (ARN) of the virtual tape that was added
	// to the tape pool.
	TapeARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s AssignTapePoolOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssignTapePool = "AssignTapePool"

// AssignTapePoolRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Assigns a tape to a tape pool for archiving. The tape assigned to a pool
// is archived in the S3 storage class that is associated with the pool. When
// you use your backup application to eject the tape, the tape is archived directly
// into the S3 storage class (Glacier or Deep Archive) that corresponds to the
// pool.
//
// Valid values: "GLACIER", "DEEP_ARCHIVE"
//
//    // Example sending a request using AssignTapePoolRequest.
//    req := client.AssignTapePoolRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/AssignTapePool
func (c *Client) AssignTapePoolRequest(input *AssignTapePoolInput) AssignTapePoolRequest {
	op := &aws.Operation{
		Name:       opAssignTapePool,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssignTapePoolInput{}
	}

	req := c.newRequest(op, input, &AssignTapePoolOutput{})
	return AssignTapePoolRequest{Request: req, Input: input, Copy: c.AssignTapePoolRequest}
}

// AssignTapePoolRequest is the request type for the
// AssignTapePool API operation.
type AssignTapePoolRequest struct {
	*aws.Request
	Input *AssignTapePoolInput
	Copy  func(*AssignTapePoolInput) AssignTapePoolRequest
}

// Send marshals and sends the AssignTapePool API request.
func (r AssignTapePoolRequest) Send(ctx context.Context) (*AssignTapePoolResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssignTapePoolResponse{
		AssignTapePoolOutput: r.Request.Data.(*AssignTapePoolOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssignTapePoolResponse is the response type for the
// AssignTapePool API operation.
type AssignTapePoolResponse struct {
	*AssignTapePoolOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssignTapePool request.
func (r *AssignTapePoolResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
