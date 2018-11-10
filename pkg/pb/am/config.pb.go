// Code generated by protoc-gen-go. DO NOT EDIT.
// source: iam/am/config.proto

package pbam // import "openpitrix.io/iam/pkg/pb/am"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Config struct {
	LogLevel             *string           `protobuf:"bytes,1,opt,name=log_level,json=logLevel" json:"log_level,omitempty"`
	Mysql                *MysqlConfig      `protobuf:"bytes,2,opt,name=mysql" json:"mysql,omitempty"`
	Extra                map[string]string `protobuf:"bytes,100,rep,name=extra" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_b1d38c7ea9e2bd7c, []int{0}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (dst *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(dst, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetLogLevel() string {
	if m != nil && m.LogLevel != nil {
		return *m.LogLevel
	}
	return ""
}

func (m *Config) GetMysql() *MysqlConfig {
	if m != nil {
		return m.Mysql
	}
	return nil
}

func (m *Config) GetExtra() map[string]string {
	if m != nil {
		return m.Extra
	}
	return nil
}

type MysqlConfig struct {
	Host                 *string  `protobuf:"bytes,1,opt,name=host,def=openpitrix-db" json:"host,omitempty"`
	Port                 *int32   `protobuf:"varint,2,opt,name=port,def=3306" json:"port,omitempty"`
	User                 *string  `protobuf:"bytes,3,opt,name=user,def=root" json:"user,omitempty"`
	Password             *string  `protobuf:"bytes,4,opt,name=password,def=password" json:"password,omitempty"`
	Database             *string  `protobuf:"bytes,5,opt,name=database,def=openpitrix" json:"database,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MysqlConfig) Reset()         { *m = MysqlConfig{} }
func (m *MysqlConfig) String() string { return proto.CompactTextString(m) }
func (*MysqlConfig) ProtoMessage()    {}
func (*MysqlConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_b1d38c7ea9e2bd7c, []int{1}
}
func (m *MysqlConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MysqlConfig.Unmarshal(m, b)
}
func (m *MysqlConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MysqlConfig.Marshal(b, m, deterministic)
}
func (dst *MysqlConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MysqlConfig.Merge(dst, src)
}
func (m *MysqlConfig) XXX_Size() int {
	return xxx_messageInfo_MysqlConfig.Size(m)
}
func (m *MysqlConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MysqlConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MysqlConfig proto.InternalMessageInfo

const Default_MysqlConfig_Host string = "openpitrix-db"
const Default_MysqlConfig_Port int32 = 3306
const Default_MysqlConfig_User string = "root"
const Default_MysqlConfig_Password string = "password"
const Default_MysqlConfig_Database string = "openpitrix"

func (m *MysqlConfig) GetHost() string {
	if m != nil && m.Host != nil {
		return *m.Host
	}
	return Default_MysqlConfig_Host
}

func (m *MysqlConfig) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return Default_MysqlConfig_Port
}

func (m *MysqlConfig) GetUser() string {
	if m != nil && m.User != nil {
		return *m.User
	}
	return Default_MysqlConfig_User
}

func (m *MysqlConfig) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return Default_MysqlConfig_Password
}

func (m *MysqlConfig) GetDatabase() string {
	if m != nil && m.Database != nil {
		return *m.Database
	}
	return Default_MysqlConfig_Database
}

func init() {
	proto.RegisterType((*Config)(nil), "iam.am.Config")
	proto.RegisterMapType((map[string]string)(nil), "iam.am.Config.ExtraEntry")
	proto.RegisterType((*MysqlConfig)(nil), "iam.am.MysqlConfig")
}

func init() { proto.RegisterFile("iam/am/config.proto", fileDescriptor_config_b1d38c7ea9e2bd7c) }

