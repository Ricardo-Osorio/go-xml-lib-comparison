package data

import (
	"google.golang.org/protobuf/runtime/protoimpl"
)

// The following vars were changed from the original code.
// From int32 to strings.
type ObjectType string
type StorageClassType string

type Owner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"ID" xml:"ID"`
	DisplayName *string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3,oneof" json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
}

type ObjectMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key          string           `protobuf:"bytes,1,opt,name=key,proto3" json:"Key" xml:"Key"`
	LastModified string           `protobuf:"bytes,2,opt,name=last_modified,json=lastModified,proto3" json:"LastModified" xml:"LastModified"`
	Etag         string           `protobuf:"bytes,3,opt,name=etag,proto3" json:"ETag" xml:"ETag"`
	Size         uint64           `protobuf:"varint,4,opt,name=size,proto3" json:"Size" xml:"Size"`
	StorageClass StorageClassType `protobuf:"varint,5,opt,name=storage_class,json=storageClass,proto3,enum=common.v1.StorageClassType" json:"StorageClass" xml:"StorageClass"`
	Type         ObjectType       `protobuf:"varint,6,opt,name=type,proto3,enum=common.v1.ObjectType" json:"Type" xml:"Type"`
	Owner        *Owner           `protobuf:"bytes,7,opt,name=owner,proto3" json:"Owner,omitempty" xml:"Owner,omitempty"`
}

type CommonPrefixes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prefix string `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
}

// Response message for ListObjectsV2
type ListObjectsV2Response struct {
	// These grpc fields aren't used in the XML parser however, since the native
	// golang xml lib uses reflection, I am leaving them in the struct so as to
	// replicate as much of the original behavior as possible.
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contents              []*ObjectMeta     `protobuf:"bytes,1,rep,name=contents,proto3" json:"Contents,omitempty" xml:"Contents,omitempty"`
	Name                  string            `protobuf:"bytes,2,opt,name=name,proto3" json:"Name" xml:"Name"`
	EncodingType          *string           `protobuf:"bytes,3,opt,name=encoding_type,json=encodingType,proto3,oneof" json:"EncodingType,omitempty" xml:"EncodingType,omitempty"`
	Prefix                string            `protobuf:"bytes,4,opt,name=prefix,proto3" json:"Prefix" xml:"Prefix"`
	Delimiter             *string           `protobuf:"bytes,5,opt,name=delimiter,proto3,oneof" json:"Delimiter" xml:"Delimiter"`
	CommonPrefixes        []*CommonPrefixes `protobuf:"bytes,6,rep,name=common_prefixes,json=commonPrefixes,proto3" json:"CommonPrefixes,omitempty" xml:"CommonPrefixes,omitempty"`
	StartAfter            *string           `protobuf:"bytes,7,opt,name=start_after,json=startAfter,proto3,oneof" json:"StartAfter,omitempty" xml:"StartAfter,omitempty"`
	ContinuationToken     *string           `protobuf:"bytes,8,opt,name=continuation_token,json=continuationToken,proto3,oneof" json:"ConfigurationToken,omitempty" xml:"ConfigurationToken,omitempty"`
	NextContinuationToken *string           `protobuf:"bytes,9,opt,name=next_continuation_token,json=nextContinuationToken,proto3,oneof" json:"NextContinuationToken,omitempty" xml:"NextContinuationToken,omitempty"`
	IsTruncated           bool              `protobuf:"varint,10,opt,name=is_truncated,json=isTruncated,proto3" json:"IsTruncated" xml:"IsTruncated"`
	MaxKeys               int32             `protobuf:"varint,11,opt,name=max_keys,json=maxKeys,proto3" json:"MaxKeys" xml:"MaxKeys"`
	KeyCount              int32             `protobuf:"varint,12,opt,name=key_count,json=keyCount,proto3" json:"KeyCount" xml:"KeyCount"`
	AllowUnordered        *bool             `protobuf:"varint,13,opt,name=allow_unordered,json=allowUnordered,proto3,oneof" json:"AllowUnordered,omitempty" xml:"AllowUnordered,omitempty"`
}
