package persist

import (
	_ "github.com/v2fly/v2ray-core/v5/common/protoext"
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

type DBconfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Path   string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Uname  string `protobuf:"bytes,3,opt,name=uname,proto3" json:"uname,omitempty"`
	Passwd string `protobuf:"bytes,4,opt,name=passwd,proto3" json:"passwd,omitempty"`
	Dbname string `protobuf:"bytes,5,opt,name=dbname,proto3" json:"dbname,omitempty"`
}

func (x *DBconfig) Reset() {
	*x = DBconfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_persist_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBconfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBconfig) ProtoMessage() {}

func (x *DBconfig) ProtoReflect() protoreflect.Message {
	mi := &file_app_persist_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBconfig.ProtoReflect.Descriptor instead.
func (*DBconfig) Descriptor() ([]byte, []int) {
	return file_app_persist_config_proto_rawDescGZIP(), []int{0}
}

func (x *DBconfig) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DBconfig) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *DBconfig) GetUname() string {
	if x != nil {
		return x.Uname
	}
	return ""
}

func (x *DBconfig) GetPasswd() string {
	if x != nil {
		return x.Passwd
	}
	return ""
}

func (x *DBconfig) GetDbname() string {
	if x != nil {
		return x.Dbname
	}
	return ""
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database *DBconfig `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Enable   bool      `protobuf:"varint,2,opt,name=enable,proto3" json:"enable,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_persist_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_app_persist_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_app_persist_config_proto_rawDescGZIP(), []int{1}
}

func (x *Config) GetDatabase() *DBconfig {
	if x != nil {
		return x.Database
	}
	return nil
}

func (x *Config) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

var File_app_persist_config_proto protoreflect.FileDescriptor

var file_app_persist_config_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x76, 0x32, 0x72, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x69,
	0x73, 0x74, 0x1a, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x65, 0x78, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x78, 0x0a, 0x08, 0x44, 0x42, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x75, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x61, 0x73, 0x73, 0x77, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x62, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x76,
	0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3c, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x76, 0x32, 0x72,
	0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x65, 0x72, 0x73,
	0x69, 0x73, 0x74, 0x2e, 0x44, 0x42, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x08, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x3a, 0x16,
	0x82, 0xb5, 0x18, 0x12, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x07, 0x70,
	0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x42, 0x63, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x32,
	0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x65, 0x72,
	0x73, 0x69, 0x73, 0x74, 0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x76, 0x32, 0x66, 0x6c, 0x79, 0x2f, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2d, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x35, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x69,
	0x73, 0x74, 0xaa, 0x02, 0x16, 0x56, 0x32, 0x52, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e,
	0x41, 0x70, 0x70, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_app_persist_config_proto_rawDescOnce sync.Once
	file_app_persist_config_proto_rawDescData = file_app_persist_config_proto_rawDesc
)

func file_app_persist_config_proto_rawDescGZIP() []byte {
	file_app_persist_config_proto_rawDescOnce.Do(func() {
		file_app_persist_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_persist_config_proto_rawDescData)
	})
	return file_app_persist_config_proto_rawDescData
}

var file_app_persist_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_app_persist_config_proto_goTypes = []interface{}{
	(*DBconfig)(nil), // 0: v2ray.core.app.persist.DBconfig
	(*Config)(nil),   // 1: v2ray.core.app.persist.Config
}
var file_app_persist_config_proto_depIdxs = []int32{
	0, // 0: v2ray.core.app.persist.Config.database:type_name -> v2ray.core.app.persist.DBconfig
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_app_persist_config_proto_init() }
func file_app_persist_config_proto_init() {
	if File_app_persist_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_persist_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBconfig); i {
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
		file_app_persist_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_app_persist_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_persist_config_proto_goTypes,
		DependencyIndexes: file_app_persist_config_proto_depIdxs,
		MessageInfos:      file_app_persist_config_proto_msgTypes,
	}.Build()
	File_app_persist_config_proto = out.File
	file_app_persist_config_proto_rawDesc = nil
	file_app_persist_config_proto_goTypes = nil
	file_app_persist_config_proto_depIdxs = nil
}
