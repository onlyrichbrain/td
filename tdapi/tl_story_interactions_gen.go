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

// StoryInteractions represents TL type `storyInteractions#f9f5d50f`.
type StoryInteractions struct {
	// Approximate total number of interactions found
	TotalCount int32
	// Approximate total number of found forwards and reposts; always 0 for chat stories
	TotalForwardCount int32
	// Approximate total number of found reactions; always 0 for chat stories
	TotalReactionCount int32
	// List of story interactions
	Interactions []StoryInteraction
	// The offset for the next request. If empty, then there are no more results
	NextOffset string
}

// StoryInteractionsTypeID is TL type id of StoryInteractions.
const StoryInteractionsTypeID = 0xf9f5d50f

// Ensuring interfaces in compile-time for StoryInteractions.
var (
	_ bin.Encoder     = &StoryInteractions{}
	_ bin.Decoder     = &StoryInteractions{}
	_ bin.BareEncoder = &StoryInteractions{}
	_ bin.BareDecoder = &StoryInteractions{}
)

func (s *StoryInteractions) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.TotalCount == 0) {
		return false
	}
	if !(s.TotalForwardCount == 0) {
		return false
	}
	if !(s.TotalReactionCount == 0) {
		return false
	}
	if !(s.Interactions == nil) {
		return false
	}
	if !(s.NextOffset == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StoryInteractions) String() string {
	if s == nil {
		return "StoryInteractions(nil)"
	}
	type Alias StoryInteractions
	return fmt.Sprintf("StoryInteractions%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoryInteractions) TypeID() uint32 {
	return StoryInteractionsTypeID
}

// TypeName returns name of type in TL schema.
func (*StoryInteractions) TypeName() string {
	return "storyInteractions"
}

// TypeInfo returns info about TL type.
func (s *StoryInteractions) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storyInteractions",
		ID:   StoryInteractionsTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "TotalCount",
			SchemaName: "total_count",
		},
		{
			Name:       "TotalForwardCount",
			SchemaName: "total_forward_count",
		},
		{
			Name:       "TotalReactionCount",
			SchemaName: "total_reaction_count",
		},
		{
			Name:       "Interactions",
			SchemaName: "interactions",
		},
		{
			Name:       "NextOffset",
			SchemaName: "next_offset",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StoryInteractions) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyInteractions#f9f5d50f as nil")
	}
	b.PutID(StoryInteractionsTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StoryInteractions) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyInteractions#f9f5d50f as nil")
	}
	b.PutInt32(s.TotalCount)
	b.PutInt32(s.TotalForwardCount)
	b.PutInt32(s.TotalReactionCount)
	b.PutInt(len(s.Interactions))
	for idx, v := range s.Interactions {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare storyInteractions#f9f5d50f: field interactions element with index %d: %w", idx, err)
		}
	}
	b.PutString(s.NextOffset)
	return nil
}

// Decode implements bin.Decoder.
func (s *StoryInteractions) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyInteractions#f9f5d50f to nil")
	}
	if err := b.ConsumeID(StoryInteractionsTypeID); err != nil {
		return fmt.Errorf("unable to decode storyInteractions#f9f5d50f: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StoryInteractions) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyInteractions#f9f5d50f to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode storyInteractions#f9f5d50f: field total_count: %w", err)
		}
		s.TotalCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode storyInteractions#f9f5d50f: field total_forward_count: %w", err)
		}
		s.TotalForwardCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode storyInteractions#f9f5d50f: field total_reaction_count: %w", err)
		}
		s.TotalReactionCount = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode storyInteractions#f9f5d50f: field interactions: %w", err)
		}

		if headerLen > 0 {
			s.Interactions = make([]StoryInteraction, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value StoryInteraction
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare storyInteractions#f9f5d50f: field interactions: %w", err)
			}
			s.Interactions = append(s.Interactions, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode storyInteractions#f9f5d50f: field next_offset: %w", err)
		}
		s.NextOffset = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StoryInteractions) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode storyInteractions#f9f5d50f as nil")
	}
	b.ObjStart()
	b.PutID("storyInteractions")
	b.Comma()
	b.FieldStart("total_count")
	b.PutInt32(s.TotalCount)
	b.Comma()
	b.FieldStart("total_forward_count")
	b.PutInt32(s.TotalForwardCount)
	b.Comma()
	b.FieldStart("total_reaction_count")
	b.PutInt32(s.TotalReactionCount)
	b.Comma()
	b.FieldStart("interactions")
	b.ArrStart()
	for idx, v := range s.Interactions {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode storyInteractions#f9f5d50f: field interactions element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("next_offset")
	b.PutString(s.NextOffset)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StoryInteractions) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode storyInteractions#f9f5d50f to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("storyInteractions"); err != nil {
				return fmt.Errorf("unable to decode storyInteractions#f9f5d50f: %w", err)
			}
		case "total_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode storyInteractions#f9f5d50f: field total_count: %w", err)
			}
			s.TotalCount = value
		case "total_forward_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode storyInteractions#f9f5d50f: field total_forward_count: %w", err)
			}
			s.TotalForwardCount = value
		case "total_reaction_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode storyInteractions#f9f5d50f: field total_reaction_count: %w", err)
			}
			s.TotalReactionCount = value
		case "interactions":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value StoryInteraction
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode storyInteractions#f9f5d50f: field interactions: %w", err)
				}
				s.Interactions = append(s.Interactions, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode storyInteractions#f9f5d50f: field interactions: %w", err)
			}
		case "next_offset":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode storyInteractions#f9f5d50f: field next_offset: %w", err)
			}
			s.NextOffset = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetTotalCount returns value of TotalCount field.
func (s *StoryInteractions) GetTotalCount() (value int32) {
	if s == nil {
		return
	}
	return s.TotalCount
}

// GetTotalForwardCount returns value of TotalForwardCount field.
func (s *StoryInteractions) GetTotalForwardCount() (value int32) {
	if s == nil {
		return
	}
	return s.TotalForwardCount
}

// GetTotalReactionCount returns value of TotalReactionCount field.
func (s *StoryInteractions) GetTotalReactionCount() (value int32) {
	if s == nil {
		return
	}
	return s.TotalReactionCount
}

// GetInteractions returns value of Interactions field.
func (s *StoryInteractions) GetInteractions() (value []StoryInteraction) {
	if s == nil {
		return
	}
	return s.Interactions
}

// GetNextOffset returns value of NextOffset field.
func (s *StoryInteractions) GetNextOffset() (value string) {
	if s == nil {
		return
	}
	return s.NextOffset
}