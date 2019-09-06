// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/campaign_extension_setting_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
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

// Request message for
// [CampaignExtensionSettingService.GetCampaignExtensionSetting][google.ads.googleads.v2.services.CampaignExtensionSettingService.GetCampaignExtensionSetting].
type GetCampaignExtensionSettingRequest struct {
	// The resource name of the campaign extension setting to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCampaignExtensionSettingRequest) Reset()         { *m = GetCampaignExtensionSettingRequest{} }
func (m *GetCampaignExtensionSettingRequest) String() string { return proto.CompactTextString(m) }
func (*GetCampaignExtensionSettingRequest) ProtoMessage()    {}
func (*GetCampaignExtensionSettingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_49140070a305e12a, []int{0}
}

func (m *GetCampaignExtensionSettingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCampaignExtensionSettingRequest.Unmarshal(m, b)
}
func (m *GetCampaignExtensionSettingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCampaignExtensionSettingRequest.Marshal(b, m, deterministic)
}
func (m *GetCampaignExtensionSettingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCampaignExtensionSettingRequest.Merge(m, src)
}
func (m *GetCampaignExtensionSettingRequest) XXX_Size() int {
	return xxx_messageInfo_GetCampaignExtensionSettingRequest.Size(m)
}
func (m *GetCampaignExtensionSettingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCampaignExtensionSettingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCampaignExtensionSettingRequest proto.InternalMessageInfo

func (m *GetCampaignExtensionSettingRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for
// [CampaignExtensionSettingService.MutateCampaignExtensionSettings][google.ads.googleads.v2.services.CampaignExtensionSettingService.MutateCampaignExtensionSettings].
type MutateCampaignExtensionSettingsRequest struct {
	// The ID of the customer whose campaign extension settings are being
	// modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual campaign extension
	// settings.
	Operations []*CampaignExtensionSettingOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignExtensionSettingsRequest) Reset() {
	*m = MutateCampaignExtensionSettingsRequest{}
}
func (m *MutateCampaignExtensionSettingsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignExtensionSettingsRequest) ProtoMessage()    {}
func (*MutateCampaignExtensionSettingsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_49140070a305e12a, []int{1}
}

func (m *MutateCampaignExtensionSettingsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignExtensionSettingsRequest.Unmarshal(m, b)
}
func (m *MutateCampaignExtensionSettingsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignExtensionSettingsRequest.Marshal(b, m, deterministic)
}
func (m *MutateCampaignExtensionSettingsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignExtensionSettingsRequest.Merge(m, src)
}
func (m *MutateCampaignExtensionSettingsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignExtensionSettingsRequest.Size(m)
}
func (m *MutateCampaignExtensionSettingsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignExtensionSettingsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignExtensionSettingsRequest proto.InternalMessageInfo

func (m *MutateCampaignExtensionSettingsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCampaignExtensionSettingsRequest) GetOperations() []*CampaignExtensionSettingOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateCampaignExtensionSettingsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateCampaignExtensionSettingsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on a campaign extension setting.
type CampaignExtensionSettingOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*CampaignExtensionSettingOperation_Create
	//	*CampaignExtensionSettingOperation_Update
	//	*CampaignExtensionSettingOperation_Remove
	Operation            isCampaignExtensionSettingOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *CampaignExtensionSettingOperation) Reset()         { *m = CampaignExtensionSettingOperation{} }
func (m *CampaignExtensionSettingOperation) String() string { return proto.CompactTextString(m) }
func (*CampaignExtensionSettingOperation) ProtoMessage()    {}
func (*CampaignExtensionSettingOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_49140070a305e12a, []int{2}
}

func (m *CampaignExtensionSettingOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignExtensionSettingOperation.Unmarshal(m, b)
}
func (m *CampaignExtensionSettingOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignExtensionSettingOperation.Marshal(b, m, deterministic)
}
func (m *CampaignExtensionSettingOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignExtensionSettingOperation.Merge(m, src)
}
func (m *CampaignExtensionSettingOperation) XXX_Size() int {
	return xxx_messageInfo_CampaignExtensionSettingOperation.Size(m)
}
func (m *CampaignExtensionSettingOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignExtensionSettingOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignExtensionSettingOperation proto.InternalMessageInfo

func (m *CampaignExtensionSettingOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isCampaignExtensionSettingOperation_Operation interface {
	isCampaignExtensionSettingOperation_Operation()
}

type CampaignExtensionSettingOperation_Create struct {
	Create *resources.CampaignExtensionSetting `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CampaignExtensionSettingOperation_Update struct {
	Update *resources.CampaignExtensionSetting `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type CampaignExtensionSettingOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*CampaignExtensionSettingOperation_Create) isCampaignExtensionSettingOperation_Operation() {}

func (*CampaignExtensionSettingOperation_Update) isCampaignExtensionSettingOperation_Operation() {}

func (*CampaignExtensionSettingOperation_Remove) isCampaignExtensionSettingOperation_Operation() {}

func (m *CampaignExtensionSettingOperation) GetOperation() isCampaignExtensionSettingOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *CampaignExtensionSettingOperation) GetCreate() *resources.CampaignExtensionSetting {
	if x, ok := m.GetOperation().(*CampaignExtensionSettingOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *CampaignExtensionSettingOperation) GetUpdate() *resources.CampaignExtensionSetting {
	if x, ok := m.GetOperation().(*CampaignExtensionSettingOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *CampaignExtensionSettingOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*CampaignExtensionSettingOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CampaignExtensionSettingOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CampaignExtensionSettingOperation_Create)(nil),
		(*CampaignExtensionSettingOperation_Update)(nil),
		(*CampaignExtensionSettingOperation_Remove)(nil),
	}
}

// Response message for a campaign extension setting mutate.
type MutateCampaignExtensionSettingsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateCampaignExtensionSettingResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *MutateCampaignExtensionSettingsResponse) Reset() {
	*m = MutateCampaignExtensionSettingsResponse{}
}
func (m *MutateCampaignExtensionSettingsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignExtensionSettingsResponse) ProtoMessage()    {}
func (*MutateCampaignExtensionSettingsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_49140070a305e12a, []int{3}
}

func (m *MutateCampaignExtensionSettingsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignExtensionSettingsResponse.Unmarshal(m, b)
}
func (m *MutateCampaignExtensionSettingsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignExtensionSettingsResponse.Marshal(b, m, deterministic)
}
func (m *MutateCampaignExtensionSettingsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignExtensionSettingsResponse.Merge(m, src)
}
func (m *MutateCampaignExtensionSettingsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignExtensionSettingsResponse.Size(m)
}
func (m *MutateCampaignExtensionSettingsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignExtensionSettingsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignExtensionSettingsResponse proto.InternalMessageInfo

func (m *MutateCampaignExtensionSettingsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateCampaignExtensionSettingsResponse) GetResults() []*MutateCampaignExtensionSettingResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the campaign extension setting mutate.
type MutateCampaignExtensionSettingResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignExtensionSettingResult) Reset()         { *m = MutateCampaignExtensionSettingResult{} }
func (m *MutateCampaignExtensionSettingResult) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignExtensionSettingResult) ProtoMessage()    {}
func (*MutateCampaignExtensionSettingResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_49140070a305e12a, []int{4}
}

func (m *MutateCampaignExtensionSettingResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignExtensionSettingResult.Unmarshal(m, b)
}
func (m *MutateCampaignExtensionSettingResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignExtensionSettingResult.Marshal(b, m, deterministic)
}
func (m *MutateCampaignExtensionSettingResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignExtensionSettingResult.Merge(m, src)
}
func (m *MutateCampaignExtensionSettingResult) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignExtensionSettingResult.Size(m)
}
func (m *MutateCampaignExtensionSettingResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignExtensionSettingResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignExtensionSettingResult proto.InternalMessageInfo

func (m *MutateCampaignExtensionSettingResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCampaignExtensionSettingRequest)(nil), "google.ads.googleads.v2.services.GetCampaignExtensionSettingRequest")
	proto.RegisterType((*MutateCampaignExtensionSettingsRequest)(nil), "google.ads.googleads.v2.services.MutateCampaignExtensionSettingsRequest")
	proto.RegisterType((*CampaignExtensionSettingOperation)(nil), "google.ads.googleads.v2.services.CampaignExtensionSettingOperation")
	proto.RegisterType((*MutateCampaignExtensionSettingsResponse)(nil), "google.ads.googleads.v2.services.MutateCampaignExtensionSettingsResponse")
	proto.RegisterType((*MutateCampaignExtensionSettingResult)(nil), "google.ads.googleads.v2.services.MutateCampaignExtensionSettingResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/campaign_extension_setting_service.proto", fileDescriptor_49140070a305e12a)
}

