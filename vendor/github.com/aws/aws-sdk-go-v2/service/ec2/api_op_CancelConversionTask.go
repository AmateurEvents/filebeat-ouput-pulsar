// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CancelConversionRequest
type CancelConversionTaskInput struct {
	_ struct{} `type:"structure"`

	// The ID of the conversion task.
	//
	// ConversionTaskId is a required field
	ConversionTaskId *string `locationName:"conversionTaskId" type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The reason for canceling the conversion task.
	ReasonMessage *string `locationName:"reasonMessage" type:"string"`
}

// String returns the string representation
func (s CancelConversionTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelConversionTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelConversionTaskInput"}

	if s.ConversionTaskId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConversionTaskId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CancelConversionTaskOutput
type CancelConversionTaskOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CancelConversionTaskOutput) String() string {
	return awsutil.Prettify(s)
}

const opCancelConversionTask = "CancelConversionTask"

// CancelConversionTaskRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Cancels an active conversion task. The task can be the import of an instance
// or volume. The action removes all artifacts of the conversion, including
// a partially uploaded volume or instance. If the conversion is complete or
// is in the process of transferring the final disk image, the command fails
// and returns an exception.
//
// For more information, see Importing a Virtual Machine Using the Amazon EC2
// CLI (https://docs.aws.amazon.com/AWSEC2/latest/CommandLineReference/ec2-cli-vmimport-export.html).
//
//    // Example sending a request using CancelConversionTaskRequest.
//    req := client.CancelConversionTaskRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CancelConversionTask
func (c *Client) CancelConversionTaskRequest(input *CancelConversionTaskInput) CancelConversionTaskRequest {
	op := &aws.Operation{
		Name:       opCancelConversionTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CancelConversionTaskInput{}
	}

	req := c.newRequest(op, input, &CancelConversionTaskOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return CancelConversionTaskRequest{Request: req, Input: input, Copy: c.CancelConversionTaskRequest}
}

// CancelConversionTaskRequest is the request type for the
// CancelConversionTask API operation.
type CancelConversionTaskRequest struct {
	*aws.Request
	Input *CancelConversionTaskInput
	Copy  func(*CancelConversionTaskInput) CancelConversionTaskRequest
}

// Send marshals and sends the CancelConversionTask API request.
func (r CancelConversionTaskRequest) Send(ctx context.Context) (*CancelConversionTaskResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelConversionTaskResponse{
		CancelConversionTaskOutput: r.Request.Data.(*CancelConversionTaskOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelConversionTaskResponse is the response type for the
// CancelConversionTask API operation.
type CancelConversionTaskResponse struct {
	*CancelConversionTaskOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelConversionTask request.
func (r *CancelConversionTaskResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
