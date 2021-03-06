// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              20.01
// source: .vppapi/plugins/rdma.api.json

// Package rdma contains generated bindings for API file rdma.api.
//
// Contents:
//   1 alias
//   7 enums
//   4 messages
//
package rdma

import (
	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	"strconv"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "rdma"
	APIVersion = "1.0.0"
	VersionCrc = 0x5ce233e0
)

// IfStatusFlags defines enum 'if_status_flags'.
type IfStatusFlags uint32

const (
	IF_STATUS_API_FLAG_ADMIN_UP IfStatusFlags = 1
	IF_STATUS_API_FLAG_LINK_UP  IfStatusFlags = 2
)

var (
	IfStatusFlags_name = map[uint32]string{
		1: "IF_STATUS_API_FLAG_ADMIN_UP",
		2: "IF_STATUS_API_FLAG_LINK_UP",
	}
	IfStatusFlags_value = map[string]uint32{
		"IF_STATUS_API_FLAG_ADMIN_UP": 1,
		"IF_STATUS_API_FLAG_LINK_UP":  2,
	}
)

func (x IfStatusFlags) String() string {
	s, ok := IfStatusFlags_name[uint32(x)]
	if ok {
		return s
	}
	str := func(n uint32) string {
		s, ok := IfStatusFlags_name[uint32(n)]
		if ok {
			return s
		}
		return "IfStatusFlags(" + strconv.Itoa(int(n)) + ")"
	}
	for i := uint32(0); i <= 32; i++ {
		val := uint32(x)
		if val&(1<<i) != 0 {
			if s != "" {
				s += "|"
			}
			s += str(1 << i)
		}
	}
	if s == "" {
		return str(uint32(x))
	}
	return s
}

// IfType defines enum 'if_type'.
type IfType uint32

const (
	IF_API_TYPE_HARDWARE IfType = 1
	IF_API_TYPE_SUB      IfType = 2
	IF_API_TYPE_P2P      IfType = 3
	IF_API_TYPE_PIPE     IfType = 4
)

var (
	IfType_name = map[uint32]string{
		1: "IF_API_TYPE_HARDWARE",
		2: "IF_API_TYPE_SUB",
		3: "IF_API_TYPE_P2P",
		4: "IF_API_TYPE_PIPE",
	}
	IfType_value = map[string]uint32{
		"IF_API_TYPE_HARDWARE": 1,
		"IF_API_TYPE_SUB":      2,
		"IF_API_TYPE_P2P":      3,
		"IF_API_TYPE_PIPE":     4,
	}
)

func (x IfType) String() string {
	s, ok := IfType_name[uint32(x)]
	if ok {
		return s
	}
	return "IfType(" + strconv.Itoa(int(x)) + ")"
}

// LinkDuplex defines enum 'link_duplex'.
type LinkDuplex uint32

const (
	LINK_DUPLEX_API_UNKNOWN LinkDuplex = 0
	LINK_DUPLEX_API_HALF    LinkDuplex = 1
	LINK_DUPLEX_API_FULL    LinkDuplex = 2
)

var (
	LinkDuplex_name = map[uint32]string{
		0: "LINK_DUPLEX_API_UNKNOWN",
		1: "LINK_DUPLEX_API_HALF",
		2: "LINK_DUPLEX_API_FULL",
	}
	LinkDuplex_value = map[string]uint32{
		"LINK_DUPLEX_API_UNKNOWN": 0,
		"LINK_DUPLEX_API_HALF":    1,
		"LINK_DUPLEX_API_FULL":    2,
	}
)

func (x LinkDuplex) String() string {
	s, ok := LinkDuplex_name[uint32(x)]
	if ok {
		return s
	}
	return "LinkDuplex(" + strconv.Itoa(int(x)) + ")"
}

// MtuProto defines enum 'mtu_proto'.
type MtuProto uint32

const (
	MTU_PROTO_API_L3   MtuProto = 1
	MTU_PROTO_API_IP4  MtuProto = 2
	MTU_PROTO_API_IP6  MtuProto = 3
	MTU_PROTO_API_MPLS MtuProto = 4
	MTU_PROTO_API_N    MtuProto = 5
)

var (
	MtuProto_name = map[uint32]string{
		1: "MTU_PROTO_API_L3",
		2: "MTU_PROTO_API_IP4",
		3: "MTU_PROTO_API_IP6",
		4: "MTU_PROTO_API_MPLS",
		5: "MTU_PROTO_API_N",
	}
	MtuProto_value = map[string]uint32{
		"MTU_PROTO_API_L3":   1,
		"MTU_PROTO_API_IP4":  2,
		"MTU_PROTO_API_IP6":  3,
		"MTU_PROTO_API_MPLS": 4,
		"MTU_PROTO_API_N":    5,
	}
)

