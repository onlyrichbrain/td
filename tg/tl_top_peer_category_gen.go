// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is

// TopPeerCategoryBotsPM represents TL type `topPeerCategoryBotsPM#ab661b5b`.
// Most used bots
//
// See https://core.telegram.org/constructor/topPeerCategoryBotsPM for reference.
type TopPeerCategoryBotsPM struct {
}

// TopPeerCategoryBotsPMTypeID is TL type id of TopPeerCategoryBotsPM.
const TopPeerCategoryBotsPMTypeID = 0xab661b5b

func (t *TopPeerCategoryBotsPM) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *TopPeerCategoryBotsPM) String() string {
	if t == nil {
		return "TopPeerCategoryBotsPM(nil)"
	}
	type Alias TopPeerCategoryBotsPM
	return fmt.Sprintf("TopPeerCategoryBotsPM%+v", Alias(*t))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (t *TopPeerCategoryBotsPM) TypeID() uint32 {
	return TopPeerCategoryBotsPMTypeID
}

// SchemaName returns MTProto type name.
func (t *TopPeerCategoryBotsPM) SchemaName() string {
	return "topPeerCategoryBotsPM"
}

// Encode implements bin.Encoder.
func (t *TopPeerCategoryBotsPM) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topPeerCategoryBotsPM#ab661b5b as nil")
	}
	b.PutID(TopPeerCategoryBotsPMTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (t *TopPeerCategoryBotsPM) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topPeerCategoryBotsPM#ab661b5b to nil")
	}
	if err := b.ConsumeID(TopPeerCategoryBotsPMTypeID); err != nil {
		return fmt.Errorf("unable to decode topPeerCategoryBotsPM#ab661b5b: %w", err)
	}
	return nil
}

// construct implements constructor of TopPeerCategoryClass.
func (t TopPeerCategoryBotsPM) construct() TopPeerCategoryClass { return &t }

// Ensuring interfaces in compile-time for TopPeerCategoryBotsPM.
var (
	_ bin.Encoder = &TopPeerCategoryBotsPM{}
	_ bin.Decoder = &TopPeerCategoryBotsPM{}

	_ TopPeerCategoryClass = &TopPeerCategoryBotsPM{}
)

// TopPeerCategoryBotsInline represents TL type `topPeerCategoryBotsInline#148677e2`.
// Most used inline bots
//
// See https://core.telegram.org/constructor/topPeerCategoryBotsInline for reference.
type TopPeerCategoryBotsInline struct {
}

// TopPeerCategoryBotsInlineTypeID is TL type id of TopPeerCategoryBotsInline.
const TopPeerCategoryBotsInlineTypeID = 0x148677e2

func (t *TopPeerCategoryBotsInline) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *TopPeerCategoryBotsInline) String() string {
	if t == nil {
		return "TopPeerCategoryBotsInline(nil)"
	}
	type Alias TopPeerCategoryBotsInline
	return fmt.Sprintf("TopPeerCategoryBotsInline%+v", Alias(*t))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (t *TopPeerCategoryBotsInline) TypeID() uint32 {
	return TopPeerCategoryBotsInlineTypeID
}

// SchemaName returns MTProto type name.
func (t *TopPeerCategoryBotsInline) SchemaName() string {
	return "topPeerCategoryBotsInline"
}

// Encode implements bin.Encoder.
func (t *TopPeerCategoryBotsInline) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topPeerCategoryBotsInline#148677e2 as nil")
	}
	b.PutID(TopPeerCategoryBotsInlineTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (t *TopPeerCategoryBotsInline) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topPeerCategoryBotsInline#148677e2 to nil")
	}
	if err := b.ConsumeID(TopPeerCategoryBotsInlineTypeID); err != nil {
		return fmt.Errorf("unable to decode topPeerCategoryBotsInline#148677e2: %w", err)
	}
	return nil
}

// construct implements constructor of TopPeerCategoryClass.
func (t TopPeerCategoryBotsInline) construct() TopPeerCategoryClass { return &t }

// Ensuring interfaces in compile-time for TopPeerCategoryBotsInline.
var (
	_ bin.Encoder = &TopPeerCategoryBotsInline{}
	_ bin.Decoder = &TopPeerCategoryBotsInline{}

	_ TopPeerCategoryClass = &TopPeerCategoryBotsInline{}
)

