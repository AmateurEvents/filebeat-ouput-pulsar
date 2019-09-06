// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// CancelArchivalInput
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/CancelArchivalInput
type CancelArchivalInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and region.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the virtual tape you want to cancel archiving
	// for.
	//
	// TapeARN is a required field
	TapeARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s CancelArchivalInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelArchivalInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelArchivalInput"}

	if s.GatewayARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayARN"))
	}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
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

// CancelArchivalOutput
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/CancelArchivalOutput
type CancelArchivalOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the virtual tape for which archiving was
	// canceled.
	TapeARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s CancelArchivalOutput) String() string {
	return awsutil.Prettify(s)
}

const opCancelArchival = "CancelArchival"

// CancelArchivalRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Cancels archiving of a virtual tape to the virtual tape shelf (VTS) after
// the archiving process is initiated. This operation is only supported in the
// tape gateway type.
//
//    // Example sending a request using CancelArchivalRequest.
//    req := client.CancelArchivalRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/CancelArchival
func (c *Client) CancelArchivalRequest(input *CancelArchivalInput) CancelArchivalRequest {
	op := &aws.Operation{
		Name:       opCancelArchival,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CancelArchivalInput{}
	}

	req := c.newRequest(op, input, &CancelArchivalOutput{})
	return CancelArchivalRequest{Request: req, Input: input, Copy: c.CancelArchivalRequest}
}

// CancelArchivalRequest is the request type for the
// CancelArchival API operation.
type CancelArchivalRequest struct {
	*aws.Request
	Input *CancelArchivalInput
	Copy  func(*CancelArchivalInput) CancelArchivalRequest
}

// Send marshals and sends the CancelArchival API request.
func (r CancelArchivalRequest) Send(ctx context.Context) (*CancelArchivalResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelArchivalResponse{
		CancelArchivalOutput: r.Request.Data.(*CancelArchivalOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelArchivalResponse is the response type for the
// CancelArchival API operation.
type CancelArchivalResponse struct {
	*CancelArchivalOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelArchival request.
func (r *CancelArchivalResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
