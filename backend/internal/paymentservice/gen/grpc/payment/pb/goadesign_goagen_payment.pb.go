// Code generated with goa v3.7.0, DO NOT EDIT.
//
// payment protocol buffer definition
//
// Command:
// $ goa gen
// github.com/JordanRad/play-j/backend/internal/design/payment-service -o
// ./internal/paymentservice

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: goadesign_goagen_payment.proto

package paymentpb

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

type GetPaymentsByAccountIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Account ID
	AccountId int32 `protobuf:"zigzag32,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// Resource array size
	Limit int32 `protobuf:"zigzag32,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetPaymentsByAccountIDRequest) Reset() {
	*x = GetPaymentsByAccountIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaymentsByAccountIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentsByAccountIDRequest) ProtoMessage() {}

func (x *GetPaymentsByAccountIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentsByAccountIDRequest.ProtoReflect.Descriptor instead.
func (*GetPaymentsByAccountIDRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_payment_proto_rawDescGZIP(), []int{0}
}

func (x *GetPaymentsByAccountIDRequest) GetAccountId() int32 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *GetPaymentsByAccountIDRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetPaymentsByAccountIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total     uint32             `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Resources []*PaymentResponse `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *GetPaymentsByAccountIDResponse) Reset() {
	*x = GetPaymentsByAccountIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaymentsByAccountIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentsByAccountIDResponse) ProtoMessage() {}

func (x *GetPaymentsByAccountIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentsByAccountIDResponse.ProtoReflect.Descriptor instead.
func (*GetPaymentsByAccountIDResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_payment_proto_rawDescGZIP(), []int{1}
}

func (x *GetPaymentsByAccountIDResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GetPaymentsByAccountIDResponse) GetResources() []*PaymentResponse {
	if x != nil {
		return x.Resources
	}
	return nil
}

type PaymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt     string  `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	PaymentNumber string  `protobuf:"bytes,3,opt,name=payment_number,json=paymentNumber,proto3" json:"payment_number,omitempty"`
	Amount        float32 `protobuf:"fixed32,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *PaymentResponse) Reset() {
	*x = PaymentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_payment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentResponse) ProtoMessage() {}

func (x *PaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_payment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentResponse.ProtoReflect.Descriptor instead.
func (*PaymentResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_payment_proto_rawDescGZIP(), []int{2}
}

func (x *PaymentResponse) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PaymentResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *PaymentResponse) GetPaymentNumber() string {
	if x != nil {
		return x.PaymentNumber
	}
	return ""
}

func (x *PaymentResponse) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_goadesign_goagen_payment_proto protoreflect.FileDescriptor

var file_goadesign_goagen_payment_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x67, 0x6f, 0x61, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x67, 0x6f, 0x61, 0x67,
	0x65, 0x6e, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x54, 0x0a, 0x1d, 0x47, 0x65, 0x74,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11, 0x52, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22,
	0x6e, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x36, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22,
	0x7f, 0x0a, 0x0f, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x32, 0x74, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x69, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x26, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goadesign_goagen_payment_proto_rawDescOnce sync.Once
	file_goadesign_goagen_payment_proto_rawDescData = file_goadesign_goagen_payment_proto_rawDesc
)

func file_goadesign_goagen_payment_proto_rawDescGZIP() []byte {
	file_goadesign_goagen_payment_proto_rawDescOnce.Do(func() {
		file_goadesign_goagen_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_goadesign_goagen_payment_proto_rawDescData)
	})
	return file_goadesign_goagen_payment_proto_rawDescData
}

var file_goadesign_goagen_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_goadesign_goagen_payment_proto_goTypes = []interface{}{
	(*GetPaymentsByAccountIDRequest)(nil),  // 0: payment.GetPaymentsByAccountIDRequest
	(*GetPaymentsByAccountIDResponse)(nil), // 1: payment.GetPaymentsByAccountIDResponse
	(*PaymentResponse)(nil),                // 2: payment.PaymentResponse
}
var file_goadesign_goagen_payment_proto_depIdxs = []int32{
	2, // 0: payment.GetPaymentsByAccountIDResponse.resources:type_name -> payment.PaymentResponse
	0, // 1: payment.Payment.GetPaymentsByAccountID:input_type -> payment.GetPaymentsByAccountIDRequest
	1, // 2: payment.Payment.GetPaymentsByAccountID:output_type -> payment.GetPaymentsByAccountIDResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_goadesign_goagen_payment_proto_init() }
func file_goadesign_goagen_payment_proto_init() {
	if File_goadesign_goagen_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goadesign_goagen_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPaymentsByAccountIDRequest); i {
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
		file_goadesign_goagen_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPaymentsByAccountIDResponse); i {
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
		file_goadesign_goagen_payment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentResponse); i {
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
			RawDescriptor: file_goadesign_goagen_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goadesign_goagen_payment_proto_goTypes,
		DependencyIndexes: file_goadesign_goagen_payment_proto_depIdxs,
		MessageInfos:      file_goadesign_goagen_payment_proto_msgTypes,
	}.Build()
	File_goadesign_goagen_payment_proto = out.File
	file_goadesign_goagen_payment_proto_rawDesc = nil
	file_goadesign_goagen_payment_proto_goTypes = nil
	file_goadesign_goagen_payment_proto_depIdxs = nil
}