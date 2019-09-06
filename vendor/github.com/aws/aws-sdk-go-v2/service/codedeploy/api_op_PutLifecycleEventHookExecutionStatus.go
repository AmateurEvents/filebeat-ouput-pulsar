// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/PutLifecycleEventHookExecutionStatusInput
type PutLifecycleEventHookExecutionStatusInput struct {
	_ struct{} `type:"structure"`

	// The unique ID of a deployment. Pass this ID to a Lambda function that validates
	// a deployment lifecycle event.
	DeploymentId *string `locationName:"deploymentId" type:"string"`

	// The execution ID of a deployment's lifecycle hook. A deployment lifecycle
	// hook is specified in the hooks section of the AppSpec file.
	LifecycleEventHookExecutionId *string `locationName:"lifecycleEventHookExecutionId" type:"string"`

	// The result of a Lambda function that validates a deployment lifecycle event
	// (Succeeded or Failed).
	Status LifecycleEventStatus `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s PutLifecycleEventHookExecutionStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/PutLifecycleEventHookExecutionStatusOutput
type PutLifecycleEventHookExecutionStatusOutput struct {
	_ struct{} `type:"structure"`

	// The execution ID of the lifecycle event hook. A hook is specified in the
	// hooks section of the deployment's AppSpec file.
	LifecycleEventHookExecutionId *string `locationName:"lifecycleEventHookExecutionId" type:"string"`
}

// String returns the string representation
func (s PutLifecycleEventHookExecutionStatusOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutLifecycleEventHookExecutionStatus = "PutLifecycleEventHookExecutionStatus"

// PutLifecycleEventHookExecutionStatusRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Sets the result of a Lambda validation function. The function validates one
// or both lifecycle events (BeforeAllowTraffic and AfterAllowTraffic) and returns
// Succeeded or Failed.
//
//    // Example sending a request using PutLifecycleEventHookExecutionStatusRequest.
//    req := client.PutLifecycleEventHookExecutionStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/PutLifecycleEventHookExecutionStatus
func (c *Client) PutLifecycleEventHookExecutionStatusRequest(input *PutLifecycleEventHookExecutionStatusInput) PutLifecycleEventHookExecutionStatusRequest {
	op := &aws.Operation{
		Name:       opPutLifecycleEventHookExecutionStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutLifecycleEventHookExecutionStatusInput{}
	}

	req := c.newRequest(op, input, &PutLifecycleEventHookExecutionStatusOutput{})
	return PutLifecycleEventHookExecutionStatusRequest{Request: req, Input: input, Copy: c.PutLifecycleEventHookExecutionStatusRequest}
}

// PutLifecycleEventHookExecutionStatusRequest is the request type for the
// PutLifecycleEventHookExecutionStatus API operation.
type PutLifecycleEventHookExecutionStatusRequest struct {
	*aws.Request
	Input *PutLifecycleEventHookExecutionStatusInput
	Copy  func(*PutLifecycleEventHookExecutionStatusInput) PutLifecycleEventHookExecutionStatusRequest
}

// Send marshals and sends the PutLifecycleEventHookExecutionStatus API request.
func (r PutLifecycleEventHookExecutionStatusRequest) Send(ctx context.Context) (*PutLifecycleEventHookExecutionStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutLifecycleEventHookExecutionStatusResponse{
		PutLifecycleEventHookExecutionStatusOutput: r.Request.Data.(*PutLifecycleEventHookExecutionStatusOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutLifecycleEventHookExecutionStatusResponse is the response type for the
// PutLifecycleEventHookExecutionStatus API operation.
type PutLifecycleEventHookExecutionStatusResponse struct {
	*PutLifecycleEventHookExecutionStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutLifecycleEventHookExecutionStatus request.
func (r *PutLifecycleEventHookExecutionStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
