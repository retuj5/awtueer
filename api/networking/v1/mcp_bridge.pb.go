// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networking/v1/mcp_bridge.proto

package v1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// <!-- crd generation tags
// +cue-gen:McpBridge:groupName:networking.higress.io
// +cue-gen:McpBridge:version:v1
// +cue-gen:McpBridge:storageVersion
// +cue-gen:McpBridge:annotations:helm.sh/resource-policy=keep
// +cue-gen:McpBridge:subresource:status
// +cue-gen:McpBridge:scope:Namespaced
// +cue-gen:McpBridge:resource:categories=higress-io,plural=mcpbridges
// +cue-gen:McpBridge:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=networking.higress.io/v1
// +genclient
// +k8s:deepcopy-gen=true
// -->
type McpBridge struct {
	Registries           []*RegistryConfig `protobuf:"bytes,1,rep,name=registries,proto3" json:"registries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *McpBridge) Reset()         { *m = McpBridge{} }
func (m *McpBridge) String() string { return proto.CompactTextString(m) }
func (*McpBridge) ProtoMessage()    {}
func (*McpBridge) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fcc59a15c34642d, []int{0}
}
func (m *McpBridge) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *McpBridge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_McpBridge.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *McpBridge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_McpBridge.Merge(m, src)
}
func (m *McpBridge) XXX_Size() int {
	return m.Size()
}
func (m *McpBridge) XXX_DiscardUnknown() {
	xxx_messageInfo_McpBridge.DiscardUnknown(m)
}

var xxx_messageInfo_McpBridge proto.InternalMessageInfo

func (m *McpBridge) GetRegistries() []*RegistryConfig {
	if m != nil {
		return m.Registries
	}
	return nil
}

type RegistryConfig struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Domain               string   `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain,omitempty"`
	Port                 uint32   `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	NacosAddressServer   string   `protobuf:"bytes,5,opt,name=nacosAddressServer,proto3" json:"nacosAddressServer,omitempty"`
	NacosAccessKey       string   `protobuf:"bytes,6,opt,name=nacosAccessKey,proto3" json:"nacosAccessKey,omitempty"`
	NacosSecretKey       string   `protobuf:"bytes,7,opt,name=nacosSecretKey,proto3" json:"nacosSecretKey,omitempty"`
	NacosNamespaceId     string   `protobuf:"bytes,8,opt,name=nacosNamespaceId,proto3" json:"nacosNamespaceId,omitempty"`
	NacosNamespace       string   `protobuf:"bytes,9,opt,name=nacosNamespace,proto3" json:"nacosNamespace,omitempty"`
	NacosGroups          []string `protobuf:"bytes,10,rep,name=nacosGroups,proto3" json:"nacosGroups,omitempty"`
	NacosRefreshInterval int64    `protobuf:"varint,11,opt,name=nacosRefreshInterval,proto3" json:"nacosRefreshInterval,omitempty"`
	ConsulNamespace      string   `protobuf:"bytes,12,opt,name=consulNamespace,proto3" json:"consulNamespace,omitempty"`
	ZkServicesPath       []string `protobuf:"bytes,13,rep,name=zkServicesPath,proto3" json:"zkServicesPath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistryConfig) Reset()         { *m = RegistryConfig{} }
func (m *RegistryConfig) String() string { return proto.CompactTextString(m) }
func (*RegistryConfig) ProtoMessage()    {}
func (*RegistryConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fcc59a15c34642d, []int{1}
}
func (m *RegistryConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegistryConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegistryConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegistryConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistryConfig.Merge(m, src)
}
func (m *RegistryConfig) XXX_Size() int {
	return m.Size()
}
func (m *RegistryConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistryConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RegistryConfig proto.InternalMessageInfo

func (m *RegistryConfig) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *RegistryConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegistryConfig) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RegistryConfig) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *RegistryConfig) GetNacosAddressServer() string {
	if m != nil {
		return m.NacosAddressServer
	}
	return ""
}

func (m *RegistryConfig) GetNacosAccessKey() string {
	if m != nil {
		return m.NacosAccessKey
	}
	return ""
}

func (m *RegistryConfig) GetNacosSecretKey() string {
	if m != nil {
		return m.NacosSecretKey
	}
	return ""
}

func (m *RegistryConfig) GetNacosNamespaceId() string {
	if m != nil {
		return m.NacosNamespaceId
	}
	return ""
}

