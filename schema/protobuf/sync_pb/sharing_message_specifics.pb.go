// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sharing_message_specifics.proto

package sync_pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

// This enum is used in histograms. Entries should not be renumbered and
// numeric values should never be reused. Also remember to update in
// tools/metrics/histograms/enums.xml SyncSharingMessageCommitErrorCode enum.
type SharingMessageCommitError_ErrorCode int32

const (
	SharingMessageCommitError_NONE               SharingMessageCommitError_ErrorCode = 0
	SharingMessageCommitError_INVALID_ARGUMENT   SharingMessageCommitError_ErrorCode = 1
	SharingMessageCommitError_NOT_FOUND          SharingMessageCommitError_ErrorCode = 2
	SharingMessageCommitError_INTERNAL           SharingMessageCommitError_ErrorCode = 3
	SharingMessageCommitError_UNAVAILABLE        SharingMessageCommitError_ErrorCode = 4
	SharingMessageCommitError_RESOURCE_EXHAUSTED SharingMessageCommitError_ErrorCode = 5
	SharingMessageCommitError_UNAUTHENTICATED    SharingMessageCommitError_ErrorCode = 6
	SharingMessageCommitError_PERMISSION_DENIED  SharingMessageCommitError_ErrorCode = 7
	// Client-specific error codes.
	SharingMessageCommitError_SYNC_TURNED_OFF    SharingMessageCommitError_ErrorCode = 8
	SharingMessageCommitError_SYNC_NETWORK_ERROR SharingMessageCommitError_ErrorCode = 9
	// Error code for server error or unparsable server response.
	SharingMessageCommitError_SYNC_SERVER_ERROR SharingMessageCommitError_ErrorCode = 10
	// Message wasn't committed before timeout.
	SharingMessageCommitError_SYNC_TIMEOUT SharingMessageCommitError_ErrorCode = 11
)

var SharingMessageCommitError_ErrorCode_name = map[int32]string{
	0:  "NONE",
	1:  "INVALID_ARGUMENT",
	2:  "NOT_FOUND",
	3:  "INTERNAL",
	4:  "UNAVAILABLE",
	5:  "RESOURCE_EXHAUSTED",
	6:  "UNAUTHENTICATED",
	7:  "PERMISSION_DENIED",
	8:  "SYNC_TURNED_OFF",
	9:  "SYNC_NETWORK_ERROR",
	10: "SYNC_SERVER_ERROR",
	11: "SYNC_TIMEOUT",
}

var SharingMessageCommitError_ErrorCode_value = map[string]int32{
	"NONE":               0,
	"INVALID_ARGUMENT":   1,
	"NOT_FOUND":          2,
	"INTERNAL":           3,
	"UNAVAILABLE":        4,
	"RESOURCE_EXHAUSTED": 5,
	"UNAUTHENTICATED":    6,
	"PERMISSION_DENIED":  7,
	"SYNC_TURNED_OFF":    8,
	"SYNC_NETWORK_ERROR": 9,
	"SYNC_SERVER_ERROR":  10,
	"SYNC_TIMEOUT":       11,
}

func (x SharingMessageCommitError_ErrorCode) Enum() *SharingMessageCommitError_ErrorCode {
	p := new(SharingMessageCommitError_ErrorCode)
	*p = x
	return p
}

func (x SharingMessageCommitError_ErrorCode) String() string {
	return proto.EnumName(SharingMessageCommitError_ErrorCode_name, int32(x))
}

func (x *SharingMessageCommitError_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SharingMessageCommitError_ErrorCode_value, data, "SharingMessageCommitError_ErrorCode")
	if err != nil {
		return err
	}
	*x = SharingMessageCommitError_ErrorCode(value)
	return nil
}

func (SharingMessageCommitError_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_490056852b5b1e8d, []int{1, 0}
}

