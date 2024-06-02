// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

package msgpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	commonpb "github.com/milvus-io/milvus-proto/go-api/v2/commonpb"
	schemapb "github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
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

type InsertDataVersion int32

const (
	// 0 must refer to row-based format, since it's the first version in Milvus.
	InsertDataVersion_RowBased    InsertDataVersion = 0
	InsertDataVersion_ColumnBased InsertDataVersion = 1
)

var InsertDataVersion_name = map[int32]string{
	0: "RowBased",
	1: "ColumnBased",
}

var InsertDataVersion_value = map[string]int32{
	"RowBased":    0,
	"ColumnBased": 1,
}

func (x InsertDataVersion) String() string {
	return proto.EnumName(InsertDataVersion_name, int32(x))
}

func (InsertDataVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

type InsertRequest struct {
	Base           *commonpb.MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	ShardName      string            `protobuf:"bytes,2,opt,name=shardName,proto3" json:"shardName,omitempty"`
	DbName         string            `protobuf:"bytes,3,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
	CollectionName string            `protobuf:"bytes,4,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
	PartitionName  string            `protobuf:"bytes,5,opt,name=partition_name,json=partitionName,proto3" json:"partition_name,omitempty"`
	DbID           int64             `protobuf:"varint,6,opt,name=dbID,proto3" json:"dbID,omitempty"`
	CollectionID   int64             `protobuf:"varint,7,opt,name=collectionID,proto3" json:"collectionID,omitempty"`
	PartitionID    int64             `protobuf:"varint,8,opt,name=partitionID,proto3" json:"partitionID,omitempty"`
	SegmentID      int64             `protobuf:"varint,9,opt,name=segmentID,proto3" json:"segmentID,omitempty"`
	Timestamps     []uint64          `protobuf:"varint,10,rep,packed,name=timestamps,proto3" json:"timestamps,omitempty"`
	RowIDs         []int64           `protobuf:"varint,11,rep,packed,name=rowIDs,proto3" json:"rowIDs,omitempty"`
	// row_data was reserved for compatibility
	RowData              []*commonpb.Blob      `protobuf:"bytes,12,rep,name=row_data,json=rowData,proto3" json:"row_data,omitempty"`
	FieldsData           []*schemapb.FieldData `protobuf:"bytes,13,rep,name=fields_data,json=fieldsData,proto3" json:"fields_data,omitempty"`
	NumRows              uint64                `protobuf:"varint,14,opt,name=num_rows,json=numRows,proto3" json:"num_rows,omitempty"`
	Version              InsertDataVersion     `protobuf:"varint,15,opt,name=version,proto3,enum=milvus.proto.msg.InsertDataVersion" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *InsertRequest) Reset()         { *m = InsertRequest{} }
func (m *InsertRequest) String() string { return proto.CompactTextString(m) }
func (*InsertRequest) ProtoMessage()    {}
func (*InsertRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

func (m *InsertRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsertRequest.Unmarshal(m, b)
}
func (m *InsertRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsertRequest.Marshal(b, m, deterministic)
}
func (m *InsertRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsertRequest.Merge(m, src)
}
func (m *InsertRequest) XXX_Size() int {
	return xxx_messageInfo_InsertRequest.Size(m)
}
func (m *InsertRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InsertRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InsertRequest proto.InternalMessageInfo

func (m *InsertRequest) GetBase() *commonpb.MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *InsertRequest) GetShardName() string {
	if m != nil {
		return m.ShardName
	}
	return ""
}

func (m *InsertRequest) GetDbName() string {
	if m != nil {
		return m.DbName
	}
	return ""
}

func (m *InsertRequest) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *InsertRequest) GetPartitionName() string {
	if m != nil {
		return m.PartitionName
	}
	return ""
}

func (m *InsertRequest) GetDbID() int64 {
	if m != nil {
		return m.DbID
	}
	return 0
}

func (m *InsertRequest) GetCollectionID() int64 {
	if m != nil {
		return m.CollectionID
	}
	return 0
}

func (m *InsertRequest) GetPartitionID() int64 {
	if m != nil {
		return m.PartitionID
	}
	return 0
}

func (m *InsertRequest) GetSegmentID() int64 {
	if m != nil {
		return m.SegmentID
	}
	return 0
}

func (m *InsertRequest) GetTimestamps() []uint64 {
	if m != nil {
		return m.Timestamps
	}
	return nil
}

func (m *InsertRequest) GetRowIDs() []int64 {
	if m != nil {
		return m.RowIDs
	}
	return nil
}

func (m *InsertRequest) GetRowData() []*commonpb.Blob {
	if m != nil {
		return m.RowData
	}
	return nil
}

func (m *InsertRequest) GetFieldsData() []*schemapb.FieldData {
	if m != nil {
		return m.FieldsData
	}
	return nil
}

func (m *InsertRequest) GetNumRows() uint64 {
	if m != nil {
		return m.NumRows
	}
	return 0
}

func (m *InsertRequest) GetVersion() InsertDataVersion {
	if m != nil {
		return m.Version
	}
	return InsertDataVersion_RowBased
}

type DeleteRequest struct {
	Base                 *commonpb.MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	ShardName            string            `protobuf:"bytes,2,opt,name=shardName,proto3" json:"shardName,omitempty"`
	DbName               string            `protobuf:"bytes,3,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
	CollectionName       string            `protobuf:"bytes,4,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
	PartitionName        string            `protobuf:"bytes,5,opt,name=partition_name,json=partitionName,proto3" json:"partition_name,omitempty"`
	DbID                 int64             `protobuf:"varint,6,opt,name=dbID,proto3" json:"dbID,omitempty"`
	CollectionID         int64             `protobuf:"varint,7,opt,name=collectionID,proto3" json:"collectionID,omitempty"`
	PartitionID          int64             `protobuf:"varint,8,opt,name=partitionID,proto3" json:"partitionID,omitempty"`
	Int64PrimaryKeys     []int64           `protobuf:"varint,9,rep,packed,name=int64_primary_keys,json=int64PrimaryKeys,proto3" json:"int64_primary_keys,omitempty"`
	Timestamps           []uint64          `protobuf:"varint,10,rep,packed,name=timestamps,proto3" json:"timestamps,omitempty"`
	NumRows              int64             `protobuf:"varint,11,opt,name=num_rows,json=numRows,proto3" json:"num_rows,omitempty"`
	PrimaryKeys          *schemapb.IDs     `protobuf:"bytes,12,opt,name=primary_keys,json=primaryKeys,proto3" json:"primary_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetBase() *commonpb.MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *DeleteRequest) GetShardName() string {
	if m != nil {
		return m.ShardName
	}
	return ""
}

func (m *DeleteRequest) GetDbName() string {
	if m != nil {
		return m.DbName
	}
	return ""
}

func (m *DeleteRequest) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *DeleteRequest) GetPartitionName() string {
	if m != nil {
		return m.PartitionName
	}
	return ""
}

