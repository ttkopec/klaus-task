// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: score_service.proto

package gen

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

// START: Request types
type TimeRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *TimeRange) Reset() {
	*x = TimeRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_score_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeRange) ProtoMessage() {}

func (x *TimeRange) ProtoReflect() protoreflect.Message {
	mi := &file_score_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeRange.ProtoReflect.Descriptor instead.
func (*TimeRange) Descriptor() ([]byte, []int) {
	return file_score_service_proto_rawDescGZIP(), []int{0}
}

func (x *TimeRange) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *TimeRange) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

type AggregatedCategoryScoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Period *TimeRange `protobuf:"bytes,1,opt,name=period,proto3" json:"period,omitempty"`
}

func (x *AggregatedCategoryScoreRequest) Reset() {
	*x = AggregatedCategoryScoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_score_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggregatedCategoryScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregatedCategoryScoreRequest) ProtoMessage() {}

func (x *AggregatedCategoryScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_score_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregatedCategoryScoreRequest.ProtoReflect.Descriptor instead.
func (*AggregatedCategoryScoreRequest) Descriptor() ([]byte, []int) {
	return file_score_service_proto_rawDescGZIP(), []int{1}
}

func (x *AggregatedCategoryScoreRequest) GetPeriod() *TimeRange {
	if x != nil {
		return x.Period
	}
	return nil
}

type ScoresByTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Period *TimeRange `protobuf:"bytes,1,opt,name=period,proto3" json:"period,omitempty"`
}

func (x *ScoresByTicketRequest) Reset() {
	*x = ScoresByTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_score_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScoresByTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScoresByTicketRequest) ProtoMessage() {}

func (x *ScoresByTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_score_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScoresByTicketRequest.ProtoReflect.Descriptor instead.
func (*ScoresByTicketRequest) Descriptor() ([]byte, []int) {
	return file_score_service_proto_rawDescGZIP(), []int{2}
}

func (x *ScoresByTicketRequest) GetPeriod() *TimeRange {
	if x != nil {
		return x.Period
	}
	return nil
}

type OverallQualityScoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Period *TimeRange `protobuf:"bytes,1,opt,name=period,proto3" json:"period,omitempty"`
}

func (x *OverallQualityScoreRequest) Reset() {
	*x = OverallQualityScoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_score_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OverallQualityScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverallQualityScoreRequest) ProtoMessage() {}

func (x *OverallQualityScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_score_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverallQualityScoreRequest.ProtoReflect.Descriptor instead.
func (*OverallQualityScoreRequest) Descriptor() ([]byte, []int) {
	return file_score_service_proto_rawDescGZIP(), []int{3}
}

func (x *OverallQualityScoreRequest) GetPeriod() *TimeRange {
	if x != nil {
		return x.Period
	}
	return nil
}

type PeriodOverPeriodScoreChangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentPeriod  *TimeRange `protobuf:"bytes,1,opt,name=currentPeriod,proto3" json:"currentPeriod,omitempty"`
	PreviousPeriod *TimeRange `protobuf:"bytes,2,opt,name=previousPeriod,proto3" json:"previousPeriod,omitempty"`
}

func (x *PeriodOverPeriodScoreChangeRequest) Reset() {
	*x = PeriodOverPeriodScoreChangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_score_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeriodOverPeriodScoreChangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeriodOverPeriodScoreChangeRequest) ProtoMessage() {}

func (x *PeriodOverPeriodScoreChangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_score_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeriodOverPeriodScoreChangeRequest.ProtoReflect.Descriptor instead.
func (*PeriodOverPeriodScoreChangeRequest) Descriptor() ([]byte, []int) {
	return file_score_service_proto_rawDescGZIP(), []int{4}
}

func (x *PeriodOverPeriodScoreChangeRequest) GetCurrentPeriod() *TimeRange {
	if x != nil {
		return x.CurrentPeriod
	}
	return nil
}

func (x *PeriodOverPeriodScoreChangeRequest) GetPreviousPeriod() *TimeRange {
	if x != nil {
		return x.PreviousPeriod
	}
	return nil
}

// START: Response types
type ScorePerPeriod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Period string  `protobuf:"bytes,1,opt,name=period,proto3" json:"period,omitempty"`
	Score  float64 `protobuf:"fixed64,2,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *ScorePerPeriod) Reset() {
	*x = ScorePerPeriod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_score_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScorePerPeriod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScorePerPeriod) ProtoMessage() {}

func (x *ScorePerPeriod) ProtoReflect() protoreflect.Message {
	mi := &file_score_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScorePerPeriod.ProtoReflect.Descriptor instead.
func (*ScorePerPeriod) Descriptor() ([]byte, []int) {
	return file_score_service_proto_rawDescGZIP(), []int{5}
}

func (x *ScorePerPeriod) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *ScorePerPeriod) GetScore() float64 {
	if x != nil {
		return x.Score
	}
	return 0
}

type CategoryScores struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category        string            `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	ScoresPerPeriod []*ScorePerPeriod `protobuf:"bytes,2,rep,name=scoresPerPeriod,proto3" json:"scoresPerPeriod,omitempty"`
}

