// Code generated by protoc-gen-gogo.
// source: internal/internal.proto
// DO NOT EDIT!

/*
Package influxql is a generated protocol buffer package.

It is generated from these files:
	internal/internal.proto

It has these top-level messages:
	Point
	Aux
	IteratorOptions
	Measurements
	Measurement
	Interval
	IteratorStats
	VarRef
*/
package influxql

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

type Point struct {
	Name             *string        `protobuf:"bytes,1,req,name=Name" json:"Name,omitempty"`
	Tags             *string        `protobuf:"bytes,2,req,name=Tags" json:"Tags,omitempty"`
	Time             *int64         `protobuf:"varint,3,req,name=Time" json:"Time,omitempty"`
	Nil              *bool          `protobuf:"varint,4,req,name=Nil" json:"Nil,omitempty"`
	Aux              []*Aux         `protobuf:"bytes,5,rep,name=Aux" json:"Aux,omitempty"`
	Aggregated       *uint32        `protobuf:"varint,6,opt,name=Aggregated" json:"Aggregated,omitempty"`
	FloatValue       *float64       `protobuf:"fixed64,7,opt,name=FloatValue" json:"FloatValue,omitempty"`
	IntegerValue     *int64         `protobuf:"varint,8,opt,name=IntegerValue" json:"IntegerValue,omitempty"`
	StringValue      *string        `protobuf:"bytes,9,opt,name=StringValue" json:"StringValue,omitempty"`
	BooleanValue     *bool          `protobuf:"varint,10,opt,name=BooleanValue" json:"BooleanValue,omitempty"`
	Stats            *IteratorStats `protobuf:"bytes,11,opt,name=Stats" json:"Stats,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{0} }

func (m *Point) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Point) GetTags() string {
	if m != nil && m.Tags != nil {
		return *m.Tags
	}
	return ""
}

func (m *Point) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *Point) GetNil() bool {
	if m != nil && m.Nil != nil {
		return *m.Nil
	}
	return false
}

func (m *Point) GetAux() []*Aux {
	if m != nil {
		return m.Aux
	}
	return nil
}

func (m *Point) GetAggregated() uint32 {
	if m != nil && m.Aggregated != nil {
		return *m.Aggregated
	}
	return 0
}

func (m *Point) GetFloatValue() float64 {
	if m != nil && m.FloatValue != nil {
		return *m.FloatValue
	}
	return 0
}

func (m *Point) GetIntegerValue() int64 {
	if m != nil && m.IntegerValue != nil {
		return *m.IntegerValue
	}
	return 0
}

func (m *Point) GetStringValue() string {
	if m != nil && m.StringValue != nil {
		return *m.StringValue
	}
	return ""
}

func (m *Point) GetBooleanValue() bool {
	if m != nil && m.BooleanValue != nil {
		return *m.BooleanValue
	}
	return false
}

func (m *Point) GetStats() *IteratorStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

type Aux struct {
	DataType         *int32   `protobuf:"varint,1,req,name=DataType" json:"DataType,omitempty"`
	FloatValue       *float64 `protobuf:"fixed64,2,opt,name=FloatValue" json:"FloatValue,omitempty"`
	IntegerValue     *int64   `protobuf:"varint,3,opt,name=IntegerValue" json:"IntegerValue,omitempty"`
	StringValue      *string  `protobuf:"bytes,4,opt,name=StringValue" json:"StringValue,omitempty"`
	BooleanValue     *bool    `protobuf:"varint,5,opt,name=BooleanValue" json:"BooleanValue,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Aux) Reset()                    { *m = Aux{} }
func (m *Aux) String() string            { return proto.CompactTextString(m) }
func (*Aux) ProtoMessage()               {}
func (*Aux) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{1} }

func (m *Aux) GetDataType() int32 {
	if m != nil && m.DataType != nil {
		return *m.DataType
	}
	return 0
}

func (m *Aux) GetFloatValue() float64 {
	if m != nil && m.FloatValue != nil {
		return *m.FloatValue
	}
	return 0
}

func (m *Aux) GetIntegerValue() int64 {
	if m != nil && m.IntegerValue != nil {
		return *m.IntegerValue
	}
	return 0
}

