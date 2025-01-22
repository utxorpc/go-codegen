// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: utxorpc/v1alpha/watch/watch.proto

package watch

import (
	cardano "github.com/utxorpc/go-codegen/utxorpc/v1alpha/cardano"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents a reference to a specific block
type BlockRef struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Index         uint64                 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"` // Height or slot number (depending on the blockchain)
	Hash          []byte                 `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`    // Hash of the content of the block
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BlockRef) Reset() {
	*x = BlockRef{}
	mi := &file_utxorpc_v1alpha_watch_watch_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BlockRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockRef) ProtoMessage() {}

func (x *BlockRef) ProtoReflect() protoreflect.Message {
	mi := &file_utxorpc_v1alpha_watch_watch_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockRef.ProtoReflect.Descriptor instead.
func (*BlockRef) Descriptor() ([]byte, []int) {
	return file_utxorpc_v1alpha_watch_watch_proto_rawDescGZIP(), []int{0}
}

func (x *BlockRef) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *BlockRef) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

// Represents a tx pattern from any supported blockchain.
type AnyChainTxPattern struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Chain:
	//
	//	*AnyChainTxPattern_Cardano
	Chain         isAnyChainTxPattern_Chain `protobuf_oneof:"chain"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AnyChainTxPattern) Reset() {
	*x = AnyChainTxPattern{}
	mi := &file_utxorpc_v1alpha_watch_watch_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnyChainTxPattern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnyChainTxPattern) ProtoMessage() {}

func (x *AnyChainTxPattern) ProtoReflect() protoreflect.Message {
	mi := &file_utxorpc_v1alpha_watch_watch_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnyChainTxPattern.ProtoReflect.Descriptor instead.
func (*AnyChainTxPattern) Descriptor() ([]byte, []int) {
	return file_utxorpc_v1alpha_watch_watch_proto_rawDescGZIP(), []int{1}
}

func (x *AnyChainTxPattern) GetChain() isAnyChainTxPattern_Chain {
	if x != nil {
		return x.Chain
	}
	return nil
}

func (x *AnyChainTxPattern) GetCardano() *cardano.TxPattern {
	if x != nil {
		if x, ok := x.Chain.(*AnyChainTxPattern_Cardano); ok {
			return x.Cardano
		}
	}
	return nil
}

type isAnyChainTxPattern_Chain interface {
	isAnyChainTxPattern_Chain()
}

type AnyChainTxPattern_Cardano struct {
	Cardano *cardano.TxPattern `protobuf:"bytes,1,opt,name=cardano,proto3,oneof"` // A Cardano tx pattern.
}

func (*AnyChainTxPattern_Cardano) isAnyChainTxPattern_Chain() {}

// Represents a simple tx predicate that can composed to create more complex ones
type TxPredicate struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Match         *AnyChainTxPattern     `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`              // Predicate is true if tx exhibits pattern.
	Not           []*TxPredicate         `protobuf:"bytes,2,rep,name=not,proto3" json:"not,omitempty"`                  // Predicate is true if tx doesn't exhibit pattern.
	AllOf         []*TxPredicate         `protobuf:"bytes,3,rep,name=all_of,json=allOf,proto3" json:"all_of,omitempty"` // Predicate is true if tx exhibits all of the patterns.
	AnyOf         []*TxPredicate         `protobuf:"bytes,4,rep,name=any_of,json=anyOf,proto3" json:"any_of,omitempty"` // Predicate is true if tx exhibits any of the patterns.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TxPredicate) Reset() {
	*x = TxPredicate{}
	mi := &file_utxorpc_v1alpha_watch_watch_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TxPredicate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxPredicate) ProtoMessage() {}

func (x *TxPredicate) ProtoReflect() protoreflect.Message {
	mi := &file_utxorpc_v1alpha_watch_watch_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxPredicate.ProtoReflect.Descriptor instead.
func (*TxPredicate) Descriptor() ([]byte, []int) {
	return file_utxorpc_v1alpha_watch_watch_proto_rawDescGZIP(), []int{2}
}

func (x *TxPredicate) GetMatch() *AnyChainTxPattern {
	if x != nil {
		return x.Match
	}
	return nil
}

func (x *TxPredicate) GetNot() []*TxPredicate {
	if x != nil {
		return x.Not
	}
	return nil
}

func (x *TxPredicate) GetAllOf() []*TxPredicate {
	if x != nil {
		return x.AllOf
	}
	return nil
}

func (x *TxPredicate) GetAnyOf() []*TxPredicate {
	if x != nil {
		return x.AnyOf
	}
	return nil
}

// Request to watch transactions from the chain based on a set of predicates.
type WatchTxRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Predicate     *TxPredicate           `protobuf:"bytes,1,opt,name=predicate,proto3" json:"predicate,omitempty"`                  // Predicate to filter transactions by.
	FieldMask     *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"` // Field mask to selectively return fields.
	Intersect     []*BlockRef            `protobuf:"bytes,3,rep,name=intersect,proto3" json:"intersect,omitempty"`                  // List of block references to find the intersection.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WatchTxRequest) Reset() {
	*x = WatchTxRequest{}
	mi := &file_utxorpc_v1alpha_watch_watch_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WatchTxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchTxRequest) ProtoMessage() {}

