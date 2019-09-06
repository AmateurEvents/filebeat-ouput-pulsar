// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/CreateWebACLRequest
type CreateWebACLInput struct {
	_ struct{} `type:"structure"`

	// The value returned by the most recent call to GetChangeToken.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`

	// The action that you want AWS WAF to take when a request doesn't match the
	// criteria specified in any of the Rule objects that are associated with the
	// WebACL.
	//
	// DefaultAction is a required field
	DefaultAction *WafAction `type:"structure" required:"true"`

	// A friendly name or description for the metrics for this WebACL.The name can
	// contain only alphanumeric characters (A-Z, a-z, 0-9), with maximum length
	// 128 and minimum length one. It can't contain whitespace or metric names reserved
	// for AWS WAF, including "All" and "Default_Action." You can't change MetricName
	// after you create the WebACL.
	//
	// MetricName is a required field
	MetricName *string `type:"string" required:"true"`

	// A friendly name or description of the WebACL. You can't change Name after
	// you create the WebACL.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	Tags []Tag `min:"1" type:"list"`
}

// String returns the string representation
func (s CreateWebACLInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateWebACLInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateWebACLInput"}

	if s.ChangeToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeToken"))
	}
	if s.ChangeToken != nil && len(*s.ChangeToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeToken", 1))
	}

	if s.DefaultAction == nil {
		invalidParams.Add(aws.NewErrParamRequired("DefaultAction"))
	}

	if s.MetricName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MetricName"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.DefaultAction != nil {
		if err := s.DefaultAction.Validate(); err != nil {
			invalidParams.AddNested("DefaultAction", err.(aws.ErrInvalidParams))
		}
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/CreateWebACLResponse
type CreateWebACLOutput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used to submit the CreateWebACL request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus.
	ChangeToken *string `min:"1" type:"string"`

	// The WebACL returned in the CreateWebACL response.
	WebACL *WebACL `type:"structure"`
}

// String returns the string representation
func (s CreateWebACLOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateWebACL = "CreateWebACL"

// CreateWebACLRequest returns a request value for making API operation for
// AWS WAF.
//
// Creates a WebACL, which contains the Rules that identify the CloudFront web
// requests that you want to allow, block, or count. AWS WAF evaluates Rules
// in order based on the value of Priority for each Rule.
//
// You also specify a default action, either ALLOW or BLOCK. If a web request
// doesn't match any of the Rules in a WebACL, AWS WAF responds to the request
// with the default action.
//
// To create and configure a WebACL, perform the following steps:
//
// Create and update the ByteMatchSet objects and other predicates that you
// want to include in Rules. For more information, see CreateByteMatchSet, UpdateByteMatchSet,
// CreateIPSet, UpdateIPSet, CreateSqlInjectionMatchSet, and UpdateSqlInjectionMatchSet.
//
// Create and update the Rules that you want to include in the WebACL. For more
// information, see CreateRule and UpdateRule.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of a CreateWebACL request.
//
// Submit a CreateWebACL request.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of an UpdateWebACL request.
//
// Submit an UpdateWebACL request to specify the Rules that you want to include
// in the WebACL, to specify the default action, and to associate the WebACL
// with a CloudFront distribution.
//
// For more information about how to use the AWS WAF API, see the AWS WAF Developer
// Guide (https://docs.aws.amazon.com/waf/latest/developerguide/).
//
//    // Example sending a request using CreateWebACLRequest.
//    req := client.CreateWebACLRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/CreateWebACL
func (c *Client) CreateWebACLRequest(input *CreateWebACLInput) CreateWebACLRequest {
	op := &aws.Operation{
		Name:       opCreateWebACL,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateWebACLInput{}
	}

	req := c.newRequest(op, input, &CreateWebACLOutput{})
	return CreateWebACLRequest{Request: req, Input: input, Copy: c.CreateWebACLRequest}
}

// CreateWebACLRequest is the request type for the
// CreateWebACL API operation.
type CreateWebACLRequest struct {
	*aws.Request
	Input *CreateWebACLInput
	Copy  func(*CreateWebACLInput) CreateWebACLRequest
}

// Send marshals and sends the CreateWebACL API request.
func (r CreateWebACLRequest) Send(ctx context.Context) (*CreateWebACLResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateWebACLResponse{
		CreateWebACLOutput: r.Request.Data.(*CreateWebACLOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateWebACLResponse is the response type for the
// CreateWebACL API operation.
type CreateWebACLResponse struct {
	*CreateWebACLOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateWebACL request.
func (r *CreateWebACLResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
