// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The input for the AttachThingPrincipal operation.
type AttachThingPrincipalInput struct {
	_ struct{} `type:"structure"`

	// The principal, such as a certificate or other credential.
	//
	// Principal is a required field
	Principal *string `location:"header" locationName:"x-amzn-principal" type:"string" required:"true"`

	// The name of the thing.
	//
	// ThingName is a required field
	ThingName *string `location:"uri" locationName:"thingName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s AttachThingPrincipalInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AttachThingPrincipalInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AttachThingPrincipalInput"}

	if s.Principal == nil {
		invalidParams.Add(aws.NewErrParamRequired("Principal"))
	}

	if s.ThingName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingName"))
	}
	if s.ThingName != nil && len(*s.ThingName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AttachThingPrincipalInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Principal != nil {
		v := *s.Principal

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amzn-principal", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingName != nil {
		v := *s.ThingName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "thingName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The output from the AttachThingPrincipal operation.
type AttachThingPrincipalOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AttachThingPrincipalOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AttachThingPrincipalOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opAttachThingPrincipal = "AttachThingPrincipal"

// AttachThingPrincipalRequest returns a request value for making API operation for
// AWS IoT.
//
// Attaches the specified principal to the specified thing. A principal can
// be X.509 certificates, IAM users, groups, and roles, Amazon Cognito identities
// or federated identities.
//
//    // Example sending a request using AttachThingPrincipalRequest.
//    req := client.AttachThingPrincipalRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) AttachThingPrincipalRequest(input *AttachThingPrincipalInput) AttachThingPrincipalRequest {
	op := &aws.Operation{
		Name:       opAttachThingPrincipal,
		HTTPMethod: "PUT",
		HTTPPath:   "/things/{thingName}/principals",
	}

	if input == nil {
		input = &AttachThingPrincipalInput{}
	}

	req := c.newRequest(op, input, &AttachThingPrincipalOutput{})
	return AttachThingPrincipalRequest{Request: req, Input: input, Copy: c.AttachThingPrincipalRequest}
}

// AttachThingPrincipalRequest is the request type for the
// AttachThingPrincipal API operation.
type AttachThingPrincipalRequest struct {
	*aws.Request
	Input *AttachThingPrincipalInput
	Copy  func(*AttachThingPrincipalInput) AttachThingPrincipalRequest
}

// Send marshals and sends the AttachThingPrincipal API request.
func (r AttachThingPrincipalRequest) Send(ctx context.Context) (*AttachThingPrincipalResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AttachThingPrincipalResponse{
		AttachThingPrincipalOutput: r.Request.Data.(*AttachThingPrincipalOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AttachThingPrincipalResponse is the response type for the
// AttachThingPrincipal API operation.
type AttachThingPrincipalResponse struct {
	*AttachThingPrincipalOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AttachThingPrincipal request.
func (r *AttachThingPrincipalResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
