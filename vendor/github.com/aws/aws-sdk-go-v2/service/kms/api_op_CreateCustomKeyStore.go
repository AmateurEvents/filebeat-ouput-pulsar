// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/CreateCustomKeyStoreRequest
type CreateCustomKeyStoreInput struct {
	_ struct{} `type:"structure"`

	// Identifies the AWS CloudHSM cluster for the custom key store. Enter the cluster
	// ID of any active AWS CloudHSM cluster that is not already associated with
	// a custom key store. To find the cluster ID, use the DescribeClusters (https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_DescribeClusters.html)
	// operation.
	//
	// CloudHsmClusterId is a required field
	CloudHsmClusterId *string `min:"19" type:"string" required:"true"`

	// Specifies a friendly name for the custom key store. The name must be unique
	// in your AWS account.
	//
	// CustomKeyStoreName is a required field
	CustomKeyStoreName *string `min:"1" type:"string" required:"true"`

	// Enter the password of the kmsuser crypto user (CU) account (https://docs.aws.amazon.com/kms/latest/developerguide/key-store-concepts.html#concept-kmsuser)
	// in the specified AWS CloudHSM cluster. AWS KMS logs into the cluster as this
	// user to manage key material on your behalf.
	//
	// This parameter tells AWS KMS the kmsuser account password; it does not change
	// the password in the AWS CloudHSM cluster.
	//
	// KeyStorePassword is a required field
	KeyStorePassword *string `min:"1" type:"string" required:"true"`

	// Enter the content of the trust anchor certificate for the cluster. This is
	// the content of the customerCA.crt file that you created when you initialized
	// the cluster (https://docs.aws.amazon.com/cloudhsm/latest/userguide/initialize-cluster.html).
	//
	// TrustAnchorCertificate is a required field
	TrustAnchorCertificate *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateCustomKeyStoreInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateCustomKeyStoreInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateCustomKeyStoreInput"}

	if s.CloudHsmClusterId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CloudHsmClusterId"))
	}
	if s.CloudHsmClusterId != nil && len(*s.CloudHsmClusterId) < 19 {
		invalidParams.Add(aws.NewErrParamMinLen("CloudHsmClusterId", 19))
	}

	if s.CustomKeyStoreName == nil {
		invalidParams.Add(aws.NewErrParamRequired("CustomKeyStoreName"))
	}
	if s.CustomKeyStoreName != nil && len(*s.CustomKeyStoreName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CustomKeyStoreName", 1))
	}

	if s.KeyStorePassword == nil {
		invalidParams.Add(aws.NewErrParamRequired("KeyStorePassword"))
	}
	if s.KeyStorePassword != nil && len(*s.KeyStorePassword) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KeyStorePassword", 1))
	}

	if s.TrustAnchorCertificate == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrustAnchorCertificate"))
	}
	if s.TrustAnchorCertificate != nil && len(*s.TrustAnchorCertificate) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TrustAnchorCertificate", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/CreateCustomKeyStoreResponse
type CreateCustomKeyStoreOutput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for the new custom key store.
	CustomKeyStoreId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateCustomKeyStoreOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateCustomKeyStore = "CreateCustomKeyStore"

// CreateCustomKeyStoreRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Creates a custom key store (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
// that is associated with an AWS CloudHSM cluster (https://docs.aws.amazon.com/cloudhsm/latest/userguide/clusters.html)
// that you own and manage.
//
// This operation is part of the Custom Key Store feature (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
// feature in AWS KMS, which combines the convenience and extensive integration
// of AWS KMS with the isolation and control of a single-tenant key store.
//
// Before you create the custom key store, you must assemble the required elements,
// including an AWS CloudHSM cluster that fulfills the requirements for a custom
// key store. For details about the required elements, see Assemble the Prerequisites
// (https://docs.aws.amazon.com/kms/latest/developerguide/create-keystore.html#before-keystore)
// in the AWS Key Management Service Developer Guide.
//
// When the operation completes successfully, it returns the ID of the new custom
// key store. Before you can use your new custom key store, you need to use
// the ConnectCustomKeyStore operation to connect the new key store to its AWS
// CloudHSM cluster. Even if you are not going to use your custom key store
// immediately, you might want to connect it to verify that all settings are
// correct and then disconnect it until you are ready to use it.
//
// For help with failures, see Troubleshooting a Custom Key Store (https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html)
// in the AWS Key Management Service Developer Guide.
//
//    // Example sending a request using CreateCustomKeyStoreRequest.
//    req := client.CreateCustomKeyStoreRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/CreateCustomKeyStore
func (c *Client) CreateCustomKeyStoreRequest(input *CreateCustomKeyStoreInput) CreateCustomKeyStoreRequest {
	op := &aws.Operation{
		Name:       opCreateCustomKeyStore,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateCustomKeyStoreInput{}
	}

	req := c.newRequest(op, input, &CreateCustomKeyStoreOutput{})
	return CreateCustomKeyStoreRequest{Request: req, Input: input, Copy: c.CreateCustomKeyStoreRequest}
}

// CreateCustomKeyStoreRequest is the request type for the
// CreateCustomKeyStore API operation.
type CreateCustomKeyStoreRequest struct {
	*aws.Request
	Input *CreateCustomKeyStoreInput
	Copy  func(*CreateCustomKeyStoreInput) CreateCustomKeyStoreRequest
}

// Send marshals and sends the CreateCustomKeyStore API request.
func (r CreateCustomKeyStoreRequest) Send(ctx context.Context) (*CreateCustomKeyStoreResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateCustomKeyStoreResponse{
		CreateCustomKeyStoreOutput: r.Request.Data.(*CreateCustomKeyStoreOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateCustomKeyStoreResponse is the response type for the
// CreateCustomKeyStore API operation.
type CreateCustomKeyStoreResponse struct {
	*CreateCustomKeyStoreOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCustomKeyStore request.
func (r *CreateCustomKeyStoreResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