type SharingMessageSpecifics struct {
	// Unique identifier of message.
	MessageId            *string                                       `protobuf:"bytes,1,opt,name=message_id,json=messageId" json:"message_id,omitempty"`
	ChannelConfiguration *SharingMessageSpecifics_ChannelConfiguration `protobuf:"bytes,2,opt,name=channel_configuration,json=channelConfiguration" json:"channel_configuration,omitempty"`
	// Payload encrypted using the target user keys according to WebPush
	// encryption scheme. The payload has to be a valid
	// chrome/browser/sharing/proto/sharing_message.proto serialized using
	// SerializeToString.
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SharingMessageSpecifics) Reset()         { *m = SharingMessageSpecifics{} }
func (m *SharingMessageSpecifics) String() string { return proto.CompactTextString(m) }
func (*SharingMessageSpecifics) ProtoMessage()    {}
func (*SharingMessageSpecifics) Descriptor() ([]byte, []int) {
	return fileDescriptor_490056852b5b1e8d, []int{0}
}

func (m *SharingMessageSpecifics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharingMessageSpecifics.Unmarshal(m, b)
}
func (m *SharingMessageSpecifics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharingMessageSpecifics.Marshal(b, m, deterministic)
}
func (m *SharingMessageSpecifics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharingMessageSpecifics.Merge(m, src)
}
func (m *SharingMessageSpecifics) XXX_Size() int {
	return xxx_messageInfo_SharingMessageSpecifics.Size(m)
}
func (m *SharingMessageSpecifics) XXX_DiscardUnknown() {
	xxx_messageInfo_SharingMessageSpecifics.DiscardUnknown(m)
}

var xxx_messageInfo_SharingMessageSpecifics proto.InternalMessageInfo

func (m *SharingMessageSpecifics) GetMessageId() string {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return ""
}

func (m *SharingMessageSpecifics) GetChannelConfiguration() *SharingMessageSpecifics_ChannelConfiguration {
	if m != nil {
		return m.ChannelConfiguration
	}
	return nil
}

func (m *SharingMessageSpecifics) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type SharingMessageSpecifics_ChannelConfiguration struct {
	// Types that are valid to be assigned to ChannelConfiguration:
	//	*SharingMessageSpecifics_ChannelConfiguration_Fcm
	//	*SharingMessageSpecifics_ChannelConfiguration_Server
	ChannelConfiguration isSharingMessageSpecifics_ChannelConfiguration_ChannelConfiguration `protobuf_oneof:"channel_configuration"`
	XXX_NoUnkeyedLiteral struct{}                                                            `json:"-"`
	XXX_unrecognized     []byte                                                              `json:"-"`
	XXX_sizecache        int32                                                               `json:"-"`
}

func (m *SharingMessageSpecifics_ChannelConfiguration) Reset() {
	*m = SharingMessageSpecifics_ChannelConfiguration{}
}
func (m *SharingMessageSpecifics_ChannelConfiguration) String() string {
	return proto.CompactTextString(m)
}
func (*SharingMessageSpecifics_ChannelConfiguration) ProtoMessage() {}
func (*SharingMessageSpecifics_ChannelConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_490056852b5b1e8d, []int{0, 0}
}

