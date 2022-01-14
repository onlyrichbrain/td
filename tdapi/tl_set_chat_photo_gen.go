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

// SetChatPhotoRequest represents TL type `setChatPhoto#e97b8d03`.
type SetChatPhotoRequest struct {
	// Chat identifier
	ChatID int64
	// New chat photo; pass null to delete the chat photo
	Photo InputChatPhotoClass
}

// SetChatPhotoRequestTypeID is TL type id of SetChatPhotoRequest.
const SetChatPhotoRequestTypeID = 0xe97b8d03

// Ensuring interfaces in compile-time for SetChatPhotoRequest.
var (
	_ bin.Encoder     = &SetChatPhotoRequest{}
	_ bin.Decoder     = &SetChatPhotoRequest{}
	_ bin.BareEncoder = &SetChatPhotoRequest{}
	_ bin.BareDecoder = &SetChatPhotoRequest{}
)

func (s *SetChatPhotoRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.Photo == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetChatPhotoRequest) String() string {
	if s == nil {
		return "SetChatPhotoRequest(nil)"
	}
	type Alias SetChatPhotoRequest
	return fmt.Sprintf("SetChatPhotoRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetChatPhotoRequest) TypeID() uint32 {
	return SetChatPhotoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetChatPhotoRequest) TypeName() string {
	return "setChatPhoto"
}

// TypeInfo returns info about TL type.
func (s *SetChatPhotoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setChatPhoto",
		ID:   SetChatPhotoRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "Photo",
			SchemaName: "photo",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetChatPhotoRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatPhoto#e97b8d03 as nil")
	}
	b.PutID(SetChatPhotoRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetChatPhotoRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatPhoto#e97b8d03 as nil")
	}
	b.PutInt53(s.ChatID)
	if s.Photo == nil {
		return fmt.Errorf("unable to encode setChatPhoto#e97b8d03: field photo is nil")
	}
	if err := s.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setChatPhoto#e97b8d03: field photo: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SetChatPhotoRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatPhoto#e97b8d03 to nil")
	}
	if err := b.ConsumeID(SetChatPhotoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setChatPhoto#e97b8d03: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetChatPhotoRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatPhoto#e97b8d03 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode setChatPhoto#e97b8d03: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		value, err := DecodeInputChatPhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode setChatPhoto#e97b8d03: field photo: %w", err)
		}
		s.Photo = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetChatPhotoRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatPhoto#e97b8d03 as nil")
	}
	b.ObjStart()
	b.PutID("setChatPhoto")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(s.ChatID)
	b.Comma()
	b.FieldStart("photo")
	if s.Photo == nil {
		return fmt.Errorf("unable to encode setChatPhoto#e97b8d03: field photo is nil")
	}
	if err := s.Photo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setChatPhoto#e97b8d03: field photo: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetChatPhotoRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatPhoto#e97b8d03 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setChatPhoto"); err != nil {
				return fmt.Errorf("unable to decode setChatPhoto#e97b8d03: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode setChatPhoto#e97b8d03: field chat_id: %w", err)
			}
			s.ChatID = value
		case "photo":
			value, err := DecodeTDLibJSONInputChatPhoto(b)
			if err != nil {
				return fmt.Errorf("unable to decode setChatPhoto#e97b8d03: field photo: %w", err)
			}
			s.Photo = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (s *SetChatPhotoRequest) GetChatID() (value int64) {
	if s == nil {
		return
	}
	return s.ChatID
}

// GetPhoto returns value of Photo field.
func (s *SetChatPhotoRequest) GetPhoto() (value InputChatPhotoClass) {
	if s == nil {
		return
	}
	return s.Photo
}

// SetChatPhoto invokes method setChatPhoto#e97b8d03 returning error if any.
func (c *Client) SetChatPhoto(ctx context.Context, request *SetChatPhotoRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
