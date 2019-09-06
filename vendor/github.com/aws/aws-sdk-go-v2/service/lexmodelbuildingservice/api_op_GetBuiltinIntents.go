// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/GetBuiltinIntentsRequest
type GetBuiltinIntentsInput struct {
	_ struct{} `type:"structure"`

	// A list of locales that the intent supports.
	Locale Locale `location:"querystring" locationName:"locale" type:"string" enum:"true"`

	// The maximum number of intents to return in the response. The default is 10.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// A pagination token that fetches the next page of intents. If this API call
	// is truncated, Amazon Lex returns a pagination token in the response. To fetch
	// the next page of intents, use the pagination token in the next request.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// Substring to match in built-in intent signatures. An intent will be returned
	// if any part of its signature matches the substring. For example, "xyz" matches
	// both "xyzabc" and "abcxyz." To find the signature for an intent, see Standard
	// Built-in Intents (https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/built-in-intent-ref/standard-intents)
	// in the Alexa Skills Kit.
	SignatureContains *string `location:"querystring" locationName:"signatureContains" type:"string"`
}

// String returns the string representation
func (s GetBuiltinIntentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBuiltinIntentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBuiltinIntentsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBuiltinIntentsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if len(s.Locale) > 0 {
		v := s.Locale

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "locale", protocol.QuotedValue{ValueMarshaler: v}, metadata)
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
	if s.SignatureContains != nil {
		v := *s.SignatureContains

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "signatureContains", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/GetBuiltinIntentsResponse
type GetBuiltinIntentsOutput struct {
	_ struct{} `type:"structure"`

	// An array of builtinIntentMetadata objects, one for each intent in the response.
	Intents []BuiltinIntentMetadata `locationName:"intents" type:"list"`

	// A pagination token that fetches the next page of intents. If the response
	// to this API call is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch the next page of intents, specify the pagination token
	// in the next request.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetBuiltinIntentsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBuiltinIntentsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Intents != nil {
		v := s.Intents

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "intents", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetBuiltinIntents = "GetBuiltinIntents"

// GetBuiltinIntentsRequest returns a request value for making API operation for
// Amazon Lex Model Building Service.
//
// Gets a list of built-in intents that meet the specified criteria.
//
// This operation requires permission for the lex:GetBuiltinIntents action.
//
//    // Example sending a request using GetBuiltinIntentsRequest.
//    req := client.GetBuiltinIntentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/GetBuiltinIntents
func (c *Client) GetBuiltinIntentsRequest(input *GetBuiltinIntentsInput) GetBuiltinIntentsRequest {
	op := &aws.Operation{
		Name:       opGetBuiltinIntents,
		HTTPMethod: "GET",
		HTTPPath:   "/builtins/intents/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetBuiltinIntentsInput{}
	}

	req := c.newRequest(op, input, &GetBuiltinIntentsOutput{})
	return GetBuiltinIntentsRequest{Request: req, Input: input, Copy: c.GetBuiltinIntentsRequest}
}

// GetBuiltinIntentsRequest is the request type for the
// GetBuiltinIntents API operation.
type GetBuiltinIntentsRequest struct {
	*aws.Request
	Input *GetBuiltinIntentsInput
	Copy  func(*GetBuiltinIntentsInput) GetBuiltinIntentsRequest
}

// Send marshals and sends the GetBuiltinIntents API request.
func (r GetBuiltinIntentsRequest) Send(ctx context.Context) (*GetBuiltinIntentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBuiltinIntentsResponse{
		GetBuiltinIntentsOutput: r.Request.Data.(*GetBuiltinIntentsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetBuiltinIntentsRequestPaginator returns a paginator for GetBuiltinIntents.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetBuiltinIntentsRequest(input)
//   p := lexmodelbuildingservice.NewGetBuiltinIntentsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetBuiltinIntentsPaginator(req GetBuiltinIntentsRequest) GetBuiltinIntentsPaginator {
	return GetBuiltinIntentsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetBuiltinIntentsInput
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

// GetBuiltinIntentsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetBuiltinIntentsPaginator struct {
	aws.Pager
}

func (p *GetBuiltinIntentsPaginator) CurrentPage() *GetBuiltinIntentsOutput {
	return p.Pager.CurrentPage().(*GetBuiltinIntentsOutput)
}

// GetBuiltinIntentsResponse is the response type for the
// GetBuiltinIntents API operation.
type GetBuiltinIntentsResponse struct {
	*GetBuiltinIntentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBuiltinIntents request.
func (r *GetBuiltinIntentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
