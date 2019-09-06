// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

type ArtifactCategory string

// Enum values for ArtifactCategory
const (
	ArtifactCategoryScreenshot ArtifactCategory = "SCREENSHOT"
	ArtifactCategoryFile       ArtifactCategory = "FILE"
	ArtifactCategoryLog        ArtifactCategory = "LOG"
)

func (enum ArtifactCategory) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ArtifactCategory) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ArtifactType string

// Enum values for ArtifactType
const (
	ArtifactTypeUnknown                ArtifactType = "UNKNOWN"
	ArtifactTypeScreenshot             ArtifactType = "SCREENSHOT"
	ArtifactTypeDeviceLog              ArtifactType = "DEVICE_LOG"
	ArtifactTypeMessageLog             ArtifactType = "MESSAGE_LOG"
	ArtifactTypeVideoLog               ArtifactType = "VIDEO_LOG"
	ArtifactTypeResultLog              ArtifactType = "RESULT_LOG"
	ArtifactTypeServiceLog             ArtifactType = "SERVICE_LOG"
	ArtifactTypeWebkitLog              ArtifactType = "WEBKIT_LOG"
	ArtifactTypeInstrumentationOutput  ArtifactType = "INSTRUMENTATION_OUTPUT"
	ArtifactTypeExerciserMonkeyOutput  ArtifactType = "EXERCISER_MONKEY_OUTPUT"
	ArtifactTypeCalabashJsonOutput     ArtifactType = "CALABASH_JSON_OUTPUT"
	ArtifactTypeCalabashPrettyOutput   ArtifactType = "CALABASH_PRETTY_OUTPUT"
	ArtifactTypeCalabashStandardOutput ArtifactType = "CALABASH_STANDARD_OUTPUT"
	ArtifactTypeCalabashJavaXmlOutput  ArtifactType = "CALABASH_JAVA_XML_OUTPUT"
	ArtifactTypeAutomationOutput       ArtifactType = "AUTOMATION_OUTPUT"
	ArtifactTypeAppiumServerOutput     ArtifactType = "APPIUM_SERVER_OUTPUT"
	ArtifactTypeAppiumJavaOutput       ArtifactType = "APPIUM_JAVA_OUTPUT"
	ArtifactTypeAppiumJavaXmlOutput    ArtifactType = "APPIUM_JAVA_XML_OUTPUT"
	ArtifactTypeAppiumPythonOutput     ArtifactType = "APPIUM_PYTHON_OUTPUT"
	ArtifactTypeAppiumPythonXmlOutput  ArtifactType = "APPIUM_PYTHON_XML_OUTPUT"
	ArtifactTypeExplorerEventLog       ArtifactType = "EXPLORER_EVENT_LOG"
	ArtifactTypeExplorerSummaryLog     ArtifactType = "EXPLORER_SUMMARY_LOG"
	ArtifactTypeApplicationCrashReport ArtifactType = "APPLICATION_CRASH_REPORT"
	ArtifactTypeXctestLog              ArtifactType = "XCTEST_LOG"
	ArtifactTypeVideo                  ArtifactType = "VIDEO"
	ArtifactTypeCustomerArtifact       ArtifactType = "CUSTOMER_ARTIFACT"
	ArtifactTypeCustomerArtifactLog    ArtifactType = "CUSTOMER_ARTIFACT_LOG"
	ArtifactTypeTestspecOutput         ArtifactType = "TESTSPEC_OUTPUT"
)

func (enum ArtifactType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ArtifactType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type BillingMethod string

// Enum values for BillingMethod
const (
	BillingMethodMetered   BillingMethod = "METERED"
	BillingMethodUnmetered BillingMethod = "UNMETERED"
)

func (enum BillingMethod) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum BillingMethod) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CurrencyCode string

// Enum values for CurrencyCode
const (
	CurrencyCodeUsd CurrencyCode = "USD"
)

