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

// GetRecommendedChatFiltersRequest represents TL type `getRecommendedChatFilters#d18b70e6`.
type GetRecommendedChatFiltersRequest struct {
}

// GetRecommendedChatFiltersRequestTypeID is TL type id of GetRecommendedChatFiltersRequest.
const GetRecommendedChatFiltersRequestTypeID = 0xd18b70e6

// Ensuring interfaces in compile-time for GetRecommendedChatFiltersRequest.
var (
	_ bin.Encoder     = &GetRecommendedChatFiltersRequest{}
	_ bin.Decoder     = &GetRecommendedChatFiltersRequest{}
	_ bin.BareEncoder = &GetRecommendedChatFiltersRequest{}
	_ bin.BareDecoder = &GetRecommendedChatFiltersRequest{}
)

func (g *GetRecommendedChatFiltersRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetRecommendedChatFiltersRequest) String() string {
	if g == nil {
		return "GetRecommendedChatFiltersRequest(nil)"
	}
	type Alias GetRecommendedChatFiltersRequest
	return fmt.Sprintf("GetRecommendedChatFiltersRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetRecommendedChatFiltersRequest) TypeID() uint32 {
	return GetRecommendedChatFiltersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetRecommendedChatFiltersRequest) TypeName() string {
	return "getRecommendedChatFilters"
}

// TypeInfo returns info about TL type.
func (g *GetRecommendedChatFiltersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getRecommendedChatFilters",
		ID:   GetRecommendedChatFiltersRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetRecommendedChatFiltersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getRecommendedChatFilters#d18b70e6 as nil")
	}
	b.PutID(GetRecommendedChatFiltersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetRecommendedChatFiltersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getRecommendedChatFilters#d18b70e6 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetRecommendedChatFiltersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getRecommendedChatFilters#d18b70e6 to nil")
	}
	if err := b.ConsumeID(GetRecommendedChatFiltersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getRecommendedChatFilters#d18b70e6: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetRecommendedChatFiltersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getRecommendedChatFilters#d18b70e6 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetRecommendedChatFiltersRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getRecommendedChatFilters#d18b70e6 as nil")
	}
	b.ObjStart()
	b.PutID("getRecommendedChatFilters")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetRecommendedChatFiltersRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getRecommendedChatFilters#d18b70e6 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getRecommendedChatFilters"); err != nil {
				return fmt.Errorf("unable to decode getRecommendedChatFilters#d18b70e6: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetRecommendedChatFilters invokes method getRecommendedChatFilters#d18b70e6 returning error if any.
func (c *Client) GetRecommendedChatFilters(ctx context.Context) (*RecommendedChatFilters, error) {
	var result RecommendedChatFilters

	request := &GetRecommendedChatFiltersRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
