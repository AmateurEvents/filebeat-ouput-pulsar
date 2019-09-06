// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/ListGroupsRequest
type ListGroupsInput struct {
	_ struct{} `type:"structure"`

	// The limit of the request to list groups.
	Limit *int64 `type:"integer"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `min:"1" type:"string"`

	// The user pool ID for the user pool.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListGroupsInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/ListGroupsResponse
type ListGroupsOutput struct {
	_ struct{} `type:"structure"`

	// The group objects for the groups.
	Groups []GroupType `type:"list"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListGroups = "ListGroups"

// ListGroupsRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Lists the groups associated with a user pool.
//
// Requires developer credentials.
//
//    // Example sending a request using ListGroupsRequest.
//    req := client.ListGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/ListGroups
func (c *Client) ListGroupsRequest(input *ListGroupsInput) ListGroupsRequest {
	op := &aws.Operation{
		Name:       opListGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "Limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListGroupsInput{}
	}

	req := c.newRequest(op, input, &ListGroupsOutput{})
	return ListGroupsRequest{Request: req, Input: input, Copy: c.ListGroupsRequest}
}

// ListGroupsRequest is the request type for the
// ListGroups API operation.
type ListGroupsRequest struct {
	*aws.Request
	Input *ListGroupsInput
	Copy  func(*ListGroupsInput) ListGroupsRequest
}

// Send marshals and sends the ListGroups API request.
func (r ListGroupsRequest) Send(ctx context.Context) (*ListGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListGroupsResponse{
		ListGroupsOutput: r.Request.Data.(*ListGroupsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListGroupsRequestPaginator returns a paginator for ListGroups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListGroupsRequest(input)
//   p := cognitoidentityprovider.NewListGroupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListGroupsPaginator(req ListGroupsRequest) ListGroupsPaginator {
	return ListGroupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListGroupsInput
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

// ListGroupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListGroupsPaginator struct {
	aws.Pager
}

func (p *ListGroupsPaginator) CurrentPage() *ListGroupsOutput {
	return p.Pager.CurrentPage().(*ListGroupsOutput)
}

// ListGroupsResponse is the response type for the
// ListGroups API operation.
type ListGroupsResponse struct {
	*ListGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListGroups request.
func (r *ListGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
