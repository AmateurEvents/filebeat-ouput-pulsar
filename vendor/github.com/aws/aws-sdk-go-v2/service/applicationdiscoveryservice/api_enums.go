// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationdiscoveryservice

type AgentStatus string

// Enum values for AgentStatus
const (
	AgentStatusHealthy     AgentStatus = "HEALTHY"
	AgentStatusUnhealthy   AgentStatus = "UNHEALTHY"
	AgentStatusRunning     AgentStatus = "RUNNING"
	AgentStatusUnknown     AgentStatus = "UNKNOWN"
	AgentStatusBlacklisted AgentStatus = "BLACKLISTED"
	AgentStatusShutdown    AgentStatus = "SHUTDOWN"
)

func (enum AgentStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AgentStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type BatchDeleteImportDataErrorCode string

// Enum values for BatchDeleteImportDataErrorCode
const (
	BatchDeleteImportDataErrorCodeNotFound            BatchDeleteImportDataErrorCode = "NOT_FOUND"
	BatchDeleteImportDataErrorCodeInternalServerError BatchDeleteImportDataErrorCode = "INTERNAL_SERVER_ERROR"
	BatchDeleteImportDataErrorCodeOverLimit           BatchDeleteImportDataErrorCode = "OVER_LIMIT"
)

func (enum BatchDeleteImportDataErrorCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum BatchDeleteImportDataErrorCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ConfigurationItemType string

// Enum values for ConfigurationItemType
const (
	ConfigurationItemTypeServer      ConfigurationItemType = "SERVER"
	ConfigurationItemTypeProcess     ConfigurationItemType = "PROCESS"
	ConfigurationItemTypeConnection  ConfigurationItemType = "CONNECTION"
	ConfigurationItemTypeApplication ConfigurationItemType = "APPLICATION"
)

func (enum ConfigurationItemType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ConfigurationItemType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ContinuousExportStatus string

// Enum values for ContinuousExportStatus
const (
	ContinuousExportStatusStartInProgress ContinuousExportStatus = "START_IN_PROGRESS"
	ContinuousExportStatusStartFailed     ContinuousExportStatus = "START_FAILED"
	ContinuousExportStatusActive          ContinuousExportStatus = "ACTIVE"
	ContinuousExportStatusError           ContinuousExportStatus = "ERROR"
	ContinuousExportStatusStopInProgress  ContinuousExportStatus = "STOP_IN_PROGRESS"
	ContinuousExportStatusStopFailed      ContinuousExportStatus = "STOP_FAILED"
	ContinuousExportStatusInactive        ContinuousExportStatus = "INACTIVE"
)

func (enum ContinuousExportStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ContinuousExportStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DataSource string

// Enum values for DataSource
const (
	DataSourceAgent DataSource = "AGENT"
)

func (enum DataSource) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DataSource) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ExportDataFormat string

// Enum values for ExportDataFormat
const (
	ExportDataFormatCsv     ExportDataFormat = "CSV"
	ExportDataFormatGraphml ExportDataFormat = "GRAPHML"
)

func (enum ExportDataFormat) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ExportDataFormat) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ExportStatus string

// Enum values for ExportStatus
const (
	ExportStatusFailed     ExportStatus = "FAILED"
	ExportStatusSucceeded  ExportStatus = "SUCCEEDED"
	ExportStatusInProgress ExportStatus = "IN_PROGRESS"
)

func (enum ExportStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ExportStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ImportStatus string

// Enum values for ImportStatus
const (
	ImportStatusImportInProgress                ImportStatus = "IMPORT_IN_PROGRESS"
	ImportStatusImportComplete                  ImportStatus = "IMPORT_COMPLETE"
	ImportStatusImportCompleteWithErrors        ImportStatus = "IMPORT_COMPLETE_WITH_ERRORS"
	ImportStatusImportFailed                    ImportStatus = "IMPORT_FAILED"
	ImportStatusImportFailedServerLimitExceeded ImportStatus = "IMPORT_FAILED_SERVER_LIMIT_EXCEEDED"
	ImportStatusImportFailedRecordLimitExceeded ImportStatus = "IMPORT_FAILED_RECORD_LIMIT_EXCEEDED"
	ImportStatusDeleteInProgress                ImportStatus = "DELETE_IN_PROGRESS"
	ImportStatusDeleteComplete                  ImportStatus = "DELETE_COMPLETE"
	ImportStatusDeleteFailed                    ImportStatus = "DELETE_FAILED"
	ImportStatusDeleteFailedLimitExceeded       ImportStatus = "DELETE_FAILED_LIMIT_EXCEEDED"
	ImportStatusInternalError                   ImportStatus = "INTERNAL_ERROR"
)

func (enum ImportStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ImportStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ImportTaskFilterName string

// Enum values for ImportTaskFilterName
const (
	ImportTaskFilterNameImportTaskId ImportTaskFilterName = "IMPORT_TASK_ID"
	ImportTaskFilterNameStatus       ImportTaskFilterName = "STATUS"
	ImportTaskFilterNameName         ImportTaskFilterName = "NAME"
)

func (enum ImportTaskFilterName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ImportTaskFilterName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type OrderString string

// Enum values for OrderString
const (
	OrderStringAsc  OrderString = "ASC"
	OrderStringDesc OrderString = "DESC"
)

func (enum OrderString) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum OrderString) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
