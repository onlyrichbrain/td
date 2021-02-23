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

// InputPrivacyValueAllowContacts represents TL type `inputPrivacyValueAllowContacts#d09e07b`.
// Allow only contacts
//
// See https://core.telegram.org/constructor/inputPrivacyValueAllowContacts for reference.
type InputPrivacyValueAllowContacts struct {
}

// InputPrivacyValueAllowContactsTypeID is TL type id of InputPrivacyValueAllowContacts.
const InputPrivacyValueAllowContactsTypeID = 0xd09e07b

func (i *InputPrivacyValueAllowContacts) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyValueAllowContacts) String() string {
	if i == nil {
		return "InputPrivacyValueAllowContacts(nil)"
	}
	type Alias InputPrivacyValueAllowContacts
	return fmt.Sprintf("InputPrivacyValueAllowContacts%+v", Alias(*i))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPrivacyValueAllowContacts) TypeID() uint32 {
	return InputPrivacyValueAllowContactsTypeID
}

// SchemaName returns MTProto type name.
func (i *InputPrivacyValueAllowContacts) SchemaName() string {
	return "inputPrivacyValueAllowContacts"
}

// Encode implements bin.Encoder.
func (i *InputPrivacyValueAllowContacts) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueAllowContacts#d09e07b as nil")
	}
	b.PutID(InputPrivacyValueAllowContactsTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyValueAllowContacts) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueAllowContacts#d09e07b to nil")
	}
	if err := b.ConsumeID(InputPrivacyValueAllowContactsTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyValueAllowContacts#d09e07b: %w", err)
	}
	return nil
}

// construct implements constructor of InputPrivacyRuleClass.
func (i InputPrivacyValueAllowContacts) construct() InputPrivacyRuleClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyValueAllowContacts.
var (
	_ bin.Encoder = &InputPrivacyValueAllowContacts{}
	_ bin.Decoder = &InputPrivacyValueAllowContacts{}

	_ InputPrivacyRuleClass = &InputPrivacyValueAllowContacts{}
)

// InputPrivacyValueAllowAll represents TL type `inputPrivacyValueAllowAll#184b35ce`.
// Allow all users
//
// See https://core.telegram.org/constructor/inputPrivacyValueAllowAll for reference.
type InputPrivacyValueAllowAll struct {
}

// InputPrivacyValueAllowAllTypeID is TL type id of InputPrivacyValueAllowAll.
const InputPrivacyValueAllowAllTypeID = 0x184b35ce

func (i *InputPrivacyValueAllowAll) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyValueAllowAll) String() string {
	if i == nil {
		return "InputPrivacyValueAllowAll(nil)"
	}
	type Alias InputPrivacyValueAllowAll
	return fmt.Sprintf("InputPrivacyValueAllowAll%+v", Alias(*i))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPrivacyValueAllowAll) TypeID() uint32 {
	return InputPrivacyValueAllowAllTypeID
}

// SchemaName returns MTProto type name.
func (i *InputPrivacyValueAllowAll) SchemaName() string {
	return "inputPrivacyValueAllowAll"
}

// Encode implements bin.Encoder.
func (i *InputPrivacyValueAllowAll) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueAllowAll#184b35ce as nil")
	}
	b.PutID(InputPrivacyValueAllowAllTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyValueAllowAll) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueAllowAll#184b35ce to nil")
	}
	if err := b.ConsumeID(InputPrivacyValueAllowAllTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyValueAllowAll#184b35ce: %w", err)
	}
	return nil
}

// construct implements constructor of InputPrivacyRuleClass.
func (i InputPrivacyValueAllowAll) construct() InputPrivacyRuleClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyValueAllowAll.
var (
	_ bin.Encoder = &InputPrivacyValueAllowAll{}
	_ bin.Decoder = &InputPrivacyValueAllowAll{}

	_ InputPrivacyRuleClass = &InputPrivacyValueAllowAll{}
)

// InputPrivacyValueAllowUsers represents TL type `inputPrivacyValueAllowUsers#131cc67f`.
// Allow only certain users
//
// See https://core.telegram.org/constructor/inputPrivacyValueAllowUsers for reference.
type InputPrivacyValueAllowUsers struct {
	// Allowed users
	Users []InputUserClass `schemaname:"users"`
}

