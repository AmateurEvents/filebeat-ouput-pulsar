// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package datasynciface provides an interface to enable mocking the AWS DataSync service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package datasynciface

import (
	"github.com/aws/aws-sdk-go-v2/service/datasync"
)

// ClientAPI provides an interface to enable mocking the
// datasync.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // DataSync.
//    func myFunc(svc datasynciface.ClientAPI) bool {
//        // Make svc.CancelTaskExecution request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := datasync.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        datasynciface.ClientPI
//    }
//    func (m *mockClientClient) CancelTaskExecution(input *datasync.CancelTaskExecutionInput) (*datasync.CancelTaskExecutionOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	CancelTaskExecutionRequest(*datasync.CancelTaskExecutionInput) datasync.CancelTaskExecutionRequest

	CreateAgentRequest(*datasync.CreateAgentInput) datasync.CreateAgentRequest

	CreateLocationEfsRequest(*datasync.CreateLocationEfsInput) datasync.CreateLocationEfsRequest

	CreateLocationNfsRequest(*datasync.CreateLocationNfsInput) datasync.CreateLocationNfsRequest

	CreateLocationS3Request(*datasync.CreateLocationS3Input) datasync.CreateLocationS3Request

	CreateLocationSmbRequest(*datasync.CreateLocationSmbInput) datasync.CreateLocationSmbRequest

	CreateTaskRequest(*datasync.CreateTaskInput) datasync.CreateTaskRequest

	DeleteAgentRequest(*datasync.DeleteAgentInput) datasync.DeleteAgentRequest

	DeleteLocationRequest(*datasync.DeleteLocationInput) datasync.DeleteLocationRequest

	DeleteTaskRequest(*datasync.DeleteTaskInput) datasync.DeleteTaskRequest

	DescribeAgentRequest(*datasync.DescribeAgentInput) datasync.DescribeAgentRequest

	DescribeLocationEfsRequest(*datasync.DescribeLocationEfsInput) datasync.DescribeLocationEfsRequest

	DescribeLocationNfsRequest(*datasync.DescribeLocationNfsInput) datasync.DescribeLocationNfsRequest

	DescribeLocationS3Request(*datasync.DescribeLocationS3Input) datasync.DescribeLocationS3Request

	DescribeLocationSmbRequest(*datasync.DescribeLocationSmbInput) datasync.DescribeLocationSmbRequest

	DescribeTaskRequest(*datasync.DescribeTaskInput) datasync.DescribeTaskRequest

	DescribeTaskExecutionRequest(*datasync.DescribeTaskExecutionInput) datasync.DescribeTaskExecutionRequest

	ListAgentsRequest(*datasync.ListAgentsInput) datasync.ListAgentsRequest

	ListLocationsRequest(*datasync.ListLocationsInput) datasync.ListLocationsRequest

	ListTagsForResourceRequest(*datasync.ListTagsForResourceInput) datasync.ListTagsForResourceRequest

	ListTaskExecutionsRequest(*datasync.ListTaskExecutionsInput) datasync.ListTaskExecutionsRequest

	ListTasksRequest(*datasync.ListTasksInput) datasync.ListTasksRequest

	StartTaskExecutionRequest(*datasync.StartTaskExecutionInput) datasync.StartTaskExecutionRequest

	TagResourceRequest(*datasync.TagResourceInput) datasync.TagResourceRequest

	UntagResourceRequest(*datasync.UntagResourceInput) datasync.UntagResourceRequest

	UpdateAgentRequest(*datasync.UpdateAgentInput) datasync.UpdateAgentRequest

	UpdateTaskRequest(*datasync.UpdateTaskInput) datasync.UpdateTaskRequest
}

var _ ClientAPI = (*datasync.Client)(nil)
