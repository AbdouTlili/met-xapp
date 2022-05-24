// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.8
// source: message.proto

package northbound

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

type Kpi_NSSMF int32

const (
	Kpi_NFVO Kpi_NSSMF = 0
	Kpi_RAN  Kpi_NSSMF = 1
	Kpi_SO   Kpi_NSSMF = 2
)

// Enum value maps for Kpi_NSSMF.
var (
	Kpi_NSSMF_name = map[int32]string{
		0: "NFVO",
		1: "RAN",
		2: "SO",
	}
	Kpi_NSSMF_value = map[string]int32{
		"NFVO": 0,
		"RAN":  1,
		"SO":   2,
	}
)

func (x Kpi_NSSMF) Enum() *Kpi_NSSMF {
	p := new(Kpi_NSSMF)
	*p = x
	return p
}

func (x Kpi_NSSMF) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Kpi_NSSMF) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[0].Descriptor()
}

func (Kpi_NSSMF) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[0]
}

func (x Kpi_NSSMF) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Kpi_NSSMF.Descriptor instead.
func (Kpi_NSSMF) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2, 0}
}

type Parameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Parameter) Reset() {
	*x = Parameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Parameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Parameter) ProtoMessage() {}

func (x *Parameter) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Parameter.ProtoReflect.Descriptor instead.
func (*Parameter) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *Parameter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Parameter) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value  float64      `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	Unit   string       `protobuf:"bytes,2,opt,name=unit,proto3" json:"unit,omitempty"`
	Params []*Parameter `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty"`
}

func (x *Payload) Reset() {
	*x = Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload) ProtoMessage() {}

func (x *Payload) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payload.ProtoReflect.Descriptor instead.
func (*Payload) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *Payload) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Payload) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *Payload) GetParams() []*Parameter {
	if x != nil {
		return x.Params
	}
	return nil
}

type Kpi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nssmf     Kpi_NSSMF `protobuf:"varint,1,opt,name=nssmf,proto3,enum=Kpi_NSSMF" json:"nssmf,omitempty"`
	Id        int64     `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Region    string    `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	Timestamp float64   `protobuf:"fixed64,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Nssid     int64     `protobuf:"varint,6,opt,name=nssid,proto3" json:"nssid,omitempty"`
	Metric    string    `protobuf:"bytes,7,opt,name=metric,proto3" json:"metric,omitempty"`
	Unit      string    `protobuf:"bytes,8,opt,name=unit,proto3" json:"unit,omitempty"`
	Payload   *Payload  `protobuf:"bytes,9,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Kpi) Reset() {
	*x = Kpi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Kpi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kpi) ProtoMessage() {}

func (x *Kpi) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kpi.ProtoReflect.Descriptor instead.
func (*Kpi) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *Kpi) GetNssmf() Kpi_NSSMF {
	if x != nil {
		return x.Nssmf
	}
	return Kpi_NFVO
}

func (x *Kpi) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Kpi) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Kpi) GetTimestamp() float64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Kpi) GetNssid() int64 {
	if x != nil {
		return x.Nssid
	}
	return 0
}

func (x *Kpi) GetMetric() string {
	if x != nil {
		return x.Metric
	}
	return ""
}

func (x *Kpi) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *Kpi) GetPayload() *Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x35, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x57, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x22, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22,
	0xf7, 0x01, 0x0a, 0x03, 0x4b, 0x70, 0x69, 0x12, 0x20, 0x0a, 0x05, 0x6e, 0x73, 0x73, 0x6d, 0x66,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x4b, 0x70, 0x69, 0x2e, 0x4e, 0x53, 0x53,
	0x4d, 0x46, 0x52, 0x05, 0x6e, 0x73, 0x73, 0x6d, 0x66, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x14, 0x0a, 0x05, 0x6e, 0x73, 0x73, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x6e, 0x73, 0x73, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69,
	0x74, 0x12, 0x22, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x22, 0x0a, 0x05, 0x4e, 0x53, 0x53, 0x4d, 0x46, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x46, 0x56, 0x4f, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x41, 0x4e, 0x10,
	0x01, 0x12, 0x06, 0x0a, 0x02, 0x53, 0x4f, 0x10, 0x02, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x62, 0x64, 0x6f, 0x75, 0x54, 0x6c, 0x69,
	0x6c, 0x69, 0x2f, 0x6d, 0x65, 0x74, 0x2d, 0x78, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x6e, 0x6f, 0x72, 0x74, 0x68, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_message_proto_goTypes = []interface{}{
	(Kpi_NSSMF)(0),    // 0: Kpi.NSSMF
	(*Parameter)(nil), // 1: Parameter
	(*Payload)(nil),   // 2: Payload
	(*Kpi)(nil),       // 3: Kpi
}
var file_message_proto_depIdxs = []int32{
	1, // 0: Payload.params:type_name -> Parameter
	0, // 1: Kpi.nssmf:type_name -> Kpi.NSSMF
	2, // 2: Kpi.payload:type_name -> Payload
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Parameter); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payload); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Kpi); i {
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
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		EnumInfos:         file_message_proto_enumTypes,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