// InputPrivacyValueAllowUsersTypeID is TL type id of InputPrivacyValueAllowUsers.
const InputPrivacyValueAllowUsersTypeID = 0x131cc67f

func (i *InputPrivacyValueAllowUsers) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyValueAllowUsers) String() string {
	if i == nil {
		return "InputPrivacyValueAllowUsers(nil)"
	}
	type Alias InputPrivacyValueAllowUsers
	return fmt.Sprintf("InputPrivacyValueAllowUsers%+v", Alias(*i))
}

// FillFrom fills InputPrivacyValueAllowUsers from given interface.
func (i *InputPrivacyValueAllowUsers) FillFrom(from interface {
	GetUsers() (value []InputUserClass)
}) {
	i.Users = from.GetUsers()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPrivacyValueAllowUsers) TypeID() uint32 {
	return InputPrivacyValueAllowUsersTypeID
}

// SchemaName returns MTProto type name.
func (i *InputPrivacyValueAllowUsers) SchemaName() string {
	return "inputPrivacyValueAllowUsers"
}

// Encode implements bin.Encoder.
func (i *InputPrivacyValueAllowUsers) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueAllowUsers#131cc67f as nil")
	}
	b.PutID(InputPrivacyValueAllowUsersTypeID)
	b.PutVectorHeader(len(i.Users))
	for idx, v := range i.Users {
		if v == nil {
			return fmt.Errorf("unable to encode inputPrivacyValueAllowUsers#131cc67f: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputPrivacyValueAllowUsers#131cc67f: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetUsers returns value of Users field.
func (i *InputPrivacyValueAllowUsers) GetUsers() (value []InputUserClass) {
	return i.Users
}

// MapUsers returns field Users wrapped in InputUserClassSlice helper.
func (i *InputPrivacyValueAllowUsers) MapUsers() (value InputUserClassSlice) {
	return InputUserClassSlice(i.Users)
}

// Decode implements bin.Decoder.
func (i *InputPrivacyValueAllowUsers) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueAllowUsers#131cc67f to nil")
	}
	if err := b.ConsumeID(InputPrivacyValueAllowUsersTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyValueAllowUsers#131cc67f: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode inputPrivacyValueAllowUsers#131cc67f: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputPrivacyValueAllowUsers#131cc67f: field users: %w", err)
			}
			i.Users = append(i.Users, value)
		}
	}
	return nil
}

// construct implements constructor of InputPrivacyRuleClass.
func (i InputPrivacyValueAllowUsers) construct() InputPrivacyRuleClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyValueAllowUsers.
var (
	_ bin.Encoder = &InputPrivacyValueAllowUsers{}
	_ bin.Decoder = &InputPrivacyValueAllowUsers{}

	_ InputPrivacyRuleClass = &InputPrivacyValueAllowUsers{}
)

// InputPrivacyValueDisallowContacts represents TL type `inputPrivacyValueDisallowContacts#ba52007`.
// Disallow only contacts
//
// See https://core.telegram.org/constructor/inputPrivacyValueDisallowContacts for reference.
type InputPrivacyValueDisallowContacts struct {
}

// InputPrivacyValueDisallowContactsTypeID is TL type id of InputPrivacyValueDisallowContacts.
const InputPrivacyValueDisallowContactsTypeID = 0xba52007

func (i *InputPrivacyValueDisallowContacts) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyValueDisallowContacts) String() string {
	if i == nil {
		return "InputPrivacyValueDisallowContacts(nil)"
	}
	type Alias InputPrivacyValueDisallowContacts
	return fmt.Sprintf("InputPrivacyValueDisallowContacts%+v", Alias(*i))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPrivacyValueDisallowContacts) TypeID() uint32 {
	return InputPrivacyValueDisallowContactsTypeID
}

// SchemaName returns MTProto type name.
func (i *InputPrivacyValueDisallowContacts) SchemaName() string {
	return "inputPrivacyValueDisallowContacts"
}

// Encode implements bin.Encoder.
func (i *InputPrivacyValueDisallowContacts) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueDisallowContacts#ba52007 as nil")
	}
	b.PutID(InputPrivacyValueDisallowContactsTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyValueDisallowContacts) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueDisallowContacts#ba52007 to nil")
	}
	if err := b.ConsumeID(InputPrivacyValueDisallowContactsTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyValueDisallowContacts#ba52007: %w", err)
	}
	return nil
}

