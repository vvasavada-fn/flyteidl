// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/schedule.proto

package admin // import "github.com/lyft/flyteidl/gen/pb-go/flyteidl/admin"

import proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Schedule_FixedRateUnit int32

const (
	Schedule_MINUTE Schedule_FixedRateUnit = 0
	Schedule_HOUR   Schedule_FixedRateUnit = 1
	Schedule_DAY    Schedule_FixedRateUnit = 2
)

var Schedule_FixedRateUnit_name = map[int32]string{
	0: "MINUTE",
	1: "HOUR",
	2: "DAY",
}
var Schedule_FixedRateUnit_value = map[string]int32{
	"MINUTE": 0,
	"HOUR":   1,
	"DAY":    2,
}

func (x Schedule_FixedRateUnit) String() string {
	return proto.EnumName(Schedule_FixedRateUnit_name, int32(x))
}
func (Schedule_FixedRateUnit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_schedule_c2ccf9a67c472f51, []int{0, 0}
}

type Schedule struct {
	// Types that are valid to be assigned to ScheduleExpression:
	//	*Schedule_CronExpression
	//	*Schedule_FixedRate_
	ScheduleExpression   isSchedule_ScheduleExpression `protobuf_oneof:"ScheduleExpression"`
	KickoffTimeInputArg  string                        `protobuf:"bytes,3,opt,name=kickoff_time_input_arg,json=kickoffTimeInputArg,proto3" json:"kickoff_time_input_arg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *Schedule) Reset()         { *m = Schedule{} }
func (m *Schedule) String() string { return proto.CompactTextString(m) }
func (*Schedule) ProtoMessage()    {}
func (*Schedule) Descriptor() ([]byte, []int) {
	return fileDescriptor_schedule_c2ccf9a67c472f51, []int{0}
}
func (m *Schedule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Schedule.Unmarshal(m, b)
}
func (m *Schedule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Schedule.Marshal(b, m, deterministic)
}
func (dst *Schedule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schedule.Merge(dst, src)
}
func (m *Schedule) XXX_Size() int {
	return xxx_messageInfo_Schedule.Size(m)
}
func (m *Schedule) XXX_DiscardUnknown() {
	xxx_messageInfo_Schedule.DiscardUnknown(m)
}

var xxx_messageInfo_Schedule proto.InternalMessageInfo

type isSchedule_ScheduleExpression interface {
	isSchedule_ScheduleExpression()
}

type Schedule_CronExpression struct {
	CronExpression string `protobuf:"bytes,1,opt,name=cron_expression,json=cronExpression,proto3,oneof"`
}

type Schedule_FixedRate_ struct {
	FixedRate *Schedule_FixedRate `protobuf:"bytes,2,opt,name=fixed_rate,json=fixedRate,proto3,oneof"`
}

func (*Schedule_CronExpression) isSchedule_ScheduleExpression() {}

func (*Schedule_FixedRate_) isSchedule_ScheduleExpression() {}

func (m *Schedule) GetScheduleExpression() isSchedule_ScheduleExpression {
	if m != nil {
		return m.ScheduleExpression
	}
	return nil
}

func (m *Schedule) GetCronExpression() string {
	if x, ok := m.GetScheduleExpression().(*Schedule_CronExpression); ok {
		return x.CronExpression
	}
	return ""
}

func (m *Schedule) GetFixedRate() *Schedule_FixedRate {
	if x, ok := m.GetScheduleExpression().(*Schedule_FixedRate_); ok {
		return x.FixedRate
	}
	return nil
}

func (m *Schedule) GetKickoffTimeInputArg() string {
	if m != nil {
		return m.KickoffTimeInputArg
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Schedule) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Schedule_OneofMarshaler, _Schedule_OneofUnmarshaler, _Schedule_OneofSizer, []interface{}{
		(*Schedule_CronExpression)(nil),
		(*Schedule_FixedRate_)(nil),
	}
}

func _Schedule_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Schedule)
	// ScheduleExpression
	switch x := m.ScheduleExpression.(type) {
	case *Schedule_CronExpression:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.CronExpression)
	case *Schedule_FixedRate_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FixedRate); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Schedule.ScheduleExpression has unexpected type %T", x)
	}
	return nil
}

func _Schedule_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Schedule)
	switch tag {
	case 1: // ScheduleExpression.cron_expression
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.ScheduleExpression = &Schedule_CronExpression{x}
		return true, err
	case 2: // ScheduleExpression.fixed_rate
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Schedule_FixedRate)
		err := b.DecodeMessage(msg)
		m.ScheduleExpression = &Schedule_FixedRate_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Schedule_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Schedule)
	// ScheduleExpression
	switch x := m.ScheduleExpression.(type) {
	case *Schedule_CronExpression:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.CronExpression)))
		n += len(x.CronExpression)
	case *Schedule_FixedRate_:
		s := proto.Size(x.FixedRate)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Schedule_FixedRate struct {
	Value                uint32                 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Unit                 Schedule_FixedRateUnit `protobuf:"varint,2,opt,name=unit,proto3,enum=admin.Schedule_FixedRateUnit" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Schedule_FixedRate) Reset()         { *m = Schedule_FixedRate{} }
