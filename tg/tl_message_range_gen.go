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

// MessageRange represents TL type `messageRange#ae30253`.
// Indicates a range of chat messages
//
// See https://core.telegram.org/constructor/messageRange for reference.
type MessageRange struct {
	// Start of range (message ID)
	MinID int `schemaname:"min_id"`
	// End of range (message ID)
	MaxID int `schemaname:"max_id"`
}

// MessageRangeTypeID is TL type id of MessageRange.
const MessageRangeTypeID = 0xae30253

func (m *MessageRange) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.MinID == 0) {
		return false
	}
	if !(m.MaxID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageRange) String() string {
	if m == nil {
		return "MessageRange(nil)"
	}
	type Alias MessageRange
	return fmt.Sprintf("MessageRange%+v", Alias(*m))
}

// FillFrom fills MessageRange from given interface.
func (m *MessageRange) FillFrom(from interface {
	GetMinID() (value int)
	GetMaxID() (value int)
}) {
	m.MinID = from.GetMinID()
	m.MaxID = from.GetMaxID()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (m *MessageRange) TypeID() uint32 {
	return MessageRangeTypeID
}

// SchemaName returns MTProto type name.
func (m *MessageRange) SchemaName() string {
	return "messageRange"
}

// Encode implements bin.Encoder.
func (m *MessageRange) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageRange#ae30253 as nil")
	}
	b.PutID(MessageRangeTypeID)
	b.PutInt(m.MinID)
	b.PutInt(m.MaxID)
	return nil
}

// GetMinID returns value of MinID field.
func (m *MessageRange) GetMinID() (value int) {
	return m.MinID
}

// GetMaxID returns value of MaxID field.
func (m *MessageRange) GetMaxID() (value int) {
	return m.MaxID
}

// Decode implements bin.Decoder.
func (m *MessageRange) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageRange#ae30253 to nil")
	}
	if err := b.ConsumeID(MessageRangeTypeID); err != nil {
		return fmt.Errorf("unable to decode messageRange#ae30253: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageRange#ae30253: field min_id: %w", err)
		}
		m.MinID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageRange#ae30253: field max_id: %w", err)
		}
		m.MaxID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessageRange.
var (
	_ bin.Encoder = &MessageRange{}
	_ bin.Decoder = &MessageRange{}
)