func (m *DeleteRequest) GetDbID() int64 {
	if m != nil {
		return m.DbID
	}
	return 0
}

func (m *DeleteRequest) GetCollectionID() int64 {
	if m != nil {
		return m.CollectionID
	}
	return 0
}

func (m *DeleteRequest) GetPartitionID() int64 {
	if m != nil {
		return m.PartitionID
	}
	return 0
}

func (m *DeleteRequest) GetInt64PrimaryKeys() []int64 {
	if m != nil {
		return m.Int64PrimaryKeys
	}
	return nil
}

func (m *DeleteRequest) GetTimestamps() []uint64 {
	if m != nil {
		return m.Timestamps
	}
	return nil
}

func (m *DeleteRequest) GetNumRows() int64 {
	if m != nil {
		return m.NumRows
	}
	return 0
}

func (m *DeleteRequest) GetPrimaryKeys() *schemapb.IDs {
	if m != nil {
		return m.PrimaryKeys
	}
	return nil
}

type MsgPosition struct {
	ChannelName          string   `protobuf:"bytes,1,opt,name=channel_name,json=channelName,proto3" json:"channel_name,omitempty"`
	MsgID                []byte   `protobuf:"bytes,2,opt,name=msgID,proto3" json:"msgID,omitempty"`
	MsgGroup             string   `protobuf:"bytes,3,opt,name=msgGroup,proto3" json:"msgGroup,omitempty"`
	Timestamp            uint64   `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgPosition) Reset()         { *m = MsgPosition{} }
func (m *MsgPosition) String() string { return proto.CompactTextString(m) }
func (*MsgPosition) ProtoMessage()    {}
func (*MsgPosition) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{2}
}

func (m *MsgPosition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgPosition.Unmarshal(m, b)
}
func (m *MsgPosition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgPosition.Marshal(b, m, deterministic)
}
func (m *MsgPosition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPosition.Merge(m, src)
}
func (m *MsgPosition) XXX_Size() int {
	return xxx_messageInfo_MsgPosition.Size(m)
}
func (m *MsgPosition) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPosition.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPosition proto.InternalMessageInfo

func (m *MsgPosition) GetChannelName() string {
	if m != nil {
		return m.ChannelName
	}
	return ""
}

func (m *MsgPosition) GetMsgID() []byte {
	if m != nil {
		return m.MsgID
	}
	return nil
}

func (m *MsgPosition) GetMsgGroup() string {
	if m != nil {
		return m.MsgGroup
	}
	return ""
}

func (m *MsgPosition) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type CreateCollectionRequest struct {
	Base           *commonpb.MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	DbName         string            `protobuf:"bytes,2,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
	CollectionName string            `protobuf:"bytes,3,opt,name=collectionName,proto3" json:"collectionName,omitempty"`
	PartitionName  string            `protobuf:"bytes,4,opt,name=partitionName,proto3" json:"partitionName,omitempty"`
	// `schema` is the serialized `schema.CollectionSchema`
	DbID                 int64    `protobuf:"varint,5,opt,name=dbID,proto3" json:"dbID,omitempty"`
	CollectionID         int64    `protobuf:"varint,6,opt,name=collectionID,proto3" json:"collectionID,omitempty"`
	PartitionID          int64    `protobuf:"varint,7,opt,name=partitionID,proto3" json:"partitionID,omitempty"`
	Schema               []byte   `protobuf:"bytes,8,opt,name=schema,proto3" json:"schema,omitempty"`
	VirtualChannelNames  []string `protobuf:"bytes,9,rep,name=virtualChannelNames,proto3" json:"virtualChannelNames,omitempty"`
	PhysicalChannelNames []string `protobuf:"bytes,10,rep,name=physicalChannelNames,proto3" json:"physicalChannelNames,omitempty"`
	PartitionIDs         []int64  `protobuf:"varint,11,rep,packed,name=partitionIDs,proto3" json:"partitionIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCollectionRequest) Reset()         { *m = CreateCollectionRequest{} }
func (m *CreateCollectionRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCollectionRequest) ProtoMessage()    {}
func (*CreateCollectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{3}
}

func (m *CreateCollectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCollectionRequest.Unmarshal(m, b)
}
func (m *CreateCollectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCollectionRequest.Marshal(b, m, deterministic)
}
func (m *CreateCollectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCollectionRequest.Merge(m, src)
}
func (m *CreateCollectionRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCollectionRequest.Size(m)
}
func (m *CreateCollectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCollectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCollectionRequest proto.InternalMessageInfo

func (m *CreateCollectionRequest) GetBase() *commonpb.MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *CreateCollectionRequest) GetDbName() string {
	if m != nil {
		return m.DbName
	}
	return ""
}

func (m *CreateCollectionRequest) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *CreateCollectionRequest) GetPartitionName() string {
	if m != nil {
		return m.PartitionName
	}
	return ""
}

func (m *CreateCollectionRequest) GetDbID() int64 {
	if m != nil {
		return m.DbID
	}
	return 0
}

func (m *CreateCollectionRequest) GetCollectionID() int64 {
	if m != nil {
		return m.CollectionID
	}
	return 0
}

func (m *CreateCollectionRequest) GetPartitionID() int64 {
	if m != nil {
		return m.PartitionID
	}
	return 0
}

func (m *CreateCollectionRequest) GetSchema() []byte {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *CreateCollectionRequest) GetVirtualChannelNames() []string {
	if m != nil {
		return m.VirtualChannelNames
	}
	return nil
}

func (m *CreateCollectionRequest) GetPhysicalChannelNames() []string {
	if m != nil {
		return m.PhysicalChannelNames
	}
	return nil
}

func (m *CreateCollectionRequest) GetPartitionIDs() []int64 {
	if m != nil {
		return m.PartitionIDs
	}
	return nil
}

type DropCollectionRequest struct {
	Base                 *commonpb.MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	DbName               string            `protobuf:"bytes,2,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
	CollectionName       string            `protobuf:"bytes,3,opt,name=collectionName,proto3" json:"collectionName,omitempty"`
	DbID                 int64             `protobuf:"varint,4,opt,name=dbID,proto3" json:"dbID,omitempty"`
	CollectionID         int64             `protobuf:"varint,5,opt,name=collectionID,proto3" json:"collectionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DropCollectionRequest) Reset()         { *m = DropCollectionRequest{} }
func (m *DropCollectionRequest) String() string { return proto.CompactTextString(m) }
func (*DropCollectionRequest) ProtoMessage()    {}
func (*DropCollectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{4}
}

func (m *DropCollectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DropCollectionRequest.Unmarshal(m, b)
}
func (m *DropCollectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DropCollectionRequest.Marshal(b, m, deterministic)
}
func (m *DropCollectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DropCollectionRequest.Merge(m, src)
}
func (m *DropCollectionRequest) XXX_Size() int {
	return xxx_messageInfo_DropCollectionRequest.Size(m)
}
func (m *DropCollectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DropCollectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DropCollectionRequest proto.InternalMessageInfo

func (m *DropCollectionRequest) GetBase() *commonpb.MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *DropCollectionRequest) GetDbName() string {
	if m != nil {
		return m.DbName
	}
	return ""
}

func (m *DropCollectionRequest) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *DropCollectionRequest) GetDbID() int64 {
	if m != nil {
		return m.DbID
	}
	return 0
}

func (m *DropCollectionRequest) GetCollectionID() int64 {
	if m != nil {
		return m.CollectionID
	}
	return 0
}

type CreatePartitionRequest struct {
	Base                 *commonpb.MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	DbName               string            `protobuf:"bytes,2,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
	CollectionName       string            `protobuf:"bytes,3,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
	PartitionName        string            `protobuf:"bytes,4,opt,name=partition_name,json=partitionName,proto3" json:"partition_name,omitempty"`
	DbID                 int64             `protobuf:"varint,5,opt,name=dbID,proto3" json:"dbID,omitempty"`
	CollectionID         int64             `protobuf:"varint,6,opt,name=collectionID,proto3" json:"collectionID,omitempty"`
	PartitionID          int64             `protobuf:"varint,7,opt,name=partitionID,proto3" json:"partitionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CreatePartitionRequest) Reset()         { *m = CreatePartitionRequest{} }
func (m *CreatePartitionRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePartitionRequest) ProtoMessage()    {}
func (*CreatePartitionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{5}
}

func (m *CreatePartitionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePartitionRequest.Unmarshal(m, b)
}
func (m *CreatePartitionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePartitionRequest.Marshal(b, m, deterministic)
}
func (m *CreatePartitionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePartitionRequest.Merge(m, src)
}
func (m *CreatePartitionRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePartitionRequest.Size(m)
}
func (m *CreatePartitionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePartitionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePartitionRequest proto.InternalMessageInfo

func (m *CreatePartitionRequest) GetBase() *commonpb.MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *CreatePartitionRequest) GetDbName() string {
	if m != nil {
		return m.DbName
	}
	return ""
}

func (m *CreatePartitionRequest) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *CreatePartitionRequest) GetPartitionName() string {
	if m != nil {
		return m.PartitionName
	}
	return ""
}

