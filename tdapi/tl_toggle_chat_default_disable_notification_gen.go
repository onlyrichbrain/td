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

// ToggleChatDefaultDisableNotificationRequest represents TL type `toggleChatDefaultDisableNotification#12c36012`.
type ToggleChatDefaultDisableNotificationRequest struct {
	// Chat identifier
	ChatID int64
	// New value of default_disable_notification
	DefaultDisableNotification bool
}

// ToggleChatDefaultDisableNotificationRequestTypeID is TL type id of ToggleChatDefaultDisableNotificationRequest.
const ToggleChatDefaultDisableNotificationRequestTypeID = 0x12c36012

// Ensuring interfaces in compile-time for ToggleChatDefaultDisableNotificationRequest.
var (
	_ bin.Encoder     = &ToggleChatDefaultDisableNotificationRequest{}
	_ bin.Decoder     = &ToggleChatDefaultDisableNotificationRequest{}
	_ bin.BareEncoder = &ToggleChatDefaultDisableNotificationRequest{}
	_ bin.BareDecoder = &ToggleChatDefaultDisableNotificationRequest{}
)

func (t *ToggleChatDefaultDisableNotificationRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.ChatID == 0) {
		return false
	}
	if !(t.DefaultDisableNotification == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ToggleChatDefaultDisableNotificationRequest) String() string {
	if t == nil {
		return "ToggleChatDefaultDisableNotificationRequest(nil)"
	}
	type Alias ToggleChatDefaultDisableNotificationRequest
	return fmt.Sprintf("ToggleChatDefaultDisableNotificationRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ToggleChatDefaultDisableNotificationRequest) TypeID() uint32 {
	return ToggleChatDefaultDisableNotificationRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ToggleChatDefaultDisableNotificationRequest) TypeName() string {
	return "toggleChatDefaultDisableNotification"
}

// TypeInfo returns info about TL type.
func (t *ToggleChatDefaultDisableNotificationRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "toggleChatDefaultDisableNotification",
		ID:   ToggleChatDefaultDisableNotificationRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "DefaultDisableNotification",
			SchemaName: "default_disable_notification",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ToggleChatDefaultDisableNotificationRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleChatDefaultDisableNotification#12c36012 as nil")
	}
	b.PutID(ToggleChatDefaultDisableNotificationRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ToggleChatDefaultDisableNotificationRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleChatDefaultDisableNotification#12c36012 as nil")
	}
	b.PutInt53(t.ChatID)
	b.PutBool(t.DefaultDisableNotification)
	return nil
}

// Decode implements bin.Decoder.
func (t *ToggleChatDefaultDisableNotificationRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleChatDefaultDisableNotification#12c36012 to nil")
	}
	if err := b.ConsumeID(ToggleChatDefaultDisableNotificationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode toggleChatDefaultDisableNotification#12c36012: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ToggleChatDefaultDisableNotificationRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleChatDefaultDisableNotification#12c36012 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode toggleChatDefaultDisableNotification#12c36012: field chat_id: %w", err)
		}
		t.ChatID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode toggleChatDefaultDisableNotification#12c36012: field default_disable_notification: %w", err)
		}
		t.DefaultDisableNotification = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *ToggleChatDefaultDisableNotificationRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleChatDefaultDisableNotification#12c36012 as nil")
	}
	b.ObjStart()
	b.PutID("toggleChatDefaultDisableNotification")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(t.ChatID)
	b.Comma()
	b.FieldStart("default_disable_notification")
	b.PutBool(t.DefaultDisableNotification)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *ToggleChatDefaultDisableNotificationRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleChatDefaultDisableNotification#12c36012 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("toggleChatDefaultDisableNotification"); err != nil {
				return fmt.Errorf("unable to decode toggleChatDefaultDisableNotification#12c36012: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode toggleChatDefaultDisableNotification#12c36012: field chat_id: %w", err)
			}
			t.ChatID = value
		case "default_disable_notification":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode toggleChatDefaultDisableNotification#12c36012: field default_disable_notification: %w", err)
			}
			t.DefaultDisableNotification = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (t *ToggleChatDefaultDisableNotificationRequest) GetChatID() (value int64) {
	if t == nil {
		return
	}
	return t.ChatID
}

// GetDefaultDisableNotification returns value of DefaultDisableNotification field.
func (t *ToggleChatDefaultDisableNotificationRequest) GetDefaultDisableNotification() (value bool) {
	if t == nil {
		return
	}
	return t.DefaultDisableNotification
}

// ToggleChatDefaultDisableNotification invokes method toggleChatDefaultDisableNotification#12c36012 returning error if any.
func (c *Client) ToggleChatDefaultDisableNotification(ctx context.Context, request *ToggleChatDefaultDisableNotificationRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
