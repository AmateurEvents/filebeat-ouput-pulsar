// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/DeleteApnsVoipSandboxChannelRequest
type DeleteApnsVoipSandboxChannelInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteApnsVoipSandboxChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteApnsVoipSandboxChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteApnsVoipSandboxChannelInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteApnsVoipSandboxChannelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/DeleteApnsVoipSandboxChannelResponse
type DeleteApnsVoipSandboxChannelOutput struct {
	_ struct{} `type:"structure" payload:"APNSVoipSandboxChannelResponse"`

	// Provides information about the status and settings of the APNs (Apple Push
	// Notification service) VoIP sandbox channel for an application.
	//
	// APNSVoipSandboxChannelResponse is a required field
	APNSVoipSandboxChannelResponse *APNSVoipSandboxChannelResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteApnsVoipSandboxChannelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteApnsVoipSandboxChannelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.APNSVoipSandboxChannelResponse != nil {
		v := s.APNSVoipSandboxChannelResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "APNSVoipSandboxChannelResponse", v, metadata)
	}
	return nil
}

const opDeleteApnsVoipSandboxChannel = "DeleteApnsVoipSandboxChannel"

// DeleteApnsVoipSandboxChannelRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Disables the APNs VoIP sandbox channel for an application and deletes any
// existing settings for the channel.
//
//    // Example sending a request using DeleteApnsVoipSandboxChannelRequest.
//    req := client.DeleteApnsVoipSandboxChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/DeleteApnsVoipSandboxChannel
func (c *Client) DeleteApnsVoipSandboxChannelRequest(input *DeleteApnsVoipSandboxChannelInput) DeleteApnsVoipSandboxChannelRequest {
	op := &aws.Operation{
		Name:       opDeleteApnsVoipSandboxChannel,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/apps/{application-id}/channels/apns_voip_sandbox",
	}

	if input == nil {
		input = &DeleteApnsVoipSandboxChannelInput{}
	}

	req := c.newRequest(op, input, &DeleteApnsVoipSandboxChannelOutput{})
	return DeleteApnsVoipSandboxChannelRequest{Request: req, Input: input, Copy: c.DeleteApnsVoipSandboxChannelRequest}
}

// DeleteApnsVoipSandboxChannelRequest is the request type for the
// DeleteApnsVoipSandboxChannel API operation.
type DeleteApnsVoipSandboxChannelRequest struct {
	*aws.Request
	Input *DeleteApnsVoipSandboxChannelInput
	Copy  func(*DeleteApnsVoipSandboxChannelInput) DeleteApnsVoipSandboxChannelRequest
}

// Send marshals and sends the DeleteApnsVoipSandboxChannel API request.
func (r DeleteApnsVoipSandboxChannelRequest) Send(ctx context.Context) (*DeleteApnsVoipSandboxChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteApnsVoipSandboxChannelResponse{
		DeleteApnsVoipSandboxChannelOutput: r.Request.Data.(*DeleteApnsVoipSandboxChannelOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteApnsVoipSandboxChannelResponse is the response type for the
// DeleteApnsVoipSandboxChannel API operation.
type DeleteApnsVoipSandboxChannelResponse struct {
	*DeleteApnsVoipSandboxChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteApnsVoipSandboxChannel request.
func (r *DeleteApnsVoipSandboxChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
