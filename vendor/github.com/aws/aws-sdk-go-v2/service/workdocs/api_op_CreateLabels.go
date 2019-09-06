// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/CreateLabelsRequest
type CreateLabelsInput struct {
	_ struct{} `type:"structure"`

	// Amazon WorkDocs authentication token. Do not set this field when using administrative
	// API actions, as in accessing the API using AWS credentials.
	AuthenticationToken *string `location:"header" locationName:"Authentication" min:"1" type:"string"`

	// List of labels to add to the resource.
	//
	// Labels is a required field
	Labels []string `type:"list" required:"true"`

	// The ID of the resource.
	//
	// ResourceId is a required field
	ResourceId *string `location:"uri" locationName:"ResourceId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateLabelsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateLabelsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateLabelsInput"}
	if s.AuthenticationToken != nil && len(*s.AuthenticationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthenticationToken", 1))
	}

	if s.Labels == nil {
		invalidParams.Add(aws.NewErrParamRequired("Labels"))
	}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}
	if s.ResourceId != nil && len(*s.ResourceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateLabelsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Labels != nil {
		v := s.Labels

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Labels", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.AuthenticationToken != nil {
		v := *s.AuthenticationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Authentication", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceId != nil {
		v := *s.ResourceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ResourceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/CreateLabelsResponse
type CreateLabelsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateLabelsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateLabelsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opCreateLabels = "CreateLabels"

// CreateLabelsRequest returns a request value for making API operation for
// Amazon WorkDocs.
//
// Adds the specified list of labels to the given resource (a document or folder)
//
//    // Example sending a request using CreateLabelsRequest.
//    req := client.CreateLabelsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/CreateLabels
func (c *Client) CreateLabelsRequest(input *CreateLabelsInput) CreateLabelsRequest {
	op := &aws.Operation{
		Name:       opCreateLabels,
		HTTPMethod: "PUT",
		HTTPPath:   "/api/v1/resources/{ResourceId}/labels",
	}

	if input == nil {
		input = &CreateLabelsInput{}
	}

	req := c.newRequest(op, input, &CreateLabelsOutput{})
	return CreateLabelsRequest{Request: req, Input: input, Copy: c.CreateLabelsRequest}
}

// CreateLabelsRequest is the request type for the
// CreateLabels API operation.
type CreateLabelsRequest struct {
	*aws.Request
	Input *CreateLabelsInput
	Copy  func(*CreateLabelsInput) CreateLabelsRequest
}

// Send marshals and sends the CreateLabels API request.
func (r CreateLabelsRequest) Send(ctx context.Context) (*CreateLabelsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateLabelsResponse{
		CreateLabelsOutput: r.Request.Data.(*CreateLabelsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateLabelsResponse is the response type for the
// CreateLabels API operation.
type CreateLabelsResponse struct {
	*CreateLabelsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateLabels request.
func (r *CreateLabelsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
