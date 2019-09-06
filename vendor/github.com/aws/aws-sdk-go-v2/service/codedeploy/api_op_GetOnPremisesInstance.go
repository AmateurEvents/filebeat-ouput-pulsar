// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a GetOnPremisesInstance operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/GetOnPremisesInstanceInput
type GetOnPremisesInstanceInput struct {
	_ struct{} `type:"structure"`

	// The name of the on-premises instance about which to get information.
	//
	// InstanceName is a required field
	InstanceName *string `locationName:"instanceName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetOnPremisesInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetOnPremisesInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetOnPremisesInstanceInput"}

	if s.InstanceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a GetOnPremisesInstance operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/GetOnPremisesInstanceOutput
type GetOnPremisesInstanceOutput struct {
	_ struct{} `type:"structure"`

	// Information about the on-premises instance.
	InstanceInfo *InstanceInfo `locationName:"instanceInfo" type:"structure"`
}

// String returns the string representation
func (s GetOnPremisesInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetOnPremisesInstance = "GetOnPremisesInstance"

// GetOnPremisesInstanceRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Gets information about an on-premises instance.
//
//    // Example sending a request using GetOnPremisesInstanceRequest.
//    req := client.GetOnPremisesInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/GetOnPremisesInstance
func (c *Client) GetOnPremisesInstanceRequest(input *GetOnPremisesInstanceInput) GetOnPremisesInstanceRequest {
	op := &aws.Operation{
		Name:       opGetOnPremisesInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetOnPremisesInstanceInput{}
	}

	req := c.newRequest(op, input, &GetOnPremisesInstanceOutput{})
	return GetOnPremisesInstanceRequest{Request: req, Input: input, Copy: c.GetOnPremisesInstanceRequest}
}

// GetOnPremisesInstanceRequest is the request type for the
// GetOnPremisesInstance API operation.
type GetOnPremisesInstanceRequest struct {
	*aws.Request
	Input *GetOnPremisesInstanceInput
	Copy  func(*GetOnPremisesInstanceInput) GetOnPremisesInstanceRequest
}

// Send marshals and sends the GetOnPremisesInstance API request.
func (r GetOnPremisesInstanceRequest) Send(ctx context.Context) (*GetOnPremisesInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetOnPremisesInstanceResponse{
		GetOnPremisesInstanceOutput: r.Request.Data.(*GetOnPremisesInstanceOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetOnPremisesInstanceResponse is the response type for the
// GetOnPremisesInstance API operation.
type GetOnPremisesInstanceResponse struct {
	*GetOnPremisesInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetOnPremisesInstance request.
func (r *GetOnPremisesInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
