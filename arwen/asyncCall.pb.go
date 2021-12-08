// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: asyncCall.proto

package arwen

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strconv "strconv"
	strings "strings"
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

type SerializableAsyncCallStatus int32

const (
	SerializableAsyncCallPending  SerializableAsyncCallStatus = 0
	SerializableAsyncCallResolved SerializableAsyncCallStatus = 1
	SerializableAsyncCallRejected SerializableAsyncCallStatus = 2
)

var SerializableAsyncCallStatus_name = map[int32]string{
	0: "SerializableAsyncCallPending",
	1: "SerializableAsyncCallResolved",
	2: "SerializableAsyncCallRejected",
}

var SerializableAsyncCallStatus_value = map[string]int32{
	"SerializableAsyncCallPending":  0,
	"SerializableAsyncCallResolved": 1,
	"SerializableAsyncCallRejected": 2,
}

func (SerializableAsyncCallStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a0e9b586d6e1f667, []int{0}
}

type SerializableAsyncCallExecutionMode int32

const (
	SerializableSyncExecution              SerializableAsyncCallExecutionMode = 0
	SerializableAsyncBuiltinFuncIntraShard SerializableAsyncCallExecutionMode = 1
	SerializableAsyncBuiltinFuncCrossShard SerializableAsyncCallExecutionMode = 2
	SerializableESDTTransferOnCallBack     SerializableAsyncCallExecutionMode = 3
	SerializableAsyncUnknown               SerializableAsyncCallExecutionMode = 4
)

var SerializableAsyncCallExecutionMode_name = map[int32]string{
	0: "SerializableSyncExecution",
	1: "SerializableAsyncBuiltinFuncIntraShard",
	2: "SerializableAsyncBuiltinFuncCrossShard",
	3: "SerializableESDTTransferOnCallBack",
	4: "SerializableAsyncUnknown",
}

var SerializableAsyncCallExecutionMode_value = map[string]int32{
	"SerializableSyncExecution":              0,
	"SerializableAsyncBuiltinFuncIntraShard": 1,
	"SerializableAsyncBuiltinFuncCrossShard": 2,
	"SerializableESDTTransferOnCallBack":     3,
	"SerializableAsyncUnknown":               4,
}

func (SerializableAsyncCallExecutionMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a0e9b586d6e1f667, []int{1}
}

type SerializableAsyncCall struct {
	CallID          []byte                             `protobuf:"bytes,1,opt,name=CallID,proto3" json:"CallID,omitempty"`
	Status          SerializableAsyncCallStatus        `protobuf:"varint,2,opt,name=Status,proto3,enum=arwen.SerializableAsyncCallStatus" json:"Status,omitempty"`
	ExecutionMode   SerializableAsyncCallExecutionMode `protobuf:"varint,3,opt,name=ExecutionMode,proto3,enum=arwen.SerializableAsyncCallExecutionMode" json:"ExecutionMode,omitempty"`
	Destination     []byte                             `protobuf:"bytes,5,opt,name=Destination,proto3" json:"Destination,omitempty"`
	Data            []byte                             `protobuf:"bytes,6,opt,name=Data,proto3" json:"Data,omitempty"`
	GasLimit        uint64                             `protobuf:"varint,7,opt,name=GasLimit,proto3" json:"GasLimit,omitempty"`
	GasLocked       uint64                             `protobuf:"varint,8,opt,name=GasLocked,proto3" json:"GasLocked,omitempty"`
	ValueBytes      []byte                             `protobuf:"bytes,9,opt,name=ValueBytes,proto3" json:"ValueBytes,omitempty"`
	SuccessCallback string                             `protobuf:"bytes,10,opt,name=SuccessCallback,proto3" json:"SuccessCallback,omitempty"`
	ErrorCallback   string                             `protobuf:"bytes,11,opt,name=ErrorCallback,proto3" json:"ErrorCallback,omitempty"`
}

func (m *SerializableAsyncCall) Reset()      { *m = SerializableAsyncCall{} }
func (*SerializableAsyncCall) ProtoMessage() {}
func (*SerializableAsyncCall) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0e9b586d6e1f667, []int{0}
}
func (m *SerializableAsyncCall) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SerializableAsyncCall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *SerializableAsyncCall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SerializableAsyncCall.Merge(m, src)
}
func (m *SerializableAsyncCall) XXX_Size() int {
	return m.Size()
}
func (m *SerializableAsyncCall) XXX_DiscardUnknown() {
	xxx_messageInfo_SerializableAsyncCall.DiscardUnknown(m)
}

