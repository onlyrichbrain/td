// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// InputBusinessChatLink represents TL type `inputBusinessChatLink#11679fa7`.
// Contains info about a business chat deep link »¹ to be created by the current
// account.
//
// Links:
//  1. https://core.telegram.org/api/business#business-chat-links
//
// See https://core.telegram.org/constructor/inputBusinessChatLink for reference.
type InputBusinessChatLink struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Message to pre-fill in the message input field.
	Message string
	// Message entities for styled text¹
	//
	// Links:
	//  1) https://core.telegram.org/api/entities
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
	// Human-readable name of the link, to simplify management in the UI (only visible to the
	// creator of the link).
	//
	// Use SetTitle and GetTitle helpers.
	Title string
}

// InputBusinessChatLinkTypeID is TL type id of InputBusinessChatLink.
const InputBusinessChatLinkTypeID = 0x11679fa7

// Ensuring interfaces in compile-time for InputBusinessChatLink.
var (
	_ bin.Encoder     = &InputBusinessChatLink{}
	_ bin.Decoder     = &InputBusinessChatLink{}
	_ bin.BareEncoder = &InputBusinessChatLink{}
	_ bin.BareDecoder = &InputBusinessChatLink{}
)

func (i *InputBusinessChatLink) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Flags.Zero()) {
		return false
	}
	if !(i.Message == "") {
		return false
	}
	if !(i.Entities == nil) {
		return false
	}
	if !(i.Title == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputBusinessChatLink) String() string {
	if i == nil {
		return "InputBusinessChatLink(nil)"
	}
	type Alias InputBusinessChatLink
	return fmt.Sprintf("InputBusinessChatLink%+v", Alias(*i))
}

// FillFrom fills InputBusinessChatLink from given interface.
func (i *InputBusinessChatLink) FillFrom(from interface {
	GetMessage() (value string)
	GetEntities() (value []MessageEntityClass, ok bool)
	GetTitle() (value string, ok bool)
}) {
	i.Message = from.GetMessage()
	if val, ok := from.GetEntities(); ok {
		i.Entities = val
	}

	if val, ok := from.GetTitle(); ok {
		i.Title = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputBusinessChatLink) TypeID() uint32 {
	return InputBusinessChatLinkTypeID
}

// TypeName returns name of type in TL schema.
func (*InputBusinessChatLink) TypeName() string {
	return "inputBusinessChatLink"
}

// TypeInfo returns info about TL type.
func (i *InputBusinessChatLink) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputBusinessChatLink",
		ID:   InputBusinessChatLinkTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Message",
			SchemaName: "message",
		},
		{
			Name:       "Entities",
			SchemaName: "entities",
			Null:       !i.Flags.Has(0),
		},
		{
			Name:       "Title",
			SchemaName: "title",
			Null:       !i.Flags.Has(1),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (i *InputBusinessChatLink) SetFlags() {
	if !(i.Entities == nil) {
		i.Flags.Set(0)
	}
	if !(i.Title == "") {
		i.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (i *InputBusinessChatLink) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputBusinessChatLink#11679fa7 as nil")
	}
	b.PutID(InputBusinessChatLinkTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputBusinessChatLink) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputBusinessChatLink#11679fa7 as nil")
	}
	i.SetFlags()
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputBusinessChatLink#11679fa7: field flags: %w", err)
	}
	b.PutString(i.Message)
	if i.Flags.Has(0) {
		b.PutVectorHeader(len(i.Entities))
		for idx, v := range i.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode inputBusinessChatLink#11679fa7: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode inputBusinessChatLink#11679fa7: field entities element with index %d: %w", idx, err)
			}
		}
	}
	if i.Flags.Has(1) {
		b.PutString(i.Title)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputBusinessChatLink) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputBusinessChatLink#11679fa7 to nil")
	}
	if err := b.ConsumeID(InputBusinessChatLinkTypeID); err != nil {
		return fmt.Errorf("unable to decode inputBusinessChatLink#11679fa7: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputBusinessChatLink) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputBusinessChatLink#11679fa7 to nil")
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputBusinessChatLink#11679fa7: field flags: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputBusinessChatLink#11679fa7: field message: %w", err)
		}
		i.Message = value
	}
	if i.Flags.Has(0) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode inputBusinessChatLink#11679fa7: field entities: %w", err)
		}

		if headerLen > 0 {
			i.Entities = make([]MessageEntityClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputBusinessChatLink#11679fa7: field entities: %w", err)
			}
			i.Entities = append(i.Entities, value)
		}
	}
	if i.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputBusinessChatLink#11679fa7: field title: %w", err)
		}
		i.Title = value
	}
	return nil
}

// GetMessage returns value of Message field.
func (i *InputBusinessChatLink) GetMessage() (value string) {
	if i == nil {
		return
	}
	return i.Message
}

// SetEntities sets value of Entities conditional field.
func (i *InputBusinessChatLink) SetEntities(value []MessageEntityClass) {
	i.Flags.Set(0)
	i.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (i *InputBusinessChatLink) GetEntities() (value []MessageEntityClass, ok bool) {
	if i == nil {
		return
	}
	if !i.Flags.Has(0) {
		return value, false
	}
	return i.Entities, true
}

// SetTitle sets value of Title conditional field.
func (i *InputBusinessChatLink) SetTitle(value string) {
	i.Flags.Set(1)
	i.Title = value
}

// GetTitle returns value of Title conditional field and
// boolean which is true if field was set.
func (i *InputBusinessChatLink) GetTitle() (value string, ok bool) {
	if i == nil {
		return
	}
	if !i.Flags.Has(1) {
		return value, false
	}
	return i.Title, true
}

// MapEntities returns field Entities wrapped in MessageEntityClassArray helper.
func (i *InputBusinessChatLink) MapEntities() (value MessageEntityClassArray, ok bool) {
	if !i.Flags.Has(0) {
		return value, false
	}
	return MessageEntityClassArray(i.Entities), true
}
