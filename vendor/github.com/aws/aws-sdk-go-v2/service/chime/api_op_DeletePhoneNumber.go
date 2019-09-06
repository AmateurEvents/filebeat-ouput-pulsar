// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/DeletePhoneNumberRequest
type DeletePhoneNumberInput struct {
	_ struct{} `type:"structure"`

	// The phone number ID.
	//
	// PhoneNumberId is a required field
	PhoneNumberId *string `location:"uri" locationName:"phoneNumberId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeletePhoneNumberInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeletePhoneNumberInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeletePhoneNumberInput"}

	if s.PhoneNumberId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PhoneNumberId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeletePhoneNumberInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.PhoneNumberId != nil {
		v := *s.PhoneNumberId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "phoneNumberId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/DeletePhoneNumberOutput
type DeletePhoneNumberOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeletePhoneNumberOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeletePhoneNumberOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeletePhoneNumber = "DeletePhoneNumber"

// DeletePhoneNumberRequest returns a request value for making API operation for
// Amazon Chime.
//
// Moves the specified phone number into the Deletion queue. A phone number
// must be disassociated from any users or Amazon Chime Voice Connectors before
// it can be deleted.
//
// Deleted phone numbers remain in the Deletion queue for 7 days before they
// are deleted permanently.
//
//    // Example sending a request using DeletePhoneNumberRequest.
//    req := client.DeletePhoneNumberRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/DeletePhoneNumber
func (c *Client) DeletePhoneNumberRequest(input *DeletePhoneNumberInput) DeletePhoneNumberRequest {
	op := &aws.Operation{
		Name:       opDeletePhoneNumber,
		HTTPMethod: "DELETE",
		HTTPPath:   "/phone-numbers/{phoneNumberId}",
	}

	if input == nil {
		input = &DeletePhoneNumberInput{}
	}

	req := c.newRequest(op, input, &DeletePhoneNumberOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeletePhoneNumberRequest{Request: req, Input: input, Copy: c.DeletePhoneNumberRequest}
}

// DeletePhoneNumberRequest is the request type for the
// DeletePhoneNumber API operation.
type DeletePhoneNumberRequest struct {
	*aws.Request
	Input *DeletePhoneNumberInput
	Copy  func(*DeletePhoneNumberInput) DeletePhoneNumberRequest
}

// Send marshals and sends the DeletePhoneNumber API request.
func (r DeletePhoneNumberRequest) Send(ctx context.Context) (*DeletePhoneNumberResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeletePhoneNumberResponse{
		DeletePhoneNumberOutput: r.Request.Data.(*DeletePhoneNumberOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeletePhoneNumberResponse is the response type for the
// DeletePhoneNumber API operation.
type DeletePhoneNumberResponse struct {
	*DeletePhoneNumberOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeletePhoneNumber request.
func (r *DeletePhoneNumberResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
