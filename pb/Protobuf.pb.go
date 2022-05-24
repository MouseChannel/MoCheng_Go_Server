// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: Protobuf.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PbMessage_CMD int32

const (
	PbMessage_login PbMessage_CMD = 0
	PbMessage_match PbMessage_CMD = 1
	PbMessage_room  PbMessage_CMD = 2
	// 这个fight只是服务器发客户端
	PbMessage_fight    PbMessage_CMD = 3
	PbMessage_chat     PbMessage_CMD = 4
	PbMessage_fightEnd PbMessage_CMD = 5
)

// Enum value maps for PbMessage_CMD.
var (
	PbMessage_CMD_name = map[int32]string{
		0: "login",
		1: "match",
		2: "room",
		3: "fight",
		4: "chat",
		5: "fightEnd",
	}
	PbMessage_CMD_value = map[string]int32{
		"login":    0,
		"match":    1,
		"room":     2,
		"fight":    3,
		"chat":     4,
		"fightEnd": 5,
	}
)

func (x PbMessage_CMD) Enum() *PbMessage_CMD {
	p := new(PbMessage_CMD)
	*p = x
	return p
}

func (x PbMessage_CMD) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PbMessage_CMD) Descriptor() protoreflect.EnumDescriptor {
	return file_Protobuf_proto_enumTypes[0].Descriptor()
}

func (PbMessage_CMD) Type() protoreflect.EnumType {
	return &file_Protobuf_proto_enumTypes[0]
}

func (x PbMessage_CMD) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PbMessage_CMD.Descriptor instead.
func (PbMessage_CMD) EnumDescriptor() ([]byte, []int) {
	return file_Protobuf_proto_rawDescGZIP(), []int{2, 0}
}

type PbMessage_CmdMatch int32

const (
	PbMessage_joinMatch PbMessage_CmdMatch = 0
	PbMessage_quitMatch PbMessage_CmdMatch = 1
)

// Enum value maps for PbMessage_CmdMatch.
var (
	PbMessage_CmdMatch_name = map[int32]string{
		0: "joinMatch",
		1: "quitMatch",
	}
	PbMessage_CmdMatch_value = map[string]int32{
		"joinMatch": 0,
		"quitMatch": 1,
	}
)

func (x PbMessage_CmdMatch) Enum() *PbMessage_CmdMatch {
	p := new(PbMessage_CmdMatch)
	*p = x
	return p
}

func (x PbMessage_CmdMatch) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PbMessage_CmdMatch) Descriptor() protoreflect.EnumDescriptor {
	return file_Protobuf_proto_enumTypes[1].Descriptor()
}

func (PbMessage_CmdMatch) Type() protoreflect.EnumType {
	return &file_Protobuf_proto_enumTypes[1]
}

func (x PbMessage_CmdMatch) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PbMessage_CmdMatch.Descriptor instead.
func (PbMessage_CmdMatch) EnumDescriptor() ([]byte, []int) {
	return file_Protobuf_proto_rawDescGZIP(), []int{2, 1}
}

type PbMessage_CmdRoom int32

const (
	PbMessage_confirm    PbMessage_CmdRoom = 0
	PbMessage_select     PbMessage_CmdRoom = 1
	PbMessage_selectDate PbMessage_CmdRoom = 2
	PbMessage_load       PbMessage_CmdRoom = 3
	PbMessage_loadData   PbMessage_CmdRoom = 4
	PbMessage_fightStart PbMessage_CmdRoom = 5
	//这个fightop只是客户端发服务器
	PbMessage_fightOp   PbMessage_CmdRoom = 6
	PbMessage_dismissed PbMessage_CmdRoom = 7
)

// Enum value maps for PbMessage_CmdRoom.
var (
	PbMessage_CmdRoom_name = map[int32]string{
		0: "confirm",
		1: "select",
		2: "selectDate",
		3: "load",
		4: "loadData",
		5: "fightStart",
		6: "fightOp",
		7: "dismissed",
	}
	PbMessage_CmdRoom_value = map[string]int32{
		"confirm":    0,
		"select":     1,
		"selectDate": 2,
		"load":       3,
		"loadData":   4,
		"fightStart": 5,
		"fightOp":    6,
		"dismissed":  7,
	}
)

