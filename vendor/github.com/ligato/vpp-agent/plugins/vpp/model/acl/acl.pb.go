// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: acl.proto

package acl

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type AclAction int32

const (
	AclAction_DENY    AclAction = 0
	AclAction_PERMIT  AclAction = 1
	AclAction_REFLECT AclAction = 2
)

var AclAction_name = map[int32]string{
	0: "DENY",
	1: "PERMIT",
	2: "REFLECT",
}
var AclAction_value = map[string]int32{
	"DENY":    0,
	"PERMIT":  1,
	"REFLECT": 2,
}

func (x AclAction) String() string {
	return proto.EnumName(AclAction_name, int32(x))
}
func (AclAction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_acl_700e6e99ae46a356, []int{0}
}

// This is a top level container for Access Control Lists.
// It can have one or more Access Control Lists.
type AccessLists struct {
	Acls                 []*AccessLists_Acl `protobuf:"bytes,1,rep,name=acls,proto3" json:"acls,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AccessLists) Reset()         { *m = AccessLists{} }
func (m *AccessLists) String() string { return proto.CompactTextString(m) }
func (*AccessLists) ProtoMessage()    {}
func (*AccessLists) Descriptor() ([]byte, []int) {
	return fileDescriptor_acl_700e6e99ae46a356, []int{0}
}
func (m *AccessLists) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessLists.Unmarshal(m, b)
}
func (m *AccessLists) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessLists.Marshal(b, m, deterministic)
}
func (dst *AccessLists) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessLists.Merge(dst, src)
}
func (m *AccessLists) XXX_Size() int {
	return xxx_messageInfo_AccessLists.Size(m)
}
func (m *AccessLists) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessLists.DiscardUnknown(m)
}

var xxx_messageInfo_AccessLists proto.InternalMessageInfo

func (m *AccessLists) GetAcls() []*AccessLists_Acl {
	if m != nil {
		return m.Acls
	}
	return nil
}

// An Access Control List (ACL) is an ordered list of Access List Rules.
type AccessLists_Acl struct {
	// The name of access list. A device MAY restrict the length
	// and value of this name, possibly spaces and special
	// characters are not allowed.
	AclName string                  `protobuf:"bytes,1,opt,name=acl_name,json=aclName,proto3" json:"acl_name,omitempty"`
	Rules   []*AccessLists_Acl_Rule `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
	// The set of interfaces that has assigned this ACL on ingres or egress.
	Interfaces           *AccessLists_Acl_Interfaces `protobuf:"bytes,3,opt,name=interfaces,proto3" json:"interfaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *AccessLists_Acl) Reset()         { *m = AccessLists_Acl{} }
func (m *AccessLists_Acl) String() string { return proto.CompactTextString(m) }
func (*AccessLists_Acl) ProtoMessage()    {}
func (*AccessLists_Acl) Descriptor() ([]byte, []int) {
	return fileDescriptor_acl_700e6e99ae46a356, []int{0, 0}
}
func (m *AccessLists_Acl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessLists_Acl.Unmarshal(m, b)
}
func (m *AccessLists_Acl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessLists_Acl.Marshal(b, m, deterministic)
}
func (dst *AccessLists_Acl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessLists_Acl.Merge(dst, src)
}
func (m *AccessLists_Acl) XXX_Size() int {
	return xxx_messageInfo_AccessLists_Acl.Size(m)
}
func (m *AccessLists_Acl) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessLists_Acl.DiscardUnknown(m)
}

var xxx_messageInfo_AccessLists_Acl proto.InternalMessageInfo

func (m *AccessLists_Acl) GetAclName() string {
	if m != nil {
		return m.AclName
	}
	return ""
}

func (m *AccessLists_Acl) GetRules() []*AccessLists_Acl_Rule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *AccessLists_Acl) GetInterfaces() *AccessLists_Acl_Interfaces {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

// List of access list entries (Rules). Each Access Control Rule has
// a list of match criteria and a list of actions.
// Access List entry that can define:
// - IPv4/IPv6 src ip prefix
// - src MAC address mask
// - src MAC address value
// - can be used only for static ACLs.
type AccessLists_Acl_Rule struct {
	// A unique name identifying this Access List Entry (Rule)
	RuleName string `protobuf:"bytes,1,opt,name=rule_name,json=ruleName,proto3" json:"rule_name,omitempty"`
	// Action for this Access List Rule
	AclAction            AclAction                   `protobuf:"varint,2,opt,name=acl_action,json=aclAction,proto3,enum=acl.AclAction" json:"acl_action,omitempty"`
	Match                *AccessLists_Acl_Rule_Match `protobuf:"bytes,3,opt,name=match,proto3" json:"match,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *AccessLists_Acl_Rule) Reset()         { *m = AccessLists_Acl_Rule{} }
func (m *AccessLists_Acl_Rule) String() string { return proto.CompactTextString(m) }
func (*AccessLists_Acl_Rule) ProtoMessage()    {}
func (*AccessLists_Acl_Rule) Descriptor() ([]byte, []int) {
	return fileDescriptor_acl_700e6e99ae46a356, []int{0, 0, 0}
}
func (m *AccessLists_Acl_Rule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessLists_Acl_Rule.Unmarshal(m, b)
}
func (m *AccessLists_Acl_Rule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessLists_Acl_Rule.Marshal(b, m, deterministic)
}
func (dst *AccessLists_Acl_Rule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessLists_Acl_Rule.Merge(dst, src)
}
func (m *AccessLists_Acl_Rule) XXX_Size() int {
	return xxx_messageInfo_AccessLists_Acl_Rule.Size(m)
}
func (m *AccessLists_Acl_Rule) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessLists_Acl_Rule.DiscardUnknown(m)
}