// TopPeerCategoryCorrespondents represents TL type `topPeerCategoryCorrespondents#637b7ed`.
// Users we've chatted most frequently with
//
// See https://core.telegram.org/constructor/topPeerCategoryCorrespondents for reference.
type TopPeerCategoryCorrespondents struct {
}

// TopPeerCategoryCorrespondentsTypeID is TL type id of TopPeerCategoryCorrespondents.
const TopPeerCategoryCorrespondentsTypeID = 0x637b7ed

func (t *TopPeerCategoryCorrespondents) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *TopPeerCategoryCorrespondents) String() string {
	if t == nil {
		return "TopPeerCategoryCorrespondents(nil)"
	}
	type Alias TopPeerCategoryCorrespondents
	return fmt.Sprintf("TopPeerCategoryCorrespondents%+v", Alias(*t))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (t *TopPeerCategoryCorrespondents) TypeID() uint32 {
	return TopPeerCategoryCorrespondentsTypeID
}

// SchemaName returns MTProto type name.
func (t *TopPeerCategoryCorrespondents) SchemaName() string {
	return "topPeerCategoryCorrespondents"
}

// Encode implements bin.Encoder.
func (t *TopPeerCategoryCorrespondents) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topPeerCategoryCorrespondents#637b7ed as nil")
	}
	b.PutID(TopPeerCategoryCorrespondentsTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (t *TopPeerCategoryCorrespondents) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topPeerCategoryCorrespondents#637b7ed to nil")
	}
	if err := b.ConsumeID(TopPeerCategoryCorrespondentsTypeID); err != nil {
		return fmt.Errorf("unable to decode topPeerCategoryCorrespondents#637b7ed: %w", err)
	}
	return nil
}

// construct implements constructor of TopPeerCategoryClass.
func (t TopPeerCategoryCorrespondents) construct() TopPeerCategoryClass { return &t }

// Ensuring interfaces in compile-time for TopPeerCategoryCorrespondents.
var (
	_ bin.Encoder = &TopPeerCategoryCorrespondents{}
	_ bin.Decoder = &TopPeerCategoryCorrespondents{}

	_ TopPeerCategoryClass = &TopPeerCategoryCorrespondents{}
)

// TopPeerCategoryGroups represents TL type `topPeerCategoryGroups#bd17a14a`.
// Often-opened groups and supergroups
//
// See https://core.telegram.org/constructor/topPeerCategoryGroups for reference.
type TopPeerCategoryGroups struct {
}

// TopPeerCategoryGroupsTypeID is TL type id of TopPeerCategoryGroups.
const TopPeerCategoryGroupsTypeID = 0xbd17a14a

func (t *TopPeerCategoryGroups) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *TopPeerCategoryGroups) String() string {
	if t == nil {
		return "TopPeerCategoryGroups(nil)"
	}
	type Alias TopPeerCategoryGroups
	return fmt.Sprintf("TopPeerCategoryGroups%+v", Alias(*t))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (t *TopPeerCategoryGroups) TypeID() uint32 {
	return TopPeerCategoryGroupsTypeID
}

// SchemaName returns MTProto type name.
func (t *TopPeerCategoryGroups) SchemaName() string {
	return "topPeerCategoryGroups"
}

// Encode implements bin.Encoder.
func (t *TopPeerCategoryGroups) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topPeerCategoryGroups#bd17a14a as nil")
	}
	b.PutID(TopPeerCategoryGroupsTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (t *TopPeerCategoryGroups) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topPeerCategoryGroups#bd17a14a to nil")
	}
	if err := b.ConsumeID(TopPeerCategoryGroupsTypeID); err != nil {
		return fmt.Errorf("unable to decode topPeerCategoryGroups#bd17a14a: %w", err)
	}
	return nil
}

// construct implements constructor of TopPeerCategoryClass.
func (t TopPeerCategoryGroups) construct() TopPeerCategoryClass { return &t }

// Ensuring interfaces in compile-time for TopPeerCategoryGroups.
var (
	_ bin.Encoder = &TopPeerCategoryGroups{}
	_ bin.Decoder = &TopPeerCategoryGroups{}

	_ TopPeerCategoryClass = &TopPeerCategoryGroups{}
)

// TopPeerCategoryChannels represents TL type `topPeerCategoryChannels#161d9628`.
// Most frequently visited channels
//
// See https://core.telegram.org/constructor/topPeerCategoryChannels for reference.
type TopPeerCategoryChannels struct {
}