func (x *WatchTxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_utxorpc_v1alpha_watch_watch_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchTxRequest.ProtoReflect.Descriptor instead.
func (*WatchTxRequest) Descriptor() ([]byte, []int) {
	return file_utxorpc_v1alpha_watch_watch_proto_rawDescGZIP(), []int{3}
}

func (x *WatchTxRequest) GetPredicate() *TxPredicate {
	if x != nil {
		return x.Predicate
	}
	return nil
}

func (x *WatchTxRequest) GetFieldMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.FieldMask
	}
	return nil
}

func (x *WatchTxRequest) GetIntersect() []*BlockRef {
	if x != nil {
		return x.Intersect
	}
	return nil
}

// Represents a transaction from any supported blockchain.
type AnyChainTx struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Chain:
	//
	//	*AnyChainTx_Cardano
	Chain         isAnyChainTx_Chain `protobuf_oneof:"chain"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AnyChainTx) Reset() {
	*x = AnyChainTx{}
	mi := &file_utxorpc_v1alpha_watch_watch_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnyChainTx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnyChainTx) ProtoMessage() {}

func (x *AnyChainTx) ProtoReflect() protoreflect.Message {
	mi := &file_utxorpc_v1alpha_watch_watch_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnyChainTx.ProtoReflect.Descriptor instead.
func (*AnyChainTx) Descriptor() ([]byte, []int) {
	return file_utxorpc_v1alpha_watch_watch_proto_rawDescGZIP(), []int{4}
}

func (x *AnyChainTx) GetChain() isAnyChainTx_Chain {
	if x != nil {
		return x.Chain
	}
	return nil
}

func (x *AnyChainTx) GetCardano() *cardano.Tx {
	if x != nil {
		if x, ok := x.Chain.(*AnyChainTx_Cardano); ok {
			return x.Cardano
		}
	}
	return nil
}

type isAnyChainTx_Chain interface {
	isAnyChainTx_Chain()
}

type AnyChainTx_Cardano struct {
	Cardano *cardano.Tx `protobuf:"bytes,1,opt,name=cardano,proto3,oneof"` // A Cardano transaction.
}

func (*AnyChainTx_Cardano) isAnyChainTx_Chain() {}

// Response containing the matching chain transactions.
type WatchTxResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Action:
	//
	//	*WatchTxResponse_Apply
	//	*WatchTxResponse_Undo
	Action        isWatchTxResponse_Action `protobuf_oneof:"action"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WatchTxResponse) Reset() {
	*x = WatchTxResponse{}
	mi := &file_utxorpc_v1alpha_watch_watch_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WatchTxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchTxResponse) ProtoMessage() {}

