// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/SearchTablesRequest
type SearchTablesInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier, consisting of account_id/datalake.
	CatalogId *string `min:"1" type:"string"`

	// A list of key-value pairs, and a comparator used to filter the search results.
	// Returns all entities matching the predicate.
	Filters []PropertyPredicate `type:"list"`

	// The maximum number of tables to return in a single response.
	MaxResults *int64 `min:"1" type:"integer"`

	// A continuation token, included if this is a continuation call.
	NextToken *string `type:"string"`

	// A string used for a text search.
	//
	// Specifying a value in quotes filters based on an exact match to the value.
	SearchText *string `type:"string"`

	// A list of criteria for sorting the results by a field name, in an ascending
	// or descending order.
	SortCriteria []SortCriterion `type:"list"`
}

// String returns the string representation
func (s SearchTablesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SearchTablesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SearchTablesInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/SearchTablesResponse
type SearchTablesOutput struct {
	_ struct{} `type:"structure"`

	// A continuation token, present if the current list segment is not the last.
	NextToken *string `type:"string"`

	// A list of the requested Table objects. The SearchTables response returns
	// only the tables that you have access to.
	TableList []Table `type:"list"`
}

// String returns the string representation
func (s SearchTablesOutput) String() string {
	return awsutil.Prettify(s)
}

const opSearchTables = "SearchTables"

// SearchTablesRequest returns a request value for making API operation for
// AWS Glue.
//
// Searches a set of tables based on properties in the table metadata as well
// as on the parent database. You can search against text or filter conditions.
//
// You can only get tables that you have access to based on the security policies
// defined in Lake Formation. You need at least a read-only access to the table
// for it to be returned. If you do not have access to all the columns in the
// table, these columns will not be searched against when returning the list
// of tables back to you. If you have access to the columns but not the data
// in the columns, those columns and the associated metadata for those columns
// will be included in the search.
//
//    // Example sending a request using SearchTablesRequest.
//    req := client.SearchTablesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/SearchTables
func (c *Client) SearchTablesRequest(input *SearchTablesInput) SearchTablesRequest {
	op := &aws.Operation{
		Name:       opSearchTables,
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
		input = &SearchTablesInput{}
	}

	req := c.newRequest(op, input, &SearchTablesOutput{})
	return SearchTablesRequest{Request: req, Input: input, Copy: c.SearchTablesRequest}
}

// SearchTablesRequest is the request type for the
// SearchTables API operation.
type SearchTablesRequest struct {
	*aws.Request
	Input *SearchTablesInput
	Copy  func(*SearchTablesInput) SearchTablesRequest
}

// Send marshals and sends the SearchTables API request.
func (r SearchTablesRequest) Send(ctx context.Context) (*SearchTablesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SearchTablesResponse{
		SearchTablesOutput: r.Request.Data.(*SearchTablesOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewSearchTablesRequestPaginator returns a paginator for SearchTables.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.SearchTablesRequest(input)
//   p := glue.NewSearchTablesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewSearchTablesPaginator(req SearchTablesRequest) SearchTablesPaginator {
	return SearchTablesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *SearchTablesInput
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

// SearchTablesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type SearchTablesPaginator struct {
	aws.Pager
}

func (p *SearchTablesPaginator) CurrentPage() *SearchTablesOutput {
	return p.Pager.CurrentPage().(*SearchTablesOutput)
}

// SearchTablesResponse is the response type for the
// SearchTables API operation.
type SearchTablesResponse struct {
	*SearchTablesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SearchTables request.
func (r *SearchTablesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