// TopPeerCategoryChannelsTypeID is TL type id of TopPeerCategoryChannels.
const TopPeerCategoryChannelsTypeID = 0x161d9628

func (t *TopPeerCategoryChannels) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *TopPeerCategoryChannels) String() string {
	if t == nil {
		return "TopPeerCategoryChannels(nil)"
	}
	type Alias TopPeerCategoryChannels
	return fmt.Sprintf("TopPeerCategoryChannels%+v", Alias(*t))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (t *TopPeerCategoryChannels) TypeID() uint32 {
	return TopPeerCategoryChannelsTypeID
}

// SchemaName returns MTProto type name.
func (t *TopPeerCategoryChannels) SchemaName() string {
	return "topPeerCategoryChannels"
}

// Encode implements bin.Encoder.
func (t *TopPeerCategoryChannels) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topPeerCategoryChannels#161d9628 as nil")
	}
	b.PutID(TopPeerCategoryChannelsTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (t *TopPeerCategoryChannels) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topPeerCategoryChannels#161d9628 to nil")
	}
	if err := b.ConsumeID(TopPeerCategoryChannelsTypeID); err != nil {
		return fmt.Errorf("unable to decode topPeerCategoryChannels#161d9628: %w", err)
	}
	return nil
}

// construct implements constructor of TopPeerCategoryClass.
func (t TopPeerCategoryChannels) construct() TopPeerCategoryClass { return &t }

// Ensuring interfaces in compile-time for TopPeerCategoryChannels.
var (
	_ bin.Encoder = &TopPeerCategoryChannels{}
	_ bin.Decoder = &TopPeerCategoryChannels{}

	_ TopPeerCategoryClass = &TopPeerCategoryChannels{}
)

// TopPeerCategoryPhoneCalls represents TL type `topPeerCategoryPhoneCalls#1e76a78c`.
// Most frequently called users
//
// See https://core.telegram.org/constructor/topPeerCategoryPhoneCalls for reference.
type TopPeerCategoryPhoneCalls struct {
}

// TopPeerCategoryPhoneCallsTypeID is TL type id of TopPeerCategoryPhoneCalls.
const TopPeerCategoryPhoneCallsTypeID = 0x1e76a78c

func (t *TopPeerCategoryPhoneCalls) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *TopPeerCategoryPhoneCalls) String() string {
	if t == nil {
		return "TopPeerCategoryPhoneCalls(nil)"
	}
	type Alias TopPeerCategoryPhoneCalls
	return fmt.Sprintf("TopPeerCategoryPhoneCalls%+v", Alias(*t))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (t *TopPeerCategoryPhoneCalls) TypeID() uint32 {
	return TopPeerCategoryPhoneCallsTypeID
}

// SchemaName returns MTProto type name.
func (t *TopPeerCategoryPhoneCalls) SchemaName() string {
	return "topPeerCategoryPhoneCalls"
}

// Encode implements bin.Encoder.
func (t *TopPeerCategoryPhoneCalls) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topPeerCategoryPhoneCalls#1e76a78c as nil")
	}
	b.PutID(TopPeerCategoryPhoneCallsTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (t *TopPeerCategoryPhoneCalls) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topPeerCategoryPhoneCalls#1e76a78c to nil")
	}
	if err := b.ConsumeID(TopPeerCategoryPhoneCallsTypeID); err != nil {
		return fmt.Errorf("unable to decode topPeerCategoryPhoneCalls#1e76a78c: %w", err)
	}
	return nil
}

// construct implements constructor of TopPeerCategoryClass.
func (t TopPeerCategoryPhoneCalls) construct() TopPeerCategoryClass { return &t }

// Ensuring interfaces in compile-time for TopPeerCategoryPhoneCalls.
var (
	_ bin.Encoder = &TopPeerCategoryPhoneCalls{}
	_ bin.Decoder = &TopPeerCategoryPhoneCalls{}

	_ TopPeerCategoryClass = &TopPeerCategoryPhoneCalls{}
)

// TopPeerCategoryForwardUsers represents TL type `topPeerCategoryForwardUsers#a8406ca9`.
// Users to which the users often forwards messages to
//
// See https://core.telegram.org/constructor/topPeerCategoryForwardUsers for reference.
type TopPeerCategoryForwardUsers struct {
}

