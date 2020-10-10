// Code generated by protoc-gen-go. DO NOT EDIT.
// source: short_url.proto

package ShortUrl

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

type Code int32

const (
	Code_OK            Code = 0
	Code_ERR_PARAMETER Code = 10
	Code_ERR_SERVICE   Code = 20
	Code_ERR_NOT_FOUND Code = 30
	Code_ERR_UNKNOWN   Code = 100
)

var Code_name = map[int32]string{
	0:   "OK",
	10:  "ERR_PARAMETER",
	20:  "ERR_SERVICE",
	30:  "ERR_NOT_FOUND",
	100: "ERR_UNKNOWN",
}

var Code_value = map[string]int32{
	"OK":            0,
	"ERR_PARAMETER": 10,
	"ERR_SERVICE":   20,
	"ERR_NOT_FOUND": 30,
	"ERR_UNKNOWN":   100,
}

func (x Code) String() string {
	return proto.EnumName(Code_name, int32(x))
}

func (Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fe3da7d805421a0a, []int{0}
}

type SetUrlReq struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetUrlReq) Reset()         { *m = SetUrlReq{} }
func (m *SetUrlReq) String() string { return proto.CompactTextString(m) }
func (*SetUrlReq) ProtoMessage()    {}
func (*SetUrlReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe3da7d805421a0a, []int{0}
}

func (m *SetUrlReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetUrlReq.Unmarshal(m, b)
}
func (m *SetUrlReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetUrlReq.Marshal(b, m, deterministic)
}
func (m *SetUrlReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetUrlReq.Merge(m, src)
}
func (m *SetUrlReq) XXX_Size() int {
	return xxx_messageInfo_SetUrlReq.Size(m)
}
func (m *SetUrlReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SetUrlReq.DiscardUnknown(m)
}

var xxx_messageInfo_SetUrlReq proto.InternalMessageInfo

