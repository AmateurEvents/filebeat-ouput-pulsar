// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/StopMonitoringMembersRequest
type StopMonitoringMembersInput struct {
	_ struct{} `type:"structure"`

	// A list of account IDs of the GuardDuty member accounts whose findings you
	// want the master account to stop monitoring.
	//
	// AccountIds is a required field
	AccountIds []string `locationName:"accountIds" min:"1" type:"list" required:"true"`

	// The unique ID of the detector of the GuardDuty account that you want to stop
	// from monitor members' findings.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StopMonitoringMembersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopMonitoringMembersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopMonitoringMembersInput"}

	if s.AccountIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountIds"))
	}
	if s.AccountIds != nil && len(s.AccountIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AccountIds", 1))
	}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StopMonitoringMembersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccountIds != nil {
		v := s.AccountIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "accountIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.DetectorId != nil {
		v := *s.DetectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/StopMonitoringMembersResponse
type StopMonitoringMembersOutput struct {
	_ struct{} `type:"structure"`

	// A list of objects containing the unprocessed account and a result string
	// explaining why it was unprocessed.
	//
	// UnprocessedAccounts is a required field
	UnprocessedAccounts []UnprocessedAccount `locationName:"unprocessedAccounts" type:"list" required:"true"`
}

// String returns the string representation
func (s StopMonitoringMembersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StopMonitoringMembersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.UnprocessedAccounts != nil {
		v := s.UnprocessedAccounts

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "unprocessedAccounts", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opStopMonitoringMembers = "StopMonitoringMembers"

// StopMonitoringMembersRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Disables GuardDuty from monitoring findings of the member accounts specified
// by the account IDs. After running this command, a master GuardDuty account
// can run StartMonitoringMembers to re-enable GuardDuty to monitor these members’
// findings.
//
//    // Example sending a request using StopMonitoringMembersRequest.
//    req := client.StopMonitoringMembersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/StopMonitoringMembers
func (c *Client) StopMonitoringMembersRequest(input *StopMonitoringMembersInput) StopMonitoringMembersRequest {
	op := &aws.Operation{
		Name:       opStopMonitoringMembers,
		HTTPMethod: "POST",
		HTTPPath:   "/detector/{detectorId}/member/stop",
	}

	if input == nil {
		input = &StopMonitoringMembersInput{}
	}

	req := c.newRequest(op, input, &StopMonitoringMembersOutput{})
	return StopMonitoringMembersRequest{Request: req, Input: input, Copy: c.StopMonitoringMembersRequest}
}

// StopMonitoringMembersRequest is the request type for the
// StopMonitoringMembers API operation.
type StopMonitoringMembersRequest struct {
	*aws.Request
	Input *StopMonitoringMembersInput
	Copy  func(*StopMonitoringMembersInput) StopMonitoringMembersRequest
}

// Send marshals and sends the StopMonitoringMembers API request.
func (r StopMonitoringMembersRequest) Send(ctx context.Context) (*StopMonitoringMembersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopMonitoringMembersResponse{
		StopMonitoringMembersOutput: r.Request.Data.(*StopMonitoringMembersOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopMonitoringMembersResponse is the response type for the
// StopMonitoringMembers API operation.
type StopMonitoringMembersResponse struct {
	*StopMonitoringMembersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopMonitoringMembers request.
func (r *StopMonitoringMembersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
