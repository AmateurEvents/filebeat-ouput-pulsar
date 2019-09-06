// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/AssociatePhoneNumberWithUserRequest
type AssociatePhoneNumberWithUserInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The phone number, in E.164 format.
	//
	// E164PhoneNumber is a required field
	E164PhoneNumber *string `type:"string" required:"true"`

	// The user ID.
	//
	// UserId is a required field
	UserId *string `location:"uri" locationName:"userId" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociatePhoneNumberWithUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociatePhoneNumberWithUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociatePhoneNumberWithUserInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.E164PhoneNumber == nil {
		invalidParams.Add(aws.NewErrParamRequired("E164PhoneNumber"))
	}

	if s.UserId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AssociatePhoneNumberWithUserInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.E164PhoneNumber != nil {
		v := *s.E164PhoneNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "E164PhoneNumber", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UserId != nil {
		v := *s.UserId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "userId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/AssociatePhoneNumberWithUserResponse
type AssociatePhoneNumberWithUserOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AssociatePhoneNumberWithUserOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AssociatePhoneNumberWithUserOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opAssociatePhoneNumberWithUser = "AssociatePhoneNumberWithUser"

// AssociatePhoneNumberWithUserRequest returns a request value for making API operation for
// Amazon Chime.
//
// Associates a phone number with the specified Amazon Chime user.
//
//    // Example sending a request using AssociatePhoneNumberWithUserRequest.
//    req := client.AssociatePhoneNumberWithUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/AssociatePhoneNumberWithUser
func (c *Client) AssociatePhoneNumberWithUserRequest(input *AssociatePhoneNumberWithUserInput) AssociatePhoneNumberWithUserRequest {
	op := &aws.Operation{
		Name:       opAssociatePhoneNumberWithUser,
		HTTPMethod: "POST",
		HTTPPath:   "/accounts/{accountId}/users/{userId}?operation=associate-phone-number",
	}

	if input == nil {
		input = &AssociatePhoneNumberWithUserInput{}
	}

	req := c.newRequest(op, input, &AssociatePhoneNumberWithUserOutput{})
	return AssociatePhoneNumberWithUserRequest{Request: req, Input: input, Copy: c.AssociatePhoneNumberWithUserRequest}
}

// AssociatePhoneNumberWithUserRequest is the request type for the
// AssociatePhoneNumberWithUser API operation.
type AssociatePhoneNumberWithUserRequest struct {
	*aws.Request
	Input *AssociatePhoneNumberWithUserInput
	Copy  func(*AssociatePhoneNumberWithUserInput) AssociatePhoneNumberWithUserRequest
}

// Send marshals and sends the AssociatePhoneNumberWithUser API request.
func (r AssociatePhoneNumberWithUserRequest) Send(ctx context.Context) (*AssociatePhoneNumberWithUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociatePhoneNumberWithUserResponse{
		AssociatePhoneNumberWithUserOutput: r.Request.Data.(*AssociatePhoneNumberWithUserOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociatePhoneNumberWithUserResponse is the response type for the
// AssociatePhoneNumberWithUser API operation.
type AssociatePhoneNumberWithUserResponse struct {
	*AssociatePhoneNumberWithUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociatePhoneNumberWithUser request.
func (r *AssociatePhoneNumberWithUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