var xxx_messageInfo_AccessLists_Acl_Rule proto.InternalMessageInfo

func (m *AccessLists_Acl_Rule) GetRuleName() string {
	if m != nil {
		return m.RuleName
	}
	return ""
}

func (m *AccessLists_Acl_Rule) GetAclAction() AclAction {
	if m != nil {
		return m.AclAction
	}
	return AclAction_DENY
}

func (m *AccessLists_Acl_Rule) GetMatch() *AccessLists_Acl_Rule_Match {
	if m != nil {
		return m.Match
	}
	return nil
}

// Definitions for match criteria for this Access List Rule
type AccessLists_Acl_Rule_Match struct {
	IpRule               *AccessLists_Acl_Rule_Match_IpRule    `protobuf:"bytes,1,opt,name=ip_rule,json=ipRule,proto3" json:"ip_rule,omitempty"`
	MacipRule            *AccessLists_Acl_Rule_Match_MacIpRule `protobuf:"bytes,2,opt,name=macip_rule,json=macipRule,proto3" json:"macip_rule,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *AccessLists_Acl_Rule_Match) Reset()         { *m = AccessLists_Acl_Rule_Match{} }
func (m *AccessLists_Acl_Rule_Match) String() string { return proto.CompactTextString(m) }
func (*AccessLists_Acl_Rule_Match) ProtoMessage()    {}
func (*AccessLists_Acl_Rule_Match) Descriptor() ([]byte, []int) {
	return fileDescriptor_acl_700e6e99ae46a356, []int{0, 0, 0, 0}
}
func (m *AccessLists_Acl_Rule_Match) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessLists_Acl_Rule_Match.Unmarshal(m, b)
}
func (m *AccessLists_Acl_Rule_Match) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessLists_Acl_Rule_Match.Marshal(b, m, deterministic)
}
func (dst *AccessLists_Acl_Rule_Match) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessLists_Acl_Rule_Match.Merge(dst, src)
}
func (m *AccessLists_Acl_Rule_Match) XXX_Size() int {
	return xxx_messageInfo_AccessLists_Acl_Rule_Match.Size(m)
}
func (m *AccessLists_Acl_Rule_Match) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessLists_Acl_Rule_Match.DiscardUnknown(m)
}

var xxx_messageInfo_AccessLists_Acl_Rule_Match proto.InternalMessageInfo

func (m *AccessLists_Acl_Rule_Match) GetIpRule() *AccessLists_Acl_Rule_Match_IpRule {
	if m != nil {
		return m.IpRule
	}
	return nil
}

func (m *AccessLists_Acl_Rule_Match) GetMacipRule() *AccessLists_Acl_Rule_Match_MacIpRule {
	if m != nil {
		return m.MacipRule
	}
	return nil
}

// Access List entry that can define:
// - IPv4/IPv6 src/dst IP prefix
// - Internet Protocol number
// - selected L4 headers:
//   * ICMP (type range)
//   * UDP (port range)
//   * TCP (port range, flags mask, flags value)
type AccessLists_Acl_Rule_Match_IpRule struct {
	Ip                   *AccessLists_Acl_Rule_Match_IpRule_Ip   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Icmp                 *AccessLists_Acl_Rule_Match_IpRule_Icmp `protobuf:"bytes,2,opt,name=icmp,proto3" json:"icmp,omitempty"`
	Tcp                  *AccessLists_Acl_Rule_Match_IpRule_Tcp  `protobuf:"bytes,3,opt,name=tcp,proto3" json:"tcp,omitempty"`
	Udp                  *AccessLists_Acl_Rule_Match_IpRule_Udp  `protobuf:"bytes,4,opt,name=udp,proto3" json:"udp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *AccessLists_Acl_Rule_Match_IpRule) Reset()         { *m = AccessLists_Acl_Rule_Match_IpRule{} }
func (m *AccessLists_Acl_Rule_Match_IpRule) String() string { return proto.CompactTextString(m) }
func (*AccessLists_Acl_Rule_Match_IpRule) ProtoMessage()    {}
func (*AccessLists_Acl_Rule_Match_IpRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_acl_700e6e99ae46a356, []int{0, 0, 0, 0, 0}
}
func (m *AccessLists_Acl_Rule_Match_IpRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule.Unmarshal(m, b)
}
func (m *AccessLists_Acl_Rule_Match_IpRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule.Marshal(b, m, deterministic)
}
func (dst *AccessLists_Acl_Rule_Match_IpRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule.Merge(dst, src)
}
func (m *AccessLists_Acl_Rule_Match_IpRule) XXX_Size() int {
	return xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule.Size(m)
}
func (m *AccessLists_Acl_Rule_Match_IpRule) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule.DiscardUnknown(m)
}

var xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule proto.InternalMessageInfo

