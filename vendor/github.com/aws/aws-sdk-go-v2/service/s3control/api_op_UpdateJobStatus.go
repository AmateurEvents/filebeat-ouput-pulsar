// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3control

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3control-2018-08-20/UpdateJobStatusRequest
type UpdateJobStatusInput struct {
	_ struct{} `type:"structure"`

	// AccountId is a required field
	AccountId *string `location:"header" locationName:"x-amz-account-id" type:"string" required:"true"`

	// The ID of the job whose status you want to update.
	//
	// JobId is a required field
	JobId *string `location:"uri" locationName:"id" min:"5" type:"string" required:"true"`

	// The status that you want to move the specified job to.
	//
	// RequestedJobStatus is a required field
	RequestedJobStatus RequestedJobStatus `location:"querystring" locationName:"requestedJobStatus" type:"string" required:"true" enum:"true"`

	// A description of the reason why you want to change the specified job's status.
	// This field can be any string up to the maximum length.
	StatusUpdateReason *string `location:"querystring" locationName:"statusUpdateReason" min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateJobStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateJobStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateJobStatusInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 5 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 5))
	}
	if len(s.RequestedJobStatus) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("RequestedJobStatus"))
	}
	if s.StatusUpdateReason != nil && len(*s.StatusUpdateReason) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StatusUpdateReason", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateJobStatusInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-account-id", protocol.StringValue(v), metadata)
	}
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "id", protocol.StringValue(v), metadata)
	}
	if len(s.RequestedJobStatus) > 0 {
		v := s.RequestedJobStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "requestedJobStatus", v, metadata)
	}
	if s.StatusUpdateReason != nil {
		v := *s.StatusUpdateReason

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "statusUpdateReason", protocol.StringValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3control-2018-08-20/UpdateJobStatusResult
type UpdateJobStatusOutput struct {
	_ struct{} `type:"structure"`

	// The ID for the job whose status was updated.
	JobId *string `min:"5" type:"string"`

	// The current status for the specified job.
	Status JobStatus `type:"string" enum:"true"`

	// The reason that the specified job's status was updated.
	StatusUpdateReason *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateJobStatusOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateJobStatusOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "JobId", protocol.StringValue(v), metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Status", v, metadata)
	}
	if s.StatusUpdateReason != nil {
		v := *s.StatusUpdateReason

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StatusUpdateReason", protocol.StringValue(v), metadata)
	}
	return nil
}

const opUpdateJobStatus = "UpdateJobStatus"

// UpdateJobStatusRequest returns a request value for making API operation for
// AWS S3 Control.
//
// Updates the status for the specified job. Use this operation to confirm that
// you want to run a job or to cancel an existing job.
//
//    // Example sending a request using UpdateJobStatusRequest.
//    req := client.UpdateJobStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3control-2018-08-20/UpdateJobStatus
func (c *Client) UpdateJobStatusRequest(input *UpdateJobStatusInput) UpdateJobStatusRequest {
	op := &aws.Operation{
		Name:       opUpdateJobStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/v20180820/jobs/{id}/status",
	}

	if input == nil {
		input = &UpdateJobStatusInput{}
	}

	req := c.newRequest(op, input, &UpdateJobStatusOutput{})
	req.Handlers.Build.PushBackNamed(buildPrefixHostHandler("AccountID", aws.StringValue(input.AccountId)))
	req.Handlers.Build.PushBackNamed(buildRemoveHeaderHandler("X-Amz-Account-Id"))
	return UpdateJobStatusRequest{Request: req, Input: input, Copy: c.UpdateJobStatusRequest}
}

// UpdateJobStatusRequest is the request type for the
// UpdateJobStatus API operation.
type UpdateJobStatusRequest struct {
	*aws.Request
	Input *UpdateJobStatusInput
	Copy  func(*UpdateJobStatusInput) UpdateJobStatusRequest
}

// Send marshals and sends the UpdateJobStatus API request.
func (r UpdateJobStatusRequest) Send(ctx context.Context) (*UpdateJobStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateJobStatusResponse{
		UpdateJobStatusOutput: r.Request.Data.(*UpdateJobStatusOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateJobStatusResponse is the response type for the
// UpdateJobStatus API operation.
type UpdateJobStatusResponse struct {
	*UpdateJobStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateJobStatus request.
func (r *UpdateJobStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
