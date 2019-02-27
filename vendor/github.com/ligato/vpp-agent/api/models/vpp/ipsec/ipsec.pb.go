// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: models/vpp/ipsec/ipsec.proto

package vpp_ipsec // import "github.com/ligato/vpp-agent/api/models/vpp/ipsec"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Cryptographic algorithm for encryption
type CryptoAlg int32

const (
	CryptoAlg_NONE_CRYPTO CryptoAlg = 0
	CryptoAlg_AES_CBC_128 CryptoAlg = 1
	CryptoAlg_AES_CBC_192 CryptoAlg = 2
	CryptoAlg_AES_CBC_256 CryptoAlg = 3
	CryptoAlg_AES_CTR_128 CryptoAlg = 4
	CryptoAlg_AES_CTR_192 CryptoAlg = 5
	CryptoAlg_AES_CTR_256 CryptoAlg = 6
	CryptoAlg_AES_GCM_128 CryptoAlg = 7
	CryptoAlg_AES_GCM_192 CryptoAlg = 8
	CryptoAlg_AES_GCM_256 CryptoAlg = 9
	CryptoAlg_DES_CBC     CryptoAlg = 10
	CryptoAlg_DES3_CBC    CryptoAlg = 11
)

var CryptoAlg_name = map[int32]string{
	0:  "NONE_CRYPTO",
	1:  "AES_CBC_128",
	2:  "AES_CBC_192",
	3:  "AES_CBC_256",
	4:  "AES_CTR_128",
	5:  "AES_CTR_192",
	6:  "AES_CTR_256",
	7:  "AES_GCM_128",
	8:  "AES_GCM_192",
	9:  "AES_GCM_256",
	10: "DES_CBC",
	11: "DES3_CBC",
}
var CryptoAlg_value = map[string]int32{
	"NONE_CRYPTO": 0,
	"AES_CBC_128": 1,
	"AES_CBC_192": 2,
	"AES_CBC_256": 3,
	"AES_CTR_128": 4,
	"AES_CTR_192": 5,
	"AES_CTR_256": 6,
	"AES_GCM_128": 7,
	"AES_GCM_192": 8,
	"AES_GCM_256": 9,
	"DES_CBC":     10,
	"DES3_CBC":    11,
}

func (x CryptoAlg) String() string {
	return proto.EnumName(CryptoAlg_name, int32(x))
}
func (CryptoAlg) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ipsec_de07facc5beafc91, []int{0}
}

// Cryptographic algorithm for authentication
type IntegAlg int32

const (
	IntegAlg_NONE_INTEG  IntegAlg = 0
	IntegAlg_MD5_96      IntegAlg = 1
	IntegAlg_SHA1_96     IntegAlg = 2
	IntegAlg_SHA_256_96  IntegAlg = 3
	IntegAlg_SHA_256_128 IntegAlg = 4
	IntegAlg_SHA_384_192 IntegAlg = 5
	IntegAlg_SHA_512_256 IntegAlg = 6
)

var IntegAlg_name = map[int32]string{
	0: "NONE_INTEG",
	1: "MD5_96",
	2: "SHA1_96",
	3: "SHA_256_96",
	4: "SHA_256_128",
	5: "SHA_384_192",
	6: "SHA_512_256",
}
var IntegAlg_value = map[string]int32{
	"NONE_INTEG":  0,
	"MD5_96":      1,
	"SHA1_96":     2,
	"SHA_256_96":  3,
	"SHA_256_128": 4,
	"SHA_384_192": 5,
	"SHA_512_256": 6,
}

func (x IntegAlg) String() string {
	return proto.EnumName(IntegAlg_name, int32(x))
}
func (IntegAlg) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ipsec_de07facc5beafc91, []int{1}
}

type SecurityPolicyDatabase_PolicyEntry_Action int32

const (
	SecurityPolicyDatabase_PolicyEntry_BYPASS  SecurityPolicyDatabase_PolicyEntry_Action = 0
	SecurityPolicyDatabase_PolicyEntry_DISCARD SecurityPolicyDatabase_PolicyEntry_Action = 1
	SecurityPolicyDatabase_PolicyEntry_RESOLVE SecurityPolicyDatabase_PolicyEntry_Action = 2
	SecurityPolicyDatabase_PolicyEntry_PROTECT SecurityPolicyDatabase_PolicyEntry_Action = 3
)

