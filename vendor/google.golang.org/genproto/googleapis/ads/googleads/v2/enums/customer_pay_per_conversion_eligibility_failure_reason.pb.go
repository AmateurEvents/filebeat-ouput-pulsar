// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/customer_pay_per_conversion_eligibility_failure_reason.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Enum describing possible reasons a customer is not eligible to use
// PaymentMode.CONVERSIONS.
type CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason int32

const (
	// Not specified.
	CustomerPayPerConversionEligibilityFailureReasonEnum_UNSPECIFIED CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason = 0
	// Used for return value only. Represents value unknown in this version.
	CustomerPayPerConversionEligibilityFailureReasonEnum_UNKNOWN CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason = 1
	// Customer does not have enough conversions.
	CustomerPayPerConversionEligibilityFailureReasonEnum_NOT_ENOUGH_CONVERSIONS CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason = 2
	// Customer's conversion lag is too high.
	CustomerPayPerConversionEligibilityFailureReasonEnum_CONVERSION_LAG_TOO_HIGH CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason = 3
	// Customer uses shared budgets.
	CustomerPayPerConversionEligibilityFailureReasonEnum_HAS_CAMPAIGN_WITH_SHARED_BUDGET CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason = 4
	// Customer has conversions with ConversionActionType.UPLOAD_CLICKS.
	CustomerPayPerConversionEligibilityFailureReasonEnum_HAS_UPLOAD_CLICKS_CONVERSION CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason = 5
	// Customer's average daily spend is too high.
	CustomerPayPerConversionEligibilityFailureReasonEnum_AVERAGE_DAILY_SPEND_TOO_HIGH CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason = 6
	// Customer's eligibility has not yet been calculated by the Google Ads
	// backend. Check back soon.
	CustomerPayPerConversionEligibilityFailureReasonEnum_ANALYSIS_NOT_COMPLETE CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason = 7
	// Customer is not eligible due to other reasons.
	CustomerPayPerConversionEligibilityFailureReasonEnum_OTHER CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason = 8
)

var CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "NOT_ENOUGH_CONVERSIONS",
	3: "CONVERSION_LAG_TOO_HIGH",
	4: "HAS_CAMPAIGN_WITH_SHARED_BUDGET",
	5: "HAS_UPLOAD_CLICKS_CONVERSION",
	6: "AVERAGE_DAILY_SPEND_TOO_HIGH",
	7: "ANALYSIS_NOT_COMPLETE",
	8: "OTHER",
}

var CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason_value = map[string]int32{
	"UNSPECIFIED":                     0,
	"UNKNOWN":                         1,
	"NOT_ENOUGH_CONVERSIONS":          2,
	"CONVERSION_LAG_TOO_HIGH":         3,
	"HAS_CAMPAIGN_WITH_SHARED_BUDGET": 4,
	"HAS_UPLOAD_CLICKS_CONVERSION":    5,
	"AVERAGE_DAILY_SPEND_TOO_HIGH":    6,
	"ANALYSIS_NOT_COMPLETE":           7,
	"OTHER":                           8,
}

func (x CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason) String() string {
	return proto.EnumName(CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason_name, int32(x))
}

func (CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9902b2b64be08123, []int{0, 0}
}

// Container for enum describing reasons why a customer is not eligible to use
// PaymentMode.CONVERSIONS.
type CustomerPayPerConversionEligibilityFailureReasonEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomerPayPerConversionEligibilityFailureReasonEnum) Reset() {
	*m = CustomerPayPerConversionEligibilityFailureReasonEnum{}
}
func (m *CustomerPayPerConversionEligibilityFailureReasonEnum) String() string {
	return proto.CompactTextString(m)
}
func (*CustomerPayPerConversionEligibilityFailureReasonEnum) ProtoMessage() {}
func (*CustomerPayPerConversionEligibilityFailureReasonEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_9902b2b64be08123, []int{0}
}

