// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: cluster.proto

package repmanv3

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_cluster_proto protoreflect.FileDescriptor

var file_cluster_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x31, 0x38, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x33,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xbd, 0x02, 0x0a, 0x14, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x8d, 0x01, 0x0a, 0x0d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x31, 0x38, 0x2e,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x2e,
	0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x31, 0x38, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x33,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x22,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x94, 0x01, 0x0a, 0x14, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x50, 0x68, 0x79,
	0x73, 0x69, 0x63, 0x61, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x28, 0x2e, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x6c, 0x31, 0x38, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3a, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x34, 0x12, 0x32, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63,
	0x61, 0x6c, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x32, 0xdb, 0x0a, 0x0a, 0x0e, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6c, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x28, 0x2e, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x31, 0x38, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x1b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x80, 0x01, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x46, 0x6f, 0x72, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x28, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x31, 0x38, 0x2e,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12,
	0x1c, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x6e,
	0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x84, 0x03,
	0x0a, 0x1b, 0x53, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x2f, 0x2e,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x31, 0x38, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x9b, 0x02, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x94, 0x02,
	0x12, 0x35, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x5a, 0x43, 0x12, 0x41, 0x2f, 0x76, 0x33, 0x2f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7d,
	0x2f, 0x7b, 0x74, 0x61, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x7d, 0x5a, 0x51, 0x12, 0x4f,
	0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x74,
	0x2f, 0x7b, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f,
	0x7b, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x7d, 0x5a,
	0x43, 0x12, 0x41, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f,
	0x7b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2f, 0x7b, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2e, 0x6e,
	0x61, 0x6d, 0x65, 0x7d, 0x12, 0xaf, 0x04, 0x0a, 0x14, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x2e,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x31, 0x38, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xce, 0x03, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0xc7, 0x03, 0x12,
	0x2c, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x7b,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x5a, 0x4c, 0x12,
	0x4a, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x7b, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x68,
	0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x7d, 0x5a, 0x4a, 0x12, 0x48, 0x2f,
	0x76, 0x33, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x7b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x7b, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x7d, 0x2f, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x7d, 0x5a, 0x58, 0x12, 0x56, 0x2f, 0x76, 0x33, 0x2f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x68, 0x6f, 0x73, 0x74, 0x7d, 0x2f, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x6f,
	0x72, 0x74, 0x7d, 0x2f, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x7d, 0x5a, 0x39, 0x12, 0x37, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d,
	0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x7d, 0x2f, 0x7b, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x7d, 0x5a, 0x2e, 0x12, 0x2c,
	0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x5a, 0x38, 0x22, 0x2e,
	0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x6f, 0x76, 0x65, 0x72, 0x3a, 0x06,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x9e, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x12,
	0x32, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x31, 0x38, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76,
	0x33, 0x2e, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x61, 0x6c, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x37, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x31, 0x12, 0x2f, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65,
	0x7d, 0x2f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2f, 0x7b, 0x72, 0x65, 0x74, 0x72,
	0x69, 0x65, 0x76, 0x65, 0x7d, 0x30, 0x01, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x31, 0x38, 0x2f, 0x72,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x3b, 0x72, 0x65, 0x70, 0x6d, 0x61, 0x6e, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_cluster_proto_goTypes = []interface{}{
	(*Cluster)(nil),           // 0: signal18.replication_manager.v3.Cluster
	(*ClusterSetting)(nil),    // 1: signal18.replication_manager.v3.ClusterSetting
	(*ClusterAction)(nil),     // 2: signal18.replication_manager.v3.ClusterAction
	(*TopologyRetrieval)(nil), // 3: signal18.replication_manager.v3.TopologyRetrieval
	(*StatusMessage)(nil),     // 4: signal18.replication_manager.v3.StatusMessage
	(*emptypb.Empty)(nil),     // 5: google.protobuf.Empty
	(*structpb.Struct)(nil),   // 6: google.protobuf.Struct
}
var file_cluster_proto_depIdxs = []int32{
	0, // 0: signal18.replication_manager.v3.ClusterPublicService.ClusterStatus:input_type -> signal18.replication_manager.v3.Cluster
	0, // 1: signal18.replication_manager.v3.ClusterPublicService.MasterPhysicalBackup:input_type -> signal18.replication_manager.v3.Cluster
	0, // 2: signal18.replication_manager.v3.ClusterService.GetCluster:input_type -> signal18.replication_manager.v3.Cluster
	0, // 3: signal18.replication_manager.v3.ClusterService.GetSettingsForCluster:input_type -> signal18.replication_manager.v3.Cluster
	1, // 4: signal18.replication_manager.v3.ClusterService.SetActionForClusterSettings:input_type -> signal18.replication_manager.v3.ClusterSetting
	2, // 5: signal18.replication_manager.v3.ClusterService.PerformClusterAction:input_type -> signal18.replication_manager.v3.ClusterAction
	3, // 6: signal18.replication_manager.v3.ClusterService.RetrieveFromTopology:input_type -> signal18.replication_manager.v3.TopologyRetrieval
	4, // 7: signal18.replication_manager.v3.ClusterPublicService.ClusterStatus:output_type -> signal18.replication_manager.v3.StatusMessage
	5, // 8: signal18.replication_manager.v3.ClusterPublicService.MasterPhysicalBackup:output_type -> google.protobuf.Empty
	6, // 9: signal18.replication_manager.v3.ClusterService.GetCluster:output_type -> google.protobuf.Struct
	6, // 10: signal18.replication_manager.v3.ClusterService.GetSettingsForCluster:output_type -> google.protobuf.Struct
	5, // 11: signal18.replication_manager.v3.ClusterService.SetActionForClusterSettings:output_type -> google.protobuf.Empty
	5, // 12: signal18.replication_manager.v3.ClusterService.PerformClusterAction:output_type -> google.protobuf.Empty
	6, // 13: signal18.replication_manager.v3.ClusterService.RetrieveFromTopology:output_type -> google.protobuf.Struct
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cluster_proto_init() }
func file_cluster_proto_init() {
	if File_cluster_proto != nil {
		return
	}
	file_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_cluster_proto_goTypes,
		DependencyIndexes: file_cluster_proto_depIdxs,
	}.Build()
	File_cluster_proto = out.File
	file_cluster_proto_rawDesc = nil
	file_cluster_proto_goTypes = nil
	file_cluster_proto_depIdxs = nil
}