var xxx_messageInfo_SerializableAsyncCall proto.InternalMessageInfo

func (m *SerializableAsyncCall) GetCallID() []byte {
	if m != nil {
		return m.CallID
	}
	return nil
}

func (m *SerializableAsyncCall) GetStatus() SerializableAsyncCallStatus {
	if m != nil {
		return m.Status
	}
	return SerializableAsyncCallPending
}

func (m *SerializableAsyncCall) GetExecutionMode() SerializableAsyncCallExecutionMode {
	if m != nil {
		return m.ExecutionMode
	}
	return SerializableSyncExecution
}

func (m *SerializableAsyncCall) GetDestination() []byte {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *SerializableAsyncCall) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SerializableAsyncCall) GetGasLimit() uint64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *SerializableAsyncCall) GetGasLocked() uint64 {
	if m != nil {
		return m.GasLocked
	}
	return 0
}

func (m *SerializableAsyncCall) GetValueBytes() []byte {
	if m != nil {
		return m.ValueBytes
	}
	return nil
}

func (m *SerializableAsyncCall) GetSuccessCallback() string {
	if m != nil {
		return m.SuccessCallback
	}
	return ""
}

func (m *SerializableAsyncCall) GetErrorCallback() string {
	if m != nil {
		return m.ErrorCallback
	}
	return ""
}

type SerializableAsyncCallGroup struct {
	Callback     string                   `protobuf:"bytes,1,opt,name=Callback,proto3" json:"Callback,omitempty"`
	GasLocked    uint64                   `protobuf:"varint,2,opt,name=GasLocked,proto3" json:"GasLocked,omitempty"`
	CallbackData []byte                   `protobuf:"bytes,3,opt,name=CallbackData,proto3" json:"CallbackData,omitempty"`
	Identifier   string                   `protobuf:"bytes,4,opt,name=Identifier,proto3" json:"Identifier,omitempty"`
	AsyncCalls   []*SerializableAsyncCall `protobuf:"bytes,5,rep,name=AsyncCalls,proto3" json:"AsyncCalls,omitempty"`
}

func (m *SerializableAsyncCallGroup) Reset()      { *m = SerializableAsyncCallGroup{} }
func (*SerializableAsyncCallGroup) ProtoMessage() {}
func (*SerializableAsyncCallGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0e9b586d6e1f667, []int{1}
}
func (m *SerializableAsyncCallGroup) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SerializableAsyncCallGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *SerializableAsyncCallGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SerializableAsyncCallGroup.Merge(m, src)
}
func (m *SerializableAsyncCallGroup) XXX_Size() int {
	return m.Size()
}
func (m *SerializableAsyncCallGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_SerializableAsyncCallGroup.DiscardUnknown(m)
}

var xxx_messageInfo_SerializableAsyncCallGroup proto.InternalMessageInfo

func (m *SerializableAsyncCallGroup) GetCallback() string {
	if m != nil {
		return m.Callback
	}
	return ""
}

func (m *SerializableAsyncCallGroup) GetGasLocked() uint64 {
	if m != nil {
		return m.GasLocked
	}
	return 0
}

func (m *SerializableAsyncCallGroup) GetCallbackData() []byte {
	if m != nil {
		return m.CallbackData
	}
	return nil
}

func (m *SerializableAsyncCallGroup) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *SerializableAsyncCallGroup) GetAsyncCalls() []*SerializableAsyncCall {
	if m != nil {
		return m.AsyncCalls
	}
	return nil
}

func init() {
	proto.RegisterEnum("arwen.SerializableAsyncCallStatus", SerializableAsyncCallStatus_name, SerializableAsyncCallStatus_value)
	proto.RegisterEnum("arwen.SerializableAsyncCallExecutionMode", SerializableAsyncCallExecutionMode_name, SerializableAsyncCallExecutionMode_value)
	proto.RegisterType((*SerializableAsyncCall)(nil), "arwen.SerializableAsyncCall")
	proto.RegisterType((*SerializableAsyncCallGroup)(nil), "arwen.SerializableAsyncCallGroup")
}

