// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/account/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type QueryIBCAccountRequest struct {
	// address is the address to query.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QueryIBCAccountRequest) Reset()         { *m = QueryIBCAccountRequest{} }
func (m *QueryIBCAccountRequest) String() string { return proto.CompactTextString(m) }
func (*QueryIBCAccountRequest) ProtoMessage()    {}
func (*QueryIBCAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5738b2fac406b004, []int{0}
}
func (m *QueryIBCAccountRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIBCAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIBCAccountRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIBCAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIBCAccountRequest.Merge(m, src)
}
func (m *QueryIBCAccountRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryIBCAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIBCAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIBCAccountRequest proto.InternalMessageInfo

type QueryIBCAccountFromDataRequest struct {
	Port    string `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty"`
	Channel string `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	Data    string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *QueryIBCAccountFromDataRequest) Reset()         { *m = QueryIBCAccountFromDataRequest{} }
func (m *QueryIBCAccountFromDataRequest) String() string { return proto.CompactTextString(m) }
func (*QueryIBCAccountFromDataRequest) ProtoMessage()    {}
func (*QueryIBCAccountFromDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5738b2fac406b004, []int{1}
}
func (m *QueryIBCAccountFromDataRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIBCAccountFromDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIBCAccountFromDataRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIBCAccountFromDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIBCAccountFromDataRequest.Merge(m, src)
}
func (m *QueryIBCAccountFromDataRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryIBCAccountFromDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIBCAccountFromDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIBCAccountFromDataRequest proto.InternalMessageInfo

type QueryIBCAccountResponse struct {
	// account defines the account of the corresponding address.
	Account *IBCAccount `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (m *QueryIBCAccountResponse) Reset()         { *m = QueryIBCAccountResponse{} }
func (m *QueryIBCAccountResponse) String() string { return proto.CompactTextString(m) }
func (*QueryIBCAccountResponse) ProtoMessage()    {}
func (*QueryIBCAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5738b2fac406b004, []int{2}
}
func (m *QueryIBCAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIBCAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIBCAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIBCAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIBCAccountResponse.Merge(m, src)
}
func (m *QueryIBCAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryIBCAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIBCAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIBCAccountResponse proto.InternalMessageInfo

func (m *QueryIBCAccountResponse) GetAccount() *IBCAccount {
	if m != nil {
		return m.Account
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryIBCAccountRequest)(nil), "ibc.account.QueryIBCAccountRequest")
	proto.RegisterType((*QueryIBCAccountFromDataRequest)(nil), "ibc.account.QueryIBCAccountFromDataRequest")
	proto.RegisterType((*QueryIBCAccountResponse)(nil), "ibc.account.QueryIBCAccountResponse")
}

func init() { proto.RegisterFile("ibc/account/query.proto", fileDescriptor_5738b2fac406b004) }

var fileDescriptor_5738b2fac406b004 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x3d, 0xaf, 0xd3, 0x30,
	0x14, 0x4d, 0xca, 0xc7, 0x03, 0xbf, 0xcd, 0x42, 0x34, 0x54, 0x28, 0x0f, 0x15, 0x06, 0x24, 0xd4,
	0x58, 0x79, 0x6f, 0x2a, 0x62, 0xa1, 0x20, 0xa4, 0x0a, 0x96, 0x96, 0x8d, 0xcd, 0x71, 0x4d, 0x1a,
	0xa9, 0xf1, 0x4d, 0x63, 0x07, 0x51, 0x45, 0x59, 0x98, 0x18, 0x11, 0xfc, 0x81, 0xfe, 0x0e, 0x7e,
	0x01, 0x63, 0x25, 0x16, 0xd8, 0x50, 0xcb, 0xc0, 0xcf, 0x40, 0x4e, 0x1c, 0x48, 0xdb, 0xa7, 0xaa,
	0x53, 0xaf, 0xef, 0x3d, 0xf7, 0x9c, 0xd3, 0x73, 0x83, 0xda, 0x51, 0xc0, 0x08, 0x65, 0x0c, 0x32,
	0xa1, 0xc8, 0x3c, 0xe3, 0xe9, 0xc2, 0x4b, 0x52, 0x50, 0x80, 0x4f, 0xa3, 0x80, 0x79, 0x66, 0xd0,
	0xb9, 0x15, 0x42, 0x08, 0x65, 0x9f, 0xe8, 0xaa, 0x82, 0x74, 0xee, 0x86, 0x00, 0xe1, 0x8c, 0x13,
	0x9a, 0x44, 0x84, 0x0a, 0x01, 0x8a, 0xaa, 0x08, 0x84, 0x34, 0xd3, 0x3b, 0x4d, 0x66, 0xf3, 0x5b,
	0x8d, 0xba, 0x4f, 0xd0, 0xed, 0x91, 0x96, 0x1a, 0x0e, 0x9e, 0x3d, 0xad, 0x06, 0x63, 0x3e, 0xcf,
	0xb8, 0x54, 0xd8, 0x41, 0x27, 0x74, 0x32, 0x49, 0xb9, 0x94, 0x8e, 0x7d, 0xcf, 0x7e, 0x78, 0x73,
	0x5c, 0x3f, 0x1f, 0xdf, 0xf8, 0xb8, 0x3c, 0xb3, 0xfe, 0x2c, 0xcf, 0xac, 0xee, 0x0c, 0xb9, 0x3b,
	0xdb, 0x2f, 0x52, 0x88, 0x9f, 0x53, 0x45, 0x6b, 0x16, 0x8c, 0xae, 0x26, 0x90, 0x2a, 0x43, 0x51,
	0xd6, 0x9a, 0x99, 0x4d, 0xa9, 0x10, 0x7c, 0xe6, 0xb4, 0x2a, 0x66, 0xf3, 0xd4, 0xe8, 0x09, 0x55,
	0xd4, 0xb9, 0x52, 0xa1, 0x75, 0xdd, 0x50, 0x7b, 0x85, 0xda, 0x7b, 0x5e, 0x65, 0x02, 0x42, 0x72,
	0xec, 0xa3, 0x13, 0xf3, 0xbf, 0x4a, 0xa5, 0xd3, 0xf3, 0xb6, 0xd7, 0x08, 0xcd, 0x6b, 0x6c, 0xd4,
	0xb8, 0xf3, 0x9f, 0x2d, 0x74, 0xad, 0xa4, 0xc3, 0x9f, 0x6d, 0x84, 0xfe, 0x23, 0xf0, 0xfd, 0xad,
	0xd5, 0xcb, 0xd3, 0xe9, 0x3c, 0x38, 0x0c, 0xaa, 0x6c, 0x75, 0xfb, 0x1f, 0xbe, 0xff, 0xfe, 0xd2,
	0xba, 0xc0, 0x3e, 0x61, 0x20, 0x63, 0x90, 0x24, 0x0a, 0x58, 0xaf, 0x3e, 0xc4, 0x3b, 0x3f, 0xe0,
	0x8a, 0xfa, 0x5b, 0xbd, 0xdc, 0x64, 0x5c, 0xe0, 0xaf, 0x36, 0xc2, 0xfb, 0xb1, 0xe2, 0x47, 0x87,
	0x74, 0x77, 0xc2, 0x3f, 0xd2, 0xe4, 0xa8, 0x34, 0xf9, 0x12, 0x0f, 0x8f, 0x34, 0xd9, 0x7b, 0x9b,
	0x42, 0xdc, 0xd3, 0xb7, 0x21, 0xb9, 0xbe, 0x67, 0x41, 0x72, 0x73, 0xbf, 0x82, 0xe4, 0xba, 0x5d,
	0x0c, 0x5e, 0x7f, 0x5b, 0xbb, 0xf6, 0x6a, 0xed, 0xda, 0xbf, 0xd6, 0xae, 0xfd, 0x69, 0xe3, 0x5a,
	0xab, 0x8d, 0x6b, 0xfd, 0xd8, 0xb8, 0xd6, 0x9b, 0x7e, 0x18, 0xa9, 0x69, 0x16, 0x78, 0x0c, 0xe2,
	0x7f, 0x72, 0x42, 0xf1, 0x94, 0x4d, 0x69, 0x24, 0x6a, 0x05, 0x49, 0xde, 0x6f, 0x99, 0x50, 0x8b,
	0x84, 0xcb, 0xe0, 0x7a, 0xf9, 0xc5, 0x5e, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x72, 0x6d, 0x10,
	0x04, 0x28, 0x03, 0x00, 0x00,
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
	IBCAccount(ctx context.Context, in *QueryIBCAccountRequest, opts ...grpc.CallOption) (*QueryIBCAccountResponse, error)
	IBCAccountFromData(ctx context.Context, in *QueryIBCAccountFromDataRequest, opts ...grpc.CallOption) (*QueryIBCAccountResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) IBCAccount(ctx context.Context, in *QueryIBCAccountRequest, opts ...grpc.CallOption) (*QueryIBCAccountResponse, error) {
	out := new(QueryIBCAccountResponse)
	err := c.cc.Invoke(ctx, "/ibc.account.Query/IBCAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) IBCAccountFromData(ctx context.Context, in *QueryIBCAccountFromDataRequest, opts ...grpc.CallOption) (*QueryIBCAccountResponse, error) {
	out := new(QueryIBCAccountResponse)
	err := c.cc.Invoke(ctx, "/ibc.account.Query/IBCAccountFromData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	IBCAccount(context.Context, *QueryIBCAccountRequest) (*QueryIBCAccountResponse, error)
	IBCAccountFromData(context.Context, *QueryIBCAccountFromDataRequest) (*QueryIBCAccountResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) IBCAccount(ctx context.Context, req *QueryIBCAccountRequest) (*QueryIBCAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IBCAccount not implemented")
}
func (*UnimplementedQueryServer) IBCAccountFromData(ctx context.Context, req *QueryIBCAccountFromDataRequest) (*QueryIBCAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IBCAccountFromData not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_IBCAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryIBCAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).IBCAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.account.Query/IBCAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IBCAccount(ctx, req.(*QueryIBCAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_IBCAccountFromData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryIBCAccountFromDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).IBCAccountFromData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.account.Query/IBCAccountFromData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IBCAccountFromData(ctx, req.(*QueryIBCAccountFromDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ibc.account.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IBCAccount",
			Handler:    _Query_IBCAccount_Handler,
		},
		{
			MethodName: "IBCAccountFromData",
			Handler:    _Query_IBCAccountFromData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc/account/query.proto",
}

func (m *QueryIBCAccountRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIBCAccountRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIBCAccountRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryIBCAccountFromDataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIBCAccountFromDataRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIBCAccountFromDataRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Channel) > 0 {
		i -= len(m.Channel)
		copy(dAtA[i:], m.Channel)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Channel)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Port) > 0 {
		i -= len(m.Port)
		copy(dAtA[i:], m.Port)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Port)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryIBCAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIBCAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIBCAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Account != nil {
		{
			size, err := m.Account.MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryIBCAccountRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryIBCAccountFromDataRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Port)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Channel)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryIBCAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Account != nil {
		l = m.Account.Size()
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
func (m *QueryIBCAccountRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryIBCAccountRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIBCAccountRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
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
func (m *QueryIBCAccountFromDataRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryIBCAccountFromDataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIBCAccountFromDataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Port = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
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
func (m *QueryIBCAccountResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryIBCAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIBCAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
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
			if m.Account == nil {
				m.Account = &IBCAccount{}
			}
			if err := m.Account.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
