// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.1
// source: internal/app/proto/grpc.proto

package proto

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

type CreateShortResponse_Status int32

const (
	CreateShortResponse_Unknown  CreateShortResponse_Status = 0
	CreateShortResponse_Created  CreateShortResponse_Status = 1
	CreateShortResponse_Conflict CreateShortResponse_Status = 2
)

// Enum value maps for CreateShortResponse_Status.
var (
	CreateShortResponse_Status_name = map[int32]string{
		0: "Unknown",
		1: "Created",
		2: "Conflict",
	}
	CreateShortResponse_Status_value = map[string]int32{
		"Unknown":  0,
		"Created":  1,
		"Conflict": 2,
	}
)

func (x CreateShortResponse_Status) Enum() *CreateShortResponse_Status {
	p := new(CreateShortResponse_Status)
	*p = x
	return p
}

func (x CreateShortResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CreateShortResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_app_proto_grpc_proto_enumTypes[0].Descriptor()
}

func (CreateShortResponse_Status) Type() protoreflect.EnumType {
	return &file_internal_app_proto_grpc_proto_enumTypes[0]
}

func (x CreateShortResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CreateShortResponse_Status.Descriptor instead.
func (CreateShortResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_internal_app_proto_grpc_proto_rawDescGZIP(), []int{2, 0}
}

type Short struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl    string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	OriginalUrl string `protobuf:"bytes,2,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
	UserId      string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Short) Reset() {
	*x = Short{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Short) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Short) ProtoMessage() {}

func (x *Short) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Short.ProtoReflect.Descriptor instead.
func (*Short) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *Short) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

func (x *Short) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

func (x *Short) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CreateShortRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Url    string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *CreateShortRequest) Reset() {
	*x = CreateShortRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShortRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShortRequest) ProtoMessage() {}

func (x *CreateShortRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShortRequest.ProtoReflect.Descriptor instead.
func (*CreateShortRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *CreateShortRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateShortRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type CreateShortResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status CreateShortResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=grpc.CreateShortResponse_Status" json:"status,omitempty"`
	Short  *Short                     `protobuf:"bytes,2,opt,name=short,proto3" json:"short,omitempty"`
}

func (x *CreateShortResponse) Reset() {
	*x = CreateShortResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShortResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShortResponse) ProtoMessage() {}

func (x *CreateShortResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShortResponse.ProtoReflect.Descriptor instead.
func (*CreateShortResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *CreateShortResponse) GetStatus() CreateShortResponse_Status {
	if x != nil {
		return x.Status
	}
	return CreateShortResponse_Unknown
}

func (x *CreateShortResponse) GetShort() *Short {
	if x != nil {
		return x.Short
	}
	return nil
}

type CreateShortBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string                                          `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Shorts []*CreateShortBatchRequest_CreateShortBatchItem `protobuf:"bytes,2,rep,name=shorts,proto3" json:"shorts,omitempty"`
}

func (x *CreateShortBatchRequest) Reset() {
	*x = CreateShortBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShortBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShortBatchRequest) ProtoMessage() {}

func (x *CreateShortBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShortBatchRequest.ProtoReflect.Descriptor instead.
func (*CreateShortBatchRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *CreateShortBatchRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateShortBatchRequest) GetShorts() []*CreateShortBatchRequest_CreateShortBatchItem {
	if x != nil {
		return x.Shorts
	}
	return nil
}

type CreateShortBatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*CreateShortBatchResponse_CreateShortBatchResponseItem `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *CreateShortBatchResponse) Reset() {
	*x = CreateShortBatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShortBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShortBatchResponse) ProtoMessage() {}

func (x *CreateShortBatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShortBatchResponse.ProtoReflect.Descriptor instead.
func (*CreateShortBatchResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *CreateShortBatchResponse) GetResults() []*CreateShortBatchResponse_CreateShortBatchResponseItem {
	if x != nil {
		return x.Results
	}
	return nil
}

type DeleteShortBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Shorts []string `protobuf:"bytes,2,rep,name=shorts,proto3" json:"shorts,omitempty"`
}

func (x *DeleteShortBatchRequest) Reset() {
	*x = DeleteShortBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_grpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteShortBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteShortBatchRequest) ProtoMessage() {}

func (x *DeleteShortBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_grpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteShortBatchRequest.ProtoReflect.Descriptor instead.
func (*DeleteShortBatchRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_grpc_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteShortBatchRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DeleteShortBatchRequest) GetShorts() []string {
	if x != nil {
		return x.Shorts
	}
	return nil
}

type DeleteShortBatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteShortBatchResponse) Reset() {
	*x = DeleteShortBatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_grpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteShortBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteShortBatchResponse) ProtoMessage() {}

func (x *DeleteShortBatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_grpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteShortBatchResponse.ProtoReflect.Descriptor instead.
func (*DeleteShortBatchResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_grpc_proto_rawDescGZIP(), []int{6}
}

type ListShortsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ListShortsRequest) Reset() {
	*x = ListShortsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_grpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListShortsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShortsRequest) ProtoMessage() {}

