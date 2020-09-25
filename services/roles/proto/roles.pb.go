// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: services/roles/proto/roles.proto

package roles

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ListRolesRequest struct {
	PageSize             int32    `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken            string   `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRolesRequest) Reset()         { *m = ListRolesRequest{} }
func (m *ListRolesRequest) String() string { return proto.CompactTextString(m) }
func (*ListRolesRequest) ProtoMessage()    {}
func (*ListRolesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a385908be3e62237, []int{0}
}
func (m *ListRolesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRolesRequest.Unmarshal(m, b)
}
func (m *ListRolesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRolesRequest.Marshal(b, m, deterministic)
}
func (m *ListRolesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRolesRequest.Merge(m, src)
}
func (m *ListRolesRequest) XXX_Size() int {
	return xxx_messageInfo_ListRolesRequest.Size(m)
}
func (m *ListRolesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRolesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRolesRequest proto.InternalMessageInfo

func (m *ListRolesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListRolesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListRolesResponse struct {
	Roles                []*Role  `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
	TotalSize            int32    `protobuf:"varint,2,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	NextPageToken        string   `protobuf:"bytes,3,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRolesResponse) Reset()         { *m = ListRolesResponse{} }
func (m *ListRolesResponse) String() string { return proto.CompactTextString(m) }
func (*ListRolesResponse) ProtoMessage()    {}
func (*ListRolesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a385908be3e62237, []int{1}
}
func (m *ListRolesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRolesResponse.Unmarshal(m, b)
}
func (m *ListRolesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRolesResponse.Marshal(b, m, deterministic)
}
func (m *ListRolesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRolesResponse.Merge(m, src)
}
func (m *ListRolesResponse) XXX_Size() int {
	return xxx_messageInfo_ListRolesResponse.Size(m)
}
func (m *ListRolesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRolesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListRolesResponse proto.InternalMessageInfo

func (m *ListRolesResponse) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *ListRolesResponse) GetTotalSize() int32 {
	if m != nil {
		return m.TotalSize
	}
	return 0
}

func (m *ListRolesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type GetRoleRequest struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRoleRequest) Reset()         { *m = GetRoleRequest{} }
func (m *GetRoleRequest) String() string { return proto.CompactTextString(m) }
func (*GetRoleRequest) ProtoMessage()    {}
func (*GetRoleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a385908be3e62237, []int{2}
}
func (m *GetRoleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRoleRequest.Unmarshal(m, b)
}
func (m *GetRoleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRoleRequest.Marshal(b, m, deterministic)
}
func (m *GetRoleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRoleRequest.Merge(m, src)
}
func (m *GetRoleRequest) XXX_Size() int {
	return xxx_messageInfo_GetRoleRequest.Size(m)
}
func (m *GetRoleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRoleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRoleRequest proto.InternalMessageInfo

func (m *GetRoleRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type CreateRoleRequest struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRoleRequest) Reset()         { *m = CreateRoleRequest{} }
func (m *CreateRoleRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRoleRequest) ProtoMessage()    {}
func (*CreateRoleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a385908be3e62237, []int{3}
}
func (m *CreateRoleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoleRequest.Unmarshal(m, b)
}
func (m *CreateRoleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoleRequest.Marshal(b, m, deterministic)
}
func (m *CreateRoleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoleRequest.Merge(m, src)
}
func (m *CreateRoleRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRoleRequest.Size(m)
}
func (m *CreateRoleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoleRequest proto.InternalMessageInfo

func (m *CreateRoleRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type UpdateRoleRequest struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRoleRequest) Reset()         { *m = UpdateRoleRequest{} }
func (m *UpdateRoleRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRoleRequest) ProtoMessage()    {}
func (*UpdateRoleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a385908be3e62237, []int{4}
}
func (m *UpdateRoleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRoleRequest.Unmarshal(m, b)
}
func (m *UpdateRoleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRoleRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRoleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRoleRequest.Merge(m, src)
}
func (m *UpdateRoleRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRoleRequest.Size(m)
}
func (m *UpdateRoleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRoleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRoleRequest proto.InternalMessageInfo

func (m *UpdateRoleRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *UpdateRoleRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type DeleteRoleRequest struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRoleRequest) Reset()         { *m = DeleteRoleRequest{} }
func (m *DeleteRoleRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRoleRequest) ProtoMessage()    {}
func (*DeleteRoleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a385908be3e62237, []int{5}
}
func (m *DeleteRoleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRoleRequest.Unmarshal(m, b)
}
func (m *DeleteRoleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRoleRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRoleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRoleRequest.Merge(m, src)
}
func (m *DeleteRoleRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRoleRequest.Size(m)
}
func (m *DeleteRoleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRoleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRoleRequest proto.InternalMessageInfo

func (m *DeleteRoleRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type Role struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_a385908be3e62237, []int{6}
}
func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (m *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(m, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Role) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func init() {
	proto.RegisterType((*ListRolesRequest)(nil), "roles.ListRolesRequest")
	proto.RegisterType((*ListRolesResponse)(nil), "roles.ListRolesResponse")
	proto.RegisterType((*GetRoleRequest)(nil), "roles.GetRoleRequest")
	proto.RegisterType((*CreateRoleRequest)(nil), "roles.CreateRoleRequest")
	proto.RegisterType((*UpdateRoleRequest)(nil), "roles.UpdateRoleRequest")
	proto.RegisterType((*DeleteRoleRequest)(nil), "roles.DeleteRoleRequest")
	proto.RegisterType((*Role)(nil), "roles.Role")
}

func init() { proto.RegisterFile("services/roles/proto/roles.proto", fileDescriptor_a385908be3e62237) }

var fileDescriptor_a385908be3e62237 = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0xc7, 0xc9, 0x6e, 0xd7, 0x9a, 0x13, 0xac, 0xcd, 0xd0, 0x6a, 0xc8, 0x2a, 0xc4, 0x41, 0x74,
	0xed, 0x45, 0xa2, 0xf5, 0x4e, 0xf0, 0xc6, 0x0f, 0x8a, 0x20, 0x52, 0x53, 0x05, 0xf1, 0xa6, 0xa4,
	0xf6, 0xb8, 0x0c, 0xc6, 0x4c, 0xcc, 0x4c, 0x8a, 0x56, 0x7a, 0xe3, 0x2b, 0xf8, 0x02, 0xbe, 0x93,
	0xaf, 0xe0, 0x83, 0xc8, 0x9c, 0xc9, 0xd7, 0x36, 0x82, 0xbd, 0x9b, 0x39, 0x1f, 0xbf, 0xf3, 0xf1,
	0x9f, 0x81, 0x48, 0x61, 0x75, 0x22, 0x3e, 0xa0, 0x4a, 0x2a, 0x99, 0xa3, 0x4a, 0xca, 0x4a, 0x6a,
	0x69, 0xcf, 0x31, 0x9d, 0xd9, 0x8c, 0x2e, 0xe1, 0x8d, 0xa5, 0x94, 0xcb, 0x1c, 0x93, 0xac, 0x14,
	0x49, 0x56, 0x14, 0x52, 0x67, 0x5a, 0xc8, 0xa2, 0x09, 0x0a, 0xe7, 0x8d, 0x97, 0x6e, 0x47, 0xf5,
	0xc7, 0x04, 0x3f, 0x97, 0xfa, 0x9b, 0x75, 0xf2, 0x57, 0xb0, 0xf9, 0x52, 0x28, 0x9d, 0x1a, 0x4e,
	0x8a, 0x5f, 0x6a, 0x54, 0x9a, 0xcd, 0xc1, 0x2d, 0xb3, 0x25, 0x1e, 0x2a, 0x71, 0x8a, 0x81, 0x13,
	0x39, 0x8b, 0x59, 0x7a, 0xd9, 0x18, 0x0e, 0xc4, 0x29, 0xb2, 0x9b, 0x00, 0xe4, 0xd4, 0xf2, 0x13,
	0x16, 0xc1, 0x24, 0x72, 0x16, 0x6e, 0x4a, 0xe1, 0x6f, 0x8c, 0x81, 0x9f, 0x81, 0x3f, 0xe0, 0xa9,
	0x52, 0x16, 0x0a, 0xd9, 0x2d, 0xb0, 0x8d, 0x06, 0x4e, 0x34, 0x5d, 0x78, 0xbb, 0x5e, 0x6c, 0x67,
	0x30, 0x41, 0xa9, 0xf5, 0x18, 0xac, 0x96, 0x3a, 0xcb, 0x6d, 0xd1, 0x09, 0x15, 0x75, 0xc9, 0x42,
	0x55, 0xef, 0xc0, 0xd5, 0x02, 0xbf, 0xea, 0xc3, 0x41, 0xe9, 0x29, 0x95, 0xbe, 0x62, 0xcc, 0xfb,
	0x5d, 0xf9, 0xdb, 0xb0, 0xb1, 0x87, 0x54, 0xbd, 0x1d, 0x86, 0xc1, 0x5a, 0x5d, 0x8b, 0x63, 0x9a,
	0xc3, 0x4d, 0xe9, 0xcc, 0xef, 0x81, 0xff, 0xb4, 0xc2, 0x4c, 0xe3, 0x30, 0x70, 0x0b, 0x66, 0x5a,
	0xe8, 0x1c, 0x9b, 0x48, 0x7b, 0xe1, 0x8f, 0xc1, 0x7f, 0x5b, 0x1e, 0x9f, 0x0b, 0xfd, 0x07, 0xb3,
	0x4f, 0x9f, 0x0c, 0xd3, 0xef, 0x82, 0xff, 0x0c, 0x73, 0xfc, 0x6f, 0x3a, 0xbf, 0x0f, 0x6b, 0x26,
	0xe4, 0xe2, 0xe8, 0xdd, 0x5f, 0x53, 0xf0, 0x4c, 0xca, 0x81, 0x7d, 0x24, 0xec, 0x35, 0xb8, 0xdd,
	0xe6, 0xd9, 0xf5, 0x66, 0xc5, 0xe7, 0xb5, 0x0d, 0x83, 0xb1, 0xc3, 0x8a, 0xc4, 0xfd, 0x1f, 0xbf,
	0xff, 0xfc, 0x9c, 0x78, 0xcc, 0x4d, 0x4e, 0x1e, 0xd8, 0x47, 0xc6, 0xf6, 0x60, 0xbd, 0xd9, 0x26,
	0xdb, 0x6e, 0xf2, 0x56, 0xb7, 0x1b, 0x0e, 0xa5, 0xe4, 0x01, 0x11, 0x18, 0xdb, 0xec, 0x08, 0xc9,
	0x77, 0x33, 0xc0, 0x19, 0x7b, 0x01, 0xd0, 0x2f, 0x9c, 0xb5, 0x3d, 0x8c, 0x34, 0x58, 0xc5, 0x6d,
	0x11, 0x6e, 0x83, 0xf7, 0x0d, 0x3d, 0x72, 0x76, 0xd8, 0x3e, 0x40, 0x2f, 0x48, 0x87, 0x1a, 0x69,
	0xb4, 0x8a, 0x9a, 0x13, 0x6a, 0x3b, 0x1c, 0x75, 0x66, 0x88, 0xef, 0x00, 0x7a, 0x8d, 0x3a, 0xe2,
	0x48, 0xb6, 0xf0, 0x5a, 0x6c, 0x3f, 0x52, 0xdc, 0x7e, 0xa4, 0xf8, 0xb9, 0xf9, 0x48, 0xed, 0xd8,
	0x3b, 0x23, 0xf8, 0x93, 0xf5, 0xf7, 0xf6, 0x75, 0x1f, 0x5d, 0xa2, 0x94, 0x87, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x14, 0xe8, 0x37, 0x82, 0xd2, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RoleServiceClient is the client API for RoleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoleServiceClient interface {
	// List Roles
	ListRoles(ctx context.Context, in *ListRolesRequest, opts ...grpc.CallOption) (*ListRolesResponse, error)
	// Get Role
	GetRole(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*Role, error)
	// Create Role object request
	CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*Role, error)
	// Update Role object request
	UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*Role, error)
	// Delete Role object request
	DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*types.Empty, error)
}

type roleServiceClient struct {
	cc *grpc.ClientConn
}

func NewRoleServiceClient(cc *grpc.ClientConn) RoleServiceClient {
	return &roleServiceClient{cc}
}

func (c *roleServiceClient) ListRoles(ctx context.Context, in *ListRolesRequest, opts ...grpc.CallOption) (*ListRolesResponse, error) {
	out := new(ListRolesResponse)
	err := c.cc.Invoke(ctx, "/roles.RoleService/ListRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) GetRole(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/roles.RoleService/GetRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/roles.RoleService/CreateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/roles.RoleService/UpdateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/roles.RoleService/DeleteRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleServiceServer is the server API for RoleService service.
type RoleServiceServer interface {
	// List Roles
	ListRoles(context.Context, *ListRolesRequest) (*ListRolesResponse, error)
	// Get Role
	GetRole(context.Context, *GetRoleRequest) (*Role, error)
	// Create Role object request
	CreateRole(context.Context, *CreateRoleRequest) (*Role, error)
	// Update Role object request
	UpdateRole(context.Context, *UpdateRoleRequest) (*Role, error)
	// Delete Role object request
	DeleteRole(context.Context, *DeleteRoleRequest) (*types.Empty, error)
}

// UnimplementedRoleServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRoleServiceServer struct {
}

func (*UnimplementedRoleServiceServer) ListRoles(ctx context.Context, req *ListRolesRequest) (*ListRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoles not implemented")
}
func (*UnimplementedRoleServiceServer) GetRole(ctx context.Context, req *GetRoleRequest) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRole not implemented")
}
func (*UnimplementedRoleServiceServer) CreateRole(ctx context.Context, req *CreateRoleRequest) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (*UnimplementedRoleServiceServer) UpdateRole(ctx context.Context, req *UpdateRoleRequest) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (*UnimplementedRoleServiceServer) DeleteRole(ctx context.Context, req *DeleteRoleRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}

func RegisterRoleServiceServer(s *grpc.Server, srv RoleServiceServer) {
	s.RegisterService(&_RoleService_serviceDesc, srv)
}

func _RoleService_ListRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).ListRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roles.RoleService/ListRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).ListRoles(ctx, req.(*ListRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_GetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).GetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roles.RoleService/GetRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).GetRole(ctx, req.(*GetRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roles.RoleService/CreateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).CreateRole(ctx, req.(*CreateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roles.RoleService/UpdateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).UpdateRole(ctx, req.(*UpdateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roles.RoleService/DeleteRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).DeleteRole(ctx, req.(*DeleteRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "roles.RoleService",
	HandlerType: (*RoleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRoles",
			Handler:    _RoleService_ListRoles_Handler,
		},
		{
			MethodName: "GetRole",
			Handler:    _RoleService_GetRole_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _RoleService_CreateRole_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _RoleService_UpdateRole_Handler,
		},
		{
			MethodName: "DeleteRole",
			Handler:    _RoleService_DeleteRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/roles/proto/roles.proto",
}