var SecurityPolicyDatabase_PolicyEntry_Action_name = map[int32]string{
	0: "BYPASS",
	1: "DISCARD",
	2: "RESOLVE",
	3: "PROTECT",
}
var SecurityPolicyDatabase_PolicyEntry_Action_value = map[string]int32{
	"BYPASS":  0,
	"DISCARD": 1,
	"RESOLVE": 2,
	"PROTECT": 3,
}

func (x SecurityPolicyDatabase_PolicyEntry_Action) String() string {
	return proto.EnumName(SecurityPolicyDatabase_PolicyEntry_Action_name, int32(x))
}
func (SecurityPolicyDatabase_PolicyEntry_Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ipsec_de07facc5beafc91, []int{0, 1, 0}
}

type SecurityAssociation_IPSecProtocol int32

const (
	SecurityAssociation_AH  SecurityAssociation_IPSecProtocol = 0
	SecurityAssociation_ESP SecurityAssociation_IPSecProtocol = 1
)

var SecurityAssociation_IPSecProtocol_name = map[int32]string{
	0: "AH",
	1: "ESP",
}
var SecurityAssociation_IPSecProtocol_value = map[string]int32{
	"AH":  0,
	"ESP": 1,
}

func (x SecurityAssociation_IPSecProtocol) String() string {
	return proto.EnumName(SecurityAssociation_IPSecProtocol_name, int32(x))
}
func (SecurityAssociation_IPSecProtocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ipsec_de07facc5beafc91, []int{1, 0}
}

// Security Policy Database (SPD)
type SecurityPolicyDatabase struct {
	Index                string                                `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Interfaces           []*SecurityPolicyDatabase_Interface   `protobuf:"bytes,2,rep,name=interfaces,proto3" json:"interfaces,omitempty"`
	PolicyEntries        []*SecurityPolicyDatabase_PolicyEntry `protobuf:"bytes,3,rep,name=policy_entries,json=policyEntries,proto3" json:"policy_entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *SecurityPolicyDatabase) Reset()         { *m = SecurityPolicyDatabase{} }
func (m *SecurityPolicyDatabase) String() string { return proto.CompactTextString(m) }
func (*SecurityPolicyDatabase) ProtoMessage()    {}
func (*SecurityPolicyDatabase) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipsec_de07facc5beafc91, []int{0}
}
func (m *SecurityPolicyDatabase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecurityPolicyDatabase.Unmarshal(m, b)
}
func (m *SecurityPolicyDatabase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecurityPolicyDatabase.Marshal(b, m, deterministic)
}
func (dst *SecurityPolicyDatabase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecurityPolicyDatabase.Merge(dst, src)
}
func (m *SecurityPolicyDatabase) XXX_Size() int {
	return xxx_messageInfo_SecurityPolicyDatabase.Size(m)
}
func (m *SecurityPolicyDatabase) XXX_DiscardUnknown() {
	xxx_messageInfo_SecurityPolicyDatabase.DiscardUnknown(m)
}

var xxx_messageInfo_SecurityPolicyDatabase proto.InternalMessageInfo

func (m *SecurityPolicyDatabase) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *SecurityPolicyDatabase) GetInterfaces() []*SecurityPolicyDatabase_Interface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

func (m *SecurityPolicyDatabase) GetPolicyEntries() []*SecurityPolicyDatabase_PolicyEntry {
	if m != nil {
		return m.PolicyEntries
	}
	return nil
}

func (*SecurityPolicyDatabase) XXX_MessageName() string {
	return "vpp.ipsec.SecurityPolicyDatabase"
}

type SecurityPolicyDatabase_Interface struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SecurityPolicyDatabase_Interface) Reset()         { *m = SecurityPolicyDatabase_Interface{} }
func (m *SecurityPolicyDatabase_Interface) String() string { return proto.CompactTextString(m) }
func (*SecurityPolicyDatabase_Interface) ProtoMessage()    {}
func (*SecurityPolicyDatabase_Interface) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipsec_de07facc5beafc91, []int{0, 0}
}
func (m *SecurityPolicyDatabase_Interface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecurityPolicyDatabase_Interface.Unmarshal(m, b)
}
func (m *SecurityPolicyDatabase_Interface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecurityPolicyDatabase_Interface.Marshal(b, m, deterministic)
}
func (dst *SecurityPolicyDatabase_Interface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecurityPolicyDatabase_Interface.Merge(dst, src)
}
func (m *SecurityPolicyDatabase_Interface) XXX_Size() int {
	return xxx_messageInfo_SecurityPolicyDatabase_Interface.Size(m)
}
func (m *SecurityPolicyDatabase_Interface) XXX_DiscardUnknown() {
	xxx_messageInfo_SecurityPolicyDatabase_Interface.DiscardUnknown(m)
}

