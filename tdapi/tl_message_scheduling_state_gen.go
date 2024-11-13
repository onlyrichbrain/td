// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// MessageSchedulingStateSendAtDate represents TL type `messageSchedulingStateSendAtDate#a773ffe7`.
type MessageSchedulingStateSendAtDate struct {
	// Point in time (Unix timestamp) when the message will be sent. The date must be within
	// 367 days in the future
	SendDate int32
}

// MessageSchedulingStateSendAtDateTypeID is TL type id of MessageSchedulingStateSendAtDate.
const MessageSchedulingStateSendAtDateTypeID = 0xa773ffe7

// construct implements constructor of MessageSchedulingStateClass.
func (m MessageSchedulingStateSendAtDate) construct() MessageSchedulingStateClass { return &m }

// Ensuring interfaces in compile-time for MessageSchedulingStateSendAtDate.
var (
	_ bin.Encoder     = &MessageSchedulingStateSendAtDate{}
	_ bin.Decoder     = &MessageSchedulingStateSendAtDate{}
	_ bin.BareEncoder = &MessageSchedulingStateSendAtDate{}
	_ bin.BareDecoder = &MessageSchedulingStateSendAtDate{}

	_ MessageSchedulingStateClass = &MessageSchedulingStateSendAtDate{}
)

func (m *MessageSchedulingStateSendAtDate) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.SendDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageSchedulingStateSendAtDate) String() string {
	if m == nil {
		return "MessageSchedulingStateSendAtDate(nil)"
	}
	type Alias MessageSchedulingStateSendAtDate
	return fmt.Sprintf("MessageSchedulingStateSendAtDate%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageSchedulingStateSendAtDate) TypeID() uint32 {
	return MessageSchedulingStateSendAtDateTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageSchedulingStateSendAtDate) TypeName() string {
	return "messageSchedulingStateSendAtDate"
}

// TypeInfo returns info about TL type.
func (m *MessageSchedulingStateSendAtDate) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageSchedulingStateSendAtDate",
		ID:   MessageSchedulingStateSendAtDateTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SendDate",
			SchemaName: "send_date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageSchedulingStateSendAtDate) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSchedulingStateSendAtDate#a773ffe7 as nil")
	}
	b.PutID(MessageSchedulingStateSendAtDateTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageSchedulingStateSendAtDate) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSchedulingStateSendAtDate#a773ffe7 as nil")
	}
	b.PutInt32(m.SendDate)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageSchedulingStateSendAtDate) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSchedulingStateSendAtDate#a773ffe7 to nil")
	}
	if err := b.ConsumeID(MessageSchedulingStateSendAtDateTypeID); err != nil {
		return fmt.Errorf("unable to decode messageSchedulingStateSendAtDate#a773ffe7: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageSchedulingStateSendAtDate) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSchedulingStateSendAtDate#a773ffe7 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode messageSchedulingStateSendAtDate#a773ffe7: field send_date: %w", err)
		}
		m.SendDate = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (m *MessageSchedulingStateSendAtDate) EncodeTDLibJSON(b tdjson.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSchedulingStateSendAtDate#a773ffe7 as nil")
	}
	b.ObjStart()
	b.PutID("messageSchedulingStateSendAtDate")
	b.Comma()
	b.FieldStart("send_date")
	b.PutInt32(m.SendDate)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (m *MessageSchedulingStateSendAtDate) DecodeTDLibJSON(b tdjson.Decoder) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSchedulingStateSendAtDate#a773ffe7 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("messageSchedulingStateSendAtDate"); err != nil {
				return fmt.Errorf("unable to decode messageSchedulingStateSendAtDate#a773ffe7: %w", err)
			}
		case "send_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode messageSchedulingStateSendAtDate#a773ffe7: field send_date: %w", err)
			}
			m.SendDate = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSendDate returns value of SendDate field.
func (m *MessageSchedulingStateSendAtDate) GetSendDate() (value int32) {
	if m == nil {
		return
	}
	return m.SendDate
}