func (m *CreatePartitionRequest) GetDbID() int64 {
	if m != nil {
		return m.DbID
	}
	return 0
}

func (m *CreatePartitionRequest) GetCollectionID() int64 {
	if m != nil {
		return m.CollectionID
	}
	return 0
}

func (m *CreatePartitionRequest) GetPartitionID() int64 {
	if m != nil {
		return m.PartitionID
	}
	return 0
}

type DropPartitionRequest struct {
	Base                 *commonpb.MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	DbName               string            `protobuf:"bytes,2,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
	CollectionName       string            `protobuf:"bytes,3,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
	PartitionName        string            `protobuf:"bytes,4,opt,name=partition_name,json=partitionName,proto3" json:"partition_name,omitempty"`
	DbID                 int64             `protobuf:"varint,5,opt,name=dbID,proto3" json:"dbID,omitempty"`
	CollectionID         int64             `protobuf:"varint,6,opt,name=collectionID,proto3" json:"collectionID,omitempty"`
	PartitionID          int64             `protobuf:"varint,7,opt,name=partitionID,proto3" json:"partitionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DropPartitionRequest) Reset()         { *m = DropPartitionRequest{} }
func (m *DropPartitionRequest) String() string { return proto.CompactTextString(m) }
func (*DropPartitionRequest) ProtoMessage()    {}
func (*DropPartitionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{6}
}

func (m *DropPartitionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DropPartitionRequest.Unmarshal(m, b)
}
func (m *DropPartitionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DropPartitionRequest.Marshal(b, m, deterministic)
}
func (m *DropPartitionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DropPartitionRequest.Merge(m, src)
}
func (m *DropPartitionRequest) XXX_Size() int {
	return xxx_messageInfo_DropPartitionRequest.Size(m)
}
func (m *DropPartitionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DropPartitionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DropPartitionRequest proto.InternalMessageInfo

func (m *DropPartitionRequest) GetBase() *commonpb.MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *DropPartitionRequest) GetDbName() string {
	if m != nil {
		return m.DbName
	}
	return ""
}

func (m *DropPartitionRequest) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *DropPartitionRequest) GetPartitionName() string {
	if m != nil {
		return m.PartitionName
	}
	return ""
}

func (m *DropPartitionRequest) GetDbID() int64 {
	if m != nil {
		return m.DbID
	}
	return 0
}

func (m *DropPartitionRequest) GetCollectionID() int64 {
	if m != nil {
		return m.CollectionID
	}
	return 0
}

func (m *DropPartitionRequest) GetPartitionID() int64 {
	if m != nil {
		return m.PartitionID
	}
	return 0
}

type TimeTickMsg struct {
	Base                 *commonpb.MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TimeTickMsg) Reset()         { *m = TimeTickMsg{} }
func (m *TimeTickMsg) String() string { return proto.CompactTextString(m) }
func (*TimeTickMsg) ProtoMessage()    {}
func (*TimeTickMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{7}
}

func (m *TimeTickMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeTickMsg.Unmarshal(m, b)
}
func (m *TimeTickMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeTickMsg.Marshal(b, m, deterministic)
}
func (m *TimeTickMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeTickMsg.Merge(m, src)
}
func (m *TimeTickMsg) XXX_Size() int {
	return xxx_messageInfo_TimeTickMsg.Size(m)
}
func (m *TimeTickMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeTickMsg.DiscardUnknown(m)
}

var xxx_messageInfo_TimeTickMsg proto.InternalMessageInfo

func (m *TimeTickMsg) GetBase() *commonpb.MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

type DataNodeTtMsg struct {
	Base                 *commonpb.MsgBase        `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	ChannelName          string                   `protobuf:"bytes,2,opt,name=channel_name,json=channelName,proto3" json:"channel_name,omitempty"`
	Timestamp            uint64                   `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	SegmentsStats        []*commonpb.SegmentStats `protobuf:"bytes,4,rep,name=segments_stats,json=segmentsStats,proto3" json:"segments_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *DataNodeTtMsg) Reset()         { *m = DataNodeTtMsg{} }
func (m *DataNodeTtMsg) String() string { return proto.CompactTextString(m) }
func (*DataNodeTtMsg) ProtoMessage()    {}
func (*DataNodeTtMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{8}
}

func (m *DataNodeTtMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataNodeTtMsg.Unmarshal(m, b)
}
func (m *DataNodeTtMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataNodeTtMsg.Marshal(b, m, deterministic)
}
func (m *DataNodeTtMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataNodeTtMsg.Merge(m, src)
}
func (m *DataNodeTtMsg) XXX_Size() int {
	return xxx_messageInfo_DataNodeTtMsg.Size(m)
}
func (m *DataNodeTtMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_DataNodeTtMsg.DiscardUnknown(m)
}

var xxx_messageInfo_DataNodeTtMsg proto.InternalMessageInfo

func (m *DataNodeTtMsg) GetBase() *commonpb.MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *DataNodeTtMsg) GetChannelName() string {
	if m != nil {
		return m.ChannelName
	}
	return ""
}

