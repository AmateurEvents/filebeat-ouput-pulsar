// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transcribe

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/transcribe-2017-10-26/GetTranscriptionJobRequest
type GetTranscriptionJobInput struct {
	_ struct{} `type:"structure"`

	// The name of the job.
	//
	// TranscriptionJobName is a required field
	TranscriptionJobName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetTranscriptionJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTranscriptionJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetTranscriptionJobInput"}

	if s.TranscriptionJobName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TranscriptionJobName"))
	}
	if s.TranscriptionJobName != nil && len(*s.TranscriptionJobName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TranscriptionJobName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/transcribe-2017-10-26/GetTranscriptionJobResponse
type GetTranscriptionJobOutput struct {
	_ struct{} `type:"structure"`

	// An object that contains the results of the transcription job.
	TranscriptionJob *TranscriptionJob `type:"structure"`
}

// String returns the string representation
func (s GetTranscriptionJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetTranscriptionJob = "GetTranscriptionJob"

// GetTranscriptionJobRequest returns a request value for making API operation for
// Amazon Transcribe Service.
//
// Returns information about a transcription job. To see the status of the job,
// check the TranscriptionJobStatus field. If the status is COMPLETED, the job
// is finished and you can find the results at the location specified in the
// TranscriptionFileUri field.
//
//    // Example sending a request using GetTranscriptionJobRequest.
//    req := client.GetTranscriptionJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transcribe-2017-10-26/GetTranscriptionJob
func (c *Client) GetTranscriptionJobRequest(input *GetTranscriptionJobInput) GetTranscriptionJobRequest {
	op := &aws.Operation{
		Name:       opGetTranscriptionJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetTranscriptionJobInput{}
	}

	req := c.newRequest(op, input, &GetTranscriptionJobOutput{})
	return GetTranscriptionJobRequest{Request: req, Input: input, Copy: c.GetTranscriptionJobRequest}
}

// GetTranscriptionJobRequest is the request type for the
// GetTranscriptionJob API operation.
type GetTranscriptionJobRequest struct {
	*aws.Request
	Input *GetTranscriptionJobInput
	Copy  func(*GetTranscriptionJobInput) GetTranscriptionJobRequest
}

// Send marshals and sends the GetTranscriptionJob API request.
func (r GetTranscriptionJobRequest) Send(ctx context.Context) (*GetTranscriptionJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetTranscriptionJobResponse{
		GetTranscriptionJobOutput: r.Request.Data.(*GetTranscriptionJobOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetTranscriptionJobResponse is the response type for the
// GetTranscriptionJob API operation.
type GetTranscriptionJobResponse struct {
	*GetTranscriptionJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetTranscriptionJob request.
func (r *GetTranscriptionJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
