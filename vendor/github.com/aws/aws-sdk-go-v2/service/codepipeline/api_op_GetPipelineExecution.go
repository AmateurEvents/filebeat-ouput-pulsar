// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codepipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a GetPipelineExecution action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/GetPipelineExecutionInput
type GetPipelineExecutionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the pipeline execution about which you want to get execution details.
	//
	// PipelineExecutionId is a required field
	PipelineExecutionId *string `locationName:"pipelineExecutionId" type:"string" required:"true"`

	// The name of the pipeline about which you want to get execution details.
	//
	// PipelineName is a required field
	PipelineName *string `locationName:"pipelineName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetPipelineExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPipelineExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPipelineExecutionInput"}

	if s.PipelineExecutionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PipelineExecutionId"))
	}

	if s.PipelineName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PipelineName"))
	}
	if s.PipelineName != nil && len(*s.PipelineName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PipelineName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a GetPipelineExecution action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/GetPipelineExecutionOutput
type GetPipelineExecutionOutput struct {
	_ struct{} `type:"structure"`

	// Represents information about the execution of a pipeline.
	PipelineExecution *PipelineExecution `locationName:"pipelineExecution" type:"structure"`
}

// String returns the string representation
func (s GetPipelineExecutionOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetPipelineExecution = "GetPipelineExecution"

// GetPipelineExecutionRequest returns a request value for making API operation for
// AWS CodePipeline.
//
// Returns information about an execution of a pipeline, including details about
// artifacts, the pipeline execution ID, and the name, version, and status of
// the pipeline.
//
//    // Example sending a request using GetPipelineExecutionRequest.
//    req := client.GetPipelineExecutionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/GetPipelineExecution
func (c *Client) GetPipelineExecutionRequest(input *GetPipelineExecutionInput) GetPipelineExecutionRequest {
	op := &aws.Operation{
		Name:       opGetPipelineExecution,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetPipelineExecutionInput{}
	}

	req := c.newRequest(op, input, &GetPipelineExecutionOutput{})
	return GetPipelineExecutionRequest{Request: req, Input: input, Copy: c.GetPipelineExecutionRequest}
}

// GetPipelineExecutionRequest is the request type for the
// GetPipelineExecution API operation.
type GetPipelineExecutionRequest struct {
	*aws.Request
	Input *GetPipelineExecutionInput
	Copy  func(*GetPipelineExecutionInput) GetPipelineExecutionRequest
}

// Send marshals and sends the GetPipelineExecution API request.
func (r GetPipelineExecutionRequest) Send(ctx context.Context) (*GetPipelineExecutionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetPipelineExecutionResponse{
		GetPipelineExecutionOutput: r.Request.Data.(*GetPipelineExecutionOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetPipelineExecutionResponse is the response type for the
// GetPipelineExecution API operation.
type GetPipelineExecutionResponse struct {
	*GetPipelineExecutionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetPipelineExecution request.
func (r *GetPipelineExecutionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
