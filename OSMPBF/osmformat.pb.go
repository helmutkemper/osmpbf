// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmformat.proto

package OSMPBF

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Relation_MemberType int32

const (
	Relation_NODE     Relation_MemberType = 0
	Relation_WAY      Relation_MemberType = 1
	Relation_RELATION Relation_MemberType = 2
)

var Relation_MemberType_name = map[int32]string{
	0: "NODE",
	1: "WAY",
	2: "RELATION",
}
var Relation_MemberType_value = map[string]int32{
	"NODE":     0,
	"WAY":      1,
	"RELATION": 2,
}

func (x Relation_MemberType) Enum() *Relation_MemberType {
	p := new(Relation_MemberType)
	*p = x
	return p
}
func (x Relation_MemberType) String() string {
	return proto.EnumName(Relation_MemberType_name, int32(x))
}
func (x *Relation_MemberType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Relation_MemberType_value, data, "Relation_MemberType")
	if err != nil {
		return err
	}
	*x = Relation_MemberType(value)
	return nil
}
func (Relation_MemberType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorOsmformat, []int{11, 0}
}

type HeaderBlock struct {
	Bbox *HeaderBBox `protobuf:"bytes,1,opt,name=bbox" json:"bbox,omitempty"`
	// Additional tags to aid in parsing this dataset
	RequiredFeatures []string `protobuf:"bytes,4,rep,name=required_features,json=requiredFeatures" json:"required_features,omitempty"`
	OptionalFeatures []string `protobuf:"bytes,5,rep,name=optional_features,json=optionalFeatures" json:"optional_features,omitempty"`
	Writingprogram   *string  `protobuf:"bytes,16,opt,name=writingprogram" json:"writingprogram,omitempty"`
	Source           *string  `protobuf:"bytes,17,opt,name=source" json:"source,omitempty"`
	// replication timestamp, expressed in seconds since the epoch,
	// otherwise the same value as in the "timestamp=..." field
	// in the state.txt file used by Osmosis
	OsmosisReplicationTimestamp *int64 `protobuf:"varint,32,opt,name=osmosis_replication_timestamp,json=osmosisReplicationTimestamp" json:"osmosis_replication_timestamp,omitempty"`
	// replication sequence number (sequenceNumber in state.txt)
	OsmosisReplicationSequenceNumber *int64 `protobuf:"varint,33,opt,name=osmosis_replication_sequence_number,json=osmosisReplicationSequenceNumber" json:"osmosis_replication_sequence_number,omitempty"`
	// replication base URL (from Osmosis' configuration.txt file)
	OsmosisReplicationBaseUrl *string `protobuf:"bytes,34,opt,name=osmosis_replication_base_url,json=osmosisReplicationBaseUrl" json:"osmosis_replication_base_url,omitempty"`
	XXX_unrecognized          []byte  `json:"-"`
}

func (m *HeaderBlock) Reset()                    { *m = HeaderBlock{} }
func (m *HeaderBlock) String() string            { return proto.CompactTextString(m) }
func (*HeaderBlock) ProtoMessage()               {}
func (*HeaderBlock) Descriptor() ([]byte, []int) { return fileDescriptorOsmformat, []int{0} }

func (m *HeaderBlock) GetBbox() *HeaderBBox {
	if m != nil {
		return m.Bbox
	}
	return nil
}

func (m *HeaderBlock) GetRequiredFeatures() []string {
	if m != nil {
		return m.RequiredFeatures
	}
	return nil
}

func (m *HeaderBlock) GetOptionalFeatures() []string {
	if m != nil {
		return m.OptionalFeatures
	}
	return nil
}

func (m *HeaderBlock) GetWritingprogram() string {
	if m != nil && m.Writingprogram != nil {
		return *m.Writingprogram
	}
	return ""
}

func (m *HeaderBlock) GetSource() string {
	if m != nil && m.Source != nil {
		return *m.Source
	}
	return ""
}

func (m *HeaderBlock) GetOsmosisReplicationTimestamp() int64 {
	if m != nil && m.OsmosisReplicationTimestamp != nil {
		return *m.OsmosisReplicationTimestamp
	}
	return 0
}

func (m *HeaderBlock) GetOsmosisReplicationSequenceNumber() int64 {
	if m != nil && m.OsmosisReplicationSequenceNumber != nil {
		return *m.OsmosisReplicationSequenceNumber
	}
	return 0
}

func (m *HeaderBlock) GetOsmosisReplicationBaseUrl() string {
	if m != nil && m.OsmosisReplicationBaseUrl != nil {
		return *m.OsmosisReplicationBaseUrl
	}
	return ""
}