// MessageSchedulingStateSendWhenOnline represents TL type `messageSchedulingStateSendWhenOnline#7cbfd808`.
type MessageSchedulingStateSendWhenOnline struct {
}

// MessageSchedulingStateSendWhenOnlineTypeID is TL type id of MessageSchedulingStateSendWhenOnline.
const MessageSchedulingStateSendWhenOnlineTypeID = 0x7cbfd808

// construct implements constructor of MessageSchedulingStateClass.
func (m MessageSchedulingStateSendWhenOnline) construct() MessageSchedulingStateClass { return &m }

// Ensuring interfaces in compile-time for MessageSchedulingStateSendWhenOnline.
var (
	_ bin.Encoder     = &MessageSchedulingStateSendWhenOnline{}
	_ bin.Decoder     = &MessageSchedulingStateSendWhenOnline{}
	_ bin.BareEncoder = &MessageSchedulingStateSendWhenOnline{}
	_ bin.BareDecoder = &MessageSchedulingStateSendWhenOnline{}

	_ MessageSchedulingStateClass = &MessageSchedulingStateSendWhenOnline{}
)

func (m *MessageSchedulingStateSendWhenOnline) Zero() bool {
	if m == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageSchedulingStateSendWhenOnline) String() string {
	if m == nil {
		return "MessageSchedulingStateSendWhenOnline(nil)"
	}
	type Alias MessageSchedulingStateSendWhenOnline
	return fmt.Sprintf("MessageSchedulingStateSendWhenOnline%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageSchedulingStateSendWhenOnline) TypeID() uint32 {
	return MessageSchedulingStateSendWhenOnlineTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageSchedulingStateSendWhenOnline) TypeName() string {
	return "messageSchedulingStateSendWhenOnline"
}

// TypeInfo returns info about TL type.
func (m *MessageSchedulingStateSendWhenOnline) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageSchedulingStateSendWhenOnline",
		ID:   MessageSchedulingStateSendWhenOnlineTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageSchedulingStateSendWhenOnline) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSchedulingStateSendWhenOnline#7cbfd808 as nil")
	}
	b.PutID(MessageSchedulingStateSendWhenOnlineTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageSchedulingStateSendWhenOnline) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSchedulingStateSendWhenOnline#7cbfd808 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageSchedulingStateSendWhenOnline) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSchedulingStateSendWhenOnline#7cbfd808 to nil")
	}
	if err := b.ConsumeID(MessageSchedulingStateSendWhenOnlineTypeID); err != nil {
		return fmt.Errorf("unable to decode messageSchedulingStateSendWhenOnline#7cbfd808: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageSchedulingStateSendWhenOnline) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSchedulingStateSendWhenOnline#7cbfd808 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (m *MessageSchedulingStateSendWhenOnline) EncodeTDLibJSON(b tdjson.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSchedulingStateSendWhenOnline#7cbfd808 as nil")
	}
	b.ObjStart()
	b.PutID("messageSchedulingStateSendWhenOnline")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (m *MessageSchedulingStateSendWhenOnline) DecodeTDLibJSON(b tdjson.Decoder) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSchedulingStateSendWhenOnline#7cbfd808 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("messageSchedulingStateSendWhenOnline"); err != nil {
				return fmt.Errorf("unable to decode messageSchedulingStateSendWhenOnline#7cbfd808: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// MessageSchedulingStateSendWhenVideoProcessed represents TL type `messageSchedulingStateSendWhenVideoProcessed#7d438bee`.
type MessageSchedulingStateSendWhenVideoProcessed struct {
	// Approximate point in time (Unix timestamp) when the message is expected to be sent
	SendDate int32
}

// MessageSchedulingStateSendWhenVideoProcessedTypeID is TL type id of MessageSchedulingStateSendWhenVideoProcessed.
const MessageSchedulingStateSendWhenVideoProcessedTypeID = 0x7d438bee

// construct implements constructor of MessageSchedulingStateClass.
func (m MessageSchedulingStateSendWhenVideoProcessed) construct() MessageSchedulingStateClass {
	return &m
}

// Ensuring interfaces in compile-time for MessageSchedulingStateSendWhenVideoProcessed.
var (
	_ bin.Encoder     = &MessageSchedulingStateSendWhenVideoProcessed{}
	_ bin.Decoder     = &MessageSchedulingStateSendWhenVideoProcessed{}
	_ bin.BareEncoder = &MessageSchedulingStateSendWhenVideoProcessed{}
	_ bin.BareDecoder = &MessageSchedulingStateSendWhenVideoProcessed{}

	_ MessageSchedulingStateClass = &MessageSchedulingStateSendWhenVideoProcessed{}
)

func (m *MessageSchedulingStateSendWhenVideoProcessed) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.SendDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageSchedulingStateSendWhenVideoProcessed) String() string {
	if m == nil {
		return "MessageSchedulingStateSendWhenVideoProcessed(nil)"
	}
	type Alias MessageSchedulingStateSendWhenVideoProcessed
	return fmt.Sprintf("MessageSchedulingStateSendWhenVideoProcessed%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageSchedulingStateSendWhenVideoProcessed) TypeID() uint32 {
	return MessageSchedulingStateSendWhenVideoProcessedTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageSchedulingStateSendWhenVideoProcessed) TypeName() string {
	return "messageSchedulingStateSendWhenVideoProcessed"
}

// TypeInfo returns info about TL type.
func (m *MessageSchedulingStateSendWhenVideoProcessed) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageSchedulingStateSendWhenVideoProcessed",
		ID:   MessageSchedulingStateSendWhenVideoProcessedTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SendDate",
			SchemaName: "send_date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageSchedulingStateSendWhenVideoProcessed) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSchedulingStateSendWhenVideoProcessed#7d438bee as nil")
	}
	b.PutID(MessageSchedulingStateSendWhenVideoProcessedTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageSchedulingStateSendWhenVideoProcessed) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSchedulingStateSendWhenVideoProcessed#7d438bee as nil")
	}
	b.PutInt32(m.SendDate)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageSchedulingStateSendWhenVideoProcessed) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSchedulingStateSendWhenVideoProcessed#7d438bee to nil")
	}
	if err := b.ConsumeID(MessageSchedulingStateSendWhenVideoProcessedTypeID); err != nil {
		return fmt.Errorf("unable to decode messageSchedulingStateSendWhenVideoProcessed#7d438bee: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageSchedulingStateSendWhenVideoProcessed) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSchedulingStateSendWhenVideoProcessed#7d438bee to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode messageSchedulingStateSendWhenVideoProcessed#7d438bee: field send_date: %w", err)
		}
		m.SendDate = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (m *MessageSchedulingStateSendWhenVideoProcessed) EncodeTDLibJSON(b tdjson.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSchedulingStateSendWhenVideoProcessed#7d438bee as nil")
	}
	b.ObjStart()
	b.PutID("messageSchedulingStateSendWhenVideoProcessed")
	b.Comma()
	b.FieldStart("send_date")
	b.PutInt32(m.SendDate)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (m *MessageSchedulingStateSendWhenVideoProcessed) DecodeTDLibJSON(b tdjson.Decoder) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSchedulingStateSendWhenVideoProcessed#7d438bee to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("messageSchedulingStateSendWhenVideoProcessed"); err != nil {
				return fmt.Errorf("unable to decode messageSchedulingStateSendWhenVideoProcessed#7d438bee: %w", err)
			}
		case "send_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode messageSchedulingStateSendWhenVideoProcessed#7d438bee: field send_date: %w", err)
			}
			m.SendDate = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSendDate returns value of SendDate field.
