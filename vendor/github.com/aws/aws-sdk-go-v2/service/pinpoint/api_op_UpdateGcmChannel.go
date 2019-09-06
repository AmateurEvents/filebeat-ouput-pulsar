// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/UpdateGcmChannelRequest
type UpdateGcmChannelInput struct {
	_ struct{} `type:"structure" payload:"GCMChannelRequest"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	// Specifies the status and settings of the GCM channel for an application.
	// This channel enables Amazon Pinpoint to send push notifications through the
	// Firebase Cloud Messaging (FCM), formerly Google Cloud Messaging (GCM), service.
	//
	// GCMChannelRequest is a required field
	GCMChannelRequest *GCMChannelRequest `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateGcmChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateGcmChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateGcmChannelInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.GCMChannelRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("GCMChannelRequest"))
	}
	if s.GCMChannelRequest != nil {
		if err := s.GCMChannelRequest.Validate(); err != nil {
			invalidParams.AddNested("GCMChannelRequest", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateGcmChannelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.GCMChannelRequest != nil {
		v := s.GCMChannelRequest

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "GCMChannelRequest", v, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/UpdateGcmChannelResponse
type UpdateGcmChannelOutput struct {
	_ struct{} `type:"structure" payload:"GCMChannelResponse"`

	// Provides information about the status and settings of the GCM channel for
	// an application. The GCM channel enables Amazon Pinpoint to send push notifications
	// through the Firebase Cloud Messaging (FCM), formerly Google Cloud Messaging
	// (GCM), service.
	//
	// GCMChannelResponse is a required field
	GCMChannelResponse *GCMChannelResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateGcmChannelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateGcmChannelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.GCMChannelResponse != nil {
		v := s.GCMChannelResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "GCMChannelResponse", v, metadata)
	}
	return nil
}

const opUpdateGcmChannel = "UpdateGcmChannel"

// UpdateGcmChannelRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Updates the status and settings of the GCM channel for an application.
//
//    // Example sending a request using UpdateGcmChannelRequest.
//    req := client.UpdateGcmChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/UpdateGcmChannel
func (c *Client) UpdateGcmChannelRequest(input *UpdateGcmChannelInput) UpdateGcmChannelRequest {
	op := &aws.Operation{
		Name:       opUpdateGcmChannel,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/apps/{application-id}/channels/gcm",
	}

	if input == nil {
		input = &UpdateGcmChannelInput{}
	}

	req := c.newRequest(op, input, &UpdateGcmChannelOutput{})
	return UpdateGcmChannelRequest{Request: req, Input: input, Copy: c.UpdateGcmChannelRequest}
}

// UpdateGcmChannelRequest is the request type for the
// UpdateGcmChannel API operation.
type UpdateGcmChannelRequest struct {
	*aws.Request
	Input *UpdateGcmChannelInput
	Copy  func(*UpdateGcmChannelInput) UpdateGcmChannelRequest
}

// Send marshals and sends the UpdateGcmChannel API request.
func (r UpdateGcmChannelRequest) Send(ctx context.Context) (*UpdateGcmChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateGcmChannelResponse{
		UpdateGcmChannelOutput: r.Request.Data.(*UpdateGcmChannelOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateGcmChannelResponse is the response type for the
// UpdateGcmChannel API operation.
type UpdateGcmChannelResponse struct {
	*UpdateGcmChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateGcmChannel request.
func (r *UpdateGcmChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