func (m *Schedule_FixedRate) String() string { return proto.CompactTextString(m) }
func (*Schedule_FixedRate) ProtoMessage()    {}
func (*Schedule_FixedRate) Descriptor() ([]byte, []int) {
	return fileDescriptor_schedule_c2ccf9a67c472f51, []int{0, 0}
}
func (m *Schedule_FixedRate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Schedule_FixedRate.Unmarshal(m, b)
}
func (m *Schedule_FixedRate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Schedule_FixedRate.Marshal(b, m, deterministic)
}
func (dst *Schedule_FixedRate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schedule_FixedRate.Merge(dst, src)
}
func (m *Schedule_FixedRate) XXX_Size() int {
	return xxx_messageInfo_Schedule_FixedRate.Size(m)
}
func (m *Schedule_FixedRate) XXX_DiscardUnknown() {
	xxx_messageInfo_Schedule_FixedRate.DiscardUnknown(m)
}

var xxx_messageInfo_Schedule_FixedRate proto.InternalMessageInfo

func (m *Schedule_FixedRate) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Schedule_FixedRate) GetUnit() Schedule_FixedRateUnit {
	if m != nil {
		return m.Unit
	}
	return Schedule_MINUTE
}

func init() {
	proto.RegisterType((*Schedule)(nil), "admin.Schedule")
	proto.RegisterType((*Schedule_FixedRate)(nil), "admin.Schedule.FixedRate")
	proto.RegisterEnum("admin.Schedule_FixedRateUnit", Schedule_FixedRateUnit_name, Schedule_FixedRateUnit_value)
}

func init() {
	proto.RegisterFile("flyteidl/admin/schedule.proto", fileDescriptor_schedule_c2ccf9a67c472f51)
}

var fileDescriptor_schedule_c2ccf9a67c472f51 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4f, 0x4f, 0xc2, 0x40,
	0x10, 0xc5, 0x29, 0xff, 0x84, 0x31, 0x20, 0x59, 0x89, 0x41, 0x13, 0x12, 0xc2, 0x09, 0x0f, 0xb6,
	0x01, 0x6e, 0xde, 0x20, 0x62, 0xe0, 0xa0, 0x26, 0x2b, 0x1c, 0xf4, 0xd2, 0x94, 0x76, 0xb6, 0x6c,
	0x68, 0x77, 0x9b, 0x65, 0x6b, 0xe0, 0xeb, 0xf9, 0xc9, 0x4c, 0x57, 0xa8, 0xe1, 0xe0, 0x71, 0xde,
	0x9b, 0xfc, 0xe6, 0x65, 0x1e, 0x74, 0x59, 0x74, 0xd0, 0xc8, 0x83, 0xc8, 0xf1, 0x82, 0x98, 0x0b,
	0x67, 0xe7, 0x6f, 0x30, 0x48, 0x23, 0xb4, 0x13, 0x25, 0xb5, 0x24, 0x15, 0xa3, 0xf6, 0xbf, 0x8b,
	0x50, 0x7b, 0x3f, 0x3a, 0xe4, 0x1e, 0xae, 0x7c, 0x25, 0x85, 0x8b, 0xfb, 0x44, 0xe1, 0x6e, 0xc7,
	0xa5, 0xe8, 0x58, 0x3d, 0x6b, 0x50, 0x9f, 0x17, 0x68, 0x33, 0x33, 0x66, 0xb9, 0x4e, 0x1e, 0x01,
	0x18, 0xdf, 0x63, 0xe0, 0x2a, 0x4f, 0x63, 0xa7, 0xd8, 0xb3, 0x06, 0x97, 0xa3, 0x5b, 0xdb, 0x30,
	0xed, 0x13, 0xcf, 0x7e, 0xce, 0x36, 0xa8, 0xa7, 0x71, 0x5e, 0xa0, 0x75, 0x76, 0x1a, 0xc8, 0x18,
	0x6e, 0xb6, 0xdc, 0xdf, 0x4a, 0xc6, 0x5c, 0xcd, 0x63, 0x74, 0xb9, 0x48, 0x52, 0xed, 0x7a, 0x2a,
	0xec, 0x94, 0xb2, 0x6b, 0xf4, 0xfa, 0xe8, 0x2e, 0x79, 0x8c, 0x8b, 0xcc, 0x9b, 0xa8, 0xf0, 0x6e,
	0x09, 0xf5, 0x1c, 0x47, 0xda, 0x50, 0xf9, 0xf2, 0xa2, 0x14, 0x4d, 0xbc, 0x06, 0xfd, 0x1d, 0xc8,
	0x10, 0xca, 0xa9, 0xe0, 0xda, 0xa4, 0x69, 0x8e, 0xba, 0xff, 0xa6, 0x59, 0x09, 0xae, 0xa9, 0x59,
	0xed, 0xdb, 0xd0, 0x38, 0x93, 0x09, 0x40, 0xf5, 0x65, 0xf1, 0xba, 0x5a, 0xce, 0x5a, 0x05, 0x52,
	0x83, 0xf2, 0xfc, 0x6d, 0x45, 0x5b, 0x16, 0xb9, 0x80, 0xd2, 0xd3, 0xe4, 0xa3, 0x55, 0x9c, 0xb6,
	0x81, 0x9c, 0x78, 0x7f, 0xcf, 0x98, 0x8e, 0x3f, 0x87, 0x21, 0xd7, 0x9b, 0x74, 0x6d, 0xfb, 0x32,
	0x76, 0xa2, 0x03, 0xd3, 0x4e, 0xfe, 0xfc, 0x10, 0x85, 0x93, 0xac, 0x1f, 0x42, 0xe9, 0x9c, 0xf7,
	0xb1, 0xae, 0x9a, 0x1e, 0xc6, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x69, 0x89, 0x95, 0xa8,
	0x01, 0x00, 0x00,
}