func (m *SharingMessageSpecifics_ChannelConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharingMessageSpecifics_ChannelConfiguration.Unmarshal(m, b)
}
func (m *SharingMessageSpecifics_ChannelConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharingMessageSpecifics_ChannelConfiguration.Marshal(b, m, deterministic)
}
func (m *SharingMessageSpecifics_ChannelConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharingMessageSpecifics_ChannelConfiguration.Merge(m, src)
}
func (m *SharingMessageSpecifics_ChannelConfiguration) XXX_Size() int {
	return xxx_messageInfo_SharingMessageSpecifics_ChannelConfiguration.Size(m)
}
func (m *SharingMessageSpecifics_ChannelConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_SharingMessageSpecifics_ChannelConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_SharingMessageSpecifics_ChannelConfiguration proto.InternalMessageInfo

type isSharingMessageSpecifics_ChannelConfiguration_ChannelConfiguration interface {
	isSharingMessageSpecifics_ChannelConfiguration_ChannelConfiguration()
}

type SharingMessageSpecifics_ChannelConfiguration_Fcm struct {
	Fcm *SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration `protobuf:"bytes,1,opt,name=fcm,oneof"`
}

type SharingMessageSpecifics_ChannelConfiguration_Server struct {
	Server []byte `protobuf:"bytes,2,opt,name=server,oneof"`
}

func (*SharingMessageSpecifics_ChannelConfiguration_Fcm) isSharingMessageSpecifics_ChannelConfiguration_ChannelConfiguration() {
}

func (*SharingMessageSpecifics_ChannelConfiguration_Server) isSharingMessageSpecifics_ChannelConfiguration_ChannelConfiguration() {
}

func (m *SharingMessageSpecifics_ChannelConfiguration) GetChannelConfiguration() isSharingMessageSpecifics_ChannelConfiguration_ChannelConfiguration {
	if m != nil {
		return m.ChannelConfiguration
	}
	return nil
}

func (m *SharingMessageSpecifics_ChannelConfiguration) GetFcm() *SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration {
	if x, ok := m.GetChannelConfiguration().(*SharingMessageSpecifics_ChannelConfiguration_Fcm); ok {
		return x.Fcm
	}
	return nil
}

func (m *SharingMessageSpecifics_ChannelConfiguration) GetServer() []byte {
	if x, ok := m.GetChannelConfiguration().(*SharingMessageSpecifics_ChannelConfiguration_Server); ok {
		return x.Server
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SharingMessageSpecifics_ChannelConfiguration) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SharingMessageSpecifics_ChannelConfiguration_Fcm)(nil),
		(*SharingMessageSpecifics_ChannelConfiguration_Server)(nil),
	}
}

type SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration struct {
	// FCM registration token of target device.
	Token *string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	// Time to live for a FCM message (in seconds) - if specified, the message
	// will expire based on the TTL.
	Ttl *int32 `protobuf:"varint,2,opt,name=ttl" json:"ttl,omitempty"`
	// Priority level of a FCM message. 5 = normal, 10 = high.
	Priority             *int32   `protobuf:"varint,3,opt,name=priority" json:"priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration) Reset() {
	*m = SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration{}
}
func (m *SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration) String() string {
	return proto.CompactTextString(m)
}
func (*SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration) ProtoMessage() {}
func (*SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_490056852b5b1e8d, []int{0, 0, 0}
}

func (m *SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration.Unmarshal(m, b)
}
func (m *SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration.Marshal(b, m, deterministic)
}
func (m *SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration.Merge(m, src)
}
func (m *SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration) XXX_Size() int {
	return xxx_messageInfo_SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration.Size(m)
}
func (m *SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration proto.InternalMessageInfo

func (m *SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

func (m *SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration) GetTtl() int32 {
	if m != nil && m.Ttl != nil {
		return *m.Ttl
	}
	return 0
}

func (m *SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration) GetPriority() int32 {
	if m != nil && m.Priority != nil {
		return *m.Priority
	}
	return 0
}

// Used for the server to return fine grained commit errors back to the client.
type SharingMessageCommitError struct {
	ErrorCode            *SharingMessageCommitError_ErrorCode `protobuf:"varint,1,opt,name=error_code,json=errorCode,enum=sync_pb.SharingMessageCommitError_ErrorCode" json:"error_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *SharingMessageCommitError) Reset()         { *m = SharingMessageCommitError{} }
func (m *SharingMessageCommitError) String() string { return proto.CompactTextString(m) }
func (*SharingMessageCommitError) ProtoMessage()    {}
func (*SharingMessageCommitError) Descriptor() ([]byte, []int) {
	return fileDescriptor_490056852b5b1e8d, []int{1}
}

func (m *SharingMessageCommitError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharingMessageCommitError.Unmarshal(m, b)
}
func (m *SharingMessageCommitError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharingMessageCommitError.Marshal(b, m, deterministic)
}
func (m *SharingMessageCommitError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharingMessageCommitError.Merge(m, src)
}
func (m *SharingMessageCommitError) XXX_Size() int {
	return xxx_messageInfo_SharingMessageCommitError.Size(m)
}
func (m *SharingMessageCommitError) XXX_DiscardUnknown() {
	xxx_messageInfo_SharingMessageCommitError.DiscardUnknown(m)
}

var xxx_messageInfo_SharingMessageCommitError proto.InternalMessageInfo

func (m *SharingMessageCommitError) GetErrorCode() SharingMessageCommitError_ErrorCode {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return SharingMessageCommitError_NONE
}

func init() {
	proto.RegisterEnum("sync_pb.SharingMessageCommitError_ErrorCode", SharingMessageCommitError_ErrorCode_name, SharingMessageCommitError_ErrorCode_value)
	proto.RegisterType((*SharingMessageSpecifics)(nil), "sync_pb.SharingMessageSpecifics")
	proto.RegisterType((*SharingMessageSpecifics_ChannelConfiguration)(nil), "sync_pb.SharingMessageSpecifics.ChannelConfiguration")
	proto.RegisterType((*SharingMessageSpecifics_ChannelConfiguration_FCMChannelConfiguration)(nil), "sync_pb.SharingMessageSpecifics.ChannelConfiguration.FCMChannelConfiguration")
	proto.RegisterType((*SharingMessageCommitError)(nil), "sync_pb.SharingMessageCommitError")
}

func init() {
	proto.RegisterFile("sharing_message_specifics.proto", fileDescriptor_490056852b5b1e8d)
}

var fileDescriptor_490056852b5b1e8d = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x5d, 0x9a, 0x75, 0x6b, 0x6e, 0x0b, 0x33, 0xa6, 0x63, 0x65, 0x12, 0x62, 0x9a, 0x84, 0x34,
	0x09, 0x94, 0x87, 0x49, 0x7c, 0x40, 0x96, 0xb8, 0x34, 0x5a, 0xe3, 0x4c, 0x4e, 0x32, 0xe0, 0x01,
	0x59, 0x21, 0xf5, 0xda, 0x40, 0x13, 0x47, 0x4e, 0x86, 0xd4, 0x6f, 0xe0, 0xd7, 0xf8, 0x06, 0x7e,
	0x82, 0x1f, 0x40, 0x49, 0xda, 0x69, 0x48, 0xed, 0x0b, 0x2f, 0xd1, 0x3d, 0x27, 0xe7, 0x1c, 0xfb,
	0x1e, 0xc9, 0xf0, 0xba, 0x5c, 0xc4, 0x2a, 0xcd, 0xe7, 0x3c, 0x13, 0x65, 0x19, 0xcf, 0x05, 0x2f,
	0x0b, 0x91, 0xa4, 0x77, 0x69, 0x52, 0x9a, 0x85, 0x92, 0x95, 0xc4, 0x87, 0xe5, 0x2a, 0x4f, 0x78,
	0xf1, 0xf5, 0xfc, 0x97, 0x0e, 0x27, 0x41, 0x2b, 0xf6, 0x5a, 0x6d, 0xb0, 0x91, 0xe2, 0x57, 0x00,
	0x1b, 0x7f, 0x3a, 0x1b, 0x69, 0x67, 0xda, 0x85, 0xc1, 0x8c, 0x35, 0xe3, 0xce, 0xf0, 0x37, 0x38,
	0x4e, 0x16, 0x71, 0x9e, 0x8b, 0x25, 0x4f, 0x64, 0x7e, 0x97, 0xce, 0xef, 0x55, 0x5c, 0xa5, 0x32,
	0x1f, 0x75, 0xce, 0xb4, 0x8b, 0xfe, 0xe5, 0x7b, 0x73, 0x7d, 0x86, 0xb9, 0x23, 0xdf, 0xb4, 0x5b,
	0xb7, 0xfd, 0xd8, 0xcc, 0x86, 0xc9, 0x16, 0x16, 0x8f, 0xe0, 0xb0, 0x88, 0x57, 0x4b, 0x19, 0xcf,
	0x46, 0xfa, 0x99, 0x76, 0x31, 0x60, 0x1b, 0x78, 0xfa, 0xb3, 0x03, 0xc3, 0x6d, 0x41, 0x38, 0x06,
	0xfd, 0x2e, 0xc9, 0x9a, 0x6b, 0xf7, 0x2f, 0xbd, 0xff, 0xba, 0x8c, 0x39, 0xb6, 0xbd, 0x6d, 0xfc,
	0x64, 0x8f, 0xd5, 0xd9, 0x78, 0x04, 0x07, 0xa5, 0x50, 0x3f, 0x84, 0x6a, 0x56, 0x1e, 0x4c, 0xf6,
	0xd8, 0x1a, 0x9f, 0x7e, 0x81, 0x93, 0x1d, 0x5e, 0x3c, 0x84, 0x6e, 0x25, 0xbf, 0x8b, 0x7c, 0x5d,
	0x68, 0x0b, 0x30, 0x02, 0xbd, 0xaa, 0x96, 0x4d, 0x4e, 0x97, 0xd5, 0x23, 0x3e, 0x85, 0x5e, 0xa1,
	0x52, 0xa9, 0xd2, 0x6a, 0xd5, 0xec, 0xdc, 0x65, 0x0f, 0xf8, 0xea, 0x64, 0x47, 0xf5, 0xe7, 0xbf,
	0x3b, 0xf0, 0xf2, 0xdf, 0x0d, 0x6d, 0x99, 0x65, 0x69, 0x45, 0x94, 0x92, 0x0a, 0x5f, 0x03, 0x88,
	0x7a, 0xe0, 0x89, 0x9c, 0x89, 0xe6, 0xfc, 0xa7, 0x97, 0xef, 0x76, 0x34, 0xf3, 0xc8, 0x67, 0x36,
	0x5f, 0x5b, 0xce, 0x04, 0x33, 0xc4, 0x66, 0x3c, 0xff, 0xa3, 0x81, 0xf1, 0xf0, 0x03, 0xf7, 0x60,
	0x9f, 0xfa, 0x94, 0xa0, 0x3d, 0x3c, 0x04, 0xe4, 0xd2, 0x5b, 0x6b, 0xea, 0x3a, 0xdc, 0x62, 0x1f,
	0x22, 0x8f, 0xd0, 0x10, 0x69, 0xf8, 0x09, 0x18, 0xd4, 0x0f, 0xf9, 0xd8, 0x8f, 0xa8, 0x83, 0x3a,
	0x78, 0x00, 0x3d, 0x97, 0x86, 0x84, 0x51, 0x6b, 0x8a, 0x74, 0x7c, 0x04, 0xfd, 0x88, 0x5a, 0xb7,
	0x96, 0x3b, 0xb5, 0xae, 0xa6, 0x04, 0xed, 0xe3, 0x17, 0x80, 0x19, 0x09, 0xfc, 0x88, 0xd9, 0x84,
	0x93, 0x4f, 0x13, 0x2b, 0x0a, 0x42, 0xe2, 0xa0, 0x2e, 0x7e, 0x0e, 0x47, 0x11, 0xb5, 0xa2, 0x70,
	0x42, 0x68, 0xe8, 0xda, 0x56, 0x4d, 0x1e, 0xe0, 0x63, 0x78, 0x76, 0x43, 0x98, 0xe7, 0x06, 0x81,
	0xeb, 0x53, 0xee, 0x10, 0xea, 0x12, 0x07, 0x1d, 0xd6, 0xda, 0xe0, 0x33, 0xb5, 0x79, 0x18, 0x31,
	0x4a, 0x1c, 0xee, 0x8f, 0xc7, 0xa8, 0x57, 0x07, 0x37, 0x24, 0x25, 0xe1, 0x47, 0x9f, 0x5d, 0x73,
	0xc2, 0x98, 0xcf, 0x90, 0x51, 0x67, 0x34, 0x7c, 0x40, 0xd8, 0x2d, 0x61, 0x6b, 0x1a, 0x30, 0x82,
	0x41, 0x9b, 0xe1, 0x7a, 0xc4, 0x8f, 0x42, 0xd4, 0xbf, 0x7a, 0x0b, 0x6f, 0xa4, 0x9a, 0x9b, 0xc9,
	0x42, 0xc9, 0x2c, 0xbd, 0xcf, 0xcc, 0x44, 0x66, 0x85, 0xcc, 0x45, 0x5e, 0x95, 0x4d, 0x8f, 0xed,
	0xf3, 0x4a, 0xe4, 0x72, 0xa2, 0xdf, 0x68, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x36, 0xad, 0x5c,
	0x40, 0x87, 0x03, 0x00, 0x00,
}
