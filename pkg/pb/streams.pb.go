// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: streams.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SegmentMeta struct {
	EncryptedKey         []byte   `protobuf:"bytes,1,opt,name=encrypted_key,json=encryptedKey,proto3" json:"encrypted_key,omitempty"`
	KeyNonce             []byte   `protobuf:"bytes,2,opt,name=key_nonce,json=keyNonce,proto3" json:"key_nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SegmentMeta) Reset()         { *m = SegmentMeta{} }
func (m *SegmentMeta) String() string { return proto.CompactTextString(m) }
func (*SegmentMeta) ProtoMessage()    {}
func (*SegmentMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_streams_c0d9754174b032dc, []int{0}
}
func (m *SegmentMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegmentMeta.Unmarshal(m, b)
}
func (m *SegmentMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegmentMeta.Marshal(b, m, deterministic)
}
func (dst *SegmentMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegmentMeta.Merge(dst, src)
}
func (m *SegmentMeta) XXX_Size() int {
	return xxx_messageInfo_SegmentMeta.Size(m)
}
func (m *SegmentMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_SegmentMeta.DiscardUnknown(m)
}

var xxx_messageInfo_SegmentMeta proto.InternalMessageInfo

func (m *SegmentMeta) GetEncryptedKey() []byte {
	if m != nil {
		return m.EncryptedKey
	}
	return nil
}

func (m *SegmentMeta) GetKeyNonce() []byte {
	if m != nil {
		return m.KeyNonce
	}
	return nil
}

type StreamInfo struct {
	NumberOfSegments     int64    `protobuf:"varint,1,opt,name=number_of_segments,json=numberOfSegments,proto3" json:"number_of_segments,omitempty"`
	SegmentsSize         int64    `protobuf:"varint,2,opt,name=segments_size,json=segmentsSize,proto3" json:"segments_size,omitempty"`
	LastSegmentSize      int64    `protobuf:"varint,3,opt,name=last_segment_size,json=lastSegmentSize,proto3" json:"last_segment_size,omitempty"`
	Metadata             []byte   `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamInfo) Reset()         { *m = StreamInfo{} }
func (m *StreamInfo) String() string { return proto.CompactTextString(m) }
func (*StreamInfo) ProtoMessage()    {}
func (*StreamInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_streams_c0d9754174b032dc, []int{1}
}
func (m *StreamInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamInfo.Unmarshal(m, b)
}
func (m *StreamInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamInfo.Marshal(b, m, deterministic)
}
func (dst *StreamInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamInfo.Merge(dst, src)
}
func (m *StreamInfo) XXX_Size() int {
	return xxx_messageInfo_StreamInfo.Size(m)
}
func (m *StreamInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamInfo.DiscardUnknown(m)
}

var xxx_messageInfo_StreamInfo proto.InternalMessageInfo

func (m *StreamInfo) GetNumberOfSegments() int64 {
	if m != nil {
		return m.NumberOfSegments
	}
	return 0
}

func (m *StreamInfo) GetSegmentsSize() int64 {
	if m != nil {
		return m.SegmentsSize
	}
	return 0
}

func (m *StreamInfo) GetLastSegmentSize() int64 {
	if m != nil {
		return m.LastSegmentSize
	}
	return 0
}

