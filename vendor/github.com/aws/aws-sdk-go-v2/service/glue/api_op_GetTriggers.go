// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetTriggersRequest
type GetTriggersInput struct {
	_ struct{} `type:"structure"`

	// The name of the job to retrieve triggers for. The trigger that can start
	// this job is returned, and if there is no such trigger, all triggers are returned.
	DependentJobName *string `min:"1" type:"string"`

	// The maximum size of the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// A continuation token, if this is a continuation call.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetTriggersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTriggersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetTriggersInput"}
	if s.DependentJobName != nil && len(*s.DependentJobName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DependentJobName", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetTriggersResponse
type GetTriggersOutput struct {
	_ struct{} `type:"structure"`

	// A continuation token, if not all the requested triggers have yet been returned.
	NextToken *string `type:"string"`

	// A list of triggers for the specified job.
	Triggers []Trigger `type:"list"`
}

// String returns the string representation
func (s GetTriggersOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetTriggers = "GetTriggers"

// GetTriggersRequest returns a request value for making API operation for
// AWS Glue.
//
// Gets all the triggers associated with a job.
//
//    // Example sending a request using GetTriggersRequest.
//    req := client.GetTriggersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetTriggers
func (c *Client) GetTriggersRequest(input *GetTriggersInput) GetTriggersRequest {
	op := &aws.Operation{
		Name:       opGetTriggers,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetTriggersInput{}
	}

	req := c.newRequest(op, input, &GetTriggersOutput{})
	return GetTriggersRequest{Request: req, Input: input, Copy: c.GetTriggersRequest}
}

// GetTriggersRequest is the request type for the
// GetTriggers API operation.
type GetTriggersRequest struct {
	*aws.Request
	Input *GetTriggersInput
	Copy  func(*GetTriggersInput) GetTriggersRequest
}

// Send marshals and sends the GetTriggers API request.
func (r GetTriggersRequest) Send(ctx context.Context) (*GetTriggersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetTriggersResponse{
		GetTriggersOutput: r.Request.Data.(*GetTriggersOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetTriggersRequestPaginator returns a paginator for GetTriggers.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetTriggersRequest(input)
//   p := glue.NewGetTriggersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetTriggersPaginator(req GetTriggersRequest) GetTriggersPaginator {
	return GetTriggersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetTriggersInput
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

// GetTriggersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetTriggersPaginator struct {
	aws.Pager
}

func (p *GetTriggersPaginator) CurrentPage() *GetTriggersOutput {
	return p.Pager.CurrentPage().(*GetTriggersOutput)
}

// GetTriggersResponse is the response type for the
// GetTriggers API operation.
type GetTriggersResponse struct {
	*GetTriggersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetTriggers request.
func (r *GetTriggersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