var xxx_messageInfo_SecurityPolicyDatabase_Interface proto.InternalMessageInfo

func (m *SecurityPolicyDatabase_Interface) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (*SecurityPolicyDatabase_Interface) XXX_MessageName() string {
	return "vpp.ipsec.SecurityPolicyDatabase.Interface"
}

type SecurityPolicyDatabase_PolicyEntry struct {
	SaIndex              string                                    `protobuf:"bytes,1,opt,name=sa_index,json=saIndex,proto3" json:"sa_index,omitempty"`
	Priority             int32                                     `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	IsOutbound           bool                                      `protobuf:"varint,3,opt,name=is_outbound,json=isOutbound,proto3" json:"is_outbound,omitempty"`
	RemoteAddrStart      string                                    `protobuf:"bytes,4,opt,name=remote_addr_start,json=remoteAddrStart,proto3" json:"remote_addr_start,omitempty"`
	RemoteAddrStop       string                                    `protobuf:"bytes,5,opt,name=remote_addr_stop,json=remoteAddrStop,proto3" json:"remote_addr_stop,omitempty"`
	LocalAddrStart       string                                    `protobuf:"bytes,6,opt,name=local_addr_start,json=localAddrStart,proto3" json:"local_addr_start,omitempty"`
	LocalAddrStop        string                                    `protobuf:"bytes,7,opt,name=local_addr_stop,json=localAddrStop,proto3" json:"local_addr_stop,omitempty"`
	Protocol             uint32                                    `protobuf:"varint,8,opt,name=protocol,proto3" json:"protocol,omitempty"`
	RemotePortStart      uint32                                    `protobuf:"varint,9,opt,name=remote_port_start,json=remotePortStart,proto3" json:"remote_port_start,omitempty"`
	RemotePortStop       uint32                                    `protobuf:"varint,10,opt,name=remote_port_stop,json=remotePortStop,proto3" json:"remote_port_stop,omitempty"`
	LocalPortStart       uint32                                    `protobuf:"varint,11,opt,name=local_port_start,json=localPortStart,proto3" json:"local_port_start,omitempty"`
	LocalPortStop        uint32                                    `protobuf:"varint,12,opt,name=local_port_stop,json=localPortStop,proto3" json:"local_port_stop,omitempty"`
	Action               SecurityPolicyDatabase_PolicyEntry_Action `protobuf:"varint,13,opt,name=action,proto3,enum=vpp.ipsec.SecurityPolicyDatabase_PolicyEntry_Action" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *SecurityPolicyDatabase_PolicyEntry) Reset()         { *m = SecurityPolicyDatabase_PolicyEntry{} }
func (m *SecurityPolicyDatabase_PolicyEntry) String() string { return proto.CompactTextString(m) }
func (*SecurityPolicyDatabase_PolicyEntry) ProtoMessage()    {}
func (*SecurityPolicyDatabase_PolicyEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipsec_de07facc5beafc91, []int{0, 1}
}
func (m *SecurityPolicyDatabase_PolicyEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecurityPolicyDatabase_PolicyEntry.Unmarshal(m, b)
}
func (m *SecurityPolicyDatabase_PolicyEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecurityPolicyDatabase_PolicyEntry.Marshal(b, m, deterministic)
}
func (dst *SecurityPolicyDatabase_PolicyEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecurityPolicyDatabase_PolicyEntry.Merge(dst, src)
}
func (m *SecurityPolicyDatabase_PolicyEntry) XXX_Size() int {
	return xxx_messageInfo_SecurityPolicyDatabase_PolicyEntry.Size(m)
}
func (m *SecurityPolicyDatabase_PolicyEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_SecurityPolicyDatabase_PolicyEntry.DiscardUnknown(m)
}

var xxx_messageInfo_SecurityPolicyDatabase_PolicyEntry proto.InternalMessageInfo

