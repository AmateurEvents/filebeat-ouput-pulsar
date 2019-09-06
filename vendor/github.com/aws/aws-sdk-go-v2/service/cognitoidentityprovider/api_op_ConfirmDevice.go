// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Confirms the device request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/ConfirmDeviceRequest
type ConfirmDeviceInput struct {
	_ struct{} `type:"structure"`

	// The access token.
	//
	// AccessToken is a required field
	AccessToken *string `type:"string" required:"true"`

	// The device key.
	//
	// DeviceKey is a required field
	DeviceKey *string `min:"1" type:"string" required:"true"`

	// The device name.
	DeviceName *string `min:"1" type:"string"`

	// The configuration of the device secret verifier.
	DeviceSecretVerifierConfig *DeviceSecretVerifierConfigType `type:"structure"`
}

// String returns the string representation
func (s ConfirmDeviceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ConfirmDeviceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ConfirmDeviceInput"}

	if s.AccessToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccessToken"))
	}

	if s.DeviceKey == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeviceKey"))
	}
	if s.DeviceKey != nil && len(*s.DeviceKey) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeviceKey", 1))
	}
	if s.DeviceName != nil && len(*s.DeviceName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeviceName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Confirms the device response.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/ConfirmDeviceResponse
type ConfirmDeviceOutput struct {
	_ struct{} `type:"structure"`

	// Indicates whether the user confirmation is necessary to confirm the device
	// response.
	UserConfirmationNecessary *bool `type:"boolean"`
}

// String returns the string representation
func (s ConfirmDeviceOutput) String() string {
	return awsutil.Prettify(s)
}

const opConfirmDevice = "ConfirmDevice"

// ConfirmDeviceRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Confirms tracking of the device. This API call is the call that begins device
// tracking.
//
//    // Example sending a request using ConfirmDeviceRequest.
//    req := client.ConfirmDeviceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/ConfirmDevice
func (c *Client) ConfirmDeviceRequest(input *ConfirmDeviceInput) ConfirmDeviceRequest {
	op := &aws.Operation{
		Name:       opConfirmDevice,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ConfirmDeviceInput{}
	}

	req := c.newRequest(op, input, &ConfirmDeviceOutput{})
	return ConfirmDeviceRequest{Request: req, Input: input, Copy: c.ConfirmDeviceRequest}
}

// ConfirmDeviceRequest is the request type for the
// ConfirmDevice API operation.
type ConfirmDeviceRequest struct {
	*aws.Request
	Input *ConfirmDeviceInput
	Copy  func(*ConfirmDeviceInput) ConfirmDeviceRequest
}

// Send marshals and sends the ConfirmDevice API request.
func (r ConfirmDeviceRequest) Send(ctx context.Context) (*ConfirmDeviceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ConfirmDeviceResponse{
		ConfirmDeviceOutput: r.Request.Data.(*ConfirmDeviceOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ConfirmDeviceResponse is the response type for the
// ConfirmDevice API operation.
type ConfirmDeviceResponse struct {
	*ConfirmDeviceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ConfirmDevice request.
func (r *ConfirmDeviceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