func (m *Aux) GetStringValue() string {
	if m != nil && m.StringValue != nil {
		return *m.StringValue
	}
	return ""
}

func (m *Aux) GetBooleanValue() bool {
	if m != nil && m.BooleanValue != nil {
		return *m.BooleanValue
	}
	return false
}

type IteratorOptions struct {
	Expr             *string        `protobuf:"bytes,1,opt,name=Expr" json:"Expr,omitempty"`
	Aux              []string       `protobuf:"bytes,2,rep,name=Aux" json:"Aux,omitempty"`
	Fields           []*VarRef      `protobuf:"bytes,17,rep,name=Fields" json:"Fields,omitempty"`
	Sources          []*Measurement `protobuf:"bytes,3,rep,name=Sources" json:"Sources,omitempty"`
	Interval         *Interval      `protobuf:"bytes,4,opt,name=Interval" json:"Interval,omitempty"`
	Dimensions       []string       `protobuf:"bytes,5,rep,name=Dimensions" json:"Dimensions,omitempty"`
	Fill             *int32         `protobuf:"varint,6,opt,name=Fill" json:"Fill,omitempty"`
	FillValue        *float64       `protobuf:"fixed64,7,opt,name=FillValue" json:"FillValue,omitempty"`
	Condition        *string        `protobuf:"bytes,8,opt,name=Condition" json:"Condition,omitempty"`
	StartTime        *int64         `protobuf:"varint,9,opt,name=StartTime" json:"StartTime,omitempty"`
	EndTime          *int64         `protobuf:"varint,10,opt,name=EndTime" json:"EndTime,omitempty"`
	Ascending        *bool          `protobuf:"varint,11,opt,name=Ascending" json:"Ascending,omitempty"`
	Limit            *int64         `protobuf:"varint,12,opt,name=Limit" json:"Limit,omitempty"`
	Offset           *int64         `protobuf:"varint,13,opt,name=Offset" json:"Offset,omitempty"`
	SLimit           *int64         `protobuf:"varint,14,opt,name=SLimit" json:"SLimit,omitempty"`
	SOffset          *int64         `protobuf:"varint,15,opt,name=SOffset" json:"SOffset,omitempty"`
	Dedupe           *bool          `protobuf:"varint,16,opt,name=Dedupe" json:"Dedupe,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *IteratorOptions) Reset()                    { *m = IteratorOptions{} }
func (m *IteratorOptions) String() string            { return proto.CompactTextString(m) }
func (*IteratorOptions) ProtoMessage()               {}
func (*IteratorOptions) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{2} }

func (m *IteratorOptions) GetExpr() string {
	if m != nil && m.Expr != nil {
		return *m.Expr
	}
	return ""
}

func (m *IteratorOptions) GetAux() []string {
	if m != nil {
		return m.Aux
	}
	return nil
}

func (m *IteratorOptions) GetFields() []*VarRef {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *IteratorOptions) GetSources() []*Measurement {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *IteratorOptions) GetInterval() *Interval {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *IteratorOptions) GetDimensions() []string {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *IteratorOptions) GetFill() int32 {
	if m != nil && m.Fill != nil {
		return *m.Fill
	}
	return 0
}

func (m *IteratorOptions) GetFillValue() float64 {
	if m != nil && m.FillValue != nil {
		return *m.FillValue
	}
	return 0
}

func (m *IteratorOptions) GetCondition() string {
	if m != nil && m.Condition != nil {
		return *m.Condition
	}
	return ""
}

func (m *IteratorOptions) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *IteratorOptions) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

func (m *IteratorOptions) GetAscending() bool {
	if m != nil && m.Ascending != nil {
		return *m.Ascending
	}
	return false
}

func (m *IteratorOptions) GetLimit() int64 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

func (m *IteratorOptions) GetOffset() int64 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *IteratorOptions) GetSLimit() int64 {
	if m != nil && m.SLimit != nil {
		return *m.SLimit
	}
	return 0
}

func (m *IteratorOptions) GetSOffset() int64 {
	if m != nil && m.SOffset != nil {
		return *m.SOffset
	}
	return 0
}

func (m *IteratorOptions) GetDedupe() bool {
	if m != nil && m.Dedupe != nil {
		return *m.Dedupe
	}
	return false
}

type Measurements struct {
	Items            []*Measurement `protobuf:"bytes,1,rep,name=Items" json:"Items,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Measurements) Reset()                    { *m = Measurements{} }
func (m *Measurements) String() string            { return proto.CompactTextString(m) }
func (*Measurements) ProtoMessage()               {}
func (*Measurements) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{3} }

