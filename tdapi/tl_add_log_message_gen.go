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

// AddLogMessageRequest represents TL type `addLogMessage#5f36cfec`.
type AddLogMessageRequest struct {
	// The minimum verbosity level needed for the message to be logged; 0-1023
	VerbosityLevel int32
	// Text of a message to log
	Text string
}

// AddLogMessageRequestTypeID is TL type id of AddLogMessageRequest.
const AddLogMessageRequestTypeID = 0x5f36cfec

// Ensuring interfaces in compile-time for AddLogMessageRequest.
var (
	_ bin.Encoder     = &AddLogMessageRequest{}
	_ bin.Decoder     = &AddLogMessageRequest{}
	_ bin.BareEncoder = &AddLogMessageRequest{}
	_ bin.BareDecoder = &AddLogMessageRequest{}
)

func (a *AddLogMessageRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.VerbosityLevel == 0) {
		return false
	}
	if !(a.Text == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AddLogMessageRequest) String() string {
	if a == nil {
		return "AddLogMessageRequest(nil)"
	}
	type Alias AddLogMessageRequest
	return fmt.Sprintf("AddLogMessageRequest%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AddLogMessageRequest) TypeID() uint32 {
	return AddLogMessageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AddLogMessageRequest) TypeName() string {
	return "addLogMessage"
}

// TypeInfo returns info about TL type.
func (a *AddLogMessageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "addLogMessage",
		ID:   AddLogMessageRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "VerbosityLevel",
			SchemaName: "verbosity_level",
		},
		{
			Name:       "Text",
			SchemaName: "text",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AddLogMessageRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode addLogMessage#5f36cfec as nil")
	}
	b.PutID(AddLogMessageRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AddLogMessageRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode addLogMessage#5f36cfec as nil")
	}
	b.PutInt32(a.VerbosityLevel)
	b.PutString(a.Text)
	return nil
}

// Decode implements bin.Decoder.
func (a *AddLogMessageRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode addLogMessage#5f36cfec to nil")
	}
	if err := b.ConsumeID(AddLogMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode addLogMessage#5f36cfec: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AddLogMessageRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode addLogMessage#5f36cfec to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode addLogMessage#5f36cfec: field verbosity_level: %w", err)
		}
		a.VerbosityLevel = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode addLogMessage#5f36cfec: field text: %w", err)
		}
		a.Text = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AddLogMessageRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode addLogMessage#5f36cfec as nil")
	}
	b.ObjStart()
	b.PutID("addLogMessage")
	b.Comma()
	b.FieldStart("verbosity_level")
	b.PutInt32(a.VerbosityLevel)
	b.Comma()
	b.FieldStart("text")
	b.PutString(a.Text)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *AddLogMessageRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode addLogMessage#5f36cfec to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("addLogMessage"); err != nil {
				return fmt.Errorf("unable to decode addLogMessage#5f36cfec: %w", err)
			}
		case "verbosity_level":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode addLogMessage#5f36cfec: field verbosity_level: %w", err)
			}
			a.VerbosityLevel = value
		case "text":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode addLogMessage#5f36cfec: field text: %w", err)
			}
			a.Text = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetVerbosityLevel returns value of VerbosityLevel field.
func (a *AddLogMessageRequest) GetVerbosityLevel() (value int32) {
	if a == nil {
		return
	}
	return a.VerbosityLevel
}

// GetText returns value of Text field.
func (a *AddLogMessageRequest) GetText() (value string) {
	if a == nil {
		return
	}
	return a.Text
}

// AddLogMessage invokes method addLogMessage#5f36cfec returning error if any.
func (c *Client) AddLogMessage(ctx context.Context, request *AddLogMessageRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