// TopPeerCategoryForwardUsersTypeID is TL type id of TopPeerCategoryForwardUsers.
const TopPeerCategoryForwardUsersTypeID = 0xa8406ca9

func (t *TopPeerCategoryForwardUsers) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *TopPeerCategoryForwardUsers) String() string {
	if t == nil {
		return "TopPeerCategoryForwardUsers(nil)"
	}
	type Alias TopPeerCategoryForwardUsers
	return fmt.Sprintf("TopPeerCategoryForwardUsers%+v", Alias(*t))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (t *TopPeerCategoryForwardUsers) TypeID() uint32 {
	return TopPeerCategoryForwardUsersTypeID
}

// SchemaName returns MTProto type name.
func (t *TopPeerCategoryForwardUsers) SchemaName() string {
	return "topPeerCategoryForwardUsers"
}

// Encode implements bin.Encoder.
func (t *TopPeerCategoryForwardUsers) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topPeerCategoryForwardUsers#a8406ca9 as nil")
	}
	b.PutID(TopPeerCategoryForwardUsersTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (t *TopPeerCategoryForwardUsers) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topPeerCategoryForwardUsers#a8406ca9 to nil")
	}
	if err := b.ConsumeID(TopPeerCategoryForwardUsersTypeID); err != nil {
		return fmt.Errorf("unable to decode topPeerCategoryForwardUsers#a8406ca9: %w", err)
	}
	return nil
}

// construct implements constructor of TopPeerCategoryClass.
func (t TopPeerCategoryForwardUsers) construct() TopPeerCategoryClass { return &t }

// Ensuring interfaces in compile-time for TopPeerCategoryForwardUsers.
var (
	_ bin.Encoder = &TopPeerCategoryForwardUsers{}
	_ bin.Decoder = &TopPeerCategoryForwardUsers{}

	_ TopPeerCategoryClass = &TopPeerCategoryForwardUsers{}
)

// TopPeerCategoryForwardChats represents TL type `topPeerCategoryForwardChats#fbeec0f0`.
// Chats to which the users often forwards messages to
//
// See https://core.telegram.org/constructor/topPeerCategoryForwardChats for reference.
type TopPeerCategoryForwardChats struct {
}

// TopPeerCategoryForwardChatsTypeID is TL type id of TopPeerCategoryForwardChats.
const TopPeerCategoryForwardChatsTypeID = 0xfbeec0f0

func (t *TopPeerCategoryForwardChats) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *TopPeerCategoryForwardChats) String() string {
	if t == nil {
		return "TopPeerCategoryForwardChats(nil)"
	}
	type Alias TopPeerCategoryForwardChats
	return fmt.Sprintf("TopPeerCategoryForwardChats%+v", Alias(*t))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (t *TopPeerCategoryForwardChats) TypeID() uint32 {
	return TopPeerCategoryForwardChatsTypeID
}

// SchemaName returns MTProto type name.
func (t *TopPeerCategoryForwardChats) SchemaName() string {
	return "topPeerCategoryForwardChats"
}

// Encode implements bin.Encoder.
func (t *TopPeerCategoryForwardChats) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topPeerCategoryForwardChats#fbeec0f0 as nil")
	}
	b.PutID(TopPeerCategoryForwardChatsTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (t *TopPeerCategoryForwardChats) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topPeerCategoryForwardChats#fbeec0f0 to nil")
	}
	if err := b.ConsumeID(TopPeerCategoryForwardChatsTypeID); err != nil {
		return fmt.Errorf("unable to decode topPeerCategoryForwardChats#fbeec0f0: %w", err)
	}
	return nil
}

// construct implements constructor of TopPeerCategoryClass.
func (t TopPeerCategoryForwardChats) construct() TopPeerCategoryClass { return &t }

// Ensuring interfaces in compile-time for TopPeerCategoryForwardChats.
var (
	_ bin.Encoder = &TopPeerCategoryForwardChats{}
	_ bin.Decoder = &TopPeerCategoryForwardChats{}

	_ TopPeerCategoryClass = &TopPeerCategoryForwardChats{}
)