func (m *Measurements) GetItems() []*Measurement {
	if m != nil {
		return m.Items
	}
	return nil
}

type Measurement struct {
	Database         *string `protobuf:"bytes,1,opt,name=Database" json:"Database,omitempty"`
	RetentionPolicy  *string `protobuf:"bytes,2,opt,name=RetentionPolicy" json:"RetentionPolicy,omitempty"`
	Name             *string `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	Regex            *string `protobuf:"bytes,4,opt,name=Regex" json:"Regex,omitempty"`
	IsTarget         *bool   `protobuf:"varint,5,opt,name=IsTarget" json:"IsTarget,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Measurement) Reset()                    { *m = Measurement{} }
func (m *Measurement) String() string            { return proto.CompactTextString(m) }
func (*Measurement) ProtoMessage()               {}
func (*Measurement) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{4} }

func (m *Measurement) GetDatabase() string {
	if m != nil && m.Database != nil {
		return *m.Database
	}
	return ""
}

func (m *Measurement) GetRetentionPolicy() string {
	if m != nil && m.RetentionPolicy != nil {
		return *m.RetentionPolicy
	}
	return ""
}

func (m *Measurement) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Measurement) GetRegex() string {
	if m != nil && m.Regex != nil {
		return *m.Regex
	}
	return ""
}

func (m *Measurement) GetIsTarget() bool {
	if m != nil && m.IsTarget != nil {
		return *m.IsTarget
	}
	return false
}

type Interval struct {
	Duration         *int64 `protobuf:"varint,1,opt,name=Duration" json:"Duration,omitempty"`
	Offset           *int64 `protobuf:"varint,2,opt,name=Offset" json:"Offset,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Interval) Reset()                    { *m = Interval{} }
func (m *Interval) String() string            { return proto.CompactTextString(m) }
func (*Interval) ProtoMessage()               {}
func (*Interval) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{5} }

func (m *Interval) GetDuration() int64 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

func (m *Interval) GetOffset() int64 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

type IteratorStats struct {
	SeriesN          *int64 `protobuf:"varint,1,opt,name=SeriesN" json:"SeriesN,omitempty"`
	PointN           *int64 `protobuf:"varint,2,opt,name=PointN" json:"PointN,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *IteratorStats) Reset()                    { *m = IteratorStats{} }
func (m *IteratorStats) String() string            { return proto.CompactTextString(m) }
func (*IteratorStats) ProtoMessage()               {}
func (*IteratorStats) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{6} }

func (m *IteratorStats) GetSeriesN() int64 {
	if m != nil && m.SeriesN != nil {
		return *m.SeriesN
	}
	return 0
}

func (m *IteratorStats) GetPointN() int64 {
	if m != nil && m.PointN != nil {
		return *m.PointN
	}
	return 0
}

