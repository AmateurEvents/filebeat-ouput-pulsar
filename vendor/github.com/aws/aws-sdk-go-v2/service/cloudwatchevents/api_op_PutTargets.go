// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchevents

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/events-2015-10-07/PutTargetsRequest
type PutTargetsInput struct {
	_ struct{} `type:"structure"`

	// The name of the event bus associated with the rule. If you omit this, the
	// default event bus is used.
	EventBusName *string `min:"1" type:"string"`

	// The name of the rule.
	//
	// Rule is a required field
	Rule *string `min:"1" type:"string" required:"true"`

	// The targets to update or add to the rule.
	//
	// Targets is a required field
	Targets []Target `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s PutTargetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutTargetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutTargetsInput"}
	if s.EventBusName != nil && len(*s.EventBusName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventBusName", 1))
	}

	if s.Rule == nil {
		invalidParams.Add(aws.NewErrParamRequired("Rule"))
	}
	if s.Rule != nil && len(*s.Rule) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Rule", 1))
	}

	if s.Targets == nil {
		invalidParams.Add(aws.NewErrParamRequired("Targets"))
	}
	if s.Targets != nil && len(s.Targets) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Targets", 1))
	}
	if s.Targets != nil {
		for i, v := range s.Targets {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Targets", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/events-2015-10-07/PutTargetsResponse
type PutTargetsOutput struct {
	_ struct{} `type:"structure"`

	// The failed target entries.
	FailedEntries []PutTargetsResultEntry `type:"list"`

	// The number of failed entries.
	FailedEntryCount *int64 `type:"integer"`
}

// String returns the string representation
func (s PutTargetsOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutTargets = "PutTargets"

// PutTargetsRequest returns a request value for making API operation for
// Amazon CloudWatch Events.
//
// Adds the specified targets to the specified rule, or updates the targets
// if they're already associated with the rule.
//
// Targets are the resources that are invoked when a rule is triggered.
//
// You can configure the following as targets in EventBridge:
//
//    * EC2 instances
//
//    * SSM Run Command
//
//    * SSM Automation
//
//    * AWS Lambda functions
//
//    * Data streams in Amazon Kinesis Data Streams
//
//    * Data delivery streams in Amazon Kinesis Data Firehose
//
//    * Amazon ECS tasks
//
//    * AWS Step Functions state machines
//
//    * AWS Batch jobs
//
//    * AWS CodeBuild projects
//
//    * Pipelines in AWS CodePipeline
//
//    * Amazon Inspector assessment templates
//
//    * Amazon SNS topics
//
//    * Amazon SQS queues, including FIFO queues
//
//    * The default event bus of another AWS account
//
// Creating rules with built-in targets is supported only on the AWS Management
// Console. The built-in targets are EC2 CreateSnapshot API call, EC2 RebootInstances
// API call, EC2 StopInstances API call, and EC2 TerminateInstances API call.
//
// For some target types, PutTargets provides target-specific parameters. If
// the target is a Kinesis data stream, you can optionally specify which shard
// the event goes to by using the KinesisParameters argument. To invoke a command
// on multiple EC2 instances with one rule, you can use the RunCommandParameters
// field.
//
// To be able to make API calls against the resources that you own, Amazon EventBridge
// needs the appropriate permissions. For AWS Lambda and Amazon SNS resources,
// EventBridge relies on resource-based policies. For EC2 instances, Kinesis
// data streams, and AWS Step Functions state machines, EventBridge relies on
// IAM roles that you specify in the RoleARN argument in PutTargets. For more
// information, see Authentication and Access Control (https://docs.aws.amazon.com/eventbridge/latest/userguide/auth-and-access-control-eventbridge.html)
// in the Amazon EventBridge User Guide.
//
// If another AWS account is in the same Region and has granted you permission
// (using PutPermission), you can send events to that account. Set that account's
// event bus as a target of the rules in your account. To send the matched events
// to the other account, specify that account's event bus as the Arn value when
// you run PutTargets. If your account sends events to another account, your
// account is charged for each sent event. Each event sent to another account
// is charged as a custom event. The account receiving the event isn't charged.
// For more information, see Amazon EventBridge Pricing (https://aws.amazon.com/eventbridge/pricing/).
//
// If you're setting an event bus in another account as the target and that
// account granted permission to your account through an organization instead
// of directly by the account ID, you must specify a RoleArn with proper permissions
// in the Target structure. For more information, see Sending and Receiving
// Events Between AWS Accounts (https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-cross-account-event-delivery.html)
// in the Amazon EventBridge User Guide.
//
// For more information about enabling cross-account events, see PutPermission.
//
// Input, InputPath, and InputTransformer are mutually exclusive and optional
// parameters of a target. When a rule is triggered due to a matched event:
//
//    * If none of the following arguments are specified for a target, the entire
//    event is passed to the target in JSON format (unless the target is Amazon
//    EC2 Run Command or Amazon ECS task, in which case nothing from the event
//    is passed to the target).
//
//    * If Input is specified in the form of valid JSON, then the matched event
//    is overridden with this constant.
//
//    * If InputPath is specified in the form of JSONPath (for example, $.detail),
//    only the part of the event specified in the path is passed to the target
//    (for example, only the detail part of the event is passed).
//
//    * If InputTransformer is specified, one or more specified JSONPaths are
//    extracted from the event and used as values in a template that you specify
//    as the input to the target.
//
// When you specify InputPath or InputTransformer, you must use JSON dot notation,
// not bracket notation.
//
// When you add targets to a rule and the associated rule triggers soon after,
// new or updated targets might not be immediately invoked. Allow a short period
// of time for changes to take effect.
//
// This action can partially fail if too many requests are made at the same
// time. If that happens, FailedEntryCount is nonzero in the response, and each
// entry in FailedEntries provides the ID of the failed target and the error
// code.
//
//    // Example sending a request using PutTargetsRequest.
//    req := client.PutTargetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/events-2015-10-07/PutTargets
func (c *Client) PutTargetsRequest(input *PutTargetsInput) PutTargetsRequest {
	op := &aws.Operation{
		Name:       opPutTargets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutTargetsInput{}
	}

	req := c.newRequest(op, input, &PutTargetsOutput{})
	return PutTargetsRequest{Request: req, Input: input, Copy: c.PutTargetsRequest}
}

// PutTargetsRequest is the request type for the
// PutTargets API operation.
type PutTargetsRequest struct {
	*aws.Request
	Input *PutTargetsInput
	Copy  func(*PutTargetsInput) PutTargetsRequest
}

// Send marshals and sends the PutTargets API request.
func (r PutTargetsRequest) Send(ctx context.Context) (*PutTargetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutTargetsResponse{
		PutTargetsOutput: r.Request.Data.(*PutTargetsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutTargetsResponse is the response type for the
// PutTargets API operation.
type PutTargetsResponse struct {
	*PutTargetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutTargets request.
func (r *PutTargetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