func (m *StreamInfo) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type StreamMeta struct {
	EncryptedStreamInfo  []byte       `protobuf:"bytes,1,opt,name=encrypted_stream_info,json=encryptedStreamInfo,proto3" json:"encrypted_stream_info,omitempty"`
	EncryptionType       int32        `protobuf:"varint,2,opt,name=encryption_type,json=encryptionType,proto3" json:"encryption_type,omitempty"`
	EncryptionBlockSize  int32        `protobuf:"varint,3,opt,name=encryption_block_size,json=encryptionBlockSize,proto3" json:"encryption_block_size,omitempty"`
	LastSegmentMeta      *SegmentMeta `protobuf:"bytes,4,opt,name=last_segment_meta,json=lastSegmentMeta" json:"last_segment_meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StreamMeta) Reset()         { *m = StreamMeta{} }
func (m *StreamMeta) String() string { return proto.CompactTextString(m) }
func (*StreamMeta) ProtoMessage()    {}
func (*StreamMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_streams_c0d9754174b032dc, []int{2}
}
func (m *StreamMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamMeta.Unmarshal(m, b)
}
func (m *StreamMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamMeta.Marshal(b, m, deterministic)
}
func (dst *StreamMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMeta.Merge(dst, src)
}
func (m *StreamMeta) XXX_Size() int {
	return xxx_messageInfo_StreamMeta.Size(m)
}
func (m *StreamMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMeta.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMeta proto.InternalMessageInfo

func (m *StreamMeta) GetEncryptedStreamInfo() []byte {
	if m != nil {
		return m.EncryptedStreamInfo
	}
	return nil
}

func (m *StreamMeta) GetEncryptionType() int32 {
	if m != nil {
		return m.EncryptionType
	}
	return 0
}

func (m *StreamMeta) GetEncryptionBlockSize() int32 {
	if m != nil {
		return m.EncryptionBlockSize
	}
	return 0
}

func (m *StreamMeta) GetLastSegmentMeta() *SegmentMeta {
	if m != nil {
		return m.LastSegmentMeta
	}
	return nil
}

func init() {
	proto.RegisterType((*SegmentMeta)(nil), "streams.SegmentMeta")
	proto.RegisterType((*StreamInfo)(nil), "streams.StreamInfo")
	proto.RegisterType((*StreamMeta)(nil), "streams.StreamMeta")
}

func init() { proto.RegisterFile("streams.proto", fileDescriptor_streams_c0d9754174b032dc) }

var fileDescriptor_streams_c0d9754174b032dc = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x51, 0xcb, 0x4e, 0xc3, 0x30,
	0x10, 0x54, 0x5f, 0x50, 0xb6, 0x29, 0x05, 0x03, 0x52, 0x04, 0x17, 0x14, 0x0e, 0x20, 0x84, 0x7a,
	0x28, 0x3f, 0x80, 0x7a, 0x43, 0x08, 0x2a, 0x25, 0x9c, 0xb8, 0x58, 0x4e, 0xba, 0x41, 0x51, 0x1a,
	0x3b, 0x8a, 0xcd, 0xc1, 0xfd, 0x21, 0x3e, 0x8b, 0x5f, 0x41, 0xb6, 0xf3, 0x82, 0xe3, 0xce, 0x8c,
	0x66, 0x67, 0x76, 0x61, 0x2e, 0x55, 0x85, 0xac, 0x90, 0xcb, 0xb2, 0x12, 0x4a, 0x90, 0xc3, 0x7a,
	0x0c, 0x36, 0x30, 0x8b, 0xf0, 0xb3, 0x40, 0xae, 0x5e, 0x51, 0x31, 0x72, 0x03, 0x73, 0xe4, 0x49,
	0xa5, 0x4b, 0x85, 0x5b, 0x9a, 0xa3, 0xf6, 0x07, 0xd7, 0x83, 0x3b, 0x2f, 0xf4, 0x5a, 0xf0, 0x05,
	0x35, 0xb9, 0x82, 0xa3, 0x1c, 0x35, 0xe5, 0x82, 0x27, 0xe8, 0x0f, 0xad, 0x60, 0x9a, 0xa3, 0x7e,
	0x33, 0x73, 0xf0, 0x3d, 0x00, 0x88, 0xac, 0xf9, 0x33, 0x4f, 0x05, 0x79, 0x00, 0xc2, 0xbf, 0x8a,
	0x18, 0x2b, 0x2a, 0x52, 0x2a, 0xdd, 0x26, 0x69, 0x5d, 0x47, 0xe1, 0x89, 0x63, 0x36, 0x69, 0x9d,
	0x40, 0x9a, 0xf5, 0x8d, 0x86, 0xca, 0x6c, 0xef, 0xdc, 0x47, 0xa1, 0xd7, 0x80, 0x51, 0xb6, 0x47,
	0x72, 0x0f, 0xa7, 0x3b, 0x26, 0x55, 0xe3, 0xe6, 0x84, 0x23, 0x2b, 0x5c, 0x18, 0xa2, 0x76, 0xb3,
	0xda, 0x4b, 0x98, 0x16, 0xa8, 0xd8, 0x96, 0x29, 0xe6, 0x8f, 0x5d, 0xd2, 0x66, 0x0e, 0x7e, 0xda,
	0xa4, 0xb6, 0xfa, 0x0a, 0x2e, 0xba, 0xea, 0xee, 0x3c, 0x34, 0xe3, 0xa9, 0xa8, 0x4f, 0x70, 0xd6,
	0x92, 0xbd, 0x76, 0xb7, 0xb0, 0xa8, 0xe1, 0x4c, 0x70, 0xaa, 0x74, 0xe9, 0x12, 0x4f, 0xc2, 0xe3,
	0x0e, 0x7e, 0xd7, 0x25, 0xf6, 0xcc, 0x8d, 0x30, 0xde, 0x89, 0x24, 0xef, 0x72, 0x4f, 0x5a, 0xf3,
	0x4c, 0xf0, 0xb5, 0xe1, 0x6c, 0xf6, 0xa7, 0x7f, 0x3d, 0x4d, 0x70, 0x5b, 0x62, 0xb6, 0x3a, 0x5f,
	0x36, 0xef, 0xec, 0x3d, 0xef, 0x4f, 0x7b, 0x03, 0xac, 0xc7, 0x1f, 0xc3, 0x32, 0x8e, 0x0f, 0xec,
	0xcb, 0x1f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x5e, 0x32, 0x31, 0x03, 0x02, 0x00, 0x00,
}