// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: route/route.proto

package route

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//AddRouteRequest добавить роут
type RemoveRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HcDestIP    string `protobuf:"bytes,1,opt,name=hcDestIP,proto3" json:"hcDestIP,omitempty"`
	HcTunDestIP string `protobuf:"bytes,2,opt,name=hcTunDestIP,proto3" json:"hcTunDestIP,omitempty"`
}

func (x *RemoveRouteRequest) Reset() {
	*x = RemoveRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRouteRequest) ProtoMessage() {}

func (x *RemoveRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_route_route_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRouteRequest.ProtoReflect.Descriptor instead.
func (*RemoveRouteRequest) Descriptor() ([]byte, []int) {
	return file_route_route_proto_rawDescGZIP(), []int{0}
}

func (x *RemoveRouteRequest) GetHcDestIP() string {
	if x != nil {
		return x.HcDestIP
	}
	return ""
}

func (x *RemoveRouteRequest) GetHcTunDestIP() string {
	if x != nil {
		return x.HcTunDestIP
	}
	return ""
}

//AddRouteRequest добавить роут
type AddRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HcDestIP    string `protobuf:"bytes,1,opt,name=hcDestIP,proto3" json:"hcDestIP,omitempty"`
	HcTunDestIP string `protobuf:"bytes,2,opt,name=hcTunDestIP,proto3" json:"hcTunDestIP,omitempty"`
}

func (x *AddRouteRequest) Reset() {
	*x = AddRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_route_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRouteRequest) ProtoMessage() {}

func (x *AddRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_route_route_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRouteRequest.ProtoReflect.Descriptor instead.
func (*AddRouteRequest) Descriptor() ([]byte, []int) {
	return file_route_route_proto_rawDescGZIP(), []int{1}
}

func (x *AddRouteRequest) GetHcDestIP() string {
	if x != nil {
		return x.HcDestIP
	}
	return ""
}

func (x *AddRouteRequest) GetHcTunDestIP() string {
	if x != nil {
		return x.HcTunDestIP
	}
	return ""
}

//Route IP route
type Route struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HcDestIP    string `protobuf:"bytes,1,opt,name=hcDestIP,proto3" json:"hcDestIP,omitempty"`
	HcTunDestIP string `protobuf:"bytes,2,opt,name=hcTunDestIP,proto3" json:"hcTunDestIP,omitempty"`
	//dev net-link device
	Dev string `protobuf:"bytes,3,opt,name=dev,proto3" json:"dev,omitempty"`
	//table route table
	Table string `protobuf:"bytes,4,opt,name=table,proto3" json:"table,omitempty"`
}

