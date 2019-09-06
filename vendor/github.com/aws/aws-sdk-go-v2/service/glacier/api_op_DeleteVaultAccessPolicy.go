// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glacier

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// DeleteVaultAccessPolicy input.
type DeleteVaultAccessPolicyInput struct {
	_ struct{} `type:"structure"`

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen),
	// in which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The name of the vault.
	//
	// VaultName is a required field
	VaultName *string `location:"uri" locationName:"vaultName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteVaultAccessPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVaultAccessPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteVaultAccessPolicyInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.VaultName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VaultName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteVaultAccessPolicyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VaultName != nil {
		v := *s.VaultName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "vaultName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteVaultAccessPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteVaultAccessPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteVaultAccessPolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteVaultAccessPolicy = "DeleteVaultAccessPolicy"

// DeleteVaultAccessPolicyRequest returns a request value for making API operation for
// Amazon Glacier.
//
// This operation deletes the access policy associated with the specified vault.
// The operation is eventually consistent; that is, it might take some time
// for Amazon S3 Glacier to completely remove the access policy, and you might
// still see the effect of the policy for a short time after you send the delete
// request.
//
// This operation is idempotent. You can invoke delete multiple times, even
// if there is no policy associated with the vault. For more information about
// vault access policies, see Amazon Glacier Access Control with Vault Access
// Policies (https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html).
//
//    // Example sending a request using DeleteVaultAccessPolicyRequest.
//    req := client.DeleteVaultAccessPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteVaultAccessPolicyRequest(input *DeleteVaultAccessPolicyInput) DeleteVaultAccessPolicyRequest {
	op := &aws.Operation{
		Name:       opDeleteVaultAccessPolicy,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{accountId}/vaults/{vaultName}/access-policy",
	}

	if input == nil {
		input = &DeleteVaultAccessPolicyInput{}
	}

	req := c.newRequest(op, input, &DeleteVaultAccessPolicyOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteVaultAccessPolicyRequest{Request: req, Input: input, Copy: c.DeleteVaultAccessPolicyRequest}
}

// DeleteVaultAccessPolicyRequest is the request type for the
// DeleteVaultAccessPolicy API operation.
type DeleteVaultAccessPolicyRequest struct {
	*aws.Request
	Input *DeleteVaultAccessPolicyInput
	Copy  func(*DeleteVaultAccessPolicyInput) DeleteVaultAccessPolicyRequest
}

// Send marshals and sends the DeleteVaultAccessPolicy API request.
func (r DeleteVaultAccessPolicyRequest) Send(ctx context.Context) (*DeleteVaultAccessPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVaultAccessPolicyResponse{
		DeleteVaultAccessPolicyOutput: r.Request.Data.(*DeleteVaultAccessPolicyOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVaultAccessPolicyResponse is the response type for the
// DeleteVaultAccessPolicy API operation.
type DeleteVaultAccessPolicyResponse struct {
	*DeleteVaultAccessPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVaultAccessPolicy request.
func (r *DeleteVaultAccessPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