func (x MtuProto) String() string {
	s, ok := MtuProto_name[uint32(x)]
	if ok {
		return s
	}
	return "MtuProto(" + strconv.Itoa(int(x)) + ")"
}

// RdmaMode defines enum 'rdma_mode'.
type RdmaMode uint32

const (
	RDMA_API_MODE_AUTO RdmaMode = 0
	RDMA_API_MODE_IBV  RdmaMode = 1
	RDMA_API_MODE_DV   RdmaMode = 2
)

var (
	RdmaMode_name = map[uint32]string{
		0: "RDMA_API_MODE_AUTO",
		1: "RDMA_API_MODE_IBV",
		2: "RDMA_API_MODE_DV",
	}
	RdmaMode_value = map[string]uint32{
		"RDMA_API_MODE_AUTO": 0,
		"RDMA_API_MODE_IBV":  1,
		"RDMA_API_MODE_DV":   2,
	}
)

func (x RdmaMode) String() string {
	s, ok := RdmaMode_name[uint32(x)]
	if ok {
		return s
	}
	return "RdmaMode(" + strconv.Itoa(int(x)) + ")"
}

// RxMode defines enum 'rx_mode'.
type RxMode uint32

const (
	RX_MODE_API_UNKNOWN   RxMode = 0
	RX_MODE_API_POLLING   RxMode = 1
	RX_MODE_API_INTERRUPT RxMode = 2
	RX_MODE_API_ADAPTIVE  RxMode = 3
	RX_MODE_API_DEFAULT   RxMode = 4
)

var (
	RxMode_name = map[uint32]string{
		0: "RX_MODE_API_UNKNOWN",
		1: "RX_MODE_API_POLLING",
		2: "RX_MODE_API_INTERRUPT",
		3: "RX_MODE_API_ADAPTIVE",
		4: "RX_MODE_API_DEFAULT",
	}
	RxMode_value = map[string]uint32{
		"RX_MODE_API_UNKNOWN":   0,
		"RX_MODE_API_POLLING":   1,
		"RX_MODE_API_INTERRUPT": 2,
		"RX_MODE_API_ADAPTIVE":  3,
		"RX_MODE_API_DEFAULT":   4,
	}
)

func (x RxMode) String() string {
	s, ok := RxMode_name[uint32(x)]
	if ok {
		return s
	}
	return "RxMode(" + strconv.Itoa(int(x)) + ")"
}

// SubIfFlags defines enum 'sub_if_flags'.
type SubIfFlags uint32

const (
	SUB_IF_API_FLAG_NO_TAGS           SubIfFlags = 1
	SUB_IF_API_FLAG_ONE_TAG           SubIfFlags = 2
	SUB_IF_API_FLAG_TWO_TAGS          SubIfFlags = 4
	SUB_IF_API_FLAG_DOT1AD            SubIfFlags = 8
	SUB_IF_API_FLAG_EXACT_MATCH       SubIfFlags = 16
	SUB_IF_API_FLAG_DEFAULT           SubIfFlags = 32
	SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY SubIfFlags = 64
	SUB_IF_API_FLAG_INNER_VLAN_ID_ANY SubIfFlags = 128
	SUB_IF_API_FLAG_MASK_VNET         SubIfFlags = 254
	SUB_IF_API_FLAG_DOT1AH            SubIfFlags = 256
)

var (
	SubIfFlags_name = map[uint32]string{
		1:   "SUB_IF_API_FLAG_NO_TAGS",
		2:   "SUB_IF_API_FLAG_ONE_TAG",
		4:   "SUB_IF_API_FLAG_TWO_TAGS",
		8:   "SUB_IF_API_FLAG_DOT1AD",
		16:  "SUB_IF_API_FLAG_EXACT_MATCH",
		32:  "SUB_IF_API_FLAG_DEFAULT",
		64:  "SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY",
		128: "SUB_IF_API_FLAG_INNER_VLAN_ID_ANY",
		254: "SUB_IF_API_FLAG_MASK_VNET",
		256: "SUB_IF_API_FLAG_DOT1AH",
	}
	SubIfFlags_value = map[string]uint32{
		"SUB_IF_API_FLAG_NO_TAGS":           1,
		"SUB_IF_API_FLAG_ONE_TAG":           2,
		"SUB_IF_API_FLAG_TWO_TAGS":          4,
		"SUB_IF_API_FLAG_DOT1AD":            8,
		"SUB_IF_API_FLAG_EXACT_MATCH":       16,
		"SUB_IF_API_FLAG_DEFAULT":           32,
		"SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY": 64,
		"SUB_IF_API_FLAG_INNER_VLAN_ID_ANY": 128,
		"SUB_IF_API_FLAG_MASK_VNET":         254,
		"SUB_IF_API_FLAG_DOT1AH":            256,
	}
)

