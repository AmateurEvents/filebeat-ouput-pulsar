// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicequotas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/service-quotas-2019-06-24/GetServiceQuotaIncreaseRequestFromTemplateRequest
type GetServiceQuotaIncreaseRequestFromTemplateInput struct {
	_ struct{} `type:"structure"`

	// Specifies the AWS Region for the quota that you want to use.
	//
	// AwsRegion is a required field
	AwsRegion *string `min:"1" type:"string" required:"true"`

	// Specifies the quota you want.
	//
	// QuotaCode is a required field
	QuotaCode *string `min:"1" type:"string" required:"true"`

	// Specifies the service that you want to use.
	//
	// ServiceCode is a required field
	ServiceCode *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetServiceQuotaIncreaseRequestFromTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetServiceQuotaIncreaseRequestFromTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetServiceQuotaIncreaseRequestFromTemplateInput"}

	if s.AwsRegion == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsRegion"))
	}
	if s.AwsRegion != nil && len(*s.AwsRegion) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsRegion", 1))
	}

	if s.QuotaCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("QuotaCode"))
	}
	if s.QuotaCode != nil && len(*s.QuotaCode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("QuotaCode", 1))
	}

	if s.ServiceCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServiceCode"))
	}
	if s.ServiceCode != nil && len(*s.ServiceCode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ServiceCode", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/service-quotas-2019-06-24/GetServiceQuotaIncreaseRequestFromTemplateResponse
type GetServiceQuotaIncreaseRequestFromTemplateOutput struct {
	_ struct{} `type:"structure"`

	// This object contains the details about the quota increase request.
	ServiceQuotaIncreaseRequestInTemplate *ServiceQuotaIncreaseRequestInTemplate `type:"structure"`
}

// String returns the string representation
func (s GetServiceQuotaIncreaseRequestFromTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetServiceQuotaIncreaseRequestFromTemplate = "GetServiceQuotaIncreaseRequestFromTemplate"

// GetServiceQuotaIncreaseRequestFromTemplateRequest returns a request value for making API operation for
// Service Quotas.
//
// Returns the details of the service quota increase request in your template.
//
//    // Example sending a request using GetServiceQuotaIncreaseRequestFromTemplateRequest.
//    req := client.GetServiceQuotaIncreaseRequestFromTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/service-quotas-2019-06-24/GetServiceQuotaIncreaseRequestFromTemplate
func (c *Client) GetServiceQuotaIncreaseRequestFromTemplateRequest(input *GetServiceQuotaIncreaseRequestFromTemplateInput) GetServiceQuotaIncreaseRequestFromTemplateRequest {
	op := &aws.Operation{
		Name:       opGetServiceQuotaIncreaseRequestFromTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetServiceQuotaIncreaseRequestFromTemplateInput{}
	}

	req := c.newRequest(op, input, &GetServiceQuotaIncreaseRequestFromTemplateOutput{})
	return GetServiceQuotaIncreaseRequestFromTemplateRequest{Request: req, Input: input, Copy: c.GetServiceQuotaIncreaseRequestFromTemplateRequest}
}

// GetServiceQuotaIncreaseRequestFromTemplateRequest is the request type for the
// GetServiceQuotaIncreaseRequestFromTemplate API operation.
type GetServiceQuotaIncreaseRequestFromTemplateRequest struct {
	*aws.Request
	Input *GetServiceQuotaIncreaseRequestFromTemplateInput
	Copy  func(*GetServiceQuotaIncreaseRequestFromTemplateInput) GetServiceQuotaIncreaseRequestFromTemplateRequest
}

// Send marshals and sends the GetServiceQuotaIncreaseRequestFromTemplate API request.
func (r GetServiceQuotaIncreaseRequestFromTemplateRequest) Send(ctx context.Context) (*GetServiceQuotaIncreaseRequestFromTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetServiceQuotaIncreaseRequestFromTemplateResponse{
		GetServiceQuotaIncreaseRequestFromTemplateOutput: r.Request.Data.(*GetServiceQuotaIncreaseRequestFromTemplateOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetServiceQuotaIncreaseRequestFromTemplateResponse is the response type for the
// GetServiceQuotaIncreaseRequestFromTemplate API operation.
type GetServiceQuotaIncreaseRequestFromTemplateResponse struct {
	*GetServiceQuotaIncreaseRequestFromTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetServiceQuotaIncreaseRequestFromTemplate request.
func (r *GetServiceQuotaIncreaseRequestFromTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
