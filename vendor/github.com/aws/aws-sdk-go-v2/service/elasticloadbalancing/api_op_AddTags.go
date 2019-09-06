// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for AddTags.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/AddTagsInput
type AddTagsInput struct {
	_ struct{} `type:"structure"`

	// The name of the load balancer. You can specify one load balancer only.
	//
	// LoadBalancerNames is a required field
	LoadBalancerNames []string `type:"list" required:"true"`

	// The tags.
	//
	// Tags is a required field
	Tags []Tag `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s AddTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddTagsInput"}

	if s.LoadBalancerNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerNames"))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
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

// Contains the output of AddTags.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/AddTagsOutput
type AddTagsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AddTagsOutput) String() string {
	return awsutil.Prettify(s)
}

const opAddTags = "AddTags"

// AddTagsRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Adds the specified tags to the specified load balancer. Each load balancer
// can have a maximum of 10 tags.
//
// Each tag consists of a key and an optional value. If a tag with the same
// key is already associated with the load balancer, AddTags updates its value.
//
// For more information, see Tag Your Classic Load Balancer (http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/add-remove-tags.html)
// in the Classic Load Balancers Guide.
//
//    // Example sending a request using AddTagsRequest.
//    req := client.AddTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/AddTags
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