func (m *SecurityPolicyDatabase_PolicyEntry) GetSaIndex() string {
	if m != nil {
		return m.SaIndex
	}
	return ""
}

func (m *SecurityPolicyDatabase_PolicyEntry) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *SecurityPolicyDatabase_PolicyEntry) GetIsOutbound() bool {
	if m != nil {
		return m.IsOutbound
	}
	return false
}

func (m *SecurityPolicyDatabase_PolicyEntry) GetRemoteAddrStart() string {
	if m != nil {
		return m.RemoteAddrStart
	}
	return ""
}

func (m *SecurityPolicyDatabase_PolicyEntry) GetRemoteAddrStop() string {
	if m != nil {
		return m.RemoteAddrStop
	}
	return ""
}

func (m *SecurityPolicyDatabase_PolicyEntry) GetLocalAddrStart() string {
	if m != nil {
		return m.LocalAddrStart
	}
	return ""
}

func (m *SecurityPolicyDatabase_PolicyEntry) GetLocalAddrStop() string {
	if m != nil {
		return m.LocalAddrStop
	}
	return ""
}

func (m *SecurityPolicyDatabase_PolicyEntry) GetProtocol() uint32 {
	if m != nil {
		return m.Protocol
	}
	return 0
}

func (m *SecurityPolicyDatabase_PolicyEntry) GetRemotePortStart() uint32 {
	if m != nil {
		return m.RemotePortStart
	}
	return 0
}

func (m *SecurityPolicyDatabase_PolicyEntry) GetRemotePortStop() uint32 {
	if m != nil {
		return m.RemotePortStop
	}
	return 0
}

func (m *SecurityPolicyDatabase_PolicyEntry) GetLocalPortStart() uint32 {
	if m != nil {
		return m.LocalPortStart
	}
	return 0
}

func (m *SecurityPolicyDatabase_PolicyEntry) GetLocalPortStop() uint32 {
	if m != nil {
		return m.LocalPortStop
	}
	return 0
}

func (m *SecurityPolicyDatabase_PolicyEntry) GetAction() SecurityPolicyDatabase_PolicyEntry_Action {
	if m != nil {
		return m.Action
	}
	return SecurityPolicyDatabase_PolicyEntry_BYPASS
}

func (*SecurityPolicyDatabase_PolicyEntry) XXX_MessageName() string {
	return "vpp.ipsec.SecurityPolicyDatabase.PolicyEntry"
}

// Security Association (SA)
type SecurityAssociation struct {
	Index                string                            `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Spi                  uint32                            `protobuf:"varint,2,opt,name=spi,proto3" json:"spi,omitempty"`
	Protocol             SecurityAssociation_IPSecProtocol `protobuf:"varint,3,opt,name=protocol,proto3,enum=vpp.ipsec.SecurityAssociation_IPSecProtocol" json:"protocol,omitempty"`
	CryptoAlg            CryptoAlg                         `protobuf:"varint,4,opt,name=crypto_alg,json=cryptoAlg,proto3,enum=vpp.ipsec.CryptoAlg" json:"crypto_alg,omitempty"`
	CryptoKey            string                            `protobuf:"bytes,5,opt,name=crypto_key,json=cryptoKey,proto3" json:"crypto_key,omitempty"`
	IntegAlg             IntegAlg                          `protobuf:"varint,6,opt,name=integ_alg,json=integAlg,proto3,enum=vpp.ipsec.IntegAlg" json:"integ_alg,omitempty"`
	IntegKey             string                            `protobuf:"bytes,7,opt,name=integ_key,json=integKey,proto3" json:"integ_key,omitempty"`
	UseEsn               bool                              `protobuf:"varint,8,opt,name=use_esn,json=useEsn,proto3" json:"use_esn,omitempty"`
	UseAntiReplay        bool                              `protobuf:"varint,9,opt,name=use_anti_replay,json=useAntiReplay,proto3" json:"use_anti_replay,omitempty"`
	TunnelSrcAddr        string                            `protobuf:"bytes,10,opt,name=tunnel_src_addr,json=tunnelSrcAddr,proto3" json:"tunnel_src_addr,omitempty"`
	TunnelDstAddr        string                            `protobuf:"bytes,11,opt,name=tunnel_dst_addr,json=tunnelDstAddr,proto3" json:"tunnel_dst_addr,omitempty"`
	EnableUdpEncap       bool                              `protobuf:"varint,12,opt,name=enable_udp_encap,json=enableUdpEncap,proto3" json:"enable_udp_encap,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *SecurityAssociation) Reset()         { *m = SecurityAssociation{} }
func (m *SecurityAssociation) String() string { return proto.CompactTextString(m) }
func (*SecurityAssociation) ProtoMessage()    {}
func (*SecurityAssociation) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipsec_de07facc5beafc91, []int{1}
}
func (m *SecurityAssociation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecurityAssociation.Unmarshal(m, b)
}
func (m *SecurityAssociation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecurityAssociation.Marshal(b, m, deterministic)
}
func (dst *SecurityAssociation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecurityAssociation.Merge(dst, src)
}
func (m *SecurityAssociation) XXX_Size() int {
	return xxx_messageInfo_SecurityAssociation.Size(m)
}
func (m *SecurityAssociation) XXX_DiscardUnknown() {
	xxx_messageInfo_SecurityAssociation.DiscardUnknown(m)
}

