// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to delete a dedicated IP pool.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/DeleteDedicatedIpPoolRequest
type DeleteDedicatedIpPoolInput struct {
	_ struct{} `type:"structure"`

	// The name of the dedicated IP pool that you want to delete.
	//
	// PoolName is a required field
	PoolName *string `location:"uri" locationName:"PoolName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDedicatedIpPoolInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDedicatedIpPoolInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDedicatedIpPoolInput"}

	if s.PoolName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PoolName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDedicatedIpPoolInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.PoolName != nil {
		v := *s.PoolName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "PoolName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An HTTP 200 response if the request succeeds, or an error message if the
// request fails.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/DeleteDedicatedIpPoolResponse
type DeleteDedicatedIpPoolOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDedicatedIpPoolOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDedicatedIpPoolOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteDedicatedIpPool = "DeleteDedicatedIpPool"

// DeleteDedicatedIpPoolRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Delete a dedicated IP pool.
//
//    // Example sending a request using DeleteDedicatedIpPoolRequest.
//    req := client.DeleteDedicatedIpPoolRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/DeleteDedicatedIpPool
func (c *Client) DeleteDedicatedIpPoolRequest(input *DeleteDedicatedIpPoolInput) DeleteDedicatedIpPoolRequest {
	op := &aws.Operation{
		Name:       opDeleteDedicatedIpPool,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/email/dedicated-ip-pools/{PoolName}",
	}

	if input == nil {
		input = &DeleteDedicatedIpPoolInput{}
	}

	req := c.newRequest(op, input, &DeleteDedicatedIpPoolOutput{})
	return DeleteDedicatedIpPoolRequest{Request: req, Input: input, Copy: c.DeleteDedicatedIpPoolRequest}
}

// DeleteDedicatedIpPoolRequest is the request type for the
// DeleteDedicatedIpPool API operation.
type DeleteDedicatedIpPoolRequest struct {
	*aws.Request
	Input *DeleteDedicatedIpPoolInput
	Copy  func(*DeleteDedicatedIpPoolInput) DeleteDedicatedIpPoolRequest
}

// Send marshals and sends the DeleteDedicatedIpPool API request.
func (r DeleteDedicatedIpPoolRequest) Send(ctx context.Context) (*DeleteDedicatedIpPoolResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDedicatedIpPoolResponse{
		DeleteDedicatedIpPoolOutput: r.Request.Data.(*DeleteDedicatedIpPoolOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDedicatedIpPoolResponse is the response type for the
// DeleteDedicatedIpPool API operation.
type DeleteDedicatedIpPoolResponse struct {
	*DeleteDedicatedIpPoolOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDedicatedIpPool request.
func (r *DeleteDedicatedIpPoolResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
