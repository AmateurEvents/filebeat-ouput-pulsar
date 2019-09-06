// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package inspector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/CreateAssessmentTargetRequest
type CreateAssessmentTargetInput struct {
	_ struct{} `type:"structure"`

	// The user-defined name that identifies the assessment target that you want
	// to create. The name must be unique within the AWS account.
	//
	// AssessmentTargetName is a required field
	AssessmentTargetName *string `locationName:"assessmentTargetName" min:"1" type:"string" required:"true"`

	// The ARN that specifies the resource group that is used to create the assessment
	// target. If resourceGroupArn is not specified, all EC2 instances in the current
	// AWS account and region are included in the assessment target.
	ResourceGroupArn *string `locationName:"resourceGroupArn" min:"1" type:"string"`
}

// String returns the string representation
func (s CreateAssessmentTargetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateAssessmentTargetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateAssessmentTargetInput"}

	if s.AssessmentTargetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssessmentTargetName"))
	}
	if s.AssessmentTargetName != nil && len(*s.AssessmentTargetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AssessmentTargetName", 1))
	}
	if s.ResourceGroupArn != nil && len(*s.ResourceGroupArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceGroupArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/CreateAssessmentTargetResponse
type CreateAssessmentTargetOutput struct {
	_ struct{} `type:"structure"`

	// The ARN that specifies the assessment target that is created.
	//
	// AssessmentTargetArn is a required field
	AssessmentTargetArn *string `locationName:"assessmentTargetArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateAssessmentTargetOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateAssessmentTarget = "CreateAssessmentTarget"

// CreateAssessmentTargetRequest returns a request value for making API operation for
// Amazon Inspector.
//
// Creates a new assessment target using the ARN of the resource group that
// is generated by CreateResourceGroup. If resourceGroupArn is not specified,
// all EC2 instances in the current AWS account and region are included in the
// assessment target. If the service-linked role (https://docs.aws.amazon.com/inspector/latest/userguide/inspector_slr.html)
// isn’t already registered, this action also creates and registers a service-linked
// role to grant Amazon Inspector access to AWS Services needed to perform security
// assessments. You can create up to 50 assessment targets per AWS account.
// You can run up to 500 concurrent agents per AWS account. For more information,
// see Amazon Inspector Assessment Targets (https://docs.aws.amazon.com/inspector/latest/userguide/inspector_applications.html).
//
//    // Example sending a request using CreateAssessmentTargetRequest.
//    req := client.CreateAssessmentTargetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/CreateAssessmentTarget
func (c *Client) CreateAssessmentTargetRequest(input *CreateAssessmentTargetInput) CreateAssessmentTargetRequest {
	op := &aws.Operation{
		Name:       opCreateAssessmentTarget,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateAssessmentTargetInput{}
	}

	req := c.newRequest(op, input, &CreateAssessmentTargetOutput{})
	return CreateAssessmentTargetRequest{Request: req, Input: input, Copy: c.CreateAssessmentTargetRequest}
}

// CreateAssessmentTargetRequest is the request type for the
// CreateAssessmentTarget API operation.
type CreateAssessmentTargetRequest struct {
	*aws.Request
	Input *CreateAssessmentTargetInput
	Copy  func(*CreateAssessmentTargetInput) CreateAssessmentTargetRequest
}

// Send marshals and sends the CreateAssessmentTarget API request.
func (r CreateAssessmentTargetRequest) Send(ctx context.Context) (*CreateAssessmentTargetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateAssessmentTargetResponse{
		CreateAssessmentTargetOutput: r.Request.Data.(*CreateAssessmentTargetOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateAssessmentTargetResponse is the response type for the
// CreateAssessmentTarget API operation.
type CreateAssessmentTargetResponse struct {
	*CreateAssessmentTargetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateAssessmentTarget request.
func (r *CreateAssessmentTargetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