func (x *WatchTxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_utxorpc_v1alpha_watch_watch_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchTxResponse.ProtoReflect.Descriptor instead.
func (*WatchTxResponse) Descriptor() ([]byte, []int) {
	return file_utxorpc_v1alpha_watch_watch_proto_rawDescGZIP(), []int{5}
}

func (x *WatchTxResponse) GetAction() isWatchTxResponse_Action {
	if x != nil {
		return x.Action
	}
	return nil
}

func (x *WatchTxResponse) GetApply() *AnyChainTx {
	if x != nil {
		if x, ok := x.Action.(*WatchTxResponse_Apply); ok {
			return x.Apply
		}
	}
	return nil
}

func (x *WatchTxResponse) GetUndo() *AnyChainTx {
	if x != nil {
		if x, ok := x.Action.(*WatchTxResponse_Undo); ok {
			return x.Undo
		}
	}
	return nil
}

type isWatchTxResponse_Action interface {
	isWatchTxResponse_Action()
}

type WatchTxResponse_Apply struct {
	Apply *AnyChainTx `protobuf:"bytes,1,opt,name=apply,proto3,oneof"` // Apply this transaction.
}

type WatchTxResponse_Undo struct {
	Undo *AnyChainTx `protobuf:"bytes,2,opt,name=undo,proto3,oneof"` // Undo this transaction.
}

func (*WatchTxResponse_Apply) isWatchTxResponse_Action() {}

func (*WatchTxResponse_Undo) isWatchTxResponse_Action() {}

var File_utxorpc_v1alpha_watch_watch_proto protoreflect.FileDescriptor

var file_utxorpc_v1alpha_watch_watch_proto_rawDesc = []byte{
	0x0a, 0x21, 0x75, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x75, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x75, 0x74,
	0x78, 0x6f, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x61,
	0x72, 0x64, 0x61, 0x6e, 0x6f, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x61, 0x6e, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x08, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x66, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x5c, 0x0a, 0x11, 0x41, 0x6e, 0x79,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x78, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x3e,
	0x0a, 0x07, 0x63, 0x61, 0x72, 0x64, 0x61, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x75, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x61, 0x6e, 0x6f, 0x2e, 0x54, 0x78, 0x50, 0x61, 0x74, 0x74,
	0x65, 0x72, 0x6e, 0x48, 0x00, 0x52, 0x07, 0x63, 0x61, 0x72, 0x64, 0x61, 0x6e, 0x6f, 0x42, 0x07,
	0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x22, 0xf9, 0x01, 0x0a, 0x0b, 0x54, 0x78, 0x50, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x3e, 0x0a, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x75, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x41,
	0x6e, 0x79, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x78, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e,
	0x52, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x34, 0x0a, 0x03, 0x6e, 0x6f, 0x74, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x75, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x54, 0x78, 0x50,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x03, 0x6e, 0x6f, 0x74, 0x12, 0x39, 0x0a,
	0x06, 0x61, 0x6c, 0x6c, 0x5f, 0x6f, 0x66, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x75, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x54, 0x78, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x52, 0x05, 0x61, 0x6c, 0x6c, 0x4f, 0x66, 0x12, 0x39, 0x0a, 0x06, 0x61, 0x6e, 0x79, 0x5f,
	0x6f, 0x66, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x75, 0x74, 0x78, 0x6f, 0x72,
	0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x2e, 0x54, 0x78, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x05, 0x61, 0x6e,
	0x79, 0x4f, 0x66, 0x22, 0xcc, 0x01, 0x0a, 0x0e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x54, 0x78, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x75, 0x74, 0x78, 0x6f,
	0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x2e, 0x54, 0x78, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x09, 0x70,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x61, 0x73, 0x6b, 0x12, 0x3d, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x63, 0x74,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x75, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x66, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65,
	0x63, 0x74, 0x22, 0x4e, 0x0a, 0x0a, 0x41, 0x6e, 0x79, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x78,
	0x12, 0x37, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x64, 0x61, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x75, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x61, 0x6e, 0x6f, 0x2e, 0x54, 0x78, 0x48, 0x00,
	0x52, 0x07, 0x63, 0x61, 0x72, 0x64, 0x61, 0x6e, 0x6f, 0x42, 0x07, 0x0a, 0x05, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x22, 0x8f, 0x01, 0x0a, 0x0f, 0x57, 0x61, 0x74, 0x63, 0x68, 0x54, 0x78, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x75, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x41, 0x6e,
	0x79, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x78, 0x48, 0x00, 0x52, 0x05, 0x61, 0x70, 0x70, 0x6c,
	0x79, 0x12, 0x37, 0x0a, 0x04, 0x75, 0x6e, 0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x75, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x41, 0x6e, 0x79, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x54, 0x78, 0x48, 0x00, 0x52, 0x04, 0x75, 0x6e, 0x64, 0x6f, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x32, 0x6a, 0x0a, 0x0c, 0x57, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x07, 0x57, 0x61, 0x74, 0x63, 0x68, 0x54, 0x78, 0x12,
	0x25, 0x2e, 0x75, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x54, 0x78, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x75, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01,
	0x42, 0xd2, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x42, 0x0a,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63,
	0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x75, 0x74, 0x78, 0x6f,
	0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x77, 0x61, 0x74, 0x63,
	0x68, 0xa2, 0x02, 0x03, 0x55, 0x56, 0x57, 0xaa, 0x02, 0x15, 0x55, 0x74, 0x78, 0x6f, 0x72, 0x70,
	0x63, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0xca,
	0x02, 0x15, 0x55, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x5c, 0x57, 0x61, 0x74, 0x63, 0x68, 0xe2, 0x02, 0x21, 0x55, 0x74, 0x78, 0x6f, 0x72, 0x70,
	0x63, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5c, 0x57, 0x61, 0x74, 0x63, 0x68, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x17, 0x55, 0x74,
	0x78, 0x6f, 0x72, 0x70, 0x63, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3a, 0x3a,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_utxorpc_v1alpha_watch_watch_proto_rawDescOnce sync.Once
	file_utxorpc_v1alpha_watch_watch_proto_rawDescData = file_utxorpc_v1alpha_watch_watch_proto_rawDesc
)

func file_utxorpc_v1alpha_watch_watch_proto_rawDescGZIP() []byte {
	file_utxorpc_v1alpha_watch_watch_proto_rawDescOnce.Do(func() {
		file_utxorpc_v1alpha_watch_watch_proto_rawDescData = protoimpl.X.CompressGZIP(file_utxorpc_v1alpha_watch_watch_proto_rawDescData)
	})
	return file_utxorpc_v1alpha_watch_watch_proto_rawDescData
}

var file_utxorpc_v1alpha_watch_watch_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_utxorpc_v1alpha_watch_watch_proto_goTypes = []any{
	(*BlockRef)(nil),              // 0: utxorpc.v1alpha.watch.BlockRef
	(*AnyChainTxPattern)(nil),     // 1: utxorpc.v1alpha.watch.AnyChainTxPattern
	(*TxPredicate)(nil),           // 2: utxorpc.v1alpha.watch.TxPredicate
	(*WatchTxRequest)(nil),        // 3: utxorpc.v1alpha.watch.WatchTxRequest
	(*AnyChainTx)(nil),            // 4: utxorpc.v1alpha.watch.AnyChainTx
	(*WatchTxResponse)(nil),       // 5: utxorpc.v1alpha.watch.WatchTxResponse
	(*cardano.TxPattern)(nil),     // 6: utxorpc.v1alpha.cardano.TxPattern
	(*fieldmaskpb.FieldMask)(nil), // 7: google.protobuf.FieldMask
	(*cardano.Tx)(nil),            // 8: utxorpc.v1alpha.cardano.Tx
}
var file_utxorpc_v1alpha_watch_watch_proto_depIdxs = []int32{
	6,  // 0: utxorpc.v1alpha.watch.AnyChainTxPattern.cardano:type_name -> utxorpc.v1alpha.cardano.TxPattern
	1,  // 1: utxorpc.v1alpha.watch.TxPredicate.match:type_name -> utxorpc.v1alpha.watch.AnyChainTxPattern
	2,  // 2: utxorpc.v1alpha.watch.TxPredicate.not:type_name -> utxorpc.v1alpha.watch.TxPredicate
	2,  // 3: utxorpc.v1alpha.watch.TxPredicate.all_of:type_name -> utxorpc.v1alpha.watch.TxPredicate
	2,  // 4: utxorpc.v1alpha.watch.TxPredicate.any_of:type_name -> utxorpc.v1alpha.watch.TxPredicate
	2,  // 5: utxorpc.v1alpha.watch.WatchTxRequest.predicate:type_name -> utxorpc.v1alpha.watch.TxPredicate
	7,  // 6: utxorpc.v1alpha.watch.WatchTxRequest.field_mask:type_name -> google.protobuf.FieldMask
	0,  // 7: utxorpc.v1alpha.watch.WatchTxRequest.intersect:type_name -> utxorpc.v1alpha.watch.BlockRef
	8,  // 8: utxorpc.v1alpha.watch.AnyChainTx.cardano:type_name -> utxorpc.v1alpha.cardano.Tx
	4,  // 9: utxorpc.v1alpha.watch.WatchTxResponse.apply:type_name -> utxorpc.v1alpha.watch.AnyChainTx
	4,  // 10: utxorpc.v1alpha.watch.WatchTxResponse.undo:type_name -> utxorpc.v1alpha.watch.AnyChainTx
	3,  // 11: utxorpc.v1alpha.watch.WatchService.WatchTx:input_type -> utxorpc.v1alpha.watch.WatchTxRequest
	5,  // 12: utxorpc.v1alpha.watch.WatchService.WatchTx:output_type -> utxorpc.v1alpha.watch.WatchTxResponse
	12, // [12:13] is the sub-list for method output_type
	11, // [11:12] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_utxorpc_v1alpha_watch_watch_proto_init() }
func file_utxorpc_v1alpha_watch_watch_proto_init() {
	if File_utxorpc_v1alpha_watch_watch_proto != nil {
		return
	}
	file_utxorpc_v1alpha_watch_watch_proto_msgTypes[1].OneofWrappers = []any{
		(*AnyChainTxPattern_Cardano)(nil),
	}
	file_utxorpc_v1alpha_watch_watch_proto_msgTypes[4].OneofWrappers = []any{
		(*AnyChainTx_Cardano)(nil),
	}
	file_utxorpc_v1alpha_watch_watch_proto_msgTypes[5].OneofWrappers = []any{
		(*WatchTxResponse_Apply)(nil),
		(*WatchTxResponse_Undo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_utxorpc_v1alpha_watch_watch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_utxorpc_v1alpha_watch_watch_proto_goTypes,
		DependencyIndexes: file_utxorpc_v1alpha_watch_watch_proto_depIdxs,
		MessageInfos:      file_utxorpc_v1alpha_watch_watch_proto_msgTypes,
	}.Build()
	File_utxorpc_v1alpha_watch_watch_proto = out.File
	file_utxorpc_v1alpha_watch_watch_proto_rawDesc = nil
	file_utxorpc_v1alpha_watch_watch_proto_goTypes = nil
	file_utxorpc_v1alpha_watch_watch_proto_depIdxs = nil
}
