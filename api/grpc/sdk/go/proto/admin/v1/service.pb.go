// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: proto/admin/v1/service.proto

package adminv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Admin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username  string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email     string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	FullName  string   `protobuf:"bytes,4,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Roles     []string `protobuf:"bytes,5,rep,name=roles,proto3" json:"roles,omitempty"`
	IsActive  bool     `protobuf:"varint,6,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	CreatedAt int64    `protobuf:"varint,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt int64    `protobuf:"varint,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Admin) Reset() {
	*x = Admin{}
	mi := &file_proto_admin_v1_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Admin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Admin) ProtoMessage() {}

func (x *Admin) ProtoReflect() protoreflect.Message {
	mi := &file_proto_admin_v1_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Admin.ProtoReflect.Descriptor instead.
func (*Admin) Descriptor() ([]byte, []int) {
	return file_proto_admin_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *Admin) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Admin) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Admin) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Admin) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Admin) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *Admin) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Admin) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Admin) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type CreateAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email    string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	FullName string   `protobuf:"bytes,4,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Roles    []string `protobuf:"bytes,5,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *CreateAdminRequest) Reset() {
	*x = CreateAdminRequest{}
	mi := &file_proto_admin_v1_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAdminRequest) ProtoMessage() {}

func (x *CreateAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_admin_v1_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAdminRequest.ProtoReflect.Descriptor instead.
func (*CreateAdminRequest) Descriptor() ([]byte, []int) {
	return file_proto_admin_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAdminRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateAdminRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateAdminRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateAdminRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *CreateAdminRequest) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

type CreateAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Admin *Admin `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
}

func (x *CreateAdminResponse) Reset() {
	*x = CreateAdminResponse{}
	mi := &file_proto_admin_v1_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAdminResponse) ProtoMessage() {}

func (x *CreateAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_admin_v1_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAdminResponse.ProtoReflect.Descriptor instead.
func (*CreateAdminResponse) Descriptor() ([]byte, []int) {
	return file_proto_admin_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAdminResponse) GetAdmin() *Admin {
	if x != nil {
		return x.Admin
	}
	return nil
}

type GetAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAdminRequest) Reset() {
	*x = GetAdminRequest{}
	mi := &file_proto_admin_v1_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdminRequest) ProtoMessage() {}

func (x *GetAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_admin_v1_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdminRequest.ProtoReflect.Descriptor instead.
func (*GetAdminRequest) Descriptor() ([]byte, []int) {
	return file_proto_admin_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetAdminRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Admin *Admin `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
}

func (x *GetAdminResponse) Reset() {
	*x = GetAdminResponse{}
	mi := &file_proto_admin_v1_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdminResponse) ProtoMessage() {}

func (x *GetAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_admin_v1_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdminResponse.ProtoReflect.Descriptor instead.
func (*GetAdminResponse) Descriptor() ([]byte, []int) {
	return file_proto_admin_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetAdminResponse) GetAdmin() *Admin {
	if x != nil {
		return x.Admin
	}
	return nil
}

type UpdateAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email    string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	FullName string   `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Roles    []string `protobuf:"bytes,4,rep,name=roles,proto3" json:"roles,omitempty"`
	IsActive bool     `protobuf:"varint,5,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
}

func (x *UpdateAdminRequest) Reset() {
	*x = UpdateAdminRequest{}
	mi := &file_proto_admin_v1_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAdminRequest) ProtoMessage() {}

func (x *UpdateAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_admin_v1_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAdminRequest.ProtoReflect.Descriptor instead.
func (*UpdateAdminRequest) Descriptor() ([]byte, []int) {
	return file_proto_admin_v1_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateAdminRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateAdminRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateAdminRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *UpdateAdminRequest) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *UpdateAdminRequest) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

type UpdateAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Admin *Admin `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
}

func (x *UpdateAdminResponse) Reset() {
	*x = UpdateAdminResponse{}
	mi := &file_proto_admin_v1_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAdminResponse) ProtoMessage() {}

func (x *UpdateAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_admin_v1_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAdminResponse.ProtoReflect.Descriptor instead.
func (*UpdateAdminResponse) Descriptor() ([]byte, []int) {
	return file_proto_admin_v1_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateAdminResponse) GetAdmin() *Admin {
	if x != nil {
		return x.Admin
	}
	return nil
}

type DeleteAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAdminRequest) Reset() {
	*x = DeleteAdminRequest{}
	mi := &file_proto_admin_v1_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAdminRequest) ProtoMessage() {}

func (x *DeleteAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_admin_v1_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAdminRequest.ProtoReflect.Descriptor instead.
func (*DeleteAdminRequest) Descriptor() ([]byte, []int) {
	return file_proto_admin_v1_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteAdminRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteAdminResponse) Reset() {
	*x = DeleteAdminResponse{}
	mi := &file_proto_admin_v1_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAdminResponse) ProtoMessage() {}

func (x *DeleteAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_admin_v1_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAdminResponse.ProtoReflect.Descriptor instead.
func (*DeleteAdminResponse) Descriptor() ([]byte, []int) {
	return file_proto_admin_v1_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteAdminResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ListAdminsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize  int32  `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListAdminsRequest) Reset() {
	*x = ListAdminsRequest{}
	mi := &file_proto_admin_v1_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAdminsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAdminsRequest) ProtoMessage() {}

