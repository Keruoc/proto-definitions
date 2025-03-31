// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: services/kss/v1/kss_service.proto

package kssv1

import (
	v1 "github.com/kerouc/proto-definitions/common/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type KSSRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Transcript     string                 `protobuf:"bytes,1,opt,name=transcript,proto3" json:"transcript,omitempty"`
	ConfirmedNodes []*v1.Node             `protobuf:"bytes,2,rep,name=confirmed_nodes,json=confirmedNodes,proto3" json:"confirmed_nodes,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *KSSRequest) Reset() {
	*x = KSSRequest{}
	mi := &file_services_kss_v1_kss_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KSSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KSSRequest) ProtoMessage() {}

func (x *KSSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_kss_v1_kss_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KSSRequest.ProtoReflect.Descriptor instead.
func (*KSSRequest) Descriptor() ([]byte, []int) {
	return file_services_kss_v1_kss_service_proto_rawDescGZIP(), []int{0}
}

func (x *KSSRequest) GetTranscript() string {
	if x != nil {
		return x.Transcript
	}
	return ""
}

func (x *KSSRequest) GetConfirmedNodes() []*v1.Node {
	if x != nil {
		return x.ConfirmedNodes
	}
	return nil
}

type SimilarNode struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title           string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	SimilarityScore float32                `protobuf:"fixed32,3,opt,name=similarity_score,json=similarityScore,proto3" json:"similarity_score,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *SimilarNode) Reset() {
	*x = SimilarNode{}
	mi := &file_services_kss_v1_kss_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SimilarNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimilarNode) ProtoMessage() {}

func (x *SimilarNode) ProtoReflect() protoreflect.Message {
	mi := &file_services_kss_v1_kss_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimilarNode.ProtoReflect.Descriptor instead.
func (*SimilarNode) Descriptor() ([]byte, []int) {
	return file_services_kss_v1_kss_service_proto_rawDescGZIP(), []int{1}
}

func (x *SimilarNode) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SimilarNode) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SimilarNode) GetSimilarityScore() float32 {
	if x != nil {
		return x.SimilarityScore
	}
	return 0
}

type KSSResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SimilarNodes  []*SimilarNode         `protobuf:"bytes,1,rep,name=similar_nodes,json=similarNodes,proto3" json:"similar_nodes,omitempty"`
	CurrentLink   []*v1.Link             `protobuf:"bytes,2,rep,name=current_link,json=currentLink,proto3" json:"current_link,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KSSResponse) Reset() {
	*x = KSSResponse{}
	mi := &file_services_kss_v1_kss_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KSSResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KSSResponse) ProtoMessage() {}

func (x *KSSResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_kss_v1_kss_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KSSResponse.ProtoReflect.Descriptor instead.
func (*KSSResponse) Descriptor() ([]byte, []int) {
	return file_services_kss_v1_kss_service_proto_rawDescGZIP(), []int{2}
}

func (x *KSSResponse) GetSimilarNodes() []*SimilarNode {
	if x != nil {
		return x.SimilarNodes
	}
	return nil
}

func (x *KSSResponse) GetCurrentLink() []*v1.Link {
	if x != nil {
		return x.CurrentLink
	}
	return nil
}

var File_services_kss_v1_kss_service_proto protoreflect.FileDescriptor

const file_services_kss_v1_kss_service_proto_rawDesc = "" +
	"\n" +
	"!services/kss/v1/kss_service.proto\x12\x06kss.v1\x1a\x14common/v1/node.proto\x1a\x14common/v1/link.proto\"f\n" +
	"\n" +
	"KSSRequest\x12\x1e\n" +
	"\n" +
	"transcript\x18\x01 \x01(\tR\n" +
	"transcript\x128\n" +
	"\x0fconfirmed_nodes\x18\x02 \x03(\v2\x0f.common.v1.NodeR\x0econfirmedNodes\"^\n" +
	"\vSimilarNode\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12)\n" +
	"\x10similarity_score\x18\x03 \x01(\x02R\x0fsimilarityScore\"{\n" +
	"\vKSSResponse\x128\n" +
	"\rsimilar_nodes\x18\x01 \x03(\v2\x13.kss.v1.SimilarNodeR\fsimilarNodes\x122\n" +
	"\fcurrent_link\x18\x02 \x03(\v2\x0f.common.v1.LinkR\vcurrentLink2_\n" +
	"\x17KnowledgeStorageService\x12D\n" +
	"\x19GetSuggestedNodesAndLinks\x12\x12.kss.v1.KSSRequest\x1a\x13.kss.v1.KSSResponseB;Z9github.com/kerouc/proto-definitions/services/kss/v1;kssv1b\x06proto3"

var (
	file_services_kss_v1_kss_service_proto_rawDescOnce sync.Once
	file_services_kss_v1_kss_service_proto_rawDescData []byte
)

func file_services_kss_v1_kss_service_proto_rawDescGZIP() []byte {
	file_services_kss_v1_kss_service_proto_rawDescOnce.Do(func() {
		file_services_kss_v1_kss_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_services_kss_v1_kss_service_proto_rawDesc), len(file_services_kss_v1_kss_service_proto_rawDesc)))
	})
	return file_services_kss_v1_kss_service_proto_rawDescData
}

var file_services_kss_v1_kss_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_services_kss_v1_kss_service_proto_goTypes = []any{
	(*KSSRequest)(nil),  // 0: kss.v1.KSSRequest
	(*SimilarNode)(nil), // 1: kss.v1.SimilarNode
	(*KSSResponse)(nil), // 2: kss.v1.KSSResponse
	(*v1.Node)(nil),     // 3: common.v1.Node
	(*v1.Link)(nil),     // 4: common.v1.Link
}
var file_services_kss_v1_kss_service_proto_depIdxs = []int32{
	3, // 0: kss.v1.KSSRequest.confirmed_nodes:type_name -> common.v1.Node
	1, // 1: kss.v1.KSSResponse.similar_nodes:type_name -> kss.v1.SimilarNode
	4, // 2: kss.v1.KSSResponse.current_link:type_name -> common.v1.Link
	0, // 3: kss.v1.KnowledgeStorageService.GetSuggestedNodesAndLinks:input_type -> kss.v1.KSSRequest
	2, // 4: kss.v1.KnowledgeStorageService.GetSuggestedNodesAndLinks:output_type -> kss.v1.KSSResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_services_kss_v1_kss_service_proto_init() }
func file_services_kss_v1_kss_service_proto_init() {
	if File_services_kss_v1_kss_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_kss_v1_kss_service_proto_rawDesc), len(file_services_kss_v1_kss_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_kss_v1_kss_service_proto_goTypes,
		DependencyIndexes: file_services_kss_v1_kss_service_proto_depIdxs,
		MessageInfos:      file_services_kss_v1_kss_service_proto_msgTypes,
	}.Build()
	File_services_kss_v1_kss_service_proto = out.File
	file_services_kss_v1_kss_service_proto_goTypes = nil
	file_services_kss_v1_kss_service_proto_depIdxs = nil
}
