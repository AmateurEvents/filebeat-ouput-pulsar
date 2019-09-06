// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListFunctionDefinitionVersionsRequest
type ListFunctionDefinitionVersionsInput struct {
	_ struct{} `type:"structure"`

	// FunctionDefinitionId is a required field
	FunctionDefinitionId *string `location:"uri" locationName:"FunctionDefinitionId" type:"string" required:"true"`

	MaxResults *string `location:"querystring" locationName:"MaxResults" type:"string"`

	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`
}

// String returns the string representation
func (s ListFunctionDefinitionVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListFunctionDefinitionVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListFunctionDefinitionVersionsInput"}

	if s.FunctionDefinitionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FunctionDefinitionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFunctionDefinitionVersionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.FunctionDefinitionId != nil {
		v := *s.FunctionDefinitionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FunctionDefinitionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MaxResults", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListFunctionDefinitionVersionsResponse
type ListFunctionDefinitionVersionsOutput struct {
	_ struct{} `type:"structure"`

	NextToken *string `type:"string"`

	Versions []VersionInformation `type:"list"`
}

// String returns the string representation
func (s ListFunctionDefinitionVersionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFunctionDefinitionVersionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Versions != nil {
		v := s.Versions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Versions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListFunctionDefinitionVersions = "ListFunctionDefinitionVersions"

// ListFunctionDefinitionVersionsRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Lists the versions of a Lambda function definition.
//
//    // Example sending a request using ListFunctionDefinitionVersionsRequest.
//    req := client.ListFunctionDefinitionVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListFunctionDefinitionVersions
func (c *Client) ListFunctionDefinitionVersionsRequest(input *ListFunctionDefinitionVersionsInput) ListFunctionDefinitionVersionsRequest {
	op := &aws.Operation{
		Name:       opListFunctionDefinitionVersions,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/definition/functions/{FunctionDefinitionId}/versions",
	}

	if input == nil {
		input = &ListFunctionDefinitionVersionsInput{}
	}

	req := c.newRequest(op, input, &ListFunctionDefinitionVersionsOutput{})
	return ListFunctionDefinitionVersionsRequest{Request: req, Input: input, Copy: c.ListFunctionDefinitionVersionsRequest}
}

// ListFunctionDefinitionVersionsRequest is the request type for the
// ListFunctionDefinitionVersions API operation.
type ListFunctionDefinitionVersionsRequest struct {
	*aws.Request
	Input *ListFunctionDefinitionVersionsInput
	Copy  func(*ListFunctionDefinitionVersionsInput) ListFunctionDefinitionVersionsRequest
}

// Send marshals and sends the ListFunctionDefinitionVersions API request.
func (r ListFunctionDefinitionVersionsRequest) Send(ctx context.Context) (*ListFunctionDefinitionVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListFunctionDefinitionVersionsResponse{
		ListFunctionDefinitionVersionsOutput: r.Request.Data.(*ListFunctionDefinitionVersionsOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListFunctionDefinitionVersionsResponse is the response type for the
// ListFunctionDefinitionVersions API operation.
type ListFunctionDefinitionVersionsResponse struct {
	*ListFunctionDefinitionVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListFunctionDefinitionVersions request.
func (r *ListFunctionDefinitionVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
