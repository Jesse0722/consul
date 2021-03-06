// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/listener/quic_config.proto

package envoy_api_v2_listener

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type QuicProtocolOptions struct {
	MaxConcurrentStreams   *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=max_concurrent_streams,json=maxConcurrentStreams,proto3" json:"max_concurrent_streams,omitempty"`
	IdleTimeout            *duration.Duration    `protobuf:"bytes,2,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	CryptoHandshakeTimeout *duration.Duration    `protobuf:"bytes,3,opt,name=crypto_handshake_timeout,json=cryptoHandshakeTimeout,proto3" json:"crypto_handshake_timeout,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}              `json:"-"`
	XXX_unrecognized       []byte                `json:"-"`
	XXX_sizecache          int32                 `json:"-"`
}

func (m *QuicProtocolOptions) Reset()         { *m = QuicProtocolOptions{} }
func (m *QuicProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*QuicProtocolOptions) ProtoMessage()    {}
func (*QuicProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f6a4a402e708e40, []int{0}
}

func (m *QuicProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuicProtocolOptions.Unmarshal(m, b)
}
func (m *QuicProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuicProtocolOptions.Marshal(b, m, deterministic)
}
func (m *QuicProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuicProtocolOptions.Merge(m, src)
}
func (m *QuicProtocolOptions) XXX_Size() int {
	return xxx_messageInfo_QuicProtocolOptions.Size(m)
}
func (m *QuicProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_QuicProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_QuicProtocolOptions proto.InternalMessageInfo

func (m *QuicProtocolOptions) GetMaxConcurrentStreams() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConcurrentStreams
	}
	return nil
}

func (m *QuicProtocolOptions) GetIdleTimeout() *duration.Duration {
	if m != nil {
		return m.IdleTimeout
	}
	return nil
}

func (m *QuicProtocolOptions) GetCryptoHandshakeTimeout() *duration.Duration {
	if m != nil {
		return m.CryptoHandshakeTimeout
	}
	return nil
}

func init() {
	proto.RegisterType((*QuicProtocolOptions)(nil), "envoy.api.v2.listener.QuicProtocolOptions")
}

func init() {
	proto.RegisterFile("envoy/api/v2/listener/quic_config.proto", fileDescriptor_1f6a4a402e708e40)
}

var fileDescriptor_1f6a4a402e708e40 = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0x86, 0x95, 0x5e, 0xdd, 0x0e, 0xe9, 0x95, 0xee, 0x55, 0x2e, 0x94, 0x50, 0x01, 0x42, 0x30,
	0xc0, 0x64, 0x4b, 0xe9, 0xca, 0x42, 0x0b, 0x12, 0x48, 0x08, 0x4a, 0x0b, 0x5d, 0xa3, 0xd3, 0xc4,
	0x4d, 0x2d, 0x12, 0xdb, 0x38, 0x76, 0x68, 0x37, 0xde, 0x80, 0x85, 0x81, 0x67, 0xe0, 0x11, 0x78,
	0x02, 0x56, 0x5e, 0x83, 0x91, 0x19, 0x21, 0x14, 0x27, 0x81, 0xa1, 0x20, 0x46, 0xeb, 0xff, 0xbe,
	0x5f, 0x3e, 0xe7, 0xd8, 0x5b, 0x84, 0x65, 0x7c, 0x86, 0x41, 0x50, 0x9c, 0x79, 0x38, 0xa6, 0xa9,
	0x22, 0x8c, 0x48, 0x7c, 0xa9, 0x69, 0xe0, 0x07, 0x9c, 0x8d, 0x69, 0x84, 0x84, 0xe4, 0x8a, 0x3b,
	0x8b, 0x06, 0x44, 0x20, 0x28, 0xca, 0x3c, 0x54, 0x81, 0xad, 0xb5, 0x88, 0xf3, 0x28, 0x26, 0xd8,
	0x40, 0x23, 0x3d, 0xc6, 0xa1, 0x96, 0xa0, 0x28, 0x67, 0x85, 0x36, 0x9f, 0x5f, 0x49, 0x10, 0x82,
	0xc8, 0xb4, 0xca, 0x75, 0x28, 0x00, 0x03, 0x63, 0x5c, 0x19, 0x2d, 0xc5, 0x09, 0x8d, 0x24, 0x28,
	0x52, 0xe6, 0xab, 0x73, 0x79, 0xaa, 0x40, 0xe9, 0x52, 0xdf, 0x78, 0xb5, 0xec, 0xff, 0xa7, 0x9a,
	0x06, 0xbd, 0xfc, 0x15, 0xf0, 0xf8, 0x44, 0x18, 0xc8, 0xe9, 0xdb, 0xcd, 0x04, 0xa6, 0xf9, 0x04,
	0x81, 0x96, 0x92, 0x30, 0xe5, 0xa7, 0x4a, 0x12, 0x48, 0x52, 0xd7, 0x5a, 0xb7, 0xb6, 0x1b, 0xde,
	0x0a, 0x2a, 0xfe, 0x85, 0xaa, 0x7f, 0xa1, 0xf3, 0x43, 0xa6, 0xda, 0xde, 0x10, 0x62, 0x4d, 0xfa,
	0x0b, 0x09, 0x4c, 0xbb, 0x1f, 0xea, 0xa0, 0x30, 0x9d, 0x1d, 0xfb, 0x0f, 0x0d, 0x63, 0xe2, 0x2b,
	0x9a, 0x10, 0xae, 0x95, 0x5b, 0x33, 0x4d, 0xcb, 0x73, 0x4d, 0x7b, 0xe5, 0x06, 0xfa, 0x8d, 0x1c,
	0x3f, 0x2b, 0x68, 0x67, 0x60, 0xbb, 0x81, 0x9c, 0x09, 0xc5, 0xfd, 0x09, 0xb0, 0x30, 0x9d, 0xc0,
	0xc5, 0x67, 0xd3, 0xaf, 0x9f, 0x9a, 0x9a, 0x85, 0x7a, 0x50, 0x99, 0x65, 0x69, 0xe7, 0xd6, 0x7a,
	0xb9, 0x7b, 0xbb, 0xf9, 0xdd, 0x72, 0xdc, 0xe2, 0x3a, 0xe5, 0xc5, 0xaa, 0xeb, 0xa0, 0xac, 0xfd,
	0x70, 0xfd, 0xf8, 0x54, 0xaf, 0xfd, 0xb3, 0xec, 0x4d, 0xca, 0x91, 0x81, 0x84, 0xe4, 0xd3, 0x19,
	0xfa, 0xf2, 0x9a, 0x9d, 0xbf, 0xf9, 0x2e, 0xbb, 0xa6, 0xc4, 0x6c, 0xb4, 0x67, 0xdd, 0xd7, 0x96,
	0xf6, 0x0d, 0xba, 0x2b, 0x28, 0x1a, 0x7a, 0xe8, 0xa8, 0x44, 0x8f, 0x07, 0xcf, 0xdf, 0x26, 0xa3,
	0xba, 0x99, 0xa0, 0xfd, 0x1e, 0x00, 0x00, 0xff, 0xff, 0x12, 0xf9, 0x6d, 0x9e, 0x5d, 0x02, 0x00,
	0x00,
}
