// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DescribeCompilationJobRequest
type DescribeCompilationJobInput struct {
	_ struct{} `type:"structure"`

	// The name of the model compilation job that you want information about.
	//
	// CompilationJobName is a required field
	CompilationJobName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeCompilationJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeCompilationJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeCompilationJobInput"}

	if s.CompilationJobName == nil {
		invalidParams.Add(aws.NewErrParamRequired("CompilationJobName"))
	}
	if s.CompilationJobName != nil && len(*s.CompilationJobName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CompilationJobName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DescribeCompilationJobResponse
type DescribeCompilationJobOutput struct {
	_ struct{} `type:"structure"`

	// The time when the model compilation job on a compilation job instance ended.
	// For a successful or stopped job, this is when the job's model artifacts have
	// finished uploading. For a failed job, this is when Amazon SageMaker detected
	// that the job failed.
	CompilationEndTime *time.Time `type:"timestamp"`

	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker assumes
	// to perform the model compilation job.
	//
	// CompilationJobArn is a required field
	CompilationJobArn *string `type:"string" required:"true"`

	// The name of the model compilation job.
	//
	// CompilationJobName is a required field
	CompilationJobName *string `min:"1" type:"string" required:"true"`

	// The status of the model compilation job.
	//
	// CompilationJobStatus is a required field
	CompilationJobStatus CompilationJobStatus `type:"string" required:"true" enum:"true"`

	// The time when the model compilation job started the CompilationJob instances.
	//
	// You are billed for the time between this timestamp and the timestamp in the
	// DescribeCompilationJobResponse$CompilationEndTime field. In Amazon CloudWatch
	// Logs, the start time might be later than this time. That's because it takes
	// time to download the compilation job, which depends on the size of the compilation
	// job container.
	CompilationStartTime *time.Time `type:"timestamp"`

	// The time that the model compilation job was created.
	//
	// CreationTime is a required field
	CreationTime *time.Time `type:"timestamp" required:"true"`

	// If a model compilation job failed, the reason it failed.
	//
	// FailureReason is a required field
	FailureReason *string `type:"string" required:"true"`

	// Information about the location in Amazon S3 of the input model artifacts,
	// the name and shape of the expected data inputs, and the framework in which
	// the model was trained.
	//
	// InputConfig is a required field
	InputConfig *InputConfig `type:"structure" required:"true"`

	// The time that the status of the model compilation job was last modified.
	//
	// LastModifiedTime is a required field
	LastModifiedTime *time.Time `type:"timestamp" required:"true"`

	// Information about the location in Amazon S3 that has been configured for
	// storing the model artifacts used in the compilation job.
	//
	// ModelArtifacts is a required field
	ModelArtifacts *ModelArtifacts `type:"structure" required:"true"`

	// Information about the output location for the compiled model and the target
	// device that the model runs on.
	//
	// OutputConfig is a required field
	OutputConfig *OutputConfig `type:"structure" required:"true"`

	// The Amazon Resource Name (ARN) of the model compilation job.
	//
	// RoleArn is a required field
	RoleArn *string `min:"20" type:"string" required:"true"`

	// Specifies a limit to how long a model compilation job can run. When the job
	// reaches the time limit, Amazon SageMaker ends the compilation job. Use this
	// API to cap model training costs.
	//
	// StoppingCondition is a required field
	StoppingCondition *StoppingCondition `type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeCompilationJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeCompilationJob = "DescribeCompilationJob"

// DescribeCompilationJobRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Returns information about a model compilation job.
//
// To create a model compilation job, use CreateCompilationJob. To get information
// about multiple model compilation jobs, use ListCompilationJobs.
//
//    // Example sending a request using DescribeCompilationJobRequest.
//    req := client.DescribeCompilationJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DescribeCompilationJob
func (c *Client) DescribeCompilationJobRequest(input *DescribeCompilationJobInput) DescribeCompilationJobRequest {
	op := &aws.Operation{
		Name:       opDescribeCompilationJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeCompilationJobInput{}
	}

	req := c.newRequest(op, input, &DescribeCompilationJobOutput{})
	return DescribeCompilationJobRequest{Request: req, Input: input, Copy: c.DescribeCompilationJobRequest}
}

// DescribeCompilationJobRequest is the request type for the
// DescribeCompilationJob API operation.
type DescribeCompilationJobRequest struct {
	*aws.Request
	Input *DescribeCompilationJobInput
	Copy  func(*DescribeCompilationJobInput) DescribeCompilationJobRequest
}

// Send marshals and sends the DescribeCompilationJob API request.
func (r DescribeCompilationJobRequest) Send(ctx context.Context) (*DescribeCompilationJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCompilationJobResponse{
		DescribeCompilationJobOutput: r.Request.Data.(*DescribeCompilationJobOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeCompilationJobResponse is the response type for the
// DescribeCompilationJob API operation.
type DescribeCompilationJobResponse struct {
	*DescribeCompilationJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCompilationJob request.
func (r *DescribeCompilationJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