func (m *CustomerPayPerConversionEligibilityFailureReasonEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerPayPerConversionEligibilityFailureReasonEnum.Unmarshal(m, b)
}
func (m *CustomerPayPerConversionEligibilityFailureReasonEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerPayPerConversionEligibilityFailureReasonEnum.Marshal(b, m, deterministic)
}
func (m *CustomerPayPerConversionEligibilityFailureReasonEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerPayPerConversionEligibilityFailureReasonEnum.Merge(m, src)
}
func (m *CustomerPayPerConversionEligibilityFailureReasonEnum) XXX_Size() int {
	return xxx_messageInfo_CustomerPayPerConversionEligibilityFailureReasonEnum.Size(m)
}
func (m *CustomerPayPerConversionEligibilityFailureReasonEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerPayPerConversionEligibilityFailureReasonEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerPayPerConversionEligibilityFailureReasonEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason", CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason_name, CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason_value)
	proto.RegisterType((*CustomerPayPerConversionEligibilityFailureReasonEnum)(nil), "google.ads.googleads.v2.enums.CustomerPayPerConversionEligibilityFailureReasonEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/customer_pay_per_conversion_eligibility_failure_reason.proto", fileDescriptor_9902b2b64be08123)
}

var fileDescriptor_9902b2b64be08123 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x6f, 0x94, 0x40,
	0x18, 0xc6, 0x2d, 0xb5, 0xad, 0x4e, 0x0f, 0x12, 0x12, 0x35, 0xd6, 0x36, 0xda, 0x7a, 0x07, 0x83,
	0x7a, 0xc1, 0xd3, 0x2c, 0x4c, 0x81, 0x94, 0xce, 0x10, 0xfe, 0x35, 0x6d, 0x36, 0x99, 0xd0, 0x65,
	0x24, 0x24, 0xec, 0x0c, 0x61, 0xd8, 0x4d, 0xf6, 0x2b, 0xf8, 0x09, 0x3c, 0x7b, 0xf4, 0xa3, 0x78,
	0xf3, 0x6b, 0xf8, 0x29, 0x0c, 0xd0, 0xdd, 0xf5, 0xa2, 0xc9, 0x5e, 0xc8, 0x93, 0x79, 0x9f, 0x79,
	0x7e, 0x13, 0x9e, 0x17, 0xdc, 0x95, 0x42, 0x94, 0x35, 0x33, 0xf2, 0x42, 0x1a, 0xa3, 0xec, 0xd5,
	0xd2, 0x34, 0x18, 0x5f, 0xcc, 0xa5, 0x31, 0x5b, 0xc8, 0x4e, 0xcc, 0x59, 0x4b, 0x9b, 0x7c, 0x45,
	0x1b, 0xd6, 0xd2, 0x99, 0xe0, 0x4b, 0xd6, 0xca, 0x4a, 0x70, 0xca, 0xea, 0xaa, 0xac, 0xee, 0xab,
	0xba, 0xea, 0x56, 0xf4, 0x4b, 0x5e, 0xd5, 0x8b, 0x96, 0xd1, 0x96, 0xe5, 0x52, 0x70, 0xbd, 0x69,
	0x45, 0x27, 0xb4, 0xb3, 0x31, 0x50, 0xcf, 0x0b, 0xa9, 0x6f, 0xb2, 0xf5, 0xa5, 0xa9, 0x0f, 0xd9,
	0x27, 0xa7, 0x6b, 0x74, 0x53, 0x19, 0x39, 0xe7, 0xa2, 0xcb, 0xbb, 0x4a, 0x70, 0x39, 0x5e, 0xbe,
	0xf8, 0xa5, 0x80, 0x8f, 0xf6, 0x03, 0x3d, 0xcc, 0x57, 0x21, 0x6b, 0xed, 0x0d, 0x1a, 0x6d, 0xc9,
	0x97, 0x23, 0x38, 0x1a, 0xb8, 0x88, 0x2f, 0xe6, 0x17, 0xdf, 0x14, 0xf0, 0x7e, 0xd7, 0x8b, 0xda,
	0x33, 0x70, 0x9c, 0xe2, 0x38, 0x44, 0xb6, 0x7f, 0xe9, 0x23, 0x47, 0x7d, 0xa4, 0x1d, 0x83, 0xa3,
	0x14, 0x5f, 0x61, 0x72, 0x83, 0xd5, 0x3d, 0xed, 0x04, 0xbc, 0xc0, 0x24, 0xa1, 0x08, 0x93, 0xd4,
	0xf5, 0xa8, 0x4d, 0x70, 0x86, 0xa2, 0xd8, 0x27, 0x38, 0x56, 0x15, 0xed, 0x35, 0x78, 0xb9, 0x3d,
	0xa0, 0x01, 0x74, 0x69, 0x42, 0x08, 0xf5, 0x7c, 0xd7, 0x53, 0xf7, 0xb5, 0x77, 0xe0, 0x8d, 0x07,
	0x63, 0x6a, 0xc3, 0xeb, 0x10, 0xfa, 0x2e, 0xa6, 0x37, 0x7e, 0xe2, 0xd1, 0xd8, 0x83, 0x11, 0x72,
	0xe8, 0x24, 0x75, 0x5c, 0x94, 0xa8, 0x8f, 0xb5, 0xb7, 0xe0, 0xb4, 0x37, 0xa5, 0x61, 0x40, 0xa0,
	0x43, 0xed, 0xc0, 0xb7, 0xaf, 0xe2, 0xbf, 0x20, 0xea, 0x41, 0xef, 0x80, 0x19, 0x8a, 0xa0, 0x8b,
	0xa8, 0x03, 0xfd, 0xe0, 0x96, 0xc6, 0x21, 0xc2, 0xce, 0x16, 0x74, 0xa8, 0xbd, 0x02, 0xcf, 0x21,
	0x86, 0xc1, 0x6d, 0xec, 0xc7, 0xb4, 0x7f, 0xaa, 0x4d, 0xae, 0xc3, 0x00, 0x25, 0x48, 0x3d, 0xd2,
	0x9e, 0x82, 0x03, 0x92, 0x78, 0x28, 0x52, 0x9f, 0x4c, 0xbe, 0x2a, 0xe0, 0x7c, 0x26, 0xe6, 0xfa,
	0x7f, 0x7b, 0x99, 0x7c, 0xda, 0xf5, 0xef, 0x85, 0x7d, 0x61, 0xe1, 0xde, 0xdd, 0xe4, 0x21, 0xb7,
	0x14, 0x75, 0xce, 0x4b, 0x5d, 0xb4, 0xa5, 0x51, 0x32, 0x3e, 0xd4, 0xb9, 0xde, 0xad, 0xa6, 0x92,
	0xff, 0x58, 0xb5, 0xcf, 0xc3, 0xf7, 0xbb, 0xb2, 0xef, 0x42, 0xf8, 0x43, 0x39, 0x73, 0xc7, 0x28,
	0x58, 0x48, 0x7d, 0x94, 0xbd, 0xca, 0x4c, 0xbd, 0xaf, 0x58, 0xfe, 0x5c, 0xcf, 0xa7, 0xb0, 0x90,
	0xd3, 0xcd, 0x7c, 0x9a, 0x99, 0xd3, 0x61, 0xfe, 0x5b, 0x39, 0x1f, 0x0f, 0x2d, 0x0b, 0x16, 0xd2,
	0xb2, 0x36, 0x0e, 0xcb, 0xca, 0x4c, 0xcb, 0x1a, 0x3c, 0xf7, 0x87, 0xc3, 0xc3, 0x3e, 0xfc, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x22, 0x76, 0xfd, 0x64, 0x02, 0x03, 0x00, 0x00,
}
