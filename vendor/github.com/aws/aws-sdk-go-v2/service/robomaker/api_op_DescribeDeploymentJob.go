// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package robomaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/DescribeDeploymentJobRequest
type DescribeDeploymentJobInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the deployment job.
	//
	// Job is a required field
	Job *string `locationName:"job" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeDeploymentJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDeploymentJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDeploymentJobInput"}

	if s.Job == nil {
		invalidParams.Add(aws.NewErrParamRequired("Job"))
	}
	if s.Job != nil && len(*s.Job) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Job", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeDeploymentJobInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Job != nil {
		v := *s.Job

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "job", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/DescribeDeploymentJobResponse
type DescribeDeploymentJobOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the deployment job.
	Arn *string `locationName:"arn" min:"1" type:"string"`

	// The time, in milliseconds since the epoch, when the deployment job was created.
	CreatedAt *time.Time `locationName:"createdAt" type:"timestamp"`

	// The deployment application configuration.
	DeploymentApplicationConfigs []DeploymentApplicationConfig `locationName:"deploymentApplicationConfigs" min:"1" type:"list"`

	// The deployment configuration.
	DeploymentConfig *DeploymentConfig `locationName:"deploymentConfig" type:"structure"`

	// The deployment job failure code.
	FailureCode DeploymentJobErrorCode `locationName:"failureCode" type:"string" enum:"true"`

	// A short description of the reason why the deployment job failed.
	FailureReason *string `locationName:"failureReason" type:"string"`

	// The Amazon Resource Name (ARN) of the fleet.
	Fleet *string `locationName:"fleet" min:"1" type:"string"`

	// A list of robot deployment summaries.
	RobotDeploymentSummary []RobotDeployment `locationName:"robotDeploymentSummary" type:"list"`

	// The status of the deployment job.
	Status DeploymentStatus `locationName:"status" type:"string" enum:"true"`

	// The list of all tags added to the specified deployment job.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s DescribeDeploymentJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeDeploymentJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreatedAt != nil {
		v := *s.CreatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdAt",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.DeploymentApplicationConfigs != nil {
		v := s.DeploymentApplicationConfigs

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "deploymentApplicationConfigs", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.DeploymentConfig != nil {
		v := s.DeploymentConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "deploymentConfig", v, metadata)
	}
	if len(s.FailureCode) > 0 {
		v := s.FailureCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "failureCode", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.FailureReason != nil {
		v := *s.FailureReason

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "failureReason", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Fleet != nil {
		v := *s.Fleet

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "fleet", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RobotDeploymentSummary != nil {
		v := s.RobotDeploymentSummary

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "robotDeploymentSummary", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

const opDescribeDeploymentJob = "DescribeDeploymentJob"

// DescribeDeploymentJobRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Describes a deployment job.
//
//    // Example sending a request using DescribeDeploymentJobRequest.
//    req := client.DescribeDeploymentJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/DescribeDeploymentJob
func (c *Client) DescribeDeploymentJobRequest(input *DescribeDeploymentJobInput) DescribeDeploymentJobRequest {
	op := &aws.Operation{
		Name:       opDescribeDeploymentJob,
		HTTPMethod: "POST",
		HTTPPath:   "/describeDeploymentJob",
	}

	if input == nil {
		input = &DescribeDeploymentJobInput{}
	}

	req := c.newRequest(op, input, &DescribeDeploymentJobOutput{})
	return DescribeDeploymentJobRequest{Request: req, Input: input, Copy: c.DescribeDeploymentJobRequest}
}

// DescribeDeploymentJobRequest is the request type for the
// DescribeDeploymentJob API operation.
type DescribeDeploymentJobRequest struct {
	*aws.Request
	Input *DescribeDeploymentJobInput
	Copy  func(*DescribeDeploymentJobInput) DescribeDeploymentJobRequest
}

// Send marshals and sends the DescribeDeploymentJob API request.
func (r DescribeDeploymentJobRequest) Send(ctx context.Context) (*DescribeDeploymentJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDeploymentJobResponse{
		DescribeDeploymentJobOutput: r.Request.Data.(*DescribeDeploymentJobOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDeploymentJobResponse is the response type for the
// DescribeDeploymentJob API operation.
type DescribeDeploymentJobResponse struct {
	*DescribeDeploymentJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDeploymentJob request.
func (r *DescribeDeploymentJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
