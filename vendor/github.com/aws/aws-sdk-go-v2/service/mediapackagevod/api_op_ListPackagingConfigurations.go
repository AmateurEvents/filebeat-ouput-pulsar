// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediapackagevod

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/ListPackagingConfigurationsRequest
type ListPackagingConfigurationsInput struct {
	_ struct{} `type:"structure"`

	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	PackagingGroupId *string `location:"querystring" locationName:"packagingGroupId" type:"string"`
}

// String returns the string representation
func (s ListPackagingConfigurationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPackagingConfigurationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListPackagingConfigurationsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListPackagingConfigurationsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

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
	if s.PackagingGroupId != nil {
		v := *s.PackagingGroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "packagingGroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/ListPackagingConfigurationsResponse
type ListPackagingConfigurationsOutput struct {
	_ struct{} `type:"structure"`

	NextToken *string `locationName:"nextToken" type:"string"`

	PackagingConfigurations []PackagingConfiguration `locationName:"packagingConfigurations" type:"list"`
}

// String returns the string representation
func (s ListPackagingConfigurationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListPackagingConfigurationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PackagingConfigurations != nil {
		v := s.PackagingConfigurations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "packagingConfigurations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListPackagingConfigurations = "ListPackagingConfigurations"

// ListPackagingConfigurationsRequest returns a request value for making API operation for
// AWS Elemental MediaPackage VOD.
//
// Returns a collection of MediaPackage VOD PackagingConfiguration resources.
//
//    // Example sending a request using ListPackagingConfigurationsRequest.
//    req := client.ListPackagingConfigurationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/ListPackagingConfigurations
func (c *Client) ListPackagingConfigurationsRequest(input *ListPackagingConfigurationsInput) ListPackagingConfigurationsRequest {
	op := &aws.Operation{
		Name:       opListPackagingConfigurations,
		HTTPMethod: "GET",
		HTTPPath:   "/packaging_configurations",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListPackagingConfigurationsInput{}
	}

	req := c.newRequest(op, input, &ListPackagingConfigurationsOutput{})
	return ListPackagingConfigurationsRequest{Request: req, Input: input, Copy: c.ListPackagingConfigurationsRequest}
}

// ListPackagingConfigurationsRequest is the request type for the
// ListPackagingConfigurations API operation.
type ListPackagingConfigurationsRequest struct {
	*aws.Request
	Input *ListPackagingConfigurationsInput
	Copy  func(*ListPackagingConfigurationsInput) ListPackagingConfigurationsRequest
}

// Send marshals and sends the ListPackagingConfigurations API request.
func (r ListPackagingConfigurationsRequest) Send(ctx context.Context) (*ListPackagingConfigurationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPackagingConfigurationsResponse{
		ListPackagingConfigurationsOutput: r.Request.Data.(*ListPackagingConfigurationsOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListPackagingConfigurationsRequestPaginator returns a paginator for ListPackagingConfigurations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListPackagingConfigurationsRequest(input)
//   p := mediapackagevod.NewListPackagingConfigurationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListPackagingConfigurationsPaginator(req ListPackagingConfigurationsRequest) ListPackagingConfigurationsPaginator {
	return ListPackagingConfigurationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListPackagingConfigurationsInput
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

// ListPackagingConfigurationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListPackagingConfigurationsPaginator struct {
	aws.Pager
}

func (p *ListPackagingConfigurationsPaginator) CurrentPage() *ListPackagingConfigurationsOutput {
	return p.Pager.CurrentPage().(*ListPackagingConfigurationsOutput)
}

// ListPackagingConfigurationsResponse is the response type for the
// ListPackagingConfigurations API operation.
type ListPackagingConfigurationsResponse struct {
	*ListPackagingConfigurationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPackagingConfigurations request.
func (r *ListPackagingConfigurationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