func init() { proto.RegisterFile("asyncCall.proto", fileDescriptor_a0e9b586d6e1f667) }

var fileDescriptor_a0e9b586d6e1f667 = []byte{
	// 569 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x41, 0x4f, 0x13, 0x41,
	0x18, 0xdd, 0x69, 0x4b, 0x85, 0x0f, 0x90, 0x66, 0x12, 0xcd, 0x88, 0x65, 0xb2, 0x36, 0x86, 0x54,
	0x12, 0x4b, 0x82, 0x37, 0x63, 0x62, 0x2c, 0x45, 0x42, 0xa2, 0xc1, 0xb4, 0xe8, 0xc1, 0xdb, 0x74,
	0x77, 0x28, 0x63, 0x97, 0x19, 0x32, 0x33, 0x2b, 0xe2, 0xc9, 0x8b, 0x77, 0x7f, 0x80, 0x3f, 0xc0,
	0x9f, 0xe2, 0x91, 0x63, 0x8f, 0x76, 0x7b, 0xf1, 0xc8, 0x4f, 0x30, 0x3b, 0x85, 0xb5, 0x0b, 0x4d,
	0x39, 0xed, 0x7c, 0xef, 0x7b, 0xef, 0xcd, 0xfb, 0x66, 0x27, 0x03, 0x2b, 0xcc, 0x9c, 0xc9, 0x60,
	0x9b, 0x45, 0x51, 0xe3, 0x44, 0x2b, 0xab, 0xf0, 0x1c, 0xd3, 0xa7, 0x5c, 0xae, 0x3e, 0xed, 0x09,
	0x7b, 0x14, 0x77, 0x1b, 0x81, 0x3a, 0xde, 0xec, 0xa9, 0x9e, 0xda, 0x74, 0xdd, 0x6e, 0x7c, 0xe8,
	0x2a, 0x57, 0xb8, 0xd5, 0x58, 0x55, 0xfb, 0x59, 0x84, 0x7b, 0x1d, 0xae, 0x05, 0x8b, 0xc4, 0x57,
	0xd6, 0x8d, 0xf8, 0xab, 0x2b, 0x57, 0x7c, 0x1f, 0xca, 0xe9, 0x77, 0xaf, 0x45, 0x90, 0x8f, 0xea,
	0x4b, 0xed, 0xcb, 0x0a, 0x3f, 0x87, 0x72, 0xc7, 0x32, 0x1b, 0x1b, 0x52, 0xf0, 0x51, 0xfd, 0xee,
	0x56, 0xad, 0xe1, 0x36, 0x6e, 0x4c, 0x75, 0x19, 0x33, 0xdb, 0x97, 0x0a, 0xbc, 0x0f, 0xcb, 0x3b,
	0x5f, 0x78, 0x10, 0x5b, 0xa1, 0xe4, 0x5b, 0x15, 0x72, 0x52, 0x74, 0x16, 0x4f, 0x66, 0x59, 0xe4,
	0x04, 0xed, 0xbc, 0x1e, 0xfb, 0xb0, 0xd8, 0xe2, 0xc6, 0x0a, 0xc9, 0x52, 0x88, 0xcc, 0xb9, 0xa4,
	0x93, 0x10, 0xc6, 0x50, 0x6a, 0x31, 0xcb, 0x48, 0xd9, 0xb5, 0xdc, 0x1a, 0xaf, 0xc2, 0xfc, 0x2e,
	0x33, 0x6f, 0xc4, 0xb1, 0xb0, 0xe4, 0x8e, 0x8f, 0xea, 0xa5, 0x76, 0x56, 0xe3, 0x2a, 0x2c, 0xa4,
	0x6b, 0x15, 0xf4, 0x79, 0x48, 0xe6, 0x5d, 0xf3, 0x3f, 0x80, 0x29, 0xc0, 0x07, 0x16, 0xc5, 0xbc,
	0x79, 0x66, 0xb9, 0x21, 0x0b, 0xce, 0x73, 0x02, 0xc1, 0x75, 0x58, 0xe9, 0xc4, 0x41, 0xc0, 0x8d,
	0x49, 0xa3, 0x77, 0x59, 0xd0, 0x27, 0xe0, 0xa3, 0xfa, 0x42, 0xfb, 0x3a, 0x8c, 0x1f, 0xc3, 0xf2,
	0x8e, 0xd6, 0x4a, 0x67, 0xbc, 0x45, 0xc7, 0xcb, 0x83, 0xb5, 0x01, 0x82, 0xd5, 0xa9, 0xa7, 0xb2,
	0xab, 0x55, 0x7c, 0x92, 0x0e, 0x92, 0xe9, 0x91, 0xd3, 0x67, 0x75, 0x7e, 0x90, 0xc2, 0xf5, 0x41,
	0x6a, 0xb0, 0x74, 0xc5, 0x74, 0xc7, 0x53, 0x74, 0xa3, 0xe4, 0xb0, 0x74, 0xd8, 0xbd, 0x90, 0x4b,
	0x2b, 0x0e, 0x05, 0xd7, 0xa4, 0xe4, 0xfc, 0x27, 0x10, 0xfc, 0x02, 0x20, 0xcb, 0x63, 0xc8, 0x9c,
	0x5f, 0xac, 0x2f, 0x6e, 0x55, 0x67, 0xfd, 0xca, 0xf6, 0x04, 0x7f, 0xe3, 0x3b, 0x82, 0x87, 0x33,
	0xee, 0x0c, 0xf6, 0xa1, 0x3a, 0xb5, 0xfd, 0x8e, 0xcb, 0x50, 0xc8, 0x5e, 0xc5, 0xc3, 0x8f, 0x60,
	0x6d, 0xfa, 0x36, 0xdc, 0xa8, 0xe8, 0x33, 0x0f, 0x2b, 0x68, 0x06, 0xe5, 0x13, 0x0f, 0x2c, 0x0f,
	0x2b, 0x85, 0x8d, 0x21, 0x82, 0xda, 0xed, 0x17, 0x0f, 0xaf, 0xc1, 0x83, 0x49, 0x56, 0xe7, 0x4c,
	0x06, 0x19, 0xa1, 0xe2, 0xe1, 0x0d, 0x58, 0xbf, 0x61, 0xd2, 0x8c, 0x45, 0x64, 0x85, 0x7c, 0x1d,
	0xcb, 0x60, 0x4f, 0x5a, 0xcd, 0x3a, 0x47, 0x4c, 0xa7, 0xa1, 0x6e, 0xe1, 0x6e, 0x6b, 0x65, 0xcc,
	0x98, 0x5b, 0xc0, 0xeb, 0xf9, 0x70, 0x3b, 0x9d, 0xd6, 0xc1, 0x81, 0x66, 0xd2, 0x1c, 0x72, 0xbd,
	0x2f, 0xd3, 0x94, 0x4d, 0x16, 0xf4, 0x2b, 0x45, 0x5c, 0x05, 0x72, 0xc3, 0xf3, 0xbd, 0xec, 0x4b,
	0x75, 0x2a, 0x2b, 0xa5, 0xe6, 0xcb, 0xf3, 0x21, 0xf5, 0x06, 0x43, 0xea, 0x5d, 0x0c, 0x29, 0xfa,
	0x96, 0x50, 0xf4, 0x2b, 0xa1, 0xe8, 0x77, 0x42, 0xd1, 0x79, 0x42, 0xd1, 0x20, 0xa1, 0xe8, 0x4f,
	0x42, 0xd1, 0xdf, 0x84, 0x7a, 0x17, 0x09, 0x45, 0x3f, 0x46, 0xd4, 0x3b, 0x1f, 0x51, 0x6f, 0x30,
	0xa2, 0xde, 0xc7, 0xf1, 0xab, 0xd2, 0x2d, 0xbb, 0xd7, 0xe2, 0xd9, 0xbf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x2a, 0x8a, 0x2d, 0xf9, 0x76, 0x04, 0x00, 0x00,
}

