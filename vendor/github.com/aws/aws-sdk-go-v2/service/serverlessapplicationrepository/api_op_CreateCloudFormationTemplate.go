// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package serverlessapplicationrepository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/CreateCloudFormationTemplateRequest
type CreateCloudFormationTemplateInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"applicationId" type:"string" required:"true"`

	SemanticVersion *string `locationName:"semanticVersion" type:"string"`
}

// String returns the string representation
func (s CreateCloudFormationTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateCloudFormationTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateCloudFormationTemplateInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateCloudFormationTemplateInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.SemanticVersion != nil {
		v := *s.SemanticVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "semanticVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "applicationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/CreateCloudFormationTemplateResponse
type CreateCloudFormationTemplateOutput struct {
	_ struct{} `type:"structure"`

	ApplicationId *string `locationName:"applicationId" type:"string"`

	CreationTime *string `locationName:"creationTime" type:"string"`

	ExpirationTime *string `locationName:"expirationTime" type:"string"`

	SemanticVersion *string `locationName:"semanticVersion" type:"string"`

	Status Status `locationName:"status" type:"string" enum:"true"`

	TemplateId *string `locationName:"templateId" type:"string"`

	TemplateUrl *string `locationName:"templateUrl" type:"string"`
}

// String returns the string representation
func (s CreateCloudFormationTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateCloudFormationTemplateOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "applicationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreationTime != nil {
		v := *s.CreationTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "creationTime", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ExpirationTime != nil {
		v := *s.ExpirationTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "expirationTime", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SemanticVersion != nil {
		v := *s.SemanticVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "semanticVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.TemplateId != nil {
		v := *s.TemplateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "templateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TemplateUrl != nil {
		v := *s.TemplateUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "templateUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateCloudFormationTemplate = "CreateCloudFormationTemplate"

// CreateCloudFormationTemplateRequest returns a request value for making API operation for
// AWSServerlessApplicationRepository.
//
// Creates an AWS CloudFormation template.
//
//    // Example sending a request using CreateCloudFormationTemplateRequest.
//    req := client.CreateCloudFormationTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/CreateCloudFormationTemplate
func (c *Client) CreateCloudFormationTemplateRequest(input *CreateCloudFormationTemplateInput) CreateCloudFormationTemplateRequest {
	op := &aws.Operation{
		Name:       opCreateCloudFormationTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/applications/{applicationId}/templates",
	}

	if input == nil {
		input = &CreateCloudFormationTemplateInput{}
	}

	req := c.newRequest(op, input, &CreateCloudFormationTemplateOutput{})
	return CreateCloudFormationTemplateRequest{Request: req, Input: input, Copy: c.CreateCloudFormationTemplateRequest}
}

// CreateCloudFormationTemplateRequest is the request type for the
// CreateCloudFormationTemplate API operation.
type CreateCloudFormationTemplateRequest struct {
	*aws.Request
	Input *CreateCloudFormationTemplateInput
	Copy  func(*CreateCloudFormationTemplateInput) CreateCloudFormationTemplateRequest
}

// Send marshals and sends the CreateCloudFormationTemplate API request.
func (r CreateCloudFormationTemplateRequest) Send(ctx context.Context) (*CreateCloudFormationTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateCloudFormationTemplateResponse{
		CreateCloudFormationTemplateOutput: r.Request.Data.(*CreateCloudFormationTemplateOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateCloudFormationTemplateResponse is the response type for the
// CreateCloudFormationTemplate API operation.
type CreateCloudFormationTemplateResponse struct {
	*CreateCloudFormationTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCloudFormationTemplate request.
func (r *CreateCloudFormationTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