// construct implements constructor of InputPrivacyRuleClass.
func (i InputPrivacyValueDisallowContacts) construct() InputPrivacyRuleClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyValueDisallowContacts.
var (
	_ bin.Encoder = &InputPrivacyValueDisallowContacts{}
	_ bin.Decoder = &InputPrivacyValueDisallowContacts{}

	_ InputPrivacyRuleClass = &InputPrivacyValueDisallowContacts{}
)

// InputPrivacyValueDisallowAll represents TL type `inputPrivacyValueDisallowAll#d66b66c9`.
// Disallow all
//
// See https://core.telegram.org/constructor/inputPrivacyValueDisallowAll for reference.
type InputPrivacyValueDisallowAll struct {
}

// InputPrivacyValueDisallowAllTypeID is TL type id of InputPrivacyValueDisallowAll.
const InputPrivacyValueDisallowAllTypeID = 0xd66b66c9

func (i *InputPrivacyValueDisallowAll) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyValueDisallowAll) String() string {
	if i == nil {
		return "InputPrivacyValueDisallowAll(nil)"
	}
	type Alias InputPrivacyValueDisallowAll
	return fmt.Sprintf("InputPrivacyValueDisallowAll%+v", Alias(*i))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPrivacyValueDisallowAll) TypeID() uint32 {
	return InputPrivacyValueDisallowAllTypeID
}

// SchemaName returns MTProto type name.
func (i *InputPrivacyValueDisallowAll) SchemaName() string {
	return "inputPrivacyValueDisallowAll"
}

// Encode implements bin.Encoder.
func (i *InputPrivacyValueDisallowAll) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueDisallowAll#d66b66c9 as nil")
	}
	b.PutID(InputPrivacyValueDisallowAllTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyValueDisallowAll) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueDisallowAll#d66b66c9 to nil")
	}
	if err := b.ConsumeID(InputPrivacyValueDisallowAllTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyValueDisallowAll#d66b66c9: %w", err)
	}
	return nil
}

// construct implements constructor of InputPrivacyRuleClass.
func (i InputPrivacyValueDisallowAll) construct() InputPrivacyRuleClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyValueDisallowAll.
var (
	_ bin.Encoder = &InputPrivacyValueDisallowAll{}
	_ bin.Decoder = &InputPrivacyValueDisallowAll{}

	_ InputPrivacyRuleClass = &InputPrivacyValueDisallowAll{}
)

// InputPrivacyValueDisallowUsers represents TL type `inputPrivacyValueDisallowUsers#90110467`.
// Disallow only certain users
//
// See https://core.telegram.org/constructor/inputPrivacyValueDisallowUsers for reference.
type InputPrivacyValueDisallowUsers struct {
	// Users to disallow
	Users []InputUserClass `schemaname:"users"`
}

// InputPrivacyValueDisallowUsersTypeID is TL type id of InputPrivacyValueDisallowUsers.
const InputPrivacyValueDisallowUsersTypeID = 0x90110467

func (i *InputPrivacyValueDisallowUsers) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyValueDisallowUsers) String() string {
	if i == nil {
		return "InputPrivacyValueDisallowUsers(nil)"
	}
	type Alias InputPrivacyValueDisallowUsers
	return fmt.Sprintf("InputPrivacyValueDisallowUsers%+v", Alias(*i))
}

// FillFrom fills InputPrivacyValueDisallowUsers from given interface.
func (i *InputPrivacyValueDisallowUsers) FillFrom(from interface {
	GetUsers() (value []InputUserClass)
}) {
	i.Users = from.GetUsers()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPrivacyValueDisallowUsers) TypeID() uint32 {
	return InputPrivacyValueDisallowUsersTypeID
}

// SchemaName returns MTProto type name.
func (i *InputPrivacyValueDisallowUsers) SchemaName() string {
	return "inputPrivacyValueDisallowUsers"
}

