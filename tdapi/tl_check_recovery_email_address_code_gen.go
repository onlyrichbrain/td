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

// CheckRecoveryEmailAddressCodeRequest represents TL type `checkRecoveryEmailAddressCode#88f7981b`.
type CheckRecoveryEmailAddressCodeRequest struct {
	// Verification code to check
	Code string
}

// CheckRecoveryEmailAddressCodeRequestTypeID is TL type id of CheckRecoveryEmailAddressCodeRequest.
const CheckRecoveryEmailAddressCodeRequestTypeID = 0x88f7981b

// Ensuring interfaces in compile-time for CheckRecoveryEmailAddressCodeRequest.
var (
	_ bin.Encoder     = &CheckRecoveryEmailAddressCodeRequest{}
	_ bin.Decoder     = &CheckRecoveryEmailAddressCodeRequest{}
	_ bin.BareEncoder = &CheckRecoveryEmailAddressCodeRequest{}
	_ bin.BareDecoder = &CheckRecoveryEmailAddressCodeRequest{}
)

func (c *CheckRecoveryEmailAddressCodeRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Code == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CheckRecoveryEmailAddressCodeRequest) String() string {
	if c == nil {
		return "CheckRecoveryEmailAddressCodeRequest(nil)"
	}
	type Alias CheckRecoveryEmailAddressCodeRequest
	return fmt.Sprintf("CheckRecoveryEmailAddressCodeRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CheckRecoveryEmailAddressCodeRequest) TypeID() uint32 {
	return CheckRecoveryEmailAddressCodeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*CheckRecoveryEmailAddressCodeRequest) TypeName() string {
	return "checkRecoveryEmailAddressCode"
}

// TypeInfo returns info about TL type.
func (c *CheckRecoveryEmailAddressCodeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "checkRecoveryEmailAddressCode",
		ID:   CheckRecoveryEmailAddressCodeRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Code",
			SchemaName: "code",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CheckRecoveryEmailAddressCodeRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkRecoveryEmailAddressCode#88f7981b as nil")
	}
	b.PutID(CheckRecoveryEmailAddressCodeRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CheckRecoveryEmailAddressCodeRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkRecoveryEmailAddressCode#88f7981b as nil")
	}
	b.PutString(c.Code)
	return nil
}

// Decode implements bin.Decoder.
func (c *CheckRecoveryEmailAddressCodeRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkRecoveryEmailAddressCode#88f7981b to nil")
	}
	if err := b.ConsumeID(CheckRecoveryEmailAddressCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode checkRecoveryEmailAddressCode#88f7981b: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CheckRecoveryEmailAddressCodeRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkRecoveryEmailAddressCode#88f7981b to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode checkRecoveryEmailAddressCode#88f7981b: field code: %w", err)
		}
		c.Code = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CheckRecoveryEmailAddressCodeRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode checkRecoveryEmailAddressCode#88f7981b as nil")
	}
	b.ObjStart()
	b.PutID("checkRecoveryEmailAddressCode")
	b.Comma()
	b.FieldStart("code")
	b.PutString(c.Code)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CheckRecoveryEmailAddressCodeRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode checkRecoveryEmailAddressCode#88f7981b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("checkRecoveryEmailAddressCode"); err != nil {
				return fmt.Errorf("unable to decode checkRecoveryEmailAddressCode#88f7981b: %w", err)
			}
		case "code":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode checkRecoveryEmailAddressCode#88f7981b: field code: %w", err)
			}
			c.Code = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCode returns value of Code field.
func (c *CheckRecoveryEmailAddressCodeRequest) GetCode() (value string) {
	if c == nil {
		return
	}
	return c.Code
}

// CheckRecoveryEmailAddressCode invokes method checkRecoveryEmailAddressCode#88f7981b returning error if any.
func (c *Client) CheckRecoveryEmailAddressCode(ctx context.Context, code string) (*PasswordState, error) {
	var result PasswordState

	request := &CheckRecoveryEmailAddressCodeRequest{
		Code: code,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
