// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListAlgorithmsInput
type ListAlgorithmsInput struct {
	_ struct{} `type:"structure"`

	// A filter that returns only algorithms created after the specified time (timestamp).
	CreationTimeAfter *time.Time `type:"timestamp"`

	// A filter that returns only algorithms created before the specified time (timestamp).
	CreationTimeBefore *time.Time `type:"timestamp"`

	// The maximum number of algorithms to return in the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// A string in the algorithm name. This filter returns only algorithms whose
	// name contains the specified string.
	NameContains *string `type:"string"`

	// If the response to a previous ListAlgorithms request was truncated, the response
	// includes a NextToken. To retrieve the next set of algorithms, use the token
	// in the next request.
	NextToken *string `type:"string"`

	// The parameter by which to sort the results. The default is CreationTime.
	SortBy AlgorithmSortBy `type:"string" enum:"true"`

	// The sort order for the results. The default is Ascending.
	SortOrder SortOrder `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListAlgorithmsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAlgorithmsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAlgorithmsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListAlgorithmsOutput
type ListAlgorithmsOutput struct {
	_ struct{} `type:"structure"`

	// >An array of AlgorithmSummary objects, each of which lists an algorithm.
	//
	// AlgorithmSummaryList is a required field
	AlgorithmSummaryList []AlgorithmSummary `type:"list" required:"true"`

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of algorithms, use it in the subsequent request.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListAlgorithmsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListAlgorithms = "ListAlgorithms"

// ListAlgorithmsRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Lists the machine learning algorithms that have been created.
//
//    // Example sending a request using ListAlgorithmsRequest.
//    req := client.ListAlgorithmsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListAlgorithms
func (c *Client) ListAlgorithmsRequest(input *ListAlgorithmsInput) ListAlgorithmsRequest {
	op := &aws.Operation{
		Name:       opListAlgorithms,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListAlgorithmsInput{}
	}

	req := c.newRequest(op, input, &ListAlgorithmsOutput{})
	return ListAlgorithmsRequest{Request: req, Input: input, Copy: c.ListAlgorithmsRequest}
}

// ListAlgorithmsRequest is the request type for the
// ListAlgorithms API operation.
type ListAlgorithmsRequest struct {
	*aws.Request
	Input *ListAlgorithmsInput
	Copy  func(*ListAlgorithmsInput) ListAlgorithmsRequest
}

// Send marshals and sends the ListAlgorithms API request.
func (r ListAlgorithmsRequest) Send(ctx context.Context) (*ListAlgorithmsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAlgorithmsResponse{
		ListAlgorithmsOutput: r.Request.Data.(*ListAlgorithmsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListAlgorithmsResponse is the response type for the
// ListAlgorithms API operation.
type ListAlgorithmsResponse struct {
	*ListAlgorithmsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAlgorithms request.
func (r *ListAlgorithmsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