func (x SubIfFlags) String() string {
	s, ok := SubIfFlags_name[uint32(x)]
	if ok {
		return s
	}
	str := func(n uint32) string {
		s, ok := SubIfFlags_name[uint32(n)]
		if ok {
			return s
		}
		return "SubIfFlags(" + strconv.Itoa(int(n)) + ")"
	}
	for i := uint32(0); i <= 32; i++ {
		val := uint32(x)
		if val&(1<<i) != 0 {
			if s != "" {
				s += "|"
			}
			s += str(1 << i)
		}
	}
	if s == "" {
		return str(uint32(x))
	}
	return s
}

// InterfaceIndex defines alias 'interface_index'.
type InterfaceIndex uint32

// RdmaCreate defines message 'rdma_create'.
type RdmaCreate struct {
	HostIf  string   `binapi:"string[64],name=host_if" json:"host_if,omitempty"`
	Name    string   `binapi:"string[64],name=name" json:"name,omitempty"`
	RxqNum  uint16   `binapi:"u16,name=rxq_num,default=1" json:"rxq_num,omitempty"`
	RxqSize uint16   `binapi:"u16,name=rxq_size,default=1024" json:"rxq_size,omitempty"`
	TxqSize uint16   `binapi:"u16,name=txq_size,default=1024" json:"txq_size,omitempty"`
	Mode    RdmaMode `binapi:"rdma_mode,name=mode,default=0" json:"mode,omitempty"`
}

func (m *RdmaCreate) Reset()               { *m = RdmaCreate{} }
func (*RdmaCreate) GetMessageName() string { return "rdma_create" }
func (*RdmaCreate) GetCrcString() string   { return "076fe418" }
func (*RdmaCreate) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *RdmaCreate) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 64 // m.HostIf
	size += 64 // m.Name
	size += 2  // m.RxqNum
	size += 2  // m.RxqSize
	size += 2  // m.TxqSize
	size += 4  // m.Mode
	return size
}
func (m *RdmaCreate) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeString(m.HostIf, 64)
	buf.EncodeString(m.Name, 64)
	buf.EncodeUint16(m.RxqNum)
	buf.EncodeUint16(m.RxqSize)
	buf.EncodeUint16(m.TxqSize)
	buf.EncodeUint32(uint32(m.Mode))
	return buf.Bytes(), nil
}
func (m *RdmaCreate) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.HostIf = buf.DecodeString(64)
	m.Name = buf.DecodeString(64)
	m.RxqNum = buf.DecodeUint16()
	m.RxqSize = buf.DecodeUint16()
	m.TxqSize = buf.DecodeUint16()
	m.Mode = RdmaMode(buf.DecodeUint32())
	return nil
}

// RdmaCreateReply defines message 'rdma_create_reply'.
type RdmaCreateReply struct {
	Retval    int32          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *RdmaCreateReply) Reset()               { *m = RdmaCreateReply{} }
func (*RdmaCreateReply) GetMessageName() string { return "rdma_create_reply" }
func (*RdmaCreateReply) GetCrcString() string   { return "5383d31f" }
func (*RdmaCreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *RdmaCreateReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *RdmaCreateReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *RdmaCreateReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = InterfaceIndex(buf.DecodeUint32())
	return nil
}

// RdmaDelete defines message 'rdma_delete'.
type RdmaDelete struct {
	SwIfIndex InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *RdmaDelete) Reset()               { *m = RdmaDelete{} }
func (*RdmaDelete) GetMessageName() string { return "rdma_delete" }
func (*RdmaDelete) GetCrcString() string   { return "f9e6675e" }
func (*RdmaDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *RdmaDelete) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *RdmaDelete) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *RdmaDelete) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = InterfaceIndex(buf.DecodeUint32())
	return nil
}

// RdmaDeleteReply defines message 'rdma_delete_reply'.
type RdmaDeleteReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *RdmaDeleteReply) Reset()               { *m = RdmaDeleteReply{} }
func (*RdmaDeleteReply) GetMessageName() string { return "rdma_delete_reply" }
func (*RdmaDeleteReply) GetCrcString() string   { return "e8d4e804" }
func (*RdmaDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *RdmaDeleteReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *RdmaDeleteReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *RdmaDeleteReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_rdma_binapi_init() }
func file_rdma_binapi_init() {
	api.RegisterMessage((*RdmaCreate)(nil), "rdma_create_076fe418")
	api.RegisterMessage((*RdmaCreateReply)(nil), "rdma_create_reply_5383d31f")
	api.RegisterMessage((*RdmaDelete)(nil), "rdma_delete_f9e6675e")
	api.RegisterMessage((*RdmaDeleteReply)(nil), "rdma_delete_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*RdmaCreate)(nil),
		(*RdmaCreateReply)(nil),
		(*RdmaDelete)(nil),
		(*RdmaDeleteReply)(nil),
	}
}
