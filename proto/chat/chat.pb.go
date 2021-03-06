// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

package chat

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type StringData struct {
	From                 int64    `protobuf:"zigzag64,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   int64    `protobuf:"zigzag64,2,opt,name=to,proto3" json:"to,omitempty"`
	Type                 int64    `protobuf:"zigzag64,3,opt,name=type,proto3" json:"type,omitempty"`
	Datatype             int64    `protobuf:"zigzag64,4,opt,name=datatype,proto3" json:"datatype,omitempty"`
	Data                 string   `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringData) Reset()         { *m = StringData{} }
func (m *StringData) String() string { return proto.CompactTextString(m) }
func (*StringData) ProtoMessage()    {}
func (*StringData) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
}

func (m *StringData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringData.Unmarshal(m, b)
}
func (m *StringData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringData.Marshal(b, m, deterministic)
}
func (m *StringData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringData.Merge(m, src)
}
func (m *StringData) XXX_Size() int {
	return xxx_messageInfo_StringData.Size(m)
}
func (m *StringData) XXX_DiscardUnknown() {
	xxx_messageInfo_StringData.DiscardUnknown(m)
}

var xxx_messageInfo_StringData proto.InternalMessageInfo

func (m *StringData) GetFrom() int64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *StringData) GetTo() int64 {
	if m != nil {
		return m.To
	}
	return 0
}

func (m *StringData) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *StringData) GetDatatype() int64 {
	if m != nil {
		return m.Datatype
	}
	return 0
}

func (m *StringData) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type FileData struct {
	File                 []byte   `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileData) Reset()         { *m = FileData{} }
func (m *FileData) String() string { return proto.CompactTextString(m) }
func (*FileData) ProtoMessage()    {}
func (*FileData) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{1}
}

func (m *FileData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileData.Unmarshal(m, b)
}
func (m *FileData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileData.Marshal(b, m, deterministic)
}
func (m *FileData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileData.Merge(m, src)
}
func (m *FileData) XXX_Size() int {
	return xxx_messageInfo_FileData.Size(m)
}
func (m *FileData) XXX_DiscardUnknown() {
	xxx_messageInfo_FileData.DiscardUnknown(m)
}

var xxx_messageInfo_FileData proto.InternalMessageInfo

func (m *FileData) GetFile() []byte {
	if m != nil {
		return m.File
	}
	return nil
}

type ChatRequest struct {
	Data                 *StringData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	File                 *FileData   `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ChatRequest) Reset()         { *m = ChatRequest{} }
func (m *ChatRequest) String() string { return proto.CompactTextString(m) }
func (*ChatRequest) ProtoMessage()    {}
func (*ChatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{2}
}

func (m *ChatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatRequest.Unmarshal(m, b)
}
func (m *ChatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatRequest.Marshal(b, m, deterministic)
}
func (m *ChatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatRequest.Merge(m, src)
}
func (m *ChatRequest) XXX_Size() int {
	return xxx_messageInfo_ChatRequest.Size(m)
}
func (m *ChatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChatRequest proto.InternalMessageInfo

func (m *ChatRequest) GetData() *StringData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ChatRequest) GetFile() *FileData {
	if m != nil {
		return m.File
	}
	return nil
}

type ChatResponse struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatResponse) Reset()         { *m = ChatResponse{} }
func (m *ChatResponse) String() string { return proto.CompactTextString(m) }
func (*ChatResponse) ProtoMessage()    {}
func (*ChatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{3}
}