func (m *AccessLists_Acl_Rule_Match_IpRule) GetIp() *AccessLists_Acl_Rule_Match_IpRule_Ip {
	if m != nil {
		return m.Ip
	}
	return nil
}

func (m *AccessLists_Acl_Rule_Match_IpRule) GetIcmp() *AccessLists_Acl_Rule_Match_IpRule_Icmp {
	if m != nil {
		return m.Icmp
	}
	return nil
}

func (m *AccessLists_Acl_Rule_Match_IpRule) GetTcp() *AccessLists_Acl_Rule_Match_IpRule_Tcp {
	if m != nil {
		return m.Tcp
	}
	return nil
}

func (m *AccessLists_Acl_Rule_Match_IpRule) GetUdp() *AccessLists_Acl_Rule_Match_IpRule_Udp {
	if m != nil {
		return m.Udp
	}
	return nil
}

// IP version used in this Access List Entry.
type AccessLists_Acl_Rule_Match_IpRule_Ip struct {
	// Destination IPv4/IPv6 network address (<ip>/<network>)
	DestinationNetwork string `protobuf:"bytes,1,opt,name=destination_network,json=destinationNetwork,proto3" json:"destination_network,omitempty"`
	// Destination IPv4/IPv6 network address (<ip>/<network>)
	SourceNetwork        string   `protobuf:"bytes,2,opt,name=source_network,json=sourceNetwork,proto3" json:"source_network,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Ip) Reset()         { *m = AccessLists_Acl_Rule_Match_IpRule_Ip{} }
func (m *AccessLists_Acl_Rule_Match_IpRule_Ip) String() string { return proto.CompactTextString(m) }
func (*AccessLists_Acl_Rule_Match_IpRule_Ip) ProtoMessage()    {}
func (*AccessLists_Acl_Rule_Match_IpRule_Ip) Descriptor() ([]byte, []int) {
	return fileDescriptor_acl_700e6e99ae46a356, []int{0, 0, 0, 0, 0, 0}
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Ip) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Ip.Unmarshal(m, b)
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Ip) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Ip.Marshal(b, m, deterministic)
}
func (dst *AccessLists_Acl_Rule_Match_IpRule_Ip) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Ip.Merge(dst, src)
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Ip) XXX_Size() int {
	return xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Ip.Size(m)
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Ip) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Ip.DiscardUnknown(m)
}

var xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Ip proto.InternalMessageInfo

func (m *AccessLists_Acl_Rule_Match_IpRule_Ip) GetDestinationNetwork() string {
	if m != nil {
		return m.DestinationNetwork
	}
	return ""
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Ip) GetSourceNetwork() string {
	if m != nil {
		return m.SourceNetwork
	}
	return ""
}

type AccessLists_Acl_Rule_Match_IpRule_Icmp struct {
	// ICMPv6 flag, if false ICMPv4 will be used
	Icmpv6 bool `protobuf:"varint,1,opt,name=icmpv6,proto3" json:"icmpv6,omitempty"`
	// Inclusive range representing icmp codes to be used.
	IcmpCodeRange        *AccessLists_Acl_Rule_Match_IpRule_Icmp_Range `protobuf:"bytes,2,opt,name=icmp_code_range,json=icmpCodeRange,proto3" json:"icmp_code_range,omitempty"`
	IcmpTypeRange        *AccessLists_Acl_Rule_Match_IpRule_Icmp_Range `protobuf:"bytes,3,opt,name=icmp_type_range,json=icmpTypeRange,proto3" json:"icmp_type_range,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp) Reset() {
	*m = AccessLists_Acl_Rule_Match_IpRule_Icmp{}
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp) String() string { return proto.CompactTextString(m) }
func (*AccessLists_Acl_Rule_Match_IpRule_Icmp) ProtoMessage()    {}
func (*AccessLists_Acl_Rule_Match_IpRule_Icmp) Descriptor() ([]byte, []int) {
	return fileDescriptor_acl_700e6e99ae46a356, []int{0, 0, 0, 0, 0, 1}
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Icmp.Unmarshal(m, b)
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Icmp.Marshal(b, m, deterministic)
}
func (dst *AccessLists_Acl_Rule_Match_IpRule_Icmp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Icmp.Merge(dst, src)
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp) XXX_Size() int {
	return xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Icmp.Size(m)
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Icmp.DiscardUnknown(m)
}

var xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Icmp proto.InternalMessageInfo

func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp) GetIcmpv6() bool {
	if m != nil {
		return m.Icmpv6
	}
	return false
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp) GetIcmpCodeRange() *AccessLists_Acl_Rule_Match_IpRule_Icmp_Range {
	if m != nil {
		return m.IcmpCodeRange
	}
	return nil
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp) GetIcmpTypeRange() *AccessLists_Acl_Rule_Match_IpRule_Icmp_Range {
	if m != nil {
		return m.IcmpTypeRange
	}
	return nil
}

