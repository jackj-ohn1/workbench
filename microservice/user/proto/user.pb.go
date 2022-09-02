// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type UpdateTeamGroupRequest struct {
	Ids                  []uint32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Value                uint32   `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	Kind                 uint32   `protobuf:"varint,3,opt,name=kind,proto3" json:"kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTeamGroupRequest) Reset()         { *m = UpdateTeamGroupRequest{} }
func (m *UpdateTeamGroupRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTeamGroupRequest) ProtoMessage()    {}
func (*UpdateTeamGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{0}
}

func (m *UpdateTeamGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTeamGroupRequest.Unmarshal(m, b)
}
func (m *UpdateTeamGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTeamGroupRequest.Marshal(b, m, deterministic)
}
func (m *UpdateTeamGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTeamGroupRequest.Merge(m, src)
}
func (m *UpdateTeamGroupRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTeamGroupRequest.Size(m)
}
func (m *UpdateTeamGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTeamGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTeamGroupRequest proto.InternalMessageInfo

func (m *UpdateTeamGroupRequest) GetIds() []uint32 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *UpdateTeamGroupRequest) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *UpdateTeamGroupRequest) GetKind() uint32 {
	if m != nil {
		return m.Kind
	}
	return 0
}

type UpdateInfoRequest struct {
	Id                   uint32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Info                 *UserInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UpdateInfoRequest) Reset()         { *m = UpdateInfoRequest{} }
func (m *UpdateInfoRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateInfoRequest) ProtoMessage()    {}
func (*UpdateInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{1}
}

func (m *UpdateInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateInfoRequest.Unmarshal(m, b)
}
func (m *UpdateInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateInfoRequest.Marshal(b, m, deterministic)
}
func (m *UpdateInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateInfoRequest.Merge(m, src)
}
func (m *UpdateInfoRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateInfoRequest.Size(m)
}
func (m *UpdateInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateInfoRequest proto.InternalMessageInfo

func (m *UpdateInfoRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateInfoRequest) GetInfo() *UserInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type GetRequest struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{2}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetInfoRequest struct {
	Ids                  []uint32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInfoRequest) Reset()         { *m = GetInfoRequest{} }
func (m *GetInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetInfoRequest) ProtoMessage()    {}
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{3}
}