func (m *MessageSchedulingStateSendWhenVideoProcessed) GetSendDate() (value int32) {
	if m == nil {
		return
	}
	return m.SendDate
}

// MessageSchedulingStateClassName is schema name of MessageSchedulingStateClass.
const MessageSchedulingStateClassName = "MessageSchedulingState"

// MessageSchedulingStateClass represents MessageSchedulingState generic type.
//
// Example:
//
//	g, err := tdapi.DecodeMessageSchedulingState(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.MessageSchedulingStateSendAtDate: // messageSchedulingStateSendAtDate#a773ffe7
//	case *tdapi.MessageSchedulingStateSendWhenOnline: // messageSchedulingStateSendWhenOnline#7cbfd808
//	case *tdapi.MessageSchedulingStateSendWhenVideoProcessed: // messageSchedulingStateSendWhenVideoProcessed#7d438bee
//	default: panic(v)
//	}
type MessageSchedulingStateClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessageSchedulingStateClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeMessageSchedulingState implements binary de-serialization for MessageSchedulingStateClass.
func DecodeMessageSchedulingState(buf *bin.Buffer) (MessageSchedulingStateClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessageSchedulingStateSendAtDateTypeID:
		// Decoding messageSchedulingStateSendAtDate#a773ffe7.
		v := MessageSchedulingStateSendAtDate{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageSchedulingStateClass: %w", err)
		}
		return &v, nil
	case MessageSchedulingStateSendWhenOnlineTypeID:
		// Decoding messageSchedulingStateSendWhenOnline#7cbfd808.
		v := MessageSchedulingStateSendWhenOnline{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageSchedulingStateClass: %w", err)
		}
		return &v, nil
	case MessageSchedulingStateSendWhenVideoProcessedTypeID:
		// Decoding messageSchedulingStateSendWhenVideoProcessed#7d438bee.
		v := MessageSchedulingStateSendWhenVideoProcessed{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageSchedulingStateClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessageSchedulingStateClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONMessageSchedulingState implements binary de-serialization for MessageSchedulingStateClass.
func DecodeTDLibJSONMessageSchedulingState(buf tdjson.Decoder) (MessageSchedulingStateClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "messageSchedulingStateSendAtDate":
		// Decoding messageSchedulingStateSendAtDate#a773ffe7.
		v := MessageSchedulingStateSendAtDate{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageSchedulingStateClass: %w", err)
		}
		return &v, nil
	case "messageSchedulingStateSendWhenOnline":
		// Decoding messageSchedulingStateSendWhenOnline#7cbfd808.
		v := MessageSchedulingStateSendWhenOnline{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageSchedulingStateClass: %w", err)
		}
		return &v, nil
	case "messageSchedulingStateSendWhenVideoProcessed":
		// Decoding messageSchedulingStateSendWhenVideoProcessed#7d438bee.
		v := MessageSchedulingStateSendWhenVideoProcessed{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageSchedulingStateClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessageSchedulingStateClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// MessageSchedulingState boxes the MessageSchedulingStateClass providing a helper.
type MessageSchedulingStateBox struct {
	MessageSchedulingState MessageSchedulingStateClass
}

// Decode implements bin.Decoder for MessageSchedulingStateBox.
func (b *MessageSchedulingStateBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessageSchedulingStateBox to nil")
	}
	v, err := DecodeMessageSchedulingState(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.MessageSchedulingState = v
	return nil
}

// Encode implements bin.Encode for MessageSchedulingStateBox.
func (b *MessageSchedulingStateBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.MessageSchedulingState == nil {
		return fmt.Errorf("unable to encode MessageSchedulingStateClass as nil")
	}
	return b.MessageSchedulingState.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for MessageSchedulingStateBox.
func (b *MessageSchedulingStateBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessageSchedulingStateBox to nil")
	}
	v, err := DecodeTDLibJSONMessageSchedulingState(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.MessageSchedulingState = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for MessageSchedulingStateBox.
func (b *MessageSchedulingStateBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.MessageSchedulingState == nil {
		return fmt.Errorf("unable to encode MessageSchedulingStateClass as nil")
	}
	return b.MessageSchedulingState.EncodeTDLibJSON(buf)
}
