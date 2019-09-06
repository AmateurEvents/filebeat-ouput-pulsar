// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentity

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Input to the UnlinkDeveloperIdentity action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-identity-2014-06-30/UnlinkDeveloperIdentityInput
type UnlinkDeveloperIdentityInput struct {
	_ struct{} `type:"structure"`

	// The "domain" by which Cognito will refer to your users.
	//
	// DeveloperProviderName is a required field
	DeveloperProviderName *string `min:"1" type:"string" required:"true"`

	// A unique ID used by your backend authentication process to identify a user.
	//
	// DeveloperUserIdentifier is a required field
	DeveloperUserIdentifier *string `min:"1" type:"string" required:"true"`

	// A unique identifier in the format REGION:GUID.
	//
	// IdentityId is a required field
	IdentityId *string `min:"1" type:"string" required:"true"`

	// An identity pool ID in the format REGION:GUID.
	//
	// IdentityPoolId is a required field
	IdentityPoolId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UnlinkDeveloperIdentityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UnlinkDeveloperIdentityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UnlinkDeveloperIdentityInput"}

	if s.DeveloperProviderName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeveloperProviderName"))
	}
	if s.DeveloperProviderName != nil && len(*s.DeveloperProviderName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeveloperProviderName", 1))
	}

	if s.DeveloperUserIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeveloperUserIdentifier"))
	}
	if s.DeveloperUserIdentifier != nil && len(*s.DeveloperUserIdentifier) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeveloperUserIdentifier", 1))
	}

	if s.IdentityId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdentityId"))
	}
	if s.IdentityId != nil && len(*s.IdentityId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdentityId", 1))
	}

	if s.IdentityPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdentityPoolId"))
	}
	if s.IdentityPoolId != nil && len(*s.IdentityPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdentityPoolId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-identity-2014-06-30/UnlinkDeveloperIdentityOutput
type UnlinkDeveloperIdentityOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UnlinkDeveloperIdentityOutput) String() string {
	return awsutil.Prettify(s)
}

const opUnlinkDeveloperIdentity = "UnlinkDeveloperIdentity"

// UnlinkDeveloperIdentityRequest returns a request value for making API operation for
// Amazon Cognito Identity.
//
// Unlinks a DeveloperUserIdentifier from an existing identity. Unlinked developer
// users will be considered new identities next time they are seen. If, for
// a given Cognito identity, you remove all federated identities as well as
// the developer user identifier, the Cognito identity becomes inaccessible.
//
// You must use AWS Developer credentials to call this API.
//
//    // Example sending a request using UnlinkDeveloperIdentityRequest.
//    req := client.UnlinkDeveloperIdentityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-identity-2014-06-30/UnlinkDeveloperIdentity
func (c *Client) UnlinkDeveloperIdentityRequest(input *UnlinkDeveloperIdentityInput) UnlinkDeveloperIdentityRequest {
	op := &aws.Operation{
		Name:       opUnlinkDeveloperIdentity,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UnlinkDeveloperIdentityInput{}
	}

	req := c.newRequest(op, input, &UnlinkDeveloperIdentityOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UnlinkDeveloperIdentityRequest{Request: req, Input: input, Copy: c.UnlinkDeveloperIdentityRequest}
}

// UnlinkDeveloperIdentityRequest is the request type for the
// UnlinkDeveloperIdentity API operation.
type UnlinkDeveloperIdentityRequest struct {
	*aws.Request
	Input *UnlinkDeveloperIdentityInput
	Copy  func(*UnlinkDeveloperIdentityInput) UnlinkDeveloperIdentityRequest
}

// Send marshals and sends the UnlinkDeveloperIdentity API request.
func (r UnlinkDeveloperIdentityRequest) Send(ctx context.Context) (*UnlinkDeveloperIdentityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UnlinkDeveloperIdentityResponse{
		UnlinkDeveloperIdentityOutput: r.Request.Data.(*UnlinkDeveloperIdentityOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UnlinkDeveloperIdentityResponse is the response type for the
// UnlinkDeveloperIdentity API operation.
type UnlinkDeveloperIdentityResponse struct {
	*UnlinkDeveloperIdentityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UnlinkDeveloperIdentity request.
func (r *UnlinkDeveloperIdentityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