var fileDescriptor_49140070a305e12a = []byte{
	// 738 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xc7, 0x7f, 0x76, 0xab, 0xfe, 0xe8, 0xa6, 0x80, 0xb4, 0x08, 0x61, 0xa5, 0x48, 0x0d, 0x26,
	0xa2, 0x55, 0x0e, 0xb6, 0x64, 0x6e, 0x8e, 0x8a, 0x94, 0x84, 0xa6, 0x8d, 0x50, 0x69, 0xe5, 0x8a,
	0x1e, 0x50, 0x24, 0xb3, 0xb5, 0xb7, 0x96, 0x55, 0xdb, 0x6b, 0x76, 0xd7, 0x11, 0x55, 0xd5, 0x03,
	0x5c, 0x39, 0xf2, 0x06, 0x1c, 0x7b, 0xe7, 0x25, 0x7a, 0xed, 0x2b, 0xc0, 0x85, 0x67, 0xe0, 0x80,
	0xec, 0xf5, 0xa6, 0x7f, 0x90, 0xeb, 0x4a, 0xed, 0x6d, 0x32, 0xfb, 0xf5, 0x67, 0x66, 0x76, 0x66,
	0x27, 0x60, 0x14, 0x10, 0x12, 0x44, 0xd8, 0x44, 0x3e, 0x33, 0x85, 0x99, 0x5b, 0x13, 0xcb, 0x64,
	0x98, 0x4e, 0x42, 0x0f, 0x33, 0xd3, 0x43, 0x71, 0x8a, 0xc2, 0x20, 0x71, 0xf1, 0x27, 0x8e, 0x13,
	0x16, 0x92, 0xc4, 0x65, 0x98, 0xf3, 0x30, 0x09, 0xdc, 0x52, 0x63, 0xa4, 0x94, 0x70, 0x02, 0x5b,
	0xe2, 0x7b, 0x03, 0xf9, 0xcc, 0x98, 0xa2, 0x8c, 0x89, 0x65, 0x48, 0x54, 0xb3, 0x5f, 0x15, 0x8c,
	0x62, 0x46, 0x32, 0x7a, 0x7d, 0x34, 0x11, 0xa5, 0xf9, 0x54, 0x32, 0xd2, 0xd0, 0x44, 0x49, 0x42,
	0x38, 0xe2, 0x21, 0x49, 0x58, 0x79, 0x5a, 0xe6, 0x60, 0x16, 0xbf, 0xf6, 0xb2, 0x7d, 0x73, 0x3f,
	0xc4, 0x91, 0xef, 0xc6, 0x88, 0x1d, 0x94, 0x8a, 0x27, 0xa5, 0x82, 0xa6, 0x9e, 0xc9, 0x38, 0xe2,
	0x19, 0xbb, 0x72, 0x90, 0x83, 0xbd, 0x28, 0xc4, 0x09, 0x17, 0x07, 0xfa, 0x08, 0xe8, 0xeb, 0x98,
	0x0f, 0xca, 0xc4, 0xd6, 0x64, 0x5e, 0x3b, 0x22, 0x2d, 0x07, 0x7f, 0xcc, 0x30, 0xe3, 0xf0, 0x39,
	0xb8, 0x2f, 0xab, 0x70, 0x13, 0x14, 0x63, 0x4d, 0x69, 0x29, 0x2b, 0xf3, 0xce, 0x82, 0x74, 0xbe,
	0x45, 0x31, 0xd6, 0xff, 0x28, 0xe0, 0xc5, 0x66, 0xc6, 0x11, 0xc7, 0x55, 0x38, 0x26, 0x79, 0x4b,
	0xa0, 0xe1, 0x65, 0x8c, 0x93, 0x18, 0x53, 0x37, 0xf4, 0x4b, 0x1a, 0x90, 0xae, 0x91, 0x0f, 0x3d,
	0x00, 0x48, 0x8a, 0xa9, 0x28, 0x5f, 0x53, 0x5b, 0x33, 0x2b, 0x0d, 0x6b, 0x60, 0xd4, 0xf5, 0xc0,
	0xa8, 0x0a, 0xbc, 0x25, 0x59, 0xce, 0x05, 0x2c, 0x5c, 0x06, 0x0f, 0x53, 0x44, 0x79, 0x88, 0x22,
	0x77, 0x1f, 0x85, 0x51, 0x46, 0xb1, 0x36, 0xd3, 0x52, 0x56, 0xee, 0x39, 0x0f, 0x4a, 0xf7, 0x50,
	0x78, 0xf3, 0xf2, 0x27, 0x28, 0x0a, 0x7d, 0xc4, 0xb1, 0x4b, 0x92, 0xe8, 0x50, 0x9b, 0x2d, 0x64,
	0x0b, 0xd2, 0xb9, 0x95, 0x44, 0x87, 0xfa, 0x0f, 0x15, 0x3c, 0xab, 0x8d, 0x0f, 0xbb, 0xa0, 0x91,
	0xa5, 0x05, 0x28, 0x6f, 0x5b, 0x01, 0x6a, 0x58, 0x4d, 0x59, 0x99, 0xec, 0xac, 0x31, 0xcc, 0x3b,
	0xbb, 0x89, 0xd8, 0x81, 0x03, 0x84, 0x3c, 0xb7, 0xe1, 0x3b, 0x30, 0xe7, 0x51, 0x8c, 0xb8, 0xb8,
	0xff, 0x86, 0xd5, 0xad, 0xbc, 0x91, 0xe9, 0xcc, 0x55, 0x5e, 0xc9, 0xc6, 0x7f, 0x4e, 0x09, 0xcb,
	0xb1, 0x22, 0x88, 0xa6, 0xde, 0x09, 0x56, 0xc0, 0xa0, 0x06, 0xe6, 0x28, 0x8e, 0xc9, 0x44, 0xdc,
	0xea, 0x7c, 0x7e, 0x22, 0x7e, 0xf7, 0x1b, 0x60, 0x7e, 0xda, 0x06, 0xfd, 0x4c, 0x01, 0xcb, 0xb5,
	0x63, 0xc3, 0x52, 0x92, 0x30, 0x0c, 0x87, 0xe0, 0xf1, 0x95, 0x8e, 0xb9, 0x98, 0x52, 0x42, 0x8b,
	0x08, 0x0d, 0x0b, 0xca, 0xc4, 0x69, 0xea, 0x19, 0x3b, 0xc5, 0xfc, 0x3b, 0x8f, 0x2e, 0xf7, 0x72,
	0x2d, 0x97, 0xc3, 0x0f, 0xe0, 0x7f, 0x8a, 0x59, 0x16, 0x71, 0x39, 0x5b, 0xc3, 0xfa, 0xd9, 0xba,
	0x3e, 0x47, 0xa7, 0xc0, 0x39, 0x12, 0xab, 0xbf, 0x01, 0xed, 0x9b, 0x7c, 0x70, 0xa3, 0x97, 0x65,
	0x9d, 0xcc, 0x82, 0xa5, 0x2a, 0xce, 0x8e, 0xc8, 0x0f, 0xfe, 0x52, 0xc0, 0xe2, 0x35, 0x2f, 0x19,
	0xbe, 0xae, 0xaf, 0xb0, 0x7e, 0x11, 0x34, 0x6f, 0x33, 0x1a, 0xfa, 0xe0, 0xcb, 0xd9, 0xcf, 0x6f,
	0xea, 0x2a, 0xec, 0xe6, 0x5b, 0xf1, 0xe8, 0x52, 0xd9, 0xab, 0xf2, 0xed, 0x33, 0xb3, 0x33, 0x5d,
	0x93, 0xff, 0xcc, 0x81, 0xd9, 0x39, 0x86, 0x9f, 0x55, 0xb0, 0x54, 0x33, 0x2e, 0x70, 0xe3, 0xb6,
	0xdd, 0x94, 0x8b, 0xaa, 0x39, 0xba, 0x03, 0x92, 0x98, 0x5d, 0x7d, 0x54, 0x54, 0x3f, 0xd0, 0x5f,
	0xe5, 0xd5, 0x9f, 0x97, 0x7b, 0x74, 0x61, 0x11, 0xae, 0x76, 0x8e, 0xab, 0x8b, 0xb7, 0xe3, 0x22,
	0x90, 0xad, 0x74, 0x9a, 0x8b, 0xa7, 0x3d, 0xed, 0x3c, 0x99, 0xd2, 0x4a, 0x43, 0x66, 0x78, 0x24,
	0xee, 0x7f, 0x55, 0x41, 0xdb, 0x23, 0x71, 0x6d, 0xe2, 0xfd, 0x76, 0xcd, 0x48, 0x6d, 0xe7, 0xcb,
	0x68, 0x5b, 0x79, 0xbf, 0x51, 0x92, 0x02, 0x12, 0xa1, 0x24, 0x30, 0x08, 0x0d, 0xcc, 0x00, 0x27,
	0xc5, 0xaa, 0x32, 0xcf, 0x63, 0x57, 0xff, 0xc9, 0x76, 0xa5, 0xf1, 0x5d, 0x9d, 0x59, 0xef, 0xf5,
	0x4e, 0xd4, 0xd6, 0xba, 0x00, 0xf6, 0x7c, 0x66, 0x08, 0x33, 0xb7, 0x76, 0x2d, 0xa3, 0x0c, 0xcc,
	0x4e, 0xa5, 0x64, 0xdc, 0xf3, 0xd9, 0x78, 0x2a, 0x19, 0xef, 0x5a, 0x63, 0x29, 0xf9, 0xad, 0xb6,
	0x85, 0xdf, 0xb6, 0x7b, 0x3e, 0xb3, 0xed, 0xa9, 0xc8, 0xb6, 0x77, 0x2d, 0xdb, 0x96, 0xb2, 0xbd,
	0xb9, 0x22, 0xcf, 0x97, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x35, 0xc6, 0xc6, 0x0b, 0x08,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CampaignExtensionSettingServiceClient is the client API for CampaignExtensionSettingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CampaignExtensionSettingServiceClient interface {
	// Returns the requested campaign extension setting in full detail.
	GetCampaignExtensionSetting(ctx context.Context, in *GetCampaignExtensionSettingRequest, opts ...grpc.CallOption) (*resources.CampaignExtensionSetting, error)
	// Creates, updates, or removes campaign extension settings. Operation
	// statuses are returned.
	MutateCampaignExtensionSettings(ctx context.Context, in *MutateCampaignExtensionSettingsRequest, opts ...grpc.CallOption) (*MutateCampaignExtensionSettingsResponse, error)
}

type campaignExtensionSettingServiceClient struct {
	cc *grpc.ClientConn
}

func NewCampaignExtensionSettingServiceClient(cc *grpc.ClientConn) CampaignExtensionSettingServiceClient {
	return &campaignExtensionSettingServiceClient{cc}
}

func (c *campaignExtensionSettingServiceClient) GetCampaignExtensionSetting(ctx context.Context, in *GetCampaignExtensionSettingRequest, opts ...grpc.CallOption) (*resources.CampaignExtensionSetting, error) {
	out := new(resources.CampaignExtensionSetting)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.CampaignExtensionSettingService/GetCampaignExtensionSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignExtensionSettingServiceClient) MutateCampaignExtensionSettings(ctx context.Context, in *MutateCampaignExtensionSettingsRequest, opts ...grpc.CallOption) (*MutateCampaignExtensionSettingsResponse, error) {
	out := new(MutateCampaignExtensionSettingsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.CampaignExtensionSettingService/MutateCampaignExtensionSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignExtensionSettingServiceServer is the server API for CampaignExtensionSettingService service.
type CampaignExtensionSettingServiceServer interface {
	// Returns the requested campaign extension setting in full detail.
	GetCampaignExtensionSetting(context.Context, *GetCampaignExtensionSettingRequest) (*resources.CampaignExtensionSetting, error)
	// Creates, updates, or removes campaign extension settings. Operation
	// statuses are returned.
	MutateCampaignExtensionSettings(context.Context, *MutateCampaignExtensionSettingsRequest) (*MutateCampaignExtensionSettingsResponse, error)
}

func RegisterCampaignExtensionSettingServiceServer(s *grpc.Server, srv CampaignExtensionSettingServiceServer) {
	s.RegisterService(&_CampaignExtensionSettingService_serviceDesc, srv)
}

func _CampaignExtensionSettingService_GetCampaignExtensionSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCampaignExtensionSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignExtensionSettingServiceServer).GetCampaignExtensionSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.CampaignExtensionSettingService/GetCampaignExtensionSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignExtensionSettingServiceServer).GetCampaignExtensionSetting(ctx, req.(*GetCampaignExtensionSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignExtensionSettingService_MutateCampaignExtensionSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignExtensionSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignExtensionSettingServiceServer).MutateCampaignExtensionSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.CampaignExtensionSettingService/MutateCampaignExtensionSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignExtensionSettingServiceServer).MutateCampaignExtensionSettings(ctx, req.(*MutateCampaignExtensionSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CampaignExtensionSettingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.CampaignExtensionSettingService",
	HandlerType: (*CampaignExtensionSettingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCampaignExtensionSetting",
			Handler:    _CampaignExtensionSettingService_GetCampaignExtensionSetting_Handler,
		},
		{
			MethodName: "MutateCampaignExtensionSettings",
			Handler:    _CampaignExtensionSettingService_MutateCampaignExtensionSettings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/campaign_extension_setting_service.proto",
}
