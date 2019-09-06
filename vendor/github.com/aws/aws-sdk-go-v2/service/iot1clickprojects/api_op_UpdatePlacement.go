// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot1clickprojects

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot1click-projects-2018-05-14/UpdatePlacementRequest
type UpdatePlacementInput struct {
	_ struct{} `type:"structure"`

	// The user-defined object of attributes used to update the placement. The maximum
	// number of key/value pairs is 50.
	Attributes map[string]string `locationName:"attributes" type:"map"`

	// The name of the placement to update.
	//
	// PlacementName is a required field
	PlacementName *string `location:"uri" locationName:"placementName" min:"1" type:"string" required:"true"`

	// The name of the project containing the placement to be updated.
	//
	// ProjectName is a required field
	ProjectName *string `location:"uri" locationName:"projectName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdatePlacementInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdatePlacementInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdatePlacementInput"}

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
func (s UpdatePlacementInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Attributes != nil {
		v := s.Attributes

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "attributes", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot1click-projects-2018-05-14/UpdatePlacementResponse
type UpdatePlacementOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdatePlacementOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdatePlacementOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdatePlacement = "UpdatePlacement"

// UpdatePlacementRequest returns a request value for making API operation for
// AWS IoT 1-Click Projects Service.
//
// Updates a placement with the given attributes. To clear an attribute, pass
// an empty value (i.e., "").
//
//    // Example sending a request using UpdatePlacementRequest.
//    req := client.UpdatePlacementRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot1click-projects-2018-05-14/UpdatePlacement
func (c *Client) UpdatePlacementRequest(input *UpdatePlacementInput) UpdatePlacementRequest {
	op := &aws.Operation{
		Name:       opUpdatePlacement,
		HTTPMethod: "PUT",
		HTTPPath:   "/projects/{projectName}/placements/{placementName}",
	}

	if input == nil {
		input = &UpdatePlacementInput{}
	}

	req := c.newRequest(op, input, &UpdatePlacementOutput{})
	return UpdatePlacementRequest{Request: req, Input: input, Copy: c.UpdatePlacementRequest}
}

// UpdatePlacementRequest is the request type for the
// UpdatePlacement API operation.
type UpdatePlacementRequest struct {
	*aws.Request
	Input *UpdatePlacementInput
	Copy  func(*UpdatePlacementInput) UpdatePlacementRequest
}

// Send marshals and sends the UpdatePlacement API request.
func (r UpdatePlacementRequest) Send(ctx context.Context) (*UpdatePlacementResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdatePlacementResponse{
		UpdatePlacementOutput: r.Request.Data.(*UpdatePlacementOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdatePlacementResponse is the response type for the
// UpdatePlacement API operation.
type UpdatePlacementResponse struct {
	*UpdatePlacementOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdatePlacement request.
func (r *UpdatePlacementResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
