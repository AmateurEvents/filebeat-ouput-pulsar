// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codepipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Represents the input of a PutJobSuccessResult action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/PutJobSuccessResultInput
type PutJobSuccessResultInput struct {
	_ struct{} `type:"structure"`

	// A token generated by a job worker, such as an AWS CodeDeploy deployment ID,
	// that a successful job provides to identify a custom action in progress. Future
	// jobs will use this token in order to identify the running instance of the
	// action. It can be reused to return additional information about the progress
	// of the custom action. When the action is complete, no continuation token
	// should be supplied.
	ContinuationToken *string `locationName:"continuationToken" min:"1" type:"string"`

	// The ID of the current revision of the artifact successfully worked upon by
	// the job.
	CurrentRevision *CurrentRevision `locationName:"currentRevision" type:"structure"`

	// The execution details of the successful job, such as the actions taken by
	// the job worker.
	ExecutionDetails *ExecutionDetails `locationName:"executionDetails" type:"structure"`

	// The unique system-generated ID of the job that succeeded. This is the same
	// ID returned from PollForJobs.
	//
	// JobId is a required field
	JobId *string `locationName:"jobId" type:"string" required:"true"`
}

// String returns the string representation
func (s PutJobSuccessResultInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutJobSuccessResultInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutJobSuccessResultInput"}
	if s.ContinuationToken != nil && len(*s.ContinuationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ContinuationToken", 1))
	}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.CurrentRevision != nil {
		if err := s.CurrentRevision.Validate(); err != nil {
			invalidParams.AddNested("CurrentRevision", err.(aws.ErrInvalidParams))
		}
	}
	if s.ExecutionDetails != nil {
		if err := s.ExecutionDetails.Validate(); err != nil {
			invalidParams.AddNested("ExecutionDetails", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/PutJobSuccessResultOutput
type PutJobSuccessResultOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutJobSuccessResultOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutJobSuccessResult = "PutJobSuccessResult"

// PutJobSuccessResultRequest returns a request value for making API operation for
// AWS CodePipeline.
//
// Represents the success of a job as returned to the pipeline by a job worker.
// Only used for custom actions.
//
//    // Example sending a request using PutJobSuccessResultRequest.
//    req := client.PutJobSuccessResultRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/PutJobSuccessResult
func (c *Client) PutJobSuccessResultRequest(input *PutJobSuccessResultInput) PutJobSuccessResultRequest {
	op := &aws.Operation{
		Name:       opPutJobSuccessResult,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutJobSuccessResultInput{}
	}

	req := c.newRequest(op, input, &PutJobSuccessResultOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PutJobSuccessResultRequest{Request: req, Input: input, Copy: c.PutJobSuccessResultRequest}
}

// PutJobSuccessResultRequest is the request type for the
// PutJobSuccessResult API operation.
type PutJobSuccessResultRequest struct {
	*aws.Request
	Input *PutJobSuccessResultInput
	Copy  func(*PutJobSuccessResultInput) PutJobSuccessResultRequest
}

// Send marshals and sends the PutJobSuccessResult API request.
func (r PutJobSuccessResultRequest) Send(ctx context.Context) (*PutJobSuccessResultResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutJobSuccessResultResponse{
		PutJobSuccessResultOutput: r.Request.Data.(*PutJobSuccessResultOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutJobSuccessResultResponse is the response type for the
// PutJobSuccessResult API operation.
type PutJobSuccessResultResponse struct {
	*PutJobSuccessResultOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutJobSuccessResult request.
func (r *PutJobSuccessResultResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