// Encode implements bin.Encoder.
func (i *InputPrivacyValueDisallowUsers) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueDisallowUsers#90110467 as nil")
	}
	b.PutID(InputPrivacyValueDisallowUsersTypeID)
	b.PutVectorHeader(len(i.Users))
	for idx, v := range i.Users {
		if v == nil {
			return fmt.Errorf("unable to encode inputPrivacyValueDisallowUsers#90110467: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputPrivacyValueDisallowUsers#90110467: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetUsers returns value of Users field.
func (i *InputPrivacyValueDisallowUsers) GetUsers() (value []InputUserClass) {
	return i.Users
}

// MapUsers returns field Users wrapped in InputUserClassSlice helper.
func (i *InputPrivacyValueDisallowUsers) MapUsers() (value InputUserClassSlice) {
	return InputUserClassSlice(i.Users)
}

// Decode implements bin.Decoder.
func (i *InputPrivacyValueDisallowUsers) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueDisallowUsers#90110467 to nil")
	}
	if err := b.ConsumeID(InputPrivacyValueDisallowUsersTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyValueDisallowUsers#90110467: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode inputPrivacyValueDisallowUsers#90110467: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputPrivacyValueDisallowUsers#90110467: field users: %w", err)
			}
			i.Users = append(i.Users, value)
		}
	}
	return nil
}

// construct implements constructor of InputPrivacyRuleClass.
func (i InputPrivacyValueDisallowUsers) construct() InputPrivacyRuleClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyValueDisallowUsers.
var (
	_ bin.Encoder = &InputPrivacyValueDisallowUsers{}
	_ bin.Decoder = &InputPrivacyValueDisallowUsers{}

	_ InputPrivacyRuleClass = &InputPrivacyValueDisallowUsers{}
)

// InputPrivacyValueAllowChatParticipants represents TL type `inputPrivacyValueAllowChatParticipants#4c81c1ba`.
// Allow only participants of certain chats
//
// See https://core.telegram.org/constructor/inputPrivacyValueAllowChatParticipants for reference.
type InputPrivacyValueAllowChatParticipants struct {
	// Allowed chat IDs
	Chats []int `schemaname:"chats"`
}

// InputPrivacyValueAllowChatParticipantsTypeID is TL type id of InputPrivacyValueAllowChatParticipants.
const InputPrivacyValueAllowChatParticipantsTypeID = 0x4c81c1ba

func (i *InputPrivacyValueAllowChatParticipants) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Chats == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyValueAllowChatParticipants) String() string {
	if i == nil {
		return "InputPrivacyValueAllowChatParticipants(nil)"
	}
	type Alias InputPrivacyValueAllowChatParticipants
	return fmt.Sprintf("InputPrivacyValueAllowChatParticipants%+v", Alias(*i))
}

// FillFrom fills InputPrivacyValueAllowChatParticipants from given interface.
func (i *InputPrivacyValueAllowChatParticipants) FillFrom(from interface {
	GetChats() (value []int)
}) {
	i.Chats = from.GetChats()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPrivacyValueAllowChatParticipants) TypeID() uint32 {
	return InputPrivacyValueAllowChatParticipantsTypeID
}

// SchemaName returns MTProto type name.
func (i *InputPrivacyValueAllowChatParticipants) SchemaName() string {
	return "inputPrivacyValueAllowChatParticipants"
}

// Encode implements bin.Encoder.
func (i *InputPrivacyValueAllowChatParticipants) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueAllowChatParticipants#4c81c1ba as nil")
	}
	b.PutID(InputPrivacyValueAllowChatParticipantsTypeID)
	b.PutVectorHeader(len(i.Chats))
	for _, v := range i.Chats {
		b.PutInt(v)
	}
	return nil
}

// GetChats returns value of Chats field.
func (i *InputPrivacyValueAllowChatParticipants) GetChats() (value []int) {
	return i.Chats
}

// Decode implements bin.Decoder.
func (i *InputPrivacyValueAllowChatParticipants) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueAllowChatParticipants#4c81c1ba to nil")
	}
	if err := b.ConsumeID(InputPrivacyValueAllowChatParticipantsTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyValueAllowChatParticipants#4c81c1ba: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode inputPrivacyValueAllowChatParticipants#4c81c1ba: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode inputPrivacyValueAllowChatParticipants#4c81c1ba: field chats: %w", err)
			}
			i.Chats = append(i.Chats, value)
		}
	}
	return nil
}

