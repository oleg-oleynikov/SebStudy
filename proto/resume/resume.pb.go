// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.1
// source: resume/resume.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Skill struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Skill string `protobuf:"bytes,1,opt,name=skill,proto3" json:"skill,omitempty"`
}

func (x *Skill) Reset() {
	*x = Skill{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resume_resume_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Skill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Skill) ProtoMessage() {}

func (x *Skill) ProtoReflect() protoreflect.Message {
	mi := &file_resume_resume_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Skill.ProtoReflect.Descriptor instead.
func (*Skill) Descriptor() ([]byte, []int) {
	return file_resume_resume_proto_rawDescGZIP(), []int{0}
}

func (x *Skill) GetSkill() string {
	if x != nil {
		return x.Skill
	}
	return ""
}

type Direction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Direction string `protobuf:"bytes,1,opt,name=direction,proto3" json:"direction,omitempty"`
}

func (x *Direction) Reset() {
	*x = Direction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resume_resume_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Direction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Direction) ProtoMessage() {}

func (x *Direction) ProtoReflect() protoreflect.Message {
	mi := &file_resume_resume_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Direction.ProtoReflect.Descriptor instead.
func (*Direction) Descriptor() ([]byte, []int) {
	return file_resume_resume_proto_rawDescGZIP(), []int{1}
}

func (x *Direction) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

type ResumeSended struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResumeId      string                 `protobuf:"bytes,1,opt,name=resumeId,proto3" json:"resumeId,omitempty"`
	FirstName     string                 `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	MiddleName    string                 `protobuf:"bytes,3,opt,name=middleName,proto3" json:"middleName,omitempty"`
	LastName      string                 `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
	PhoneNumber   string                 `protobuf:"bytes,5,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Education     string                 `protobuf:"bytes,6,opt,name=education,proto3" json:"education,omitempty"`
	AboutMe       string                 `protobuf:"bytes,7,opt,name=aboutMe,proto3" json:"aboutMe,omitempty"`
	Skills        []*Skill               `protobuf:"bytes,8,rep,name=skills,proto3" json:"skills,omitempty"`
	Photo         []byte                 `protobuf:"bytes,9,opt,name=photo,proto3" json:"photo,omitempty"`
	Directions    []*Direction           `protobuf:"bytes,10,rep,name=directions,proto3" json:"directions,omitempty"`
	AboutProjects string                 `protobuf:"bytes,11,opt,name=aboutProjects,proto3" json:"aboutProjects,omitempty"`
	Portfolio     string                 `protobuf:"bytes,12,opt,name=portfolio,proto3" json:"portfolio,omitempty"`
	StudentGroup  string                 `protobuf:"bytes,13,opt,name=studentGroup,proto3" json:"studentGroup,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *ResumeSended) Reset() {
	*x = ResumeSended{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resume_resume_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResumeSended) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResumeSended) ProtoMessage() {}

func (x *ResumeSended) ProtoReflect() protoreflect.Message {
	mi := &file_resume_resume_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResumeSended.ProtoReflect.Descriptor instead.
func (*ResumeSended) Descriptor() ([]byte, []int) {
	return file_resume_resume_proto_rawDescGZIP(), []int{2}
}

func (x *ResumeSended) GetResumeId() string {
	if x != nil {
		return x.ResumeId
	}
	return ""
}

func (x *ResumeSended) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *ResumeSended) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

func (x *ResumeSended) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *ResumeSended) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *ResumeSended) GetEducation() string {
	if x != nil {
		return x.Education
	}
	return ""
}

func (x *ResumeSended) GetAboutMe() string {
	if x != nil {
		return x.AboutMe
	}
	return ""
}

func (x *ResumeSended) GetSkills() []*Skill {
	if x != nil {
		return x.Skills
	}
	return nil
}

func (x *ResumeSended) GetPhoto() []byte {
	if x != nil {
		return x.Photo
	}
	return nil
}

func (x *ResumeSended) GetDirections() []*Direction {
	if x != nil {
		return x.Directions
	}
	return nil
}

func (x *ResumeSended) GetAboutProjects() string {
	if x != nil {
		return x.AboutProjects
	}
	return ""
}

func (x *ResumeSended) GetPortfolio() string {
	if x != nil {
		return x.Portfolio
	}
	return ""
}

func (x *ResumeSended) GetStudentGroup() string {
	if x != nil {
		return x.StudentGroup
	}
	return ""
}

func (x *ResumeSended) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_resume_resume_proto protoreflect.FileDescriptor

var file_resume_resume_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d,
	0x0a, 0x05, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x22, 0x29, 0x0a,
	0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xf0, 0x03, 0x0a, 0x0c, 0x52, 0x65, 0x73,
	0x75, 0x6d, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x75, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x75, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x4d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x4d, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x6b, 0x69,
	0x6c, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x65, 0x73, 0x75,
	0x6d, 0x65, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x31, 0x0a, 0x0a, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x65, 0x73,
	0x75, 0x6d, 0x65, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x62, 0x6f,
	0x75, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x12, 0x22, 0x0a,
	0x0c, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x09, 0x5a, 0x07, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resume_resume_proto_rawDescOnce sync.Once
	file_resume_resume_proto_rawDescData = file_resume_resume_proto_rawDesc
)

func file_resume_resume_proto_rawDescGZIP() []byte {
	file_resume_resume_proto_rawDescOnce.Do(func() {
		file_resume_resume_proto_rawDescData = protoimpl.X.CompressGZIP(file_resume_resume_proto_rawDescData)
	})
	return file_resume_resume_proto_rawDescData
}

var file_resume_resume_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_resume_resume_proto_goTypes = []interface{}{
	(*Skill)(nil),                 // 0: resume.Skill
	(*Direction)(nil),             // 1: resume.Direction
	(*ResumeSended)(nil),          // 2: resume.ResumeSended
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_resume_resume_proto_depIdxs = []int32{
	0, // 0: resume.ResumeSended.skills:type_name -> resume.Skill
	1, // 1: resume.ResumeSended.directions:type_name -> resume.Direction
	3, // 2: resume.ResumeSended.createdAt:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_resume_resume_proto_init() }
func file_resume_resume_proto_init() {
	if File_resume_resume_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resume_resume_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Skill); i {
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
		file_resume_resume_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Direction); i {
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
		file_resume_resume_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResumeSended); i {
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
			RawDescriptor: file_resume_resume_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resume_resume_proto_goTypes,
		DependencyIndexes: file_resume_resume_proto_depIdxs,
		MessageInfos:      file_resume_resume_proto_msgTypes,
	}.Build()
	File_resume_resume_proto = out.File
	file_resume_resume_proto_rawDesc = nil
	file_resume_resume_proto_goTypes = nil
	file_resume_resume_proto_depIdxs = nil
}
