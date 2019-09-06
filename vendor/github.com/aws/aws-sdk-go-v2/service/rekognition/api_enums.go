// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

type Attribute string

// Enum values for Attribute
const (
	AttributeDefault Attribute = "DEFAULT"
	AttributeAll     Attribute = "ALL"
)

func (enum Attribute) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Attribute) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CelebrityRecognitionSortBy string

// Enum values for CelebrityRecognitionSortBy
const (
	CelebrityRecognitionSortById        CelebrityRecognitionSortBy = "ID"
	CelebrityRecognitionSortByTimestamp CelebrityRecognitionSortBy = "TIMESTAMP"
)

func (enum CelebrityRecognitionSortBy) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CelebrityRecognitionSortBy) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ContentModerationSortBy string

// Enum values for ContentModerationSortBy
const (
	ContentModerationSortByName      ContentModerationSortBy = "NAME"
	ContentModerationSortByTimestamp ContentModerationSortBy = "TIMESTAMP"
)

func (enum ContentModerationSortBy) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ContentModerationSortBy) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type EmotionName string

// Enum values for EmotionName
const (
	EmotionNameHappy     EmotionName = "HAPPY"
	EmotionNameSad       EmotionName = "SAD"
	EmotionNameAngry     EmotionName = "ANGRY"
	EmotionNameConfused  EmotionName = "CONFUSED"
	EmotionNameDisgusted EmotionName = "DISGUSTED"
	EmotionNameSurprised EmotionName = "SURPRISED"
	EmotionNameCalm      EmotionName = "CALM"
	EmotionNameUnknown   EmotionName = "UNKNOWN"
	EmotionNameFear      EmotionName = "FEAR"
)

func (enum EmotionName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum EmotionName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type FaceAttributes string

// Enum values for FaceAttributes
const (
	FaceAttributesDefault FaceAttributes = "DEFAULT"
	FaceAttributesAll     FaceAttributes = "ALL"
)

func (enum FaceAttributes) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum FaceAttributes) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type FaceSearchSortBy string

// Enum values for FaceSearchSortBy
const (
	FaceSearchSortByIndex     FaceSearchSortBy = "INDEX"
	FaceSearchSortByTimestamp FaceSearchSortBy = "TIMESTAMP"
)

func (enum FaceSearchSortBy) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum FaceSearchSortBy) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type GenderType string

// Enum values for GenderType
const (
	GenderTypeMale   GenderType = "Male"
	GenderTypeFemale GenderType = "Female"
)

func (enum GenderType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum GenderType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type LabelDetectionSortBy string

// Enum values for LabelDetectionSortBy
const (
	LabelDetectionSortByName      LabelDetectionSortBy = "NAME"
	LabelDetectionSortByTimestamp LabelDetectionSortBy = "TIMESTAMP"
)

func (enum LabelDetectionSortBy) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum LabelDetectionSortBy) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type LandmarkType string

