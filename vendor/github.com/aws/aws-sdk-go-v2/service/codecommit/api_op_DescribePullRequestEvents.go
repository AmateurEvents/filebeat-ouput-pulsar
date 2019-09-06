// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/DescribePullRequestEventsInput
type DescribePullRequestEventsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the user whose actions resulted in the
	// event. Examples include updating the pull request with additional commits
	// or changing the status of a pull request.
	ActorArn *string `locationName:"actorArn" type:"string"`

	// A non-negative integer used to limit the number of returned results. The
	// default is 100 events, which is also the maximum number of events that can
	// be returned in a result.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// An enumeration token that when provided in a request, returns the next batch
	// of the results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Optional. The pull request event type about which you want to return information.
	PullRequestEventType PullRequestEventType `locationName:"pullRequestEventType" type:"string" enum:"true"`

	// The system-generated ID of the pull request. To get this ID, use ListPullRequests.
	//
	// PullRequestId is a required field
	PullRequestId *string `locationName:"pullRequestId" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribePullRequestEventsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribePullRequestEventsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribePullRequestEventsInput"}

	if s.PullRequestId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PullRequestId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/DescribePullRequestEventsOutput
type DescribePullRequestEventsOutput struct {
	_ struct{} `type:"structure"`

	// An enumeration token that can be used in a request to return the next batch
	// of the results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the pull request events.
	//
	// PullRequestEvents is a required field
	PullRequestEvents []PullRequestEvent `locationName:"pullRequestEvents" type:"list" required:"true"`
}

// String returns the string representation
func (s DescribePullRequestEventsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribePullRequestEvents = "DescribePullRequestEvents"

// DescribePullRequestEventsRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Returns information about one or more pull request events.
//
//    // Example sending a request using DescribePullRequestEventsRequest.
//    req := client.DescribePullRequestEventsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/DescribePullRequestEvents
func (c *Client) DescribePullRequestEventsRequest(input *DescribePullRequestEventsInput) DescribePullRequestEventsRequest {
	op := &aws.Operation{
		Name:       opDescribePullRequestEvents,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribePullRequestEventsInput{}
	}

	req := c.newRequest(op, input, &DescribePullRequestEventsOutput{})
	return DescribePullRequestEventsRequest{Request: req, Input: input, Copy: c.DescribePullRequestEventsRequest}
}

// DescribePullRequestEventsRequest is the request type for the
// DescribePullRequestEvents API operation.
type DescribePullRequestEventsRequest struct {
	*aws.Request
	Input *DescribePullRequestEventsInput
	Copy  func(*DescribePullRequestEventsInput) DescribePullRequestEventsRequest
}

// Send marshals and sends the DescribePullRequestEvents API request.
func (r DescribePullRequestEventsRequest) Send(ctx context.Context) (*DescribePullRequestEventsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribePullRequestEventsResponse{
		DescribePullRequestEventsOutput: r.Request.Data.(*DescribePullRequestEventsOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribePullRequestEventsRequestPaginator returns a paginator for DescribePullRequestEvents.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribePullRequestEventsRequest(input)
//   p := codecommit.NewDescribePullRequestEventsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribePullRequestEventsPaginator(req DescribePullRequestEventsRequest) DescribePullRequestEventsPaginator {
	return DescribePullRequestEventsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribePullRequestEventsInput
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

// DescribePullRequestEventsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribePullRequestEventsPaginator struct {
	aws.Pager
}

func (p *DescribePullRequestEventsPaginator) CurrentPage() *DescribePullRequestEventsOutput {
	return p.Pager.CurrentPage().(*DescribePullRequestEventsOutput)
}

// DescribePullRequestEventsResponse is the response type for the
// DescribePullRequestEvents API operation.
type DescribePullRequestEventsResponse struct {
	*DescribePullRequestEventsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribePullRequestEvents request.
func (r *DescribePullRequestEventsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