func (m *RegistryConfig) GetNacosNamespace() string {
	if m != nil {
		return m.NacosNamespace
	}
	return ""
}

func (m *RegistryConfig) GetNacosGroups() []string {
	if m != nil {
		return m.NacosGroups
	}
	return nil
}

func (m *RegistryConfig) GetNacosRefreshInterval() int64 {
	if m != nil {
		return m.NacosRefreshInterval
	}
	return 0
}

func (m *RegistryConfig) GetConsulNamespace() string {
	if m != nil {
		return m.ConsulNamespace
	}
	return ""
}

func (m *RegistryConfig) GetZkServicesPath() []string {
	if m != nil {
		return m.ZkServicesPath
	}
	return nil
}

func init() {
	proto.RegisterType((*McpBridge)(nil), "higress.networking.v1.McpBridge")
	proto.RegisterType((*RegistryConfig)(nil), "higress.networking.v1.RegistryConfig")
}

func init() { proto.RegisterFile("networking/v1/mcp_bridge.proto", fileDescriptor_3fcc59a15c34642d) }

var fileDescriptor_3fcc59a15c34642d = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xd1, 0x8a, 0xd4, 0x30,
	0x14, 0x86, 0xc9, 0x76, 0x1c, 0x9d, 0x8c, 0xbb, 0x4a, 0x50, 0x08, 0x22, 0x63, 0x59, 0x50, 0x8a,
	0x48, 0xcb, 0xae, 0x77, 0xde, 0xed, 0x88, 0xc8, 0x22, 0x8a, 0x74, 0xef, 0xbc, 0x59, 0xd2, 0xf4,
	0x4c, 0x1a, 0xb6, 0x4d, 0x42, 0x92, 0xa9, 0x8c, 0x4f, 0xe8, 0xa5, 0x8f, 0x20, 0x7d, 0x04, 0x9f,
	0x40, 0x9a, 0x59, 0xbb, 0x9d, 0x71, 0xee, 0xda, 0xef, 0x7c, 0xf9, 0x73, 0x08, 0x3f, 0x5e, 0x28,
	0xf0, 0xdf, 0xb5, 0xbd, 0x91, 0x4a, 0x64, 0xed, 0x59, 0xd6, 0x70, 0x73, 0x5d, 0x58, 0x59, 0x0a,
	0x48, 0x8d, 0xd5, 0x5e, 0x93, 0xa7, 0x95, 0x14, 0x16, 0x9c, 0x4b, 0xef, 0xbc, 0xb4, 0x3d, 0x7b,
	0xf6, 0x42, 0x68, 0x2d, 0x6a, 0xc8, 0x98, 0x91, 0xd9, 0x4a, 0x42, 0x5d, 0x5e, 0x17, 0x50, 0xb1,
	0x56, 0x6a, 0xbb, 0x3d, 0x77, 0x9a, 0xe3, 0xd9, 0x67, 0x6e, 0x96, 0x21, 0x8a, 0x7c, 0xc0, 0xd8,
	0x82, 0x90, 0xce, 0x5b, 0x09, 0x8e, 0xa2, 0x38, 0x4a, 0xe6, 0xe7, 0x2f, 0xd3, 0x83, 0xc9, 0x69,
	0xbe, 0x15, 0x37, 0xef, 0xb5, 0x5a, 0x49, 0x91, 0x8f, 0x0e, 0x9e, 0xfe, 0x89, 0xf0, 0xc9, 0xee,
	0x98, 0x50, 0x3c, 0xf1, 0x1b, 0x03, 0x14, 0xc5, 0x28, 0x99, 0x2d, 0x27, 0xdd, 0x05, 0x3a, 0xca,
	0x03, 0x21, 0x04, 0x4f, 0x14, 0x6b, 0x80, 0x1e, 0xf5, 0x93, 0x3c, 0x7c, 0x93, 0xe7, 0x78, 0x5a,
	0xea, 0x86, 0x49, 0x45, 0xa3, 0x91, 0x7f, 0xcb, 0xfa, 0x2c, 0xa3, 0xad, 0xa7, 0x93, 0x18, 0x25,
	0xc7, 0xff, 0xb2, 0x7a, 0x42, 0x52, 0x4c, 0x14, 0xe3, 0xda, 0x5d, 0x94, 0x65, 0xbf, 0xf1, 0x15,
	0xd8, 0x16, 0x2c, 0xbd, 0x17, 0x92, 0x0f, 0x4c, 0xc8, 0x2b, 0x7c, 0xb2, 0xa5, 0x9c, 0x83, 0x73,
	0x9f, 0x60, 0x43, 0xa7, 0xc1, 0xdd, 0xa3, 0x83, 0x77, 0x05, 0xdc, 0x82, 0xef, 0xbd, 0xfb, 0x23,
	0x6f, 0xa0, 0xe4, 0x35, 0x7e, 0x1c, 0xc8, 0x17, 0xd6, 0x80, 0x33, 0x8c, 0xc3, 0x65, 0x49, 0x1f,
	0x04, 0xf3, 0x3f, 0x3e, 0x64, 0x0e, 0x8c, 0xce, 0x46, 0x99, 0x03, 0x25, 0x31, 0x9e, 0x07, 0xf2,
	0xd1, 0xea, 0xb5, 0x71, 0x14, 0xc7, 0x51, 0x32, 0xcb, 0xc7, 0x88, 0x9c, 0xe3, 0x27, 0xe1, 0x37,
	0x87, 0x95, 0x05, 0x57, 0x5d, 0x2a, 0x0f, 0xb6, 0x65, 0x35, 0x9d, 0xc7, 0x28, 0x89, 0xf2, 0x83,
	0x33, 0x92, 0xe0, 0x47, 0x5c, 0x2b, 0xb7, 0xae, 0xef, 0xae, 0x7f, 0x18, 0xae, 0xdf, 0xc7, 0xfd,
	0x9e, 0x3f, 0x6e, 0xfa, 0xf7, 0x92, 0x1c, 0xdc, 0x57, 0xe6, 0x2b, 0x7a, 0x1c, 0x56, 0xd8, 0xa3,
	0xcb, 0x77, 0x3f, 0xbb, 0x05, 0xfa, 0xd5, 0x2d, 0xd0, 0xef, 0x6e, 0x81, 0xbe, 0xbd, 0x11, 0xd2,
	0x57, 0xeb, 0x22, 0xe5, 0xba, 0xc9, 0x58, 0x2d, 0x0b, 0x56, 0xb0, 0xec, 0xb6, 0x47, 0xa1, 0x8b,
	0x3b, 0x6d, 0x2e, 0xa6, 0xa1, 0x8b, 0x6f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x6f, 0xf7,
	0xf0, 0xe5, 0x02, 0x00, 0x00,
}