var xxx_messageInfo_SecurityAssociation proto.InternalMessageInfo

func (m *SecurityAssociation) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *SecurityAssociation) GetSpi() uint32 {
	if m != nil {
		return m.Spi
	}
	return 0
}

func (m *SecurityAssociation) GetProtocol() SecurityAssociation_IPSecProtocol {
	if m != nil {
		return m.Protocol
	}
	return SecurityAssociation_AH
}

func (m *SecurityAssociation) GetCryptoAlg() CryptoAlg {
	if m != nil {
		return m.CryptoAlg
	}
	return CryptoAlg_NONE_CRYPTO
}

func (m *SecurityAssociation) GetCryptoKey() string {
	if m != nil {
		return m.CryptoKey
	}
	return ""
}

func (m *SecurityAssociation) GetIntegAlg() IntegAlg {
	if m != nil {
		return m.IntegAlg
	}
	return IntegAlg_NONE_INTEG
}

func (m *SecurityAssociation) GetIntegKey() string {
	if m != nil {
		return m.IntegKey
	}
	return ""
}

func (m *SecurityAssociation) GetUseEsn() bool {
	if m != nil {
		return m.UseEsn
	}
	return false
}

func (m *SecurityAssociation) GetUseAntiReplay() bool {
	if m != nil {
		return m.UseAntiReplay
	}
	return false
}

func (m *SecurityAssociation) GetTunnelSrcAddr() string {
	if m != nil {
		return m.TunnelSrcAddr
	}
	return ""
}

func (m *SecurityAssociation) GetTunnelDstAddr() string {
	if m != nil {
		return m.TunnelDstAddr
	}
	return ""
}

func (m *SecurityAssociation) GetEnableUdpEncap() bool {
	if m != nil {
		return m.EnableUdpEncap
	}
	return false
}

func (*SecurityAssociation) XXX_MessageName() string {
	return "vpp.ipsec.SecurityAssociation"
}
func init() {
	proto.RegisterType((*SecurityPolicyDatabase)(nil), "vpp.ipsec.SecurityPolicyDatabase")
	proto.RegisterType((*SecurityPolicyDatabase_Interface)(nil), "vpp.ipsec.SecurityPolicyDatabase.Interface")
	proto.RegisterType((*SecurityPolicyDatabase_PolicyEntry)(nil), "vpp.ipsec.SecurityPolicyDatabase.PolicyEntry")
	proto.RegisterType((*SecurityAssociation)(nil), "vpp.ipsec.SecurityAssociation")
	proto.RegisterEnum("vpp.ipsec.CryptoAlg", CryptoAlg_name, CryptoAlg_value)
	proto.RegisterEnum("vpp.ipsec.IntegAlg", IntegAlg_name, IntegAlg_value)
	proto.RegisterEnum("vpp.ipsec.SecurityPolicyDatabase_PolicyEntry_Action", SecurityPolicyDatabase_PolicyEntry_Action_name, SecurityPolicyDatabase_PolicyEntry_Action_value)
	proto.RegisterEnum("vpp.ipsec.SecurityAssociation_IPSecProtocol", SecurityAssociation_IPSecProtocol_name, SecurityAssociation_IPSecProtocol_value)
}