type AccessLists_Acl_Rule_Match_IpRule_Icmp_Range struct {
	// Lower boundary for range
	First uint32 `protobuf:"varint,1,opt,name=first,proto3" json:"first,omitempty"`
	// Upper boundary for range
	Last                 uint32   `protobuf:"varint,2,opt,name=last,proto3" json:"last,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp_Range) Reset() {
	*m = AccessLists_Acl_Rule_Match_IpRule_Icmp_Range{}
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp_Range) String() string {
	return proto.CompactTextString(m)
}
func (*AccessLists_Acl_Rule_Match_IpRule_Icmp_Range) ProtoMessage() {}
func (*AccessLists_Acl_Rule_Match_IpRule_Icmp_Range) Descriptor() ([]byte, []int) {
	return fileDescriptor_acl_700e6e99ae46a356, []int{0, 0, 0, 0, 0, 1, 0}
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp_Range) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Icmp_Range.Unmarshal(m, b)
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp_Range) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Icmp_Range.Marshal(b, m, deterministic)
}
func (dst *AccessLists_Acl_Rule_Match_IpRule_Icmp_Range) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Icmp_Range.Merge(dst, src)
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp_Range) XXX_Size() int {
	return xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Icmp_Range.Size(m)
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp_Range) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Icmp_Range.DiscardUnknown(m)
}

var xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Icmp_Range proto.InternalMessageInfo

func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp_Range) GetFirst() uint32 {
	if m != nil {
		return m.First
	}
	return 0
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Icmp_Range) GetLast() uint32 {
	if m != nil {
		return m.Last
	}
	return 0
}

// Inclusive range representing destination ports to be used. When
// only lower-port is present, it represents a single port.
type AccessLists_Acl_Rule_Match_IpRule_PortRange struct {
	// Lower boundary for port.
	LowerPort uint32 `protobuf:"varint,1,opt,name=lower_port,json=lowerPort,proto3" json:"lower_port,omitempty"`
	// Upper boundary for port. If existing, the upper port must
	// be greater or equal to lower-port
	UpperPort            uint32   `protobuf:"varint,2,opt,name=upper_port,json=upperPort,proto3" json:"upper_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessLists_Acl_Rule_Match_IpRule_PortRange) Reset() {
	*m = AccessLists_Acl_Rule_Match_IpRule_PortRange{}
}
func (m *AccessLists_Acl_Rule_Match_IpRule_PortRange) String() string {
	return proto.CompactTextString(m)
}
func (*AccessLists_Acl_Rule_Match_IpRule_PortRange) ProtoMessage() {}
func (*AccessLists_Acl_Rule_Match_IpRule_PortRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_acl_700e6e99ae46a356, []int{0, 0, 0, 0, 0, 2}
}
func (m *AccessLists_Acl_Rule_Match_IpRule_PortRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_PortRange.Unmarshal(m, b)
}
func (m *AccessLists_Acl_Rule_Match_IpRule_PortRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_PortRange.Marshal(b, m, deterministic)
}
func (dst *AccessLists_Acl_Rule_Match_IpRule_PortRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_PortRange.Merge(dst, src)
}
func (m *AccessLists_Acl_Rule_Match_IpRule_PortRange) XXX_Size() int {
	return xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_PortRange.Size(m)
}
func (m *AccessLists_Acl_Rule_Match_IpRule_PortRange) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_PortRange.DiscardUnknown(m)
}

var xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_PortRange proto.InternalMessageInfo

func (m *AccessLists_Acl_Rule_Match_IpRule_PortRange) GetLowerPort() uint32 {
	if m != nil {
		return m.LowerPort
	}
	return 0
}

func (m *AccessLists_Acl_Rule_Match_IpRule_PortRange) GetUpperPort() uint32 {
	if m != nil {
		return m.UpperPort
	}
	return 0
}