func (x PbMessage_CmdRoom) Enum() *PbMessage_CmdRoom {
	p := new(PbMessage_CmdRoom)
	*p = x
	return p
}

func (x PbMessage_CmdRoom) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PbMessage_CmdRoom) Descriptor() protoreflect.EnumDescriptor {
	return file_Protobuf_proto_enumTypes[2].Descriptor()
}

func (PbMessage_CmdRoom) Type() protoreflect.EnumType {
	return &file_Protobuf_proto_enumTypes[2]
}

func (x PbMessage_CmdRoom) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PbMessage_CmdRoom.Descriptor instead.
func (PbMessage_CmdRoom) EnumDescriptor() ([]byte, []int) {
	return file_Protobuf_proto_rawDescGZIP(), []int{2, 2}
}

type FightMessage_BattleCMD int32

const (
	FightMessage_move     FightMessage_BattleCMD = 0
	FightMessage_fight    FightMessage_BattleCMD = 1
	FightMessage_interact FightMessage_BattleCMD = 2
)

// Enum value maps for FightMessage_BattleCMD.
var (
	FightMessage_BattleCMD_name = map[int32]string{
		0: "move",
		1: "fight",
		2: "interact",
	}
	FightMessage_BattleCMD_value = map[string]int32{
		"move":     0,
		"fight":    1,
		"interact": 2,
	}
)

func (x FightMessage_BattleCMD) Enum() *FightMessage_BattleCMD {
	p := new(FightMessage_BattleCMD)
	*p = x
	return p
}

func (x FightMessage_BattleCMD) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FightMessage_BattleCMD) Descriptor() protoreflect.EnumDescriptor {
	return file_Protobuf_proto_enumTypes[3].Descriptor()
}

func (FightMessage_BattleCMD) Type() protoreflect.EnumType {
	return &file_Protobuf_proto_enumTypes[3]
}

func (x FightMessage_BattleCMD) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FightMessage_BattleCMD.Descriptor instead.
func (FightMessage_BattleCMD) EnumDescriptor() ([]byte, []int) {
	return file_Protobuf_proto_rawDescGZIP(), []int{3, 0}
}

type AccountData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AccountData) Reset() {
	*x = AccountData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Protobuf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountData) ProtoMessage() {}

