// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blog/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// this line is used by starport scaffolding # 3
type QueryGetHelloRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryGetHelloRequest) Reset()         { *m = QueryGetHelloRequest{} }
func (m *QueryGetHelloRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetHelloRequest) ProtoMessage()    {}
func (*QueryGetHelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e22cd6d352f3384, []int{0}
}
func (m *QueryGetHelloRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetHelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetHelloRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetHelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetHelloRequest.Merge(m, src)
}
func (m *QueryGetHelloRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetHelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetHelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetHelloRequest proto.InternalMessageInfo

func (m *QueryGetHelloRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type QueryGetHelloResponse struct {
	Hello *Hello `protobuf:"bytes,1,opt,name=Hello,proto3" json:"Hello,omitempty"`
}

func (m *QueryGetHelloResponse) Reset()         { *m = QueryGetHelloResponse{} }
func (m *QueryGetHelloResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetHelloResponse) ProtoMessage()    {}
func (*QueryGetHelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e22cd6d352f3384, []int{1}
}
func (m *QueryGetHelloResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetHelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetHelloResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetHelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetHelloResponse.Merge(m, src)
}
func (m *QueryGetHelloResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetHelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetHelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetHelloResponse proto.InternalMessageInfo

func (m *QueryGetHelloResponse) GetHello() *Hello {
	if m != nil {
		return m.Hello
	}
	return nil
}

type QueryAllHelloRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllHelloRequest) Reset()         { *m = QueryAllHelloRequest{} }
func (m *QueryAllHelloRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllHelloRequest) ProtoMessage()    {}
func (*QueryAllHelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e22cd6d352f3384, []int{2}
}
func (m *QueryAllHelloRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllHelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllHelloRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllHelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllHelloRequest.Merge(m, src)
}
func (m *QueryAllHelloRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllHelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllHelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllHelloRequest proto.InternalMessageInfo

func (m *QueryAllHelloRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllHelloResponse struct {
	Hello      []*Hello            `protobuf:"bytes,1,rep,name=Hello,proto3" json:"Hello,omitempty"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllHelloResponse) Reset()         { *m = QueryAllHelloResponse{} }
func (m *QueryAllHelloResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllHelloResponse) ProtoMessage()    {}
func (*QueryAllHelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e22cd6d352f3384, []int{3}
}
func (m *QueryAllHelloResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllHelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllHelloResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllHelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllHelloResponse.Merge(m, src)
}
func (m *QueryAllHelloResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllHelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllHelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllHelloResponse proto.InternalMessageInfo

func (m *QueryAllHelloResponse) GetHello() []*Hello {
	if m != nil {
		return m.Hello
	}
	return nil
}

func (m *QueryAllHelloResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetHelloRequest)(nil), "alice.blog.blog.QueryGetHelloRequest")
	proto.RegisterType((*QueryGetHelloResponse)(nil), "alice.blog.blog.QueryGetHelloResponse")
	proto.RegisterType((*QueryAllHelloRequest)(nil), "alice.blog.blog.QueryAllHelloRequest")
	proto.RegisterType((*QueryAllHelloResponse)(nil), "alice.blog.blog.QueryAllHelloResponse")
}

func init() { proto.RegisterFile("blog/query.proto", fileDescriptor_6e22cd6d352f3384) }

var fileDescriptor_6e22cd6d352f3384 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x4b, 0x02, 0x41,
	0x14, 0xc7, 0xdd, 0x2d, 0x23, 0x26, 0xa8, 0x18, 0x4a, 0xc4, 0x6a, 0x89, 0x8d, 0x2c, 0x22, 0x66,
	0xd0, 0xae, 0x5d, 0x0c, 0xca, 0x8e, 0xe5, 0xb1, 0x43, 0x30, 0xab, 0xc3, 0x3a, 0x30, 0xee, 0xac,
	0xce, 0x18, 0x89, 0x74, 0xe9, 0x03, 0x44, 0xd0, 0x97, 0xea, 0x28, 0x74, 0xe9, 0x18, 0xda, 0xf7,
	0x28, 0x76, 0x66, 0x40, 0x5d, 0x15, 0x2f, 0xbb, 0xf0, 0xf8, 0xbf, 0xf9, 0xfd, 0xde, 0xbc, 0x01,
	0xdb, 0x01, 0x17, 0x21, 0x6e, 0x77, 0x69, 0xa7, 0x87, 0xe2, 0x8e, 0x50, 0x02, 0x6e, 0x11, 0xce,
	0xea, 0x14, 0x25, 0x75, 0xfd, 0x29, 0xec, 0x87, 0x42, 0x84, 0x9c, 0x62, 0x12, 0x33, 0x4c, 0xa2,
	0x48, 0x28, 0xa2, 0x98, 0x88, 0xa4, 0x89, 0x17, 0xce, 0xea, 0x42, 0xb6, 0x84, 0xc4, 0x01, 0x91,
	0xd4, 0x9c, 0x83, 0x9f, 0x4a, 0x01, 0x55, 0xa4, 0x84, 0x63, 0x12, 0xb2, 0x48, 0x87, 0x6d, 0xd6,
	0xc0, 0x9a, 0x94, 0x73, 0x61, 0x2a, 0x7e, 0x11, 0xec, 0xdc, 0x27, 0x3d, 0x55, 0xaa, 0x6e, 0x93,
	0x72, 0x8d, 0xb6, 0xbb, 0x54, 0x2a, 0xb8, 0x09, 0x5c, 0xd6, 0xc8, 0x3b, 0x87, 0xce, 0xe9, 0x6a,
	0xcd, 0x65, 0x0d, 0xff, 0x1a, 0xec, 0xa6, 0x72, 0x32, 0x16, 0x91, 0xa4, 0xf0, 0x1c, 0x64, 0x75,
	0x41, 0x67, 0x37, 0xca, 0x39, 0x94, 0xb2, 0x47, 0x26, 0x6e, 0x42, 0xfe, 0xa3, 0xc5, 0x55, 0x38,
	0x9f, 0xc2, 0xdd, 0x00, 0x30, 0x96, 0xb5, 0x47, 0x15, 0x91, 0x99, 0x0c, 0x25, 0x93, 0x21, 0x73,
	0x43, 0x76, 0x32, 0x74, 0x47, 0x42, 0x6a, 0x7b, 0x6b, 0x13, 0x9d, 0xfe, 0x9b, 0x63, 0x3d, 0xc7,
	0x80, 0x59, 0xcf, 0x95, 0xa5, 0x9e, 0xb0, 0x3a, 0xe5, 0xe3, 0x6a, 0x9f, 0x93, 0xa5, 0x3e, 0x06,
	0x35, 0x29, 0x54, 0xfe, 0x73, 0x40, 0x56, 0x0b, 0xc1, 0xbe, 0x15, 0x80, 0xc7, 0x33, 0xe8, 0x79,
	0x1b, 0x28, 0x14, 0x97, 0xc5, 0x0c, 0xcd, 0x3f, 0x7a, 0xfd, 0xfa, 0xfd, 0x70, 0x0f, 0xe0, 0x1e,
	0xd6, 0x79, 0xac, 0x57, 0x3c, 0xde, 0x33, 0xee, 0xb3, 0xc6, 0x0b, 0xec, 0x81, 0x75, 0xdd, 0x55,
	0xe1, 0x7c, 0x11, 0x3f, 0xb5, 0x92, 0x45, 0xfc, 0xf4, 0xc5, 0xfa, 0x9e, 0xe6, 0xe7, 0x61, 0x6e,
	0x3e, 0xff, 0xea, 0xf2, 0x73, 0xe8, 0x39, 0x83, 0xa1, 0xe7, 0xfc, 0x0c, 0x3d, 0xe7, 0x7d, 0xe4,
	0x65, 0x06, 0x23, 0x2f, 0xf3, 0x3d, 0xf2, 0x32, 0x0f, 0x7e, 0xc8, 0x54, 0xb3, 0x1b, 0xa0, 0xba,
	0x68, 0x4d, 0xf6, 0x3e, 0x9b, 0x9f, 0xea, 0xc5, 0x54, 0x06, 0x6b, 0xfa, 0x99, 0x5e, 0xfc, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x55, 0xc2, 0x55, 0xcf, 0x27, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Queries a hello by id.
	Hello(ctx context.Context, in *QueryGetHelloRequest, opts ...grpc.CallOption) (*QueryGetHelloResponse, error)
	// Queries a list of hello items.
	HelloAll(ctx context.Context, in *QueryAllHelloRequest, opts ...grpc.CallOption) (*QueryAllHelloResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Hello(ctx context.Context, in *QueryGetHelloRequest, opts ...grpc.CallOption) (*QueryGetHelloResponse, error) {
	out := new(QueryGetHelloResponse)
	err := c.cc.Invoke(ctx, "/alice.blog.blog.Query/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) HelloAll(ctx context.Context, in *QueryAllHelloRequest, opts ...grpc.CallOption) (*QueryAllHelloResponse, error) {
	out := new(QueryAllHelloResponse)
	err := c.cc.Invoke(ctx, "/alice.blog.blog.Query/HelloAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a hello by id.
	Hello(context.Context, *QueryGetHelloRequest) (*QueryGetHelloResponse, error)
	// Queries a list of hello items.
	HelloAll(context.Context, *QueryAllHelloRequest) (*QueryAllHelloResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Hello(ctx context.Context, req *QueryGetHelloRequest) (*QueryGetHelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (*UnimplementedQueryServer) HelloAll(ctx context.Context, req *QueryAllHelloRequest) (*QueryAllHelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetHelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alice.blog.blog.Query/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Hello(ctx, req.(*QueryGetHelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_HelloAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllHelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).HelloAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alice.blog.blog.Query/HelloAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).HelloAll(ctx, req.(*QueryAllHelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "alice.blog.blog.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Query_Hello_Handler,
		},
		{
			MethodName: "HelloAll",
			Handler:    _Query_HelloAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blog/query.proto",
}

func (m *QueryGetHelloRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetHelloRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetHelloRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetHelloResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetHelloResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetHelloResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Hello != nil {
		{
			size, err := m.Hello.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllHelloRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllHelloRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllHelloRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllHelloResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllHelloResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllHelloResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Hello) > 0 {
		for iNdEx := len(m.Hello) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Hello[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetHelloRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	return n
}

func (m *QueryGetHelloResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Hello != nil {
		l = m.Hello.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllHelloRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllHelloResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Hello) > 0 {
		for _, e := range m.Hello {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetHelloRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetHelloRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetHelloRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetHelloResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetHelloResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetHelloResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hello", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Hello == nil {
				m.Hello = &Hello{}
			}
			if err := m.Hello.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllHelloRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllHelloRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllHelloRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllHelloResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllHelloResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllHelloResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hello", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hello = append(m.Hello, &Hello{})
			if err := m.Hello[len(m.Hello)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
