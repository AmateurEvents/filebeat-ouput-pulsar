// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lakeformation

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lakeformation-2017-03-31/BatchRevokePermissionsRequest
type BatchRevokePermissionsInput struct {
	_ struct{} `type:"structure"`

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your AWS Lake
	// Formation environment.
	CatalogId *string `min:"1" type:"string"`

	// A list of up to 20 entries for resource permissions to be revoked by batch
	// operation to the principal.
	//
	// Entries is a required field
	Entries []BatchPermissionsRequestEntry `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchRevokePermissionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchRevokePermissionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchRevokePermissionsInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.Entries == nil {
		invalidParams.Add(aws.NewErrParamRequired("Entries"))
	}
	if s.Entries != nil {
		for i, v := range s.Entries {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Entries", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lakeformation-2017-03-31/BatchRevokePermissionsResponse
type BatchRevokePermissionsOutput struct {
	_ struct{} `type:"structure"`

	// A list of failures to revoke permissions to the resources.
	Failures []BatchPermissionsFailureEntry `type:"list"`
}

// String returns the string representation
func (s BatchRevokePermissionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchRevokePermissions = "BatchRevokePermissions"

// BatchRevokePermissionsRequest returns a request value for making API operation for
// AWS Lake Formation.
//
// Batch operation to revoke permissions from the principal.
//
//    // Example sending a request using BatchRevokePermissionsRequest.
//    req := client.BatchRevokePermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lakeformation-2017-03-31/BatchRevokePermissions
func (c *Client) BatchRevokePermissionsRequest(input *BatchRevokePermissionsInput) BatchRevokePermissionsRequest {
	op := &aws.Operation{
		Name:       opBatchRevokePermissions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchRevokePermissionsInput{}
	}

	req := c.newRequest(op, input, &BatchRevokePermissionsOutput{})
	return BatchRevokePermissionsRequest{Request: req, Input: input, Copy: c.BatchRevokePermissionsRequest}
}

// BatchRevokePermissionsRequest is the request type for the
// BatchRevokePermissions API operation.
type BatchRevokePermissionsRequest struct {
	*aws.Request
	Input *BatchRevokePermissionsInput
	Copy  func(*BatchRevokePermissionsInput) BatchRevokePermissionsRequest
}

// Send marshals and sends the BatchRevokePermissions API request.
func (r BatchRevokePermissionsRequest) Send(ctx context.Context) (*BatchRevokePermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchRevokePermissionsResponse{
		BatchRevokePermissionsOutput: r.Request.Data.(*BatchRevokePermissionsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchRevokePermissionsResponse is the response type for the
// BatchRevokePermissions API operation.
type BatchRevokePermissionsResponse struct {
	*BatchRevokePermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchRevokePermissions request.
func (r *BatchRevokePermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
