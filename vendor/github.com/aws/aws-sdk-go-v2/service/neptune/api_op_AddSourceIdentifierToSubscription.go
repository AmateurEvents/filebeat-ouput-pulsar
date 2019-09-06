// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/AddSourceIdentifierToSubscriptionMessage
type AddSourceIdentifierToSubscriptionInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the event source to be added.
	//
	// Constraints:
	//
	//    * If the source type is a DB instance, then a DBInstanceIdentifier must
	//    be supplied.
	//
	//    * If the source type is a DB security group, a DBSecurityGroupName must
	//    be supplied.
	//
	//    * If the source type is a DB parameter group, a DBParameterGroupName must
	//    be supplied.
	//
	//    * If the source type is a DB snapshot, a DBSnapshotIdentifier must be
	//    supplied.
	//
	// SourceIdentifier is a required field
	SourceIdentifier *string `type:"string" required:"true"`

	// The name of the event notification subscription you want to add a source
	// identifier to.
	//
	// SubscriptionName is a required field
	SubscriptionName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AddSourceIdentifierToSubscriptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddSourceIdentifierToSubscriptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddSourceIdentifierToSubscriptionInput"}

	if s.SourceIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceIdentifier"))
	}

	if s.SubscriptionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubscriptionName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/AddSourceIdentifierToSubscriptionResult
type AddSourceIdentifierToSubscriptionOutput struct {
	_ struct{} `type:"structure"`

	// Contains the results of a successful invocation of the DescribeEventSubscriptions
	// action.
	EventSubscription *EventSubscription `type:"structure"`
}

// String returns the string representation
func (s AddSourceIdentifierToSubscriptionOutput) String() string {
	return awsutil.Prettify(s)
}

const opAddSourceIdentifierToSubscription = "AddSourceIdentifierToSubscription"

// AddSourceIdentifierToSubscriptionRequest returns a request value for making API operation for
// Amazon Neptune.
//
// Adds a source identifier to an existing event notification subscription.
//
//    // Example sending a request using AddSourceIdentifierToSubscriptionRequest.
//    req := client.AddSourceIdentifierToSubscriptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/AddSourceIdentifierToSubscription
func (c *Client) AddSourceIdentifierToSubscriptionRequest(input *AddSourceIdentifierToSubscriptionInput) AddSourceIdentifierToSubscriptionRequest {
	op := &aws.Operation{
		Name:       opAddSourceIdentifierToSubscription,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddSourceIdentifierToSubscriptionInput{}
	}

	req := c.newRequest(op, input, &AddSourceIdentifierToSubscriptionOutput{})
	return AddSourceIdentifierToSubscriptionRequest{Request: req, Input: input, Copy: c.AddSourceIdentifierToSubscriptionRequest}
}

// AddSourceIdentifierToSubscriptionRequest is the request type for the
// AddSourceIdentifierToSubscription API operation.
type AddSourceIdentifierToSubscriptionRequest struct {
	*aws.Request
	Input *AddSourceIdentifierToSubscriptionInput
	Copy  func(*AddSourceIdentifierToSubscriptionInput) AddSourceIdentifierToSubscriptionRequest
}

// Send marshals and sends the AddSourceIdentifierToSubscription API request.
func (r AddSourceIdentifierToSubscriptionRequest) Send(ctx context.Context) (*AddSourceIdentifierToSubscriptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddSourceIdentifierToSubscriptionResponse{
		AddSourceIdentifierToSubscriptionOutput: r.Request.Data.(*AddSourceIdentifierToSubscriptionOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddSourceIdentifierToSubscriptionResponse is the response type for the
// AddSourceIdentifierToSubscription API operation.
type AddSourceIdentifierToSubscriptionResponse struct {
	*AddSourceIdentifierToSubscriptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddSourceIdentifierToSubscription request.
func (r *AddSourceIdentifierToSubscriptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
