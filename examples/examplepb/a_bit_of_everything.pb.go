// Code generated by protoc-gen-go.
// source: examples/examplepb/a_bit_of_everything.proto
// DO NOT EDIT!

package examplepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/shilkin/grpc-gateway/third_party/googleapis/google/api"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
import grpc_gateway_examples_sub "github.com/shilkin/grpc-gateway/examples/sub"
import sub2 "github.com/shilkin/grpc-gateway/examples/sub2"
import google_protobuf2 "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// NumericEnum is one or zero.
type NumericEnum int32

const (
	// ZERO means 0
	NumericEnum_ZERO NumericEnum = 0
	// ONE means 1
	NumericEnum_ONE NumericEnum = 1
)

var NumericEnum_name = map[int32]string{
	0: "ZERO",
	1: "ONE",
}
var NumericEnum_value = map[string]int32{
	"ZERO": 0,
	"ONE":  1,
}

func (x NumericEnum) String() string {
	return proto.EnumName(NumericEnum_name, int32(x))
}
func (NumericEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// DeepEnum is one or zero.
type ABitOfEverything_Nested_DeepEnum int32

const (
	// FALSE is false.
	ABitOfEverything_Nested_FALSE ABitOfEverything_Nested_DeepEnum = 0
	// TRUE is true.
	ABitOfEverything_Nested_TRUE ABitOfEverything_Nested_DeepEnum = 1
)

var ABitOfEverything_Nested_DeepEnum_name = map[int32]string{
	0: "FALSE",
	1: "TRUE",
}
var ABitOfEverything_Nested_DeepEnum_value = map[string]int32{
	"FALSE": 0,
	"TRUE":  1,
}

func (x ABitOfEverything_Nested_DeepEnum) String() string {
	return proto.EnumName(ABitOfEverything_Nested_DeepEnum_name, int32(x))
}
func (ABitOfEverything_Nested_DeepEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{0, 0, 0}
}

// Intentionaly complicated message type to cover much features of Protobuf.
// NEXT ID: 27
type ABitOfEverything struct {
	SingleNested *ABitOfEverything_Nested   `protobuf:"bytes,25,opt,name=single_nested,json=singleNested" json:"single_nested,omitempty"`
	Uuid         string                     `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Nested       []*ABitOfEverything_Nested `protobuf:"bytes,2,rep,name=nested" json:"nested,omitempty"`
	FloatValue   float32                    `protobuf:"fixed32,3,opt,name=float_value,json=floatValue" json:"float_value,omitempty"`
	DoubleValue  float64                    `protobuf:"fixed64,4,opt,name=double_value,json=doubleValue" json:"double_value,omitempty"`
	Int64Value   int64                      `protobuf:"varint,5,opt,name=int64_value,json=int64Value" json:"int64_value,omitempty"`
	Uint64Value  uint64                     `protobuf:"varint,6,opt,name=uint64_value,json=uint64Value" json:"uint64_value,omitempty"`
	Int32Value   int32                      `protobuf:"varint,7,opt,name=int32_value,json=int32Value" json:"int32_value,omitempty"`
	Fixed64Value uint64                     `protobuf:"fixed64,8,opt,name=fixed64_value,json=fixed64Value" json:"fixed64_value,omitempty"`
	Fixed32Value uint32                     `protobuf:"fixed32,9,opt,name=fixed32_value,json=fixed32Value" json:"fixed32_value,omitempty"`
	BoolValue    bool                       `protobuf:"varint,10,opt,name=bool_value,json=boolValue" json:"bool_value,omitempty"`
	StringValue  string                     `protobuf:"bytes,11,opt,name=string_value,json=stringValue" json:"string_value,omitempty"`
	// TODO(yugui) add bytes_value
	Uint32Value         uint32      `protobuf:"varint,13,opt,name=uint32_value,json=uint32Value" json:"uint32_value,omitempty"`
	EnumValue           NumericEnum `protobuf:"varint,14,opt,name=enum_value,json=enumValue,enum=grpc.gateway.examples.examplepb.NumericEnum" json:"enum_value,omitempty"`
	Sfixed32Value       int32       `protobuf:"fixed32,15,opt,name=sfixed32_value,json=sfixed32Value" json:"sfixed32_value,omitempty"`
	Sfixed64Value       int64       `protobuf:"fixed64,16,opt,name=sfixed64_value,json=sfixed64Value" json:"sfixed64_value,omitempty"`
	Sint32Value         int32       `protobuf:"zigzag32,17,opt,name=sint32_value,json=sint32Value" json:"sint32_value,omitempty"`
	Sint64Value         int64       `protobuf:"zigzag64,18,opt,name=sint64_value,json=sint64Value" json:"sint64_value,omitempty"`
	RepeatedStringValue []string    `protobuf:"bytes,19,rep,name=repeated_string_value,json=repeatedStringValue" json:"repeated_string_value,omitempty"`
	// Types that are valid to be assigned to OneofValue:
	//	*ABitOfEverything_OneofEmpty
	//	*ABitOfEverything_OneofString
	OneofValue               isABitOfEverything_OneofValue       `protobuf_oneof:"oneof_value"`
	MapValue                 map[string]NumericEnum              `protobuf:"bytes,22,rep,name=map_value,json=mapValue" json:"map_value,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value,enum=grpc.gateway.examples.examplepb.NumericEnum"`
	MappedStringValue        map[string]string                   `protobuf:"bytes,23,rep,name=mapped_string_value,json=mappedStringValue" json:"mapped_string_value,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	MappedNestedValue        map[string]*ABitOfEverything_Nested `protobuf:"bytes,24,rep,name=mapped_nested_value,json=mappedNestedValue" json:"mapped_nested_value,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	NonConventionalNameValue string                              `protobuf:"bytes,26,opt,name=nonConventionalNameValue" json:"nonConventionalNameValue,omitempty"`
	TimestampValue           *google_protobuf2.Timestamp         `protobuf:"bytes,27,opt,name=timestamp_value,json=timestampValue" json:"timestamp_value,omitempty"`
	// repeated enum value. it is comma-separated in query
	RepeatedEnumValue []NumericEnum `protobuf:"varint,28,rep,packed,name=repeated_enum_value,json=repeatedEnumValue,enum=grpc.gateway.examples.examplepb.NumericEnum" json:"repeated_enum_value,omitempty"`
}

func (m *ABitOfEverything) Reset()                    { *m = ABitOfEverything{} }
func (m *ABitOfEverything) String() string            { return proto.CompactTextString(m) }
func (*ABitOfEverything) ProtoMessage()               {}
func (*ABitOfEverything) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type isABitOfEverything_OneofValue interface {
	isABitOfEverything_OneofValue()
}

type ABitOfEverything_OneofEmpty struct {
	OneofEmpty *google_protobuf1.Empty `protobuf:"bytes,20,opt,name=oneof_empty,json=oneofEmpty,oneof"`
}
type ABitOfEverything_OneofString struct {
	OneofString string `protobuf:"bytes,21,opt,name=oneof_string,json=oneofString,oneof"`
}

func (*ABitOfEverything_OneofEmpty) isABitOfEverything_OneofValue()  {}
func (*ABitOfEverything_OneofString) isABitOfEverything_OneofValue() {}

func (m *ABitOfEverything) GetOneofValue() isABitOfEverything_OneofValue {
	if m != nil {
		return m.OneofValue
	}
	return nil
}

func (m *ABitOfEverything) GetSingleNested() *ABitOfEverything_Nested {
	if m != nil {
		return m.SingleNested
	}
	return nil
}

func (m *ABitOfEverything) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *ABitOfEverything) GetNested() []*ABitOfEverything_Nested {
	if m != nil {
		return m.Nested
	}
	return nil
}

