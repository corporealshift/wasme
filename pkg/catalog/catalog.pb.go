// Code generated by protoc-gen-go. DO NOT EDIT.
// source: catalog.proto

package catalog

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ExtensionSpec struct {
	// The name of the extension. Must be unique within a repository.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Module logo URL. optional
	LogoUrl string `protobuf:"bytes,2,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
	// Short description of the extension.
	ShortDescription string `protobuf:"bytes,3,opt,name=short_description,json=shortDescription,proto3" json:"short_description,omitempty"`
	// Long description of the extension.
	// TODO support colocated markdown file OR link to markdown file
	LongDescription string `protobuf:"bytes,4,opt,name=long_description,json=longDescription,proto3" json:"long_description,omitempty"`
	// Link to documentation for the extension.
	DocumentationUrl string `protobuf:"bytes,5,opt,name=documentation_url,json=documentationUrl,proto3" json:"documentation_url,omitempty"`
	// Link to repository containing the extension's source code.
	RepositoryUrl string `protobuf:"bytes,6,opt,name=repository_url,json=repositoryUrl,proto3" json:"repository_url,omitempty"`
	// Link to repository containing the extension's source code.
	ExtensionRef string `protobuf:"bytes,7,opt,name=extension_ref,json=extensionRef,proto3" json:"extension_ref,omitempty"`
	// The name of the entity that authored the extension.
	CreatorName string `protobuf:"bytes,8,opt,name=creator_name,json=creatorName,proto3" json:"creator_name,omitempty"`
	// Link to the entity that authored the extension.
	CreatorUrl string `protobuf:"bytes,9,opt,name=creator_url,json=creatorUrl,proto3" json:"creator_url,omitempty"`
	// Install instructions for this extension.
	InstallInstructions  *InstallInstructions `protobuf:"bytes,10,opt,name=install_instructions,json=installInstructions,proto3" json:"install_instructions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ExtensionSpec) Reset()         { *m = ExtensionSpec{} }
func (m *ExtensionSpec) String() string { return proto.CompactTextString(m) }
func (*ExtensionSpec) ProtoMessage()    {}
func (*ExtensionSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{0}
}

func (m *ExtensionSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtensionSpec.Unmarshal(m, b)
}
func (m *ExtensionSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtensionSpec.Marshal(b, m, deterministic)
}
func (m *ExtensionSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtensionSpec.Merge(m, src)
}
func (m *ExtensionSpec) XXX_Size() int {
	return xxx_messageInfo_ExtensionSpec.Size(m)
}
func (m *ExtensionSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtensionSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ExtensionSpec proto.InternalMessageInfo

func (m *ExtensionSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ExtensionSpec) GetLogoUrl() string {
	if m != nil {
		return m.LogoUrl
	}
	return ""
}

func (m *ExtensionSpec) GetShortDescription() string {
	if m != nil {
		return m.ShortDescription
	}
	return ""
}

func (m *ExtensionSpec) GetLongDescription() string {
	if m != nil {
		return m.LongDescription
	}
	return ""
}

func (m *ExtensionSpec) GetDocumentationUrl() string {
	if m != nil {
		return m.DocumentationUrl
	}
	return ""
}

func (m *ExtensionSpec) GetRepositoryUrl() string {
	if m != nil {
		return m.RepositoryUrl
	}
	return ""
}

func (m *ExtensionSpec) GetExtensionRef() string {
	if m != nil {
		return m.ExtensionRef
	}
	return ""
}

func (m *ExtensionSpec) GetCreatorName() string {
	if m != nil {
		return m.CreatorName
	}
	return ""
}

func (m *ExtensionSpec) GetCreatorUrl() string {
	if m != nil {
		return m.CreatorUrl
	}
	return ""
}

func (m *ExtensionSpec) GetInstallInstructions() *InstallInstructions {
	if m != nil {
		return m.InstallInstructions
	}
	return nil
}

// Set of install instructions for each supported envoy flavor
type InstallInstructions struct {
	Gloo                 *InstallInstructions_Gloo         `protobuf:"bytes,1,opt,name=gloo,proto3" json:"gloo,omitempty"`
	VanillaEnvoy         *InstallInstructions_VanillaEnvoy `protobuf:"bytes,3,opt,name=vanilla_envoy,json=vanillaEnvoy,proto3" json:"vanilla_envoy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *InstallInstructions) Reset()         { *m = InstallInstructions{} }
func (m *InstallInstructions) String() string { return proto.CompactTextString(m) }
func (*InstallInstructions) ProtoMessage()    {}
func (*InstallInstructions) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{1}
}

func (m *InstallInstructions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallInstructions.Unmarshal(m, b)
}
func (m *InstallInstructions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallInstructions.Marshal(b, m, deterministic)
}
func (m *InstallInstructions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallInstructions.Merge(m, src)
}
func (m *InstallInstructions) XXX_Size() int {
	return xxx_messageInfo_InstallInstructions.Size(m)
}
func (m *InstallInstructions) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallInstructions.DiscardUnknown(m)
}

var xxx_messageInfo_InstallInstructions proto.InternalMessageInfo

func (m *InstallInstructions) GetGloo() *InstallInstructions_Gloo {
	if m != nil {
		return m.Gloo
	}
	return nil
}

func (m *InstallInstructions) GetVanillaEnvoy() *InstallInstructions_VanillaEnvoy {
	if m != nil {
		return m.VanillaEnvoy
	}
	return nil
}

type InstallInstructions_Gloo struct {
	Glooctl              string   `protobuf:"bytes,1,opt,name=glooctl,proto3" json:"glooctl,omitempty"`
	GatewayYaml          string   `protobuf:"bytes,2,opt,name=gateway_yaml,json=gatewayYaml,proto3" json:"gateway_yaml,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstallInstructions_Gloo) Reset()         { *m = InstallInstructions_Gloo{} }
func (m *InstallInstructions_Gloo) String() string { return proto.CompactTextString(m) }
func (*InstallInstructions_Gloo) ProtoMessage()    {}
func (*InstallInstructions_Gloo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{1, 0}
}

func (m *InstallInstructions_Gloo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallInstructions_Gloo.Unmarshal(m, b)
}
func (m *InstallInstructions_Gloo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallInstructions_Gloo.Marshal(b, m, deterministic)
}
func (m *InstallInstructions_Gloo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallInstructions_Gloo.Merge(m, src)
}
func (m *InstallInstructions_Gloo) XXX_Size() int {
	return xxx_messageInfo_InstallInstructions_Gloo.Size(m)
}
func (m *InstallInstructions_Gloo) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallInstructions_Gloo.DiscardUnknown(m)
}

var xxx_messageInfo_InstallInstructions_Gloo proto.InternalMessageInfo

func (m *InstallInstructions_Gloo) GetGlooctl() string {
	if m != nil {
		return m.Glooctl
	}
	return ""
}

func (m *InstallInstructions_Gloo) GetGatewayYaml() string {
	if m != nil {
		return m.GatewayYaml
	}
	return ""
}

type InstallInstructions_VanillaEnvoy struct {
	Config               string   `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstallInstructions_VanillaEnvoy) Reset()         { *m = InstallInstructions_VanillaEnvoy{} }
func (m *InstallInstructions_VanillaEnvoy) String() string { return proto.CompactTextString(m) }
func (*InstallInstructions_VanillaEnvoy) ProtoMessage()    {}
func (*InstallInstructions_VanillaEnvoy) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{1, 1}
}

func (m *InstallInstructions_VanillaEnvoy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallInstructions_VanillaEnvoy.Unmarshal(m, b)
}
func (m *InstallInstructions_VanillaEnvoy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallInstructions_VanillaEnvoy.Marshal(b, m, deterministic)
}
func (m *InstallInstructions_VanillaEnvoy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallInstructions_VanillaEnvoy.Merge(m, src)
}
func (m *InstallInstructions_VanillaEnvoy) XXX_Size() int {
	return xxx_messageInfo_InstallInstructions_VanillaEnvoy.Size(m)
}
func (m *InstallInstructions_VanillaEnvoy) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallInstructions_VanillaEnvoy.DiscardUnknown(m)
}

var xxx_messageInfo_InstallInstructions_VanillaEnvoy proto.InternalMessageInfo

func (m *InstallInstructions_VanillaEnvoy) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

func init() {
	proto.RegisterType((*ExtensionSpec)(nil), "ExtensionSpec")
	proto.RegisterType((*InstallInstructions)(nil), "InstallInstructions")
	proto.RegisterType((*InstallInstructions_Gloo)(nil), "InstallInstructions.Gloo")
	proto.RegisterType((*InstallInstructions_VanillaEnvoy)(nil), "InstallInstructions.VanillaEnvoy")
}

func init() { proto.RegisterFile("catalog.proto", fileDescriptor_0abbfcf058acdf89) }

var fileDescriptor_0abbfcf058acdf89 = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0x80, 0x15, 0x16, 0x9a, 0xf5, 0x92, 0xc0, 0xe6, 0x4d, 0x28, 0xdb, 0x0b, 0xdb, 0x10, 0x68,
	0xa8, 0xa2, 0x0f, 0xe5, 0x1f, 0x00, 0xa5, 0xea, 0x0b, 0x0f, 0x41, 0x20, 0xc1, 0x4b, 0x64, 0x5c,
	0x37, 0x58, 0x72, 0x7c, 0x91, 0xe3, 0x16, 0xf2, 0x5f, 0xf8, 0xa3, 0xbc, 0x21, 0x5f, 0x13, 0x9a,
	0x6a, 0x7d, 0x4a, 0xee, 0xbb, 0x2f, 0x77, 0xce, 0xf9, 0x20, 0x15, 0xdc, 0x71, 0x8d, 0xe5, 0xb4,
	0xb6, 0xe8, 0xf0, 0xee, 0xcf, 0x09, 0xa4, 0xf3, 0xdf, 0x4e, 0x9a, 0x46, 0xa1, 0xf9, 0x5c, 0x4b,
	0xc1, 0x18, 0x84, 0x86, 0x57, 0x32, 0x0b, 0x6e, 0x82, 0xfb, 0x71, 0x4e, 0xef, 0xec, 0x0a, 0x4e,
	0x35, 0x96, 0x58, 0x6c, 0xac, 0xce, 0x1e, 0x11, 0x8f, 0x7c, 0xfc, 0xc5, 0x6a, 0x36, 0x81, 0xf3,
	0xe6, 0x27, 0x5a, 0x57, 0xac, 0x64, 0x23, 0xac, 0xaa, 0x9d, 0x42, 0x93, 0x9d, 0x90, 0x73, 0x46,
	0x89, 0x0f, 0x7b, 0xce, 0x5e, 0xc3, 0x99, 0x46, 0x53, 0x1e, 0xb8, 0x21, 0xb9, 0x4f, 0x3d, 0x1f,
	0xaa, 0x13, 0x38, 0x5f, 0xa1, 0xd8, 0x54, 0xd2, 0x38, 0xee, 0x01, 0xf5, 0x7e, 0xbc, 0xab, 0x7b,
	0x90, 0xf0, 0x87, 0x78, 0x09, 0x4f, 0xac, 0xac, 0xb1, 0x51, 0x0e, 0x6d, 0x4b, 0xe6, 0x88, 0xcc,
	0x74, 0x4f, 0xbd, 0xf6, 0x02, 0x52, 0xd9, 0xff, 0x6b, 0x61, 0xe5, 0x3a, 0x8b, 0xc8, 0x4a, 0xfe,
	0xc3, 0x5c, 0xae, 0xd9, 0x2d, 0x24, 0xc2, 0x4a, 0xee, 0xd0, 0x16, 0x34, 0x87, 0x53, 0x72, 0xe2,
	0x8e, 0x7d, 0xf2, 0xe3, 0x78, 0x0e, 0x7d, 0x48, 0xbd, 0xc6, 0x64, 0x40, 0x87, 0x7c, 0xa3, 0x05,
	0x5c, 0x2a, 0xd3, 0x38, 0xae, 0x75, 0xe1, 0x9f, 0x76, 0x23, 0xfc, 0x49, 0x9b, 0x0c, 0x6e, 0x82,
	0xfb, 0x78, 0x76, 0x39, 0x5d, 0xee, 0x92, 0xcb, 0x41, 0x2e, 0xbf, 0x50, 0x0f, 0xe1, 0xdd, 0xdf,
	0x00, 0x2e, 0x8e, 0xc8, 0xec, 0x0d, 0x84, 0xa5, 0x46, 0xa4, 0x4b, 0x8a, 0x67, 0x57, 0xc7, 0x0a,
	0x4e, 0x17, 0x1a, 0x31, 0x27, 0x8d, 0x7d, 0x84, 0x74, 0xcb, 0x8d, 0xd2, 0x9a, 0x17, 0xd2, 0x6c,
	0xb1, 0xa5, 0x0b, 0x8a, 0x67, 0xb7, 0x47, 0xbf, 0xfb, 0xba, 0x33, 0xe7, 0x5e, 0xcc, 0x93, 0xed,
	0x20, 0xba, 0x7e, 0x0f, 0xa1, 0xaf, 0xca, 0x32, 0x88, 0x7c, 0x5d, 0xe1, 0x74, 0xb7, 0x26, 0x7d,
	0xe8, 0xa7, 0x57, 0x72, 0x27, 0x7f, 0xf1, 0xb6, 0x68, 0x79, 0xd5, 0x6f, 0x4b, 0xdc, 0xb1, 0x6f,
	0xbc, 0xd2, 0xd7, 0xaf, 0x20, 0x19, 0xb6, 0x60, 0xcf, 0x60, 0x24, 0xd0, 0xac, 0x55, 0xd9, 0xd5,
	0xea, 0xa2, 0x77, 0xe3, 0xef, 0x51, 0xb7, 0xab, 0x3f, 0x46, 0xb4, 0xac, 0x6f, 0xff, 0x05, 0x00,
	0x00, 0xff, 0xff, 0xd7, 0x76, 0x9b, 0x4f, 0xbd, 0x02, 0x00, 0x00,
}