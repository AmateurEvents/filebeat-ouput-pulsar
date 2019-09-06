// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// This input determines which instance groups to retrieve.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/ListInstanceGroupsInput
type ListInstanceGroupsInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the cluster for which to list the instance groups.
	//
	// ClusterId is a required field
	ClusterId *string `type:"string" required:"true"`

	// The pagination token that indicates the next set of results to retrieve.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s ListInstanceGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListInstanceGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListInstanceGroupsInput"}

	if s.ClusterId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// This input determines which instance groups to retrieve.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/ListInstanceGroupsOutput
type ListInstanceGroupsOutput struct {
	_ struct{} `type:"structure"`

	// The list of instance groups for the cluster and given filters.
	InstanceGroups []InstanceGroup `type:"list"`

	// The pagination token that indicates the next set of results to retrieve.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s ListInstanceGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListInstanceGroups = "ListInstanceGroups"

// ListInstanceGroupsRequest returns a request value for making API operation for
// Amazon Elastic MapReduce.
//
// Provides all available details about the instance groups in a cluster.
//
//    // Example sending a request using ListInstanceGroupsRequest.
//    req := client.ListInstanceGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/ListInstanceGroups
func (c *Client) ListInstanceGroupsRequest(input *ListInstanceGroupsInput) ListInstanceGroupsRequest {
	op := &aws.Operation{
		Name:       opListInstanceGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListInstanceGroupsInput{}
	}

	req := c.newRequest(op, input, &ListInstanceGroupsOutput{})
	return ListInstanceGroupsRequest{Request: req, Input: input, Copy: c.ListInstanceGroupsRequest}
}

// ListInstanceGroupsRequest is the request type for the
// ListInstanceGroups API operation.
type ListInstanceGroupsRequest struct {
	*aws.Request
	Input *ListInstanceGroupsInput
	Copy  func(*ListInstanceGroupsInput) ListInstanceGroupsRequest
}

// Send marshals and sends the ListInstanceGroups API request.
func (r ListInstanceGroupsRequest) Send(ctx context.Context) (*ListInstanceGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListInstanceGroupsResponse{
		ListInstanceGroupsOutput: r.Request.Data.(*ListInstanceGroupsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListInstanceGroupsRequestPaginator returns a paginator for ListInstanceGroups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListInstanceGroupsRequest(input)
//   p := emr.NewListInstanceGroupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListInstanceGroupsPaginator(req ListInstanceGroupsRequest) ListInstanceGroupsPaginator {
	return ListInstanceGroupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListInstanceGroupsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListInstanceGroupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListInstanceGroupsPaginator struct {
	aws.Pager
}

func (p *ListInstanceGroupsPaginator) CurrentPage() *ListInstanceGroupsOutput {
	return p.Pager.CurrentPage().(*ListInstanceGroupsOutput)
}

// ListInstanceGroupsResponse is the response type for the
// ListInstanceGroups API operation.
type ListInstanceGroupsResponse struct {
	*ListInstanceGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListInstanceGroups request.
func (r *ListInstanceGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