func (m *DataNodeTtMsg) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *DataNodeTtMsg) GetSegmentsStats() []*commonpb.SegmentStats {
	if m != nil {
		return m.SegmentsStats
	}
	return nil
}

func init() {
	proto.RegisterEnum("milvus.proto.msg.InsertDataVersion", InsertDataVersion_name, InsertDataVersion_value)
	proto.RegisterType((*InsertRequest)(nil), "milvus.proto.msg.InsertRequest")
	proto.RegisterType((*DeleteRequest)(nil), "milvus.proto.msg.DeleteRequest")
	proto.RegisterType((*MsgPosition)(nil), "milvus.proto.msg.MsgPosition")
	proto.RegisterType((*CreateCollectionRequest)(nil), "milvus.proto.msg.CreateCollectionRequest")
	proto.RegisterType((*DropCollectionRequest)(nil), "milvus.proto.msg.DropCollectionRequest")
	proto.RegisterType((*CreatePartitionRequest)(nil), "milvus.proto.msg.CreatePartitionRequest")
	proto.RegisterType((*DropPartitionRequest)(nil), "milvus.proto.msg.DropPartitionRequest")
	proto.RegisterType((*TimeTickMsg)(nil), "milvus.proto.msg.TimeTickMsg")
	proto.RegisterType((*DataNodeTtMsg)(nil), "milvus.proto.msg.DataNodeTtMsg")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 832 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x56, 0x4d, 0x8f, 0xdb, 0x44,
	0x18, 0xc6, 0xb1, 0x37, 0x1f, 0xaf, 0x9d, 0xec, 0x32, 0x2c, 0x5b, 0x77, 0x55, 0x55, 0xae, 0xf9,
	0x8a, 0x10, 0xcd, 0x96, 0xb4, 0xe2, 0x82, 0x50, 0xa5, 0x5d, 0x0b, 0x88, 0xd0, 0x56, 0xab, 0xe9,
	0x8a, 0x03, 0x97, 0x68, 0x92, 0x0c, 0x8e, 0x55, 0x8f, 0xc7, 0x78, 0xc6, 0x1b, 0xe5, 0xc6, 0x1d,
	0xf1, 0x57, 0xf8, 0x11, 0x9c, 0x40, 0xe2, 0x47, 0x21, 0xcf, 0x38, 0x71, 0x9c, 0xa4, 0x2d, 0x54,
	0x02, 0xf5, 0xb0, 0x37, 0xbf, 0xcf, 0xfb, 0x35, 0xf3, 0x3e, 0x8f, 0x5e, 0x0f, 0x74, 0x98, 0x08,
	0x07, 0x69, 0xc6, 0x25, 0x47, 0x47, 0x2c, 0x8a, 0x6f, 0x72, 0xa1, 0xad, 0x01, 0x13, 0xe1, 0xa9,
	0x33, 0xe5, 0x8c, 0xf1, 0x44, 0x23, 0xa7, 0x8e, 0x98, 0xce, 0x29, 0x23, 0xda, 0xf2, 0xff, 0xb4,
	0xa0, 0x3b, 0x4a, 0x04, 0xcd, 0x24, 0xa6, 0x3f, 0xe5, 0x54, 0x48, 0xf4, 0x08, 0xac, 0x09, 0x11,
	0xd4, 0x35, 0x3c, 0xa3, 0x6f, 0x0f, 0xef, 0x0d, 0x6a, 0xe5, 0xca, 0x4a, 0x97, 0x22, 0x3c, 0x27,
	0x82, 0x62, 0x15, 0x89, 0xee, 0x41, 0x47, 0xcc, 0x49, 0x36, 0x7b, 0x46, 0x18, 0x75, 0x1b, 0x9e,
	0xd1, 0xef, 0xe0, 0x0a, 0x40, 0x77, 0xa0, 0x35, 0x9b, 0x8c, 0x93, 0xc2, 0x67, 0x2a, 0x5f, 0x73,
	0x36, 0x51, 0x8e, 0x4f, 0xe0, 0x70, 0xca, 0xe3, 0x98, 0x4e, 0x65, 0xc4, 0x13, 0x1d, 0x60, 0xa9,
	0x80, 0x5e, 0x05, 0xab, 0xc0, 0x8f, 0xa0, 0x97, 0x92, 0x4c, 0x46, 0x55, 0xdc, 0x81, 0x8a, 0xeb,
	0xae, 0x51, 0x15, 0x86, 0xc0, 0x9a, 0x4d, 0x46, 0x81, 0xdb, 0xf4, 0x8c, 0xbe, 0x89, 0xd5, 0x37,
	0xf2, 0xc1, 0xa9, 0x8a, 0x8d, 0x02, 0xb7, 0xa5, 0x7c, 0x35, 0x0c, 0x79, 0x60, 0xaf, 0x0b, 0x8d,
	0x02, 0xb7, 0xad, 0x42, 0x36, 0x21, 0x75, 0x41, 0x1a, 0x32, 0x9a, 0xc8, 0x51, 0xe0, 0x76, 0x94,
	0xbf, 0x02, 0xd0, 0x7d, 0x00, 0x19, 0x31, 0x2a, 0x24, 0x61, 0xa9, 0x70, 0xc1, 0x33, 0xfb, 0x16,
	0xde, 0x40, 0xd0, 0x09, 0x34, 0x33, 0xbe, 0x18, 0x05, 0xc2, 0xb5, 0x3d, 0xb3, 0x6f, 0xe2, 0xd2,
	0x42, 0x4f, 0xa0, 0x9d, 0xf1, 0xc5, 0x78, 0x46, 0x24, 0x71, 0x1d, 0xcf, 0xec, 0xdb, 0xc3, 0xbb,
	0x7b, 0x87, 0x7d, 0x1e, 0xf3, 0x09, 0x6e, 0x65, 0x7c, 0x11, 0x10, 0x49, 0xd0, 0x53, 0xb0, 0x7f,
	0x8c, 0x68, 0x3c, 0x13, 0x3a, 0xb1, 0xab, 0x12, 0xef, 0xd7, 0x13, 0x4b, 0x86, 0xbf, 0x2e, 0xe2,
	0x8a, 0x24, 0x0c, 0x3a, 0x45, 0x15, 0xb8, 0x0b, 0xed, 0x24, 0x67, 0xe3, 0x8c, 0x2f, 0x84, 0xdb,
	0xf3, 0x8c, 0xbe, 0x85, 0x5b, 0x49, 0xce, 0x30, 0x5f, 0x08, 0xf4, 0x15, 0xb4, 0x6e, 0x68, 0x26,
	0x22, 0x9e, 0xb8, 0x87, 0x9e, 0xd1, 0xef, 0x0d, 0x3f, 0x18, 0x6c, 0x8b, 0x69, 0xa0, 0xc5, 0x52,
	0x54, 0xfa, 0x5e, 0x87, 0xe2, 0x55, 0x8e, 0xff, 0x87, 0x09, 0xdd, 0x80, 0xc6, 0x54, 0xd2, 0x5b,
	0x2d, 0xbd, 0x54, 0x4b, 0x9f, 0x01, 0x8a, 0x12, 0xf9, 0xc5, 0x93, 0x71, 0x9a, 0x45, 0x8c, 0x64,
	0xcb, 0xf1, 0x0b, 0xba, 0x14, 0x6e, 0x47, 0x29, 0xe3, 0x48, 0x79, 0xae, 0xb4, 0xe3, 0x3b, 0xba,
	0x14, 0xaf, 0xd5, 0xd6, 0x26, 0x99, 0xb6, 0x6a, 0xb6, 0x26, 0xf3, 0x4b, 0x70, 0x6a, 0x2d, 0x1c,
	0xc5, 0x81, 0xbb, 0x57, 0x29, 0xa3, 0x40, 0x60, 0x3b, 0xad, 0xfa, 0xfa, 0x3f, 0x1b, 0x60, 0x5f,
	0x8a, 0xf0, 0x8a, 0x0b, 0x75, 0x6e, 0xf4, 0x00, 0x9c, 0xe9, 0x9c, 0x24, 0x09, 0x8d, 0xf5, 0xd0,
	0x0c, 0x35, 0x34, 0xbb, 0xc4, 0xd4, 0xc8, 0x8e, 0xe1, 0x80, 0x89, 0x70, 0x14, 0x28, 0xd6, 0x1c,
	0xac, 0x0d, 0x74, 0x0a, 0x6d, 0x26, 0xc2, 0x6f, 0x32, 0x9e, 0xa7, 0x25, 0x65, 0x6b, 0xbb, 0xe0,
	0x7a, 0x7d, 0x15, 0x45, 0x97, 0x85, 0x2b, 0xc0, 0xff, 0xcd, 0x84, 0x3b, 0x17, 0x19, 0x25, 0x92,
	0x5e, 0xac, 0x27, 0xfc, 0xe6, 0xba, 0xda, 0x50, 0x4e, 0xa3, 0xa6, 0x9c, 0x8f, 0x61, 0x4b, 0x22,
	0xe5, 0x31, 0xb7, 0x85, 0xf3, 0x21, 0xd4, 0x25, 0x52, 0xea, 0xeb, 0x25, 0xba, 0x39, 0x78, 0x85,
	0x6e, 0x9a, 0xaf, 0xd7, 0x4d, 0x6b, 0x57, 0x37, 0x27, 0xd0, 0xd4, 0x64, 0x29, 0x51, 0x39, 0xb8,
	0xb4, 0xd0, 0x23, 0x78, 0xef, 0x26, 0xca, 0x64, 0x4e, 0xe2, 0x8b, 0x8a, 0x0c, 0x2d, 0xa8, 0x0e,
	0xde, 0xe7, 0x42, 0x43, 0x38, 0x4e, 0xe7, 0x4b, 0x11, 0x4d, 0xb7, 0x52, 0x40, 0xa5, 0xec, 0xf5,
	0x15, 0x77, 0xd8, 0x38, 0xcc, 0x6a, 0x93, 0xd5, 0x30, 0xff, 0x77, 0x03, 0xde, 0x0f, 0x32, 0x9e,
	0xbe, 0x15, 0x74, 0xad, 0x88, 0xb0, 0x5e, 0x41, 0xc4, 0xc1, 0x2e, 0x11, 0xfe, 0xaf, 0x0d, 0x38,
	0xd1, 0xaa, 0xbb, 0x5a, 0xdd, 0xed, 0x3f, 0xb8, 0xc5, 0x9e, 0x75, 0x65, 0xfe, 0xc3, 0x75, 0xf5,
	0xff, 0xca, 0xce, 0xff, 0xa5, 0x01, 0xc7, 0x05, 0xa9, 0xb7, 0xd3, 0x28, 0xa6, 0xf1, 0x14, 0xec,
	0xeb, 0x88, 0xd1, 0xeb, 0x68, 0xfa, 0xe2, 0x52, 0x84, 0xff, 0x7e, 0x06, 0xfe, 0x5f, 0x06, 0x74,
	0x8b, 0x7f, 0xe7, 0x33, 0x3e, 0xa3, 0xd7, 0xf2, 0x8d, 0x6a, 0xec, 0xec, 0xe2, 0xc6, 0xee, 0x2e,
	0xae, 0x6d, 0x56, 0x73, 0x6b, 0xb3, 0xa2, 0x6f, 0xa1, 0x57, 0xbe, 0x5e, 0xc4, 0x58, 0x48, 0x22,
	0x85, 0x6b, 0xa9, 0x57, 0xc4, 0x83, 0xbd, 0xcd, 0x9f, 0xeb, 0xd0, 0xe7, 0x45, 0x20, 0xee, 0xae,
	0x12, 0x95, 0xf9, 0xe9, 0x10, 0xde, 0xdd, 0x79, 0x0f, 0x20, 0x07, 0xda, 0x98, 0x2f, 0x8a, 0x03,
	0xcf, 0x8e, 0xde, 0x41, 0x87, 0x60, 0x5f, 0xf0, 0x38, 0x67, 0x89, 0x06, 0x8c, 0xf3, 0xc7, 0x3f,
	0x7c, 0x1e, 0x46, 0x72, 0x9e, 0x4f, 0x8a, 0x06, 0x67, 0xba, 0xe3, 0xc3, 0x88, 0xaf, 0xbe, 0x54,
	0xef, 0xb3, 0x90, 0x3f, 0x24, 0x69, 0x74, 0x76, 0x33, 0x3c, 0x63, 0x22, 0x4c, 0x27, 0x93, 0xa6,
	0x82, 0x1f, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x14, 0x3c, 0x92, 0xe8, 0x0a, 0x00, 0x00,
}
