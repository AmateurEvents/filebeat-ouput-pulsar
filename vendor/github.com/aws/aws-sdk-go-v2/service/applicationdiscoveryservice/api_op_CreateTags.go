// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/CreateTagsRequest
type CreateTagsInput struct {
	_ struct{} `type:"structure"`

	// A list of configuration items that you want to tag.
	//
	// ConfigurationIds is a required field
	ConfigurationIds []string `locationName:"configurationIds" type:"list" required:"true"`

	// Tags that you want to associate with one or more configuration items. Specify
	// the tags that you want to create in a key-value format. For example:
	//
	// {"key": "serverType", "value": "webServer"}
	//
	// Tags is a required field
	Tags []Tag `locationName:"tags" type:"list" required:"true"`
}

// String returns the string representation
func (s CreateTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateTagsInput"}

	if s.ConfigurationIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationIds"))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/CreateTagsResponse
type CreateTagsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateTagsOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateTags = "CreateTags"

// CreateTagsRequest returns a request value for making API operation for
// AWS Application Discovery Service.
//
// Creates one or more tags for configuration items. Tags are metadata that
// help you categorize IT assets. This API accepts a list of multiple configuration
// items.
//
//    // Example sending a request using CreateTagsRequest.
//    req := client.CreateTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/CreateTags
func (c *Client) CreateTagsRequest(input *CreateTagsInput) CreateTagsRequest {
	op := &aws.Operation{
		Name:       opCreateTags,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateTagsInput{}
	}

	req := c.newRequest(op, input, &CreateTagsOutput{})
	return CreateTagsRequest{Request: req, Input: input, Copy: c.CreateTagsRequest}
}

// CreateTagsRequest is the request type for the
// CreateTags API operation.
type CreateTagsRequest struct {
	*aws.Request
	Input *CreateTagsInput
	Copy  func(*CreateTagsInput) CreateTagsRequest
}

// Send marshals and sends the CreateTags API request.
func (r CreateTagsRequest) Send(ctx context.Context) (*CreateTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTagsResponse{
		CreateTagsOutput: r.Request.Data.(*CreateTagsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTagsResponse is the response type for the
// CreateTags API operation.
type CreateTagsResponse struct {
	*CreateTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateTags request.
func (r *CreateTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
