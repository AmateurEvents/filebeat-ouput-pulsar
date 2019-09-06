// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package robomaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/DescribeFleetRequest
type DescribeFleetInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the fleet.
	//
	// Fleet is a required field
	Fleet *string `locationName:"fleet" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeFleetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeFleetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeFleetInput"}

	if s.Fleet == nil {
		invalidParams.Add(aws.NewErrParamRequired("Fleet"))
	}
	if s.Fleet != nil && len(*s.Fleet) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Fleet", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeFleetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Fleet != nil {
		v := *s.Fleet

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "fleet", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/DescribeFleetResponse
type DescribeFleetOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the fleet.
	Arn *string `locationName:"arn" min:"1" type:"string"`

	// The time, in milliseconds since the epoch, when the fleet was created.
	CreatedAt *time.Time `locationName:"createdAt" type:"timestamp"`

	// The Amazon Resource Name (ARN) of the last deployment job.
	LastDeploymentJob *string `locationName:"lastDeploymentJob" min:"1" type:"string"`

	// The status of the last deployment.
	LastDeploymentStatus DeploymentStatus `locationName:"lastDeploymentStatus" type:"string" enum:"true"`

	// The time of the last deployment.
	LastDeploymentTime *time.Time `locationName:"lastDeploymentTime" type:"timestamp"`

	// The name of the fleet.
	Name *string `locationName:"name" min:"1" type:"string"`

	// A list of robots.
	Robots []Robot `locationName:"robots" type:"list"`

	// The list of all tags added to the specified fleet.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s DescribeFleetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeFleetOutput) MarshalFields(e protocol.FieldEncoder) error {
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
	if s.LastDeploymentJob != nil {
		v := *s.LastDeploymentJob

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastDeploymentJob", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.LastDeploymentStatus) > 0 {
		v := s.LastDeploymentStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastDeploymentStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.LastDeploymentTime != nil {
		v := *s.LastDeploymentTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastDeploymentTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Robots != nil {
		v := s.Robots

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "robots", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

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

const opDescribeFleet = "DescribeFleet"

// DescribeFleetRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Describes a fleet.
//
//    // Example sending a request using DescribeFleetRequest.
//    req := client.DescribeFleetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/DescribeFleet
func (c *Client) DescribeFleetRequest(input *DescribeFleetInput) DescribeFleetRequest {
	op := &aws.Operation{
		Name:       opDescribeFleet,
		HTTPMethod: "POST",
		HTTPPath:   "/describeFleet",
	}

	if input == nil {
		input = &DescribeFleetInput{}
	}

	req := c.newRequest(op, input, &DescribeFleetOutput{})
	return DescribeFleetRequest{Request: req, Input: input, Copy: c.DescribeFleetRequest}
}

// DescribeFleetRequest is the request type for the
// DescribeFleet API operation.
type DescribeFleetRequest struct {
	*aws.Request
	Input *DescribeFleetInput
	Copy  func(*DescribeFleetInput) DescribeFleetRequest
}

// Send marshals and sends the DescribeFleet API request.
func (r DescribeFleetRequest) Send(ctx context.Context) (*DescribeFleetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeFleetResponse{
		DescribeFleetOutput: r.Request.Data.(*DescribeFleetOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeFleetResponse is the response type for the
// DescribeFleet API operation.
type DescribeFleetResponse struct {
	*DescribeFleetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeFleet request.
func (r *DescribeFleetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