func (x *ListShortsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_grpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShortsRequest.ProtoReflect.Descriptor instead.
func (*ListShortsRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_grpc_proto_rawDescGZIP(), []int{7}
}

func (x *ListShortsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ListShortsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shorts []*Short `protobuf:"bytes,1,rep,name=shorts,proto3" json:"shorts,omitempty"`
}

func (x *ListShortsResponse) Reset() {
	*x = ListShortsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_grpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListShortsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShortsResponse) ProtoMessage() {}

func (x *ListShortsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_grpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShortsResponse.ProtoReflect.Descriptor instead.
func (*ListShortsResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_grpc_proto_rawDescGZIP(), []int{8}
}

func (x *ListShortsResponse) GetShorts() []*Short {
	if x != nil {
		return x.Shorts
	}
	return nil
}

type InternalStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UrlsCount  int32 `protobuf:"varint,1,opt,name=urls_count,json=urlsCount,proto3" json:"urls_count,omitempty"`    // количество сокращённых URL в сервисе
	UsersCount int32 `protobuf:"varint,2,opt,name=users_count,json=usersCount,proto3" json:"users_count,omitempty"` // количество пользователей в сервисе
}

func (x *InternalStatsResponse) Reset() {
	*x = InternalStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_grpc_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalStatsResponse) ProtoMessage() {}

func (x *InternalStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_grpc_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalStatsResponse.ProtoReflect.Descriptor instead.
func (*InternalStatsResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_grpc_proto_rawDescGZIP(), []int{9}
}

func (x *InternalStatsResponse) GetUrlsCount() int32 {
	if x != nil {
		return x.UrlsCount
	}
	return 0
}

func (x *InternalStatsResponse) GetUsersCount() int32 {
	if x != nil {
		return x.UsersCount
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_grpc_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_grpc_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_grpc_proto_rawDescGZIP(), []int{10}
}

type CreateShortBatchRequest_CreateShortBatchItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationId string `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	OriginalUrl   string `protobuf:"bytes,2,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *CreateShortBatchRequest_CreateShortBatchItem) Reset() {
	*x = CreateShortBatchRequest_CreateShortBatchItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_grpc_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShortBatchRequest_CreateShortBatchItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShortBatchRequest_CreateShortBatchItem) ProtoMessage() {}

func (x *CreateShortBatchRequest_CreateShortBatchItem) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_grpc_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShortBatchRequest_CreateShortBatchItem.ProtoReflect.Descriptor instead.
func (*CreateShortBatchRequest_CreateShortBatchItem) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_grpc_proto_rawDescGZIP(), []int{3, 0}
}

func (x *CreateShortBatchRequest_CreateShortBatchItem) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *CreateShortBatchRequest_CreateShortBatchItem) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type CreateShortBatchResponse_CreateShortBatchResponseItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationId string `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	Short         *Short `protobuf:"bytes,2,opt,name=short,proto3" json:"short,omitempty"`
}

func (x *CreateShortBatchResponse_CreateShortBatchResponseItem) Reset() {
	*x = CreateShortBatchResponse_CreateShortBatchResponseItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_proto_grpc_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShortBatchResponse_CreateShortBatchResponseItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShortBatchResponse_CreateShortBatchResponseItem) ProtoMessage() {}

func (x *CreateShortBatchResponse_CreateShortBatchResponseItem) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_proto_grpc_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShortBatchResponse_CreateShortBatchResponseItem.ProtoReflect.Descriptor instead.
func (*CreateShortBatchResponse_CreateShortBatchResponseItem) Descriptor() ([]byte, []int) {
	return file_internal_app_proto_grpc_proto_rawDescGZIP(), []int{4, 0}
}

func (x *CreateShortBatchResponse_CreateShortBatchResponseItem) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *CreateShortBatchResponse_CreateShortBatchResponseItem) GetShort() *Short {
	if x != nil {
		return x.Short
	}
	return nil
}

var File_internal_app_proto_grpc_proto protoreflect.FileDescriptor

var file_internal_app_proto_grpc_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x60, 0x0a, 0x05, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0xa4, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x38, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x20, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x05, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x22, 0x30, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x10,
	0x01, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x10, 0x02, 0x22,
	0xe0, 0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x4a, 0x0a, 0x06, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x73,
	0x1a, 0x60, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55,
	0x72, 0x6c, 0x22, 0xdb, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x55, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x3b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0x68, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x22, 0x4a, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x73, 0x22, 0x1a, 0x0a, 0x18,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x06,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x73, 0x22, 0x57, 0x0a, 0x15, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x72,
	0x6c, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x75, 0x72, 0x6c, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x32, 0xfe, 0x02, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x12, 0x42, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x12, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1d, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x73, 0x46, 0x6f, 0x72, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x1d, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0d, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x0b, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x75, 0x74, 0x61, 0x6c, 0x65, 0x78, 0x65, 0x79, 0x2f, 0x67, 0x6f, 0x2d,
	0x70, 0x72, 0x61, 0x63, 0x74, 0x69, 0x63, 0x75, 0x6d, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_app_proto_grpc_proto_rawDescOnce sync.Once
	file_internal_app_proto_grpc_proto_rawDescData = file_internal_app_proto_grpc_proto_rawDesc
)

func file_internal_app_proto_grpc_proto_rawDescGZIP() []byte {
	file_internal_app_proto_grpc_proto_rawDescOnce.Do(func() {
		file_internal_app_proto_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_app_proto_grpc_proto_rawDescData)
	})
	return file_internal_app_proto_grpc_proto_rawDescData
}

var file_internal_app_proto_grpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_app_proto_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_internal_app_proto_grpc_proto_goTypes = []interface{}{
	(CreateShortResponse_Status)(0),                               // 0: grpc.CreateShortResponse.Status
	(*Short)(nil),                                                 // 1: grpc.Short
	(*CreateShortRequest)(nil),                                    // 2: grpc.CreateShortRequest
	(*CreateShortResponse)(nil),                                   // 3: grpc.CreateShortResponse
	(*CreateShortBatchRequest)(nil),                               // 4: grpc.CreateShortBatchRequest
	(*CreateShortBatchResponse)(nil),                              // 5: grpc.CreateShortBatchResponse
	(*DeleteShortBatchRequest)(nil),                               // 6: grpc.DeleteShortBatchRequest
	(*DeleteShortBatchResponse)(nil),                              // 7: grpc.DeleteShortBatchResponse
	(*ListShortsRequest)(nil),                                     // 8: grpc.ListShortsRequest
	(*ListShortsResponse)(nil),                                    // 9: grpc.ListShortsResponse
	(*InternalStatsResponse)(nil),                                 // 10: grpc.InternalStatsResponse
	(*Empty)(nil),                                                 // 11: grpc.Empty
	(*CreateShortBatchRequest_CreateShortBatchItem)(nil),          // 12: grpc.CreateShortBatchRequest.CreateShortBatchItem
	(*CreateShortBatchResponse_CreateShortBatchResponseItem)(nil), // 13: grpc.CreateShortBatchResponse.CreateShortBatchResponseItem
}
var file_internal_app_proto_grpc_proto_depIdxs = []int32{
	0,  // 0: grpc.CreateShortResponse.status:type_name -> grpc.CreateShortResponse.Status
	1,  // 1: grpc.CreateShortResponse.short:type_name -> grpc.Short
	12, // 2: grpc.CreateShortBatchRequest.shorts:type_name -> grpc.CreateShortBatchRequest.CreateShortBatchItem
	13, // 3: grpc.CreateShortBatchResponse.results:type_name -> grpc.CreateShortBatchResponse.CreateShortBatchResponseItem
	1,  // 4: grpc.ListShortsResponse.shorts:type_name -> grpc.Short
	1,  // 5: grpc.CreateShortBatchResponse.CreateShortBatchResponseItem.short:type_name -> grpc.Short
	2,  // 6: grpc.Shortener.CreateShort:input_type -> grpc.CreateShortRequest
	4,  // 7: grpc.Shortener.CreateShortBatch:input_type -> grpc.CreateShortBatchRequest
	8,  // 8: grpc.Shortener.GetShortsForCurrentUser:input_type -> grpc.ListShortsRequest
	6,  // 9: grpc.Shortener.DeleteUserShorts:input_type -> grpc.DeleteShortBatchRequest
	11, // 10: grpc.Shortener.InternalStats:input_type -> grpc.Empty
	3,  // 11: grpc.Shortener.CreateShort:output_type -> grpc.CreateShortResponse
	5,  // 12: grpc.Shortener.CreateShortBatch:output_type -> grpc.CreateShortBatchResponse
	9,  // 13: grpc.Shortener.GetShortsForCurrentUser:output_type -> grpc.ListShortsResponse
	7,  // 14: grpc.Shortener.DeleteUserShorts:output_type -> grpc.DeleteShortBatchResponse
	10, // 15: grpc.Shortener.InternalStats:output_type -> grpc.InternalStatsResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_internal_app_proto_grpc_proto_init() }
func file_internal_app_proto_grpc_proto_init() {
	if File_internal_app_proto_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_app_proto_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Short); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_app_proto_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShortRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_app_proto_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShortResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_app_proto_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShortBatchRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_app_proto_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShortBatchResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_app_proto_grpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteShortBatchRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_app_proto_grpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteShortBatchResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_app_proto_grpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListShortsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_app_proto_grpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListShortsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_app_proto_grpc_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalStatsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_app_proto_grpc_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_app_proto_grpc_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShortBatchRequest_CreateShortBatchItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_app_proto_grpc_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShortBatchResponse_CreateShortBatchResponseItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_app_proto_grpc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_app_proto_grpc_proto_goTypes,
		DependencyIndexes: file_internal_app_proto_grpc_proto_depIdxs,
		EnumInfos:         file_internal_app_proto_grpc_proto_enumTypes,
		MessageInfos:      file_internal_app_proto_grpc_proto_msgTypes,
	}.Build()
	File_internal_app_proto_grpc_proto = out.File
	file_internal_app_proto_grpc_proto_rawDesc = nil
	file_internal_app_proto_grpc_proto_goTypes = nil
	file_internal_app_proto_grpc_proto_depIdxs = nil
}
