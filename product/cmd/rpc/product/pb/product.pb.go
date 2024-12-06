// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0
// source: product.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name              string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ProductType       int32                  `protobuf:"varint,3,opt,name=product_type,json=productType,proto3" json:"product_type,omitempty"`
	CategoryId        int32                  `protobuf:"varint,4,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	StartingPrice     float64                `protobuf:"fixed64,5,opt,name=starting_price,json=startingPrice,proto3" json:"starting_price,omitempty"`
	TotalStock        int32                  `protobuf:"varint,6,opt,name=total_stock,json=totalStock,proto3" json:"total_stock,omitempty"`
	MainPicture       string                 `protobuf:"bytes,7,opt,name=main_picture,json=mainPicture,proto3" json:"main_picture,omitempty"`
	RemoteAreaPostage float64                `protobuf:"fixed64,8,opt,name=remote_area_postage,json=remoteAreaPostage,proto3" json:"remote_area_postage,omitempty"`
	SingleBuyLimit    int32                  `protobuf:"varint,9,opt,name=single_buy_limit,json=singleBuyLimit,proto3" json:"single_buy_limit,omitempty"`
	Remark            string                 `protobuf:"bytes,10,opt,name=remark,proto3" json:"remark,omitempty"`
	CreateTime        *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime        *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	IsEnable          int32                  `protobuf:"varint,13,opt,name=is_enable,json=isEnable,proto3" json:"is_enable,omitempty"`
	CreateUser        string                 `protobuf:"bytes,14,opt,name=create_user,json=createUser,proto3" json:"create_user,omitempty"`
	UpdateUser        string                 `protobuf:"bytes,15,opt,name=update_user,json=updateUser,proto3" json:"update_user,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	mi := &file_product_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetProductType() int32 {
	if x != nil {
		return x.ProductType
	}
	return 0
}

func (x *Product) GetCategoryId() int32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *Product) GetStartingPrice() float64 {
	if x != nil {
		return x.StartingPrice
	}
	return 0
}

func (x *Product) GetTotalStock() int32 {
	if x != nil {
		return x.TotalStock
	}
	return 0
}

func (x *Product) GetMainPicture() string {
	if x != nil {
		return x.MainPicture
	}
	return ""
}

func (x *Product) GetRemoteAreaPostage() float64 {
	if x != nil {
		return x.RemoteAreaPostage
	}
	return 0
}

func (x *Product) GetSingleBuyLimit() int32 {
	if x != nil {
		return x.SingleBuyLimit
	}
	return 0
}

func (x *Product) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *Product) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Product) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Product) GetIsEnable() int32 {
	if x != nil {
		return x.IsEnable
	}
	return 0
}

func (x *Product) GetCreateUser() string {
	if x != nil {
		return x.CreateUser
	}
	return ""
}

func (x *Product) GetUpdateUser() string {
	if x != nil {
		return x.UpdateUser
	}
	return ""
}

type CreateProductReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name              string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ProductType       int32   `protobuf:"varint,2,opt,name=product_type,json=productType,proto3" json:"product_type,omitempty"`
	CategoryId        int32   `protobuf:"varint,3,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	StartingPrice     float64 `protobuf:"fixed64,4,opt,name=starting_price,json=startingPrice,proto3" json:"starting_price,omitempty"`
	TotalStock        int32   `protobuf:"varint,5,opt,name=total_stock,json=totalStock,proto3" json:"total_stock,omitempty"`
	MainPicture       string  `protobuf:"bytes,6,opt,name=main_picture,json=mainPicture,proto3" json:"main_picture,omitempty"`
	RemoteAreaPostage float64 `protobuf:"fixed64,7,opt,name=remote_area_postage,json=remoteAreaPostage,proto3" json:"remote_area_postage,omitempty"`
	SingleBuyLimit    int32   `protobuf:"varint,8,opt,name=single_buy_limit,json=singleBuyLimit,proto3" json:"single_buy_limit,omitempty"`
	Remark            string  `protobuf:"bytes,9,opt,name=remark,proto3" json:"remark,omitempty"`
}

func (x *CreateProductReq) Reset() {
	*x = CreateProductReq{}
	mi := &file_product_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductReq) ProtoMessage() {}

func (x *CreateProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductReq.ProtoReflect.Descriptor instead.
func (*CreateProductReq) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProductReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductReq) GetProductType() int32 {
	if x != nil {
		return x.ProductType
	}
	return 0
}

func (x *CreateProductReq) GetCategoryId() int32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *CreateProductReq) GetStartingPrice() float64 {
	if x != nil {
		return x.StartingPrice
	}
	return 0
}

func (x *CreateProductReq) GetTotalStock() int32 {
	if x != nil {
		return x.TotalStock
	}
	return 0
}

func (x *CreateProductReq) GetMainPicture() string {
	if x != nil {
		return x.MainPicture
	}
	return ""
}

func (x *CreateProductReq) GetRemoteAreaPostage() float64 {
	if x != nil {
		return x.RemoteAreaPostage
	}
	return 0
}

func (x *CreateProductReq) GetSingleBuyLimit() int32 {
	if x != nil {
		return x.SingleBuyLimit
	}
	return 0
}

func (x *CreateProductReq) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type CreateProductResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateProductResp) Reset() {
	*x = CreateProductResp{}
	mi := &file_product_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProductResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductResp) ProtoMessage() {}

func (x *CreateProductResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductResp.ProtoReflect.Descriptor instead.
func (*CreateProductResp) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{2}
}

type GetProductListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetProductListReq) Reset() {
	*x = GetProductListReq{}
	mi := &file_product_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductListReq) ProtoMessage() {}

func (x *GetProductListReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductListReq.ProtoReflect.Descriptor instead.
func (*GetProductListReq) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{3}
}

type GetProductListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductList []*Product `protobuf:"bytes,1,rep,name=ProductList,proto3" json:"ProductList,omitempty"`
}

func (x *GetProductListResp) Reset() {
	*x = GetProductListResp{}
	mi := &file_product_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductListResp) ProtoMessage() {}

func (x *GetProductListResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductListResp.ProtoReflect.Descriptor instead.
func (*GetProductListResp) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{4}
}

func (x *GetProductListResp) GetProductList() []*Product {
	if x != nil {
		return x.ProductList
	}
	return nil
}

type DeleteProductReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteProductReq) Reset() {
	*x = DeleteProductReq{}
	mi := &file_product_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductReq) ProtoMessage() {}

func (x *DeleteProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductReq.ProtoReflect.Descriptor instead.
func (*DeleteProductReq) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteProductReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteProductResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteProductResp) Reset() {
	*x = DeleteProductResp{}
	mi := &file_product_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteProductResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductResp) ProtoMessage() {}

func (x *DeleteProductResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductResp.ProtoReflect.Descriptor instead.
func (*DeleteProductResp) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{6}
}

type GetProductReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetProductReq) Reset() {
	*x = GetProductReq{}
	mi := &file_product_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductReq) ProtoMessage() {}

func (x *GetProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductReq.ProtoReflect.Descriptor instead.
func (*GetProductReq) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{7}
}

func (x *GetProductReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetProductResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetProductResp) Reset() {
	*x = GetProductResp{}
	mi := &file_product_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductResp) ProtoMessage() {}

func (x *GetProductResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductResp.ProtoReflect.Descriptor instead.
func (*GetProductResp) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{8}
}

var File_product_proto protoreflect.FileDescriptor

var file_product_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x04, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x61, 0x69, 0x6e, 0x50, 0x69, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x72,
	0x65, 0x61, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x11, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x72, 0x65, 0x61, 0x50, 0x6f, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x5f, 0x62, 0x75,
	0x79, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x73,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x42, 0x75, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x69, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x22, 0xc7,
	0x02, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x70, 0x69, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x61, 0x69, 0x6e, 0x50,
	0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x5f, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x11, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x72, 0x65, 0x61, 0x50,
	0x6f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65,
	0x5f, 0x62, 0x75, 0x79, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0e, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x42, 0x75, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x13, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x22, 0x43, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2d, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x0b, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x22, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x1f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x32, 0xfb, 0x01, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x3c, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3f, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x15, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3c,
	0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x14, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x33, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_product_proto_rawDescOnce sync.Once
	file_product_proto_rawDescData = file_product_proto_rawDesc
)

func file_product_proto_rawDescGZIP() []byte {
	file_product_proto_rawDescOnce.Do(func() {
		file_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_proto_rawDescData)
	})
	return file_product_proto_rawDescData
}

var file_product_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_product_proto_goTypes = []any{
	(*Product)(nil),               // 0: pb.Product
	(*CreateProductReq)(nil),      // 1: pb.CreateProductReq
	(*CreateProductResp)(nil),     // 2: pb.CreateProductResp
	(*GetProductListReq)(nil),     // 3: pb.GetProductListReq
	(*GetProductListResp)(nil),    // 4: pb.GetProductListResp
	(*DeleteProductReq)(nil),      // 5: pb.DeleteProductReq
	(*DeleteProductResp)(nil),     // 6: pb.DeleteProductResp
	(*GetProductReq)(nil),         // 7: pb.GetProductReq
	(*GetProductResp)(nil),        // 8: pb.GetProductResp
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_product_proto_depIdxs = []int32{
	9, // 0: pb.Product.create_time:type_name -> google.protobuf.Timestamp
	9, // 1: pb.Product.update_time:type_name -> google.protobuf.Timestamp
	0, // 2: pb.GetProductListResp.ProductList:type_name -> pb.Product
	1, // 3: pb.product.CreateProduct:input_type -> pb.CreateProductReq
	3, // 4: pb.product.GetProductList:input_type -> pb.GetProductListReq
	5, // 5: pb.product.DeleteProduct:input_type -> pb.DeleteProductReq
	7, // 6: pb.product.GetProduct:input_type -> pb.GetProductReq
	2, // 7: pb.product.CreateProduct:output_type -> pb.CreateProductResp
	4, // 8: pb.product.GetProductList:output_type -> pb.GetProductListResp
	6, // 9: pb.product.DeleteProduct:output_type -> pb.DeleteProductResp
	8, // 10: pb.product.GetProduct:output_type -> pb.GetProductResp
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_product_proto_init() }
func file_product_proto_init() {
	if File_product_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_proto_goTypes,
		DependencyIndexes: file_product_proto_depIdxs,
		MessageInfos:      file_product_proto_msgTypes,
	}.Build()
	File_product_proto = out.File
	file_product_proto_rawDesc = nil
	file_product_proto_goTypes = nil
	file_product_proto_depIdxs = nil
}
