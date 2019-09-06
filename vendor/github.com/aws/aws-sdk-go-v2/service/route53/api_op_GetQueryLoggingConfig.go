// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/GetQueryLoggingConfigRequest
type GetQueryLoggingConfigInput struct {
	_ struct{} `type:"structure"`

	// The ID of the configuration for DNS query logging that you want to get information
	// about.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetQueryLoggingConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetQueryLoggingConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetQueryLoggingConfigInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetQueryLoggingConfigInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Id", protocol.StringValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/GetQueryLoggingConfigResponse
type GetQueryLoggingConfigOutput struct {
	_ struct{} `type:"structure"`

	// A complex type that contains information about the query logging configuration
	// that you specified in a GetQueryLoggingConfig (https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetQueryLoggingConfig.html)
	// request.
	//
	// QueryLoggingConfig is a required field
	QueryLoggingConfig *QueryLoggingConfig `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetQueryLoggingConfigOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetQueryLoggingConfigOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.QueryLoggingConfig != nil {
		v := s.QueryLoggingConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "QueryLoggingConfig", v, metadata)
	}
	return nil
}

const opGetQueryLoggingConfig = "GetQueryLoggingConfig"

// GetQueryLoggingConfigRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Gets information about a specified configuration for DNS query logging.
//
// For more information about DNS query logs, see CreateQueryLoggingConfig (https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateQueryLoggingConfig.html)
// and Logging DNS Queries (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/query-logs.html).
//
//    // Example sending a request using GetQueryLoggingConfigRequest.
//    req := client.GetQueryLoggingConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/GetQueryLoggingConfig
func (c *Client) GetQueryLoggingConfigRequest(input *GetQueryLoggingConfigInput) GetQueryLoggingConfigRequest {
	op := &aws.Operation{
		Name:       opGetQueryLoggingConfig,
		HTTPMethod: "GET",
		HTTPPath:   "/2013-04-01/queryloggingconfig/{Id}",
	}

	if input == nil {
		input = &GetQueryLoggingConfigInput{}
	}

	req := c.newRequest(op, input, &GetQueryLoggingConfigOutput{})
	return GetQueryLoggingConfigRequest{Request: req, Input: input, Copy: c.GetQueryLoggingConfigRequest}
}

// GetQueryLoggingConfigRequest is the request type for the
// GetQueryLoggingConfig API operation.
type GetQueryLoggingConfigRequest struct {
	*aws.Request
	Input *GetQueryLoggingConfigInput
	Copy  func(*GetQueryLoggingConfigInput) GetQueryLoggingConfigRequest
}

// Send marshals and sends the GetQueryLoggingConfig API request.
func (r GetQueryLoggingConfigRequest) Send(ctx context.Context) (*GetQueryLoggingConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetQueryLoggingConfigResponse{
		GetQueryLoggingConfigOutput: r.Request.Data.(*GetQueryLoggingConfigOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetQueryLoggingConfigResponse is the response type for the
// GetQueryLoggingConfig API operation.
type GetQueryLoggingConfigResponse struct {
	*GetQueryLoggingConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetQueryLoggingConfig request.
func (r *GetQueryLoggingConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