func (m *ABitOfEverything) GetFloatValue() float32 {
	if m != nil {
		return m.FloatValue
	}
	return 0
}

func (m *ABitOfEverything) GetDoubleValue() float64 {
	if m != nil {
		return m.DoubleValue
	}
	return 0
}

func (m *ABitOfEverything) GetInt64Value() int64 {
	if m != nil {
		return m.Int64Value
	}
	return 0
}

func (m *ABitOfEverything) GetUint64Value() uint64 {
	if m != nil {
		return m.Uint64Value
	}
	return 0
}

func (m *ABitOfEverything) GetInt32Value() int32 {
	if m != nil {
		return m.Int32Value
	}
	return 0
}

func (m *ABitOfEverything) GetFixed64Value() uint64 {
	if m != nil {
		return m.Fixed64Value
	}
	return 0
}

func (m *ABitOfEverything) GetFixed32Value() uint32 {
	if m != nil {
		return m.Fixed32Value
	}
	return 0
}

func (m *ABitOfEverything) GetBoolValue() bool {
	if m != nil {
		return m.BoolValue
	}
	return false
}

func (m *ABitOfEverything) GetStringValue() string {
	if m != nil {
		return m.StringValue
	}
	return ""
}

func (m *ABitOfEverything) GetUint32Value() uint32 {
	if m != nil {
		return m.Uint32Value
	}
	return 0
}

func (m *ABitOfEverything) GetEnumValue() NumericEnum {
	if m != nil {
		return m.EnumValue
	}
	return NumericEnum_ZERO
}

func (m *ABitOfEverything) GetSfixed32Value() int32 {
	if m != nil {
		return m.Sfixed32Value
	}
	return 0
}

func (m *ABitOfEverything) GetSfixed64Value() int64 {
	if m != nil {
		return m.Sfixed64Value
	}
	return 0
}

func (m *ABitOfEverything) GetSint32Value() int32 {
	if m != nil {
		return m.Sint32Value
	}
	return 0
}

func (m *ABitOfEverything) GetSint64Value() int64 {
	if m != nil {
		return m.Sint64Value
	}
	return 0
}

func (m *ABitOfEverything) GetRepeatedStringValue() []string {
	if m != nil {
		return m.RepeatedStringValue
	}
	return nil
}

func (m *ABitOfEverything) GetOneofEmpty() *google_protobuf1.Empty {
	if x, ok := m.GetOneofValue().(*ABitOfEverything_OneofEmpty); ok {
		return x.OneofEmpty
	}
	return nil
}

func (m *ABitOfEverything) GetOneofString() string {
	if x, ok := m.GetOneofValue().(*ABitOfEverything_OneofString); ok {
		return x.OneofString
	}
	return ""
}