func (x *AccountData) ProtoReflect() protoreflect.Message {
	mi := &file_Protobuf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountData.ProtoReflect.Descriptor instead.
func (*AccountData) Descriptor() ([]byte, []int) {
	return file_Protobuf_proto_rawDescGZIP(), []int{0}
}

func (x *AccountData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SelectData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerName string `protobuf:"bytes,2,opt,name=playerName,proto3" json:"playerName,omitempty"`
	ChatMes    string `protobuf:"bytes,3,opt,name=chatMes,proto3" json:"chatMes,omitempty"`
	Faction    int32  `protobuf:"varint,6,opt,name=faction,proto3" json:"faction,omitempty"`
	IsReady    bool   `protobuf:"varint,7,opt,name=isReady,proto3" json:"isReady,omitempty"`
}

func (x *SelectData) Reset() {
	*x = SelectData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Protobuf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectData) ProtoMessage() {}

func (x *SelectData) ProtoReflect() protoreflect.Message {
	mi := &file_Protobuf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectData.ProtoReflect.Descriptor instead.
func (*SelectData) Descriptor() ([]byte, []int) {
	return file_Protobuf_proto_rawDescGZIP(), []int{1}
}

func (x *SelectData) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

func (x *SelectData) GetChatMes() string {
	if x != nil {
		return x.ChatMes
	}
	return ""
}

func (x *SelectData) GetFaction() int32 {
	if x != nil {
		return x.Faction
	}
	return 0
}

func (x *SelectData) GetIsReady() bool {
	if x != nil {
		return x.IsReady
	}
	return false
}

type PbMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Cmd              PbMessage_CMD      `protobuf:"varint,2,opt,name=cmd,proto3,enum=pb.PbMessage_CMD" json:"cmd,omitempty"`
	CmdMatch         PbMessage_CmdMatch `protobuf:"varint,3,opt,name=cmdMatch,proto3,enum=pb.PbMessage_CmdMatch" json:"cmdMatch,omitempty"`
	CmdRoom          PbMessage_CmdRoom  `protobuf:"varint,4,opt,name=cmdRoom,proto3,enum=pb.PbMessage_CmdRoom" json:"cmdRoom,omitempty"`
	RoomIndex        int32              `protobuf:"varint,5,opt,name=roomIndex,proto3" json:"roomIndex,omitempty"`
	ChatMes          string             `protobuf:"bytes,6,opt,name=chatMes,proto3" json:"chatMes,omitempty"`
	SelectData       *SelectData        `protobuf:"bytes,7,opt,name=selectData,proto3" json:"selectData,omitempty"`
	RoomSelectData   []*SelectData      `protobuf:"bytes,8,rep,name=roomSelectData,proto3" json:"roomSelectData,omitempty"`
	LoadPercent      int32              `protobuf:"varint,9,opt,name=loadPercent,proto3" json:"loadPercent,omitempty"`
	RoomLoadPercent  []int32            `protobuf:"varint,10,rep,packed,name=roomLoadPercent,proto3" json:"roomLoadPercent,omitempty"`
	FrameId          int32              `protobuf:"varint,11,opt,name=frameId,proto3" json:"frameId,omitempty"`
	SendFightMessage *FightMessage      `protobuf:"bytes,12,opt,name=sendFightMessage,proto3" json:"sendFightMessage,omitempty"`
	FightMessage     []*FightMessage    `protobuf:"bytes,13,rep,name=fightMessage,proto3" json:"fightMessage,omitempty"`
}

func (x *PbMessage) Reset() {
	*x = PbMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Protobuf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PbMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PbMessage) ProtoMessage() {}

func (x *PbMessage) ProtoReflect() protoreflect.Message {
	mi := &file_Protobuf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PbMessage.ProtoReflect.Descriptor instead.
func (*PbMessage) Descriptor() ([]byte, []int) {
	return file_Protobuf_proto_rawDescGZIP(), []int{2}
}

func (x *PbMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PbMessage) GetCmd() PbMessage_CMD {
	if x != nil {
		return x.Cmd
	}
	return PbMessage_login
}

func (x *PbMessage) GetCmdMatch() PbMessage_CmdMatch {
	if x != nil {
		return x.CmdMatch
	}
	return PbMessage_joinMatch
}

func (x *PbMessage) GetCmdRoom() PbMessage_CmdRoom {
	if x != nil {
		return x.CmdRoom
	}
	return PbMessage_confirm
}

func (x *PbMessage) GetRoomIndex() int32 {
	if x != nil {
		return x.RoomIndex
	}
	return 0
}

func (x *PbMessage) GetChatMes() string {
	if x != nil {
		return x.ChatMes
	}
	return ""
}

func (x *PbMessage) GetSelectData() *SelectData {
	if x != nil {
		return x.SelectData
	}
	return nil
}

func (x *PbMessage) GetRoomSelectData() []*SelectData {
	if x != nil {
		return x.RoomSelectData
	}
	return nil
}

func (x *PbMessage) GetLoadPercent() int32 {
	if x != nil {
		return x.LoadPercent
	}
	return 0
}

func (x *PbMessage) GetRoomLoadPercent() []int32 {
	if x != nil {
		return x.RoomLoadPercent
	}
	return nil
}

func (x *PbMessage) GetFrameId() int32 {
	if x != nil {
		return x.FrameId
	}
	return 0
}

func (x *PbMessage) GetSendFightMessage() *FightMessage {
	if x != nil {
		return x.SendFightMessage
	}
	return nil
}

func (x *PbMessage) GetFightMessage() []*FightMessage {
	if x != nil {
		return x.FightMessage
	}
	return nil
}

type FightMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BattleCMD    FightMessage_BattleCMD `protobuf:"varint,1,opt,name=battleCMD,proto3,enum=pb.FightMessage_BattleCMD" json:"battleCMD,omitempty"`
	SelectedUnit []int32                `protobuf:"varint,3,rep,packed,name=selectedUnit,proto3" json:"selectedUnit,omitempty"`
	// move
	StartPos int32 `protobuf:"varint,4,opt,name=startPos,proto3" json:"startPos,omitempty"`
	EndPos   int32 `protobuf:"varint,5,opt,name=endPos,proto3" json:"endPos,omitempty"`
	//fight
	EnemyUnit int32 `protobuf:"varint,6,opt,name=enemyUnit,proto3" json:"enemyUnit,omitempty"`
	//interact
	InteractObject int32 `protobuf:"varint,7,opt,name=interactObject,proto3" json:"interactObject,omitempty"`
}

func (x *FightMessage) Reset() {
	*x = FightMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Protobuf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FightMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FightMessage) ProtoMessage() {}

func (x *FightMessage) ProtoReflect() protoreflect.Message {
	mi := &file_Protobuf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FightMessage.ProtoReflect.Descriptor instead.
func (*FightMessage) Descriptor() ([]byte, []int) {
	return file_Protobuf_proto_rawDescGZIP(), []int{3}
}

func (x *FightMessage) GetBattleCMD() FightMessage_BattleCMD {
	if x != nil {
		return x.BattleCMD
	}
	return FightMessage_move
}

func (x *FightMessage) GetSelectedUnit() []int32 {
	if x != nil {
		return x.SelectedUnit
	}
	return nil
}

func (x *FightMessage) GetStartPos() int32 {
	if x != nil {
		return x.StartPos
	}
	return 0
}

func (x *FightMessage) GetEndPos() int32 {
	if x != nil {
		return x.EndPos
	}
	return 0
}

func (x *FightMessage) GetEnemyUnit() int32 {
	if x != nil {
		return x.EnemyUnit
	}
	return 0
}

func (x *FightMessage) GetInteractObject() int32 {
	if x != nil {
		return x.InteractObject
	}
	return 0
}

var File_Protobuf_proto protoreflect.FileDescriptor

var file_Protobuf_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0x21, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x7a, 0x0a, 0x0a, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x66, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x66, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x52,
	0x65, 0x61, 0x64, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x52, 0x65,
	0x61, 0x64, 0x79, 0x22, 0x8f, 0x06, 0x0a, 0x09, 0x50, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x43, 0x4d, 0x44, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x32, 0x0a, 0x08, 0x63, 0x6d,
	0x64, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70,
	0x62, 0x2e, 0x50, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6d, 0x64, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x08, 0x63, 0x6d, 0x64, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x2f,
	0x0a, 0x07, 0x63, 0x6d, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x15, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43,
	0x6d, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x07, 0x63, 0x6d, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x12,
	0x1c, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x0a, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62,
	0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x36, 0x0a, 0x0e, 0x72, 0x6f, 0x6f, 0x6d, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x0e, 0x72, 0x6f, 0x6f, 0x6d, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x20, 0x0a, 0x0b, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x6f, 0x6f, 0x6d, 0x4c, 0x6f, 0x61, 0x64, 0x50, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0f, 0x72, 0x6f, 0x6f, 0x6d,
	0x4c, 0x6f, 0x61, 0x64, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x66, 0x72,
	0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x10, 0x73, 0x65, 0x6e, 0x64, 0x46, 0x69, 0x67,
	0x68, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x67, 0x68, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x10, 0x73, 0x65, 0x6e, 0x64, 0x46, 0x69, 0x67, 0x68, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x0c, 0x66, 0x69, 0x67, 0x68, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x46,
	0x69, 0x67, 0x68, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0c, 0x66, 0x69, 0x67,
	0x68, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x48, 0x0a, 0x03, 0x43, 0x4d, 0x44,
	0x12, 0x09, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x10, 0x02,
	0x12, 0x09, 0x0a, 0x05, 0x66, 0x69, 0x67, 0x68, 0x74, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x63,
	0x68, 0x61, 0x74, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x66, 0x69, 0x67, 0x68, 0x74, 0x45, 0x6e,
	0x64, 0x10, 0x05, 0x22, 0x28, 0x0a, 0x08, 0x43, 0x6d, 0x64, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x0d, 0x0a, 0x09, 0x6a, 0x6f, 0x69, 0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x10, 0x00, 0x12, 0x0d,
	0x0a, 0x09, 0x71, 0x75, 0x69, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x10, 0x01, 0x22, 0x76, 0x0a,
	0x07, 0x43, 0x6d, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x0b, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x10,
	0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x44, 0x61, 0x74, 0x65, 0x10,
	0x02, 0x12, 0x08, 0x0a, 0x04, 0x6c, 0x6f, 0x61, 0x64, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x6c,
	0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x66, 0x69, 0x67,
	0x68, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x66, 0x69, 0x67,
	0x68, 0x74, 0x4f, 0x70, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x6d, 0x69, 0x73,
	0x73, 0x65, 0x64, 0x10, 0x07, 0x22, 0x96, 0x02, 0x0a, 0x0c, 0x46, 0x69, 0x67, 0x68, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x43, 0x4d, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x46,
	0x69, 0x67, 0x68, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x42, 0x61, 0x74, 0x74,
	0x6c, 0x65, 0x43, 0x4d, 0x44, 0x52, 0x09, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x43, 0x4d, 0x44,
	0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x55, 0x6e, 0x69, 0x74,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x55, 0x6e, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x6f, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x6f, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x50, 0x6f, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x65, 0x6e, 0x64, 0x50, 0x6f, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x65, 0x6d,
	0x79, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x6e, 0x65,
	0x6d, 0x79, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2e,
	0x0a, 0x09, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x43, 0x4d, 0x44, 0x12, 0x08, 0x0a, 0x04, 0x6d,
	0x6f, 0x76, 0x65, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x66, 0x69, 0x67, 0x68, 0x74, 0x10, 0x01,
	0x12, 0x0c, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x10, 0x02, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Protobuf_proto_rawDescOnce sync.Once
	file_Protobuf_proto_rawDescData = file_Protobuf_proto_rawDesc
)

func file_Protobuf_proto_rawDescGZIP() []byte {
	file_Protobuf_proto_rawDescOnce.Do(func() {
		file_Protobuf_proto_rawDescData = protoimpl.X.CompressGZIP(file_Protobuf_proto_rawDescData)
	})
	return file_Protobuf_proto_rawDescData
}

var file_Protobuf_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_Protobuf_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Protobuf_proto_goTypes = []interface{}{
	(PbMessage_CMD)(0),          // 0: pb.PbMessage.CMD
	(PbMessage_CmdMatch)(0),     // 1: pb.PbMessage.CmdMatch
	(PbMessage_CmdRoom)(0),      // 2: pb.PbMessage.CmdRoom
	(FightMessage_BattleCMD)(0), // 3: pb.FightMessage.BattleCMD
	(*AccountData)(nil),         // 4: pb.AccountData
	(*SelectData)(nil),          // 5: pb.SelectData
	(*PbMessage)(nil),           // 6: pb.PbMessage
	(*FightMessage)(nil),        // 7: pb.FightMessage
}
var file_Protobuf_proto_depIdxs = []int32{
	0, // 0: pb.PbMessage.cmd:type_name -> pb.PbMessage.CMD
	1, // 1: pb.PbMessage.cmdMatch:type_name -> pb.PbMessage.CmdMatch
	2, // 2: pb.PbMessage.cmdRoom:type_name -> pb.PbMessage.CmdRoom
	5, // 3: pb.PbMessage.selectData:type_name -> pb.SelectData
	5, // 4: pb.PbMessage.roomSelectData:type_name -> pb.SelectData
	7, // 5: pb.PbMessage.sendFightMessage:type_name -> pb.FightMessage
	7, // 6: pb.PbMessage.fightMessage:type_name -> pb.FightMessage
	3, // 7: pb.FightMessage.battleCMD:type_name -> pb.FightMessage.BattleCMD
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_Protobuf_proto_init() }
func file_Protobuf_proto_init() {
	if File_Protobuf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Protobuf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountData); i {
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
		file_Protobuf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectData); i {
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
		file_Protobuf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PbMessage); i {
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
		file_Protobuf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FightMessage); i {
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
			RawDescriptor: file_Protobuf_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Protobuf_proto_goTypes,
		DependencyIndexes: file_Protobuf_proto_depIdxs,
		EnumInfos:         file_Protobuf_proto_enumTypes,
		MessageInfos:      file_Protobuf_proto_msgTypes,
	}.Build()
	File_Protobuf_proto = out.File
	file_Protobuf_proto_rawDesc = nil
	file_Protobuf_proto_goTypes = nil
	file_Protobuf_proto_depIdxs = nil
}
