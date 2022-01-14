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

// CheckAuthenticationBotTokenRequest represents TL type `checkAuthenticationBotToken#261b4476`.
type CheckAuthenticationBotTokenRequest struct {
	// The bot token
	Token string
}

// CheckAuthenticationBotTokenRequestTypeID is TL type id of CheckAuthenticationBotTokenRequest.
const CheckAuthenticationBotTokenRequestTypeID = 0x261b4476

// Ensuring interfaces in compile-time for CheckAuthenticationBotTokenRequest.
var (
	_ bin.Encoder     = &CheckAuthenticationBotTokenRequest{}
	_ bin.Decoder     = &CheckAuthenticationBotTokenRequest{}
	_ bin.BareEncoder = &CheckAuthenticationBotTokenRequest{}
	_ bin.BareDecoder = &CheckAuthenticationBotTokenRequest{}
)

func (c *CheckAuthenticationBotTokenRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Token == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CheckAuthenticationBotTokenRequest) String() string {
	if c == nil {
		return "CheckAuthenticationBotTokenRequest(nil)"
	}
	type Alias CheckAuthenticationBotTokenRequest
	return fmt.Sprintf("CheckAuthenticationBotTokenRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CheckAuthenticationBotTokenRequest) TypeID() uint32 {
	return CheckAuthenticationBotTokenRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*CheckAuthenticationBotTokenRequest) TypeName() string {
	return "checkAuthenticationBotToken"
}

// TypeInfo returns info about TL type.
func (c *CheckAuthenticationBotTokenRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "checkAuthenticationBotToken",
		ID:   CheckAuthenticationBotTokenRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Token",
			SchemaName: "token",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CheckAuthenticationBotTokenRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkAuthenticationBotToken#261b4476 as nil")
	}
	b.PutID(CheckAuthenticationBotTokenRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CheckAuthenticationBotTokenRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkAuthenticationBotToken#261b4476 as nil")
	}
	b.PutString(c.Token)
	return nil
}

// Decode implements bin.Decoder.
func (c *CheckAuthenticationBotTokenRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkAuthenticationBotToken#261b4476 to nil")
	}
	if err := b.ConsumeID(CheckAuthenticationBotTokenRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode checkAuthenticationBotToken#261b4476: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CheckAuthenticationBotTokenRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkAuthenticationBotToken#261b4476 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode checkAuthenticationBotToken#261b4476: field token: %w", err)
		}
		c.Token = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CheckAuthenticationBotTokenRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode checkAuthenticationBotToken#261b4476 as nil")
	}
	b.ObjStart()
	b.PutID("checkAuthenticationBotToken")
	b.Comma()
	b.FieldStart("token")
	b.PutString(c.Token)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CheckAuthenticationBotTokenRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode checkAuthenticationBotToken#261b4476 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("checkAuthenticationBotToken"); err != nil {
				return fmt.Errorf("unable to decode checkAuthenticationBotToken#261b4476: %w", err)
			}
		case "token":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode checkAuthenticationBotToken#261b4476: field token: %w", err)
			}
			c.Token = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetToken returns value of Token field.
func (c *CheckAuthenticationBotTokenRequest) GetToken() (value string) {
	if c == nil {
		return
	}
	return c.Token
}

// CheckAuthenticationBotToken invokes method checkAuthenticationBotToken#261b4476 returning error if any.
func (c *Client) CheckAuthenticationBotToken(ctx context.Context, token string) error {
	var ok Ok

	request := &CheckAuthenticationBotTokenRequest{
		Token: token,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
