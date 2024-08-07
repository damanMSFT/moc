// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_cloudagent_galleryimage.proto

package compute

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/microsoft/moc/rpc/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GalleryImageOSType int32

const (
	GalleryImageOSType_UNKNOWN GalleryImageOSType = 0
	GalleryImageOSType_LINUX   GalleryImageOSType = 1
	GalleryImageOSType_WINDOWS GalleryImageOSType = 2
)

var GalleryImageOSType_name = map[int32]string{
	0: "UNKNOWN",
	1: "LINUX",
	2: "WINDOWS",
}

var GalleryImageOSType_value = map[string]int32{
	"UNKNOWN": 0,
	"LINUX":   1,
	"WINDOWS": 2,
}

func (x GalleryImageOSType) String() string {
	return proto.EnumName(GalleryImageOSType_name, int32(x))
}

func (GalleryImageOSType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4b5fad170433da5e, []int{0}
}

type GalleryImageRequest struct {
	GalleryImages        []*GalleryImage  `protobuf:"bytes,1,rep,name=GalleryImages,proto3" json:"GalleryImages,omitempty"`
	OperationType        common.Operation `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GalleryImageRequest) Reset()         { *m = GalleryImageRequest{} }
func (m *GalleryImageRequest) String() string { return proto.CompactTextString(m) }
func (*GalleryImageRequest) ProtoMessage()    {}
func (*GalleryImageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b5fad170433da5e, []int{0}
}

func (m *GalleryImageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GalleryImageRequest.Unmarshal(m, b)
}
func (m *GalleryImageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GalleryImageRequest.Marshal(b, m, deterministic)
}
func (m *GalleryImageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GalleryImageRequest.Merge(m, src)
}
func (m *GalleryImageRequest) XXX_Size() int {
	return xxx_messageInfo_GalleryImageRequest.Size(m)
}
func (m *GalleryImageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GalleryImageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GalleryImageRequest proto.InternalMessageInfo

func (m *GalleryImageRequest) GetGalleryImages() []*GalleryImage {
	if m != nil {
		return m.GalleryImages
	}
	return nil
}

func (m *GalleryImageRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type GalleryImageResponse struct {
	GalleryImages        []*GalleryImage     `protobuf:"bytes,1,rep,name=GalleryImages,proto3" json:"GalleryImages,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GalleryImageResponse) Reset()         { *m = GalleryImageResponse{} }
func (m *GalleryImageResponse) String() string { return proto.CompactTextString(m) }
func (*GalleryImageResponse) ProtoMessage()    {}
func (*GalleryImageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b5fad170433da5e, []int{1}
}

func (m *GalleryImageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GalleryImageResponse.Unmarshal(m, b)
}
func (m *GalleryImageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GalleryImageResponse.Marshal(b, m, deterministic)
}
func (m *GalleryImageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GalleryImageResponse.Merge(m, src)
}
func (m *GalleryImageResponse) XXX_Size() int {
	return xxx_messageInfo_GalleryImageResponse.Size(m)
}
func (m *GalleryImageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GalleryImageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GalleryImageResponse proto.InternalMessageInfo

func (m *GalleryImageResponse) GetGalleryImages() []*GalleryImage {
	if m != nil {
		return m.GalleryImages
	}
	return nil
}

func (m *GalleryImageResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *GalleryImageResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type GalleryImagePrecheckRequest struct {
	GalleryImages        []*GalleryImage `protobuf:"bytes,1,rep,name=GalleryImages,proto3" json:"GalleryImages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GalleryImagePrecheckRequest) Reset()         { *m = GalleryImagePrecheckRequest{} }
func (m *GalleryImagePrecheckRequest) String() string { return proto.CompactTextString(m) }
func (*GalleryImagePrecheckRequest) ProtoMessage()    {}
func (*GalleryImagePrecheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b5fad170433da5e, []int{2}
}

func (m *GalleryImagePrecheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GalleryImagePrecheckRequest.Unmarshal(m, b)
}
func (m *GalleryImagePrecheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GalleryImagePrecheckRequest.Marshal(b, m, deterministic)
}
func (m *GalleryImagePrecheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GalleryImagePrecheckRequest.Merge(m, src)
}
func (m *GalleryImagePrecheckRequest) XXX_Size() int {
	return xxx_messageInfo_GalleryImagePrecheckRequest.Size(m)
}
func (m *GalleryImagePrecheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GalleryImagePrecheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GalleryImagePrecheckRequest proto.InternalMessageInfo

func (m *GalleryImagePrecheckRequest) GetGalleryImages() []*GalleryImage {
	if m != nil {
		return m.GalleryImages
	}
	return nil
}

type GalleryImagePrecheckResponse struct {
	// The precheck result: true if the precheck criteria is passed; otherwise, false
	Result *wrappers.BoolValue `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
	// The error message if the precheck is not passed; otherwise, empty string
	Error                string   `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GalleryImagePrecheckResponse) Reset()         { *m = GalleryImagePrecheckResponse{} }
func (m *GalleryImagePrecheckResponse) String() string { return proto.CompactTextString(m) }
func (*GalleryImagePrecheckResponse) ProtoMessage()    {}
func (*GalleryImagePrecheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b5fad170433da5e, []int{3}
}

func (m *GalleryImagePrecheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GalleryImagePrecheckResponse.Unmarshal(m, b)
}
func (m *GalleryImagePrecheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GalleryImagePrecheckResponse.Marshal(b, m, deterministic)
}
func (m *GalleryImagePrecheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GalleryImagePrecheckResponse.Merge(m, src)
}
func (m *GalleryImagePrecheckResponse) XXX_Size() int {
	return xxx_messageInfo_GalleryImagePrecheckResponse.Size(m)
}
func (m *GalleryImagePrecheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GalleryImagePrecheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GalleryImagePrecheckResponse proto.InternalMessageInfo

func (m *GalleryImagePrecheckResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *GalleryImagePrecheckResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type SFSImageProperties struct {
	CatalogName          string   `protobuf:"bytes,1,opt,name=catalogName,proto3" json:"catalogName,omitempty"`
	Audience             string   `protobuf:"bytes,2,opt,name=audience,proto3" json:"audience,omitempty"`
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	ReleaseName          string   `protobuf:"bytes,4,opt,name=releaseName,proto3" json:"releaseName,omitempty"`
	DestinationDir       string   `protobuf:"bytes,5,opt,name=destinationDir,proto3" json:"destinationDir,omitempty"`
	Parts                uint32   `protobuf:"varint,6,opt,name=parts,proto3" json:"parts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SFSImageProperties) Reset()         { *m = SFSImageProperties{} }
func (m *SFSImageProperties) String() string { return proto.CompactTextString(m) }
func (*SFSImageProperties) ProtoMessage()    {}
func (*SFSImageProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b5fad170433da5e, []int{4}
}

func (m *SFSImageProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SFSImageProperties.Unmarshal(m, b)
}
func (m *SFSImageProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SFSImageProperties.Marshal(b, m, deterministic)
}
func (m *SFSImageProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SFSImageProperties.Merge(m, src)
}
func (m *SFSImageProperties) XXX_Size() int {
	return xxx_messageInfo_SFSImageProperties.Size(m)
}
func (m *SFSImageProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_SFSImageProperties.DiscardUnknown(m)
}

var xxx_messageInfo_SFSImageProperties proto.InternalMessageInfo

func (m *SFSImageProperties) GetCatalogName() string {
	if m != nil {
		return m.CatalogName
	}
	return ""
}

func (m *SFSImageProperties) GetAudience() string {
	if m != nil {
		return m.Audience
	}
	return ""
}

func (m *SFSImageProperties) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *SFSImageProperties) GetReleaseName() string {
	if m != nil {
		return m.ReleaseName
	}
	return ""
}

func (m *SFSImageProperties) GetDestinationDir() string {
	if m != nil {
		return m.DestinationDir
	}
	return ""
}

func (m *SFSImageProperties) GetParts() uint32 {
	if m != nil {
		return m.Parts
	}
	return 0
}

type LocalImageProperties struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalImageProperties) Reset()         { *m = LocalImageProperties{} }
func (m *LocalImageProperties) String() string { return proto.CompactTextString(m) }
func (*LocalImageProperties) ProtoMessage()    {}
func (*LocalImageProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b5fad170433da5e, []int{5}
}

func (m *LocalImageProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalImageProperties.Unmarshal(m, b)
}
func (m *LocalImageProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalImageProperties.Marshal(b, m, deterministic)
}
func (m *LocalImageProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalImageProperties.Merge(m, src)
}
func (m *LocalImageProperties) XXX_Size() int {
	return xxx_messageInfo_LocalImageProperties.Size(m)
}
func (m *LocalImageProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalImageProperties.DiscardUnknown(m)
}

var xxx_messageInfo_LocalImageProperties proto.InternalMessageInfo

func (m *LocalImageProperties) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type CloneImageProperties struct {
	CloneSource          string   `protobuf:"bytes,1,opt,name=cloneSource,proto3" json:"cloneSource,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloneImageProperties) Reset()         { *m = CloneImageProperties{} }
func (m *CloneImageProperties) String() string { return proto.CompactTextString(m) }
func (*CloneImageProperties) ProtoMessage()    {}
func (*CloneImageProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b5fad170433da5e, []int{6}
}

func (m *CloneImageProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloneImageProperties.Unmarshal(m, b)
}
func (m *CloneImageProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloneImageProperties.Marshal(b, m, deterministic)
}
func (m *CloneImageProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloneImageProperties.Merge(m, src)
}
func (m *CloneImageProperties) XXX_Size() int {
	return xxx_messageInfo_CloneImageProperties.Size(m)
}
func (m *CloneImageProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_CloneImageProperties.DiscardUnknown(m)
}

var xxx_messageInfo_CloneImageProperties proto.InternalMessageInfo

func (m *CloneImageProperties) GetCloneSource() string {
	if m != nil {
		return m.CloneSource
	}
	return ""
}

type HttpImageProperties struct {
	HttpURL              string   `protobuf:"bytes,1,opt,name=httpURL,proto3" json:"httpURL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HttpImageProperties) Reset()         { *m = HttpImageProperties{} }
func (m *HttpImageProperties) String() string { return proto.CompactTextString(m) }
func (*HttpImageProperties) ProtoMessage()    {}
func (*HttpImageProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b5fad170433da5e, []int{7}
}

func (m *HttpImageProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpImageProperties.Unmarshal(m, b)
}
func (m *HttpImageProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpImageProperties.Marshal(b, m, deterministic)
}
func (m *HttpImageProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpImageProperties.Merge(m, src)
}
func (m *HttpImageProperties) XXX_Size() int {
	return xxx_messageInfo_HttpImageProperties.Size(m)
}
func (m *HttpImageProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpImageProperties.DiscardUnknown(m)
}

var xxx_messageInfo_HttpImageProperties proto.InternalMessageInfo

func (m *HttpImageProperties) GetHttpURL() string {
	if m != nil {
		return m.HttpURL
	}
	return ""
}

type AzureGalleryImageProperties struct {
	SasURI               string   `protobuf:"bytes,1,opt,name=sasURI,proto3" json:"sasURI,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AzureGalleryImageProperties) Reset()         { *m = AzureGalleryImageProperties{} }
func (m *AzureGalleryImageProperties) String() string { return proto.CompactTextString(m) }
func (*AzureGalleryImageProperties) ProtoMessage()    {}
func (*AzureGalleryImageProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b5fad170433da5e, []int{8}
}

func (m *AzureGalleryImageProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AzureGalleryImageProperties.Unmarshal(m, b)
}
func (m *AzureGalleryImageProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AzureGalleryImageProperties.Marshal(b, m, deterministic)
}
func (m *AzureGalleryImageProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AzureGalleryImageProperties.Merge(m, src)
}
func (m *AzureGalleryImageProperties) XXX_Size() int {
	return xxx_messageInfo_AzureGalleryImageProperties.Size(m)
}
func (m *AzureGalleryImageProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_AzureGalleryImageProperties.DiscardUnknown(m)
}

var xxx_messageInfo_AzureGalleryImageProperties proto.InternalMessageInfo

func (m *AzureGalleryImageProperties) GetSasURI() string {
	if m != nil {
		return m.SasURI
	}
	return ""
}

func (m *AzureGalleryImageProperties) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type GalleryImage struct {
	Name        string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id          string             `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	ImageOSType GalleryImageOSType `protobuf:"varint,3,opt,name=imageOSType,proto3,enum=moc.cloudagent.compute.GalleryImageOSType" json:"imageOSType,omitempty"`
	// Path of the image on the cloud
	Path          string         `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Status        *common.Status `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	LocationName  string         `protobuf:"bytes,6,opt,name=locationName,proto3" json:"locationName,omitempty"`
	Publisher     string         `protobuf:"bytes,7,opt,name=publisher,proto3" json:"publisher,omitempty"`
	Sku           string         `protobuf:"bytes,8,opt,name=sku,proto3" json:"sku,omitempty"`
	Offer         string         `protobuf:"bytes,9,opt,name=offer,proto3" json:"offer,omitempty"`
	ContainerName string         `protobuf:"bytes,10,opt,name=containerName,proto3" json:"containerName,omitempty"`
	// Source of the GalleryImage from where we can copy the image from.
	SourcePath           string                     `protobuf:"bytes,11,opt,name=sourcePath,proto3" json:"sourcePath,omitempty"`
	Tags                 *common.Tags               `protobuf:"bytes,12,opt,name=tags,proto3" json:"tags,omitempty"`
	SourceType           common.ImageSource         `protobuf:"varint,13,opt,name=sourceType,proto3,enum=moc.ImageSource" json:"sourceType,omitempty"`
	HyperVGeneration     common.HyperVGeneration    `protobuf:"varint,14,opt,name=hyperVGeneration,proto3,enum=moc.HyperVGeneration" json:"hyperVGeneration,omitempty"`
	CloudInitDataSource  common.CloudInitDataSource `protobuf:"varint,15,opt,name=cloudInitDataSource,proto3,enum=moc.CloudInitDataSource" json:"cloudInitDataSource,omitempty"`
	SourceVM             string                     `protobuf:"bytes,16,opt,name=sourceVM,proto3" json:"sourceVM,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *GalleryImage) Reset()         { *m = GalleryImage{} }
func (m *GalleryImage) String() string { return proto.CompactTextString(m) }
func (*GalleryImage) ProtoMessage()    {}
func (*GalleryImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b5fad170433da5e, []int{9}
}

func (m *GalleryImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GalleryImage.Unmarshal(m, b)
}
func (m *GalleryImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GalleryImage.Marshal(b, m, deterministic)
}
func (m *GalleryImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GalleryImage.Merge(m, src)
}
func (m *GalleryImage) XXX_Size() int {
	return xxx_messageInfo_GalleryImage.Size(m)
}
func (m *GalleryImage) XXX_DiscardUnknown() {
	xxx_messageInfo_GalleryImage.DiscardUnknown(m)
}

var xxx_messageInfo_GalleryImage proto.InternalMessageInfo

func (m *GalleryImage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GalleryImage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GalleryImage) GetImageOSType() GalleryImageOSType {
	if m != nil {
		return m.ImageOSType
	}
	return GalleryImageOSType_UNKNOWN
}

func (m *GalleryImage) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *GalleryImage) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GalleryImage) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func (m *GalleryImage) GetPublisher() string {
	if m != nil {
		return m.Publisher
	}
	return ""
}

func (m *GalleryImage) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func (m *GalleryImage) GetOffer() string {
	if m != nil {
		return m.Offer
	}
	return ""
}

func (m *GalleryImage) GetContainerName() string {
	if m != nil {
		return m.ContainerName
	}
	return ""
}

func (m *GalleryImage) GetSourcePath() string {
	if m != nil {
		return m.SourcePath
	}
	return ""
}

func (m *GalleryImage) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *GalleryImage) GetSourceType() common.ImageSource {
	if m != nil {
		return m.SourceType
	}
	return common.ImageSource_LOCAL_SOURCE
}

func (m *GalleryImage) GetHyperVGeneration() common.HyperVGeneration {
	if m != nil {
		return m.HyperVGeneration
	}
	return common.HyperVGeneration_HyperVGenerationV2
}

func (m *GalleryImage) GetCloudInitDataSource() common.CloudInitDataSource {
	if m != nil {
		return m.CloudInitDataSource
	}
	return common.CloudInitDataSource_NoCloud
}

func (m *GalleryImage) GetSourceVM() string {
	if m != nil {
		return m.SourceVM
	}
	return ""
}

func init() {
	proto.RegisterEnum("moc.cloudagent.compute.GalleryImageOSType", GalleryImageOSType_name, GalleryImageOSType_value)
	proto.RegisterType((*GalleryImageRequest)(nil), "moc.cloudagent.compute.GalleryImageRequest")
	proto.RegisterType((*GalleryImageResponse)(nil), "moc.cloudagent.compute.GalleryImageResponse")
	proto.RegisterType((*GalleryImagePrecheckRequest)(nil), "moc.cloudagent.compute.GalleryImagePrecheckRequest")
	proto.RegisterType((*GalleryImagePrecheckResponse)(nil), "moc.cloudagent.compute.GalleryImagePrecheckResponse")
	proto.RegisterType((*SFSImageProperties)(nil), "moc.cloudagent.compute.SFSImageProperties")
	proto.RegisterType((*LocalImageProperties)(nil), "moc.cloudagent.compute.LocalImageProperties")
	proto.RegisterType((*CloneImageProperties)(nil), "moc.cloudagent.compute.CloneImageProperties")
	proto.RegisterType((*HttpImageProperties)(nil), "moc.cloudagent.compute.HttpImageProperties")
	proto.RegisterType((*AzureGalleryImageProperties)(nil), "moc.cloudagent.compute.AzureGalleryImageProperties")
	proto.RegisterType((*GalleryImage)(nil), "moc.cloudagent.compute.GalleryImage")
}

func init() { proto.RegisterFile("moc_cloudagent_galleryimage.proto", fileDescriptor_4b5fad170433da5e) }

var fileDescriptor_4b5fad170433da5e = []byte{
	// 859 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xd1, 0x72, 0xdb, 0x44,
	0x14, 0xad, 0x1c, 0xc7, 0x89, 0xaf, 0x63, 0x63, 0x36, 0x01, 0x34, 0x6e, 0x60, 0x8c, 0xc8, 0x30,
	0x99, 0xc2, 0x58, 0xc5, 0xed, 0x03, 0x3c, 0x26, 0x0d, 0xb4, 0x2e, 0xc1, 0xe9, 0xac, 0xeb, 0x84,
	0xe1, 0xa5, 0xb3, 0x91, 0x6f, 0x64, 0x4d, 0x24, 0xad, 0xd8, 0x5d, 0xb5, 0x13, 0xbe, 0x80, 0x3f,
	0xe0, 0x27, 0x78, 0xe5, 0x2f, 0xf8, 0x19, 0x86, 0x1f, 0x60, 0x74, 0x25, 0xc7, 0x92, 0x93, 0x87,
	0x64, 0x86, 0x3e, 0xd9, 0x7b, 0xf6, 0x9e, 0x73, 0xef, 0x3d, 0x7b, 0x75, 0xe1, 0xf3, 0x48, 0x7a,
	0x6f, 0xbc, 0x50, 0xa6, 0x33, 0xe1, 0x63, 0x6c, 0xde, 0xf8, 0x22, 0x0c, 0x51, 0x5d, 0x05, 0x91,
	0xf0, 0x71, 0x90, 0x28, 0x69, 0x24, 0xfb, 0x38, 0x92, 0xde, 0x60, 0x19, 0x32, 0xf0, 0x64, 0x94,
	0xa4, 0x06, 0x7b, 0x9f, 0xf9, 0x52, 0xfa, 0x21, 0xba, 0x14, 0x75, 0x9e, 0x5e, 0xb8, 0xef, 0x94,
	0x48, 0x12, 0x54, 0x3a, 0xe7, 0xf5, 0x3e, 0x21, 0x69, 0x19, 0x45, 0x32, 0x2e, 0x7e, 0xf2, 0x0b,
	0xe7, 0x0f, 0x0b, 0xb6, 0x9f, 0xe7, 0x79, 0x46, 0x59, 0x1e, 0x8e, 0xbf, 0xa6, 0xa8, 0x0d, 0x7b,
	0x09, 0xed, 0x32, 0xac, 0x6d, 0xab, 0xbf, 0xb6, 0xdf, 0x1a, 0xee, 0x0d, 0x6e, 0x2f, 0x60, 0x50,
	0xd1, 0xa8, 0x52, 0xd9, 0x53, 0x68, 0x9f, 0x24, 0xa8, 0x84, 0x09, 0x64, 0xfc, 0xfa, 0x2a, 0x41,
	0xbb, 0xd6, 0xb7, 0xf6, 0x3b, 0xc3, 0x0e, 0x69, 0x5d, 0xdf, 0xf0, 0x6a, 0x90, 0xf3, 0xa7, 0x05,
	0x3b, 0xd5, 0xca, 0x74, 0x22, 0x63, 0x8d, 0xff, 0x6b, 0x69, 0x43, 0x68, 0x70, 0xd4, 0x69, 0x68,
	0xa8, 0xa6, 0xd6, 0xb0, 0x37, 0xc8, 0x8d, 0x1c, 0x2c, 0x8c, 0x1c, 0x1c, 0x4a, 0x19, 0x9e, 0x8a,
	0x30, 0x45, 0x5e, 0x44, 0xb2, 0x1d, 0x58, 0xff, 0x5e, 0x29, 0xa9, 0xec, 0xb5, 0xbe, 0xb5, 0xdf,
	0xe4, 0xf9, 0xc1, 0x09, 0xe0, 0x61, 0x59, 0xfa, 0x95, 0x42, 0x6f, 0x8e, 0xde, 0xe5, 0x7b, 0xf0,
	0xd3, 0x99, 0xc3, 0xee, 0xed, 0xa9, 0x0a, 0x83, 0x96, 0x4d, 0x59, 0xf7, 0x6f, 0xaa, 0x56, 0x6e,
	0xea, 0x6f, 0x0b, 0xd8, 0xe4, 0x87, 0x49, 0x91, 0x46, 0x26, 0xa8, 0x4c, 0x80, 0x9a, 0xf5, 0xa1,
	0xe5, 0x09, 0x23, 0x42, 0xe9, 0x8f, 0x45, 0x84, 0x94, 0xa5, 0xc9, 0xcb, 0x10, 0xeb, 0xc1, 0xa6,
	0x48, 0x67, 0x01, 0xc6, 0x1e, 0x16, 0x8a, 0xd7, 0x67, 0x66, 0xc3, 0xc6, 0x5b, 0x54, 0x3a, 0x90,
	0x71, 0xe1, 0xe0, 0xe2, 0x98, 0xe9, 0x2a, 0x0c, 0x51, 0x68, 0x24, 0xdd, 0x7a, 0xae, 0x5b, 0x82,
	0xd8, 0x97, 0xd0, 0x99, 0xa1, 0x36, 0x41, 0x4c, 0x73, 0x72, 0x14, 0x28, 0x7b, 0x9d, 0x82, 0x56,
	0xd0, 0xac, 0x9d, 0x44, 0x28, 0xa3, 0xed, 0x46, 0xdf, 0xda, 0x6f, 0xf3, 0xfc, 0xe0, 0x3c, 0x86,
	0x9d, 0x63, 0xe9, 0x89, 0x70, 0xb5, 0x1f, 0x1b, 0xea, 0x89, 0x30, 0xf3, 0xbc, 0x91, 0xc3, 0xfa,
	0xef, 0x7f, 0xd9, 0x16, 0x27, 0xc4, 0xf9, 0x16, 0x76, 0x9e, 0x85, 0x32, 0xc6, 0xdb, 0x1c, 0xc8,
	0xf0, 0x89, 0x4c, 0x95, 0xb7, 0x74, 0x60, 0x09, 0x39, 0x2e, 0x6c, 0xbf, 0x30, 0x26, 0xb9, 0x99,
	0x6a, 0x63, 0x6e, 0x4c, 0x32, 0xe5, 0xc7, 0x05, 0x69, 0x71, 0x74, 0xa6, 0xf0, 0xf0, 0xe0, 0xb7,
	0x54, 0x61, 0xf5, 0x69, 0xaf, 0x89, 0xbb, 0xd0, 0xd0, 0x42, 0x4f, 0xf9, 0xa8, 0x52, 0x65, 0x81,
	0x95, 0x3d, 0xad, 0x55, 0x3c, 0x75, 0xfe, 0xa9, 0xc3, 0x56, 0x59, 0x92, 0x31, 0xa8, 0xc7, 0xcb,
	0x57, 0xa3, 0xff, 0xac, 0x03, 0xb5, 0x60, 0x56, 0x30, 0x6b, 0xc1, 0x8c, 0x1d, 0x43, 0x8b, 0xb6,
	0xce, 0xc9, 0x84, 0xbe, 0xd7, 0x35, 0xfa, 0x5e, 0x1f, 0xdd, 0x65, 0x56, 0x73, 0x06, 0x2f, 0xd3,
	0xaf, 0xed, 0xad, 0xaf, 0xda, 0xcb, 0xbe, 0x80, 0x86, 0x36, 0xc2, 0xa4, 0x9a, 0x9e, 0xb1, 0x35,
	0x6c, 0x51, 0x8a, 0x09, 0x41, 0xbc, 0xb8, 0x62, 0x0e, 0x6c, 0x85, 0xd2, 0xa3, 0xa7, 0xa5, 0xb1,
	0x68, 0x50, 0x99, 0x15, 0x8c, 0xed, 0x42, 0x33, 0x49, 0xcf, 0xc3, 0x40, 0xcf, 0x51, 0xd9, 0x1b,
	0x14, 0xb0, 0x04, 0x58, 0x17, 0xd6, 0xf4, 0x65, 0x6a, 0x6f, 0x12, 0x9e, 0xfd, 0xcd, 0xe6, 0x43,
	0x5e, 0x5c, 0xa0, 0xb2, 0x9b, 0xf9, 0xb8, 0xd3, 0x81, 0xed, 0x41, 0xdb, 0x93, 0xb1, 0x11, 0x41,
	0x8c, 0x8a, 0x52, 0x01, 0xdd, 0x56, 0x41, 0xb6, 0x07, 0xa0, 0xe9, 0x8d, 0x5f, 0x65, 0x4d, 0xb5,
	0x4a, 0x4d, 0x95, 0x70, 0xf6, 0x29, 0xd4, 0x8d, 0xf0, 0xb5, 0xbd, 0x45, 0x8d, 0x35, 0xa9, 0xb1,
	0xd7, 0xc2, 0xd7, 0x9c, 0x60, 0xf6, 0x78, 0x21, 0x42, 0x06, 0xb7, 0xc9, 0xe0, 0x2e, 0x05, 0x91,
	0x8d, 0xf9, 0x10, 0xf1, 0x52, 0x0c, 0x3b, 0x80, 0xee, 0xfc, 0x2a, 0x41, 0x75, 0xfa, 0x1c, 0xe3,
	0x62, 0x4f, 0xda, 0x1d, 0xe2, 0x7d, 0x44, 0xbc, 0x17, 0x2b, 0x97, 0xfc, 0x46, 0x38, 0x7b, 0x09,
	0xdb, 0xf4, 0x7c, 0xa3, 0x38, 0x30, 0x47, 0xc2, 0x88, 0x62, 0x7a, 0x3f, 0x20, 0x15, 0x9b, 0x54,
	0x9e, 0xdd, 0xbc, 0xe7, 0xb7, 0x91, 0xb2, 0x2f, 0x3c, 0x2f, 0xee, 0xf4, 0x27, 0xbb, 0x9b, 0x7f,
	0xe1, 0x8b, 0xf3, 0xa3, 0xef, 0x80, 0xdd, 0x9c, 0x09, 0xd6, 0x82, 0x8d, 0xe9, 0xf8, 0xc7, 0xf1,
	0xc9, 0xd9, 0xb8, 0xfb, 0x80, 0x35, 0x61, 0xfd, 0x78, 0x34, 0x9e, 0xfe, 0xdc, 0xb5, 0x32, 0xfc,
	0x6c, 0x34, 0x3e, 0x3a, 0x39, 0x9b, 0x74, 0x6b, 0xc3, 0x7f, 0x2d, 0xf8, 0xb0, 0xcc, 0x3d, 0xc8,
	0x26, 0x8d, 0x21, 0x34, 0x46, 0xf1, 0x5b, 0x79, 0x89, 0xec, 0xab, 0x3b, 0x2d, 0xcc, 0x7c, 0xe9,
	0xf6, 0xbe, 0xbe, 0x5b, 0x70, 0xbe, 0x36, 0x9d, 0x07, 0xec, 0x1d, 0x6c, 0x2e, 0x96, 0x29, 0x7b,
	0x72, 0x17, 0xee, 0xca, 0x96, 0xef, 0x3d, 0xbd, 0x1f, 0x69, 0x91, 0xf8, 0xf0, 0x9b, 0x5f, 0x5c,
	0x3f, 0x30, 0xf3, 0xf4, 0x3c, 0x23, 0xb8, 0x51, 0xe0, 0x29, 0xa9, 0xe5, 0x85, 0x71, 0x23, 0xe9,
	0xb9, 0x2a, 0xf1, 0xdc, 0xa5, 0xa2, 0x5b, 0x28, 0x9e, 0x37, 0x68, 0x99, 0x3f, 0xf9, 0x2f, 0x00,
	0x00, 0xff, 0xff, 0x5a, 0xb9, 0x52, 0x5e, 0x35, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GalleryImageAgentClient is the client API for GalleryImageAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GalleryImageAgentClient interface {
	Invoke(ctx context.Context, in *GalleryImageRequest, opts ...grpc.CallOption) (*GalleryImageResponse, error)
	// Prechecks whether the system is able to create specified gallery images (but does not actually create them).
	Precheck(ctx context.Context, in *GalleryImagePrecheckRequest, opts ...grpc.CallOption) (*GalleryImagePrecheckResponse, error)
}

type galleryImageAgentClient struct {
	cc *grpc.ClientConn
}

func NewGalleryImageAgentClient(cc *grpc.ClientConn) GalleryImageAgentClient {
	return &galleryImageAgentClient{cc}
}

func (c *galleryImageAgentClient) Invoke(ctx context.Context, in *GalleryImageRequest, opts ...grpc.CallOption) (*GalleryImageResponse, error) {
	out := new(GalleryImageResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.compute.GalleryImageAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *galleryImageAgentClient) Precheck(ctx context.Context, in *GalleryImagePrecheckRequest, opts ...grpc.CallOption) (*GalleryImagePrecheckResponse, error) {
	out := new(GalleryImagePrecheckResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.compute.GalleryImageAgent/Precheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GalleryImageAgentServer is the server API for GalleryImageAgent service.
type GalleryImageAgentServer interface {
	Invoke(context.Context, *GalleryImageRequest) (*GalleryImageResponse, error)
	// Prechecks whether the system is able to create specified gallery images (but does not actually create them).
	Precheck(context.Context, *GalleryImagePrecheckRequest) (*GalleryImagePrecheckResponse, error)
}

// UnimplementedGalleryImageAgentServer can be embedded to have forward compatible implementations.
type UnimplementedGalleryImageAgentServer struct {
}

func (*UnimplementedGalleryImageAgentServer) Invoke(ctx context.Context, req *GalleryImageRequest) (*GalleryImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}
func (*UnimplementedGalleryImageAgentServer) Precheck(ctx context.Context, req *GalleryImagePrecheckRequest) (*GalleryImagePrecheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Precheck not implemented")
}

func RegisterGalleryImageAgentServer(s *grpc.Server, srv GalleryImageAgentServer) {
	s.RegisterService(&_GalleryImageAgent_serviceDesc, srv)
}

func _GalleryImageAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GalleryImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GalleryImageAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.compute.GalleryImageAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GalleryImageAgentServer).Invoke(ctx, req.(*GalleryImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GalleryImageAgent_Precheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GalleryImagePrecheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GalleryImageAgentServer).Precheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.compute.GalleryImageAgent/Precheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GalleryImageAgentServer).Precheck(ctx, req.(*GalleryImagePrecheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GalleryImageAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.compute.GalleryImageAgent",
	HandlerType: (*GalleryImageAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _GalleryImageAgent_Invoke_Handler,
		},
		{
			MethodName: "Precheck",
			Handler:    _GalleryImageAgent_Precheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moc_cloudagent_galleryimage.proto",
}