type AccessLists_Acl_Rule_Match_IpRule_Tcp struct {
	DestinationPortRange *AccessLists_Acl_Rule_Match_IpRule_PortRange `protobuf:"bytes,1,opt,name=destination_port_range,json=destinationPortRange,proto3" json:"destination_port_range,omitempty"`
	SourcePortRange      *AccessLists_Acl_Rule_Match_IpRule_PortRange `protobuf:"bytes,2,opt,name=source_port_range,json=sourcePortRange,proto3" json:"source_port_range,omitempty"`
	// Binary mask for tcp flags to match. MSB order (FIN at position 0).
	// Applied as logical AND to tcp flags field of the packet being matched,
	// before it is compared with tcp-flags-value.
	TcpFlagsMask uint32 `protobuf:"varint,3,opt,name=tcp_flags_mask,json=tcpFlagsMask,proto3" json:"tcp_flags_mask,omitempty"`
	// Binary value for tcp flags to match. MSB order (FIN at position 0).
	// Before tcp-flags-value is compared with tcp flags field of the packet being matched,
	// tcp-flags-mask is applied to packet field value.
	TcpFlagsValue        uint32   `protobuf:"varint,4,opt,name=tcp_flags_value,json=tcpFlagsValue,proto3" json:"tcp_flags_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Tcp) Reset()         { *m = AccessLists_Acl_Rule_Match_IpRule_Tcp{} }
func (m *AccessLists_Acl_Rule_Match_IpRule_Tcp) String() string { return proto.CompactTextString(m) }
func (*AccessLists_Acl_Rule_Match_IpRule_Tcp) ProtoMessage()    {}
func (*AccessLists_Acl_Rule_Match_IpRule_Tcp) Descriptor() ([]byte, []int) {
	return fileDescriptor_acl_700e6e99ae46a356, []int{0, 0, 0, 0, 0, 3}
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Tcp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Tcp.Unmarshal(m, b)
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Tcp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Tcp.Marshal(b, m, deterministic)
}
func (dst *AccessLists_Acl_Rule_Match_IpRule_Tcp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Tcp.Merge(dst, src)
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Tcp) XXX_Size() int {
	return xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Tcp.Size(m)
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Tcp) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Tcp.DiscardUnknown(m)
}

var xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Tcp proto.InternalMessageInfo

func (m *AccessLists_Acl_Rule_Match_IpRule_Tcp) GetDestinationPortRange() *AccessLists_Acl_Rule_Match_IpRule_PortRange {
	if m != nil {
		return m.DestinationPortRange
	}
	return nil
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Tcp) GetSourcePortRange() *AccessLists_Acl_Rule_Match_IpRule_PortRange {
	if m != nil {
		return m.SourcePortRange
	}
	return nil
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Tcp) GetTcpFlagsMask() uint32 {
	if m != nil {
		return m.TcpFlagsMask
	}
	return 0
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Tcp) GetTcpFlagsValue() uint32 {
	if m != nil {
		return m.TcpFlagsValue
	}
	return 0
}

type AccessLists_Acl_Rule_Match_IpRule_Udp struct {
	DestinationPortRange *AccessLists_Acl_Rule_Match_IpRule_PortRange `protobuf:"bytes,1,opt,name=destination_port_range,json=destinationPortRange,proto3" json:"destination_port_range,omitempty"`
	SourcePortRange      *AccessLists_Acl_Rule_Match_IpRule_PortRange `protobuf:"bytes,2,opt,name=source_port_range,json=sourcePortRange,proto3" json:"source_port_range,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Udp) Reset()         { *m = AccessLists_Acl_Rule_Match_IpRule_Udp{} }
func (m *AccessLists_Acl_Rule_Match_IpRule_Udp) String() string { return proto.CompactTextString(m) }
func (*AccessLists_Acl_Rule_Match_IpRule_Udp) ProtoMessage()    {}
func (*AccessLists_Acl_Rule_Match_IpRule_Udp) Descriptor() ([]byte, []int) {
	return fileDescriptor_acl_700e6e99ae46a356, []int{0, 0, 0, 0, 0, 4}
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Udp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Udp.Unmarshal(m, b)
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Udp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Udp.Marshal(b, m, deterministic)
}
func (dst *AccessLists_Acl_Rule_Match_IpRule_Udp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Udp.Merge(dst, src)
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Udp) XXX_Size() int {
	return xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Udp.Size(m)
}
func (m *AccessLists_Acl_Rule_Match_IpRule_Udp) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Udp.DiscardUnknown(m)
}

var xxx_messageInfo_AccessLists_Acl_Rule_Match_IpRule_Udp proto.InternalMessageInfo

func (m *AccessLists_Acl_Rule_Match_IpRule_Udp) GetDestinationPortRange() *AccessLists_Acl_Rule_Match_IpRule_PortRange {
	if m != nil {
		return m.DestinationPortRange
	}
	return nil
}

func (m *AccessLists_Acl_Rule_Match_IpRule_Udp) GetSourcePortRange() *AccessLists_Acl_Rule_Match_IpRule_PortRange {
	if m != nil {
		return m.SourcePortRange
	}
	return nil
}

type AccessLists_Acl_Rule_Match_MacIpRule struct {
	// Source IP address.
	SourceAddress string `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	// Source IP address prefix.
	SourceAddressPrefix uint32 `protobuf:"varint,2,opt,name=source_address_prefix,json=sourceAddressPrefix,proto3" json:"source_address_prefix,omitempty"`
	// Source MAC address.
	// Before source-mac-address is compared with source mac address field of the packet
	// being matched, source-mac-address-mask is applied to packet field value.
	SourceMacAddress string `protobuf:"bytes,3,opt,name=source_mac_address,json=sourceMacAddress,proto3" json:"source_mac_address,omitempty"`
	// Source MAC address mask.
	// Applied as logical AND with source mac address field of the packet being matched,
	// before it is compared with source-mac-address.
	SourceMacAddressMask string   `protobuf:"bytes,4,opt,name=source_mac_address_mask,json=sourceMacAddressMask,proto3" json:"source_mac_address_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessLists_Acl_Rule_Match_MacIpRule) Reset()         { *m = AccessLists_Acl_Rule_Match_MacIpRule{} }
func (m *AccessLists_Acl_Rule_Match_MacIpRule) String() string { return proto.CompactTextString(m) }
func (*AccessLists_Acl_Rule_Match_MacIpRule) ProtoMessage()    {}
func (*AccessLists_Acl_Rule_Match_MacIpRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_acl_700e6e99ae46a356, []int{0, 0, 0, 0, 1}
}
func (m *AccessLists_Acl_Rule_Match_MacIpRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessLists_Acl_Rule_Match_MacIpRule.Unmarshal(m, b)
}
func (m *AccessLists_Acl_Rule_Match_MacIpRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessLists_Acl_Rule_Match_MacIpRule.Marshal(b, m, deterministic)
}
func (dst *AccessLists_Acl_Rule_Match_MacIpRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessLists_Acl_Rule_Match_MacIpRule.Merge(dst, src)
}
func (m *AccessLists_Acl_Rule_Match_MacIpRule) XXX_Size() int {
	return xxx_messageInfo_AccessLists_Acl_Rule_Match_MacIpRule.Size(m)
}
func (m *AccessLists_Acl_Rule_Match_MacIpRule) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessLists_Acl_Rule_Match_MacIpRule.DiscardUnknown(m)
}