var fileDescriptor_config_b1d38c7ea9e2bd7c = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xcf, 0x4e, 0x2a, 0x31,
	0x14, 0x87, 0x33, 0x30, 0x43, 0xe0, 0x90, 0x9b, 0xdc, 0x94, 0xbb, 0x98, 0x8b, 0x0b, 0x47, 0x62,
	0x0c, 0x2e, 0x9c, 0x1a, 0x48, 0x0c, 0xc1, 0x9d, 0x86, 0x9d, 0x6e, 0xba, 0x74, 0x63, 0x3a, 0x4e,
	0x19, 0x27, 0xb4, 0x9c, 0xda, 0x16, 0x84, 0xd7, 0x72, 0xeb, 0xcb, 0x99, 0x69, 0x15, 0xdc, 0xf5,
	0x7c, 0xbf, 0xef, 0xfc, 0x49, 0x0a, 0x83, 0x9a, 0x2b, 0xca, 0x15, 0x7d, 0xc1, 0xf5, 0xb2, 0xae,
	0x72, 0x6d, 0xd0, 0x21, 0xe9, 0xd4, 0x5c, 0xe5, 0x5c, 0x0d, 0x4f, 0x2b, 0xc4, 0x4a, 0x0a, 0xea,
	0x69, 0xb1, 0x59, 0x52, 0x57, 0x2b, 0x61, 0x1d, 0x57, 0x3a, 0x88, 0xa3, 0xcf, 0x08, 0x3a, 0xf7,
	0xbe, 0x93, 0x9c, 0x40, 0x4f, 0x62, 0xf5, 0x2c, 0xc5, 0x56, 0xc8, 0x34, 0xca, 0xa2, 0x71, 0x8f,
	0x75, 0x25, 0x56, 0x0f, 0x4d, 0x4d, 0x2e, 0x21, 0x51, 0x7b, 0xfb, 0x26, 0xd3, 0x56, 0x16, 0x8d,
	0xfb, 0x93, 0x41, 0x1e, 0x16, 0xe4, 0x8f, 0x0d, 0x0c, 0x03, 0x58, 0x30, 0x08, 0x85, 0x44, 0xec,
	0x9c, 0xe1, 0x69, 0x99, 0xb5, 0xc7, 0xfd, 0xc9, 0xff, 0x1f, 0x35, 0x58, 0xf9, 0xa2, 0xc9, 0x16,
	0x6b, 0x67, 0xf6, 0x2c, 0x78, 0xc3, 0x19, 0xc0, 0x11, 0x92, 0xbf, 0xd0, 0x5e, 0x89, 0xfd, 0xf7,
	0x01, 0xcd, 0x93, 0xfc, 0x83, 0x64, 0xcb, 0xe5, 0x46, 0xf8, 0xdd, 0x3d, 0x16, 0x8a, 0x79, 0x6b,
	0x16, 0x8d, 0x3e, 0x22, 0xe8, 0xff, 0xba, 0x80, 0x9c, 0x41, 0xfc, 0x8a, 0xd6, 0x85, 0xe6, 0xf9,
	0x1f, 0xd4, 0x62, 0xad, 0x6b, 0x67, 0xea, 0xdd, 0x55, 0x59, 0x30, 0x1f, 0x91, 0x14, 0x62, 0x8d,
	0xc6, 0xf9, 0x59, 0xc9, 0x3c, 0x9e, 0x4e, 0xaf, 0x6f, 0x98, 0x27, 0x4d, 0xb2, 0xb1, 0xc2, 0xa4,
	0x6d, 0xdf, 0x1c, 0x1b, 0x44, 0xc7, 0x3c, 0x21, 0xe7, 0xd0, 0xd5, 0xdc, 0xda, 0x77, 0x34, 0x65,
	0x1a, 0xfb, 0xf4, 0x50, 0xb3, 0xc3, 0x8b, 0x5c, 0x40, 0xb7, 0xe4, 0x8e, 0x17, 0xdc, 0x8a, 0x34,
	0xf1, 0x16, 0x1c, 0x0f, 0x60, 0x87, 0xec, 0x6e, 0xf4, 0x94, 0x1d, 0x79, 0x5e, 0x23, 0x6d, 0x3e,
	0x50, 0xaf, 0x2a, 0xaa, 0x0b, 0xca, 0xd5, 0xad, 0x2e, 0xb8, 0xfa, 0x0a, 0x00, 0x00, 0xff, 0xff,
	0x8f, 0xac, 0xf4, 0x57, 0xd5, 0x01, 0x00, 0x00,
}