// construct implements constructor of InputPrivacyRuleClass.
func (i InputPrivacyValueAllowChatParticipants) construct() InputPrivacyRuleClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyValueAllowChatParticipants.
var (
	_ bin.Encoder = &InputPrivacyValueAllowChatParticipants{}
	_ bin.Decoder = &InputPrivacyValueAllowChatParticipants{}

	_ InputPrivacyRuleClass = &InputPrivacyValueAllowChatParticipants{}
)

// InputPrivacyValueDisallowChatParticipants represents TL type `inputPrivacyValueDisallowChatParticipants#d82363af`.
// Disallow only participants of certain chats
//
// See https://core.telegram.org/constructor/inputPrivacyValueDisallowChatParticipants for reference.
type InputPrivacyValueDisallowChatParticipants struct {
	// Disallowed chat IDs
	Chats []int `schemaname:"chats"`
}

// InputPrivacyValueDisallowChatParticipantsTypeID is TL type id of InputPrivacyValueDisallowChatParticipants.
const InputPrivacyValueDisallowChatParticipantsTypeID = 0xd82363af

func (i *InputPrivacyValueDisallowChatParticipants) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Chats == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyValueDisallowChatParticipants) String() string {
	if i == nil {
		return "InputPrivacyValueDisallowChatParticipants(nil)"
	}
	type Alias InputPrivacyValueDisallowChatParticipants
	return fmt.Sprintf("InputPrivacyValueDisallowChatParticipants%+v", Alias(*i))
}

// FillFrom fills InputPrivacyValueDisallowChatParticipants from given interface.
func (i *InputPrivacyValueDisallowChatParticipants) FillFrom(from interface {
	GetChats() (value []int)
}) {
	i.Chats = from.GetChats()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPrivacyValueDisallowChatParticipants) TypeID() uint32 {
	return InputPrivacyValueDisallowChatParticipantsTypeID
}

// SchemaName returns MTProto type name.
func (i *InputPrivacyValueDisallowChatParticipants) SchemaName() string {
	return "inputPrivacyValueDisallowChatParticipants"
}

// Encode implements bin.Encoder.
func (i *InputPrivacyValueDisallowChatParticipants) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueDisallowChatParticipants#d82363af as nil")
	}
	b.PutID(InputPrivacyValueDisallowChatParticipantsTypeID)
	b.PutVectorHeader(len(i.Chats))
	for _, v := range i.Chats {
		b.PutInt(v)
	}
	return nil
}

// GetChats returns value of Chats field.
func (i *InputPrivacyValueDisallowChatParticipants) GetChats() (value []int) {
	return i.Chats
}

// Decode implements bin.Decoder.
func (i *InputPrivacyValueDisallowChatParticipants) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueDisallowChatParticipants#d82363af to nil")
	}
	if err := b.ConsumeID(InputPrivacyValueDisallowChatParticipantsTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyValueDisallowChatParticipants#d82363af: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode inputPrivacyValueDisallowChatParticipants#d82363af: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode inputPrivacyValueDisallowChatParticipants#d82363af: field chats: %w", err)
			}
			i.Chats = append(i.Chats, value)
		}
	}
	return nil
}

// construct implements constructor of InputPrivacyRuleClass.
func (i InputPrivacyValueDisallowChatParticipants) construct() InputPrivacyRuleClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyValueDisallowChatParticipants.
var (
	_ bin.Encoder = &InputPrivacyValueDisallowChatParticipants{}
	_ bin.Decoder = &InputPrivacyValueDisallowChatParticipants{}

	_ InputPrivacyRuleClass = &InputPrivacyValueDisallowChatParticipants{}
)