func (m *ABitOfEverything) GetMapValue() map[string]NumericEnum {
	if m != nil {
		return m.MapValue
	}
	return nil
}

func (m *ABitOfEverything) GetMappedStringValue() map[string]string {
	if m != nil {
		return m.MappedStringValue
	}
	return nil
}

func (m *ABitOfEverything) GetMappedNestedValue() map[string]*ABitOfEverything_Nested {
	if m != nil {
		return m.MappedNestedValue
	}
	return nil
}

func (m *ABitOfEverything) GetNonConventionalNameValue() string {
	if m != nil {
		return m.NonConventionalNameValue
	}
	return ""
}

func (m *ABitOfEverything) GetTimestampValue() *google_protobuf2.Timestamp {
	if m != nil {
		return m.TimestampValue
	}
	return nil
}

func (m *ABitOfEverything) GetRepeatedEnumValue() []NumericEnum {
	if m != nil {
		return m.RepeatedEnumValue
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ABitOfEverything) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ABitOfEverything_OneofMarshaler, _ABitOfEverything_OneofUnmarshaler, _ABitOfEverything_OneofSizer, []interface{}{
		(*ABitOfEverything_OneofEmpty)(nil),
		(*ABitOfEverything_OneofString)(nil),
	}
}

func _ABitOfEverything_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ABitOfEverything)
	// oneof_value
	switch x := m.OneofValue.(type) {
	case *ABitOfEverything_OneofEmpty:
		b.EncodeVarint(20<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.OneofEmpty); err != nil {
			return err
		}
	case *ABitOfEverything_OneofString:
		b.EncodeVarint(21<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.OneofString)
	case nil:
	default:
		return fmt.Errorf("ABitOfEverything.OneofValue has unexpected type %T", x)
	}
	return nil
}

func _ABitOfEverything_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ABitOfEverything)
	switch tag {
	case 20: // oneof_value.oneof_empty
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf1.Empty)
		err := b.DecodeMessage(msg)
		m.OneofValue = &ABitOfEverything_OneofEmpty{msg}
		return true, err
	case 21: // oneof_value.oneof_string
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.OneofValue = &ABitOfEverything_OneofString{x}
		return true, err
	default:
		return false, nil
	}
}

func _ABitOfEverything_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ABitOfEverything)
	// oneof_value
	switch x := m.OneofValue.(type) {
	case *ABitOfEverything_OneofEmpty:
		s := proto.Size(x.OneofEmpty)
		n += proto.SizeVarint(20<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ABitOfEverything_OneofString:
		n += proto.SizeVarint(21<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.OneofString)))
		n += len(x.OneofString)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Nested is nested type.
type ABitOfEverything_Nested struct {
	// name is nested field.
	Name   string                           `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Amount uint32                           `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
	Ok     ABitOfEverything_Nested_DeepEnum `protobuf:"varint,3,opt,name=ok,enum=grpc.gateway.examples.examplepb.ABitOfEverything_Nested_DeepEnum" json:"ok,omitempty"`
}

func (m *ABitOfEverything_Nested) Reset()                    { *m = ABitOfEverything_Nested{} }
func (m *ABitOfEverything_Nested) String() string            { return proto.CompactTextString(m) }
func (*ABitOfEverything_Nested) ProtoMessage()               {}
func (*ABitOfEverything_Nested) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

func (m *ABitOfEverything_Nested) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ABitOfEverything_Nested) GetAmount() uint32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ABitOfEverything_Nested) GetOk() ABitOfEverything_Nested_DeepEnum {
	if m != nil {
		return m.Ok
	}
	return ABitOfEverything_Nested_FALSE
}

