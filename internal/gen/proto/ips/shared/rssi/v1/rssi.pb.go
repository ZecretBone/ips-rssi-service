// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: ips/shared/rssi/v1/rssi.proto

package rssiv1

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

type RSSI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ssid        string                 `protobuf:"bytes,1,opt,name=ssid,proto3" json:"ssid,omitempty"`
	MacAddress  string                 `protobuf:"bytes,2,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	Strength    float32                `protobuf:"fixed32,3,opt,name=strength,proto3" json:"strength,omitempty"`
	PollingRate int32                  `protobuf:"varint,4,opt,name=polling_rate,json=pollingRate,proto3" json:"polling_rate,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *RSSI) Reset() {
	*x = RSSI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ips_shared_rssi_v1_rssi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RSSI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RSSI) ProtoMessage() {}

func (x *RSSI) ProtoReflect() protoreflect.Message {
	mi := &file_ips_shared_rssi_v1_rssi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RSSI.ProtoReflect.Descriptor instead.
func (*RSSI) Descriptor() ([]byte, []int) {
	return file_ips_shared_rssi_v1_rssi_proto_rawDescGZIP(), []int{0}
}

func (x *RSSI) GetSsid() string {
	if x != nil {
		return x.Ssid
	}
	return ""
}

func (x *RSSI) GetMacAddress() string {
	if x != nil {
		return x.MacAddress
	}
	return ""
}

func (x *RSSI) GetStrength() float32 {
	if x != nil {
		return x.Strength
	}
	return 0
}

func (x *RSSI) GetPollingRate() int32 {
	if x != nil {
		return x.PollingRate
	}
	return 0
}

func (x *RSSI) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	Z float32 `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ips_shared_rssi_v1_rssi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_ips_shared_rssi_v1_rssi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_ips_shared_rssi_v1_rssi_proto_rawDescGZIP(), []int{1}
}

func (x *Position) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Position) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Position) GetZ() float32 {
	if x != nil {
		return x.Z
	}
	return 0
}

type DeviceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Models   string `protobuf:"bytes,2,opt,name=models,proto3" json:"models,omitempty"`
}

func (x *DeviceInfo) Reset() {
	*x = DeviceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ips_shared_rssi_v1_rssi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceInfo) ProtoMessage() {}

func (x *DeviceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ips_shared_rssi_v1_rssi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceInfo.ProtoReflect.Descriptor instead.
func (*DeviceInfo) Descriptor() ([]byte, []int) {
	return file_ips_shared_rssi_v1_rssi_proto_rawDescGZIP(), []int{2}
}

func (x *DeviceInfo) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *DeviceInfo) GetModels() string {
	if x != nil {
		return x.Models
	}
	return ""
}

var File_ips_shared_rssi_v1_rssi_proto protoreflect.FileDescriptor

var file_ips_shared_rssi_v1_rssi_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x69, 0x70, 0x73, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x72, 0x73, 0x73,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x73, 0x73, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x69, 0x70, 0x73, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x72, 0x73, 0x73, 0x69,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x01, 0x0a, 0x04, 0x52, 0x53, 0x53, 0x49, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x73, 0x69,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x73, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74,
	0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x34, 0x0a, 0x08,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x01, 0x7a, 0x22, 0x41, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x42, 0xe3, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x70,
	0x73, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x72, 0x73, 0x73, 0x69, 0x2e, 0x76, 0x31,
	0x42, 0x09, 0x52, 0x73, 0x73, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x53, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x5a, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x42, 0x6f, 0x6e, 0x65, 0x2f, 0x69, 0x70, 0x73, 0x2d, 0x72, 0x73, 0x73, 0x69, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x70, 0x73, 0x2f, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2f, 0x72, 0x73, 0x73, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x73, 0x73, 0x69,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x49, 0x53, 0x52, 0xaa, 0x02, 0x12, 0x49, 0x70, 0x73, 0x2e, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x52, 0x73, 0x73, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x12,
	0x49, 0x70, 0x73, 0x5c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5c, 0x52, 0x73, 0x73, 0x69, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x1e, 0x49, 0x70, 0x73, 0x5c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5c,
	0x52, 0x73, 0x73, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x49, 0x70, 0x73, 0x3a, 0x3a, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x3a, 0x3a, 0x52, 0x73, 0x73, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_ips_shared_rssi_v1_rssi_proto_rawDescOnce sync.Once
	file_ips_shared_rssi_v1_rssi_proto_rawDescData = file_ips_shared_rssi_v1_rssi_proto_rawDesc
)

func file_ips_shared_rssi_v1_rssi_proto_rawDescGZIP() []byte {
	file_ips_shared_rssi_v1_rssi_proto_rawDescOnce.Do(func() {
		file_ips_shared_rssi_v1_rssi_proto_rawDescData = protoimpl.X.CompressGZIP(file_ips_shared_rssi_v1_rssi_proto_rawDescData)
	})
	return file_ips_shared_rssi_v1_rssi_proto_rawDescData
}

var file_ips_shared_rssi_v1_rssi_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ips_shared_rssi_v1_rssi_proto_goTypes = []interface{}{
	(*RSSI)(nil),                  // 0: ips.shared.rssi.v1.RSSI
	(*Position)(nil),              // 1: ips.shared.rssi.v1.Position
	(*DeviceInfo)(nil),            // 2: ips.shared.rssi.v1.DeviceInfo
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_ips_shared_rssi_v1_rssi_proto_depIdxs = []int32{
	3, // 0: ips.shared.rssi.v1.RSSI.created_at:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ips_shared_rssi_v1_rssi_proto_init() }
func file_ips_shared_rssi_v1_rssi_proto_init() {
	if File_ips_shared_rssi_v1_rssi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ips_shared_rssi_v1_rssi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RSSI); i {
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
		file_ips_shared_rssi_v1_rssi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
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
		file_ips_shared_rssi_v1_rssi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceInfo); i {
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
			RawDescriptor: file_ips_shared_rssi_v1_rssi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ips_shared_rssi_v1_rssi_proto_goTypes,
		DependencyIndexes: file_ips_shared_rssi_v1_rssi_proto_depIdxs,
		MessageInfos:      file_ips_shared_rssi_v1_rssi_proto_msgTypes,
	}.Build()
	File_ips_shared_rssi_v1_rssi_proto = out.File
	file_ips_shared_rssi_v1_rssi_proto_rawDesc = nil
	file_ips_shared_rssi_v1_rssi_proto_goTypes = nil
	file_ips_shared_rssi_v1_rssi_proto_depIdxs = nil
}