type HeaderBBox struct {
	Left             *int64 `protobuf:"zigzag64,1,req,name=left" json:"left,omitempty"`
	Right            *int64 `protobuf:"zigzag64,2,req,name=right" json:"right,omitempty"`
	Top              *int64 `protobuf:"zigzag64,3,req,name=top" json:"top,omitempty"`
	Bottom           *int64 `protobuf:"zigzag64,4,req,name=bottom" json:"bottom,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *HeaderBBox) Reset()                    { *m = HeaderBBox{} }
func (m *HeaderBBox) String() string            { return proto.CompactTextString(m) }
func (*HeaderBBox) ProtoMessage()               {}
func (*HeaderBBox) Descriptor() ([]byte, []int) { return fileDescriptorOsmformat, []int{1} }

func (m *HeaderBBox) GetLeft() int64 {
	if m != nil && m.Left != nil {
		return *m.Left
	}
	return 0
}

func (m *HeaderBBox) GetRight() int64 {
	if m != nil && m.Right != nil {
		return *m.Right
	}
	return 0
}

func (m *HeaderBBox) GetTop() int64 {
	if m != nil && m.Top != nil {
		return *m.Top
	}
	return 0
}

func (m *HeaderBBox) GetBottom() int64 {
	if m != nil && m.Bottom != nil {
		return *m.Bottom
	}
	return 0
}

type PrimitiveBlock struct {
	Stringtable    *StringTable      `protobuf:"bytes,1,req,name=stringtable" json:"stringtable,omitempty"`
	Primitivegroup []*PrimitiveGroup `protobuf:"bytes,2,rep,name=primitivegroup" json:"primitivegroup,omitempty"`
	// Granularity, units of nanodegrees, used to store coordinates in this block
	Granularity *int32 `protobuf:"varint,17,opt,name=granularity,def=100" json:"granularity,omitempty"`
	// Offset value between the output coordinates coordinates and the granularity grid in unites of nanodegrees.
	LatOffset *int64 `protobuf:"varint,19,opt,name=lat_offset,json=latOffset,def=0" json:"lat_offset,omitempty"`
	LonOffset *int64 `protobuf:"varint,20,opt,name=lon_offset,json=lonOffset,def=0" json:"lon_offset,omitempty"`
	// Granularity of dates, normally represented in units of milliseconds since the 1970 epoch.
	DateGranularity  *int32 `protobuf:"varint,18,opt,name=date_granularity,json=dateGranularity,def=1000" json:"date_granularity,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PrimitiveBlock) Reset()                    { *m = PrimitiveBlock{} }
func (m *PrimitiveBlock) String() string            { return proto.CompactTextString(m) }
func (*PrimitiveBlock) ProtoMessage()               {}
func (*PrimitiveBlock) Descriptor() ([]byte, []int) { return fileDescriptorOsmformat, []int{2} }

const Default_PrimitiveBlock_Granularity int32 = 100
const Default_PrimitiveBlock_LatOffset int64 = 0
const Default_PrimitiveBlock_LonOffset int64 = 0
const Default_PrimitiveBlock_DateGranularity int32 = 1000

func (m *PrimitiveBlock) GetStringtable() *StringTable {
	if m != nil {
		return m.Stringtable
	}
	return nil
}

func (m *PrimitiveBlock) GetPrimitivegroup() []*PrimitiveGroup {
	if m != nil {
		return m.Primitivegroup
	}
	return nil
}

func (m *PrimitiveBlock) GetGranularity() int32 {
	if m != nil && m.Granularity != nil {
		return *m.Granularity
	}
	return Default_PrimitiveBlock_Granularity
}

func (m *PrimitiveBlock) GetLatOffset() int64 {
	if m != nil && m.LatOffset != nil {
		return *m.LatOffset
	}
	return Default_PrimitiveBlock_LatOffset
}

func (m *PrimitiveBlock) GetLonOffset() int64 {
	if m != nil && m.LonOffset != nil {
		return *m.LonOffset
	}
	return Default_PrimitiveBlock_LonOffset
}

func (m *PrimitiveBlock) GetDateGranularity() int32 {
	if m != nil && m.DateGranularity != nil {
		return *m.DateGranularity
	}
	return Default_PrimitiveBlock_DateGranularity
}

// Group of OSMPrimitives. All primitives in a group must be the same type.
type PrimitiveGroup struct {
	Nodes            []*Node      `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
	Dense            *DenseNodes  `protobuf:"bytes,2,opt,name=dense" json:"dense,omitempty"`
	Ways             []*Way       `protobuf:"bytes,3,rep,name=ways" json:"ways,omitempty"`
	Relations        []*Relation  `protobuf:"bytes,4,rep,name=relations" json:"relations,omitempty"`
	Changesets       []*ChangeSet `protobuf:"bytes,5,rep,name=changesets" json:"changesets,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PrimitiveGroup) Reset()                    { *m = PrimitiveGroup{} }
func (m *PrimitiveGroup) String() string            { return proto.CompactTextString(m) }
func (*PrimitiveGroup) ProtoMessage()               {}
func (*PrimitiveGroup) Descriptor() ([]byte, []int) { return fileDescriptorOsmformat, []int{3} }

func (m *PrimitiveGroup) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *PrimitiveGroup) GetDense() *DenseNodes {
	if m != nil {
		return m.Dense
	}
	return nil
}

func (m *PrimitiveGroup) GetWays() []*Way {
	if m != nil {
		return m.Ways
	}
	return nil
}

func (m *PrimitiveGroup) GetRelations() []*Relation {
	if m != nil {
		return m.Relations
	}
	return nil
}

func (m *PrimitiveGroup) GetChangesets() []*ChangeSet {
	if m != nil {
		return m.Changesets
	}
	return nil
}

// * String table, contains the common strings in each block.
//
// Note that we reserve index '0' as a delimiter, so the entry at that
// index in the table is ALWAYS blank and unused.
//
type StringTable struct {
	S                []string `protobuf:"bytes,1,rep,name=s" json:"s,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *StringTable) Reset()                    { *m = StringTable{} }
func (m *StringTable) String() string            { return proto.CompactTextString(m) }
func (*StringTable) ProtoMessage()               {}
func (*StringTable) Descriptor() ([]byte, []int) { return fileDescriptorOsmformat, []int{4} }

func (m *StringTable) GetS() []string {
	if m != nil {
		return m.S
	}
	return nil
}

// Optional metadata that may be included into each primitive.
type Info struct {
	Version   *int32  `protobuf:"varint,1,opt,name=version,def=-1" json:"version,omitempty"`
	Timestamp *int64  `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Changeset *int64  `protobuf:"varint,3,opt,name=changeset" json:"changeset,omitempty"`
	Uid       *int32  `protobuf:"varint,4,opt,name=uid" json:"uid,omitempty"`
	UserSid   *uint32 `protobuf:"varint,5,opt,name=user_sid,json=userSid" json:"user_sid,omitempty"`
	// The visible flag is used to store history information. It indicates that
	// the current object version has been created by a delete operation on the
	// OSM API.
	// When a writer sets this flag, it MUST add a required_features tag with
	// value "HistoricalInformation" to the HeaderBlock.
	// If this flag is not available for some object it MUST be assumed to be
	// true if the file has the required_features tag "HistoricalInformation"
	// set.
	Visible          *bool  `protobuf:"varint,6,opt,name=visible" json:"visible,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Info) Reset()                    { *m = Info{} }
func (m *Info) String() string            { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()               {}
func (*Info) Descriptor() ([]byte, []int) { return fileDescriptorOsmformat, []int{5} }

const Default_Info_Version int32 = -1

func (m *Info) GetVersion() int32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return Default_Info_Version
}

func (m *Info) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *Info) GetChangeset() int64 {
	if m != nil && m.Changeset != nil {
		return *m.Changeset
	}
	return 0
}

func (m *Info) GetUid() int32 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *Info) GetUserSid() uint32 {
	if m != nil && m.UserSid != nil {
		return *m.UserSid
	}
	return 0
}

func (m *Info) GetVisible() bool {
	if m != nil && m.Visible != nil {
		return *m.Visible
	}
	return false
}

// * Optional metadata that may be included into each primitive. Special dense format used in DenseNodes.
type DenseInfo struct {
	Version   []int32 `protobuf:"varint,1,rep,packed,name=version" json:"version,omitempty"`
	Timestamp []int64 `protobuf:"zigzag64,2,rep,packed,name=timestamp" json:"timestamp,omitempty"`
	Changeset []int64 `protobuf:"zigzag64,3,rep,packed,name=changeset" json:"changeset,omitempty"`
	Uid       []int32 `protobuf:"zigzag32,4,rep,packed,name=uid" json:"uid,omitempty"`
	UserSid   []int32 `protobuf:"zigzag32,5,rep,packed,name=user_sid,json=userSid" json:"user_sid,omitempty"`
	// The visible flag is used to store history information. It indicates that
	// the current object version has been created by a delete operation on the
	// OSM API.
	// When a writer sets this flag, it MUST add a required_features tag with
	// value "HistoricalInformation" to the HeaderBlock.
	// If this flag is not available for some object it MUST be assumed to be
	// true if the file has the required_features tag "HistoricalInformation"
	// set.
	Visible          []bool `protobuf:"varint,6,rep,packed,name=visible" json:"visible,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DenseInfo) Reset()                    { *m = DenseInfo{} }
func (m *DenseInfo) String() string            { return proto.CompactTextString(m) }
func (*DenseInfo) ProtoMessage()               {}
func (*DenseInfo) Descriptor() ([]byte, []int) { return fileDescriptorOsmformat, []int{6} }

func (m *DenseInfo) GetVersion() []int32 {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *DenseInfo) GetTimestamp() []int64 {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *DenseInfo) GetChangeset() []int64 {
	if m != nil {
		return m.Changeset
	}
	return nil
}

func (m *DenseInfo) GetUid() []int32 {
	if m != nil {
		return m.Uid
	}
	return nil
}

func (m *DenseInfo) GetUserSid() []int32 {
	if m != nil {
		return m.UserSid
	}
	return nil
}

func (m *DenseInfo) GetVisible() []bool {
	if m != nil {
		return m.Visible
	}
	return nil
}

// THIS IS STUB DESIGN FOR CHANGESETS. NOT USED RIGHT NOW.
// TODO:    REMOVE THIS?
type ChangeSet struct {
	Id               *int64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ChangeSet) Reset()                    { *m = ChangeSet{} }
func (m *ChangeSet) String() string            { return proto.CompactTextString(m) }
func (*ChangeSet) ProtoMessage()               {}
func (*ChangeSet) Descriptor() ([]byte, []int) { return fileDescriptorOsmformat, []int{7} }

func (m *ChangeSet) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

type Node struct {
	Id *int64 `protobuf:"zigzag64,1,req,name=id" json:"id,omitempty"`
	// Parallel arrays.
	Keys             []uint32 `protobuf:"varint,2,rep,packed,name=keys" json:"keys,omitempty"`
	Vals             []uint32 `protobuf:"varint,3,rep,packed,name=vals" json:"vals,omitempty"`
	Info             *Info    `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`
	Lat              *int64   `protobuf:"zigzag64,8,req,name=lat" json:"lat,omitempty"`
	Lon              *int64   `protobuf:"zigzag64,9,req,name=lon" json:"lon,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptorOsmformat, []int{8} }

func (m *Node) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Node) GetKeys() []uint32 {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *Node) GetVals() []uint32 {
	if m != nil {
		return m.Vals
	}
	return nil
}

func (m *Node) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Node) GetLat() int64 {
	if m != nil && m.Lat != nil {
		return *m.Lat
	}
	return 0
}

func (m *Node) GetLon() int64 {
	if m != nil && m.Lon != nil {
		return *m.Lon
	}
	return 0
}

type DenseNodes struct {
	Id []int64 `protobuf:"zigzag64,1,rep,packed,name=id" json:"id,omitempty"`
	// repeated Info info = 4;
	Denseinfo *DenseInfo `protobuf:"bytes,5,opt,name=denseinfo" json:"denseinfo,omitempty"`
	Lat       []int64    `protobuf:"zigzag64,8,rep,packed,name=lat" json:"lat,omitempty"`
	Lon       []int64    `protobuf:"zigzag64,9,rep,packed,name=lon" json:"lon,omitempty"`
	// Special packing of keys and vals into one array. May be empty if all nodes in this block are tagless.
	KeysVals         []int32 `protobuf:"varint,10,rep,packed,name=keys_vals,json=keysVals" json:"keys_vals,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DenseNodes) Reset()                    { *m = DenseNodes{} }
func (m *DenseNodes) String() string            { return proto.CompactTextString(m) }
func (*DenseNodes) ProtoMessage()               {}
func (*DenseNodes) Descriptor() ([]byte, []int) { return fileDescriptorOsmformat, []int{9} }

func (m *DenseNodes) GetId() []int64 {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *DenseNodes) GetDenseinfo() *DenseInfo {
	if m != nil {
		return m.Denseinfo
	}
	return nil
}

func (m *DenseNodes) GetLat() []int64 {
	if m != nil {
		return m.Lat
	}
	return nil
}

func (m *DenseNodes) GetLon() []int64 {
	if m != nil {
		return m.Lon
	}
	return nil
}

func (m *DenseNodes) GetKeysVals() []int32 {
	if m != nil {
		return m.KeysVals
	}
	return nil
}

type Way struct {
	Id *int64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	// Parallel arrays.
	Keys             []uint32 `protobuf:"varint,2,rep,packed,name=keys" json:"keys,omitempty"`
	Vals             []uint32 `protobuf:"varint,3,rep,packed,name=vals" json:"vals,omitempty"`
	Info             *Info    `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`
	Refs             []int64  `protobuf:"zigzag64,8,rep,packed,name=refs" json:"refs,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Way) Reset()                    { *m = Way{} }
func (m *Way) String() string            { return proto.CompactTextString(m) }
func (*Way) ProtoMessage()               {}
func (*Way) Descriptor() ([]byte, []int) { return fileDescriptorOsmformat, []int{10} }

func (m *Way) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Way) GetKeys() []uint32 {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *Way) GetVals() []uint32 {
	if m != nil {
		return m.Vals
	}
	return nil
}

func (m *Way) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Way) GetRefs() []int64 {
	if m != nil {
		return m.Refs
	}
	return nil
}

type Relation struct {
	Id *int64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	// Parallel arrays.
	Keys []uint32 `protobuf:"varint,2,rep,packed,name=keys" json:"keys,omitempty"`
	Vals []uint32 `protobuf:"varint,3,rep,packed,name=vals" json:"vals,omitempty"`
	Info *Info    `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`
	// Parallel arrays
	RolesSid         []int32               `protobuf:"varint,8,rep,packed,name=roles_sid,json=rolesSid" json:"roles_sid,omitempty"`
	Memids           []int64               `protobuf:"zigzag64,9,rep,packed,name=memids" json:"memids,omitempty"`
	Types            []Relation_MemberType `protobuf:"varint,10,rep,packed,name=types,enum=OSMPBF.Relation_MemberType" json:"types,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *Relation) Reset()                    { *m = Relation{} }
func (m *Relation) String() string            { return proto.CompactTextString(m) }
func (*Relation) ProtoMessage()               {}
func (*Relation) Descriptor() ([]byte, []int) { return fileDescriptorOsmformat, []int{11} }

func (m *Relation) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Relation) GetKeys() []uint32 {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *Relation) GetVals() []uint32 {
	if m != nil {
		return m.Vals
	}
	return nil
}

func (m *Relation) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Relation) GetRolesSid() []int32 {
	if m != nil {
		return m.RolesSid
	}
	return nil
}

func (m *Relation) GetMemids() []int64 {
	if m != nil {
		return m.Memids
	}
	return nil
}

func (m *Relation) GetTypes() []Relation_MemberType {
	if m != nil {
		return m.Types
	}
	return nil
}

func init() {
	proto.RegisterType((*HeaderBlock)(nil), "OSMPBF.HeaderBlock")
	proto.RegisterType((*HeaderBBox)(nil), "OSMPBF.HeaderBBox")
	proto.RegisterType((*PrimitiveBlock)(nil), "OSMPBF.PrimitiveBlock")
	proto.RegisterType((*PrimitiveGroup)(nil), "OSMPBF.PrimitiveGroup")
	proto.RegisterType((*StringTable)(nil), "OSMPBF.StringTable")
	proto.RegisterType((*Info)(nil), "OSMPBF.Info")
	proto.RegisterType((*DenseInfo)(nil), "OSMPBF.DenseInfo")
	proto.RegisterType((*ChangeSet)(nil), "OSMPBF.ChangeSet")
	proto.RegisterType((*Node)(nil), "OSMPBF.Node")
	proto.RegisterType((*DenseNodes)(nil), "OSMPBF.DenseNodes")
	proto.RegisterType((*Way)(nil), "OSMPBF.Way")
	proto.RegisterType((*Relation)(nil), "OSMPBF.Relation")
	proto.RegisterEnum("OSMPBF.Relation_MemberType", Relation_MemberType_name, Relation_MemberType_value)
}

func init() { proto.RegisterFile("osmformat.proto", fileDescriptorOsmformat) }

var fileDescriptorOsmformat = []byte{
	// 1019 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xcb, 0x6e, 0xdb, 0x46,
	0x14, 0x2d, 0x1f, 0xb2, 0xc5, 0xab, 0xc4, 0x91, 0x26, 0x86, 0xc1, 0xd4, 0x0e, 0xc2, 0xb2, 0x68,
	0x40, 0xa0, 0x88, 0x22, 0x0b, 0xc8, 0xc6, 0x8b, 0x16, 0x51, 0xf3, 0x04, 0x6a, 0x3b, 0x18, 0xbb,
	0x35, 0xba, 0x52, 0x47, 0xe2, 0x48, 0x19, 0x84, 0xe2, 0x28, 0x33, 0x23, 0x27, 0x5a, 0xf5, 0x07,
	0x0a, 0xf4, 0x0f, 0xba, 0xc9, 0xa2, 0xbf, 0xd1, 0x6f, 0xe9, 0x97, 0x14, 0x73, 0x49, 0x8a, 0xb4,
	0x9d, 0x6d, 0xbb, 0xe3, 0x9c, 0x73, 0x78, 0xef, 0x99, 0xfb, 0x20, 0xe1, 0x8e, 0xd4, 0x8b, 0x99,
	0x54, 0x0b, 0x66, 0xfa, 0x4b, 0x25, 0x8d, 0x24, 0x5b, 0xa7, 0x67, 0xc7, 0x6f, 0x46, 0x2f, 0xe2,
	0x4f, 0x1e, 0x74, 0x5e, 0x71, 0x96, 0x72, 0x35, 0xca, 0xe4, 0xf4, 0x1d, 0x79, 0x08, 0xfe, 0x64,
	0x22, 0x3f, 0x86, 0x4e, 0xe4, 0x24, 0x9d, 0x21, 0xe9, 0x17, 0xb2, 0x7e, 0x29, 0x19, 0xc9, 0x8f,
	0x14, 0x79, 0xf2, 0x2d, 0xf4, 0x14, 0x7f, 0xbf, 0x12, 0x8a, 0xa7, 0xe3, 0x19, 0x67, 0x66, 0xa5,
	0xb8, 0x0e, 0xfd, 0xc8, 0x4b, 0x02, 0xda, 0xad, 0x88, 0x17, 0x25, 0x6e, 0xc5, 0x72, 0x69, 0x84,
	0xcc, 0x59, 0x56, 0x8b, 0x5b, 0x85, 0xb8, 0x22, 0x36, 0xe2, 0x87, 0xb0, 0xf3, 0x41, 0x09, 0x23,
	0xf2, 0xf9, 0x52, 0xc9, 0xb9, 0x62, 0x8b, 0xb0, 0x1b, 0x39, 0x49, 0x40, 0xaf, 0xa1, 0x64, 0x0f,
	0xb6, 0xb4, 0x5c, 0xa9, 0x29, 0x0f, 0x7b, 0xc8, 0x97, 0x27, 0x32, 0x82, 0xfb, 0x52, 0x2f, 0xa4,
	0x16, 0x7a, 0xac, 0xf8, 0x32, 0x13, 0x53, 0x66, 0x13, 0x8c, 0x8d, 0x58, 0x70, 0x6d, 0xd8, 0x62,
	0x19, 0x46, 0x91, 0x93, 0x78, 0x74, 0xbf, 0x14, 0xd1, 0x5a, 0x73, 0x5e, 0x49, 0xc8, 0x31, 0x7c,
	0xfd, 0xb9, 0x18, 0x9a, 0xbf, 0x5f, 0xf1, 0x7c, 0xca, 0xc7, 0xf9, 0x6a, 0x31, 0xe1, 0x2a, 0xfc,
	0x0a, 0x23, 0x45, 0x37, 0x23, 0x9d, 0x95, 0xc2, 0x13, 0xd4, 0x91, 0xef, 0xe1, 0xe0, 0x73, 0xe1,
	0x26, 0x4c, 0xf3, 0xf1, 0x4a, 0x65, 0x61, 0x8c, 0x17, 0xb8, 0x77, 0x33, 0xce, 0x88, 0x69, 0xfe,
	0x93, 0xca, 0xe2, 0x5f, 0x01, 0xea, 0x0e, 0x10, 0x02, 0x7e, 0xc6, 0x67, 0x26, 0x74, 0x22, 0x37,
	0x21, 0x14, 0x9f, 0xc9, 0x2e, 0xb4, 0x94, 0x98, 0xbf, 0x35, 0xa1, 0x8b, 0x60, 0x71, 0x20, 0x5d,
	0xf0, 0x8c, 0x5c, 0x86, 0x1e, 0x62, 0xf6, 0xd1, 0x56, 0x6d, 0x22, 0x8d, 0x91, 0x8b, 0xd0, 0x47,
	0xb0, 0x3c, 0xc5, 0x9f, 0x5c, 0xd8, 0x79, 0xa3, 0xc4, 0x42, 0x18, 0x71, 0xc9, 0x8b, 0x51, 0x78,
	0x02, 0x1d, 0x6d, 0x94, 0xc8, 0xe7, 0x86, 0x4d, 0x32, 0x8e, 0xd9, 0x3a, 0xc3, 0xbb, 0xd5, 0x44,
	0x9c, 0x21, 0x75, 0x6e, 0x29, 0xda, 0xd4, 0x91, 0xef, 0x60, 0x67, 0x59, 0x05, 0x9a, 0x2b, 0xb9,
	0x5a, 0x86, 0x6e, 0xe4, 0x25, 0x9d, 0xe1, 0x5e, 0xf5, 0xe6, 0x26, 0xcd, 0x4b, 0xcb, 0xd2, 0x6b,
	0x6a, 0xf2, 0x0d, 0x74, 0xe6, 0x8a, 0xe5, 0xab, 0x8c, 0x29, 0x61, 0xd6, 0xd8, 0xdc, 0xd6, 0x91,
	0x77, 0x38, 0x18, 0xd0, 0x26, 0x4e, 0x22, 0x80, 0x8c, 0x99, 0xb1, 0x9c, 0xcd, 0x34, 0x37, 0xe1,
	0x5d, 0xdb, 0x89, 0x23, 0x67, 0x40, 0x83, 0x8c, 0x99, 0x53, 0xc4, 0x50, 0x21, 0xf3, 0x4a, 0xb1,
	0x5b, 0x2b, 0x64, 0x5e, 0x2a, 0x1e, 0x43, 0x37, 0x65, 0x86, 0x8f, 0x9b, 0xf9, 0x08, 0xe6, 0xf3,
	0x0f, 0x07, 0x83, 0x01, 0xbd, 0x63, 0xd9, 0x97, 0x35, 0x19, 0xff, 0xe3, 0x34, 0xaa, 0x84, 0xf6,
	0x49, 0x0c, 0xad, 0x5c, 0xa6, 0x5c, 0x87, 0x0e, 0xde, 0xf2, 0x56, 0x75, 0xcb, 0x13, 0x99, 0x72,
	0x5a, 0x50, 0x24, 0x81, 0x56, 0xca, 0x73, 0xcd, 0x43, 0xf7, 0xea, 0x56, 0x3d, 0xb3, 0xa0, 0x15,
	0x6a, 0x5a, 0x08, 0xc8, 0x03, 0xf0, 0x3f, 0xb0, 0xb5, 0x0e, 0x3d, 0x0c, 0xd6, 0xa9, 0x84, 0x17,
	0x6c, 0x4d, 0x91, 0x20, 0x7d, 0x08, 0x14, 0xcf, 0x70, 0x38, 0x8a, 0x7d, 0xeb, 0x0c, 0xbb, 0x95,
	0x8a, 0x96, 0x04, 0xad, 0x25, 0xe4, 0x10, 0x60, 0xfa, 0x96, 0xe5, 0x73, 0xae, 0xb9, 0x29, 0x76,
	0xae, 0x33, 0xec, 0x55, 0x2f, 0xfc, 0x80, 0xcc, 0x19, 0x37, 0xb4, 0x21, 0x8a, 0xf7, 0xa1, 0xd3,
	0x68, 0x2e, 0xb9, 0x05, 0x4e, 0x71, 0xb9, 0x80, 0x3a, 0x3a, 0xfe, 0xcb, 0x01, 0xff, 0x75, 0x3e,
	0x93, 0xe4, 0x00, 0xb6, 0x2f, 0xb9, 0xd2, 0x42, 0xe6, 0xf8, 0xad, 0x68, 0x1d, 0xb9, 0x8f, 0x0e,
	0x69, 0x05, 0x91, 0x03, 0x08, 0xea, 0x85, 0x73, 0x71, 0x4d, 0x6a, 0xc0, 0xb2, 0x9b, 0x7c, 0xa1,
	0x57, 0xb0, 0x1b, 0xc0, 0x0e, 0xed, 0x4a, 0xa4, 0xa1, 0x6f, 0xa3, 0x52, 0xfb, 0x48, 0xee, 0x41,
	0x7b, 0xa5, 0xb9, 0x1a, 0x6b, 0x91, 0x86, 0xad, 0xc8, 0x49, 0x6e, 0xd3, 0x6d, 0x7b, 0x3e, 0x13,
	0x29, 0x09, 0x61, 0xfb, 0x52, 0x68, 0x61, 0x07, 0x74, 0x2b, 0x72, 0x92, 0x36, 0xad, 0x8e, 0xf1,
	0xdf, 0x0e, 0x04, 0x58, 0xe0, 0x9b, 0x76, 0xbd, 0xa4, 0x35, 0x72, 0xbb, 0x4e, 0x6d, 0x37, 0xba,
	0x6a, 0xd7, 0x4b, 0x08, 0xf2, 0x0d, 0xcb, 0xd1, 0x55, 0xcb, 0x1b, 0x45, 0x6d, 0x7b, 0xb7, 0xb2,
	0xed, 0x25, 0x3d, 0xe4, 0xd0, 0xfa, 0xfd, 0x2b, 0xd6, 0x2b, 0x6a, 0x63, 0xff, 0xa0, 0x69, 0xdf,
	0x4b, 0xda, 0xa5, 0xad, 0xf2, 0x0a, 0xfb, 0x10, 0x6c, 0x5a, 0x44, 0x76, 0xc0, 0x15, 0x29, 0x6e,
	0xa1, 0x47, 0x5d, 0x91, 0xc6, 0xbf, 0x3b, 0xe0, 0xdb, 0xd9, 0x69, 0x10, 0xc4, 0x12, 0x64, 0x0f,
	0xfc, 0x77, 0x7c, 0xad, 0xf1, 0x1e, 0xb7, 0x31, 0x20, 0x9e, 0x2d, 0x7e, 0xc9, 0xb2, 0x62, 0xb6,
	0x4a, 0xdc, 0x9e, 0x49, 0x04, 0xbe, 0xc8, 0x67, 0x12, 0x0b, 0xde, 0x18, 0x60, 0x5b, 0x36, 0x8a,
	0x8c, 0xed, 0x48, 0xc6, 0x4c, 0xd8, 0x2e, 0x3e, 0x23, 0x19, 0xc3, 0x1e, 0x65, 0x32, 0x0f, 0x83,
	0x12, 0x91, 0x79, 0xfc, 0xa7, 0x03, 0x50, 0xcf, 0x33, 0x21, 0xa5, 0xa9, 0xaa, 0x50, 0xd6, 0xd8,
	0x63, 0x08, 0x70, 0xca, 0x31, 0x5b, 0x0b, 0xb3, 0xf5, 0xae, 0xac, 0x02, 0xa6, 0xac, 0x35, 0xb6,
	0xa4, 0x45, 0xde, 0x2a, 0x0a, 0xe6, 0xde, 0xad, 0x72, 0xd7, 0xa8, 0xcc, 0xc9, 0x03, 0x08, 0xec,
	0x2d, 0xc7, 0x78, 0x45, 0xd8, 0xb4, 0xb8, 0x6d, 0xc1, 0x9f, 0x59, 0xa6, 0xe3, 0xdf, 0xc0, 0xbb,
	0x60, 0xeb, 0xeb, 0x65, 0xfc, 0x0f, 0xaa, 0xb5, 0x07, 0xbe, 0xe2, 0x33, 0xdd, 0xb0, 0x8d, 0xe7,
	0xf8, 0x0f, 0x17, 0xda, 0xd5, 0x8a, 0xfe, 0x0f, 0x36, 0x1e, 0x40, 0xa0, 0x64, 0xc6, 0x35, 0x8e,
	0x5e, 0xbb, 0x2e, 0x08, 0x82, 0x76, 0xf6, 0xbe, 0x84, 0xad, 0x05, 0x5f, 0x88, 0x54, 0x37, 0x4a,
	0x59, 0x22, 0xe4, 0x09, 0xb4, 0xcc, 0x7a, 0xc9, 0x8b, 0x4a, 0xee, 0x0c, 0xf7, 0xaf, 0x7f, 0x62,
	0xfa, 0xc7, 0xdc, 0xfe, 0xd9, 0xce, 0xd7, 0x4b, 0x8e, 0xef, 0x15, 0xea, 0xf8, 0x11, 0x40, 0x4d,
	0x90, 0x36, 0xf8, 0x27, 0xa7, 0xcf, 0x9e, 0x77, 0xbf, 0x20, 0xdb, 0xe0, 0x5d, 0x3c, 0xfd, 0xa5,
	0xeb, 0x90, 0x5b, 0xd0, 0xa6, 0xcf, 0x7f, 0x7c, 0x7a, 0xfe, 0xfa, 0xf4, 0xa4, 0xeb, 0x8e, 0x7a,
	0x70, 0x7b, 0xaa, 0xa4, 0x9e, 0xac, 0xfb, 0x13, 0x91, 0x33, 0xb5, 0x7e, 0xe5, 0xfd, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0x75, 0xc1, 0x01, 0x6f, 0xa9, 0x08, 0x00, 0x00,
}