func (x SerializableAsyncCallStatus) String() string {
	s, ok := SerializableAsyncCallStatus_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x SerializableAsyncCallExecutionMode) String() string {
	s, ok := SerializableAsyncCallExecutionMode_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *SerializableAsyncCall) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SerializableAsyncCall)
	if !ok {
		that2, ok := that.(SerializableAsyncCall)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.CallID, that1.CallID) {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	if this.ExecutionMode != that1.ExecutionMode {
		return false
	}
	if !bytes.Equal(this.Destination, that1.Destination) {
		return false
	}
	if !bytes.Equal(this.Data, that1.Data) {
		return false
	}
	if this.GasLimit != that1.GasLimit {
		return false
	}
	if this.GasLocked != that1.GasLocked {
		return false
	}
	if !bytes.Equal(this.ValueBytes, that1.ValueBytes) {
		return false
	}
	if this.SuccessCallback != that1.SuccessCallback {
		return false
	}
	if this.ErrorCallback != that1.ErrorCallback {
		return false
	}
	return true
}
func (this *SerializableAsyncCallGroup) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SerializableAsyncCallGroup)
	if !ok {
		that2, ok := that.(SerializableAsyncCallGroup)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Callback != that1.Callback {
		return false
	}
	if this.GasLocked != that1.GasLocked {
		return false
	}
	if !bytes.Equal(this.CallbackData, that1.CallbackData) {
		return false
	}
	if this.Identifier != that1.Identifier {
		return false
	}
	if len(this.AsyncCalls) != len(that1.AsyncCalls) {
		return false
	}
	for i := range this.AsyncCalls {
		if !this.AsyncCalls[i].Equal(that1.AsyncCalls[i]) {
			return false
		}
	}
	return true
}
func (this *SerializableAsyncCall) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 14)
	s = append(s, "&arwen.SerializableAsyncCall{")
	s = append(s, "CallID: "+fmt.Sprintf("%#v", this.CallID)+",\n")
	s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	s = append(s, "ExecutionMode: "+fmt.Sprintf("%#v", this.ExecutionMode)+",\n")
	s = append(s, "Destination: "+fmt.Sprintf("%#v", this.Destination)+",\n")
	s = append(s, "Data: "+fmt.Sprintf("%#v", this.Data)+",\n")
	s = append(s, "GasLimit: "+fmt.Sprintf("%#v", this.GasLimit)+",\n")
	s = append(s, "GasLocked: "+fmt.Sprintf("%#v", this.GasLocked)+",\n")
	s = append(s, "ValueBytes: "+fmt.Sprintf("%#v", this.ValueBytes)+",\n")
	s = append(s, "SuccessCallback: "+fmt.Sprintf("%#v", this.SuccessCallback)+",\n")
	s = append(s, "ErrorCallback: "+fmt.Sprintf("%#v", this.ErrorCallback)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SerializableAsyncCallGroup) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&arwen.SerializableAsyncCallGroup{")
	s = append(s, "Callback: "+fmt.Sprintf("%#v", this.Callback)+",\n")
	s = append(s, "GasLocked: "+fmt.Sprintf("%#v", this.GasLocked)+",\n")
	s = append(s, "CallbackData: "+fmt.Sprintf("%#v", this.CallbackData)+",\n")
	s = append(s, "Identifier: "+fmt.Sprintf("%#v", this.Identifier)+",\n")
	if this.AsyncCalls != nil {
		s = append(s, "AsyncCalls: "+fmt.Sprintf("%#v", this.AsyncCalls)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringAsyncCall(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *SerializableAsyncCall) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SerializableAsyncCall) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SerializableAsyncCall) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ErrorCallback) > 0 {
		i -= len(m.ErrorCallback)
		copy(dAtA[i:], m.ErrorCallback)
		i = encodeVarintAsyncCall(dAtA, i, uint64(len(m.ErrorCallback)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.SuccessCallback) > 0 {
		i -= len(m.SuccessCallback)
		copy(dAtA[i:], m.SuccessCallback)
		i = encodeVarintAsyncCall(dAtA, i, uint64(len(m.SuccessCallback)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.ValueBytes) > 0 {
		i -= len(m.ValueBytes)
		copy(dAtA[i:], m.ValueBytes)
		i = encodeVarintAsyncCall(dAtA, i, uint64(len(m.ValueBytes)))
		i--
		dAtA[i] = 0x4a
	}
	if m.GasLocked != 0 {
		i = encodeVarintAsyncCall(dAtA, i, uint64(m.GasLocked))
		i--
		dAtA[i] = 0x40
	}
	if m.GasLimit != 0 {
		i = encodeVarintAsyncCall(dAtA, i, uint64(m.GasLimit))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintAsyncCall(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Destination) > 0 {
		i -= len(m.Destination)
		copy(dAtA[i:], m.Destination)
		i = encodeVarintAsyncCall(dAtA, i, uint64(len(m.Destination)))
		i--
		dAtA[i] = 0x2a
	}
	if m.ExecutionMode != 0 {
		i = encodeVarintAsyncCall(dAtA, i, uint64(m.ExecutionMode))
		i--
		dAtA[i] = 0x18
	}
	if m.Status != 0 {
		i = encodeVarintAsyncCall(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if len(m.CallID) > 0 {
		i -= len(m.CallID)
		copy(dAtA[i:], m.CallID)
		i = encodeVarintAsyncCall(dAtA, i, uint64(len(m.CallID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SerializableAsyncCallGroup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SerializableAsyncCallGroup) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SerializableAsyncCallGroup) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AsyncCalls) > 0 {
		for iNdEx := len(m.AsyncCalls) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AsyncCalls[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAsyncCall(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Identifier) > 0 {
		i -= len(m.Identifier)
		copy(dAtA[i:], m.Identifier)
		i = encodeVarintAsyncCall(dAtA, i, uint64(len(m.Identifier)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.CallbackData) > 0 {
		i -= len(m.CallbackData)
		copy(dAtA[i:], m.CallbackData)
		i = encodeVarintAsyncCall(dAtA, i, uint64(len(m.CallbackData)))
		i--
		dAtA[i] = 0x1a
	}
	if m.GasLocked != 0 {
		i = encodeVarintAsyncCall(dAtA, i, uint64(m.GasLocked))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Callback) > 0 {
		i -= len(m.Callback)
		copy(dAtA[i:], m.Callback)
		i = encodeVarintAsyncCall(dAtA, i, uint64(len(m.Callback)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAsyncCall(dAtA []byte, offset int, v uint64) int {
	offset -= sovAsyncCall(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SerializableAsyncCall) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CallID)
	if l > 0 {
		n += 1 + l + sovAsyncCall(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovAsyncCall(uint64(m.Status))
	}
	if m.ExecutionMode != 0 {
		n += 1 + sovAsyncCall(uint64(m.ExecutionMode))
	}
	l = len(m.Destination)
	if l > 0 {
		n += 1 + l + sovAsyncCall(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovAsyncCall(uint64(l))
	}
	if m.GasLimit != 0 {
		n += 1 + sovAsyncCall(uint64(m.GasLimit))
	}
	if m.GasLocked != 0 {
		n += 1 + sovAsyncCall(uint64(m.GasLocked))
	}
	l = len(m.ValueBytes)
	if l > 0 {
		n += 1 + l + sovAsyncCall(uint64(l))
	}
	l = len(m.SuccessCallback)
	if l > 0 {
		n += 1 + l + sovAsyncCall(uint64(l))
	}
	l = len(m.ErrorCallback)
	if l > 0 {
		n += 1 + l + sovAsyncCall(uint64(l))
	}
	return n
}

func (m *SerializableAsyncCallGroup) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Callback)
	if l > 0 {
		n += 1 + l + sovAsyncCall(uint64(l))
	}
	if m.GasLocked != 0 {
		n += 1 + sovAsyncCall(uint64(m.GasLocked))
	}
	l = len(m.CallbackData)
	if l > 0 {
		n += 1 + l + sovAsyncCall(uint64(l))
	}
	l = len(m.Identifier)
	if l > 0 {
		n += 1 + l + sovAsyncCall(uint64(l))
	}
	if len(m.AsyncCalls) > 0 {
		for _, e := range m.AsyncCalls {
			l = e.Size()
			n += 1 + l + sovAsyncCall(uint64(l))
		}
	}
	return n
}

func sovAsyncCall(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAsyncCall(x uint64) (n int) {
	return sovAsyncCall(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *SerializableAsyncCall) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SerializableAsyncCall{`,
		`CallID:` + fmt.Sprintf("%v", this.CallID) + `,`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`ExecutionMode:` + fmt.Sprintf("%v", this.ExecutionMode) + `,`,
		`Destination:` + fmt.Sprintf("%v", this.Destination) + `,`,
		`Data:` + fmt.Sprintf("%v", this.Data) + `,`,
		`GasLimit:` + fmt.Sprintf("%v", this.GasLimit) + `,`,
		`GasLocked:` + fmt.Sprintf("%v", this.GasLocked) + `,`,
		`ValueBytes:` + fmt.Sprintf("%v", this.ValueBytes) + `,`,
		`SuccessCallback:` + fmt.Sprintf("%v", this.SuccessCallback) + `,`,
		`ErrorCallback:` + fmt.Sprintf("%v", this.ErrorCallback) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SerializableAsyncCallGroup) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForAsyncCalls := "[]*SerializableAsyncCall{"
	for _, f := range this.AsyncCalls {
		repeatedStringForAsyncCalls += strings.Replace(f.String(), "SerializableAsyncCall", "SerializableAsyncCall", 1) + ","
	}
	repeatedStringForAsyncCalls += "}"
	s := strings.Join([]string{`&SerializableAsyncCallGroup{`,
		`Callback:` + fmt.Sprintf("%v", this.Callback) + `,`,
		`GasLocked:` + fmt.Sprintf("%v", this.GasLocked) + `,`,
		`CallbackData:` + fmt.Sprintf("%v", this.CallbackData) + `,`,
		`Identifier:` + fmt.Sprintf("%v", this.Identifier) + `,`,
		`AsyncCalls:` + repeatedStringForAsyncCalls + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringAsyncCall(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *SerializableAsyncCall) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAsyncCall
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
			return fmt.Errorf("proto: SerializableAsyncCall: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SerializableAsyncCall: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsyncCall
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAsyncCall
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAsyncCall
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CallID = append(m.CallID[:0], dAtA[iNdEx:postIndex]...)
			if m.CallID == nil {
				m.CallID = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsyncCall
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= SerializableAsyncCallStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutionMode", wireType)
			}
			m.ExecutionMode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsyncCall
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExecutionMode |= SerializableAsyncCallExecutionMode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Destination", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsyncCall
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAsyncCall
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAsyncCall
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Destination = append(m.Destination[:0], dAtA[iNdEx:postIndex]...)
			if m.Destination == nil {
				m.Destination = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsyncCall
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAsyncCall
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAsyncCall
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasLimit", wireType)
			}
			m.GasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsyncCall
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasLocked", wireType)
			}
			m.GasLocked = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsyncCall
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasLocked |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValueBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsyncCall
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAsyncCall
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAsyncCall
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValueBytes = append(m.ValueBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.ValueBytes == nil {
				m.ValueBytes = []byte{}
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuccessCallback", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsyncCall
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
				return ErrInvalidLengthAsyncCall
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsyncCall
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SuccessCallback = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorCallback", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsyncCall
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
				return ErrInvalidLengthAsyncCall
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsyncCall
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrorCallback = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAsyncCall(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAsyncCall
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAsyncCall
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SerializableAsyncCallGroup) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAsyncCall
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
			return fmt.Errorf("proto: SerializableAsyncCallGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SerializableAsyncCallGroup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Callback", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsyncCall
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
				return ErrInvalidLengthAsyncCall
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsyncCall
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Callback = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasLocked", wireType)
			}
			m.GasLocked = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsyncCall
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasLocked |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallbackData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsyncCall
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAsyncCall
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAsyncCall
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CallbackData = append(m.CallbackData[:0], dAtA[iNdEx:postIndex]...)
			if m.CallbackData == nil {
				m.CallbackData = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsyncCall
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
				return ErrInvalidLengthAsyncCall
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsyncCall
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AsyncCalls", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsyncCall
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
				return ErrInvalidLengthAsyncCall
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAsyncCall
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AsyncCalls = append(m.AsyncCalls, &SerializableAsyncCall{})
			if err := m.AsyncCalls[len(m.AsyncCalls)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAsyncCall(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAsyncCall
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAsyncCall
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAsyncCall(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAsyncCall
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
					return 0, ErrIntOverflowAsyncCall
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
					return 0, ErrIntOverflowAsyncCall
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
				return 0, ErrInvalidLengthAsyncCall
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAsyncCall
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAsyncCall
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAsyncCall        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAsyncCall          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAsyncCall = fmt.Errorf("proto: unexpected end of group")
)