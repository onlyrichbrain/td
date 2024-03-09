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

// DeleteDefaultBackgroundRequest represents TL type `deleteDefaultBackground#b2a4ed3e`.
type DeleteDefaultBackgroundRequest struct {
	// Pass true if the background is deleted for a dark theme
	ForDarkTheme bool
}

// DeleteDefaultBackgroundRequestTypeID is TL type id of DeleteDefaultBackgroundRequest.
const DeleteDefaultBackgroundRequestTypeID = 0xb2a4ed3e

// Ensuring interfaces in compile-time for DeleteDefaultBackgroundRequest.
var (
	_ bin.Encoder     = &DeleteDefaultBackgroundRequest{}
	_ bin.Decoder     = &DeleteDefaultBackgroundRequest{}
	_ bin.BareEncoder = &DeleteDefaultBackgroundRequest{}
	_ bin.BareDecoder = &DeleteDefaultBackgroundRequest{}
)

func (d *DeleteDefaultBackgroundRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.ForDarkTheme == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DeleteDefaultBackgroundRequest) String() string {
	if d == nil {
		return "DeleteDefaultBackgroundRequest(nil)"
	}
	type Alias DeleteDefaultBackgroundRequest
	return fmt.Sprintf("DeleteDefaultBackgroundRequest%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DeleteDefaultBackgroundRequest) TypeID() uint32 {
	return DeleteDefaultBackgroundRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*DeleteDefaultBackgroundRequest) TypeName() string {
	return "deleteDefaultBackground"
}

// TypeInfo returns info about TL type.
func (d *DeleteDefaultBackgroundRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "deleteDefaultBackground",
		ID:   DeleteDefaultBackgroundRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ForDarkTheme",
			SchemaName: "for_dark_theme",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DeleteDefaultBackgroundRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteDefaultBackground#b2a4ed3e as nil")
	}
	b.PutID(DeleteDefaultBackgroundRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DeleteDefaultBackgroundRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteDefaultBackground#b2a4ed3e as nil")
	}
	b.PutBool(d.ForDarkTheme)
	return nil
}

// Decode implements bin.Decoder.
func (d *DeleteDefaultBackgroundRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteDefaultBackground#b2a4ed3e to nil")
	}
	if err := b.ConsumeID(DeleteDefaultBackgroundRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode deleteDefaultBackground#b2a4ed3e: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DeleteDefaultBackgroundRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteDefaultBackground#b2a4ed3e to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode deleteDefaultBackground#b2a4ed3e: field for_dark_theme: %w", err)
		}
		d.ForDarkTheme = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (d *DeleteDefaultBackgroundRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteDefaultBackground#b2a4ed3e as nil")
	}
	b.ObjStart()
	b.PutID("deleteDefaultBackground")
	b.Comma()
	b.FieldStart("for_dark_theme")
	b.PutBool(d.ForDarkTheme)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (d *DeleteDefaultBackgroundRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteDefaultBackground#b2a4ed3e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("deleteDefaultBackground"); err != nil {
				return fmt.Errorf("unable to decode deleteDefaultBackground#b2a4ed3e: %w", err)
			}
		case "for_dark_theme":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode deleteDefaultBackground#b2a4ed3e: field for_dark_theme: %w", err)
			}
			d.ForDarkTheme = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetForDarkTheme returns value of ForDarkTheme field.
func (d *DeleteDefaultBackgroundRequest) GetForDarkTheme() (value bool) {
	if d == nil {
		return
	}
	return d.ForDarkTheme
}

// DeleteDefaultBackground invokes method deleteDefaultBackground#b2a4ed3e returning error if any.
func (c *Client) DeleteDefaultBackground(ctx context.Context, fordarktheme bool) error {
	var ok Ok

	request := &DeleteDefaultBackgroundRequest{
		ForDarkTheme: fordarktheme,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}