// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotjobsdataplane

type JobExecutionStatus string

// Enum values for JobExecutionStatus
const (
	JobExecutionStatusQueued     JobExecutionStatus = "QUEUED"
	JobExecutionStatusInProgress JobExecutionStatus = "IN_PROGRESS"
	JobExecutionStatusSucceeded  JobExecutionStatus = "SUCCEEDED"
	JobExecutionStatusFailed     JobExecutionStatus = "FAILED"
	JobExecutionStatusTimedOut   JobExecutionStatus = "TIMED_OUT"
	JobExecutionStatusRejected   JobExecutionStatus = "REJECTED"
	JobExecutionStatusRemoved    JobExecutionStatus = "REMOVED"
	JobExecutionStatusCanceled   JobExecutionStatus = "CANCELED"
)

func (enum JobExecutionStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum JobExecutionStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