func (x *ListAdminsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_admin_v1_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAdminsRequest.ProtoReflect.Descriptor instead.
func (*ListAdminsRequest) Descriptor() ([]byte, []int) {
	return file_proto_admin_v1_service_proto_rawDescGZIP(), []int{9}
}

func (x *ListAdminsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListAdminsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListAdminsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Admins        []*Admin `protobuf:"bytes,1,rep,name=admins,proto3" json:"admins,omitempty"`
	NextPageToken string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	TotalCount    int32    `protobuf:"varint,3,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *ListAdminsResponse) Reset() {
	*x = ListAdminsResponse{}
	mi := &file_proto_admin_v1_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAdminsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAdminsResponse) ProtoMessage() {}

func (x *ListAdminsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_admin_v1_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAdminsResponse.ProtoReflect.Descriptor instead.
func (*ListAdminsResponse) Descriptor() ([]byte, []int) {
	return file_proto_admin_v1_service_proto_rawDescGZIP(), []int{10}
}

func (x *ListAdminsResponse) GetAdmins() []*Admin {
	if x != nil {
		return x.Admins
	}
	return nil
}

func (x *ListAdminsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListAdminsResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_proto_admin_v1_service_proto protoreflect.FileDescriptor

var file_proto_admin_v1_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0xd7, 0x01, 0x0a, 0x05, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x95, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x3c, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x25, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x39, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x25, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52,
	0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x8a, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x22, 0x3c, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2f, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x4f, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x86, 0x01, 0x0a, 0x12, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x27, 0x0a, 0x06, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x52, 0x06, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78,
	0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x32, 0x88, 0x03, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x12, 0x1c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x43, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x19, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x12, 0x1c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x49, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x73,
	0x12, 0x1b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2c, 0x5a,
	0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x72, 0x6c, 0x34,
	0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2f, 0x76, 0x31, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_admin_v1_service_proto_rawDescOnce sync.Once
	file_proto_admin_v1_service_proto_rawDescData = file_proto_admin_v1_service_proto_rawDesc
)

func file_proto_admin_v1_service_proto_rawDescGZIP() []byte {
	file_proto_admin_v1_service_proto_rawDescOnce.Do(func() {
		file_proto_admin_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_admin_v1_service_proto_rawDescData)
	})
	return file_proto_admin_v1_service_proto_rawDescData
}

var file_proto_admin_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_admin_v1_service_proto_goTypes = []any{
	(*Admin)(nil),               // 0: admin.v1.Admin
	(*CreateAdminRequest)(nil),  // 1: admin.v1.CreateAdminRequest
	(*CreateAdminResponse)(nil), // 2: admin.v1.CreateAdminResponse
	(*GetAdminRequest)(nil),     // 3: admin.v1.GetAdminRequest
	(*GetAdminResponse)(nil),    // 4: admin.v1.GetAdminResponse
	(*UpdateAdminRequest)(nil),  // 5: admin.v1.UpdateAdminRequest
	(*UpdateAdminResponse)(nil), // 6: admin.v1.UpdateAdminResponse
	(*DeleteAdminRequest)(nil),  // 7: admin.v1.DeleteAdminRequest
	(*DeleteAdminResponse)(nil), // 8: admin.v1.DeleteAdminResponse
	(*ListAdminsRequest)(nil),   // 9: admin.v1.ListAdminsRequest
	(*ListAdminsResponse)(nil),  // 10: admin.v1.ListAdminsResponse
}
var file_proto_admin_v1_service_proto_depIdxs = []int32{
	0,  // 0: admin.v1.CreateAdminResponse.admin:type_name -> admin.v1.Admin
	0,  // 1: admin.v1.GetAdminResponse.admin:type_name -> admin.v1.Admin
	0,  // 2: admin.v1.UpdateAdminResponse.admin:type_name -> admin.v1.Admin
	0,  // 3: admin.v1.ListAdminsResponse.admins:type_name -> admin.v1.Admin
	1,  // 4: admin.v1.AdminService.CreateAdmin:input_type -> admin.v1.CreateAdminRequest
	3,  // 5: admin.v1.AdminService.GetAdmin:input_type -> admin.v1.GetAdminRequest
	5,  // 6: admin.v1.AdminService.UpdateAdmin:input_type -> admin.v1.UpdateAdminRequest
	7,  // 7: admin.v1.AdminService.DeleteAdmin:input_type -> admin.v1.DeleteAdminRequest
	9,  // 8: admin.v1.AdminService.ListAdmins:input_type -> admin.v1.ListAdminsRequest
	2,  // 9: admin.v1.AdminService.CreateAdmin:output_type -> admin.v1.CreateAdminResponse
	4,  // 10: admin.v1.AdminService.GetAdmin:output_type -> admin.v1.GetAdminResponse
	6,  // 11: admin.v1.AdminService.UpdateAdmin:output_type -> admin.v1.UpdateAdminResponse
	8,  // 12: admin.v1.AdminService.DeleteAdmin:output_type -> admin.v1.DeleteAdminResponse
	10, // 13: admin.v1.AdminService.ListAdmins:output_type -> admin.v1.ListAdminsResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_admin_v1_service_proto_init() }
func file_proto_admin_v1_service_proto_init() {
	if File_proto_admin_v1_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_admin_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_admin_v1_service_proto_goTypes,
		DependencyIndexes: file_proto_admin_v1_service_proto_depIdxs,
		MessageInfos:      file_proto_admin_v1_service_proto_msgTypes,
	}.Build()
	File_proto_admin_v1_service_proto = out.File
	file_proto_admin_v1_service_proto_rawDesc = nil
	file_proto_admin_v1_service_proto_goTypes = nil
	file_proto_admin_v1_service_proto_depIdxs = nil
}