func (m *GetInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInfoRequest.Unmarshal(m, b)
}
func (m *GetInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoRequest.Merge(m, src)
}
func (m *GetInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetInfoRequest.Size(m)
}
func (m *GetInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoRequest proto.InternalMessageInfo

func (m *GetInfoRequest) GetIds() []uint32 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type RegisterRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{4}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginRequest struct {
	OauthCode            string   `protobuf:"bytes,1,opt,name=oauth_code,json=oauthCode,proto3" json:"oauth_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{5}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetOauthCode() string {
	if m != nil {
		return m.OauthCode
	}
	return ""
}

type LoginResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	RedirectUrl          string   `protobuf:"bytes,2,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{6}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginResponse) GetRedirectUrl() string {
	if m != nil {
		return m.RedirectUrl
	}
	return ""
}

type UserInfo struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	RealName             string   `protobuf:"bytes,3,opt,name=real_name,json=realName,proto3" json:"real_name,omitempty"`
	AvatarUrl            string   `protobuf:"bytes,4,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	Email                string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Tel                  string   `protobuf:"bytes,6,opt,name=tel,proto3" json:"tel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{7}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfo) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

func (m *UserInfo) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func (m *UserInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserInfo) GetTel() string {
	if m != nil {
		return m.Tel
	}
	return ""
}

type UserInfoResponse struct {
	List                 []*UserInfo `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UserInfoResponse) Reset()         { *m = UserInfoResponse{} }
func (m *UserInfoResponse) String() string { return proto.CompactTextString(m) }
func (*UserInfoResponse) ProtoMessage()    {}
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{8}
}

func (m *UserInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoResponse.Unmarshal(m, b)
}
func (m *UserInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoResponse.Marshal(b, m, deterministic)
}
func (m *UserInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoResponse.Merge(m, src)
}
func (m *UserInfoResponse) XXX_Size() int {
	return xxx_messageInfo_UserInfoResponse.Size(m)
}
func (m *UserInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoResponse proto.InternalMessageInfo

func (m *UserInfoResponse) GetList() []*UserInfo {
	if m != nil {
		return m.List
	}
	return nil
}

type UserProfile struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	RealName             string   `protobuf:"bytes,3,opt,name=real_name,json=realName,proto3" json:"real_name,omitempty"`
	Avatar               string   `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Email                string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Tel                  string   `protobuf:"bytes,6,opt,name=tel,proto3" json:"tel,omitempty"`
	Role                 uint32   `protobuf:"varint,7,opt,name=role,proto3" json:"role,omitempty"`
	Team                 uint32   `protobuf:"varint,8,opt,name=team,proto3" json:"team,omitempty"`
	Group                uint32   `protobuf:"varint,9,opt,name=group,proto3" json:"group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserProfile) Reset()         { *m = UserProfile{} }
func (m *UserProfile) String() string { return proto.CompactTextString(m) }
func (*UserProfile) ProtoMessage()    {}
func (*UserProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{9}
}

func (m *UserProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserProfile.Unmarshal(m, b)
}
func (m *UserProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserProfile.Marshal(b, m, deterministic)
}
func (m *UserProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserProfile.Merge(m, src)
}
func (m *UserProfile) XXX_Size() int {
	return xxx_messageInfo_UserProfile.Size(m)
}
func (m *UserProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_UserProfile.DiscardUnknown(m)
}

var xxx_messageInfo_UserProfile proto.InternalMessageInfo

func (m *UserProfile) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserProfile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserProfile) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

func (m *UserProfile) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *UserProfile) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserProfile) GetTel() string {
	if m != nil {
		return m.Tel
	}
	return ""
}

func (m *UserProfile) GetRole() uint32 {
	if m != nil {
		return m.Role
	}
	return 0
}

func (m *UserProfile) GetTeam() uint32 {
	if m != nil {
		return m.Team
	}
	return 0
}

func (m *UserProfile) GetGroup() uint32 {
	if m != nil {
		return m.Group
	}
	return 0
}

type GetResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{10}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type ListRequest struct {
	LastId               uint32   `protobuf:"varint,1,opt,name=last_id,json=lastId,proto3" json:"last_id,omitempty"`
	Offset               uint32   `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                uint32   `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Team                 uint32   `protobuf:"varint,4,opt,name=team,proto3" json:"team,omitempty"`
	Group                uint32   `protobuf:"varint,5,opt,name=group,proto3" json:"group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{11}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetLastId() uint32 {
	if m != nil {
		return m.LastId
	}
	return 0
}

func (m *ListRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListRequest) GetTeam() uint32 {
	if m != nil {
		return m.Team
	}
	return 0
}

func (m *ListRequest) GetGroup() uint32 {
	if m != nil {
		return m.Group
	}
	return 0
}

type ListResponse struct {
	Count                uint32   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	List                 []*User  `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{12}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ListResponse) GetList() []*User {
	if m != nil {
		return m.List
	}
	return nil
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{13}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Response struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{14}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

type User struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	RealName             string   `protobuf:"bytes,3,opt,name=real_name,json=realName,proto3" json:"real_name,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Avatar               string   `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Role                 uint32   `protobuf:"varint,6,opt,name=role,proto3" json:"role,omitempty"`
	Team                 uint32   `protobuf:"varint,7,opt,name=team,proto3" json:"team,omitempty"`
	Group                uint32   `protobuf:"varint,8,opt,name=group,proto3" json:"group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{15}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *User) GetRole() uint32 {
	if m != nil {
		return m.Role
	}
	return 0
}

func (m *User) GetTeam() uint32 {
	if m != nil {
		return m.Team
	}
	return 0
}

func (m *User) GetGroup() uint32 {
	if m != nil {
		return m.Group
	}
	return 0
}

func init() {
	proto.RegisterType((*UpdateTeamGroupRequest)(nil), "user.UpdateTeamGroupRequest")
	proto.RegisterType((*UpdateInfoRequest)(nil), "user.UpdateInfoRequest")
	proto.RegisterType((*GetRequest)(nil), "user.GetRequest")
	proto.RegisterType((*GetInfoRequest)(nil), "user.GetInfoRequest")
	proto.RegisterType((*RegisterRequest)(nil), "user.RegisterRequest")
	proto.RegisterType((*LoginRequest)(nil), "user.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "user.LoginResponse")
	proto.RegisterType((*UserInfo)(nil), "user.UserInfo")
	proto.RegisterType((*UserInfoResponse)(nil), "user.UserInfoResponse")
	proto.RegisterType((*UserProfile)(nil), "user.UserProfile")
	proto.RegisterType((*GetResponse)(nil), "user.GetResponse")
	proto.RegisterType((*ListRequest)(nil), "user.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "user.ListResponse")
	proto.RegisterType((*Request)(nil), "user.Request")
	proto.RegisterType((*Response)(nil), "user.Response")
	proto.RegisterType((*User)(nil), "user.User")
}

func init() { proto.RegisterFile("proto/user.proto", fileDescriptor_d570e3e37e5899c5) }

var fileDescriptor_d570e3e37e5899c5 = []byte{
	// 704 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0x6e, 0x92, 0xcd, 0xdf, 0xa4, 0x0d, 0xa9, 0x29, 0xed, 0x2a, 0xb4, 0xa8, 0xf8, 0xd4, 0x4b,
	0x5b, 0x29, 0x95, 0x40, 0x1c, 0x11, 0x88, 0x50, 0xa9, 0x42, 0x68, 0x69, 0xc4, 0x31, 0x32, 0x59,
	0xa7, 0x58, 0xdd, 0xac, 0x83, 0xed, 0x2d, 0x37, 0x1e, 0x82, 0x07, 0xe0, 0x25, 0x78, 0x0d, 0x1e,
	0x0a, 0x79, 0xec, 0xfd, 0x49, 0x1a, 0x24, 0xa4, 0xde, 0x66, 0xc6, 0xf6, 0xf8, 0xfb, 0xe6, 0xf3,
	0x8c, 0x61, 0xb0, 0x54, 0xd2, 0xc8, 0xf3, 0x4c, 0x73, 0x75, 0x86, 0x26, 0x09, 0xac, 0x4d, 0xaf,
	0x61, 0x7f, 0xb2, 0x8c, 0x99, 0xe1, 0xd7, 0x9c, 0x2d, 0xc6, 0x4a, 0x66, 0xcb, 0x88, 0x7f, 0xcb,
	0xb8, 0x36, 0x64, 0x00, 0x0d, 0x11, 0xeb, 0xb0, 0x76, 0xdc, 0x38, 0xd9, 0x89, 0xac, 0x49, 0xf6,
	0xa0, 0x79, 0xc7, 0x92, 0x8c, 0x87, 0xf5, 0xe3, 0xda, 0xc9, 0x4e, 0xe4, 0x1c, 0x42, 0x20, 0xb8,
	0x15, 0x69, 0x1c, 0x36, 0x30, 0x88, 0x36, 0x1d, 0xc3, 0xae, 0xcb, 0x7a, 0x99, 0xce, 0x65, 0x9e,
	0xb0, 0x0f, 0x75, 0x11, 0x87, 0x35, 0xdc, 0x56, 0x17, 0x31, 0xa1, 0x10, 0x88, 0x74, 0x2e, 0x31,
	0x5b, 0x6f, 0xd4, 0x3f, 0x43, 0x6c, 0x13, 0xcd, 0x15, 0x1e, 0xc2, 0x35, 0x7a, 0x08, 0x30, 0xe6,
	0xe6, 0x1f, 0x19, 0x28, 0x85, 0xfe, 0x98, 0x9b, 0xea, 0x1d, 0xf7, 0x40, 0xd3, 0xcf, 0xf0, 0x28,
	0xe2, 0x37, 0x42, 0x1b, 0xae, 0xf2, 0x4d, 0x7b, 0xd0, 0xe4, 0x0b, 0x26, 0x12, 0xcc, 0xd4, 0x8d,
	0x9c, 0x63, 0x79, 0xa4, 0x6c, 0xe1, 0xc8, 0x75, 0x23, 0xb4, 0xc9, 0x10, 0x3a, 0x4b, 0xa6, 0xf5,
	0x77, 0xa9, 0x1c, 0xbf, 0x6e, 0x54, 0xf8, 0xf4, 0x14, 0xb6, 0xaf, 0xe4, 0x8d, 0x48, 0xf3, 0xac,
	0x47, 0x00, 0x92, 0x65, 0xe6, 0xeb, 0x74, 0x26, 0x63, 0xee, 0x53, 0x77, 0x31, 0xf2, 0x46, 0xc6,
	0x9c, 0xbe, 0x87, 0x1d, 0xbf, 0x5d, 0x2f, 0x65, 0xaa, 0xb9, 0x45, 0x61, 0xe4, 0x2d, 0x4f, 0x73,
	0x14, 0xe8, 0x90, 0xe7, 0xb0, 0xad, 0x78, 0x2c, 0x14, 0x9f, 0x99, 0x69, 0xa6, 0x12, 0x8f, 0xa6,
	0x97, 0xc7, 0x26, 0x2a, 0xa1, 0x3f, 0x6b, 0xd0, 0xc9, 0xcb, 0x74, 0xaf, 0xa8, 0x9b, 0x58, 0x3c,
	0x85, 0xae, 0xe2, 0x2c, 0x99, 0xe2, 0x82, 0xa7, 0x61, 0x03, 0x1f, 0xec, 0xe2, 0x11, 0x00, 0xbb,
	0x63, 0x86, 0x29, 0xbc, 0x2e, 0x70, 0xb0, 0x5d, 0x64, 0xa2, 0x92, 0xb2, 0x56, 0xcd, 0x6a, 0xad,
	0x06, 0xd0, 0x30, 0x3c, 0x09, 0x5b, 0x18, 0xb3, 0x26, 0x7d, 0x01, 0x83, 0x42, 0xba, 0x9c, 0x21,
	0x85, 0x20, 0x11, 0xda, 0xa0, 0x1a, 0x1b, 0x04, 0xb6, 0x6b, 0xf4, 0x4f, 0x0d, 0x7a, 0x36, 0xf4,
	0x51, 0xc9, 0xb9, 0x48, 0xf8, 0xc3, 0xf9, 0xec, 0x43, 0xcb, 0xa1, 0xf7, 0x5c, 0xbc, 0xf7, 0xbf,
	0x44, 0xec, 0x85, 0x4a, 0x26, 0x3c, 0x6c, 0xbb, 0xe7, 0x6c, 0x6d, 0x1b, 0x33, 0x9c, 0x2d, 0xc2,
	0x8e, 0x8b, 0x59, 0xdb, 0xe6, 0xbb, 0xb1, 0xed, 0x12, 0x76, 0x5d, 0x33, 0xa0, 0x43, 0x4f, 0xa1,
	0x87, 0xef, 0xd5, 0x57, 0xe0, 0x19, 0x60, 0x97, 0x21, 0x9f, 0xde, 0x08, 0xca, 0x0a, 0x44, 0xae,
	0xfb, 0x7e, 0x40, 0xef, 0x4a, 0xe8, 0xe2, 0x7d, 0x1f, 0x40, 0x3b, 0x61, 0xda, 0x4c, 0x8b, 0x0a,
	0xb4, 0xac, 0x7b, 0x19, 0x5b, 0x52, 0x72, 0x3e, 0xd7, 0xdc, 0xf8, 0xd6, 0xf3, 0x9e, 0x05, 0x91,
	0x88, 0x85, 0x30, 0xbe, 0xf9, 0x9c, 0x53, 0xc0, 0x0d, 0x36, 0xc1, 0x6d, 0x56, 0xe1, 0xbe, 0x85,
	0x6d, 0x77, 0x7f, 0xf9, 0x26, 0x67, 0x32, 0x4b, 0x8d, 0xbf, 0xde, 0x39, 0x96, 0x05, 0xea, 0x58,
	0x47, 0x1d, 0x57, 0x58, 0xa0, 0x86, 0x5d, 0x68, 0x7b, 0x06, 0x14, 0xa0, 0x93, 0x27, 0xa3, 0xbf,
	0x6b, 0x10, 0xd8, 0x5d, 0x0f, 0xd7, 0xb4, 0xd0, 0x2e, 0xa8, 0x6a, 0x57, 0x2a, 0xdd, 0x5c, 0x51,
	0x3a, 0x57, 0xb0, 0xb5, 0x41, 0xc1, 0xf6, 0xa6, 0x92, 0x74, 0x2a, 0x25, 0x19, 0xfd, 0x6a, 0xb8,
	0x07, 0xf9, 0x89, 0xab, 0x3b, 0x31, 0xe3, 0xe4, 0xc2, 0x32, 0x72, 0xf3, 0x83, 0x3c, 0x71, 0xd4,
	0xd7, 0xe6, 0xc9, 0xb0, 0x9f, 0x87, 0x3d, 0xf1, 0x2d, 0x32, 0x82, 0x26, 0x36, 0x3b, 0x21, 0x6e,
	0xa9, 0x3a, 0x28, 0x86, 0x8f, 0x57, 0x62, 0xc5, 0x99, 0x57, 0xd0, 0xf6, 0xc3, 0x8c, 0xec, 0xb9,
	0x1d, 0xab, 0xb3, 0x6d, 0xb8, 0xbf, 0xd6, 0x40, 0xe5, 0xd1, 0x0b, 0x9c, 0x92, 0x79, 0x0b, 0x0d,
	0x8a, 0xd3, 0xf9, 0xc9, 0xdd, 0xf2, 0xa4, 0xdf, 0x44, 0xb7, 0xc8, 0x39, 0x04, 0x56, 0x7b, 0xe2,
	0x17, 0x2b, 0xef, 0x70, 0x48, 0xaa, 0xa1, 0xe2, 0x96, 0x97, 0x00, 0xe5, 0x50, 0x27, 0x07, 0x3e,
	0xe7, 0xfa, 0x98, 0xdf, 0x50, 0x8d, 0x2b, 0x18, 0x96, 0x7f, 0xcc, 0xeb, 0x34, 0xc6, 0x6f, 0xe6,
	0x9d, 0x54, 0x16, 0x8f, 0x26, 0x87, 0xd5, 0x44, 0xeb, 0xbf, 0xd0, 0xfd, 0x6c, 0x5f, 0x5a, 0xf8,
	0x7d, 0x5d, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x06, 0x67, 0xfc, 0xd2, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*Response, error)
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	GetInfo(ctx context.Context, in *GetInfoRequest, opts ...client.CallOption) (*UserInfoResponse, error)
	GetProfile(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*UserProfile, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	UpdateInfo(ctx context.Context, in *UpdateInfoRequest, opts ...client.CallOption) (*Response, error)
	UpdateTeamAndGroupForUsers(ctx context.Context, in *UpdateTeamGroupRequest, opts ...client.CallOption) (*Response, error)
}

type userServiceClient struct {
	c           client.Client
	serviceName string
}

func NewUserServiceClient(serviceName string, c client.Client) UserServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "user"
	}
	return &userServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Register", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...client.CallOption) (*UserInfoResponse, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetInfo", in)
	out := new(UserInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetProfile(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*UserProfile, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetProfile", in)
	out := new(UserProfile)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateInfo(ctx context.Context, in *UpdateInfoRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.UpdateInfo", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateTeamAndGroupForUsers(ctx context.Context, in *UpdateTeamGroupRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.UpdateTeamAndGroupForUsers", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Register(context.Context, *RegisterRequest, *Response) error
	Login(context.Context, *LoginRequest, *LoginResponse) error
	GetInfo(context.Context, *GetInfoRequest, *UserInfoResponse) error
	GetProfile(context.Context, *GetRequest, *UserProfile) error
	List(context.Context, *ListRequest, *ListResponse) error
	UpdateInfo(context.Context, *UpdateInfoRequest, *Response) error
	UpdateTeamAndGroupForUsers(context.Context, *UpdateTeamGroupRequest, *Response) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserService{hdlr}, opts...))
}

type UserService struct {
	UserServiceHandler
}

func (h *UserService) Register(ctx context.Context, in *RegisterRequest, out *Response) error {
	return h.UserServiceHandler.Register(ctx, in, out)
}

func (h *UserService) Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.UserServiceHandler.Login(ctx, in, out)
}

func (h *UserService) GetInfo(ctx context.Context, in *GetInfoRequest, out *UserInfoResponse) error {
	return h.UserServiceHandler.GetInfo(ctx, in, out)
}

func (h *UserService) GetProfile(ctx context.Context, in *GetRequest, out *UserProfile) error {
	return h.UserServiceHandler.GetProfile(ctx, in, out)
}

func (h *UserService) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.UserServiceHandler.List(ctx, in, out)
}

func (h *UserService) UpdateInfo(ctx context.Context, in *UpdateInfoRequest, out *Response) error {
	return h.UserServiceHandler.UpdateInfo(ctx, in, out)
}

func (h *UserService) UpdateTeamAndGroupForUsers(ctx context.Context, in *UpdateTeamGroupRequest, out *Response) error {
	return h.UserServiceHandler.UpdateTeamAndGroupForUsers(ctx, in, out)
}