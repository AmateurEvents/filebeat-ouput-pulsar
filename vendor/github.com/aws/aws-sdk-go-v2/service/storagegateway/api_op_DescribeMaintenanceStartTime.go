// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A JSON object containing the of the gateway.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/DescribeMaintenanceStartTimeInput
type DescribeMaintenanceStartTimeInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and region.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeMaintenanceStartTimeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeMaintenanceStartTimeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeMaintenanceStartTimeInput"}

	if s.GatewayARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayARN"))
	}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A JSON object containing the following fields:
//
//    * DescribeMaintenanceStartTimeOutput$DayOfMonth
//
//    * DescribeMaintenanceStartTimeOutput$DayOfWeek
//
//    * DescribeMaintenanceStartTimeOutput$HourOfDay
//
//    * DescribeMaintenanceStartTimeOutput$MinuteOfHour
//
//    * DescribeMaintenanceStartTimeOutput$Timezone
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/DescribeMaintenanceStartTimeOutput
type DescribeMaintenanceStartTimeOutput struct {
	_ struct{} `type:"structure"`

	// The day of the month component of the maintenance start time represented
	// as an ordinal number from 1 to 28, where 1 represents the first day of the
	// month and 28 represents the last day of the month.
	//
	// This value is only available for tape and volume gateways.
	DayOfMonth *int64 `min:"1" type:"integer"`

	// An ordinal number between 0 and 6 that represents the day of the week, where
	// 0 represents Sunday and 6 represents Saturday. The day of week is in the
	// time zone of the gateway.
	DayOfWeek *int64 `type:"integer"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and region.
	GatewayARN *string `min:"50" type:"string"`

	// The hour component of the maintenance start time represented as hh, where
	// hh is the hour (0 to 23). The hour of the day is in the time zone of the
	// gateway.
	HourOfDay *int64 `type:"integer"`

	// The minute component of the maintenance start time represented as mm, where
	// mm is the minute (0 to 59). The minute of the hour is in the time zone of
	// the gateway.
	MinuteOfHour *int64 `type:"integer"`

	// A value that indicates the time zone that is set for the gateway. The start
	// time and day of week specified should be in the time zone of the gateway.
	Timezone *string `min:"3" type:"string"`
}

// String returns the string representation
func (s DescribeMaintenanceStartTimeOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeMaintenanceStartTime = "DescribeMaintenanceStartTime"

// DescribeMaintenanceStartTimeRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Returns your gateway's weekly maintenance start time including the day and
// time of the week. Note that values are in terms of the gateway's time zone.
//
//    // Example sending a request using DescribeMaintenanceStartTimeRequest.
//    req := client.DescribeMaintenanceStartTimeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/DescribeMaintenanceStartTime
func (c *Client) DescribeMaintenanceStartTimeRequest(input *DescribeMaintenanceStartTimeInput) DescribeMaintenanceStartTimeRequest {
	op := &aws.Operation{
		Name:       opDescribeMaintenanceStartTime,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeMaintenanceStartTimeInput{}
	}

	req := c.newRequest(op, input, &DescribeMaintenanceStartTimeOutput{})
	return DescribeMaintenanceStartTimeRequest{Request: req, Input: input, Copy: c.DescribeMaintenanceStartTimeRequest}
}

// DescribeMaintenanceStartTimeRequest is the request type for the
// DescribeMaintenanceStartTime API operation.
type DescribeMaintenanceStartTimeRequest struct {
	*aws.Request
	Input *DescribeMaintenanceStartTimeInput
	Copy  func(*DescribeMaintenanceStartTimeInput) DescribeMaintenanceStartTimeRequest
}

// Send marshals and sends the DescribeMaintenanceStartTime API request.
func (r DescribeMaintenanceStartTimeRequest) Send(ctx context.Context) (*DescribeMaintenanceStartTimeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeMaintenanceStartTimeResponse{
		DescribeMaintenanceStartTimeOutput: r.Request.Data.(*DescribeMaintenanceStartTimeOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeMaintenanceStartTimeResponse is the response type for the
// DescribeMaintenanceStartTime API operation.
type DescribeMaintenanceStartTimeResponse struct {
	*DescribeMaintenanceStartTimeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeMaintenanceStartTime request.
func (r *DescribeMaintenanceStartTimeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