// InputPrivacyRuleClass represents InputPrivacyRule generic type.
//
// See https://core.telegram.org/type/InputPrivacyRule for reference.
//
// Example:
//  g, err := tg.DecodeInputPrivacyRule(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.InputPrivacyValueAllowContacts: // inputPrivacyValueAllowContacts#d09e07b
//  case *tg.InputPrivacyValueAllowAll: // inputPrivacyValueAllowAll#184b35ce
//  case *tg.InputPrivacyValueAllowUsers: // inputPrivacyValueAllowUsers#131cc67f
//  case *tg.InputPrivacyValueDisallowContacts: // inputPrivacyValueDisallowContacts#ba52007
//  case *tg.InputPrivacyValueDisallowAll: // inputPrivacyValueDisallowAll#d66b66c9
//  case *tg.InputPrivacyValueDisallowUsers: // inputPrivacyValueDisallowUsers#90110467
//  case *tg.InputPrivacyValueAllowChatParticipants: // inputPrivacyValueAllowChatParticipants#4c81c1ba
//  case *tg.InputPrivacyValueDisallowChatParticipants: // inputPrivacyValueDisallowChatParticipants#d82363af
//  default: panic(v)
//  }
type InputPrivacyRuleClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputPrivacyRuleClass

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

// DecodeInputPrivacyRule implements binary de-serialization for InputPrivacyRuleClass.
func DecodeInputPrivacyRule(buf *bin.Buffer) (InputPrivacyRuleClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputPrivacyValueAllowContactsTypeID:
		// Decoding inputPrivacyValueAllowContacts#d09e07b.
		v := InputPrivacyValueAllowContacts{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyRuleClass: %w", err)
		}
		return &v, nil
	case InputPrivacyValueAllowAllTypeID:
		// Decoding inputPrivacyValueAllowAll#184b35ce.
		v := InputPrivacyValueAllowAll{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyRuleClass: %w", err)
		}
		return &v, nil
	case InputPrivacyValueAllowUsersTypeID:
		// Decoding inputPrivacyValueAllowUsers#131cc67f.
		v := InputPrivacyValueAllowUsers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyRuleClass: %w", err)
		}
		return &v, nil
	case InputPrivacyValueDisallowContactsTypeID:
		// Decoding inputPrivacyValueDisallowContacts#ba52007.
		v := InputPrivacyValueDisallowContacts{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyRuleClass: %w", err)
		}
		return &v, nil
	case InputPrivacyValueDisallowAllTypeID:
		// Decoding inputPrivacyValueDisallowAll#d66b66c9.
		v := InputPrivacyValueDisallowAll{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyRuleClass: %w", err)
		}
		return &v, nil
	case InputPrivacyValueDisallowUsersTypeID:
		// Decoding inputPrivacyValueDisallowUsers#90110467.
		v := InputPrivacyValueDisallowUsers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyRuleClass: %w", err)
		}
		return &v, nil
	case InputPrivacyValueAllowChatParticipantsTypeID:
		// Decoding inputPrivacyValueAllowChatParticipants#4c81c1ba.
		v := InputPrivacyValueAllowChatParticipants{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyRuleClass: %w", err)
		}
		return &v, nil
	case InputPrivacyValueDisallowChatParticipantsTypeID:
		// Decoding inputPrivacyValueDisallowChatParticipants#d82363af.
		v := InputPrivacyValueDisallowChatParticipants{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyRuleClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputPrivacyRuleClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputPrivacyRule boxes the InputPrivacyRuleClass providing a helper.
type InputPrivacyRuleBox struct {
	InputPrivacyRule InputPrivacyRuleClass
}

// Decode implements bin.Decoder for InputPrivacyRuleBox.
func (b *InputPrivacyRuleBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputPrivacyRuleBox to nil")
	}
	v, err := DecodeInputPrivacyRule(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputPrivacyRule = v
	return nil
}

// Encode implements bin.Encode for InputPrivacyRuleBox.
func (b *InputPrivacyRuleBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputPrivacyRule == nil {
		return fmt.Errorf("unable to encode InputPrivacyRuleClass as nil")
	}
	return b.InputPrivacyRule.Encode(buf)
}

// InputPrivacyRuleClassSlice is adapter for slice of InputPrivacyRuleClass.
type InputPrivacyRuleClassSlice []InputPrivacyRuleClass

// First returns first element of slice (if exists).
func (s InputPrivacyRuleClassSlice) First() (v InputPrivacyRuleClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputPrivacyRuleClassSlice) Last() (v InputPrivacyRuleClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputPrivacyRuleClassSlice) PopFirst() (v InputPrivacyRuleClass, ok bool) {
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
func (s *InputPrivacyRuleClassSlice) Pop() (v InputPrivacyRuleClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
