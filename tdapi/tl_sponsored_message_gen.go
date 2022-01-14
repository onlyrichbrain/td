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

// SponsoredMessage represents TL type `sponsoredMessage#ad7c99d8`.
type SponsoredMessage struct {
	// Message identifier; unique for the chat to which the sponsored message belongs among
	// both ordinary and sponsored messages
	MessageID int64
	// Sponsor chat identifier; 0 if the sponsor chat is accessible through an invite link
	SponsorChatID int64
	// Information about the sponsor chat; may be null unless sponsor_chat_id == 0
	SponsorChatInfo ChatInviteLinkInfo
	// An internal link to be opened when the sponsored message is clicked; may be null. If
	// null, the sponsor chat needs to be opened instead
	Link InternalLinkTypeClass
	// Content of the message. Currently, can be only of the type messageText
	Content MessageContentClass
}

// SponsoredMessageTypeID is TL type id of SponsoredMessage.
const SponsoredMessageTypeID = 0xad7c99d8

// Ensuring interfaces in compile-time for SponsoredMessage.
var (
	_ bin.Encoder     = &SponsoredMessage{}
	_ bin.Decoder     = &SponsoredMessage{}
	_ bin.BareEncoder = &SponsoredMessage{}
	_ bin.BareDecoder = &SponsoredMessage{}
)

func (s *SponsoredMessage) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.MessageID == 0) {
		return false
	}
	if !(s.SponsorChatID == 0) {
		return false
	}
	if !(s.SponsorChatInfo.Zero()) {
		return false
	}
	if !(s.Link == nil) {
		return false
	}
	if !(s.Content == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SponsoredMessage) String() string {
	if s == nil {
		return "SponsoredMessage(nil)"
	}
	type Alias SponsoredMessage
	return fmt.Sprintf("SponsoredMessage%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SponsoredMessage) TypeID() uint32 {
	return SponsoredMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*SponsoredMessage) TypeName() string {
	return "sponsoredMessage"
}

// TypeInfo returns info about TL type.
func (s *SponsoredMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "sponsoredMessage",
		ID:   SponsoredMessageTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "SponsorChatID",
			SchemaName: "sponsor_chat_id",
		},
		{
			Name:       "SponsorChatInfo",
			SchemaName: "sponsor_chat_info",
		},
		{
			Name:       "Link",
			SchemaName: "link",
		},
		{
			Name:       "Content",
			SchemaName: "content",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SponsoredMessage) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sponsoredMessage#ad7c99d8 as nil")
	}
	b.PutID(SponsoredMessageTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SponsoredMessage) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sponsoredMessage#ad7c99d8 as nil")
	}
	b.PutInt53(s.MessageID)
	b.PutInt53(s.SponsorChatID)
	if err := s.SponsorChatInfo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sponsoredMessage#ad7c99d8: field sponsor_chat_info: %w", err)
	}
	if s.Link == nil {
		return fmt.Errorf("unable to encode sponsoredMessage#ad7c99d8: field link is nil")
	}
	if err := s.Link.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sponsoredMessage#ad7c99d8: field link: %w", err)
	}
	if s.Content == nil {
		return fmt.Errorf("unable to encode sponsoredMessage#ad7c99d8: field content is nil")
	}
	if err := s.Content.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sponsoredMessage#ad7c99d8: field content: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SponsoredMessage) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sponsoredMessage#ad7c99d8 to nil")
	}
	if err := b.ConsumeID(SponsoredMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode sponsoredMessage#ad7c99d8: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SponsoredMessage) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sponsoredMessage#ad7c99d8 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#ad7c99d8: field message_id: %w", err)
		}
		s.MessageID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#ad7c99d8: field sponsor_chat_id: %w", err)
		}
		s.SponsorChatID = value
	}
	{
		if err := s.SponsorChatInfo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#ad7c99d8: field sponsor_chat_info: %w", err)
		}
	}
	{
		value, err := DecodeInternalLinkType(b)
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#ad7c99d8: field link: %w", err)
		}
		s.Link = value
	}
	{
		value, err := DecodeMessageContent(b)
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#ad7c99d8: field content: %w", err)
		}
		s.Content = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SponsoredMessage) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode sponsoredMessage#ad7c99d8 as nil")
	}
	b.ObjStart()
	b.PutID("sponsoredMessage")
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(s.MessageID)
	b.Comma()
	b.FieldStart("sponsor_chat_id")
	b.PutInt53(s.SponsorChatID)
	b.Comma()
	b.FieldStart("sponsor_chat_info")
	if err := s.SponsorChatInfo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sponsoredMessage#ad7c99d8: field sponsor_chat_info: %w", err)
	}
	b.Comma()
	b.FieldStart("link")
	if s.Link == nil {
		return fmt.Errorf("unable to encode sponsoredMessage#ad7c99d8: field link is nil")
	}
	if err := s.Link.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sponsoredMessage#ad7c99d8: field link: %w", err)
	}
	b.Comma()
	b.FieldStart("content")
	if s.Content == nil {
		return fmt.Errorf("unable to encode sponsoredMessage#ad7c99d8: field content is nil")
	}
	if err := s.Content.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sponsoredMessage#ad7c99d8: field content: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SponsoredMessage) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode sponsoredMessage#ad7c99d8 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("sponsoredMessage"); err != nil {
				return fmt.Errorf("unable to decode sponsoredMessage#ad7c99d8: %w", err)
			}
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode sponsoredMessage#ad7c99d8: field message_id: %w", err)
			}
			s.MessageID = value
		case "sponsor_chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode sponsoredMessage#ad7c99d8: field sponsor_chat_id: %w", err)
			}
			s.SponsorChatID = value
		case "sponsor_chat_info":
			if err := s.SponsorChatInfo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode sponsoredMessage#ad7c99d8: field sponsor_chat_info: %w", err)
			}
		case "link":
			value, err := DecodeTDLibJSONInternalLinkType(b)
			if err != nil {
				return fmt.Errorf("unable to decode sponsoredMessage#ad7c99d8: field link: %w", err)
			}
			s.Link = value
		case "content":
			value, err := DecodeTDLibJSONMessageContent(b)
			if err != nil {
				return fmt.Errorf("unable to decode sponsoredMessage#ad7c99d8: field content: %w", err)
			}
			s.Content = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetMessageID returns value of MessageID field.
func (s *SponsoredMessage) GetMessageID() (value int64) {
	if s == nil {
		return
	}
	return s.MessageID
}

// GetSponsorChatID returns value of SponsorChatID field.
func (s *SponsoredMessage) GetSponsorChatID() (value int64) {
	if s == nil {
		return
	}
	return s.SponsorChatID
}

// GetSponsorChatInfo returns value of SponsorChatInfo field.
func (s *SponsoredMessage) GetSponsorChatInfo() (value ChatInviteLinkInfo) {
	if s == nil {
		return
	}
	return s.SponsorChatInfo
}

// GetLink returns value of Link field.
func (s *SponsoredMessage) GetLink() (value InternalLinkTypeClass) {
	if s == nil {
		return
	}
	return s.Link
}

// GetContent returns value of Content field.
func (s *SponsoredMessage) GetContent() (value MessageContentClass) {
	if s == nil {
		return
	}
	return s.Content
}
