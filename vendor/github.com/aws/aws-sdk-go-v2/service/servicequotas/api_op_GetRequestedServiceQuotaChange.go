// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicequotas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/service-quotas-2019-06-24/GetRequestedServiceQuotaChangeRequest
type GetRequestedServiceQuotaChangeInput struct {
	_ struct{} `type:"structure"`

	// Identifies the quota increase request.
	//
	// RequestId is a required field
	RequestId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRequestedServiceQuotaChangeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRequestedServiceQuotaChangeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRequestedServiceQuotaChangeInput"}

	if s.RequestId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RequestId"))
	}
	if s.RequestId != nil && len(*s.RequestId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RequestId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/service-quotas-2019-06-24/GetRequestedServiceQuotaChangeResponse
type GetRequestedServiceQuotaChangeOutput struct {
	_ struct{} `type:"structure"`

	// Returns the RequestedServiceQuotaChange object for the specific increase
	// request.
	RequestedQuota *RequestedServiceQuotaChange `type:"structure"`
}

// String returns the string representation
func (s GetRequestedServiceQuotaChangeOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetRequestedServiceQuotaChange = "GetRequestedServiceQuotaChange"

// GetRequestedServiceQuotaChangeRequest returns a request value for making API operation for
// Service Quotas.
//
// Retrieves the details for a particular increase request.
//
//    // Example sending a request using GetRequestedServiceQuotaChangeRequest.
//    req := client.GetRequestedServiceQuotaChangeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/service-quotas-2019-06-24/GetRequestedServiceQuotaChange
func (c *Client) GetRequestedServiceQuotaChangeRequest(input *GetRequestedServiceQuotaChangeInput) GetRequestedServiceQuotaChangeRequest {
	op := &aws.Operation{
		Name:       opGetRequestedServiceQuotaChange,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRequestedServiceQuotaChangeInput{}
	}

	req := c.newRequest(op, input, &GetRequestedServiceQuotaChangeOutput{})
	return GetRequestedServiceQuotaChangeRequest{Request: req, Input: input, Copy: c.GetRequestedServiceQuotaChangeRequest}
}

// GetRequestedServiceQuotaChangeRequest is the request type for the
// GetRequestedServiceQuotaChange API operation.
type GetRequestedServiceQuotaChangeRequest struct {
	*aws.Request
	Input *GetRequestedServiceQuotaChangeInput
	Copy  func(*GetRequestedServiceQuotaChangeInput) GetRequestedServiceQuotaChangeRequest
}

// Send marshals and sends the GetRequestedServiceQuotaChange API request.
func (r GetRequestedServiceQuotaChangeRequest) Send(ctx context.Context) (*GetRequestedServiceQuotaChangeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRequestedServiceQuotaChangeResponse{
		GetRequestedServiceQuotaChangeOutput: r.Request.Data.(*GetRequestedServiceQuotaChangeOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRequestedServiceQuotaChangeResponse is the response type for the
// GetRequestedServiceQuotaChange API operation.
type GetRequestedServiceQuotaChangeResponse struct {
	*GetRequestedServiceQuotaChangeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRequestedServiceQuotaChange request.
func (r *GetRequestedServiceQuotaChangeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