func init() {
	proto.RegisterType((*ABitOfEverything)(nil), "grpc.gateway.examples.examplepb.ABitOfEverything")
	proto.RegisterType((*ABitOfEverything_Nested)(nil), "grpc.gateway.examples.examplepb.ABitOfEverything.Nested")
	proto.RegisterEnum("grpc.gateway.examples.examplepb.NumericEnum", NumericEnum_name, NumericEnum_value)
	proto.RegisterEnum("grpc.gateway.examples.examplepb.ABitOfEverything_Nested_DeepEnum", ABitOfEverything_Nested_DeepEnum_name, ABitOfEverything_Nested_DeepEnum_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ABitOfEverythingService service

type ABitOfEverythingServiceClient interface {
	Create(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*ABitOfEverything, error)
	CreateBody(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*ABitOfEverything, error)
	Lookup(ctx context.Context, in *sub2.IdMessage, opts ...grpc.CallOption) (*ABitOfEverything, error)
	Update(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	Delete(ctx context.Context, in *sub2.IdMessage, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	GetQuery(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	Echo(ctx context.Context, in *grpc_gateway_examples_sub.StringMessage, opts ...grpc.CallOption) (*grpc_gateway_examples_sub.StringMessage, error)
	DeepPathEcho(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*ABitOfEverything, error)
	NoBindings(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	Timeout(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type aBitOfEverythingServiceClient struct {
	cc *grpc.ClientConn
}

func NewABitOfEverythingServiceClient(cc *grpc.ClientConn) ABitOfEverythingServiceClient {
	return &aBitOfEverythingServiceClient{cc}
}

func (c *aBitOfEverythingServiceClient) Create(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*ABitOfEverything, error) {
	out := new(ABitOfEverything)
	err := grpc.Invoke(ctx, "/grpc.gateway.examples.examplepb.ABitOfEverythingService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBitOfEverythingServiceClient) CreateBody(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*ABitOfEverything, error) {
	out := new(ABitOfEverything)
	err := grpc.Invoke(ctx, "/grpc.gateway.examples.examplepb.ABitOfEverythingService/CreateBody", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBitOfEverythingServiceClient) Lookup(ctx context.Context, in *sub2.IdMessage, opts ...grpc.CallOption) (*ABitOfEverything, error) {
	out := new(ABitOfEverything)
	err := grpc.Invoke(ctx, "/grpc.gateway.examples.examplepb.ABitOfEverythingService/Lookup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBitOfEverythingServiceClient) Update(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/grpc.gateway.examples.examplepb.ABitOfEverythingService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBitOfEverythingServiceClient) Delete(ctx context.Context, in *sub2.IdMessage, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/grpc.gateway.examples.examplepb.ABitOfEverythingService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBitOfEverythingServiceClient) GetQuery(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/grpc.gateway.examples.examplepb.ABitOfEverythingService/GetQuery", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBitOfEverythingServiceClient) Echo(ctx context.Context, in *grpc_gateway_examples_sub.StringMessage, opts ...grpc.CallOption) (*grpc_gateway_examples_sub.StringMessage, error) {
	out := new(grpc_gateway_examples_sub.StringMessage)
	err := grpc.Invoke(ctx, "/grpc.gateway.examples.examplepb.ABitOfEverythingService/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBitOfEverythingServiceClient) DeepPathEcho(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*ABitOfEverything, error) {
	out := new(ABitOfEverything)
	err := grpc.Invoke(ctx, "/grpc.gateway.examples.examplepb.ABitOfEverythingService/DeepPathEcho", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBitOfEverythingServiceClient) NoBindings(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/grpc.gateway.examples.examplepb.ABitOfEverythingService/NoBindings", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBitOfEverythingServiceClient) Timeout(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/grpc.gateway.examples.examplepb.ABitOfEverythingService/Timeout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ABitOfEverythingService service

type ABitOfEverythingServiceServer interface {
	Create(context.Context, *ABitOfEverything) (*ABitOfEverything, error)
	CreateBody(context.Context, *ABitOfEverything) (*ABitOfEverything, error)
	Lookup(context.Context, *sub2.IdMessage) (*ABitOfEverything, error)
	Update(context.Context, *ABitOfEverything) (*google_protobuf1.Empty, error)
	Delete(context.Context, *sub2.IdMessage) (*google_protobuf1.Empty, error)
	GetQuery(context.Context, *ABitOfEverything) (*google_protobuf1.Empty, error)
	Echo(context.Context, *grpc_gateway_examples_sub.StringMessage) (*grpc_gateway_examples_sub.StringMessage, error)
	DeepPathEcho(context.Context, *ABitOfEverything) (*ABitOfEverything, error)
	NoBindings(context.Context, *google_protobuf1.Empty) (*google_protobuf1.Empty, error)
	Timeout(context.Context, *google_protobuf1.Empty) (*google_protobuf1.Empty, error)
}

func RegisterABitOfEverythingServiceServer(s *grpc.Server, srv ABitOfEverythingServiceServer) {
	s.RegisterService(&_ABitOfEverythingService_serviceDesc, srv)
}

func _ABitOfEverythingService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ABitOfEverything)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABitOfEverythingServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.examplepb.ABitOfEverythingService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABitOfEverythingServiceServer).Create(ctx, req.(*ABitOfEverything))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABitOfEverythingService_CreateBody_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ABitOfEverything)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABitOfEverythingServiceServer).CreateBody(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.examplepb.ABitOfEverythingService/CreateBody",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABitOfEverythingServiceServer).CreateBody(ctx, req.(*ABitOfEverything))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABitOfEverythingService_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sub2.IdMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABitOfEverythingServiceServer).Lookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.examplepb.ABitOfEverythingService/Lookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABitOfEverythingServiceServer).Lookup(ctx, req.(*sub2.IdMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABitOfEverythingService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ABitOfEverything)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABitOfEverythingServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.examplepb.ABitOfEverythingService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABitOfEverythingServiceServer).Update(ctx, req.(*ABitOfEverything))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABitOfEverythingService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sub2.IdMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABitOfEverythingServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.examplepb.ABitOfEverythingService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABitOfEverythingServiceServer).Delete(ctx, req.(*sub2.IdMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABitOfEverythingService_GetQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ABitOfEverything)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABitOfEverythingServiceServer).GetQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.examplepb.ABitOfEverythingService/GetQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABitOfEverythingServiceServer).GetQuery(ctx, req.(*ABitOfEverything))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABitOfEverythingService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(grpc_gateway_examples_sub.StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABitOfEverythingServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.examplepb.ABitOfEverythingService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABitOfEverythingServiceServer).Echo(ctx, req.(*grpc_gateway_examples_sub.StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABitOfEverythingService_DeepPathEcho_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ABitOfEverything)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABitOfEverythingServiceServer).DeepPathEcho(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.examplepb.ABitOfEverythingService/DeepPathEcho",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABitOfEverythingServiceServer).DeepPathEcho(ctx, req.(*ABitOfEverything))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABitOfEverythingService_NoBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABitOfEverythingServiceServer).NoBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.examplepb.ABitOfEverythingService/NoBindings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABitOfEverythingServiceServer).NoBindings(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABitOfEverythingService_Timeout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABitOfEverythingServiceServer).Timeout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.examplepb.ABitOfEverythingService/Timeout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABitOfEverythingServiceServer).Timeout(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ABitOfEverythingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gateway.examples.examplepb.ABitOfEverythingService",
	HandlerType: (*ABitOfEverythingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ABitOfEverythingService_Create_Handler,
		},
		{
			MethodName: "CreateBody",
			Handler:    _ABitOfEverythingService_CreateBody_Handler,
		},
		{
			MethodName: "Lookup",
			Handler:    _ABitOfEverythingService_Lookup_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ABitOfEverythingService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ABitOfEverythingService_Delete_Handler,
		},
		{
			MethodName: "GetQuery",
			Handler:    _ABitOfEverythingService_GetQuery_Handler,
		},
		{
			MethodName: "Echo",
			Handler:    _ABitOfEverythingService_Echo_Handler,
		},
		{
			MethodName: "DeepPathEcho",
			Handler:    _ABitOfEverythingService_DeepPathEcho_Handler,
		},
		{
			MethodName: "NoBindings",
			Handler:    _ABitOfEverythingService_NoBindings_Handler,
		},
		{
			MethodName: "Timeout",
			Handler:    _ABitOfEverythingService_Timeout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "examples/examplepb/a_bit_of_everything.proto",
}

// Client API for AnotherServiceWithNoBindings service

type AnotherServiceWithNoBindingsClient interface {
	NoBindings(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type anotherServiceWithNoBindingsClient struct {
	cc *grpc.ClientConn
}

func NewAnotherServiceWithNoBindingsClient(cc *grpc.ClientConn) AnotherServiceWithNoBindingsClient {
	return &anotherServiceWithNoBindingsClient{cc}
}

func (c *anotherServiceWithNoBindingsClient) NoBindings(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/grpc.gateway.examples.examplepb.AnotherServiceWithNoBindings/NoBindings", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AnotherServiceWithNoBindings service

type AnotherServiceWithNoBindingsServer interface {
	NoBindings(context.Context, *google_protobuf1.Empty) (*google_protobuf1.Empty, error)
}

func RegisterAnotherServiceWithNoBindingsServer(s *grpc.Server, srv AnotherServiceWithNoBindingsServer) {
	s.RegisterService(&_AnotherServiceWithNoBindings_serviceDesc, srv)
}

func _AnotherServiceWithNoBindings_NoBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnotherServiceWithNoBindingsServer).NoBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.examplepb.AnotherServiceWithNoBindings/NoBindings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnotherServiceWithNoBindingsServer).NoBindings(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _AnotherServiceWithNoBindings_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gateway.examples.examplepb.AnotherServiceWithNoBindings",
	HandlerType: (*AnotherServiceWithNoBindingsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NoBindings",
			Handler:    _AnotherServiceWithNoBindings_NoBindings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "examples/examplepb/a_bit_of_everything.proto",
}

func init() { proto.RegisterFile("examples/examplepb/a_bit_of_everything.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 1278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x57, 0xcf, 0x6e, 0xdb, 0x46,
	0x13, 0xf7, 0x4a, 0xb6, 0x2c, 0x8d, 0x24, 0x5b, 0x5e, 0x27, 0x8e, 0xa2, 0xe4, 0x83, 0xf8, 0x29,
	0x6d, 0x41, 0xb8, 0x01, 0x89, 0x28, 0x41, 0x91, 0x18, 0x68, 0x03, 0x2b, 0x51, 0x9b, 0xa2, 0x89,
	0x92, 0xd0, 0x49, 0x0a, 0x18, 0x49, 0x05, 0x4a, 0x5a, 0x49, 0x84, 0x45, 0x2e, 0x4b, 0x2e, 0x55,
	0x0b, 0xaa, 0x7a, 0xe8, 0xa1, 0x97, 0x1e, 0x7b, 0xcf, 0xa5, 0x40, 0xd1, 0x4b, 0x8f, 0x3d, 0xb7,
	0xef, 0xd0, 0x57, 0x68, 0xdf, 0xa3, 0xe0, 0x2e, 0x49, 0x53, 0xb2, 0x05, 0xf9, 0x4f, 0x91, 0x1b,
	0x77, 0x67, 0xe6, 0x37, 0xbf, 0x99, 0xd9, 0x99, 0x5d, 0xc2, 0x4d, 0x72, 0xa8, 0x9b, 0xf6, 0x80,
	0xb8, 0x6a, 0xf0, 0x61, 0xb7, 0x54, 0xbd, 0xd9, 0x32, 0x58, 0x93, 0x76, 0x9b, 0x64, 0x48, 0x9c,
	0x11, 0xeb, 0x1b, 0x56, 0x4f, 0xb1, 0x1d, 0xca, 0x28, 0x2e, 0xf7, 0x1c, 0xbb, 0xad, 0xf4, 0x74,
	0x46, 0xbe, 0xd1, 0x47, 0x4a, 0x68, 0xaa, 0x44, 0xa6, 0xa5, 0xeb, 0x3d, 0x4a, 0x7b, 0x03, 0xa2,
	0xea, 0xb6, 0xa1, 0xea, 0x96, 0x45, 0x99, 0xce, 0x0c, 0x6a, 0xb9, 0xc2, 0xbc, 0x74, 0x2d, 0x90,
	0xf2, 0x55, 0xcb, 0xeb, 0xaa, 0xc4, 0xb4, 0xd9, 0x28, 0x10, 0x96, 0x22, 0x26, 0xae, 0xd7, 0x52,
	0x4d, 0xe2, 0xba, 0x7a, 0x8f, 0x84, 0x86, 0x71, 0x59, 0x75, 0x46, 0x58, 0x9e, 0x45, 0x65, 0x86,
	0x49, 0x5c, 0xa6, 0x9b, 0xb6, 0x50, 0xa8, 0xfc, 0xb9, 0x0e, 0x85, 0xdd, 0x9a, 0xc1, 0x9e, 0x76,
	0xeb, 0x51, 0x40, 0xf8, 0x0d, 0xe4, 0x5d, 0xc3, 0xea, 0x0d, 0x48, 0xd3, 0x22, 0x2e, 0x23, 0x9d,
	0xe2, 0x55, 0x09, 0xc9, 0xd9, 0xea, 0x5d, 0x65, 0x41, 0x88, 0xca, 0x2c, 0x92, 0xd2, 0xe0, 0xf6,
	0x5a, 0x4e, 0xc0, 0x89, 0x15, 0xc6, 0xb0, 0xec, 0x79, 0x46, 0xa7, 0x88, 0x24, 0x24, 0x67, 0x34,
	0xfe, 0x8d, 0x9f, 0x41, 0x2a, 0xf0, 0x95, 0x90, 0x92, 0x17, 0xf2, 0x15, 0xe0, 0xe0, 0x32, 0x64,
	0xbb, 0x03, 0xaa, 0xb3, 0xe6, 0x50, 0x1f, 0x78, 0xa4, 0x98, 0x94, 0x90, 0x9c, 0xd0, 0x80, 0x6f,
	0xbd, 0xf2, 0x77, 0xf0, 0xff, 0x21, 0xd7, 0xa1, 0x5e, 0x6b, 0x40, 0x02, 0x8d, 0x65, 0x09, 0xc9,
	0x48, 0xcb, 0x8a, 0x3d, 0xa1, 0x52, 0x86, 0xac, 0x61, 0xb1, 0x8f, 0xee, 0x04, 0x1a, 0x2b, 0x12,
	0x92, 0x93, 0x1a, 0xf0, 0xad, 0x08, 0xc3, 0x8b, 0x6b, 0xa4, 0x24, 0x24, 0x2f, 0x6b, 0x59, 0x2f,
	0xa6, 0x22, 0x30, 0x6e, 0x57, 0x03, 0x8d, 0x55, 0x09, 0xc9, 0x2b, 0x1c, 0xe3, 0x76, 0x55, 0x28,
	0xdc, 0x80, 0x7c, 0xd7, 0x38, 0x24, 0x9d, 0x08, 0x24, 0x2d, 0x21, 0x39, 0xa5, 0xe5, 0x82, 0xcd,
	0x69, 0xa5, 0x08, 0x27, 0x23, 0x21, 0x79, 0x35, 0x50, 0x0a, 0x91, 0xfe, 0x07, 0xd0, 0xa2, 0x74,
	0x10, 0x68, 0x80, 0x84, 0xe4, 0xb4, 0x96, 0xf1, 0x77, 0x22, 0xb2, 0x2e, 0x73, 0x0c, 0xab, 0x17,
	0x28, 0x64, 0x79, 0xfe, 0xb3, 0x62, 0x6f, 0x2a, 0x9e, 0xc8, 0x4b, 0x5e, 0x42, 0x72, 0x5e, 0xc4,
	0x13, 0x3a, 0xf9, 0x02, 0x80, 0x58, 0x9e, 0x19, 0x28, 0xac, 0x49, 0x48, 0x5e, 0xab, 0xde, 0x5c,
	0x58, 0xad, 0x86, 0x67, 0x12, 0xc7, 0x68, 0xd7, 0x2d, 0xcf, 0xd4, 0x32, 0xbe, 0xbd, 0x00, 0x7b,
	0x1f, 0xd6, 0xdc, 0xe9, 0xb8, 0xd6, 0x25, 0x24, 0xaf, 0x6b, 0x79, 0x77, 0x2a, 0xb0, 0x48, 0x2d,
	0xca, 0x51, 0x41, 0x42, 0x72, 0x21, 0x54, 0x8b, 0x55, 0xc3, 0x8d, 0xb3, 0xdf, 0x90, 0x90, 0xbc,
	0xa1, 0x65, 0xdd, 0x18, 0xfb, 0x40, 0x25, 0xc2, 0xc1, 0x12, 0x92, 0xb1, 0x50, 0x09, 0x51, 0xaa,
	0x70, 0xd9, 0x21, 0x36, 0xd1, 0x19, 0xe9, 0x34, 0xa7, 0xf2, 0xb5, 0x29, 0x25, 0xe5, 0x8c, 0xb6,
	0x19, 0x0a, 0xf7, 0x62, 0x79, 0xbb, 0x07, 0x59, 0x6a, 0x11, 0x7f, 0x2c, 0xf8, 0x5d, 0x5b, 0xbc,
	0xc4, 0xfb, 0x65, 0x4b, 0x11, 0xdd, 0xa7, 0x84, 0xdd, 0xa7, 0xd4, 0x7d, 0xe9, 0xa3, 0x25, 0x0d,
	0xb8, 0x32, 0x5f, 0xe1, 0x1b, 0x90, 0x13, 0xa6, 0xc2, 0x57, 0xf1, 0xb2, 0x5f, 0x95, 0x47, 0x4b,
	0x9a, 0x00, 0x14, 0x4e, 0xf0, 0x6b, 0xc8, 0x98, 0xba, 0x1d, 0xf0, 0xd8, 0xe2, 0x1d, 0x72, 0xff,
	0xec, 0x1d, 0xf2, 0x44, 0xb7, 0x39, 0xdd, 0xba, 0xc5, 0x9c, 0x91, 0x96, 0x36, 0x83, 0x25, 0x3e,
	0x84, 0x4d, 0x53, 0xb7, 0xed, 0xd9, 0x78, 0xaf, 0x70, 0x3f, 0x8f, 0xce, 0xe5, 0xc7, 0x9e, 0xca,
	0x8f, 0x70, 0xb8, 0x61, 0xce, 0xee, 0xc7, 0x3c, 0x8b, 0xae, 0x0d, 0x3c, 0x17, 0x2f, 0xe6, 0x59,
	0x4c, 0x82, 0xe3, 0x9e, 0x63, 0xfb, 0x78, 0x07, 0x8a, 0x16, 0xb5, 0x1e, 0x50, 0x6b, 0x48, 0x2c,
	0x7f, 0x0e, 0xeb, 0x83, 0x86, 0x6e, 0x8a, 0xb6, 0x2f, 0x96, 0x78, 0x63, 0xcc, 0x95, 0xe3, 0x07,
	0xb0, 0x1e, 0xcd, 0xd1, 0x80, 0xf1, 0x35, 0x5e, 0xf1, 0xd2, 0xb1, 0x8a, 0xbf, 0x08, 0xf5, 0xb4,
	0xb5, 0xc8, 0x44, 0x80, 0xbc, 0x86, 0xe8, 0x24, 0x35, 0x63, 0x0d, 0x75, 0x5d, 0x4a, 0x9e, 0xb9,
	0xa1, 0x36, 0x42, 0xa0, 0x7a, 0xd8, 0x58, 0xa5, 0x5f, 0x11, 0xa4, 0x8e, 0xc6, 0xad, 0xa5, 0x9b,
	0x24, 0x1c, 0xb7, 0xfe, 0x37, 0xde, 0x82, 0x94, 0x6e, 0x52, 0xcf, 0x62, 0xc5, 0x04, 0xef, 0xf0,
	0x60, 0x85, 0x9f, 0x43, 0x82, 0x1e, 0xf0, 0x59, 0xb9, 0x56, 0xdd, 0x3d, 0xef, 0x08, 0x56, 0x1e,
	0x12, 0x62, 0x73, 0x62, 0x09, 0x7a, 0x50, 0x29, 0x43, 0x3a, 0x5c, 0xe3, 0x0c, 0xac, 0x7c, 0xba,
	0xfb, 0x78, 0xaf, 0x5e, 0x58, 0xc2, 0x69, 0x58, 0x7e, 0xa1, 0xbd, 0xac, 0x17, 0x50, 0xc9, 0x80,
	0xfc, 0xd4, 0xc1, 0xc4, 0x05, 0x48, 0x1e, 0x90, 0x51, 0xc0, 0xd7, 0xff, 0xc4, 0x35, 0x58, 0x11,
	0xd9, 0x49, 0x9c, 0x63, 0xdc, 0x08, 0xd3, 0x9d, 0xc4, 0x5d, 0x54, 0x7a, 0x08, 0x5b, 0x27, 0x9f,
	0xcd, 0x13, 0x7c, 0x5e, 0x8a, 0xfb, 0xcc, 0xc4, 0x51, 0xbe, 0x0b, 0x51, 0x66, 0xcf, 0xd9, 0x09,
	0x28, 0x8d, 0x38, 0xca, 0x45, 0xae, 0xb5, 0x23, 0xff, 0xb5, 0x7c, 0x38, 0x6c, 0xf8, 0xd6, 0xb6,
	0x04, 0xd9, 0x58, 0xb8, 0x7e, 0x62, 0xf7, 0xeb, 0xda, 0xd3, 0xc2, 0x12, 0x5e, 0x85, 0xe4, 0xd3,
	0x46, 0xbd, 0x80, 0xaa, 0xff, 0xe4, 0xe0, 0xca, 0x2c, 0xee, 0x1e, 0x71, 0x86, 0x46, 0x9b, 0xe0,
	0xb7, 0x49, 0x48, 0x3d, 0x70, 0xfc, 0xd3, 0x83, 0x6f, 0x9d, 0x99, 0x5c, 0xe9, 0xec, 0x26, 0x95,
	0xdf, 0x12, 0xdf, 0xff, 0xf5, 0xf7, 0x4f, 0x89, 0x5f, 0x12, 0x95, 0x9f, 0x13, 0xea, 0xf0, 0x56,
	0xf8, 0xb6, 0x3a, 0xe9, 0x65, 0xa5, 0x8e, 0x63, 0x37, 0xf8, 0x44, 0x1d, 0xc7, 0xaf, 0xeb, 0x89,
	0x3a, 0x8e, 0xcd, 0xf1, 0x89, 0xea, 0x12, 0x5b, 0x77, 0x74, 0x46, 0x1d, 0x75, 0xec, 0x4d, 0x09,
	0xc6, 0xb1, 0x1b, 0x61, 0xa2, 0x8e, 0xa7, 0xae, 0x91, 0x70, 0x1d, 0x93, 0x1f, 0x5d, 0xa0, 0x13,
	0x75, 0x1c, 0x1f, 0x87, 0x1f, 0xbb, 0xcc, 0xb1, 0x1d, 0xd2, 0x35, 0x0e, 0xd5, 0xed, 0x89, 0x70,
	0x12, 0x33, 0x73, 0x67, 0x71, 0xdc, 0x59, 0x47, 0xee, 0x8c, 0xc1, 0x34, 0xc9, 0x79, 0xb3, 0x66,
	0x82, 0xdf, 0x22, 0x00, 0x51, 0xa0, 0x1a, 0xed, 0x8c, 0xde, 0x51, 0x91, 0xb6, 0x79, 0x8d, 0xde,
	0xab, 0x94, 0x17, 0x54, 0x68, 0x07, 0x6d, 0xe3, 0x6f, 0x21, 0xf5, 0x98, 0xd2, 0x03, 0xcf, 0xc6,
	0xeb, 0x8a, 0xff, 0x04, 0x55, 0x3e, 0xef, 0x3c, 0x11, 0x8f, 0xd0, 0xf3, 0x78, 0x56, 0xb8, 0x67,
	0x19, 0x7f, 0xb0, 0xf0, 0x6c, 0xf8, 0xef, 0xc6, 0x09, 0xfe, 0x01, 0x41, 0xea, 0xa5, 0xdd, 0x39,
	0xe7, 0xf9, 0x9d, 0x73, 0x45, 0x57, 0x6e, 0x71, 0x16, 0x1f, 0x96, 0x4e, 0xc9, 0xc2, 0x4f, 0x83,
	0x0e, 0xa9, 0x87, 0x64, 0x40, 0x18, 0x39, 0x9e, 0x86, 0x79, 0x5e, 0x82, 0x58, 0xb7, 0x4f, 0x1b,
	0xeb, 0x8f, 0x08, 0xd2, 0x9f, 0x11, 0xf6, 0xdc, 0x23, 0xce, 0xe8, 0xbf, 0x8c, 0xf6, 0x0e, 0xe7,
	0xa1, 0xe0, 0x9b, 0x8b, 0x78, 0x7c, 0xed, 0x7b, 0x0e, 0xd9, 0xfc, 0x81, 0x60, 0xb9, 0xde, 0xee,
	0x53, 0x2c, 0xcf, 0x61, 0xe2, 0x7a, 0x2d, 0x45, 0x0c, 0xda, 0x30, 0x11, 0xa7, 0xd6, 0xac, 0xb4,
	0x39, 0xa5, 0x37, 0x8b, 0x29, 0x91, 0x76, 0x9f, 0xaa, 0x63, 0xd1, 0x46, 0xfb, 0x57, 0x2b, 0x05,
	0x75, 0x58, 0x8d, 0xf4, 0x7d, 0xd9, 0x8e, 0x18, 0x9c, 0xfb, 0x18, 0x1f, 0x13, 0xe1, 0xdf, 0x11,
	0xe4, 0xfc, 0xbb, 0xe9, 0x99, 0xce, 0xfa, 0x3c, 0x92, 0x77, 0xd3, 0x5c, 0xf7, 0x79, 0x6c, 0xf7,
	0x2a, 0x77, 0x16, 0x96, 0x7d, 0xea, 0x2f, 0x4c, 0xf1, 0x6f, 0x6e, 0x7e, 0xd4, 0x3e, 0x01, 0x68,
	0xd0, 0x9a, 0x61, 0x75, 0x0c, 0xab, 0xe7, 0xe2, 0x39, 0x55, 0x9d, 0x5b, 0xed, 0x25, 0xfc, 0x0a,
	0x56, 0xfd, 0x77, 0x09, 0xf5, 0xd8, 0x99, 0x8d, 0xaf, 0x71, 0xee, 0x97, 0xf1, 0x66, 0x3c, 0x99,
	0x4c, 0x80, 0x55, 0xbf, 0x82, 0xeb, 0xbb, 0x16, 0x65, 0x7d, 0xe2, 0x04, 0xb7, 0xcb, 0x97, 0x06,
	0xeb, 0xc7, 0x98, 0x5e, 0x90, 0x77, 0x2d, 0xbb, 0x9f, 0x89, 0xd2, 0xda, 0x4a, 0x71, 0xf1, 0xed,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe6, 0xf7, 0xb8, 0x85, 0xa6, 0x0f, 0x00, 0x00,
}
