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

// ProcessPushNotificationRequest represents TL type `processPushNotification#2ee3c890`.
type ProcessPushNotificationRequest struct {
	// JSON-encoded push notification payload with all fields sent by the server, and "google
	// sent_time" and "google.notification.sound" fields added
	Payload string
}

// ProcessPushNotificationRequestTypeID is TL type id of ProcessPushNotificationRequest.
const ProcessPushNotificationRequestTypeID = 0x2ee3c890

// Ensuring interfaces in compile-time for ProcessPushNotificationRequest.
var (
	_ bin.Encoder     = &ProcessPushNotificationRequest{}
	_ bin.Decoder     = &ProcessPushNotificationRequest{}
	_ bin.BareEncoder = &ProcessPushNotificationRequest{}
	_ bin.BareDecoder = &ProcessPushNotificationRequest{}
)

func (p *ProcessPushNotificationRequest) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Payload == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *ProcessPushNotificationRequest) String() string {
	if p == nil {
		return "ProcessPushNotificationRequest(nil)"
	}
	type Alias ProcessPushNotificationRequest
	return fmt.Sprintf("ProcessPushNotificationRequest%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ProcessPushNotificationRequest) TypeID() uint32 {
	return ProcessPushNotificationRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ProcessPushNotificationRequest) TypeName() string {
	return "processPushNotification"
}

// TypeInfo returns info about TL type.
func (p *ProcessPushNotificationRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "processPushNotification",
		ID:   ProcessPushNotificationRequestTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Payload",
			SchemaName: "payload",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *ProcessPushNotificationRequest) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode processPushNotification#2ee3c890 as nil")
	}
	b.PutID(ProcessPushNotificationRequestTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *ProcessPushNotificationRequest) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode processPushNotification#2ee3c890 as nil")
	}
	b.PutString(p.Payload)
	return nil
}

// Decode implements bin.Decoder.
func (p *ProcessPushNotificationRequest) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode processPushNotification#2ee3c890 to nil")
	}
	if err := b.ConsumeID(ProcessPushNotificationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode processPushNotification#2ee3c890: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *ProcessPushNotificationRequest) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode processPushNotification#2ee3c890 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode processPushNotification#2ee3c890: field payload: %w", err)
		}
		p.Payload = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *ProcessPushNotificationRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode processPushNotification#2ee3c890 as nil")
	}
	b.ObjStart()
	b.PutID("processPushNotification")
	b.Comma()
	b.FieldStart("payload")
	b.PutString(p.Payload)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *ProcessPushNotificationRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode processPushNotification#2ee3c890 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("processPushNotification"); err != nil {
				return fmt.Errorf("unable to decode processPushNotification#2ee3c890: %w", err)
			}
		case "payload":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode processPushNotification#2ee3c890: field payload: %w", err)
			}
			p.Payload = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetPayload returns value of Payload field.
func (p *ProcessPushNotificationRequest) GetPayload() (value string) {
	if p == nil {
		return
	}
	return p.Payload
}

// ProcessPushNotification invokes method processPushNotification#2ee3c890 returning error if any.
func (c *Client) ProcessPushNotification(ctx context.Context, payload string) error {
	var ok Ok

	request := &ProcessPushNotificationRequest{
		Payload: payload,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
