// Code generated by protoc-gen-go.
// source: api/api.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api/api.proto

It has these top-level messages:
	Client
	CreateClientReq
	CreateClientResp
	DeleteClientReq
	DeleteClientResp
	Password
	CreatePasswordReq
	CreatePasswordResp
	UpdatePasswordReq
	UpdatePasswordResp
	DeletePasswordReq
	DeletePasswordResp
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Client represents an OAuth2 client.
type Client struct {
	Id           string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Secret       string   `protobuf:"bytes,2,opt,name=secret" json:"secret,omitempty"`
	RedirectUris []string `protobuf:"bytes,3,rep,name=redirect_uris,json=redirectUris" json:"redirect_uris,omitempty"`
	TrustedPeers []string `protobuf:"bytes,4,rep,name=trusted_peers,json=trustedPeers" json:"trusted_peers,omitempty"`
	Public       bool     `protobuf:"varint,5,opt,name=public" json:"public,omitempty"`
	Name         string   `protobuf:"bytes,6,opt,name=name" json:"name,omitempty"`
	LogoUrl      string   `protobuf:"bytes,7,opt,name=logo_url,json=logoUrl" json:"logo_url,omitempty"`
}

func (m *Client) Reset()                    { *m = Client{} }
func (m *Client) String() string            { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()               {}
func (*Client) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// CreateClientReq is a request to make a client.
type CreateClientReq struct {
	Client *Client `protobuf:"bytes,1,opt,name=client" json:"client,omitempty"`
}

func (m *CreateClientReq) Reset()                    { *m = CreateClientReq{} }
func (m *CreateClientReq) String() string            { return proto.CompactTextString(m) }
func (*CreateClientReq) ProtoMessage()               {}
func (*CreateClientReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateClientReq) GetClient() *Client {
	if m != nil {
		return m.Client
	}
	return nil
}

// CreateClientResp returns the response from creating a client.
type CreateClientResp struct {
	AlreadyExists bool    `protobuf:"varint,1,opt,name=already_exists,json=alreadyExists" json:"already_exists,omitempty"`
	Client        *Client `protobuf:"bytes,2,opt,name=client" json:"client,omitempty"`
}

func (m *CreateClientResp) Reset()                    { *m = CreateClientResp{} }
func (m *CreateClientResp) String() string            { return proto.CompactTextString(m) }
func (*CreateClientResp) ProtoMessage()               {}
func (*CreateClientResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateClientResp) GetClient() *Client {
	if m != nil {
		return m.Client
	}
	return nil
}

// DeleteClientReq is a request to delete a client.
type DeleteClientReq struct {
	// The ID of the client.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteClientReq) Reset()                    { *m = DeleteClientReq{} }
func (m *DeleteClientReq) String() string            { return proto.CompactTextString(m) }
func (*DeleteClientReq) ProtoMessage()               {}
func (*DeleteClientReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// DeleteClientResp determines if the.
type DeleteClientResp struct {
	NotFound bool `protobuf:"varint,1,opt,name=not_found,json=notFound" json:"not_found,omitempty"`
}

func (m *DeleteClientResp) Reset()                    { *m = DeleteClientResp{} }
func (m *DeleteClientResp) String() string            { return proto.CompactTextString(m) }
func (*DeleteClientResp) ProtoMessage()               {}
func (*DeleteClientResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// Password is an email for password mapping managed by the storage.
type Password struct {
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	// Currently we do not accept plain text passwords. Could be an option in the future.
	Hash     []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username" json:"username,omitempty"`
	UserId   string `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *Password) Reset()                    { *m = Password{} }
func (m *Password) String() string            { return proto.CompactTextString(m) }
func (*Password) ProtoMessage()               {}
func (*Password) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// CreatePasswordReq is a request to make a password.
type CreatePasswordReq struct {
	Password *Password `protobuf:"bytes,1,opt,name=password" json:"password,omitempty"`
}

func (m *CreatePasswordReq) Reset()                    { *m = CreatePasswordReq{} }
func (m *CreatePasswordReq) String() string            { return proto.CompactTextString(m) }
func (*CreatePasswordReq) ProtoMessage()               {}
func (*CreatePasswordReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CreatePasswordReq) GetPassword() *Password {
	if m != nil {
		return m.Password
	}
	return nil
}

// CreatePasswordResp returns the response from creating a password.
type CreatePasswordResp struct {
	AlreadyExists bool `protobuf:"varint,1,opt,name=already_exists,json=alreadyExists" json:"already_exists,omitempty"`
}

func (m *CreatePasswordResp) Reset()                    { *m = CreatePasswordResp{} }
func (m *CreatePasswordResp) String() string            { return proto.CompactTextString(m) }
func (*CreatePasswordResp) ProtoMessage()               {}
func (*CreatePasswordResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// UpdatePasswordReq is a request to modify an existing password.
type UpdatePasswordReq struct {
	// The email used to lookup the password. This field cannot be modified
	Email       string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	NewHash     []byte `protobuf:"bytes,2,opt,name=new_hash,json=newHash,proto3" json:"new_hash,omitempty"`
	NewUsername string `protobuf:"bytes,3,opt,name=new_username,json=newUsername" json:"new_username,omitempty"`
}

func (m *UpdatePasswordReq) Reset()                    { *m = UpdatePasswordReq{} }
func (m *UpdatePasswordReq) String() string            { return proto.CompactTextString(m) }
func (*UpdatePasswordReq) ProtoMessage()               {}
func (*UpdatePasswordReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// UpdatePasswordResp returns the response from modifying an existing password.
type UpdatePasswordResp struct {
	NotFound bool `protobuf:"varint,1,opt,name=not_found,json=notFound" json:"not_found,omitempty"`
}

func (m *UpdatePasswordResp) Reset()                    { *m = UpdatePasswordResp{} }
func (m *UpdatePasswordResp) String() string            { return proto.CompactTextString(m) }
func (*UpdatePasswordResp) ProtoMessage()               {}
func (*UpdatePasswordResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

// DeletePasswordReq is a request to delete a password.
type DeletePasswordReq struct {
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
}

func (m *DeletePasswordReq) Reset()                    { *m = DeletePasswordReq{} }
func (m *DeletePasswordReq) String() string            { return proto.CompactTextString(m) }
func (*DeletePasswordReq) ProtoMessage()               {}
func (*DeletePasswordReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

// DeletePasswordResp returns the response from deleting a password.
type DeletePasswordResp struct {
	NotFound bool `protobuf:"varint,1,opt,name=not_found,json=notFound" json:"not_found,omitempty"`
}

func (m *DeletePasswordResp) Reset()                    { *m = DeletePasswordResp{} }
func (m *DeletePasswordResp) String() string            { return proto.CompactTextString(m) }
func (*DeletePasswordResp) ProtoMessage()               {}
func (*DeletePasswordResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func init() {
	proto.RegisterType((*Client)(nil), "api.Client")
	proto.RegisterType((*CreateClientReq)(nil), "api.CreateClientReq")
	proto.RegisterType((*CreateClientResp)(nil), "api.CreateClientResp")
	proto.RegisterType((*DeleteClientReq)(nil), "api.DeleteClientReq")
	proto.RegisterType((*DeleteClientResp)(nil), "api.DeleteClientResp")
	proto.RegisterType((*Password)(nil), "api.Password")
	proto.RegisterType((*CreatePasswordReq)(nil), "api.CreatePasswordReq")
	proto.RegisterType((*CreatePasswordResp)(nil), "api.CreatePasswordResp")
	proto.RegisterType((*UpdatePasswordReq)(nil), "api.UpdatePasswordReq")
	proto.RegisterType((*UpdatePasswordResp)(nil), "api.UpdatePasswordResp")
	proto.RegisterType((*DeletePasswordReq)(nil), "api.DeletePasswordReq")
	proto.RegisterType((*DeletePasswordResp)(nil), "api.DeletePasswordResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Dex service

type DexClient interface {
	// CreateClient attempts to create the client.
	CreateClient(ctx context.Context, in *CreateClientReq, opts ...grpc.CallOption) (*CreateClientResp, error)
	// DeleteClient attempts to delete the provided client.
	DeleteClient(ctx context.Context, in *DeleteClientReq, opts ...grpc.CallOption) (*DeleteClientResp, error)
	// CreatePassword attempts to create the password.
	CreatePassword(ctx context.Context, in *CreatePasswordReq, opts ...grpc.CallOption) (*CreatePasswordResp, error)
	// UpdatePassword attempts to modify existing password.
	UpdatePassword(ctx context.Context, in *UpdatePasswordReq, opts ...grpc.CallOption) (*UpdatePasswordResp, error)
	// DeletePassword attempts to delete the password.
	DeletePassword(ctx context.Context, in *DeletePasswordReq, opts ...grpc.CallOption) (*DeletePasswordResp, error)
}

type dexClient struct {
	cc *grpc.ClientConn
}

func NewDexClient(cc *grpc.ClientConn) DexClient {
	return &dexClient{cc}
}

func (c *dexClient) CreateClient(ctx context.Context, in *CreateClientReq, opts ...grpc.CallOption) (*CreateClientResp, error) {
	out := new(CreateClientResp)
	err := grpc.Invoke(ctx, "/api.Dex/CreateClient", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) DeleteClient(ctx context.Context, in *DeleteClientReq, opts ...grpc.CallOption) (*DeleteClientResp, error) {
	out := new(DeleteClientResp)
	err := grpc.Invoke(ctx, "/api.Dex/DeleteClient", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) CreatePassword(ctx context.Context, in *CreatePasswordReq, opts ...grpc.CallOption) (*CreatePasswordResp, error) {
	out := new(CreatePasswordResp)
	err := grpc.Invoke(ctx, "/api.Dex/CreatePassword", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) UpdatePassword(ctx context.Context, in *UpdatePasswordReq, opts ...grpc.CallOption) (*UpdatePasswordResp, error) {
	out := new(UpdatePasswordResp)
	err := grpc.Invoke(ctx, "/api.Dex/UpdatePassword", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) DeletePassword(ctx context.Context, in *DeletePasswordReq, opts ...grpc.CallOption) (*DeletePasswordResp, error) {
	out := new(DeletePasswordResp)
	err := grpc.Invoke(ctx, "/api.Dex/DeletePassword", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Dex service

type DexServer interface {
	// CreateClient attempts to create the client.
	CreateClient(context.Context, *CreateClientReq) (*CreateClientResp, error)
	// DeleteClient attempts to delete the provided client.
	DeleteClient(context.Context, *DeleteClientReq) (*DeleteClientResp, error)
	// CreatePassword attempts to create the password.
	CreatePassword(context.Context, *CreatePasswordReq) (*CreatePasswordResp, error)
	// UpdatePassword attempts to modify existing password.
	UpdatePassword(context.Context, *UpdatePasswordReq) (*UpdatePasswordResp, error)
	// DeletePassword attempts to delete the password.
	DeletePassword(context.Context, *DeletePasswordReq) (*DeletePasswordResp, error)
}

func RegisterDexServer(s *grpc.Server, srv DexServer) {
	s.RegisterService(&_Dex_serviceDesc, srv)
}

func _Dex_CreateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).CreateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Dex/CreateClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).CreateClient(ctx, req.(*CreateClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_DeleteClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).DeleteClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Dex/DeleteClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).DeleteClient(ctx, req.(*DeleteClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_CreatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).CreatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Dex/CreatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).CreatePassword(ctx, req.(*CreatePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Dex/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).UpdatePassword(ctx, req.(*UpdatePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_DeletePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).DeletePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Dex/DeletePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).DeletePassword(ctx, req.(*DeletePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dex_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Dex",
	HandlerType: (*DexServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClient",
			Handler:    _Dex_CreateClient_Handler,
		},
		{
			MethodName: "DeleteClient",
			Handler:    _Dex_DeleteClient_Handler,
		},
		{
			MethodName: "CreatePassword",
			Handler:    _Dex_CreatePassword_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _Dex_UpdatePassword_Handler,
		},
		{
			MethodName: "DeletePassword",
			Handler:    _Dex_DeletePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("api/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0xcb, 0x6e, 0xdb, 0x30,
	0x10, 0x8c, 0x1f, 0xb1, 0xe5, 0xf5, 0x23, 0x31, 0x91, 0x26, 0x8a, 0x7b, 0x71, 0x18, 0x14, 0x70,
	0x2e, 0x09, 0x9a, 0x02, 0xbd, 0x14, 0xed, 0xc5, 0x69, 0xd1, 0xde, 0x02, 0x01, 0xbe, 0x56, 0x60,
	0xcc, 0x6d, 0x42, 0x40, 0x91, 0x58, 0x92, 0x82, 0xd3, 0xcf, 0xeb, 0x2f, 0xf4, 0x8b, 0x0a, 0x52,
	0xb4, 0x21, 0xc9, 0x2e, 0xdc, 0x9b, 0x66, 0xb8, 0x9c, 0xe5, 0xcc, 0x2e, 0x04, 0x43, 0x26, 0xc5,
	0x0d, 0x93, 0xe2, 0x5a, 0xaa, 0xcc, 0x64, 0xa4, 0xc5, 0xa4, 0xa0, 0xbf, 0x1b, 0xd0, 0x99, 0x27,
	0x02, 0x53, 0x43, 0x46, 0xd0, 0x14, 0x3c, 0x6c, 0x4c, 0x1b, 0xb3, 0x5e, 0xd4, 0x14, 0x9c, 0x9c,
	0x42, 0x47, 0xe3, 0x52, 0xa1, 0x09, 0x9b, 0x8e, 0xf3, 0x88, 0x5c, 0xc2, 0x50, 0x21, 0x17, 0x0a,
	0x97, 0x26, 0xce, 0x95, 0xd0, 0x61, 0x6b, 0xda, 0x9a, 0xf5, 0xa2, 0xc1, 0x9a, 0x5c, 0x28, 0xa1,
	0x6d, 0x91, 0x51, 0xb9, 0x36, 0xc8, 0x63, 0x89, 0xa8, 0x74, 0xd8, 0x2e, 0x8a, 0x3c, 0x79, 0x6f,
	0x39, 0xdb, 0x41, 0xe6, 0x0f, 0x89, 0x58, 0x86, 0x87, 0xd3, 0xc6, 0x2c, 0x88, 0x3c, 0x22, 0x04,
	0xda, 0x29, 0x7b, 0xc6, 0xb0, 0xe3, 0xfa, 0xba, 0x6f, 0x72, 0x0e, 0x41, 0x92, 0x3d, 0x66, 0x71,
	0xae, 0x92, 0xb0, 0xeb, 0xf8, 0xae, 0xc5, 0x0b, 0x95, 0xd0, 0xf7, 0x70, 0x34, 0x57, 0xc8, 0x0c,
	0x16, 0x46, 0x22, 0xfc, 0x49, 0x2e, 0xa1, 0xb3, 0x74, 0xc0, 0xf9, 0xe9, 0xdf, 0xf6, 0xaf, 0xad,
	0x6f, 0x7f, 0xee, 0x8f, 0xe8, 0x77, 0x38, 0xae, 0xde, 0xd3, 0x92, 0xbc, 0x81, 0x11, 0x4b, 0x14,
	0x32, 0xfe, 0x2b, 0xc6, 0x17, 0xa1, 0x8d, 0x76, 0x02, 0x41, 0x34, 0xf4, 0xec, 0x67, 0x47, 0x96,
	0xf4, 0x9b, 0xff, 0xd6, 0xbf, 0x80, 0xa3, 0x3b, 0x4c, 0xb0, 0xfc, 0xae, 0x5a, 0xc6, 0xf4, 0x06,
	0x8e, 0xab, 0x25, 0x5a, 0x92, 0xd7, 0xd0, 0x4b, 0x33, 0x13, 0xff, 0xc8, 0xf2, 0x94, 0xfb, 0xee,
	0x41, 0x9a, 0x99, 0x2f, 0x16, 0x53, 0x01, 0xc1, 0x3d, 0xd3, 0x7a, 0x95, 0x29, 0x4e, 0x4e, 0xe0,
	0x10, 0x9f, 0x99, 0x48, 0xbc, 0x5e, 0x01, 0x6c, 0x78, 0x4f, 0x4c, 0x3f, 0xb9, 0x87, 0x0d, 0x22,
	0xf7, 0x4d, 0x26, 0x10, 0xe4, 0x1a, 0x95, 0x0b, 0xb5, 0xe5, 0x8a, 0x37, 0x98, 0x9c, 0x41, 0xd7,
	0x7e, 0xc7, 0x82, 0x87, 0xed, 0x62, 0xce, 0x16, 0x7e, 0xe3, 0xf4, 0x13, 0x8c, 0x8b, 0x78, 0xd6,
	0x0d, 0xad, 0x81, 0x2b, 0x08, 0xa4, 0x87, 0x3e, 0xda, 0xa1, 0xb3, 0xbe, 0xa9, 0xd9, 0x1c, 0xd3,
	0x0f, 0x40, 0xea, 0xf7, 0xff, 0x3b, 0x60, 0xfa, 0x08, 0xe3, 0x85, 0xe4, 0xb5, 0xe6, 0xbb, 0x0d,
	0x9f, 0x43, 0x90, 0xe2, 0x2a, 0x2e, 0x99, 0xee, 0xa6, 0xb8, 0xfa, 0x6a, 0x7d, 0x5f, 0xc0, 0xc0,
	0x1e, 0xd5, 0xbc, 0xf7, 0x53, 0x5c, 0x2d, 0x3c, 0x45, 0xdf, 0x02, 0xa9, 0x37, 0xda, 0x37, 0x83,
	0x2b, 0x18, 0x17, 0x43, 0xdb, 0xfb, 0x36, 0xab, 0x5e, 0x2f, 0xdd, 0xa3, 0x7e, 0xfb, 0xa7, 0x09,
	0xad, 0x3b, 0x7c, 0x21, 0x1f, 0x61, 0x50, 0xde, 0x4e, 0x72, 0x52, 0xac, 0x58, 0x75, 0xd1, 0x27,
	0xaf, 0x76, 0xb0, 0x5a, 0xd2, 0x03, 0x7b, 0xbd, 0xbc, 0x59, 0xfe, 0x7a, 0x6d, 0x1f, 0xfd, 0xf5,
	0xfa, 0x0a, 0xd2, 0x03, 0x32, 0x87, 0x51, 0x75, 0x78, 0xe4, 0xb4, 0xd4, 0xa9, 0x64, 0x7c, 0x72,
	0xb6, 0x93, 0x5f, 0x8b, 0x54, 0xb3, 0xf5, 0x22, 0x5b, 0x93, 0xf5, 0x22, 0xdb, 0x83, 0x28, 0x44,
	0xaa, 0x11, 0x7a, 0x91, 0xad, 0x11, 0x78, 0x91, 0xed, 0xbc, 0xe9, 0xc1, 0x43, 0xc7, 0xfd, 0xf2,
	0xde, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x63, 0xfb, 0xcd, 0x93, 0x03, 0x05, 0x00, 0x00,
}