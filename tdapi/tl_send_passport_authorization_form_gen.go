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

// SendPassportAuthorizationFormRequest represents TL type `sendPassportAuthorizationForm#eee589b5`.
type SendPassportAuthorizationFormRequest struct {
	// Authorization form identifier
	AutorizationFormID int32
	// Types of Telegram Passport elements chosen by user to complete the authorization form
	Types []PassportElementTypeClass
}

// SendPassportAuthorizationFormRequestTypeID is TL type id of SendPassportAuthorizationFormRequest.
const SendPassportAuthorizationFormRequestTypeID = 0xeee589b5

// Ensuring interfaces in compile-time for SendPassportAuthorizationFormRequest.
var (
	_ bin.Encoder     = &SendPassportAuthorizationFormRequest{}
	_ bin.Decoder     = &SendPassportAuthorizationFormRequest{}
	_ bin.BareEncoder = &SendPassportAuthorizationFormRequest{}
	_ bin.BareDecoder = &SendPassportAuthorizationFormRequest{}
)

func (s *SendPassportAuthorizationFormRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.AutorizationFormID == 0) {
		return false
	}
	if !(s.Types == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SendPassportAuthorizationFormRequest) String() string {
	if s == nil {
		return "SendPassportAuthorizationFormRequest(nil)"
	}
	type Alias SendPassportAuthorizationFormRequest
	return fmt.Sprintf("SendPassportAuthorizationFormRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SendPassportAuthorizationFormRequest) TypeID() uint32 {
	return SendPassportAuthorizationFormRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SendPassportAuthorizationFormRequest) TypeName() string {
	return "sendPassportAuthorizationForm"
}

// TypeInfo returns info about TL type.
func (s *SendPassportAuthorizationFormRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "sendPassportAuthorizationForm",
		ID:   SendPassportAuthorizationFormRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "AutorizationFormID",
			SchemaName: "autorization_form_id",
		},
		{
			Name:       "Types",
			SchemaName: "types",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SendPassportAuthorizationFormRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendPassportAuthorizationForm#eee589b5 as nil")
	}
	b.PutID(SendPassportAuthorizationFormRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SendPassportAuthorizationFormRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendPassportAuthorizationForm#eee589b5 as nil")
	}
	b.PutInt32(s.AutorizationFormID)
	b.PutInt(len(s.Types))
	for idx, v := range s.Types {
		if v == nil {
			return fmt.Errorf("unable to encode sendPassportAuthorizationForm#eee589b5: field types element with index %d is nil", idx)
		}
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare sendPassportAuthorizationForm#eee589b5: field types element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SendPassportAuthorizationFormRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendPassportAuthorizationForm#eee589b5 to nil")
	}
	if err := b.ConsumeID(SendPassportAuthorizationFormRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode sendPassportAuthorizationForm#eee589b5: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SendPassportAuthorizationFormRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendPassportAuthorizationForm#eee589b5 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode sendPassportAuthorizationForm#eee589b5: field autorization_form_id: %w", err)
		}
		s.AutorizationFormID = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode sendPassportAuthorizationForm#eee589b5: field types: %w", err)
		}

		if headerLen > 0 {
			s.Types = make([]PassportElementTypeClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePassportElementType(b)
			if err != nil {
				return fmt.Errorf("unable to decode sendPassportAuthorizationForm#eee589b5: field types: %w", err)
			}
			s.Types = append(s.Types, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SendPassportAuthorizationFormRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode sendPassportAuthorizationForm#eee589b5 as nil")
	}
	b.ObjStart()
	b.PutID("sendPassportAuthorizationForm")
	b.Comma()
	b.FieldStart("autorization_form_id")
	b.PutInt32(s.AutorizationFormID)
	b.Comma()
	b.FieldStart("types")
	b.ArrStart()
	for idx, v := range s.Types {
		if v == nil {
			return fmt.Errorf("unable to encode sendPassportAuthorizationForm#eee589b5: field types element with index %d is nil", idx)
		}
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode sendPassportAuthorizationForm#eee589b5: field types element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SendPassportAuthorizationFormRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode sendPassportAuthorizationForm#eee589b5 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("sendPassportAuthorizationForm"); err != nil {
				return fmt.Errorf("unable to decode sendPassportAuthorizationForm#eee589b5: %w", err)
			}
		case "autorization_form_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode sendPassportAuthorizationForm#eee589b5: field autorization_form_id: %w", err)
			}
			s.AutorizationFormID = value
		case "types":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := DecodeTDLibJSONPassportElementType(b)
				if err != nil {
					return fmt.Errorf("unable to decode sendPassportAuthorizationForm#eee589b5: field types: %w", err)
				}
				s.Types = append(s.Types, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode sendPassportAuthorizationForm#eee589b5: field types: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetAutorizationFormID returns value of AutorizationFormID field.
func (s *SendPassportAuthorizationFormRequest) GetAutorizationFormID() (value int32) {
	if s == nil {
		return
	}
	return s.AutorizationFormID
}

// GetTypes returns value of Types field.
func (s *SendPassportAuthorizationFormRequest) GetTypes() (value []PassportElementTypeClass) {
	if s == nil {
		return
	}
	return s.Types
}

// SendPassportAuthorizationForm invokes method sendPassportAuthorizationForm#eee589b5 returning error if any.
func (c *Client) SendPassportAuthorizationForm(ctx context.Context, request *SendPassportAuthorizationFormRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
