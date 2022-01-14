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

// RemoveSavedAnimationRequest represents TL type `removeSavedAnimation#e275a919`.
type RemoveSavedAnimationRequest struct {
	// Animation file to be removed
	Animation InputFileClass
}

// RemoveSavedAnimationRequestTypeID is TL type id of RemoveSavedAnimationRequest.
const RemoveSavedAnimationRequestTypeID = 0xe275a919

// Ensuring interfaces in compile-time for RemoveSavedAnimationRequest.
var (
	_ bin.Encoder     = &RemoveSavedAnimationRequest{}
	_ bin.Decoder     = &RemoveSavedAnimationRequest{}
	_ bin.BareEncoder = &RemoveSavedAnimationRequest{}
	_ bin.BareDecoder = &RemoveSavedAnimationRequest{}
)

func (r *RemoveSavedAnimationRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Animation == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RemoveSavedAnimationRequest) String() string {
	if r == nil {
		return "RemoveSavedAnimationRequest(nil)"
	}
	type Alias RemoveSavedAnimationRequest
	return fmt.Sprintf("RemoveSavedAnimationRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RemoveSavedAnimationRequest) TypeID() uint32 {
	return RemoveSavedAnimationRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*RemoveSavedAnimationRequest) TypeName() string {
	return "removeSavedAnimation"
}

// TypeInfo returns info about TL type.
func (r *RemoveSavedAnimationRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "removeSavedAnimation",
		ID:   RemoveSavedAnimationRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Animation",
			SchemaName: "animation",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RemoveSavedAnimationRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode removeSavedAnimation#e275a919 as nil")
	}
	b.PutID(RemoveSavedAnimationRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RemoveSavedAnimationRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode removeSavedAnimation#e275a919 as nil")
	}
	if r.Animation == nil {
		return fmt.Errorf("unable to encode removeSavedAnimation#e275a919: field animation is nil")
	}
	if err := r.Animation.Encode(b); err != nil {
		return fmt.Errorf("unable to encode removeSavedAnimation#e275a919: field animation: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *RemoveSavedAnimationRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode removeSavedAnimation#e275a919 to nil")
	}
	if err := b.ConsumeID(RemoveSavedAnimationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode removeSavedAnimation#e275a919: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RemoveSavedAnimationRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode removeSavedAnimation#e275a919 to nil")
	}
	{
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode removeSavedAnimation#e275a919: field animation: %w", err)
		}
		r.Animation = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *RemoveSavedAnimationRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode removeSavedAnimation#e275a919 as nil")
	}
	b.ObjStart()
	b.PutID("removeSavedAnimation")
	b.Comma()
	b.FieldStart("animation")
	if r.Animation == nil {
		return fmt.Errorf("unable to encode removeSavedAnimation#e275a919: field animation is nil")
	}
	if err := r.Animation.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode removeSavedAnimation#e275a919: field animation: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *RemoveSavedAnimationRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode removeSavedAnimation#e275a919 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("removeSavedAnimation"); err != nil {
				return fmt.Errorf("unable to decode removeSavedAnimation#e275a919: %w", err)
			}
		case "animation":
			value, err := DecodeTDLibJSONInputFile(b)
			if err != nil {
				return fmt.Errorf("unable to decode removeSavedAnimation#e275a919: field animation: %w", err)
			}
			r.Animation = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetAnimation returns value of Animation field.
func (r *RemoveSavedAnimationRequest) GetAnimation() (value InputFileClass) {
	if r == nil {
		return
	}
	return r.Animation
}

// RemoveSavedAnimation invokes method removeSavedAnimation#e275a919 returning error if any.
func (c *Client) RemoveSavedAnimation(ctx context.Context, animation InputFileClass) error {
	var ok Ok

	request := &RemoveSavedAnimationRequest{
		Animation: animation,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
