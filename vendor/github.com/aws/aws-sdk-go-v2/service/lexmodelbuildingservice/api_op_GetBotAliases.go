// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/GetBotAliasesRequest
type GetBotAliasesInput struct {
	_ struct{} `type:"structure"`

	// The name of the bot.
	//
	// BotName is a required field
	BotName *string `location:"uri" locationName:"botName" min:"2" type:"string" required:"true"`

	// The maximum number of aliases to return in the response. The default is 50. .
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// Substring to match in bot alias names. An alias will be returned if any part
	// of its name matches the substring. For example, "xyz" matches both "xyzabc"
	// and "abcxyz."
	NameContains *string `location:"querystring" locationName:"nameContains" min:"1" type:"string"`

	// A pagination token for fetching the next page of aliases. If the response
	// to this call is truncated, Amazon Lex returns a pagination token in the response.
	// To fetch the next page of aliases, specify the pagination token in the next
	// request.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetBotAliasesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBotAliasesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBotAliasesInput"}

	if s.BotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BotName"))
	}
	if s.BotName != nil && len(*s.BotName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("BotName", 2))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NameContains != nil && len(*s.NameContains) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NameContains", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBotAliasesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.BotName != nil {
		v := *s.BotName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "botName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NameContains != nil {
		v := *s.NameContains

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nameContains", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/GetBotAliasesResponse
type GetBotAliasesOutput struct {
	_ struct{} `type:"structure"`

	// An array of BotAliasMetadata objects, each describing a bot alias.
	BotAliases []BotAliasMetadata `type:"list"`

	// A pagination token for fetching next page of aliases. If the response to
	// this call is truncated, Amazon Lex returns a pagination token in the response.
	// To fetch the next page of aliases, specify the pagination token in the next
	// request.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetBotAliasesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBotAliasesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BotAliases != nil {
		v := s.BotAliases

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "BotAliases", metadata)
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

const opGetBotAliases = "GetBotAliases"

// GetBotAliasesRequest returns a request value for making API operation for
// Amazon Lex Model Building Service.
//
// Returns a list of aliases for a specified Amazon Lex bot.
//
// This operation requires permissions for the lex:GetBotAliases action.
//
//    // Example sending a request using GetBotAliasesRequest.
//    req := client.GetBotAliasesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/GetBotAliases
func (c *Client) GetBotAliasesRequest(input *GetBotAliasesInput) GetBotAliasesRequest {
	op := &aws.Operation{
		Name:       opGetBotAliases,
		HTTPMethod: "GET",
		HTTPPath:   "/bots/{botName}/aliases/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetBotAliasesInput{}
	}

	req := c.newRequest(op, input, &GetBotAliasesOutput{})
	return GetBotAliasesRequest{Request: req, Input: input, Copy: c.GetBotAliasesRequest}
}

// GetBotAliasesRequest is the request type for the
// GetBotAliases API operation.
type GetBotAliasesRequest struct {
	*aws.Request
	Input *GetBotAliasesInput
	Copy  func(*GetBotAliasesInput) GetBotAliasesRequest
}

// Send marshals and sends the GetBotAliases API request.
func (r GetBotAliasesRequest) Send(ctx context.Context) (*GetBotAliasesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBotAliasesResponse{
		GetBotAliasesOutput: r.Request.Data.(*GetBotAliasesOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetBotAliasesRequestPaginator returns a paginator for GetBotAliases.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetBotAliasesRequest(input)
//   p := lexmodelbuildingservice.NewGetBotAliasesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetBotAliasesPaginator(req GetBotAliasesRequest) GetBotAliasesPaginator {
	return GetBotAliasesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetBotAliasesInput
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

// GetBotAliasesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetBotAliasesPaginator struct {
	aws.Pager
}

func (p *GetBotAliasesPaginator) CurrentPage() *GetBotAliasesOutput {
	return p.Pager.CurrentPage().(*GetBotAliasesOutput)
}

// GetBotAliasesResponse is the response type for the
// GetBotAliases API operation.
type GetBotAliasesResponse struct {
	*GetBotAliasesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBotAliases request.
func (r *GetBotAliasesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