func (m *ChatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatResponse.Unmarshal(m, b)
}
func (m *ChatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatResponse.Marshal(b, m, deterministic)
}
func (m *ChatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatResponse.Merge(m, src)
}
func (m *ChatResponse) XXX_Size() int {
	return xxx_messageInfo_ChatResponse.Size(m)
}
func (m *ChatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ChatResponse proto.InternalMessageInfo

func (m *ChatResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ChatResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*StringData)(nil), "chat.StringData")
	proto.RegisterType((*FileData)(nil), "chat.FileData")
	proto.RegisterType((*ChatRequest)(nil), "chat.ChatRequest")
	proto.RegisterType((*ChatResponse)(nil), "chat.ChatResponse")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_8c585a45e2093e54) }

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4f, 0x02, 0x31,
	0x10, 0x85, 0xdd, 0x75, 0x41, 0x18, 0x08, 0xd1, 0x39, 0x98, 0x86, 0x83, 0x21, 0x8d, 0x07, 0x4e,
	0x1c, 0x30, 0x1e, 0x4d, 0x4c, 0x34, 0xfe, 0x80, 0x7a, 0xf0, 0x5c, 0x71, 0x84, 0x4d, 0x80, 0xae,
	0xdb, 0xc1, 0xc4, 0x7f, 0x6f, 0x3a, 0x53, 0x60, 0x6f, 0xef, 0xbd, 0xbc, 0x7c, 0xaf, 0x53, 0x80,
	0xd5, 0xc6, 0xf3, 0xa2, 0x69, 0x03, 0x07, 0xac, 0x92, 0xb6, 0x0c, 0xf0, 0xce, 0x6d, 0xbd, 0x5f,
	0xbf, 0x7a, 0xf6, 0x88, 0x50, 0x7d, 0xb7, 0x61, 0x67, 0x8a, 0x59, 0x31, 0x47, 0x27, 0x1a, 0x27,
	0x50, 0x72, 0x30, 0xa5, 0x24, 0x25, 0x87, 0xd4, 0xe1, 0xbf, 0x86, 0xcc, 0xa5, 0x76, 0x92, 0xc6,
	0x29, 0x0c, 0xbe, 0x3c, 0x7b, 0xc9, 0x2b, 0xc9, 0x4f, 0x3e, 0xf5, 0x93, 0x36, 0xbd, 0x59, 0x31,
	0x1f, 0x3a, 0xd1, 0xf6, 0x0e, 0x06, 0x6f, 0xf5, 0x96, 0x4e, 0x9b, 0xf5, 0x96, 0x64, 0x73, 0xec,
	0x44, 0xdb, 0x0f, 0x18, 0xbd, 0x6c, 0x3c, 0x3b, 0xfa, 0x39, 0x50, 0x64, 0xbc, 0xcf, 0x88, 0x54,
	0x19, 0x2d, 0xaf, 0x17, 0x72, 0xc5, 0xf9, 0xd9, 0x0a, 0x45, 0x9b, 0x41, 0xa5, 0xb4, 0x26, 0xda,
	0x3a, 0xce, 0x64, 0xf0, 0x33, 0x8c, 0x15, 0x1c, 0x9b, 0xb0, 0x8f, 0x84, 0xb7, 0xd0, 0x8f, 0xec,
	0xf9, 0x10, 0x85, 0xdd, 0x73, 0xd9, 0xa1, 0x81, 0xab, 0x1d, 0xc5, 0xe8, 0xd7, 0x8a, 0x1b, 0xba,
	0xa3, 0x5d, 0x3e, 0x41, 0x95, 0x08, 0xf8, 0x08, 0xe0, 0x68, 0x55, 0xff, 0xea, 0x11, 0x37, 0xba,
	0xd6, 0x79, 0xf4, 0x14, 0xbb, 0x91, 0xce, 0xd9, 0x8b, 0xcf, 0xbe, 0x7c, 0xfe, 0xc3, 0x7f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x4c, 0xee, 0x19, 0x5f, 0x8a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatClient interface {
	ReciveData(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*ChatResponse, error)
}

type chatClient struct {
	cc *grpc.ClientConn
}

func NewChatClient(cc *grpc.ClientConn) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) ReciveData(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*ChatResponse, error) {
	out := new(ChatResponse)
	err := c.cc.Invoke(ctx, "/chat.Chat/ReciveData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServer is the server API for Chat service.
type ChatServer interface {
	ReciveData(context.Context, *ChatRequest) (*ChatResponse, error)
}

// UnimplementedChatServer can be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (*UnimplementedChatServer) ReciveData(ctx context.Context, req *ChatRequest) (*ChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReciveData not implemented")
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_ReciveData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).ReciveData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.Chat/ReciveData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).ReciveData(ctx, req.(*ChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReciveData",
			Handler:    _Chat_ReciveData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
