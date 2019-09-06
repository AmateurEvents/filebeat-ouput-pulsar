// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/AssociatePhoneNumbersWithVoiceConnectorRequest
type AssociatePhoneNumbersWithVoiceConnectorInput struct {
	_ struct{} `type:"structure"`

	// List of phone numbers, in E.164 format.
	E164PhoneNumbers []string `type:"list"`

	// The Amazon Chime Voice Connector ID.
	//
	// VoiceConnectorId is a required field
	VoiceConnectorId *string `location:"uri" locationName:"voiceConnectorId" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociatePhoneNumbersWithVoiceConnectorInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociatePhoneNumbersWithVoiceConnectorInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociatePhoneNumbersWithVoiceConnectorInput"}

	if s.VoiceConnectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VoiceConnectorId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AssociatePhoneNumbersWithVoiceConnectorInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.E164PhoneNumbers != nil {
		v := s.E164PhoneNumbers

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "E164PhoneNumbers", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.VoiceConnectorId != nil {
		v := *s.VoiceConnectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "voiceConnectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/AssociatePhoneNumbersWithVoiceConnectorResponse
type AssociatePhoneNumbersWithVoiceConnectorOutput struct {
	_ struct{} `type:"structure"`

	// If the action fails for one or more of the phone numbers in the request,
	// a list of the phone numbers is returned, along with error codes and error
	// messages.
	PhoneNumberErrors []PhoneNumberError `type:"list"`
}

// String returns the string representation
func (s AssociatePhoneNumbersWithVoiceConnectorOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AssociatePhoneNumbersWithVoiceConnectorOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.PhoneNumberErrors != nil {
		v := s.PhoneNumberErrors

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "PhoneNumberErrors", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opAssociatePhoneNumbersWithVoiceConnector = "AssociatePhoneNumbersWithVoiceConnector"

// AssociatePhoneNumbersWithVoiceConnectorRequest returns a request value for making API operation for
// Amazon Chime.
//
// Associates a phone number with the specified Amazon Chime Voice Connector.
//
//    // Example sending a request using AssociatePhoneNumbersWithVoiceConnectorRequest.
//    req := client.AssociatePhoneNumbersWithVoiceConnectorRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/AssociatePhoneNumbersWithVoiceConnector
func (c *Client) AssociatePhoneNumbersWithVoiceConnectorRequest(input *AssociatePhoneNumbersWithVoiceConnectorInput) AssociatePhoneNumbersWithVoiceConnectorRequest {
	op := &aws.Operation{
		Name:       opAssociatePhoneNumbersWithVoiceConnector,
		HTTPMethod: "POST",
		HTTPPath:   "/voice-connectors/{voiceConnectorId}?operation=associate-phone-numbers",
	}

	if input == nil {
		input = &AssociatePhoneNumbersWithVoiceConnectorInput{}
	}

	req := c.newRequest(op, input, &AssociatePhoneNumbersWithVoiceConnectorOutput{})
	return AssociatePhoneNumbersWithVoiceConnectorRequest{Request: req, Input: input, Copy: c.AssociatePhoneNumbersWithVoiceConnectorRequest}
}

// AssociatePhoneNumbersWithVoiceConnectorRequest is the request type for the
// AssociatePhoneNumbersWithVoiceConnector API operation.
type AssociatePhoneNumbersWithVoiceConnectorRequest struct {
	*aws.Request
	Input *AssociatePhoneNumbersWithVoiceConnectorInput
	Copy  func(*AssociatePhoneNumbersWithVoiceConnectorInput) AssociatePhoneNumbersWithVoiceConnectorRequest
}

// Send marshals and sends the AssociatePhoneNumbersWithVoiceConnector API request.
func (r AssociatePhoneNumbersWithVoiceConnectorRequest) Send(ctx context.Context) (*AssociatePhoneNumbersWithVoiceConnectorResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociatePhoneNumbersWithVoiceConnectorResponse{
		AssociatePhoneNumbersWithVoiceConnectorOutput: r.Request.Data.(*AssociatePhoneNumbersWithVoiceConnectorOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociatePhoneNumbersWithVoiceConnectorResponse is the response type for the
// AssociatePhoneNumbersWithVoiceConnector API operation.
type AssociatePhoneNumbersWithVoiceConnectorResponse struct {
	*AssociatePhoneNumbersWithVoiceConnectorOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociatePhoneNumbersWithVoiceConnector request.
func (r *AssociatePhoneNumbersWithVoiceConnectorResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