func (x *CategoryScores) Reset() {
	*x = CategoryScores{}
	if protoimpl.UnsafeEnabled {
		mi := &file_score_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryScores) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryScores) ProtoMessage() {}

func (x *CategoryScores) ProtoReflect() protoreflect.Message {
	mi := &file_score_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryScores.ProtoReflect.Descriptor instead.
func (*CategoryScores) Descriptor() ([]byte, []int) {
	return file_score_service_proto_rawDescGZIP(), []int{6}
}

func (x *CategoryScores) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *CategoryScores) GetScoresPerPeriod() []*ScorePerPeriod {
	if x != nil {
		return x.ScoresPerPeriod
	}
	return nil
}

type AggregatedCategoryScoresResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scores []*CategoryScores `protobuf:"bytes,1,rep,name=scores,proto3" json:"scores,omitempty"`
}

func (x *AggregatedCategoryScoresResponse) Reset() {
	*x = AggregatedCategoryScoresResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_score_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggregatedCategoryScoresResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregatedCategoryScoresResponse) ProtoMessage() {}

func (x *AggregatedCategoryScoresResponse) ProtoReflect() protoreflect.Message {
	mi := &file_score_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregatedCategoryScoresResponse.ProtoReflect.Descriptor instead.
func (*AggregatedCategoryScoresResponse) Descriptor() ([]byte, []int) {
	return file_score_service_proto_rawDescGZIP(), []int{7}
}

func (x *AggregatedCategoryScoresResponse) GetScores() []*CategoryScores {
	if x != nil {
		return x.Scores
	}
	return nil
}

type TicketScore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TicketId       string             `protobuf:"bytes,1,opt,name=ticketId,proto3" json:"ticketId,omitempty"`
	CategoryScores map[string]float64 `protobuf:"bytes,2,rep,name=categoryScores,proto3" json:"categoryScores,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (x *TicketScore) Reset() {
	*x = TicketScore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_score_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketScore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketScore) ProtoMessage() {}

func (x *TicketScore) ProtoReflect() protoreflect.Message {
	mi := &file_score_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketScore.ProtoReflect.Descriptor instead.
func (*TicketScore) Descriptor() ([]byte, []int) {
	return file_score_service_proto_rawDescGZIP(), []int{8}
}

func (x *TicketScore) GetTicketId() string {
	if x != nil {
		return x.TicketId
	}
	return ""
}

func (x *TicketScore) GetCategoryScores() map[string]float64 {
	if x != nil {
		return x.CategoryScores
	}
	return nil
}

type ScoresByTicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TicketScores []*TicketScore `protobuf:"bytes,1,rep,name=ticketScores,proto3" json:"ticketScores,omitempty"`
}

func (x *ScoresByTicketResponse) Reset() {
	*x = ScoresByTicketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_score_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScoresByTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScoresByTicketResponse) ProtoMessage() {}

func (x *ScoresByTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_score_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScoresByTicketResponse.ProtoReflect.Descriptor instead.
func (*ScoresByTicketResponse) Descriptor() ([]byte, []int) {
	return file_score_service_proto_rawDescGZIP(), []int{9}
}

func (x *ScoresByTicketResponse) GetTicketScores() []*TicketScore {
	if x != nil {
		return x.TicketScores
	}
	return nil
}

type OverallQualityScoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OverallScore float64 `protobuf:"fixed64,1,opt,name=overallScore,proto3" json:"overallScore,omitempty"`
}

func (x *OverallQualityScoreResponse) Reset() {
	*x = OverallQualityScoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_score_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OverallQualityScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverallQualityScoreResponse) ProtoMessage() {}

func (x *OverallQualityScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_score_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverallQualityScoreResponse.ProtoReflect.Descriptor instead.
func (*OverallQualityScoreResponse) Descriptor() ([]byte, []int) {
	return file_score_service_proto_rawDescGZIP(), []int{10}
}

func (x *OverallQualityScoreResponse) GetOverallScore() float64 {
	if x != nil {
		return x.OverallScore
	}
	return 0
}

type PeriodOverPeriodScoreChangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PercentageChange float64 `protobuf:"fixed64,1,opt,name=percentageChange,proto3" json:"percentageChange,omitempty"`
}

func (x *PeriodOverPeriodScoreChangeResponse) Reset() {
	*x = PeriodOverPeriodScoreChangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_score_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeriodOverPeriodScoreChangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeriodOverPeriodScoreChangeResponse) ProtoMessage() {}

func (x *PeriodOverPeriodScoreChangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_score_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeriodOverPeriodScoreChangeResponse.ProtoReflect.Descriptor instead.
func (*PeriodOverPeriodScoreChangeResponse) Descriptor() ([]byte, []int) {
	return file_score_service_proto_rawDescGZIP(), []int{11}
}

func (x *PeriodOverPeriodScoreChangeResponse) GetPercentageChange() float64 {
	if x != nil {
		return x.PercentageChange
	}
	return 0
}

var File_score_service_proto protoreflect.FileDescriptor

var file_score_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65,
	0x6e, 0x64, 0x22, 0x52, 0x0a, 0x1e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x06,
	0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x22, 0x49, 0x0a, 0x15, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x42, 0x79, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x30, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x22, 0x4e, 0x0a, 0x1a, 0x4f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x51, 0x75, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x30, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x22, 0xa6, 0x01, 0x0a, 0x22, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x4f, 0x76, 0x65, 0x72,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x40, 0x0a, 0x0e, 0x70, 0x72, 0x65, 0x76,
	0x69, 0x6f, 0x75, 0x73, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0e, 0x70, 0x72, 0x65, 0x76,
	0x69, 0x6f, 0x75, 0x73, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x22, 0x3e, 0x0a, 0x0e, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x50, 0x65, 0x72, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x75, 0x0a, 0x0e, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x47, 0x0a, 0x0f, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x50, 0x65, 0x72, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x50, 0x65, 0x72, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x52, 0x0f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x50, 0x65, 0x72, 0x50, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x22, 0x59, 0x0a, 0x20, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x73, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x22, 0xc4, 0x01, 0x0a,
	0x0b, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x56, 0x0a, 0x0e, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x1a, 0x41, 0x0a, 0x13, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x58, 0x0a, 0x16, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x42, 0x79, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a,
	0x0c, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52,
	0x0c, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x22, 0x41, 0x0a,
	0x1b, 0x4f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0c, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x22, 0x51, 0x0a, 0x23, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x4f, 0x76, 0x65, 0x72, 0x50, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x61, 0x67, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x10, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x32, 0xea, 0x03, 0x0a, 0x0c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x7d, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x73, 0x12, 0x2d, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x42, 0x79, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x24, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x42,
	0x79, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x73, 0x42, 0x79, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6f, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72,
	0x61, 0x6c, 0x6c, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12,
	0x29, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x4f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x61,
	0x6c, 0x6c, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x87, 0x01, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x50, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x4f, 0x76, 0x65, 0x72, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x31, 0x2e, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x4f, 0x76, 0x65, 0x72, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x4f, 0x76, 0x65, 0x72, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_score_service_proto_rawDescOnce sync.Once
	file_score_service_proto_rawDescData = file_score_service_proto_rawDesc
)

func file_score_service_proto_rawDescGZIP() []byte {
	file_score_service_proto_rawDescOnce.Do(func() {
		file_score_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_score_service_proto_rawDescData)
	})
	return file_score_service_proto_rawDescData
}

var file_score_service_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_score_service_proto_goTypes = []any{
	(*TimeRange)(nil),                           // 0: score_service.TimeRange
	(*AggregatedCategoryScoreRequest)(nil),      // 1: score_service.AggregatedCategoryScoreRequest
	(*ScoresByTicketRequest)(nil),               // 2: score_service.ScoresByTicketRequest
	(*OverallQualityScoreRequest)(nil),          // 3: score_service.OverallQualityScoreRequest
	(*PeriodOverPeriodScoreChangeRequest)(nil),  // 4: score_service.PeriodOverPeriodScoreChangeRequest
	(*ScorePerPeriod)(nil),                      // 5: score_service.ScorePerPeriod
	(*CategoryScores)(nil),                      // 6: score_service.CategoryScores
	(*AggregatedCategoryScoresResponse)(nil),    // 7: score_service.AggregatedCategoryScoresResponse
	(*TicketScore)(nil),                         // 8: score_service.TicketScore
	(*ScoresByTicketResponse)(nil),              // 9: score_service.ScoresByTicketResponse
	(*OverallQualityScoreResponse)(nil),         // 10: score_service.OverallQualityScoreResponse
	(*PeriodOverPeriodScoreChangeResponse)(nil), // 11: score_service.PeriodOverPeriodScoreChangeResponse
	nil,                           // 12: score_service.TicketScore.CategoryScoresEntry
	(*timestamppb.Timestamp)(nil), // 13: google.protobuf.Timestamp
}
var file_score_service_proto_depIdxs = []int32{
	13, // 0: score_service.TimeRange.start:type_name -> google.protobuf.Timestamp
	13, // 1: score_service.TimeRange.end:type_name -> google.protobuf.Timestamp
	0,  // 2: score_service.AggregatedCategoryScoreRequest.period:type_name -> score_service.TimeRange
	0,  // 3: score_service.ScoresByTicketRequest.period:type_name -> score_service.TimeRange
	0,  // 4: score_service.OverallQualityScoreRequest.period:type_name -> score_service.TimeRange
	0,  // 5: score_service.PeriodOverPeriodScoreChangeRequest.currentPeriod:type_name -> score_service.TimeRange
	0,  // 6: score_service.PeriodOverPeriodScoreChangeRequest.previousPeriod:type_name -> score_service.TimeRange
	5,  // 7: score_service.CategoryScores.scoresPerPeriod:type_name -> score_service.ScorePerPeriod
	6,  // 8: score_service.AggregatedCategoryScoresResponse.scores:type_name -> score_service.CategoryScores
	12, // 9: score_service.TicketScore.categoryScores:type_name -> score_service.TicketScore.CategoryScoresEntry
	8,  // 10: score_service.ScoresByTicketResponse.ticketScores:type_name -> score_service.TicketScore
	1,  // 11: score_service.ScoreService.GetAggregatedCategoryScores:input_type -> score_service.AggregatedCategoryScoreRequest
	2,  // 12: score_service.ScoreService.GetScoresByTicket:input_type -> score_service.ScoresByTicketRequest
	3,  // 13: score_service.ScoreService.GetOverallQualityScore:input_type -> score_service.OverallQualityScoreRequest
	4,  // 14: score_service.ScoreService.GetPeriodOverPeriodScoreChange:input_type -> score_service.PeriodOverPeriodScoreChangeRequest
	7,  // 15: score_service.ScoreService.GetAggregatedCategoryScores:output_type -> score_service.AggregatedCategoryScoresResponse
	9,  // 16: score_service.ScoreService.GetScoresByTicket:output_type -> score_service.ScoresByTicketResponse
	10, // 17: score_service.ScoreService.GetOverallQualityScore:output_type -> score_service.OverallQualityScoreResponse
	11, // 18: score_service.ScoreService.GetPeriodOverPeriodScoreChange:output_type -> score_service.PeriodOverPeriodScoreChangeResponse
	15, // [15:19] is the sub-list for method output_type
	11, // [11:15] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_score_service_proto_init() }
func file_score_service_proto_init() {
	if File_score_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_score_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TimeRange); i {
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
		file_score_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AggregatedCategoryScoreRequest); i {
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
		file_score_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ScoresByTicketRequest); i {
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
		file_score_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*OverallQualityScoreRequest); i {
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
		file_score_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*PeriodOverPeriodScoreChangeRequest); i {
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
		file_score_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ScorePerPeriod); i {
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
		file_score_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CategoryScores); i {
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
		file_score_service_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*AggregatedCategoryScoresResponse); i {
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
		file_score_service_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*TicketScore); i {
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
		file_score_service_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*ScoresByTicketResponse); i {
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
		file_score_service_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*OverallQualityScoreResponse); i {
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
		file_score_service_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*PeriodOverPeriodScoreChangeResponse); i {
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
			RawDescriptor: file_score_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_score_service_proto_goTypes,
		DependencyIndexes: file_score_service_proto_depIdxs,
		MessageInfos:      file_score_service_proto_msgTypes,
	}.Build()
	File_score_service_proto = out.File
	file_score_service_proto_rawDesc = nil
	file_score_service_proto_goTypes = nil
	file_score_service_proto_depIdxs = nil
}
