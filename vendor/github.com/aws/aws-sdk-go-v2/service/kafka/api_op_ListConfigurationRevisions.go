// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/ListConfigurationRevisionsRequest
type ListConfigurationRevisionsInput struct {
	_ struct{} `type:"structure"`

	// Arn is a required field
	Arn *string `location:"uri" locationName:"arn" type:"string" required:"true"`

	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListConfigurationRevisionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListConfigurationRevisionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListConfigurationRevisionsInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListConfigurationRevisionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Information about revisions of an MSK configuration.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/ListConfigurationRevisionsResponse
type ListConfigurationRevisionsOutput struct {
	_ struct{} `type:"structure"`

	// Paginated results marker.
	NextToken *string `locationName:"nextToken" type:"string"`

	// List of ConfigurationRevision objects.
	Revisions []ConfigurationRevision `locationName:"revisions" type:"list"`
}

// String returns the string representation
func (s ListConfigurationRevisionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListConfigurationRevisionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Revisions != nil {
		v := s.Revisions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "revisions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListConfigurationRevisions = "ListConfigurationRevisions"

// ListConfigurationRevisionsRequest returns a request value for making API operation for
// Managed Streaming for Kafka.
//
// Returns a list of all the MSK configurations in this Region.
//
//    // Example sending a request using ListConfigurationRevisionsRequest.
//    req := client.ListConfigurationRevisionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/ListConfigurationRevisions
func (c *Client) ListConfigurationRevisionsRequest(input *ListConfigurationRevisionsInput) ListConfigurationRevisionsRequest {
	op := &aws.Operation{
		Name:       opListConfigurationRevisions,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/configurations/{arn}/revisions",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListConfigurationRevisionsInput{}
	}

	req := c.newRequest(op, input, &ListConfigurationRevisionsOutput{})
	return ListConfigurationRevisionsRequest{Request: req, Input: input, Copy: c.ListConfigurationRevisionsRequest}
}

// ListConfigurationRevisionsRequest is the request type for the
// ListConfigurationRevisions API operation.
type ListConfigurationRevisionsRequest struct {
	*aws.Request
	Input *ListConfigurationRevisionsInput
	Copy  func(*ListConfigurationRevisionsInput) ListConfigurationRevisionsRequest
}

// Send marshals and sends the ListConfigurationRevisions API request.
func (r ListConfigurationRevisionsRequest) Send(ctx context.Context) (*ListConfigurationRevisionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListConfigurationRevisionsResponse{
		ListConfigurationRevisionsOutput: r.Request.Data.(*ListConfigurationRevisionsOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListConfigurationRevisionsRequestPaginator returns a paginator for ListConfigurationRevisions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListConfigurationRevisionsRequest(input)
//   p := kafka.NewListConfigurationRevisionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListConfigurationRevisionsPaginator(req ListConfigurationRevisionsRequest) ListConfigurationRevisionsPaginator {
	return ListConfigurationRevisionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListConfigurationRevisionsInput
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

// ListConfigurationRevisionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListConfigurationRevisionsPaginator struct {
	aws.Pager
}

func (p *ListConfigurationRevisionsPaginator) CurrentPage() *ListConfigurationRevisionsOutput {
	return p.Pager.CurrentPage().(*ListConfigurationRevisionsOutput)
}

// ListConfigurationRevisionsResponse is the response type for the
// ListConfigurationRevisions API operation.
type ListConfigurationRevisionsResponse struct {
	*ListConfigurationRevisionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListConfigurationRevisions request.
func (r *ListConfigurationRevisionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
