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

// DeleteBusinessChatLinkRequest represents TL type `deleteBusinessChatLink#be526747`.
type DeleteBusinessChatLinkRequest struct {
	// The link to delete
	Link string
}

// DeleteBusinessChatLinkRequestTypeID is TL type id of DeleteBusinessChatLinkRequest.
const DeleteBusinessChatLinkRequestTypeID = 0xbe526747

// Ensuring interfaces in compile-time for DeleteBusinessChatLinkRequest.
var (
	_ bin.Encoder     = &DeleteBusinessChatLinkRequest{}
	_ bin.Decoder     = &DeleteBusinessChatLinkRequest{}
	_ bin.BareEncoder = &DeleteBusinessChatLinkRequest{}
	_ bin.BareDecoder = &DeleteBusinessChatLinkRequest{}
)

func (d *DeleteBusinessChatLinkRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Link == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DeleteBusinessChatLinkRequest) String() string {
	if d == nil {
		return "DeleteBusinessChatLinkRequest(nil)"
	}
	type Alias DeleteBusinessChatLinkRequest
	return fmt.Sprintf("DeleteBusinessChatLinkRequest%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DeleteBusinessChatLinkRequest) TypeID() uint32 {
	return DeleteBusinessChatLinkRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*DeleteBusinessChatLinkRequest) TypeName() string {
	return "deleteBusinessChatLink"
}

// TypeInfo returns info about TL type.
func (d *DeleteBusinessChatLinkRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "deleteBusinessChatLink",
		ID:   DeleteBusinessChatLinkRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Link",
			SchemaName: "link",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DeleteBusinessChatLinkRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteBusinessChatLink#be526747 as nil")
	}
	b.PutID(DeleteBusinessChatLinkRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DeleteBusinessChatLinkRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteBusinessChatLink#be526747 as nil")
	}
	b.PutString(d.Link)
	return nil
}

// Decode implements bin.Decoder.
func (d *DeleteBusinessChatLinkRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteBusinessChatLink#be526747 to nil")
	}
	if err := b.ConsumeID(DeleteBusinessChatLinkRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode deleteBusinessChatLink#be526747: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DeleteBusinessChatLinkRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteBusinessChatLink#be526747 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode deleteBusinessChatLink#be526747: field link: %w", err)
		}
		d.Link = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (d *DeleteBusinessChatLinkRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteBusinessChatLink#be526747 as nil")
	}
	b.ObjStart()
	b.PutID("deleteBusinessChatLink")
	b.Comma()
	b.FieldStart("link")
	b.PutString(d.Link)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (d *DeleteBusinessChatLinkRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteBusinessChatLink#be526747 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("deleteBusinessChatLink"); err != nil {
				return fmt.Errorf("unable to decode deleteBusinessChatLink#be526747: %w", err)
			}
		case "link":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode deleteBusinessChatLink#be526747: field link: %w", err)
			}
			d.Link = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLink returns value of Link field.
func (d *DeleteBusinessChatLinkRequest) GetLink() (value string) {
	if d == nil {
		return
	}
	return d.Link
}

// DeleteBusinessChatLink invokes method deleteBusinessChatLink#be526747 returning error if any.
func (c *Client) DeleteBusinessChatLink(ctx context.Context, link string) error {
	var ok Ok

	request := &DeleteBusinessChatLinkRequest{
		Link: link,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}