type VarRef struct {
	Val              *string `protobuf:"bytes,1,req,name=Val" json:"Val,omitempty"`
	Type             *int32  `protobuf:"varint,2,opt,name=Type" json:"Type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *VarRef) Reset()                    { *m = VarRef{} }
func (m *VarRef) String() string            { return proto.CompactTextString(m) }
func (*VarRef) ProtoMessage()               {}
func (*VarRef) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{7} }

func (m *VarRef) GetVal() string {
	if m != nil && m.Val != nil {
		return *m.Val
	}
	return ""
}

func (m *VarRef) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func init() {
	proto.RegisterType((*Point)(nil), "influxql.Point")
	proto.RegisterType((*Aux)(nil), "influxql.Aux")
	proto.RegisterType((*IteratorOptions)(nil), "influxql.IteratorOptions")
	proto.RegisterType((*Measurements)(nil), "influxql.Measurements")
	proto.RegisterType((*Measurement)(nil), "influxql.Measurement")
	proto.RegisterType((*Interval)(nil), "influxql.Interval")
	proto.RegisterType((*IteratorStats)(nil), "influxql.IteratorStats")
	proto.RegisterType((*VarRef)(nil), "influxql.VarRef")
}

func init() { proto.RegisterFile("internal/internal.proto", fileDescriptorInternal) }

var fileDescriptorInternal = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x53, 0x51, 0x8f, 0xda, 0x3c,
	0x10, 0x14, 0x84, 0x40, 0xb2, 0x81, 0x83, 0xf3, 0xf7, 0x55, 0x44, 0x7d, 0x42, 0xb4, 0xaa, 0x78,
	0xa8, 0x68, 0x85, 0xfa, 0x07, 0x68, 0x39, 0x24, 0xa4, 0x96, 0x3b, 0x01, 0xba, 0x77, 0x17, 0x4c,
	0x64, 0xc9, 0xd8, 0xd4, 0x76, 0x2a, 0xee, 0x37, 0xdf, 0x9f, 0xe8, 0xda, 0x49, 0x0e, 0x7a, 0xa2,
	0x6f, 0xd9, 0xd9, 0xb5, 0x33, 0x3b, 0x33, 0x86, 0x3e, 0x97, 0x96, 0x69, 0x49, 0xc5, 0xa7, 0xea,
	0x63, 0x7c, 0xd4, 0xca, 0x2a, 0x12, 0x71, 0xb9, 0x17, 0xf9, 0xe9, 0x97, 0x18, 0x3e, 0xd7, 0x20,
	0x7c, 0x50, 0xd8, 0x26, 0x6d, 0x68, 0x2c, 0xe9, 0x81, 0xa5, 0xb5, 0x41, 0x7d, 0x14, 0xbb, 0x6a,
	0x43, 0x33, 0x93, 0xd6, 0x5f, 0x2a, 0x8e, 0xbd, 0x00, 0xab, 0x80, 0x24, 0x10, 0x2c, 0xb9, 0x48,
	0x1b, 0x58, 0x44, 0xe4, 0x2d, 0x04, 0xd3, 0xfc, 0x94, 0x86, 0x83, 0x60, 0x94, 0x4c, 0x3a, 0xe3,
	0xea, 0xe2, 0x31, 0x82, 0x84, 0x00, 0x4c, 0xb3, 0x4c, 0xb3, 0x8c, 0x5a, 0xb6, 0x4b, 0x9b, 0x83,
	0xda, 0xa8, 0xe3, 0xb0, 0xb9, 0x50, 0xd4, 0x3e, 0x52, 0x91, 0xb3, 0xb4, 0x85, 0x58, 0x8d, 0xfc,
	0x0f, 0xed, 0x05, 0x12, 0xcc, 0x98, 0x2e, 0xd0, 0x08, 0xd1, 0x80, 0xfc, 0x07, 0xc9, 0xda, 0x6a,
	0x2e, 0xb3, 0x02, 0x8c, 0x11, 0x8c, 0xdd, 0xe8, 0x57, 0xa5, 0x04, 0xa3, 0xb2, 0x40, 0x01, 0xd1,
	0x88, 0x7c, 0x80, 0x70, 0x6d, 0xa9, 0x35, 0x69, 0x82, 0x65, 0x32, 0xe9, 0x9f, 0x69, 0x2c, 0x70,
	0x6f, 0x6a, 0x95, 0xf6, 0xed, 0xa1, 0xf0, 0x64, 0x49, 0x0f, 0xa2, 0x19, 0xb5, 0x74, 0xf3, 0x74,
	0x2c, 0xd6, 0x0d, 0x5f, 0xb1, 0xaa, 0x5f, 0x65, 0x15, 0x5c, 0x63, 0xd5, 0xb8, 0xca, 0x2a, 0x74,
	0xac, 0x86, 0xcf, 0x75, 0xe8, 0x56, 0xff, 0xbf, 0x3f, 0x5a, 0xae, 0xa4, 0x71, 0x4a, 0xde, 0x9d,
	0x8e, 0x1a, 0x7f, 0xeb, 0xce, 0x25, 0x85, 0x78, 0x75, 0x14, 0x2f, 0x26, 0x03, 0x68, 0xce, 0x39,
	0x13, 0x3b, 0x93, 0xde, 0x7a, 0x31, 0x7b, 0xe7, 0x2d, 0x1e, 0xa9, 0x5e, 0xb1, 0x3d, 0xae, 0xd9,
	0x5a, 0xab, 0x5c, 0x6f, 0x99, 0x41, 0x32, 0x6e, 0xe4, 0xcd, 0x79, 0xe4, 0x07, 0xa3, 0x26, 0xd7,
	0xec, 0xc0, 0xd0, 0xca, 0xf7, 0x10, 0x39, 0xe6, 0xfa, 0x37, 0x15, 0x9e, 0x60, 0x32, 0x21, 0x17,
	0x8a, 0x94, 0x1d, 0xb7, 0xf3, 0x0c, 0x4d, 0x95, 0xc6, 0x11, 0xf3, 0x06, 0x7a, 0xa3, 0xe7, 0x5c,
	0x08, 0xef, 0x55, 0x48, 0x6e, 0x21, 0x76, 0xd5, 0xa5, 0x55, 0x08, 0x7d, 0x53, 0x72, 0xc7, 0xdd,
	0x36, 0xde, 0xa7, 0xd8, 0x41, 0xa8, 0xae, 0xb6, 0x3e, 0x21, 0xb1, 0x17, 0xa9, 0x0b, 0xad, 0x3b,
	0xb9, 0xf3, 0x00, 0x78, 0x00, 0x67, 0xa6, 0x66, 0xcb, 0xf0, 0xa0, 0xcc, 0xbc, 0x49, 0x11, 0xe9,
	0x40, 0xf8, 0x9d, 0x1f, 0xb8, 0x4d, 0xdb, 0x7e, 0xe2, 0x06, 0x9a, 0xf7, 0xfb, 0xbd, 0x61, 0x36,
	0xed, 0x54, 0xf5, 0xba, 0xe8, 0xdf, 0x54, 0x57, 0xae, 0xcb, 0x81, 0x6e, 0x35, 0x30, 0x63, 0xbb,
	0x1c, 0x2d, 0xec, 0x79, 0xb5, 0xbf, 0x40, 0xfb, 0x42, 0x03, 0x83, 0x22, 0x84, 0x28, 0xfe, 0xc1,
	0xa0, 0xd4, 0xff, 0x96, 0x6a, 0x98, 0x41, 0x72, 0xa9, 0x5c, 0x99, 0x8c, 0x9f, 0xd4, 0xb0, 0xd2,
	0xa2, 0x3e, 0x74, 0x57, 0xcc, 0x62, 0x0f, 0x17, 0x7e, 0x50, 0x82, 0x6f, 0x9f, 0x7c, 0x3c, 0xe2,
	0x97, 0xf7, 0x12, 0xf8, 0x0a, 0xb7, 0x59, 0x61, 0x54, 0x4e, 0x65, 0x20, 0xf0, 0x9e, 0x85, 0xd9,
	0x50, 0x9d, 0x21, 0xdd, 0x22, 0x0c, 0x1f, 0xcf, 0x9e, 0xf8, 0xbf, 0xe4, 0x18, 0x0b, 0xa7, 0x61,
	0xed, 0xd5, 0xf6, 0xee, 0xf2, 0x60, 0xf8, 0x19, 0x3a, 0x7f, 0x25, 0xd7, 0xaf, 0xcf, 0x34, 0x67,
	0x66, 0x79, 0x3e, 0xe1, 0xdf, 0xed, 0xb2, 0x3c, 0xf1, 0x0e, 0x9a, 0x65, 0x4a, 0x30, 0x54, 0xe8,
	0xd8, 0xc5, 0x3b, 0x76, 0x31, 0x77, 0x43, 0xe1, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x06,
	0x51, 0x0d, 0x11, 0x04, 0x00, 0x00,
}