func (enum CurrencyCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CurrencyCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DeviceAttribute string

// Enum values for DeviceAttribute
const (
	DeviceAttributeArn                 DeviceAttribute = "ARN"
	DeviceAttributePlatform            DeviceAttribute = "PLATFORM"
	DeviceAttributeFormFactor          DeviceAttribute = "FORM_FACTOR"
	DeviceAttributeManufacturer        DeviceAttribute = "MANUFACTURER"
	DeviceAttributeRemoteAccessEnabled DeviceAttribute = "REMOTE_ACCESS_ENABLED"
	DeviceAttributeRemoteDebugEnabled  DeviceAttribute = "REMOTE_DEBUG_ENABLED"
	DeviceAttributeAppiumVersion       DeviceAttribute = "APPIUM_VERSION"
	DeviceAttributeInstanceArn         DeviceAttribute = "INSTANCE_ARN"
	DeviceAttributeInstanceLabels      DeviceAttribute = "INSTANCE_LABELS"
	DeviceAttributeFleetType           DeviceAttribute = "FLEET_TYPE"
	DeviceAttributeOsVersion           DeviceAttribute = "OS_VERSION"
	DeviceAttributeModel               DeviceAttribute = "MODEL"
	DeviceAttributeAvailability        DeviceAttribute = "AVAILABILITY"
)

func (enum DeviceAttribute) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DeviceAttribute) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DeviceAvailability string

// Enum values for DeviceAvailability
const (
	DeviceAvailabilityTemporaryNotAvailable DeviceAvailability = "TEMPORARY_NOT_AVAILABLE"
	DeviceAvailabilityBusy                  DeviceAvailability = "BUSY"
	DeviceAvailabilityAvailable             DeviceAvailability = "AVAILABLE"
	DeviceAvailabilityHighlyAvailable       DeviceAvailability = "HIGHLY_AVAILABLE"
)

func (enum DeviceAvailability) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DeviceAvailability) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DeviceFilterAttribute string

// Enum values for DeviceFilterAttribute
const (
	DeviceFilterAttributeArn                 DeviceFilterAttribute = "ARN"
	DeviceFilterAttributePlatform            DeviceFilterAttribute = "PLATFORM"
	DeviceFilterAttributeOsVersion           DeviceFilterAttribute = "OS_VERSION"
	DeviceFilterAttributeModel               DeviceFilterAttribute = "MODEL"
	DeviceFilterAttributeAvailability        DeviceFilterAttribute = "AVAILABILITY"
	DeviceFilterAttributeFormFactor          DeviceFilterAttribute = "FORM_FACTOR"
	DeviceFilterAttributeManufacturer        DeviceFilterAttribute = "MANUFACTURER"
	DeviceFilterAttributeRemoteAccessEnabled DeviceFilterAttribute = "REMOTE_ACCESS_ENABLED"
	DeviceFilterAttributeRemoteDebugEnabled  DeviceFilterAttribute = "REMOTE_DEBUG_ENABLED"
	DeviceFilterAttributeInstanceArn         DeviceFilterAttribute = "INSTANCE_ARN"
	DeviceFilterAttributeInstanceLabels      DeviceFilterAttribute = "INSTANCE_LABELS"
	DeviceFilterAttributeFleetType           DeviceFilterAttribute = "FLEET_TYPE"
)

func (enum DeviceFilterAttribute) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DeviceFilterAttribute) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DeviceFormFactor string

// Enum values for DeviceFormFactor
const (
	DeviceFormFactorPhone  DeviceFormFactor = "PHONE"
	DeviceFormFactorTablet DeviceFormFactor = "TABLET"
)

func (enum DeviceFormFactor) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DeviceFormFactor) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DevicePlatform string

// Enum values for DevicePlatform
const (
	DevicePlatformAndroid DevicePlatform = "ANDROID"
	DevicePlatformIos     DevicePlatform = "IOS"
)

func (enum DevicePlatform) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DevicePlatform) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DevicePoolType string

// Enum values for DevicePoolType
const (
	DevicePoolTypeCurated DevicePoolType = "CURATED"
	DevicePoolTypePrivate DevicePoolType = "PRIVATE"
)

