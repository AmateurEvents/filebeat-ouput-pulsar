// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloud9

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloud9-2017-09-23/DeleteEnvironmentMembershipRequest
type DeleteEnvironmentMembershipInput struct {
	_ struct{} `type:"structure"`

	// The ID of the environment to delete the environment member from.
	//
	// EnvironmentId is a required field
	EnvironmentId *string `locationName:"environmentId" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the environment member to delete from the
	// environment.
	//
	// UserArn is a required field
	UserArn *string `locationName:"userArn" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteEnvironmentMembershipInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteEnvironmentMembershipInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteEnvironmentMembershipInput"}

	if s.EnvironmentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EnvironmentId"))
	}

	if s.UserArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloud9-2017-09-23/DeleteEnvironmentMembershipResult
type DeleteEnvironmentMembershipOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteEnvironmentMembershipOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteEnvironmentMembership = "DeleteEnvironmentMembership"

// DeleteEnvironmentMembershipRequest returns a request value for making API operation for
// AWS Cloud9.
//
// Deletes an environment member from an AWS Cloud9 development environment.
//
//    // Example sending a request using DeleteEnvironmentMembershipRequest.
//    req := client.DeleteEnvironmentMembershipRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloud9-2017-09-23/DeleteEnvironmentMembership
func (c *Client) DeleteEnvironmentMembershipRequest(input *DeleteEnvironmentMembershipInput) DeleteEnvironmentMembershipRequest {
	op := &aws.Operation{
		Name:       opDeleteEnvironmentMembership,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteEnvironmentMembershipInput{}
	}

	req := c.newRequest(op, input, &DeleteEnvironmentMembershipOutput{})
	return DeleteEnvironmentMembershipRequest{Request: req, Input: input, Copy: c.DeleteEnvironmentMembershipRequest}
}

// DeleteEnvironmentMembershipRequest is the request type for the
// DeleteEnvironmentMembership API operation.
type DeleteEnvironmentMembershipRequest struct {
	*aws.Request
	Input *DeleteEnvironmentMembershipInput
	Copy  func(*DeleteEnvironmentMembershipInput) DeleteEnvironmentMembershipRequest
}

// Send marshals and sends the DeleteEnvironmentMembership API request.
func (r DeleteEnvironmentMembershipRequest) Send(ctx context.Context) (*DeleteEnvironmentMembershipResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteEnvironmentMembershipResponse{
		DeleteEnvironmentMembershipOutput: r.Request.Data.(*DeleteEnvironmentMembershipOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteEnvironmentMembershipResponse is the response type for the
// DeleteEnvironmentMembership API operation.
type DeleteEnvironmentMembershipResponse struct {
	*DeleteEnvironmentMembershipOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteEnvironmentMembership request.
func (r *DeleteEnvironmentMembershipResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
