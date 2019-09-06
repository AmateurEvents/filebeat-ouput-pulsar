// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/AddTagsInput
type AddTagsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the resource that you want to tag.
	//
	// ResourceArn is a required field
	ResourceArn *string `type:"string" required:"true"`

	// An array of Tag objects. Each tag is a key-value pair. Only the key parameter
	// is required. If you don't specify a value, Amazon SageMaker sets the value
	// to an empty string.
	//
	// Tags is a required field
	Tags []Tag `type:"list" required:"true"`
}

// String returns the string representation
func (s AddTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddTagsInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/AddTagsOutput
type AddTagsOutput struct {
	_ struct{} `type:"structure"`

	// A list of tags associated with the Amazon SageMaker resource.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s AddTagsOutput) String() string {
	return awsutil.Prettify(s)
}

const opAddTags = "AddTags"

// AddTagsRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Adds or overwrites one or more tags for the specified Amazon SageMaker resource.
// You can add tags to notebook instances, training jobs, hyperparameter tuning
// jobs, batch transform jobs, models, labeling jobs, work teams, endpoint configurations,
// and endpoints.
//
// Each tag consists of a key and an optional value. Tag keys must be unique
// per resource. For more information about tags, see For more information,
// see AWS Tagging Strategies (https://aws.amazon.com/answers/account-management/aws-tagging-strategies/).
//
// Tags that you add to a hyperparameter tuning job by calling this API are
// also added to any training jobs that the hyperparameter tuning job launches
// after you call this API, but not to training jobs that the hyperparameter
// tuning job launched before you called this API. To make sure that the tags
// associated with a hyperparameter tuning job are also added to all training
// jobs that the hyperparameter tuning job launches, add the tags when you first
// create the tuning job by specifying them in the Tags parameter of CreateHyperParameterTuningJob
//
//    // Example sending a request using AddTagsRequest.
//    req := client.AddTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/AddTags
func (c *Client) AddTagsRequest(input *AddTagsInput) AddTagsRequest {
	op := &aws.Operation{
		Name:       opAddTags,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddTagsInput{}
	}

	req := c.newRequest(op, input, &AddTagsOutput{})
	return AddTagsRequest{Request: req, Input: input, Copy: c.AddTagsRequest}
}

// AddTagsRequest is the request type for the
// AddTags API operation.
type AddTagsRequest struct {
	*aws.Request
	Input *AddTagsInput
	Copy  func(*AddTagsInput) AddTagsRequest
}

// Send marshals and sends the AddTags API request.
func (r AddTagsRequest) Send(ctx context.Context) (*AddTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddTagsResponse{
		AddTagsOutput: r.Request.Data.(*AddTagsOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddTagsResponse is the response type for the
// AddTags API operation.
type AddTagsResponse struct {
	*AddTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddTags request.
func (r *AddTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
