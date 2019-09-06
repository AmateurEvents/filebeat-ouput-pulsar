// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package textract

type BlockType string

// Enum values for BlockType
const (
	BlockTypeKeyValueSet      BlockType = "KEY_VALUE_SET"
	BlockTypePage             BlockType = "PAGE"
	BlockTypeLine             BlockType = "LINE"
	BlockTypeWord             BlockType = "WORD"
	BlockTypeTable            BlockType = "TABLE"
	BlockTypeCell             BlockType = "CELL"
	BlockTypeSelectionElement BlockType = "SELECTION_ELEMENT"
)

func (enum BlockType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum BlockType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type EntityType string

// Enum values for EntityType
const (
	EntityTypeKey   EntityType = "KEY"
	EntityTypeValue EntityType = "VALUE"
)

func (enum EntityType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum EntityType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type FeatureType string

// Enum values for FeatureType
const (
	FeatureTypeTables FeatureType = "TABLES"
	FeatureTypeForms  FeatureType = "FORMS"
)

func (enum FeatureType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum FeatureType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type JobStatus string

// Enum values for JobStatus
const (
	JobStatusInProgress     JobStatus = "IN_PROGRESS"
	JobStatusSucceeded      JobStatus = "SUCCEEDED"
	JobStatusFailed         JobStatus = "FAILED"
	JobStatusPartialSuccess JobStatus = "PARTIAL_SUCCESS"
)

func (enum JobStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum JobStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RelationshipType string

// Enum values for RelationshipType
const (
	RelationshipTypeValue RelationshipType = "VALUE"
	RelationshipTypeChild RelationshipType = "CHILD"
)

func (enum RelationshipType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RelationshipType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SelectionStatus string

// Enum values for SelectionStatus
const (
	SelectionStatusSelected    SelectionStatus = "SELECTED"
	SelectionStatusNotSelected SelectionStatus = "NOT_SELECTED"
)

func (enum SelectionStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SelectionStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
