// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DeleteDBClusterSnapshotMessage
type DeleteDBClusterSnapshotInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the DB cluster snapshot to delete.
	//
	// Constraints: Must be the name of an existing DB cluster snapshot in the available
	// state.
	//
	// DBClusterSnapshotIdentifier is a required field
	DBClusterSnapshotIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDBClusterSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDBClusterSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDBClusterSnapshotInput"}

	if s.DBClusterSnapshotIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterSnapshotIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DeleteDBClusterSnapshotResult
type DeleteDBClusterSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details for an Amazon RDS DB cluster snapshot
	//
	// This data type is used as a response element in the DescribeDBClusterSnapshots
	// action.
	DBClusterSnapshot *DBClusterSnapshot `type:"structure"`
}

// String returns the string representation
func (s DeleteDBClusterSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteDBClusterSnapshot = "DeleteDBClusterSnapshot"

// DeleteDBClusterSnapshotRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Deletes a DB cluster snapshot. If the snapshot is being copied, the copy
// operation is terminated.
//
// The DB cluster snapshot must be in the available state to be deleted.
//
// For more information on Amazon Aurora, see What Is Amazon Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide.
//
// This action only applies to Aurora DB clusters.
//
//    // Example sending a request using DeleteDBClusterSnapshotRequest.
//    req := client.DeleteDBClusterSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DeleteDBClusterSnapshot
func (c *Client) DeleteDBClusterSnapshotRequest(input *DeleteDBClusterSnapshotInput) DeleteDBClusterSnapshotRequest {
	op := &aws.Operation{
		Name:       opDeleteDBClusterSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDBClusterSnapshotInput{}
	}

	req := c.newRequest(op, input, &DeleteDBClusterSnapshotOutput{})
	return DeleteDBClusterSnapshotRequest{Request: req, Input: input, Copy: c.DeleteDBClusterSnapshotRequest}
}

// DeleteDBClusterSnapshotRequest is the request type for the
// DeleteDBClusterSnapshot API operation.
type DeleteDBClusterSnapshotRequest struct {
	*aws.Request
	Input *DeleteDBClusterSnapshotInput
	Copy  func(*DeleteDBClusterSnapshotInput) DeleteDBClusterSnapshotRequest
}

// Send marshals and sends the DeleteDBClusterSnapshot API request.
func (r DeleteDBClusterSnapshotRequest) Send(ctx context.Context) (*DeleteDBClusterSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDBClusterSnapshotResponse{
		DeleteDBClusterSnapshotOutput: r.Request.Data.(*DeleteDBClusterSnapshotOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDBClusterSnapshotResponse is the response type for the
// DeleteDBClusterSnapshot API operation.
type DeleteDBClusterSnapshotResponse struct {
	*DeleteDBClusterSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDBClusterSnapshot request.
func (r *DeleteDBClusterSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
