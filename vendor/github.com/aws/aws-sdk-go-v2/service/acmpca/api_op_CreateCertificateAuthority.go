// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acmpca

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/CreateCertificateAuthorityRequest
type CreateCertificateAuthorityInput struct {
	_ struct{} `type:"structure"`

	// Name and bit size of the private key algorithm, the name of the signing algorithm,
	// and X.500 certificate subject information.
	//
	// CertificateAuthorityConfiguration is a required field
	CertificateAuthorityConfiguration *CertificateAuthorityConfiguration `type:"structure" required:"true"`

	// The type of the certificate authority.
	//
	// CertificateAuthorityType is a required field
	CertificateAuthorityType CertificateAuthorityType `type:"string" required:"true" enum:"true"`

	// Alphanumeric string that can be used to distinguish between calls to CreateCertificateAuthority.
	// Idempotency tokens time out after five minutes. Therefore, if you call CreateCertificateAuthority
	// multiple times with the same idempotency token within a five minute period,
	// ACM Private CA recognizes that you are requesting only one certificate. As
	// a result, ACM Private CA issues only one. If you change the idempotency token
	// for each call, however, ACM Private CA recognizes that you are requesting
	// multiple certificates.
	IdempotencyToken *string `min:"1" type:"string"`

	// Contains a Boolean value that you can use to enable a certification revocation
	// list (CRL) for the CA, the name of the S3 bucket to which ACM Private CA
	// will write the CRL, and an optional CNAME alias that you can use to hide
	// the name of your bucket in the CRL Distribution Points extension of your
	// CA certificate. For more information, see the CrlConfiguration structure.
	RevocationConfiguration *RevocationConfiguration `type:"structure"`

	// Key-value pairs that will be attached to the new private CA. You can associate
	// up to 50 tags with a private CA. For information using tags with
	//
	// IAM to manage permissions, see Controlling Access Using IAM Tags (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_iam-tags.html).
	Tags []Tag `min:"1" type:"list"`
}

// String returns the string representation
func (s CreateCertificateAuthorityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateCertificateAuthorityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateCertificateAuthorityInput"}

	if s.CertificateAuthorityConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateAuthorityConfiguration"))
	}
	if len(s.CertificateAuthorityType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("CertificateAuthorityType"))
	}
	if s.IdempotencyToken != nil && len(*s.IdempotencyToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdempotencyToken", 1))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.CertificateAuthorityConfiguration != nil {
		if err := s.CertificateAuthorityConfiguration.Validate(); err != nil {
			invalidParams.AddNested("CertificateAuthorityConfiguration", err.(aws.ErrInvalidParams))
		}
	}
	if s.RevocationConfiguration != nil {
		if err := s.RevocationConfiguration.Validate(); err != nil {
			invalidParams.AddNested("RevocationConfiguration", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/CreateCertificateAuthorityResponse
type CreateCertificateAuthorityOutput struct {
	_ struct{} `type:"structure"`

	// If successful, the Amazon Resource Name (ARN) of the certificate authority
	// (CA). This is of the form:
	//
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012 .
	CertificateAuthorityArn *string `min:"5" type:"string"`
}

// String returns the string representation
func (s CreateCertificateAuthorityOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateCertificateAuthority = "CreateCertificateAuthority"

// CreateCertificateAuthorityRequest returns a request value for making API operation for
// AWS Certificate Manager Private Certificate Authority.
//
// Creates a root or subordinate private certificate authority (CA). You must
// specify the CA configuration, the certificate revocation list (CRL) configuration,
// the CA type, and an optional idempotency token to avoid accidental creation
// of multiple CAs. The CA configuration specifies the name of the algorithm
// and key size to be used to create the CA private key, the type of signing
// algorithm that the CA uses, and X.500 subject information. The CRL configuration
// specifies the CRL expiration period in days (the validity period of the CRL),
// the Amazon S3 bucket that will contain the CRL, and a CNAME alias for the
// S3 bucket that is included in certificates issued by the CA. If successful,
// this action returns the Amazon Resource Name (ARN) of the CA.
//
//    // Example sending a request using CreateCertificateAuthorityRequest.
//    req := client.CreateCertificateAuthorityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/CreateCertificateAuthority
func (c *Client) CreateCertificateAuthorityRequest(input *CreateCertificateAuthorityInput) CreateCertificateAuthorityRequest {
	op := &aws.Operation{
		Name:       opCreateCertificateAuthority,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateCertificateAuthorityInput{}
	}

	req := c.newRequest(op, input, &CreateCertificateAuthorityOutput{})
	return CreateCertificateAuthorityRequest{Request: req, Input: input, Copy: c.CreateCertificateAuthorityRequest}
}

// CreateCertificateAuthorityRequest is the request type for the
// CreateCertificateAuthority API operation.
type CreateCertificateAuthorityRequest struct {
	*aws.Request
	Input *CreateCertificateAuthorityInput
	Copy  func(*CreateCertificateAuthorityInput) CreateCertificateAuthorityRequest
}

// Send marshals and sends the CreateCertificateAuthority API request.
func (r CreateCertificateAuthorityRequest) Send(ctx context.Context) (*CreateCertificateAuthorityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateCertificateAuthorityResponse{
		CreateCertificateAuthorityOutput: r.Request.Data.(*CreateCertificateAuthorityOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateCertificateAuthorityResponse is the response type for the
// CreateCertificateAuthority API operation.
type CreateCertificateAuthorityResponse struct {
	*CreateCertificateAuthorityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCertificateAuthority request.
func (r *CreateCertificateAuthorityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
