// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/UpdateAccountRequest
type UpdateAccountInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The new name for the specified Amazon Chime account.
	Name *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateAccountInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAccountInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateAccountInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateAccountInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/UpdateAccountResponse
type UpdateAccountOutput struct {
	_ struct{} `type:"structure"`

	// The updated Amazon Chime account details.
	Account *Account `type:"structure"`
}

// String returns the string representation
func (s UpdateAccountOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateAccountOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Account != nil {
		v := s.Account

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Account", v, metadata)
	}
	return nil
}

const opUpdateAccount = "UpdateAccount"

// UpdateAccountRequest returns a request value for making API operation for
// Amazon Chime.
//
// Updates account details for the specified Amazon Chime account. Currently,
// only account name updates are supported for this action.
//
//    // Example sending a request using UpdateAccountRequest.
//    req := client.UpdateAccountRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/UpdateAccount
func (c *Client) UpdateAccountRequest(input *UpdateAccountInput) UpdateAccountRequest {
	op := &aws.Operation{
		Name:       opUpdateAccount,
		HTTPMethod: "POST",
		HTTPPath:   "/accounts/{accountId}",
	}

	if input == nil {
		input = &UpdateAccountInput{}
	}

	req := c.newRequest(op, input, &UpdateAccountOutput{})
	return UpdateAccountRequest{Request: req, Input: input, Copy: c.UpdateAccountRequest}
}

// UpdateAccountRequest is the request type for the
// UpdateAccount API operation.
type UpdateAccountRequest struct {
	*aws.Request
	Input *UpdateAccountInput
	Copy  func(*UpdateAccountInput) UpdateAccountRequest
}

// Send marshals and sends the UpdateAccount API request.
func (r UpdateAccountRequest) Send(ctx context.Context) (*UpdateAccountResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateAccountResponse{
		UpdateAccountOutput: r.Request.Data.(*UpdateAccountOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateAccountResponse is the response type for the
// UpdateAccount API operation.
type UpdateAccountResponse struct {
	*UpdateAccountOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateAccount request.
func (r *UpdateAccountResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
