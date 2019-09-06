// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a ListDeploymentConfigs operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/ListDeploymentConfigsInput
type ListDeploymentConfigsInput struct {
	_ struct{} `type:"structure"`

	// An identifier returned from the previous ListDeploymentConfigs call. It can
	// be used to return the next set of deployment configurations in the list.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDeploymentConfigsInput) String() string {
	return awsutil.Prettify(s)
}

// Represents the output of a ListDeploymentConfigs operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/ListDeploymentConfigsOutput
type ListDeploymentConfigsOutput struct {
	_ struct{} `type:"structure"`

	// A list of deployment configurations, including built-in configurations such
	// as CodeDeployDefault.OneAtATime.
	DeploymentConfigsList []string `locationName:"deploymentConfigsList" type:"list"`

	// If a large amount of information is returned, an identifier is also returned.
	// It can be used in a subsequent list deployment configurations call to return
	// the next set of deployment configurations in the list.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDeploymentConfigsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListDeploymentConfigs = "ListDeploymentConfigs"

// ListDeploymentConfigsRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Lists the deployment configurations with the IAM user or AWS account.
//
//    // Example sending a request using ListDeploymentConfigsRequest.
//    req := client.ListDeploymentConfigsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/ListDeploymentConfigs
func (c *Client) ListDeploymentConfigsRequest(input *ListDeploymentConfigsInput) ListDeploymentConfigsRequest {
	op := &aws.Operation{
		Name:       opListDeploymentConfigs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDeploymentConfigsInput{}
	}

	req := c.newRequest(op, input, &ListDeploymentConfigsOutput{})
	return ListDeploymentConfigsRequest{Request: req, Input: input, Copy: c.ListDeploymentConfigsRequest}
}

// ListDeploymentConfigsRequest is the request type for the
// ListDeploymentConfigs API operation.
type ListDeploymentConfigsRequest struct {
	*aws.Request
	Input *ListDeploymentConfigsInput
	Copy  func(*ListDeploymentConfigsInput) ListDeploymentConfigsRequest
}

// Send marshals and sends the ListDeploymentConfigs API request.
func (r ListDeploymentConfigsRequest) Send(ctx context.Context) (*ListDeploymentConfigsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDeploymentConfigsResponse{
		ListDeploymentConfigsOutput: r.Request.Data.(*ListDeploymentConfigsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDeploymentConfigsRequestPaginator returns a paginator for ListDeploymentConfigs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDeploymentConfigsRequest(input)
//   p := codedeploy.NewListDeploymentConfigsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDeploymentConfigsPaginator(req ListDeploymentConfigsRequest) ListDeploymentConfigsPaginator {
	return ListDeploymentConfigsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDeploymentConfigsInput
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

// ListDeploymentConfigsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDeploymentConfigsPaginator struct {
	aws.Pager
}

func (p *ListDeploymentConfigsPaginator) CurrentPage() *ListDeploymentConfigsOutput {
	return p.Pager.CurrentPage().(*ListDeploymentConfigsOutput)
}

// ListDeploymentConfigsResponse is the response type for the
// ListDeploymentConfigs API operation.
type ListDeploymentConfigsResponse struct {
	*ListDeploymentConfigsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDeploymentConfigs request.
func (r *ListDeploymentConfigsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
