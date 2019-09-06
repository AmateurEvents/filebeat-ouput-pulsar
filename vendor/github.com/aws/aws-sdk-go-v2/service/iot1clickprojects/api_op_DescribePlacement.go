// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot1clickprojects

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot1click-projects-2018-05-14/DescribePlacementRequest
type DescribePlacementInput struct {
	_ struct{} `type:"structure"`

	// The name of the placement within a project.
	//
	// PlacementName is a required field
	PlacementName *string `location:"uri" locationName:"placementName" min:"1" type:"string" required:"true"`

	// The project containing the placement to be described.
	//
	// ProjectName is a required field
	ProjectName *string `location:"uri" locationName:"projectName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribePlacementInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribePlacementInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribePlacementInput"}

	if s.PlacementName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PlacementName"))
	}
	if s.PlacementName != nil && len(*s.PlacementName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PlacementName", 1))
	}

	if s.ProjectName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectName"))
	}
	if s.ProjectName != nil && len(*s.ProjectName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProjectName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribePlacementInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.PlacementName != nil {
		v := *s.PlacementName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "placementName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ProjectName != nil {
		v := *s.ProjectName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "projectName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot1click-projects-2018-05-14/DescribePlacementResponse
type DescribePlacementOutput struct {
	_ struct{} `type:"structure"`

	// An object describing the placement.
	//
	// Placement is a required field
	Placement *PlacementDescription `locationName:"placement" type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribePlacementOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribePlacementOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Placement != nil {
		v := s.Placement

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "placement", v, metadata)
	}
	return nil
}

const opDescribePlacement = "DescribePlacement"

// DescribePlacementRequest returns a request value for making API operation for
// AWS IoT 1-Click Projects Service.
//
// Describes a placement in a project.
//
//    // Example sending a request using DescribePlacementRequest.
//    req := client.DescribePlacementRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot1click-projects-2018-05-14/DescribePlacement
func (c *Client) DescribePlacementRequest(input *DescribePlacementInput) DescribePlacementRequest {
	op := &aws.Operation{
		Name:       opDescribePlacement,
		HTTPMethod: "GET",
		HTTPPath:   "/projects/{projectName}/placements/{placementName}",
	}

	if input == nil {
		input = &DescribePlacementInput{}
	}

	req := c.newRequest(op, input, &DescribePlacementOutput{})
	return DescribePlacementRequest{Request: req, Input: input, Copy: c.DescribePlacementRequest}
}

// DescribePlacementRequest is the request type for the
// DescribePlacement API operation.
type DescribePlacementRequest struct {
	*aws.Request
	Input *DescribePlacementInput
	Copy  func(*DescribePlacementInput) DescribePlacementRequest
}

// Send marshals and sends the DescribePlacement API request.
func (r DescribePlacementRequest) Send(ctx context.Context) (*DescribePlacementResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribePlacementResponse{
		DescribePlacementOutput: r.Request.Data.(*DescribePlacementOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribePlacementResponse is the response type for the
// DescribePlacement API operation.
type DescribePlacementResponse struct {
	*DescribePlacementOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribePlacement request.
func (r *DescribePlacementResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