// Enum values for LandmarkType
const (
	LandmarkTypeEyeLeft           LandmarkType = "eyeLeft"
	LandmarkTypeEyeRight          LandmarkType = "eyeRight"
	LandmarkTypeNose              LandmarkType = "nose"
	LandmarkTypeMouthLeft         LandmarkType = "mouthLeft"
	LandmarkTypeMouthRight        LandmarkType = "mouthRight"
	LandmarkTypeLeftEyeBrowLeft   LandmarkType = "leftEyeBrowLeft"
	LandmarkTypeLeftEyeBrowRight  LandmarkType = "leftEyeBrowRight"
	LandmarkTypeLeftEyeBrowUp     LandmarkType = "leftEyeBrowUp"
	LandmarkTypeRightEyeBrowLeft  LandmarkType = "rightEyeBrowLeft"
	LandmarkTypeRightEyeBrowRight LandmarkType = "rightEyeBrowRight"
	LandmarkTypeRightEyeBrowUp    LandmarkType = "rightEyeBrowUp"
	LandmarkTypeLeftEyeLeft       LandmarkType = "leftEyeLeft"
	LandmarkTypeLeftEyeRight      LandmarkType = "leftEyeRight"
	LandmarkTypeLeftEyeUp         LandmarkType = "leftEyeUp"
	LandmarkTypeLeftEyeDown       LandmarkType = "leftEyeDown"
	LandmarkTypeRightEyeLeft      LandmarkType = "rightEyeLeft"
	LandmarkTypeRightEyeRight     LandmarkType = "rightEyeRight"
	LandmarkTypeRightEyeUp        LandmarkType = "rightEyeUp"
	LandmarkTypeRightEyeDown      LandmarkType = "rightEyeDown"
	LandmarkTypeNoseLeft          LandmarkType = "noseLeft"
	LandmarkTypeNoseRight         LandmarkType = "noseRight"
	LandmarkTypeMouthUp           LandmarkType = "mouthUp"
	LandmarkTypeMouthDown         LandmarkType = "mouthDown"
	LandmarkTypeLeftPupil         LandmarkType = "leftPupil"
	LandmarkTypeRightPupil        LandmarkType = "rightPupil"
	LandmarkTypeUpperJawlineLeft  LandmarkType = "upperJawlineLeft"
	LandmarkTypeMidJawlineLeft    LandmarkType = "midJawlineLeft"
	LandmarkTypeChinBottom        LandmarkType = "chinBottom"
	LandmarkTypeMidJawlineRight   LandmarkType = "midJawlineRight"
	LandmarkTypeUpperJawlineRight LandmarkType = "upperJawlineRight"
)

func (enum LandmarkType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum LandmarkType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type OrientationCorrection string

// Enum values for OrientationCorrection
const (
	OrientationCorrectionRotate0   OrientationCorrection = "ROTATE_0"
	OrientationCorrectionRotate90  OrientationCorrection = "ROTATE_90"
	OrientationCorrectionRotate180 OrientationCorrection = "ROTATE_180"
	OrientationCorrectionRotate270 OrientationCorrection = "ROTATE_270"
)

func (enum OrientationCorrection) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum OrientationCorrection) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PersonTrackingSortBy string

// Enum values for PersonTrackingSortBy
const (
	PersonTrackingSortByIndex     PersonTrackingSortBy = "INDEX"
	PersonTrackingSortByTimestamp PersonTrackingSortBy = "TIMESTAMP"
)

func (enum PersonTrackingSortBy) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PersonTrackingSortBy) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type QualityFilter string

// Enum values for QualityFilter
const (
	QualityFilterNone QualityFilter = "NONE"
	QualityFilterAuto QualityFilter = "AUTO"
)

func (enum QualityFilter) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum QualityFilter) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Reason string

// Enum values for Reason
const (
	ReasonExceedsMaxFaces  Reason = "EXCEEDS_MAX_FACES"
	ReasonExtremePose      Reason = "EXTREME_POSE"
	ReasonLowBrightness    Reason = "LOW_BRIGHTNESS"
	ReasonLowSharpness     Reason = "LOW_SHARPNESS"
	ReasonLowConfidence    Reason = "LOW_CONFIDENCE"
	ReasonSmallBoundingBox Reason = "SMALL_BOUNDING_BOX"
)

func (enum Reason) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Reason) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type StreamProcessorStatus string

// Enum values for StreamProcessorStatus
const (
	StreamProcessorStatusStopped  StreamProcessorStatus = "STOPPED"
	StreamProcessorStatusStarting StreamProcessorStatus = "STARTING"
	StreamProcessorStatusRunning  StreamProcessorStatus = "RUNNING"
	StreamProcessorStatusFailed   StreamProcessorStatus = "FAILED"
	StreamProcessorStatusStopping StreamProcessorStatus = "STOPPING"
)

func (enum StreamProcessorStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum StreamProcessorStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TextTypes string

// Enum values for TextTypes
const (
	TextTypesLine TextTypes = "LINE"
	TextTypesWord TextTypes = "WORD"
)

func (enum TextTypes) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TextTypes) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type VideoJobStatus string

// Enum values for VideoJobStatus
const (
	VideoJobStatusInProgress VideoJobStatus = "IN_PROGRESS"
	VideoJobStatusSucceeded  VideoJobStatus = "SUCCEEDED"
	VideoJobStatusFailed     VideoJobStatus = "FAILED"
)

func (enum VideoJobStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum VideoJobStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