func (m *SetUrlReq) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type SetUrlRsp struct {
	Code                 Code     `protobuf:"varint,1,opt,name=code,proto3,enum=ShortUrl.Code" json:"code,omitempty"`
	ShortUrl             string   `protobuf:"bytes,2,opt,name=shortUrl,proto3" json:"shortUrl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetUrlRsp) Reset()         { *m = SetUrlRsp{} }
func (m *SetUrlRsp) String() string { return proto.CompactTextString(m) }
func (*SetUrlRsp) ProtoMessage()    {}
func (*SetUrlRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe3da7d805421a0a, []int{1}
}

func (m *SetUrlRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetUrlRsp.Unmarshal(m, b)
}
func (m *SetUrlRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetUrlRsp.Marshal(b, m, deterministic)
}
func (m *SetUrlRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetUrlRsp.Merge(m, src)
}
func (m *SetUrlRsp) XXX_Size() int {
	return xxx_messageInfo_SetUrlRsp.Size(m)
}
func (m *SetUrlRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_SetUrlRsp.DiscardUnknown(m)
}

var xxx_messageInfo_SetUrlRsp proto.InternalMessageInfo

func (m *SetUrlRsp) GetCode() Code {
	if m != nil {
		return m.Code
	}
	return Code_OK
}

func (m *SetUrlRsp) GetShortUrl() string {
	if m != nil {
		return m.ShortUrl
	}
	return ""
}

type GetAfterForwardUrlReq struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAfterForwardUrlReq) Reset()         { *m = GetAfterForwardUrlReq{} }
func (m *GetAfterForwardUrlReq) String() string { return proto.CompactTextString(m) }
func (*GetAfterForwardUrlReq) ProtoMessage()    {}
func (*GetAfterForwardUrlReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe3da7d805421a0a, []int{2}
}

func (m *GetAfterForwardUrlReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAfterForwardUrlReq.Unmarshal(m, b)
}
func (m *GetAfterForwardUrlReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAfterForwardUrlReq.Marshal(b, m, deterministic)
}
func (m *GetAfterForwardUrlReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAfterForwardUrlReq.Merge(m, src)
}
func (m *GetAfterForwardUrlReq) XXX_Size() int {
	return xxx_messageInfo_GetAfterForwardUrlReq.Size(m)
}
func (m *GetAfterForwardUrlReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAfterForwardUrlReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetAfterForwardUrlReq proto.InternalMessageInfo

func (m *GetAfterForwardUrlReq) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type GetAfterForwardUrlRsp struct {
	Code                 Code     `protobuf:"varint,1,opt,name=code,proto3,enum=ShortUrl.Code" json:"code,omitempty"`
	LongUrl              string   `protobuf:"bytes,2,opt,name=longUrl,proto3" json:"longUrl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAfterForwardUrlRsp) Reset()         { *m = GetAfterForwardUrlRsp{} }
func (m *GetAfterForwardUrlRsp) String() string { return proto.CompactTextString(m) }
func (*GetAfterForwardUrlRsp) ProtoMessage()    {}
func (*GetAfterForwardUrlRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe3da7d805421a0a, []int{3}
}

func (m *GetAfterForwardUrlRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAfterForwardUrlRsp.Unmarshal(m, b)
}
func (m *GetAfterForwardUrlRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAfterForwardUrlRsp.Marshal(b, m, deterministic)
}
func (m *GetAfterForwardUrlRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAfterForwardUrlRsp.Merge(m, src)
}
func (m *GetAfterForwardUrlRsp) XXX_Size() int {
	return xxx_messageInfo_GetAfterForwardUrlRsp.Size(m)
}
func (m *GetAfterForwardUrlRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAfterForwardUrlRsp.DiscardUnknown(m)
}

var xxx_messageInfo_GetAfterForwardUrlRsp proto.InternalMessageInfo

func (m *GetAfterForwardUrlRsp) GetCode() Code {
	if m != nil {
		return m.Code
	}
	return Code_OK
}

func (m *GetAfterForwardUrlRsp) GetLongUrl() string {
	if m != nil {
		return m.LongUrl
	}
	return ""
}

func init() {
	proto.RegisterEnum("ShortUrl.Code", Code_name, Code_value)
	proto.RegisterType((*SetUrlReq)(nil), "ShortUrl.SetUrlReq")
	proto.RegisterType((*SetUrlRsp)(nil), "ShortUrl.SetUrlRsp")
	proto.RegisterType((*GetAfterForwardUrlReq)(nil), "ShortUrl.GetAfterForwardUrlReq")
	proto.RegisterType((*GetAfterForwardUrlRsp)(nil), "ShortUrl.GetAfterForwardUrlRsp")
}

func init() {
	proto.RegisterFile("short_url.proto", fileDescriptor_fe3da7d805421a0a)
}

var fileDescriptor_fe3da7d805421a0a = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4d, 0x4f, 0xb3, 0x40,
	0x14, 0x85, 0x4b, 0xdf, 0xa6, 0x6f, 0x7b, 0x1b, 0xcb, 0x78, 0xd5, 0x84, 0x90, 0xf8, 0x11, 0x56,
	0xea, 0x82, 0x45, 0x5d, 0xb9, 0x24, 0x95, 0x1a, 0x43, 0x04, 0x33, 0x53, 0xaa, 0x3b, 0xa2, 0x65,
	0xfc, 0x48, 0x88, 0x83, 0xc3, 0x54, 0xff, 0x91, 0xbf, 0xd3, 0x0c, 0x71, 0x60, 0xd1, 0x6a, 0xdc,
	0xdd, 0x7b, 0xce, 0x93, 0xc3, 0xe1, 0x0e, 0xd8, 0xd5, 0xb3, 0x90, 0x2a, 0x5b, 0xc9, 0xc2, 0x2f,
	0xa5, 0x50, 0x02, 0x07, 0x4c, 0x0b, 0xa9, 0x2c, 0xbc, 0x7d, 0x18, 0x32, 0xae, 0x27, 0xca, 0xdf,
	0x90, 0xc0, 0xbf, 0x95, 0x2c, 0x1c, 0xeb, 0xc8, 0x3a, 0x1e, 0x52, 0x3d, 0x7a, 0x51, 0x63, 0x57,
	0x25, 0x7a, 0xd0, 0x5b, 0x8a, 0x9c, 0xd7, 0xfe, 0x78, 0x32, 0xf6, 0x4d, 0x88, 0x3f, 0x15, 0x39,
	0xa7, 0xb5, 0x87, 0x2e, 0x0c, 0xaa, 0x6f, 0xd9, 0xe9, 0xd6, 0x39, 0xcd, 0xee, 0x9d, 0xc0, 0xde,
	0x25, 0x57, 0xc1, 0xa3, 0xe2, 0x72, 0x26, 0xe4, 0xc7, 0xbd, 0xcc, 0x7f, 0xfc, 0x6e, 0xba, 0x11,
	0xfd, 0x63, 0x07, 0x07, 0xfe, 0x17, 0xe2, 0xf5, 0xa9, 0xad, 0x60, 0xd6, 0xd3, 0x05, 0xf4, 0x34,
	0x87, 0x7d, 0xe8, 0x26, 0x11, 0xe9, 0xe0, 0x36, 0x6c, 0x85, 0x94, 0x66, 0x37, 0x01, 0x0d, 0xae,
	0xc3, 0x79, 0x48, 0x09, 0xa0, 0x0d, 0x23, 0x2d, 0xb1, 0x90, 0x2e, 0xae, 0xa6, 0x21, 0xd9, 0x35,
	0x4c, 0x9c, 0xcc, 0xb3, 0x59, 0x92, 0xc6, 0x17, 0xe4, 0xc0, 0x30, 0x69, 0x1c, 0xc5, 0xc9, 0x6d,
	0x4c, 0xf2, 0xc9, 0xa7, 0x05, 0xb6, 0x69, 0xc2, 0xb8, 0x7c, 0x7f, 0x59, 0x72, 0x3c, 0x87, 0x11,
	0xe3, 0xca, 0xa8, 0xb8, 0xd3, 0x56, 0x6d, 0x0e, 0xee, 0xae, 0x8b, 0x55, 0xe9, 0x75, 0xf0, 0x0e,
	0x70, 0xfd, 0xef, 0xf1, 0xb0, 0x85, 0x37, 0x9e, 0xd1, 0xfd, 0x1d, 0xd0, 0xc9, 0x0f, 0xfd, 0xfa,
	0xfd, 0xcf, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xbb, 0xbe, 0x71, 0x12, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ShortUrlServiceClient is the client API for ShortUrlService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShortUrlServiceClient interface {
	SetShortUrl(ctx context.Context, in *SetUrlReq, opts ...grpc.CallOption) (*SetUrlRsp, error)
	GetAfterForwardUrl(ctx context.Context, in *GetAfterForwardUrlReq, opts ...grpc.CallOption) (*GetAfterForwardUrlRsp, error)
}

type shortUrlServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShortUrlServiceClient(cc grpc.ClientConnInterface) ShortUrlServiceClient {
	return &shortUrlServiceClient{cc}
}

func (c *shortUrlServiceClient) SetShortUrl(ctx context.Context, in *SetUrlReq, opts ...grpc.CallOption) (*SetUrlRsp, error) {
	out := new(SetUrlRsp)
	err := c.cc.Invoke(ctx, "/ShortUrl.ShortUrlService/SetShortUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortUrlServiceClient) GetAfterForwardUrl(ctx context.Context, in *GetAfterForwardUrlReq, opts ...grpc.CallOption) (*GetAfterForwardUrlRsp, error) {
	out := new(GetAfterForwardUrlRsp)
	err := c.cc.Invoke(ctx, "/ShortUrl.ShortUrlService/GetAfterForwardUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortUrlServiceServer is the server API for ShortUrlService service.
type ShortUrlServiceServer interface {
	SetShortUrl(context.Context, *SetUrlReq) (*SetUrlRsp, error)
	GetAfterForwardUrl(context.Context, *GetAfterForwardUrlReq) (*GetAfterForwardUrlRsp, error)
}

// UnimplementedShortUrlServiceServer can be embedded to have forward compatible implementations.
type UnimplementedShortUrlServiceServer struct {
}

func (*UnimplementedShortUrlServiceServer) SetShortUrl(ctx context.Context, req *SetUrlReq) (*SetUrlRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetShortUrl not implemented")
}
func (*UnimplementedShortUrlServiceServer) GetAfterForwardUrl(ctx context.Context, req *GetAfterForwardUrlReq) (*GetAfterForwardUrlRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAfterForwardUrl not implemented")
}

func RegisterShortUrlServiceServer(s *grpc.Server, srv ShortUrlServiceServer) {
	s.RegisterService(&_ShortUrlService_serviceDesc, srv)
}

func _ShortUrlService_SetShortUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUrlReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortUrlServiceServer).SetShortUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShortUrl.ShortUrlService/SetShortUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortUrlServiceServer).SetShortUrl(ctx, req.(*SetUrlReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortUrlService_GetAfterForwardUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAfterForwardUrlReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortUrlServiceServer).GetAfterForwardUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShortUrl.ShortUrlService/GetAfterForwardUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortUrlServiceServer).GetAfterForwardUrl(ctx, req.(*GetAfterForwardUrlReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShortUrlService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ShortUrl.ShortUrlService",
	HandlerType: (*ShortUrlServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetShortUrl",
			Handler:    _ShortUrlService_SetShortUrl_Handler,
		},
		{
			MethodName: "GetAfterForwardUrl",
			Handler:    _ShortUrlService_GetAfterForwardUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "short_url.proto",
}
