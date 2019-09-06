// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package firehose

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/firehose-2015-08-04/StopDeliveryStreamEncryptionInput
type StopDeliveryStreamEncryptionInput struct {
	_ struct{} `type:"structure"`

	// The name of the delivery stream for which you want to disable server-side
	// encryption (SSE).
	//
	// DeliveryStreamName is a required field
	DeliveryStreamName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StopDeliveryStreamEncryptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopDeliveryStreamEncryptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopDeliveryStreamEncryptionInput"}

	if s.DeliveryStreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeliveryStreamName"))
	}
	if s.DeliveryStreamName != nil && len(*s.DeliveryStreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeliveryStreamName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/firehose-2015-08-04/StopDeliveryStreamEncryptionOutput
type StopDeliveryStreamEncryptionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StopDeliveryStreamEncryptionOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopDeliveryStreamEncryption = "StopDeliveryStreamEncryption"

// StopDeliveryStreamEncryptionRequest returns a request value for making API operation for
// Amazon Kinesis Firehose.
//
// Disables server-side encryption (SSE) for the delivery stream.
//
// This operation is asynchronous. It returns immediately. When you invoke it,
// Kinesis Data Firehose first sets the status of the stream to DISABLING, and
// then to DISABLED. You can continue to read and write data to your stream
// while its status is DISABLING. It can take up to 5 seconds after the encryption
// status changes to DISABLED before all records written to the delivery stream
// are no longer subject to encryption. To find out whether a record or a batch
// of records was encrypted, check the response elements PutRecordOutput$Encrypted
// and PutRecordBatchOutput$Encrypted, respectively.
//
// To check the encryption state of a delivery stream, use DescribeDeliveryStream.
//
// The StartDeliveryStreamEncryption and StopDeliveryStreamEncryption operations
// have a combined limit of 25 calls per delivery stream per 24 hours. For example,
// you reach the limit if you call StartDeliveryStreamEncryption 13 times and
// StopDeliveryStreamEncryption 12 times for the same delivery stream in a 24-hour
// period.
//
//    // Example sending a request using StopDeliveryStreamEncryptionRequest.
//    req := client.StopDeliveryStreamEncryptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/firehose-2015-08-04/StopDeliveryStreamEncryption
func (c *Client) StopDeliveryStreamEncryptionRequest(input *StopDeliveryStreamEncryptionInput) StopDeliveryStreamEncryptionRequest {
	op := &aws.Operation{
		Name:       opStopDeliveryStreamEncryption,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopDeliveryStreamEncryptionInput{}
	}

	req := c.newRequest(op, input, &StopDeliveryStreamEncryptionOutput{})
	return StopDeliveryStreamEncryptionRequest{Request: req, Input: input, Copy: c.StopDeliveryStreamEncryptionRequest}
}

// StopDeliveryStreamEncryptionRequest is the request type for the
// StopDeliveryStreamEncryption API operation.
type StopDeliveryStreamEncryptionRequest struct {
	*aws.Request
	Input *StopDeliveryStreamEncryptionInput
	Copy  func(*StopDeliveryStreamEncryptionInput) StopDeliveryStreamEncryptionRequest
}

// Send marshals and sends the StopDeliveryStreamEncryption API request.
func (r StopDeliveryStreamEncryptionRequest) Send(ctx context.Context) (*StopDeliveryStreamEncryptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopDeliveryStreamEncryptionResponse{
		StopDeliveryStreamEncryptionOutput: r.Request.Data.(*StopDeliveryStreamEncryptionOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopDeliveryStreamEncryptionResponse is the response type for the
// StopDeliveryStreamEncryption API operation.
type StopDeliveryStreamEncryptionResponse struct {
	*StopDeliveryStreamEncryptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopDeliveryStreamEncryption request.
func (r *StopDeliveryStreamEncryptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
