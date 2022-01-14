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

// ReportChatRequest represents TL type `reportChat#a19024af`.
type ReportChatRequest struct {
	// Chat identifier
	ChatID int64
	// Identifiers of reported messages, if any
	MessageIDs []int64
	// The reason for reporting the chat
	Reason ChatReportReasonClass
	// Additional report details; 0-1024 characters
	Text string
}

// ReportChatRequestTypeID is TL type id of ReportChatRequest.
const ReportChatRequestTypeID = 0xa19024af

// Ensuring interfaces in compile-time for ReportChatRequest.
var (
	_ bin.Encoder     = &ReportChatRequest{}
	_ bin.Decoder     = &ReportChatRequest{}
	_ bin.BareEncoder = &ReportChatRequest{}
	_ bin.BareDecoder = &ReportChatRequest{}
)

func (r *ReportChatRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.ChatID == 0) {
		return false
	}
	if !(r.MessageIDs == nil) {
		return false
	}
	if !(r.Reason == nil) {
		return false
	}
	if !(r.Text == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReportChatRequest) String() string {
	if r == nil {
		return "ReportChatRequest(nil)"
	}
	type Alias ReportChatRequest
	return fmt.Sprintf("ReportChatRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReportChatRequest) TypeID() uint32 {
	return ReportChatRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ReportChatRequest) TypeName() string {
	return "reportChat"
}

// TypeInfo returns info about TL type.
func (r *ReportChatRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reportChat",
		ID:   ReportChatRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageIDs",
			SchemaName: "message_ids",
		},
		{
			Name:       "Reason",
			SchemaName: "reason",
		},
		{
			Name:       "Text",
			SchemaName: "text",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReportChatRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChat#a19024af as nil")
	}
	b.PutID(ReportChatRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReportChatRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChat#a19024af as nil")
	}
	b.PutInt53(r.ChatID)
	b.PutInt(len(r.MessageIDs))
	for _, v := range r.MessageIDs {
		b.PutInt53(v)
	}
	if r.Reason == nil {
		return fmt.Errorf("unable to encode reportChat#a19024af: field reason is nil")
	}
	if err := r.Reason.Encode(b); err != nil {
		return fmt.Errorf("unable to encode reportChat#a19024af: field reason: %w", err)
	}
	b.PutString(r.Text)
	return nil
}

// Decode implements bin.Decoder.
func (r *ReportChatRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChat#a19024af to nil")
	}
	if err := b.ConsumeID(ReportChatRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode reportChat#a19024af: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReportChatRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChat#a19024af to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode reportChat#a19024af: field chat_id: %w", err)
		}
		r.ChatID = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode reportChat#a19024af: field message_ids: %w", err)
		}

		if headerLen > 0 {
			r.MessageIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode reportChat#a19024af: field message_ids: %w", err)
			}
			r.MessageIDs = append(r.MessageIDs, value)
		}
	}
	{
		value, err := DecodeChatReportReason(b)
		if err != nil {
			return fmt.Errorf("unable to decode reportChat#a19024af: field reason: %w", err)
		}
		r.Reason = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode reportChat#a19024af: field text: %w", err)
		}
		r.Text = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReportChatRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChat#a19024af as nil")
	}
	b.ObjStart()
	b.PutID("reportChat")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(r.ChatID)
	b.Comma()
	b.FieldStart("message_ids")
	b.ArrStart()
	for _, v := range r.MessageIDs {
		b.PutInt53(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("reason")
	if r.Reason == nil {
		return fmt.Errorf("unable to encode reportChat#a19024af: field reason is nil")
	}
	if err := r.Reason.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode reportChat#a19024af: field reason: %w", err)
	}
	b.Comma()
	b.FieldStart("text")
	b.PutString(r.Text)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReportChatRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChat#a19024af to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("reportChat"); err != nil {
				return fmt.Errorf("unable to decode reportChat#a19024af: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode reportChat#a19024af: field chat_id: %w", err)
			}
			r.ChatID = value
		case "message_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int53()
				if err != nil {
					return fmt.Errorf("unable to decode reportChat#a19024af: field message_ids: %w", err)
				}
				r.MessageIDs = append(r.MessageIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode reportChat#a19024af: field message_ids: %w", err)
			}
		case "reason":
			value, err := DecodeTDLibJSONChatReportReason(b)
			if err != nil {
				return fmt.Errorf("unable to decode reportChat#a19024af: field reason: %w", err)
			}
			r.Reason = value
		case "text":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode reportChat#a19024af: field text: %w", err)
			}
			r.Text = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (r *ReportChatRequest) GetChatID() (value int64) {
	if r == nil {
		return
	}
	return r.ChatID
}

// GetMessageIDs returns value of MessageIDs field.
func (r *ReportChatRequest) GetMessageIDs() (value []int64) {
	if r == nil {
		return
	}
	return r.MessageIDs
}

// GetReason returns value of Reason field.
func (r *ReportChatRequest) GetReason() (value ChatReportReasonClass) {
	if r == nil {
		return
	}
	return r.Reason
}

// GetText returns value of Text field.
func (r *ReportChatRequest) GetText() (value string) {
	if r == nil {
		return
	}
	return r.Text
}

// ReportChat invokes method reportChat#a19024af returning error if any.
func (c *Client) ReportChat(ctx context.Context, request *ReportChatRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