func init() { proto.RegisterFile("models/vpp/ipsec/ipsec.proto", fileDescriptor_ipsec_de07facc5beafc91) }

var fileDescriptor_ipsec_de07facc5beafc91 = []byte{
	// 919 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x5d, 0x6f, 0xdb, 0x36,
	0x14, 0x8d, 0xec, 0x44, 0x96, 0xe9, 0xda, 0xd1, 0xd8, 0x62, 0xd3, 0xb2, 0x8f, 0x1a, 0x79, 0x18,
	0x8c, 0x6c, 0xb1, 0x17, 0xa7, 0x29, 0x9a, 0xf5, 0x65, 0x8a, 0x2d, 0x34, 0x46, 0xdb, 0xd8, 0x90,
	0xbc, 0x01, 0xdd, 0x8b, 0x40, 0x4b, 0xac, 0x27, 0x4c, 0x11, 0x09, 0x91, 0x0a, 0xe6, 0x5f, 0xb0,
	0xbf, 0xb6, 0xc7, 0xed, 0x7d, 0x6f, 0xfb, 0x23, 0x03, 0xaf, 0x64, 0x85, 0x2e, 0x02, 0x0c, 0x7d,
	0x31, 0x78, 0x0e, 0xcf, 0x3d, 0xf4, 0x11, 0xef, 0x05, 0xd1, 0x97, 0xb7, 0x2c, 0xa6, 0xa9, 0x18,
	0xdd, 0x71, 0x3e, 0x4a, 0xb8, 0xa0, 0x51, 0xf9, 0x3b, 0xe4, 0x39, 0x93, 0x0c, 0xb7, 0xef, 0x38,
	0x1f, 0x02, 0x71, 0x74, 0xba, 0x4e, 0xe4, 0xaf, 0xc5, 0x6a, 0x18, 0xb1, 0xdb, 0xd1, 0x9a, 0xad,
	0xd9, 0x08, 0x14, 0xab, 0xe2, 0x3d, 0x20, 0x00, 0xb0, 0x2a, 0x2b, 0x8f, 0xff, 0x36, 0xd1, 0xa7,
	0x01, 0x8d, 0x8a, 0x3c, 0x91, 0x9b, 0x05, 0x4b, 0x93, 0x68, 0x33, 0x25, 0x92, 0xac, 0x88, 0xa0,
	0xf8, 0x09, 0x3a, 0x48, 0xb2, 0x98, 0xfe, 0xee, 0x18, 0x7d, 0x63, 0xd0, 0xf6, 0x4b, 0x80, 0x5f,
	0x23, 0x94, 0x64, 0x92, 0xe6, 0xef, 0x49, 0x44, 0x85, 0xd3, 0xe8, 0x37, 0x07, 0x9d, 0xf1, 0xb7,
	0xc3, 0xfa, 0xfc, 0xe1, 0xc3, 0x66, 0xc3, 0xd9, 0xb6, 0xc6, 0xd7, 0xca, 0xf1, 0x12, 0xf5, 0x38,
	0xe8, 0x42, 0x9a, 0xc9, 0x3c, 0xa1, 0xc2, 0x69, 0x82, 0xe1, 0xe9, 0xff, 0x1b, 0x96, 0xd0, 0xcb,
	0x64, 0xbe, 0xf1, 0xbb, 0xbc, 0x06, 0x09, 0x15, 0x47, 0x4f, 0x51, 0xbb, 0x3e, 0x0e, 0x63, 0xb4,
	0x9f, 0x91, 0x5b, 0x5a, 0x85, 0x80, 0xf5, 0xd1, 0x3f, 0xfb, 0xa8, 0xa3, 0xd5, 0xe3, 0xcf, 0x91,
	0x25, 0x48, 0xa8, 0x87, 0x6d, 0x09, 0x32, 0x83, 0xb8, 0x47, 0xc8, 0xe2, 0x79, 0xc2, 0xd4, 0x1f,
	0x70, 0x1a, 0x7d, 0x63, 0x70, 0xe0, 0xd7, 0x18, 0x3f, 0x45, 0x9d, 0x44, 0x84, 0xac, 0x90, 0x2b,
	0x56, 0x64, 0xb1, 0xd3, 0xec, 0x1b, 0x03, 0xcb, 0x47, 0x89, 0x98, 0x57, 0x0c, 0x3e, 0x41, 0x9f,
	0xe4, 0xf4, 0x96, 0x49, 0x1a, 0x92, 0x38, 0xce, 0x43, 0x21, 0x49, 0x2e, 0x9d, 0x7d, 0x38, 0xe0,
	0xb0, 0xdc, 0x70, 0xe3, 0x38, 0x0f, 0x14, 0x8d, 0x07, 0xc8, 0xde, 0xd5, 0x32, 0xee, 0x1c, 0x80,
	0xb4, 0xa7, 0x4b, 0x19, 0x57, 0xca, 0x94, 0x45, 0x24, 0xd5, 0x4d, 0xcd, 0x52, 0x09, 0xfc, 0xbd,
	0xe7, 0x37, 0xe8, 0x70, 0x47, 0xc9, 0xb8, 0xd3, 0x02, 0x61, 0x57, 0x13, 0x32, 0x5e, 0x86, 0x64,
	0x92, 0x45, 0x2c, 0x75, 0xac, 0xbe, 0x31, 0xe8, 0xfa, 0x35, 0xd6, 0x32, 0x70, 0x96, 0xcb, 0xea,
	0xb8, 0x36, 0x88, 0xaa, 0x0c, 0x0b, 0x96, 0xcb, 0x0f, 0x33, 0x54, 0x5a, 0xc6, 0x1d, 0x04, 0xd2,
	0x9e, 0x2e, 0xd5, 0x33, 0x68, 0xa6, 0x9d, 0x52, 0x09, 0xfc, 0xbd, 0x67, 0x9d, 0xe1, 0xde, 0xf2,
	0x11, 0x08, 0xbb, 0x9a, 0x90, 0x71, 0xfc, 0x06, 0x99, 0x24, 0x92, 0x09, 0xcb, 0x9c, 0x6e, 0xdf,
	0x18, 0xf4, 0xc6, 0xcf, 0x3e, 0xaa, 0x85, 0x86, 0x2e, 0xd4, 0xfa, 0x95, 0xc7, 0xf1, 0x4b, 0x64,
	0x96, 0x0c, 0x46, 0xc8, 0xbc, 0x7a, 0xb7, 0x70, 0x83, 0xc0, 0xde, 0xc3, 0x1d, 0xd4, 0x9a, 0xce,
	0x82, 0x89, 0xeb, 0x4f, 0x6d, 0x43, 0x01, 0xdf, 0x0b, 0xe6, 0x6f, 0x7e, 0xf6, 0xec, 0x86, 0x02,
	0x0b, 0x7f, 0xbe, 0xf4, 0x26, 0x4b, 0xbb, 0x79, 0xfc, 0xc7, 0x3e, 0x7a, 0xbc, 0x3d, 0xd2, 0x15,
	0x82, 0x45, 0x09, 0x01, 0xab, 0x87, 0x07, 0xca, 0x46, 0x4d, 0xc1, 0x13, 0x68, 0xae, 0xae, 0xaf,
	0x96, 0xf8, 0x5a, 0xbb, 0x8e, 0x26, 0x84, 0xf9, 0xee, 0x81, 0x30, 0x9a, 0xf3, 0x70, 0xb6, 0x08,
	0x68, 0xb4, 0xa8, 0x6a, 0xb4, 0xcb, 0x3b, 0x47, 0x28, 0xca, 0x37, 0x5c, 0xb2, 0x90, 0xa4, 0x6b,
	0xe8, 0xbc, 0xde, 0xf8, 0x89, 0xe6, 0x35, 0x81, 0x4d, 0x37, 0x5d, 0xfb, 0xed, 0x68, 0xbb, 0xc4,
	0x5f, 0xd5, 0x45, 0xbf, 0xd1, 0x4d, 0xd5, 0x83, 0xd5, 0xf6, 0x6b, 0xba, 0xc1, 0xdf, 0xa3, 0xb6,
	0x9a, 0xe0, 0x35, 0x58, 0x9a, 0x60, 0xf9, 0x58, 0xb3, 0x54, 0x93, 0xb7, 0x56, 0x8e, 0x56, 0x52,
	0xad, 0xf0, 0x17, 0xdb, 0x0a, 0xe5, 0x57, 0x36, 0x60, 0xb9, 0xa9, 0xec, 0x3e, 0x43, 0xad, 0x42,
	0xd0, 0x90, 0x8a, 0x0c, 0x5a, 0xcf, 0xf2, 0xcd, 0x42, 0x50, 0x4f, 0x64, 0xea, 0xe2, 0xd5, 0x06,
	0xc9, 0x64, 0x12, 0xe6, 0x94, 0xa7, 0x64, 0x03, 0x6d, 0x67, 0xf9, 0xdd, 0x42, 0x50, 0x37, 0x93,
	0x89, 0x0f, 0xa4, 0xd2, 0xc9, 0x22, 0xcb, 0x68, 0x1a, 0x8a, 0x3c, 0x82, 0x4e, 0x87, 0x9e, 0x6b,
	0xfb, 0xdd, 0x92, 0x0e, 0xf2, 0x48, 0x35, 0xba, 0xa6, 0x8b, 0x85, 0x2c, 0x75, 0x1d, 0x5d, 0x37,
	0x15, 0x12, 0x74, 0x03, 0x64, 0xd3, 0x8c, 0xac, 0x52, 0x1a, 0x16, 0x31, 0x0f, 0x69, 0x16, 0x91,
	0xb2, 0xe3, 0x2c, 0xbf, 0x57, 0xf2, 0x3f, 0xc5, 0xdc, 0x53, 0xec, 0x71, 0x1f, 0x75, 0x77, 0x3e,
	0x3c, 0x36, 0x51, 0xc3, 0xbd, 0xb6, 0xf7, 0x70, 0x0b, 0x35, 0xbd, 0x60, 0x61, 0x1b, 0x27, 0x7f,
	0x19, 0xa8, 0x5d, 0x7f, 0x63, 0x7c, 0x88, 0x3a, 0x37, 0xf3, 0x1b, 0x2f, 0x9c, 0xf8, 0xef, 0x16,
	0xcb, 0xb9, 0xbd, 0xa7, 0x08, 0xd7, 0x0b, 0xc2, 0xc9, 0xd5, 0x24, 0x3c, 0x1b, 0xbf, 0xb0, 0x8d,
	0x1d, 0xe2, 0x72, 0x6c, 0x37, 0x74, 0x62, 0x7c, 0xf1, 0xdc, 0x6e, 0xd6, 0xc4, 0xd2, 0x87, 0x92,
	0xfd, 0x1d, 0xe2, 0x72, 0x6c, 0x1f, 0xe8, 0x84, 0x2a, 0x31, 0xb7, 0xc4, 0xab, 0xc9, 0x5b, 0x28,
	0x69, 0xed, 0x10, 0x97, 0x63, 0xdb, 0xd2, 0x09, 0x55, 0xd2, 0x86, 0x46, 0x2f, 0x8f, 0xb5, 0x11,
	0x7e, 0x84, 0xac, 0xa9, 0x17, 0x9c, 0x03, 0xea, 0x9c, 0xdc, 0x21, 0x6b, 0x7b, 0xc5, 0xb8, 0x87,
	0x10, 0x04, 0x9a, 0xdd, 0x2c, 0xbd, 0x57, 0xf6, 0x9e, 0x9a, 0x95, 0xb7, 0xd3, 0x8b, 0xf0, 0xf2,
	0x79, 0x39, 0x1e, 0xc1, 0xb5, 0x7b, 0xa6, 0x40, 0x43, 0x09, 0x83, 0x6b, 0x57, 0x99, 0x2b, 0x0c,
	0x29, 0xb6, 0xb8, 0x4e, 0xa1, 0x88, 0xf3, 0x17, 0xcf, 0xee, 0x53, 0x28, 0xe2, 0xe2, 0x6c, 0x5c,
	0xa6, 0xb8, 0xfa, 0xf1, 0xcf, 0x7f, 0xbf, 0x36, 0x7e, 0xf9, 0x41, 0x7b, 0xdd, 0xd2, 0x64, 0x4d,
	0x24, 0x53, 0x2f, 0xe2, 0x29, 0x59, 0xd3, 0x4c, 0x8e, 0x08, 0x4f, 0x46, 0x1f, 0x3e, 0x93, 0x2f,
	0xef, 0x38, 0x0f, 0x61, 0xb5, 0x32, 0x61, 0x2c, 0xce, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x51,
	0x77, 0xc6, 0x1a, 0x4b, 0x07, 0x00, 0x00,
}