func (enum DevicePoolType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DevicePoolType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ExecutionResult string

// Enum values for ExecutionResult
const (
	ExecutionResultPending ExecutionResult = "PENDING"
	ExecutionResultPassed  ExecutionResult = "PASSED"
	ExecutionResultWarned  ExecutionResult = "WARNED"
	ExecutionResultFailed  ExecutionResult = "FAILED"
	ExecutionResultSkipped ExecutionResult = "SKIPPED"
	ExecutionResultErrored ExecutionResult = "ERRORED"
	ExecutionResultStopped ExecutionResult = "STOPPED"
)

func (enum ExecutionResult) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ExecutionResult) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ExecutionResultCode string

// Enum values for ExecutionResultCode
const (
	ExecutionResultCodeParsingFailed          ExecutionResultCode = "PARSING_FAILED"
	ExecutionResultCodeVpcEndpointSetupFailed ExecutionResultCode = "VPC_ENDPOINT_SETUP_FAILED"
)

func (enum ExecutionResultCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ExecutionResultCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ExecutionStatus string

// Enum values for ExecutionStatus
const (
	ExecutionStatusPending            ExecutionStatus = "PENDING"
	ExecutionStatusPendingConcurrency ExecutionStatus = "PENDING_CONCURRENCY"
	ExecutionStatusPendingDevice      ExecutionStatus = "PENDING_DEVICE"
	ExecutionStatusProcessing         ExecutionStatus = "PROCESSING"
	ExecutionStatusScheduling         ExecutionStatus = "SCHEDULING"
	ExecutionStatusPreparing          ExecutionStatus = "PREPARING"
	ExecutionStatusRunning            ExecutionStatus = "RUNNING"
	ExecutionStatusCompleted          ExecutionStatus = "COMPLETED"
	ExecutionStatusStopping           ExecutionStatus = "STOPPING"
)

func (enum ExecutionStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ExecutionStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type InstanceStatus string

// Enum values for InstanceStatus
const (
	InstanceStatusInUse        InstanceStatus = "IN_USE"
	InstanceStatusPreparing    InstanceStatus = "PREPARING"
	InstanceStatusAvailable    InstanceStatus = "AVAILABLE"
	InstanceStatusNotAvailable InstanceStatus = "NOT_AVAILABLE"
)

func (enum InstanceStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum InstanceStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type InteractionMode string

// Enum values for InteractionMode
const (
	InteractionModeInteractive InteractionMode = "INTERACTIVE"
	InteractionModeNoVideo     InteractionMode = "NO_VIDEO"
	InteractionModeVideoOnly   InteractionMode = "VIDEO_ONLY"
)

func (enum InteractionMode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum InteractionMode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type NetworkProfileType string

// Enum values for NetworkProfileType
const (
	NetworkProfileTypeCurated NetworkProfileType = "CURATED"
	NetworkProfileTypePrivate NetworkProfileType = "PRIVATE"
)

func (enum NetworkProfileType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum NetworkProfileType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type OfferingTransactionType string

// Enum values for OfferingTransactionType
const (
	OfferingTransactionTypePurchase OfferingTransactionType = "PURCHASE"
	OfferingTransactionTypeRenew    OfferingTransactionType = "RENEW"
	OfferingTransactionTypeSystem   OfferingTransactionType = "SYSTEM"
)

func (enum OfferingTransactionType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum OfferingTransactionType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type OfferingType string

// Enum values for OfferingType
const (
	OfferingTypeRecurring OfferingType = "RECURRING"
)

func (enum OfferingType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum OfferingType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RecurringChargeFrequency string

// Enum values for RecurringChargeFrequency
const (
	RecurringChargeFrequencyMonthly RecurringChargeFrequency = "MONTHLY"
)

func (enum RecurringChargeFrequency) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RecurringChargeFrequency) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RuleOperator string

// Enum values for RuleOperator
const (
	RuleOperatorEquals              RuleOperator = "EQUALS"
	RuleOperatorLessThan            RuleOperator = "LESS_THAN"
	RuleOperatorLessThanOrEquals    RuleOperator = "LESS_THAN_OR_EQUALS"
	RuleOperatorGreaterThan         RuleOperator = "GREATER_THAN"
	RuleOperatorGreaterThanOrEquals RuleOperator = "GREATER_THAN_OR_EQUALS"
	RuleOperatorIn                  RuleOperator = "IN"
	RuleOperatorNotIn               RuleOperator = "NOT_IN"
	RuleOperatorContains            RuleOperator = "CONTAINS"
)

func (enum RuleOperator) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RuleOperator) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SampleType string

// Enum values for SampleType
const (
	SampleTypeCpu               SampleType = "CPU"
	SampleTypeMemory            SampleType = "MEMORY"
	SampleTypeThreads           SampleType = "THREADS"
	SampleTypeRxRate            SampleType = "RX_RATE"
	SampleTypeTxRate            SampleType = "TX_RATE"
	SampleTypeRx                SampleType = "RX"
	SampleTypeTx                SampleType = "TX"
	SampleTypeNativeFrames      SampleType = "NATIVE_FRAMES"
	SampleTypeNativeFps         SampleType = "NATIVE_FPS"
	SampleTypeNativeMinDrawtime SampleType = "NATIVE_MIN_DRAWTIME"
	SampleTypeNativeAvgDrawtime SampleType = "NATIVE_AVG_DRAWTIME"
	SampleTypeNativeMaxDrawtime SampleType = "NATIVE_MAX_DRAWTIME"
	SampleTypeOpenglFrames      SampleType = "OPENGL_FRAMES"
	SampleTypeOpenglFps         SampleType = "OPENGL_FPS"
	SampleTypeOpenglMinDrawtime SampleType = "OPENGL_MIN_DRAWTIME"
	SampleTypeOpenglAvgDrawtime SampleType = "OPENGL_AVG_DRAWTIME"
	SampleTypeOpenglMaxDrawtime SampleType = "OPENGL_MAX_DRAWTIME"
)

func (enum SampleType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SampleType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TestType string

// Enum values for TestType
const (
	TestTypeBuiltinFuzz           TestType = "BUILTIN_FUZZ"
	TestTypeBuiltinExplorer       TestType = "BUILTIN_EXPLORER"
	TestTypeWebPerformanceProfile TestType = "WEB_PERFORMANCE_PROFILE"
	TestTypeAppiumJavaJunit       TestType = "APPIUM_JAVA_JUNIT"
	TestTypeAppiumJavaTestng      TestType = "APPIUM_JAVA_TESTNG"
	TestTypeAppiumPython          TestType = "APPIUM_PYTHON"
	TestTypeAppiumNode            TestType = "APPIUM_NODE"
	TestTypeAppiumRuby            TestType = "APPIUM_RUBY"
	TestTypeAppiumWebJavaJunit    TestType = "APPIUM_WEB_JAVA_JUNIT"
	TestTypeAppiumWebJavaTestng   TestType = "APPIUM_WEB_JAVA_TESTNG"
	TestTypeAppiumWebPython       TestType = "APPIUM_WEB_PYTHON"
	TestTypeAppiumWebNode         TestType = "APPIUM_WEB_NODE"
	TestTypeAppiumWebRuby         TestType = "APPIUM_WEB_RUBY"
	TestTypeCalabash              TestType = "CALABASH"
	TestTypeInstrumentation       TestType = "INSTRUMENTATION"
	TestTypeUiautomation          TestType = "UIAUTOMATION"
	TestTypeUiautomator           TestType = "UIAUTOMATOR"
	TestTypeXctest                TestType = "XCTEST"
	TestTypeXctestUi              TestType = "XCTEST_UI"
	TestTypeRemoteAccessRecord    TestType = "REMOTE_ACCESS_RECORD"
	TestTypeRemoteAccessReplay    TestType = "REMOTE_ACCESS_REPLAY"
)

func (enum TestType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TestType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type UploadCategory string

// Enum values for UploadCategory
const (
	UploadCategoryCurated UploadCategory = "CURATED"
	UploadCategoryPrivate UploadCategory = "PRIVATE"
)

func (enum UploadCategory) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum UploadCategory) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type UploadStatus string

// Enum values for UploadStatus
const (
	UploadStatusInitialized UploadStatus = "INITIALIZED"
	UploadStatusProcessing  UploadStatus = "PROCESSING"
	UploadStatusSucceeded   UploadStatus = "SUCCEEDED"
	UploadStatusFailed      UploadStatus = "FAILED"
)

func (enum UploadStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum UploadStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type UploadType string

// Enum values for UploadType
const (
	UploadTypeAndroidApp                     UploadType = "ANDROID_APP"
	UploadTypeIosApp                         UploadType = "IOS_APP"
	UploadTypeWebApp                         UploadType = "WEB_APP"
	UploadTypeExternalData                   UploadType = "EXTERNAL_DATA"
	UploadTypeAppiumJavaJunitTestPackage     UploadType = "APPIUM_JAVA_JUNIT_TEST_PACKAGE"
	UploadTypeAppiumJavaTestngTestPackage    UploadType = "APPIUM_JAVA_TESTNG_TEST_PACKAGE"
	UploadTypeAppiumPythonTestPackage        UploadType = "APPIUM_PYTHON_TEST_PACKAGE"
	UploadTypeAppiumNodeTestPackage          UploadType = "APPIUM_NODE_TEST_PACKAGE"
	UploadTypeAppiumRubyTestPackage          UploadType = "APPIUM_RUBY_TEST_PACKAGE"
	UploadTypeAppiumWebJavaJunitTestPackage  UploadType = "APPIUM_WEB_JAVA_JUNIT_TEST_PACKAGE"
	UploadTypeAppiumWebJavaTestngTestPackage UploadType = "APPIUM_WEB_JAVA_TESTNG_TEST_PACKAGE"
	UploadTypeAppiumWebPythonTestPackage     UploadType = "APPIUM_WEB_PYTHON_TEST_PACKAGE"
	UploadTypeAppiumWebNodeTestPackage       UploadType = "APPIUM_WEB_NODE_TEST_PACKAGE"
	UploadTypeAppiumWebRubyTestPackage       UploadType = "APPIUM_WEB_RUBY_TEST_PACKAGE"
	UploadTypeCalabashTestPackage            UploadType = "CALABASH_TEST_PACKAGE"
	UploadTypeInstrumentationTestPackage     UploadType = "INSTRUMENTATION_TEST_PACKAGE"
	UploadTypeUiautomationTestPackage        UploadType = "UIAUTOMATION_TEST_PACKAGE"
	UploadTypeUiautomatorTestPackage         UploadType = "UIAUTOMATOR_TEST_PACKAGE"
	UploadTypeXctestTestPackage              UploadType = "XCTEST_TEST_PACKAGE"
	UploadTypeXctestUiTestPackage            UploadType = "XCTEST_UI_TEST_PACKAGE"
	UploadTypeAppiumJavaJunitTestSpec        UploadType = "APPIUM_JAVA_JUNIT_TEST_SPEC"
	UploadTypeAppiumJavaTestngTestSpec       UploadType = "APPIUM_JAVA_TESTNG_TEST_SPEC"
	UploadTypeAppiumPythonTestSpec           UploadType = "APPIUM_PYTHON_TEST_SPEC"
	UploadTypeAppiumNodeTestSpec             UploadType = "APPIUM_NODE_TEST_SPEC"
	UploadTypeAppiumRubyTestSpec             UploadType = "APPIUM_RUBY_TEST_SPEC"
	UploadTypeAppiumWebJavaJunitTestSpec     UploadType = "APPIUM_WEB_JAVA_JUNIT_TEST_SPEC"
	UploadTypeAppiumWebJavaTestngTestSpec    UploadType = "APPIUM_WEB_JAVA_TESTNG_TEST_SPEC"
	UploadTypeAppiumWebPythonTestSpec        UploadType = "APPIUM_WEB_PYTHON_TEST_SPEC"
	UploadTypeAppiumWebNodeTestSpec          UploadType = "APPIUM_WEB_NODE_TEST_SPEC"
	UploadTypeAppiumWebRubyTestSpec          UploadType = "APPIUM_WEB_RUBY_TEST_SPEC"
	UploadTypeInstrumentationTestSpec        UploadType = "INSTRUMENTATION_TEST_SPEC"
	UploadTypeXctestUiTestSpec               UploadType = "XCTEST_UI_TEST_SPEC"
)

func (enum UploadType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum UploadType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