var xxx_messageInfo_AccessLists_Acl_Rule_Match_MacIpRule proto.InternalMessageInfo

func (m *AccessLists_Acl_Rule_Match_MacIpRule) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *AccessLists_Acl_Rule_Match_MacIpRule) GetSourceAddressPrefix() uint32 {
	if m != nil {
		return m.SourceAddressPrefix
	}
	return 0
}

func (m *AccessLists_Acl_Rule_Match_MacIpRule) GetSourceMacAddress() string {
	if m != nil {
		return m.SourceMacAddress
	}
	return ""
}

func (m *AccessLists_Acl_Rule_Match_MacIpRule) GetSourceMacAddressMask() string {
	if m != nil {
		return m.SourceMacAddressMask
	}
	return ""
}

type AccessLists_Acl_Interfaces struct {
	Egress               []string `protobuf:"bytes,1,rep,name=egress,proto3" json:"egress,omitempty"`
	Ingress              []string `protobuf:"bytes,2,rep,name=ingress,proto3" json:"ingress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessLists_Acl_Interfaces) Reset()         { *m = AccessLists_Acl_Interfaces{} }
func (m *AccessLists_Acl_Interfaces) String() string { return proto.CompactTextString(m) }
func (*AccessLists_Acl_Interfaces) ProtoMessage()    {}
func (*AccessLists_Acl_Interfaces) Descriptor() ([]byte, []int) {
	return fileDescriptor_acl_700e6e99ae46a356, []int{0, 0, 1}
}
func (m *AccessLists_Acl_Interfaces) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessLists_Acl_Interfaces.Unmarshal(m, b)
}
func (m *AccessLists_Acl_Interfaces) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessLists_Acl_Interfaces.Marshal(b, m, deterministic)
}
func (dst *AccessLists_Acl_Interfaces) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessLists_Acl_Interfaces.Merge(dst, src)
}
func (m *AccessLists_Acl_Interfaces) XXX_Size() int {
	return xxx_messageInfo_AccessLists_Acl_Interfaces.Size(m)
}
func (m *AccessLists_Acl_Interfaces) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessLists_Acl_Interfaces.DiscardUnknown(m)
}

var xxx_messageInfo_AccessLists_Acl_Interfaces proto.InternalMessageInfo

func (m *AccessLists_Acl_Interfaces) GetEgress() []string {
	if m != nil {
		return m.Egress
	}
	return nil
}

func (m *AccessLists_Acl_Interfaces) GetIngress() []string {
	if m != nil {
		return m.Ingress
	}
	return nil
}

func init() {
	proto.RegisterType((*AccessLists)(nil), "acl.AccessLists")
	proto.RegisterType((*AccessLists_Acl)(nil), "acl.AccessLists.Acl")
	proto.RegisterType((*AccessLists_Acl_Rule)(nil), "acl.AccessLists.Acl.Rule")
	proto.RegisterType((*AccessLists_Acl_Rule_Match)(nil), "acl.AccessLists.Acl.Rule.Match")
	proto.RegisterType((*AccessLists_Acl_Rule_Match_IpRule)(nil), "acl.AccessLists.Acl.Rule.Match.IpRule")
	proto.RegisterType((*AccessLists_Acl_Rule_Match_IpRule_Ip)(nil), "acl.AccessLists.Acl.Rule.Match.IpRule.Ip")
	proto.RegisterType((*AccessLists_Acl_Rule_Match_IpRule_Icmp)(nil), "acl.AccessLists.Acl.Rule.Match.IpRule.Icmp")
	proto.RegisterType((*AccessLists_Acl_Rule_Match_IpRule_Icmp_Range)(nil), "acl.AccessLists.Acl.Rule.Match.IpRule.Icmp.Range")
	proto.RegisterType((*AccessLists_Acl_Rule_Match_IpRule_PortRange)(nil), "acl.AccessLists.Acl.Rule.Match.IpRule.PortRange")
	proto.RegisterType((*AccessLists_Acl_Rule_Match_IpRule_Tcp)(nil), "acl.AccessLists.Acl.Rule.Match.IpRule.Tcp")
	proto.RegisterType((*AccessLists_Acl_Rule_Match_IpRule_Udp)(nil), "acl.AccessLists.Acl.Rule.Match.IpRule.Udp")
	proto.RegisterType((*AccessLists_Acl_Rule_Match_MacIpRule)(nil), "acl.AccessLists.Acl.Rule.Match.MacIpRule")
	proto.RegisterType((*AccessLists_Acl_Interfaces)(nil), "acl.AccessLists.Acl.Interfaces")
	proto.RegisterEnum("acl.AclAction", AclAction_name, AclAction_value)
}

func init() { proto.RegisterFile("acl.proto", fileDescriptor_acl_700e6e99ae46a356) }

var fileDescriptor_acl_700e6e99ae46a356 = []byte{
	// 779 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x26, 0xb6, 0xf3, 0xe3, 0x93, 0x75, 0x12, 0x66, 0xc3, 0x92, 0x35, 0x42, 0x44, 0x08, 0x56,
	0xd9, 0x05, 0xbc, 0x34, 0xa8, 0x48, 0x48, 0x88, 0x2a, 0x2a, 0xa9, 0x88, 0xd4, 0x54, 0xd5, 0x28,
	0x45, 0xaa, 0x54, 0xc9, 0x1a, 0xc6, 0x93, 0x62, 0xd5, 0x8e, 0x47, 0x1e, 0xa7, 0xa5, 0x77, 0xdc,
	0x71, 0xc3, 0x0b, 0x20, 0xc4, 0x13, 0xf1, 0x14, 0xbc, 0x09, 0x9a, 0x1f, 0x27, 0x2e, 0x54, 0x34,
	0x85, 0x3b, 0xae, 0x3c, 0x73, 0xce, 0xf7, 0x7d, 0xe7, 0x6f, 0xc6, 0x03, 0x2e, 0xa1, 0x49, 0xc0,
	0xf3, 0xac, 0xc8, 0x90, 0x4d, 0x68, 0xf2, 0xfe, 0xaf, 0x5d, 0x68, 0x4f, 0x28, 0x65, 0x42, 0x1c,
	0xc7, 0xa2, 0x10, 0x68, 0x04, 0x0e, 0xa1, 0x89, 0x18, 0xd4, 0x86, 0xf6, 0xa8, 0x3d, 0xee, 0x07,
	0x12, 0x5e, 0xf1, 0x07, 0x13, 0x9a, 0x60, 0x85, 0xf0, 0xff, 0xe8, 0x80, 0x3d, 0xa1, 0x09, 0x7a,
	0x0e, 0x2d, 0x42, 0x93, 0x70, 0x45, 0x52, 0x36, 0xa8, 0x0d, 0x6b, 0x23, 0x17, 0x37, 0x09, 0x4d,
	0x4e, 0x48, 0xca, 0xd0, 0x6b, 0xa8, 0xe7, 0xeb, 0x84, 0x89, 0x81, 0xa5, 0xd4, 0x9e, 0xdf, 0xa7,
	0x16, 0xe0, 0x75, 0xc2, 0xb0, 0xc6, 0xa1, 0x03, 0x80, 0x78, 0x55, 0xb0, 0x7c, 0x49, 0x28, 0x13,
	0x03, 0x7b, 0x58, 0x1b, 0xb5, 0xc7, 0xef, 0xdd, 0xcb, 0x9a, 0x6d, 0x60, 0xb8, 0x42, 0xf1, 0x7f,
	0xf6, 0xc0, 0x91, 0x82, 0xe8, 0x1d, 0x70, 0xa5, 0x64, 0x35, 0xad, 0x96, 0x34, 0xa8, 0xbc, 0x3e,
	0x01, 0x90, 0x29, 0x13, 0x5a, 0xc4, 0xd9, 0x6a, 0x60, 0x0d, 0x6b, 0xa3, 0xce, 0xb8, 0x63, 0xc2,
	0x24, 0x13, 0x65, 0xc5, 0xb2, 0x51, 0x7a, 0x89, 0xf6, 0xa1, 0x9e, 0x92, 0x82, 0x7e, 0xff, 0x8f,
	0x09, 0xc9, 0xa8, 0xc1, 0x5c, 0xc2, 0xb0, 0x46, 0xfb, 0x3f, 0x3e, 0x81, 0xba, 0x32, 0xa0, 0x03,
	0x68, 0xc6, 0x3c, 0x94, 0xe1, 0x55, 0x2a, 0xed, 0xf1, 0x8b, 0x07, 0x24, 0x82, 0x19, 0x57, 0x6d,
	0x69, 0xc4, 0xea, 0x8b, 0xbe, 0x01, 0x48, 0x09, 0x2d, 0x35, 0x2c, 0xa5, 0xf1, 0xf2, 0x21, 0x8d,
	0x39, 0xa1, 0x46, 0xc6, 0x55, 0x64, 0xb9, 0xf4, 0x7f, 0x73, 0xa1, 0xa1, 0xad, 0xe8, 0x0b, 0xb0,
	0x62, 0x6e, 0x12, 0x7a, 0xb9, 0x5b, 0x42, 0xf2, 0x63, 0xc5, 0x1c, 0x1d, 0x80, 0x13, 0xd3, 0x94,
	0x9b, 0x4c, 0x3e, 0xda, 0x95, 0x4c, 0x53, 0x8e, 0x15, 0x11, 0x7d, 0x09, 0x76, 0x41, 0xb9, 0x69,
	0xe8, 0xab, 0x1d, 0xf9, 0x0b, 0xca, 0xb1, 0xa4, 0x49, 0xf6, 0x3a, 0xe2, 0x03, 0xe7, 0x51, 0xec,
	0xb3, 0x88, 0x63, 0x49, 0xf3, 0x2f, 0xc0, 0x9a, 0x71, 0xf4, 0x1a, 0x9e, 0x46, 0x4c, 0x14, 0xf1,
	0x8a, 0xc8, 0x19, 0x87, 0x2b, 0x56, 0xdc, 0x64, 0xf9, 0x95, 0x39, 0x2a, 0xa8, 0xe2, 0x3a, 0xd1,
	0x1e, 0xf4, 0x21, 0x74, 0x44, 0xb6, 0xce, 0x29, 0xdb, 0x60, 0x2d, 0x85, 0xf5, 0xb4, 0xd5, 0xc0,
	0xfc, 0x9f, 0x2c, 0x70, 0x64, 0xa1, 0xe8, 0x19, 0x34, 0x64, 0xa9, 0xd7, 0x9f, 0x2b, 0xcd, 0x16,
	0x36, 0x3b, 0x74, 0x0e, 0x5d, 0xb9, 0x0a, 0x69, 0x16, 0xb1, 0x30, 0x27, 0xab, 0xcb, 0x72, 0xa0,
	0x7b, 0x8f, 0x68, 0x63, 0x80, 0x25, 0x11, 0x7b, 0x52, 0xe9, 0x30, 0x8b, 0x98, 0xda, 0x6e, 0xa4,
	0x8b, 0x5b, 0x5e, 0x4a, 0xdb, 0xff, 0x49, 0x7a, 0x71, 0xcb, 0xb5, 0xb4, 0xbf, 0x07, 0x75, 0x1d,
	0xa3, 0x0f, 0xf5, 0x65, 0x9c, 0x8b, 0x42, 0x55, 0xe5, 0x61, 0xbd, 0x41, 0x08, 0x9c, 0x84, 0x88,
	0x42, 0x55, 0xe2, 0x61, 0xb5, 0xf6, 0x67, 0xe0, 0x9e, 0x66, 0x79, 0xa1, 0x69, 0xef, 0x02, 0x24,
	0xd9, 0x0d, 0xcb, 0x43, 0x9e, 0xe5, 0x25, 0xd7, 0x55, 0x16, 0x89, 0x91, 0xee, 0x35, 0xe7, 0xa5,
	0x5b, 0xab, 0xb8, 0xca, 0x22, 0xdd, 0xfe, 0x2f, 0x16, 0xd8, 0x0b, 0xca, 0xd1, 0x12, 0x9e, 0x55,
	0x87, 0x26, 0xc1, 0xa6, 0x4e, 0x7d, 0x8c, 0x3f, 0xdd, 0xb1, 0xce, 0x4d, 0x5e, 0xb8, 0x5f, 0xd1,
	0xdb, 0x66, 0x7b, 0x01, 0x6f, 0x9a, 0x59, 0x57, 0x42, 0x58, 0xff, 0x32, 0x44, 0x57, 0x4b, 0x6d,
	0xd5, 0x3f, 0x80, 0x4e, 0x41, 0x79, 0xb8, 0x4c, 0xc8, 0xa5, 0x08, 0x53, 0x22, 0xae, 0xd4, 0x94,
	0x3c, 0xfc, 0xa4, 0xa0, 0xfc, 0x48, 0x1a, 0xe7, 0x44, 0x5c, 0xa1, 0x17, 0xd0, 0xdd, 0xa2, 0xae,
	0x49, 0xb2, 0x66, 0xea, 0xc0, 0x7b, 0xd8, 0x2b, 0x61, 0xdf, 0x4a, 0xa3, 0xff, 0x7b, 0x0d, 0xec,
	0xb3, 0xe8, 0x7f, 0xd2, 0x1b, 0x59, 0x8d, 0xbb, 0xf9, 0x71, 0x55, 0xee, 0x1c, 0x89, 0xa2, 0x9c,
	0x09, 0x61, 0xee, 0xa7, 0xb9, 0x73, 0x13, 0x6d, 0x44, 0x63, 0x78, 0xeb, 0x2e, 0x2c, 0xe4, 0x39,
	0x5b, 0xc6, 0x3f, 0x98, 0x83, 0xf4, 0xf4, 0x0e, 0xfa, 0x54, 0xb9, 0xd0, 0xc7, 0x80, 0x0c, 0x27,
	0x25, 0x74, 0x23, 0x6f, 0x2b, 0xf9, 0x9e, 0xf6, 0xcc, 0x09, 0x2d, 0x23, 0xec, 0xc3, 0xdb, 0x7f,
	0x47, 0xeb, 0xd9, 0x39, 0x8a, 0xd2, 0xff, 0x2b, 0x45, 0xce, 0xd0, 0xff, 0x0a, 0x60, 0xfb, 0x50,
	0xc9, 0x3f, 0x02, 0xbb, 0x34, 0x55, 0xd8, 0x23, 0x17, 0x9b, 0x1d, 0x1a, 0x40, 0x33, 0x5e, 0x69,
	0x87, 0xa5, 0x1c, 0xe5, 0xf6, 0x55, 0x00, 0xee, 0xe6, 0x45, 0x42, 0x2d, 0x70, 0xbe, 0x9e, 0x9e,
	0x9c, 0xf7, 0xde, 0x40, 0x00, 0x8d, 0xd3, 0x29, 0x9e, 0xcf, 0x16, 0xbd, 0x1a, 0x6a, 0x43, 0x13,
	0x4f, 0x8f, 0x8e, 0xa7, 0x87, 0x8b, 0x9e, 0xf5, 0x5d, 0x43, 0xbd, 0xec, 0x9f, 0xfd, 0x19, 0x00,
	0x00, 0xff, 0xff, 0xa8, 0x62, 0xf9, 0x79, 0xe6, 0x07, 0x00, 0x00,
}
