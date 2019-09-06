// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/BatchGetAggregateResourceConfigRequest
type BatchGetAggregateResourceConfigInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration aggregator.
	//
	// ConfigurationAggregatorName is a required field
	ConfigurationAggregatorName *string `min:"1" type:"string" required:"true"`

	// A list of aggregate ResourceIdentifiers objects.
	//
	// ResourceIdentifiers is a required field
	ResourceIdentifiers []AggregateResourceIdentifier `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchGetAggregateResourceConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchGetAggregateResourceConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchGetAggregateResourceConfigInput"}

	if s.ConfigurationAggregatorName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationAggregatorName"))
	}
	if s.ConfigurationAggregatorName != nil && len(*s.ConfigurationAggregatorName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ConfigurationAggregatorName", 1))
	}

	if s.ResourceIdentifiers == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceIdentifiers"))
	}
	if s.ResourceIdentifiers != nil && len(s.ResourceIdentifiers) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceIdentifiers", 1))
	}
	if s.ResourceIdentifiers != nil {
		for i, v := range s.ResourceIdentifiers {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ResourceIdentifiers", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/BatchGetAggregateResourceConfigResponse
type BatchGetAggregateResourceConfigOutput struct {
	_ struct{} `type:"structure"`

	// A list that contains the current configuration of one or more resources.
	BaseConfigurationItems []BaseConfigurationItem `type:"list"`

	// A list of resource identifiers that were not processed with current scope.
	// The list is empty if all the resources are processed.
	UnprocessedResourceIdentifiers []AggregateResourceIdentifier `type:"list"`
}

// String returns the string representation
func (s BatchGetAggregateResourceConfigOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchGetAggregateResourceConfig = "BatchGetAggregateResourceConfig"

// BatchGetAggregateResourceConfigRequest returns a request value for making API operation for
// AWS Config.
//
// Returns the current configuration items for resources that are present in
// your AWS Config aggregator. The operation also returns a list of resources
// that are not processed in the current request. If there are no unprocessed
// resources, the operation returns an empty unprocessedResourceIdentifiers
// list.
//
//    * The API does not return results for deleted resources.
//
//    * The API does not return tags and relationships.
//
//    // Example sending a request using BatchGetAggregateResourceConfigRequest.
//    req := client.BatchGetAggregateResourceConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/BatchGetAggregateResourceConfig
func (c *Client) BatchGetAggregateResourceConfigRequest(input *BatchGetAggregateResourceConfigInput) BatchGetAggregateResourceConfigRequest {
	op := &aws.Operation{
		Name:       opBatchGetAggregateResourceConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchGetAggregateResourceConfigInput{}
	}

	req := c.newRequest(op, input, &BatchGetAggregateResourceConfigOutput{})
	return BatchGetAggregateResourceConfigRequest{Request: req, Input: input, Copy: c.BatchGetAggregateResourceConfigRequest}
}

// BatchGetAggregateResourceConfigRequest is the request type for the
// BatchGetAggregateResourceConfig API operation.
type BatchGetAggregateResourceConfigRequest struct {
	*aws.Request
	Input *BatchGetAggregateResourceConfigInput
	Copy  func(*BatchGetAggregateResourceConfigInput) BatchGetAggregateResourceConfigRequest
}

// Send marshals and sends the BatchGetAggregateResourceConfig API request.
func (r BatchGetAggregateResourceConfigRequest) Send(ctx context.Context) (*BatchGetAggregateResourceConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchGetAggregateResourceConfigResponse{
		BatchGetAggregateResourceConfigOutput: r.Request.Data.(*BatchGetAggregateResourceConfigOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchGetAggregateResourceConfigResponse is the response type for the
// BatchGetAggregateResourceConfig API operation.
type BatchGetAggregateResourceConfigResponse struct {
	*BatchGetAggregateResourceConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchGetAggregateResourceConfig request.
func (r *BatchGetAggregateResourceConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