// TopPeerCategoryClass represents TopPeerCategory generic type.
//
// See https://core.telegram.org/type/TopPeerCategory for reference.
//
// Example:
//  g, err := tg.DecodeTopPeerCategory(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.TopPeerCategoryBotsPM: // topPeerCategoryBotsPM#ab661b5b
//  case *tg.TopPeerCategoryBotsInline: // topPeerCategoryBotsInline#148677e2
//  case *tg.TopPeerCategoryCorrespondents: // topPeerCategoryCorrespondents#637b7ed
//  case *tg.TopPeerCategoryGroups: // topPeerCategoryGroups#bd17a14a
//  case *tg.TopPeerCategoryChannels: // topPeerCategoryChannels#161d9628
//  case *tg.TopPeerCategoryPhoneCalls: // topPeerCategoryPhoneCalls#1e76a78c
//  case *tg.TopPeerCategoryForwardUsers: // topPeerCategoryForwardUsers#a8406ca9
//  case *tg.TopPeerCategoryForwardChats: // topPeerCategoryForwardChats#fbeec0f0
//  default: panic(v)
//  }
type TopPeerCategoryClass interface {
	bin.Encoder
	bin.Decoder
	construct() TopPeerCategoryClass

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// SchemaName returns MTProto type name.
	SchemaName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeTopPeerCategory implements binary de-serialization for TopPeerCategoryClass.
func DecodeTopPeerCategory(buf *bin.Buffer) (TopPeerCategoryClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case TopPeerCategoryBotsPMTypeID:
		// Decoding topPeerCategoryBotsPM#ab661b5b.
		v := TopPeerCategoryBotsPM{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopPeerCategoryClass: %w", err)
		}
		return &v, nil
	case TopPeerCategoryBotsInlineTypeID:
		// Decoding topPeerCategoryBotsInline#148677e2.
		v := TopPeerCategoryBotsInline{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopPeerCategoryClass: %w", err)
		}
		return &v, nil
	case TopPeerCategoryCorrespondentsTypeID:
		// Decoding topPeerCategoryCorrespondents#637b7ed.
		v := TopPeerCategoryCorrespondents{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopPeerCategoryClass: %w", err)
		}
		return &v, nil
	case TopPeerCategoryGroupsTypeID:
		// Decoding topPeerCategoryGroups#bd17a14a.
		v := TopPeerCategoryGroups{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopPeerCategoryClass: %w", err)
		}
		return &v, nil
	case TopPeerCategoryChannelsTypeID:
		// Decoding topPeerCategoryChannels#161d9628.
		v := TopPeerCategoryChannels{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopPeerCategoryClass: %w", err)
		}
		return &v, nil
	case TopPeerCategoryPhoneCallsTypeID:
		// Decoding topPeerCategoryPhoneCalls#1e76a78c.
		v := TopPeerCategoryPhoneCalls{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopPeerCategoryClass: %w", err)
		}
		return &v, nil
	case TopPeerCategoryForwardUsersTypeID:
		// Decoding topPeerCategoryForwardUsers#a8406ca9.
		v := TopPeerCategoryForwardUsers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopPeerCategoryClass: %w", err)
		}
		return &v, nil
	case TopPeerCategoryForwardChatsTypeID:
		// Decoding topPeerCategoryForwardChats#fbeec0f0.
		v := TopPeerCategoryForwardChats{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopPeerCategoryClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode TopPeerCategoryClass: %w", bin.NewUnexpectedID(id))
	}
}

// TopPeerCategory boxes the TopPeerCategoryClass providing a helper.
type TopPeerCategoryBox struct {
	TopPeerCategory TopPeerCategoryClass
}

// Decode implements bin.Decoder for TopPeerCategoryBox.
func (b *TopPeerCategoryBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode TopPeerCategoryBox to nil")
	}
	v, err := DecodeTopPeerCategory(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.TopPeerCategory = v
	return nil
}

// Encode implements bin.Encode for TopPeerCategoryBox.
func (b *TopPeerCategoryBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.TopPeerCategory == nil {
		return fmt.Errorf("unable to encode TopPeerCategoryClass as nil")
	}
	return b.TopPeerCategory.Encode(buf)
}

// TopPeerCategoryClassSlice is adapter for slice of TopPeerCategoryClass.
type TopPeerCategoryClassSlice []TopPeerCategoryClass

// First returns first element of slice (if exists).
func (s TopPeerCategoryClassSlice) First() (v TopPeerCategoryClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s TopPeerCategoryClassSlice) Last() (v TopPeerCategoryClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *TopPeerCategoryClassSlice) PopFirst() (v TopPeerCategoryClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	a[len(a)-1] = nil
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *TopPeerCategoryClassSlice) Pop() (v TopPeerCategoryClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