func (x *Route) Reset() {
	*x = Route{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_route_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route) ProtoMessage() {}

func (x *Route) ProtoReflect() protoreflect.Message {
	mi := &file_route_route_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route.ProtoReflect.Descriptor instead.
func (*Route) Descriptor() ([]byte, []int) {
	return file_route_route_proto_rawDescGZIP(), []int{2}
}

func (x *Route) GetHcDestIP() string {
	if x != nil {
		return x.HcDestIP
	}
	return ""
}

func (x *Route) GetHcTunDestIP() string {
	if x != nil {
		return x.HcTunDestIP
	}
	return ""
}

func (x *Route) GetDev() string {
	if x != nil {
		return x.Dev
	}
	return ""
}

func (x *Route) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

//GetStateResponse выдаем все роуты
type GetStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//routes список роутов
	Routes []*Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
}

func (x *GetStateResponse) Reset() {
	*x = GetStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_route_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStateResponse) ProtoMessage() {}

func (x *GetStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_route_route_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStateResponse.ProtoReflect.Descriptor instead.
func (*GetStateResponse) Descriptor() ([]byte, []int) {
	return file_route_route_proto_rawDescGZIP(), []int{3}
}

func (x *GetStateResponse) GetRoutes() []*Route {
	if x != nil {
		return x.Routes
	}
	return nil
}

var File_route_route_proto protoreflect.FileDescriptor

var file_route_route_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x68,
	0x63, 0x44, 0x65, 0x73, 0x74, 0x49, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68,
	0x63, 0x44, 0x65, 0x73, 0x74, 0x49, 0x50, 0x12, 0x20, 0x0a, 0x0b, 0x68, 0x63, 0x54, 0x75, 0x6e,
	0x44, 0x65, 0x73, 0x74, 0x49, 0x50, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x63,
	0x54, 0x75, 0x6e, 0x44, 0x65, 0x73, 0x74, 0x49, 0x50, 0x22, 0x4f, 0x0a, 0x0f, 0x41, 0x64, 0x64,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x68, 0x63, 0x44, 0x65, 0x73, 0x74, 0x49, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x68, 0x63, 0x44, 0x65, 0x73, 0x74, 0x49, 0x50, 0x12, 0x20, 0x0a, 0x0b, 0x68, 0x63, 0x54, 0x75,
	0x6e, 0x44, 0x65, 0x73, 0x74, 0x49, 0x50, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68,
	0x63, 0x54, 0x75, 0x6e, 0x44, 0x65, 0x73, 0x74, 0x49, 0x50, 0x22, 0x6d, 0x0a, 0x05, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x63, 0x44, 0x65, 0x73, 0x74, 0x49, 0x50, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x63, 0x44, 0x65, 0x73, 0x74, 0x49, 0x50, 0x12,
	0x20, 0x0a, 0x0b, 0x68, 0x63, 0x54, 0x75, 0x6e, 0x44, 0x65, 0x73, 0x74, 0x49, 0x50, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x63, 0x54, 0x75, 0x6e, 0x44, 0x65, 0x73, 0x74, 0x49,
	0x50, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x65, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x64, 0x65, 0x76, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x38, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a,
	0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x06, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x32, 0x99, 0x02, 0x0a, 0x0c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x12, 0x16, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x2f, 0x61, 0x64, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x5d, 0x0a, 0x0b, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2f,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x54, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f,
	0x2f, 0x76, 0x32, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x42,
	0xa7, 0x01, 0x5a, 0x06, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x92, 0x41, 0x9b, 0x01, 0x12, 0x72,
	0x0a, 0x10, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x20, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x20, 0x41,
	0x50, 0x49, 0x22, 0x59, 0x0a, 0x01, 0x45, 0x12, 0x54, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f,
	0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x62, 0x75, 0x6c, 0x6c, 0x67, 0x61, 0x72, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x32, 0x30, 0x32, 0x30, 0x2f, 0x30, 0x37, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2d, 0x6f, 0x66, 0x2d, 0x73, 0x77, 0x61, 0x67,
	0x67, 0x65, 0x72, 0x2d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x74, 0x6f, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x32, 0x03, 0x32,
	0x2e, 0x30, 0x2a, 0x01, 0x01, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_route_route_proto_rawDescOnce sync.Once
	file_route_route_proto_rawDescData = file_route_route_proto_rawDesc
)

func file_route_route_proto_rawDescGZIP() []byte {
	file_route_route_proto_rawDescOnce.Do(func() {
		file_route_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_route_route_proto_rawDescData)
	})
	return file_route_route_proto_rawDescData
}

var file_route_route_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_route_route_proto_goTypes = []interface{}{
	(*RemoveRouteRequest)(nil), // 0: route.RemoveRouteRequest
	(*AddRouteRequest)(nil),    // 1: route.AddRouteRequest
	(*Route)(nil),              // 2: route.Route
	(*GetStateResponse)(nil),   // 3: route.GetStateResponse
	(*emptypb.Empty)(nil),      // 4: google.protobuf.Empty
}
var file_route_route_proto_depIdxs = []int32{
	2, // 0: route.GetStateResponse.routes:type_name -> route.Route
	1, // 1: route.RouteService.AddRoute:input_type -> route.AddRouteRequest
	0, // 2: route.RouteService.RemoveRoute:input_type -> route.RemoveRouteRequest
	4, // 3: route.RouteService.GetState:input_type -> google.protobuf.Empty
	4, // 4: route.RouteService.AddRoute:output_type -> google.protobuf.Empty
	4, // 5: route.RouteService.RemoveRoute:output_type -> google.protobuf.Empty
	3, // 6: route.RouteService.GetState:output_type -> route.GetStateResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_route_route_proto_init() }
func file_route_route_proto_init() {
	if File_route_route_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_route_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRouteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_route_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRouteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_route_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Route); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_route_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_route_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_route_route_proto_goTypes,
		DependencyIndexes: file_route_route_proto_depIdxs,
		MessageInfos:      file_route_route_proto_msgTypes,
	}.Build()
	File_route_route_proto = out.File
	file_route_route_proto_rawDesc = nil
	file_route_route_proto_goTypes = nil
	file_route_route_proto_depIdxs = nil
}
