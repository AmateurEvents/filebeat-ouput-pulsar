// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/GenerateRandomRequest
type GenerateRandomInput struct {
	_ struct{} `type:"structure"`

	// Generates the random byte string in the AWS CloudHSM cluster that is associated
	// with the specified custom key store (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html).
	// To find the ID of a custom key store, use the DescribeCustomKeyStores operation.
	CustomKeyStoreId *string `min:"1" type:"string"`

	// The length of the byte string.
	NumberOfBytes *int64 `min:"1" type:"integer"`
}

// String returns the string representation
func (s GenerateRandomInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GenerateRandomInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GenerateRandomInput"}
	if s.CustomKeyStoreId != nil && len(*s.CustomKeyStoreId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CustomKeyStoreId", 1))
	}
	if s.NumberOfBytes != nil && *s.NumberOfBytes < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("NumberOfBytes", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/GenerateRandomResponse
type GenerateRandomOutput struct {
	_ struct{} `type:"structure"`

	// The random byte string. When you use the HTTP API or the AWS CLI, the value
	// is Base64-encoded. Otherwise, it is not encoded.
	//
	// Plaintext is automatically base64 encoded/decoded by the SDK.
	Plaintext []byte `min:"1" type:"blob"`
}

// String returns the string representation
func (s GenerateRandomOutput) String() string {
	return awsutil.Prettify(s)
}

const opGenerateRandom = "GenerateRandom"

// GenerateRandomRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Returns a random byte string that is cryptographically secure.
//
// By default, the random byte string is generated in AWS KMS. To generate the
// byte string in the AWS CloudHSM cluster that is associated with a custom
// key store (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html),
// specify the custom key store ID.
//
// For more information about entropy and random number generation, see the
// AWS Key Management Service Cryptographic Details (https://d0.awsstatic.com/whitepapers/KMS-Cryptographic-Details.pdf)
// whitepaper.
//
//    // Example sending a request using GenerateRandomRequest.
//    req := client.GenerateRandomRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/GenerateRandom
func (c *Client) GenerateRandomRequest(input *GenerateRandomInput) GenerateRandomRequest {
	op := &aws.Operation{
		Name:       opGenerateRandom,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GenerateRandomInput{}
	}

	req := c.newRequest(op, input, &GenerateRandomOutput{})
	return GenerateRandomRequest{Request: req, Input: input, Copy: c.GenerateRandomRequest}
}

// GenerateRandomRequest is the request type for the
// GenerateRandom API operation.
type GenerateRandomRequest struct {
	*aws.Request
	Input *GenerateRandomInput
	Copy  func(*GenerateRandomInput) GenerateRandomRequest
}

// Send marshals and sends the GenerateRandom API request.
func (r GenerateRandomRequest) Send(ctx context.Context) (*GenerateRandomResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GenerateRandomResponse{
		GenerateRandomOutput: r.Request.Data.(*GenerateRandomOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GenerateRandomResponse is the response type for the
// GenerateRandom API operation.
type GenerateRandomResponse struct {
	*GenerateRandomOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GenerateRandom request.
func (r *GenerateRandomResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
