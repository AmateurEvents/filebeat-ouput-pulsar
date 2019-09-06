// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package inspector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/StartAssessmentRunRequest
type StartAssessmentRunInput struct {
	_ struct{} `type:"structure"`

	// You can specify the name for the assessment run. The name must be unique
	// for the assessment template whose ARN is used to start the assessment run.
	AssessmentRunName *string `locationName:"assessmentRunName" min:"1" type:"string"`

	// The ARN of the assessment template of the assessment run that you want to
	// start.
	//
	// AssessmentTemplateArn is a required field
	AssessmentTemplateArn *string `locationName:"assessmentTemplateArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StartAssessmentRunInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartAssessmentRunInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartAssessmentRunInput"}
	if s.AssessmentRunName != nil && len(*s.AssessmentRunName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AssessmentRunName", 1))
	}

	if s.AssessmentTemplateArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssessmentTemplateArn"))
	}
	if s.AssessmentTemplateArn != nil && len(*s.AssessmentTemplateArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AssessmentTemplateArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/StartAssessmentRunResponse
type StartAssessmentRunOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the assessment run that has been started.
	//
	// AssessmentRunArn is a required field
	AssessmentRunArn *string `locationName:"assessmentRunArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StartAssessmentRunOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartAssessmentRun = "StartAssessmentRun"

// StartAssessmentRunRequest returns a request value for making API operation for
// Amazon Inspector.
//
// Starts the assessment run specified by the ARN of the assessment template.
// For this API to function properly, you must not exceed the limit of running
// up to 500 concurrent agents per AWS account.
//
//    // Example sending a request using StartAssessmentRunRequest.
//    req := client.StartAssessmentRunRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/StartAssessmentRun
func (c *Client) StartAssessmentRunRequest(input *StartAssessmentRunInput) StartAssessmentRunRequest {
	op := &aws.Operation{
		Name:       opStartAssessmentRun,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartAssessmentRunInput{}
	}

	req := c.newRequest(op, input, &StartAssessmentRunOutput{})
	return StartAssessmentRunRequest{Request: req, Input: input, Copy: c.StartAssessmentRunRequest}
}

// StartAssessmentRunRequest is the request type for the
// StartAssessmentRun API operation.
type StartAssessmentRunRequest struct {
	*aws.Request
	Input *StartAssessmentRunInput
	Copy  func(*StartAssessmentRunInput) StartAssessmentRunRequest
}

// Send marshals and sends the StartAssessmentRun API request.
func (r StartAssessmentRunRequest) Send(ctx context.Context) (*StartAssessmentRunResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartAssessmentRunResponse{
		StartAssessmentRunOutput: r.Request.Data.(*StartAssessmentRunOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartAssessmentRunResponse is the response type for the
// StartAssessmentRun API operation.
type StartAssessmentRunResponse struct {
	*StartAssessmentRunOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartAssessmentRun request.
func (r *StartAssessmentRunResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
