// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/GetAddressBookRequest
type GetAddressBookInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the address book for which to request details.
	//
	// AddressBookArn is a required field
	AddressBookArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetAddressBookInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAddressBookInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetAddressBookInput"}

	if s.AddressBookArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AddressBookArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/GetAddressBookResponse
type GetAddressBookOutput struct {
	_ struct{} `type:"structure"`

	// The details of the requested address book.
	AddressBook *AddressBook `type:"structure"`
}

// String returns the string representation
func (s GetAddressBookOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetAddressBook = "GetAddressBook"

// GetAddressBookRequest returns a request value for making API operation for
// Alexa For Business.
//
// Gets address the book details by the address book ARN.
//
//    // Example sending a request using GetAddressBookRequest.
//    req := client.GetAddressBookRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/GetAddressBook
func (c *Client) GetAddressBookRequest(input *GetAddressBookInput) GetAddressBookRequest {
	op := &aws.Operation{
		Name:       opGetAddressBook,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetAddressBookInput{}
	}

	req := c.newRequest(op, input, &GetAddressBookOutput{})
	return GetAddressBookRequest{Request: req, Input: input, Copy: c.GetAddressBookRequest}
}

// GetAddressBookRequest is the request type for the
// GetAddressBook API operation.
type GetAddressBookRequest struct {
	*aws.Request
	Input *GetAddressBookInput
	Copy  func(*GetAddressBookInput) GetAddressBookRequest
}

// Send marshals and sends the GetAddressBook API request.
func (r GetAddressBookRequest) Send(ctx context.Context) (*GetAddressBookResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAddressBookResponse{
		GetAddressBookOutput: r.Request.Data.(*GetAddressBookOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAddressBookResponse is the response type for the
// GetAddressBook API operation.
type GetAddressBookResponse struct {
	*GetAddressBookOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAddressBook request.
func (r *GetAddressBookResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