func (m *McpBridge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *McpBridge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *McpBridge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Registries) > 0 {
		for iNdEx := len(m.Registries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Registries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMcpBridge(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *RegistryConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegistryConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegistryConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ZkServicesPath) > 0 {
		for iNdEx := len(m.ZkServicesPath) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ZkServicesPath[iNdEx])
			copy(dAtA[i:], m.ZkServicesPath[iNdEx])
			i = encodeVarintMcpBridge(dAtA, i, uint64(len(m.ZkServicesPath[iNdEx])))
			i--
			dAtA[i] = 0x6a
		}
	}
	if len(m.ConsulNamespace) > 0 {
		i -= len(m.ConsulNamespace)
		copy(dAtA[i:], m.ConsulNamespace)
		i = encodeVarintMcpBridge(dAtA, i, uint64(len(m.ConsulNamespace)))
		i--
		dAtA[i] = 0x62
	}
	if m.NacosRefreshInterval != 0 {
		i = encodeVarintMcpBridge(dAtA, i, uint64(m.NacosRefreshInterval))
		i--
		dAtA[i] = 0x58
	}
	if len(m.NacosGroups) > 0 {
		for iNdEx := len(m.NacosGroups) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.NacosGroups[iNdEx])
			copy(dAtA[i:], m.NacosGroups[iNdEx])
			i = encodeVarintMcpBridge(dAtA, i, uint64(len(m.NacosGroups[iNdEx])))
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.NacosNamespace) > 0 {
		i -= len(m.NacosNamespace)
		copy(dAtA[i:], m.NacosNamespace)
		i = encodeVarintMcpBridge(dAtA, i, uint64(len(m.NacosNamespace)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.NacosNamespaceId) > 0 {
		i -= len(m.NacosNamespaceId)
		copy(dAtA[i:], m.NacosNamespaceId)
		i = encodeVarintMcpBridge(dAtA, i, uint64(len(m.NacosNamespaceId)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.NacosSecretKey) > 0 {
		i -= len(m.NacosSecretKey)
		copy(dAtA[i:], m.NacosSecretKey)
		i = encodeVarintMcpBridge(dAtA, i, uint64(len(m.NacosSecretKey)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.NacosAccessKey) > 0 {
		i -= len(m.NacosAccessKey)
		copy(dAtA[i:], m.NacosAccessKey)
		i = encodeVarintMcpBridge(dAtA, i, uint64(len(m.NacosAccessKey)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.NacosAddressServer) > 0 {
		i -= len(m.NacosAddressServer)
		copy(dAtA[i:], m.NacosAddressServer)
		i = encodeVarintMcpBridge(dAtA, i, uint64(len(m.NacosAddressServer)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Port != 0 {
		i = encodeVarintMcpBridge(dAtA, i, uint64(m.Port))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Domain) > 0 {
		i -= len(m.Domain)
		copy(dAtA[i:], m.Domain)
		i = encodeVarintMcpBridge(dAtA, i, uint64(len(m.Domain)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintMcpBridge(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintMcpBridge(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMcpBridge(dAtA []byte, offset int, v uint64) int {
	offset -= sovMcpBridge(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *McpBridge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Registries) > 0 {
		for _, e := range m.Registries {
			l = e.Size()
			n += 1 + l + sovMcpBridge(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RegistryConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovMcpBridge(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMcpBridge(uint64(l))
	}
	l = len(m.Domain)
	if l > 0 {
		n += 1 + l + sovMcpBridge(uint64(l))
	}
	if m.Port != 0 {
		n += 1 + sovMcpBridge(uint64(m.Port))
	}
	l = len(m.NacosAddressServer)
	if l > 0 {
		n += 1 + l + sovMcpBridge(uint64(l))
	}
	l = len(m.NacosAccessKey)
	if l > 0 {
		n += 1 + l + sovMcpBridge(uint64(l))
	}
	l = len(m.NacosSecretKey)
	if l > 0 {
		n += 1 + l + sovMcpBridge(uint64(l))
	}
	l = len(m.NacosNamespaceId)
	if l > 0 {
		n += 1 + l + sovMcpBridge(uint64(l))
	}
	l = len(m.NacosNamespace)
	if l > 0 {
		n += 1 + l + sovMcpBridge(uint64(l))
	}
	if len(m.NacosGroups) > 0 {
		for _, s := range m.NacosGroups {
			l = len(s)
			n += 1 + l + sovMcpBridge(uint64(l))
		}
	}
	if m.NacosRefreshInterval != 0 {
		n += 1 + sovMcpBridge(uint64(m.NacosRefreshInterval))
	}
	l = len(m.ConsulNamespace)
	if l > 0 {
		n += 1 + l + sovMcpBridge(uint64(l))
	}
	if len(m.ZkServicesPath) > 0 {
		for _, s := range m.ZkServicesPath {
			l = len(s)
			n += 1 + l + sovMcpBridge(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMcpBridge(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMcpBridge(x uint64) (n int) {
	return sovMcpBridge(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *McpBridge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMcpBridge
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: McpBridge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: McpBridge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Registries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMcpBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMcpBridge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMcpBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Registries = append(m.Registries, &RegistryConfig{})
			if err := m.Registries[len(m.Registries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMcpBridge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMcpBridge
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RegistryConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMcpBridge
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegistryConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegistryConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMcpBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMcpBridge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMcpBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMcpBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMcpBridge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMcpBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMcpBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMcpBridge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMcpBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMcpBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NacosAddressServer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMcpBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMcpBridge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMcpBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NacosAddressServer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NacosAccessKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMcpBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMcpBridge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMcpBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NacosAccessKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NacosSecretKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMcpBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMcpBridge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMcpBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NacosSecretKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NacosNamespaceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMcpBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMcpBridge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMcpBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NacosNamespaceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NacosNamespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMcpBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMcpBridge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMcpBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NacosNamespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NacosGroups", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMcpBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMcpBridge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMcpBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NacosGroups = append(m.NacosGroups, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NacosRefreshInterval", wireType)
			}
			m.NacosRefreshInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMcpBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NacosRefreshInterval |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsulNamespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMcpBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMcpBridge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMcpBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConsulNamespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZkServicesPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMcpBridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMcpBridge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMcpBridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ZkServicesPath = append(m.ZkServicesPath, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMcpBridge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMcpBridge
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMcpBridge(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMcpBridge
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMcpBridge
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMcpBridge
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMcpBridge
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMcpBridge
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMcpBridge
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMcpBridge        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMcpBridge          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMcpBridge = fmt.Errorf("proto: unexpected end of group")
)