// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is

// ChatBannedRights represents TL type `chatBannedRights#9f120418`.
// Represents the rights of a normal user in a supergroup/channel/chat¹. In this case, the flags are inverted: if set, a flag does not allow a user to do X.
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/constructor/chatBannedRights for reference.
type ChatBannedRights struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `schemaname:"flags"`
	// If set, does not allow a user to view messages in a supergroup/channel/chat¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	ViewMessages bool `schemaname:"view_messages"`
	// If set, does not allow a user to send messages in a supergroup/chat¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	SendMessages bool `schemaname:"send_messages"`
	// If set, does not allow a user to send any media in a supergroup/chat¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	SendMedia bool `schemaname:"send_media"`
	// If set, does not allow a user to send stickers in a supergroup/chat¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	SendStickers bool `schemaname:"send_stickers"`
	// If set, does not allow a user to send gifs in a supergroup/chat¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	SendGifs bool `schemaname:"send_gifs"`
	// If set, does not allow a user to send games in a supergroup/chat¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	SendGames bool `schemaname:"send_games"`
	// If set, does not allow a user to use inline bots in a supergroup/chat¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	SendInline bool `schemaname:"send_inline"`
	// If set, does not allow a user to embed links in the messages of a supergroup/chat¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	EmbedLinks bool `schemaname:"embed_links"`
	// If set, does not allow a user to send stickers in a supergroup/chat¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	SendPolls bool `schemaname:"send_polls"`
	// If set, does not allow any user to change the description of a supergroup/chat¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	ChangeInfo bool `schemaname:"change_info"`
	// If set, does not allow any user to invite users in a supergroup/chat¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	InviteUsers bool `schemaname:"invite_users"`
	// If set, does not allow any user to pin messages in a supergroup/chat¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	PinMessages bool `schemaname:"pin_messages"`
	// Validity of said permissions (it is considered forever any value less then 30 seconds or more then 366 days).
	UntilDate int `schemaname:"until_date"`
}

// ChatBannedRightsTypeID is TL type id of ChatBannedRights.
const ChatBannedRightsTypeID = 0x9f120418

func (c *ChatBannedRights) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.ViewMessages == false) {
		return false
	}
	if !(c.SendMessages == false) {
		return false
	}
	if !(c.SendMedia == false) {
		return false
	}
	if !(c.SendStickers == false) {
		return false
	}
	if !(c.SendGifs == false) {
		return false
	}
	if !(c.SendGames == false) {
		return false
	}
	if !(c.SendInline == false) {
		return false
	}
	if !(c.EmbedLinks == false) {
		return false
	}
	if !(c.SendPolls == false) {
		return false
	}
	if !(c.ChangeInfo == false) {
		return false
	}
	if !(c.InviteUsers == false) {
		return false
	}
	if !(c.PinMessages == false) {
		return false
	}
	if !(c.UntilDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatBannedRights) String() string {
	if c == nil {
		return "ChatBannedRights(nil)"
	}
	type Alias ChatBannedRights
	return fmt.Sprintf("ChatBannedRights%+v", Alias(*c))
}

// FillFrom fills ChatBannedRights from given interface.
func (c *ChatBannedRights) FillFrom(from interface {
	GetViewMessages() (value bool)
	GetSendMessages() (value bool)
	GetSendMedia() (value bool)
	GetSendStickers() (value bool)
	GetSendGifs() (value bool)
	GetSendGames() (value bool)
	GetSendInline() (value bool)
	GetEmbedLinks() (value bool)
	GetSendPolls() (value bool)
	GetChangeInfo() (value bool)
	GetInviteUsers() (value bool)
	GetPinMessages() (value bool)
	GetUntilDate() (value int)
}) {
	c.ViewMessages = from.GetViewMessages()
	c.SendMessages = from.GetSendMessages()
	c.SendMedia = from.GetSendMedia()
	c.SendStickers = from.GetSendStickers()
	c.SendGifs = from.GetSendGifs()
	c.SendGames = from.GetSendGames()
	c.SendInline = from.GetSendInline()
	c.EmbedLinks = from.GetEmbedLinks()
	c.SendPolls = from.GetSendPolls()
	c.ChangeInfo = from.GetChangeInfo()
	c.InviteUsers = from.GetInviteUsers()
	c.PinMessages = from.GetPinMessages()
	c.UntilDate = from.GetUntilDate()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *ChatBannedRights) TypeID() uint32 {
	return ChatBannedRightsTypeID
}

// SchemaName returns MTProto type name.
func (c *ChatBannedRights) SchemaName() string {
	return "chatBannedRights"
}

// Encode implements bin.Encoder.
func (c *ChatBannedRights) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatBannedRights#9f120418 as nil")
	}
	b.PutID(ChatBannedRightsTypeID)
	if !(c.ViewMessages == false) {
		c.Flags.Set(0)
	}
	if !(c.SendMessages == false) {
		c.Flags.Set(1)
	}
	if !(c.SendMedia == false) {
		c.Flags.Set(2)
	}
	if !(c.SendStickers == false) {
		c.Flags.Set(3)
	}
	if !(c.SendGifs == false) {
		c.Flags.Set(4)
	}
	if !(c.SendGames == false) {
		c.Flags.Set(5)
	}
	if !(c.SendInline == false) {
		c.Flags.Set(6)
	}
	if !(c.EmbedLinks == false) {
		c.Flags.Set(7)
	}
	if !(c.SendPolls == false) {
		c.Flags.Set(8)
	}
	if !(c.ChangeInfo == false) {
		c.Flags.Set(10)
	}
	if !(c.InviteUsers == false) {
		c.Flags.Set(15)
	}
	if !(c.PinMessages == false) {
		c.Flags.Set(17)
	}
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatBannedRights#9f120418: field flags: %w", err)
	}
	b.PutInt(c.UntilDate)
	return nil
}

// SetViewMessages sets value of ViewMessages conditional field.
func (c *ChatBannedRights) SetViewMessages(value bool) {
	if value {
		c.Flags.Set(0)
		c.ViewMessages = true
	} else {
		c.Flags.Unset(0)
		c.ViewMessages = false
	}
}

// GetViewMessages returns value of ViewMessages conditional field.
func (c *ChatBannedRights) GetViewMessages() (value bool) {
	return c.Flags.Has(0)
}

// SetSendMessages sets value of SendMessages conditional field.
func (c *ChatBannedRights) SetSendMessages(value bool) {
	if value {
		c.Flags.Set(1)
		c.SendMessages = true
	} else {
		c.Flags.Unset(1)
		c.SendMessages = false
	}
}

// GetSendMessages returns value of SendMessages conditional field.
func (c *ChatBannedRights) GetSendMessages() (value bool) {
	return c.Flags.Has(1)
}

// SetSendMedia sets value of SendMedia conditional field.
func (c *ChatBannedRights) SetSendMedia(value bool) {
	if value {
		c.Flags.Set(2)
		c.SendMedia = true
	} else {
		c.Flags.Unset(2)
		c.SendMedia = false
	}
}

// GetSendMedia returns value of SendMedia conditional field.
func (c *ChatBannedRights) GetSendMedia() (value bool) {
	return c.Flags.Has(2)
}

// SetSendStickers sets value of SendStickers conditional field.
func (c *ChatBannedRights) SetSendStickers(value bool) {
	if value {
		c.Flags.Set(3)
		c.SendStickers = true
	} else {
		c.Flags.Unset(3)
		c.SendStickers = false
	}
}

// GetSendStickers returns value of SendStickers conditional field.
func (c *ChatBannedRights) GetSendStickers() (value bool) {
	return c.Flags.Has(3)
}

// SetSendGifs sets value of SendGifs conditional field.
func (c *ChatBannedRights) SetSendGifs(value bool) {
	if value {
		c.Flags.Set(4)
		c.SendGifs = true
	} else {
		c.Flags.Unset(4)
		c.SendGifs = false
	}
}

// GetSendGifs returns value of SendGifs conditional field.
func (c *ChatBannedRights) GetSendGifs() (value bool) {
	return c.Flags.Has(4)
}

// SetSendGames sets value of SendGames conditional field.
func (c *ChatBannedRights) SetSendGames(value bool) {
	if value {
		c.Flags.Set(5)
		c.SendGames = true
	} else {
		c.Flags.Unset(5)
		c.SendGames = false
	}
}

// GetSendGames returns value of SendGames conditional field.
func (c *ChatBannedRights) GetSendGames() (value bool) {
	return c.Flags.Has(5)
}

// SetSendInline sets value of SendInline conditional field.
func (c *ChatBannedRights) SetSendInline(value bool) {
	if value {
		c.Flags.Set(6)
		c.SendInline = true
	} else {
		c.Flags.Unset(6)
		c.SendInline = false
	}
}

// GetSendInline returns value of SendInline conditional field.
func (c *ChatBannedRights) GetSendInline() (value bool) {
	return c.Flags.Has(6)
}

// SetEmbedLinks sets value of EmbedLinks conditional field.
func (c *ChatBannedRights) SetEmbedLinks(value bool) {
	if value {
		c.Flags.Set(7)
		c.EmbedLinks = true
	} else {
		c.Flags.Unset(7)
		c.EmbedLinks = false
	}
}

// GetEmbedLinks returns value of EmbedLinks conditional field.
func (c *ChatBannedRights) GetEmbedLinks() (value bool) {
	return c.Flags.Has(7)
}

// SetSendPolls sets value of SendPolls conditional field.
func (c *ChatBannedRights) SetSendPolls(value bool) {
	if value {
		c.Flags.Set(8)
		c.SendPolls = true
	} else {
		c.Flags.Unset(8)
		c.SendPolls = false
	}
}

// GetSendPolls returns value of SendPolls conditional field.
func (c *ChatBannedRights) GetSendPolls() (value bool) {
	return c.Flags.Has(8)
}

// SetChangeInfo sets value of ChangeInfo conditional field.
func (c *ChatBannedRights) SetChangeInfo(value bool) {
	if value {
		c.Flags.Set(10)
		c.ChangeInfo = true
	} else {
		c.Flags.Unset(10)
		c.ChangeInfo = false
	}
}

// GetChangeInfo returns value of ChangeInfo conditional field.
func (c *ChatBannedRights) GetChangeInfo() (value bool) {
	return c.Flags.Has(10)
}

// SetInviteUsers sets value of InviteUsers conditional field.
func (c *ChatBannedRights) SetInviteUsers(value bool) {
	if value {
		c.Flags.Set(15)
		c.InviteUsers = true
	} else {
		c.Flags.Unset(15)
		c.InviteUsers = false
	}
}

// GetInviteUsers returns value of InviteUsers conditional field.
func (c *ChatBannedRights) GetInviteUsers() (value bool) {
	return c.Flags.Has(15)
}

// SetPinMessages sets value of PinMessages conditional field.
func (c *ChatBannedRights) SetPinMessages(value bool) {
	if value {
		c.Flags.Set(17)
		c.PinMessages = true
	} else {
		c.Flags.Unset(17)
		c.PinMessages = false
	}
}

// GetPinMessages returns value of PinMessages conditional field.
func (c *ChatBannedRights) GetPinMessages() (value bool) {
	return c.Flags.Has(17)
}

// GetUntilDate returns value of UntilDate field.
func (c *ChatBannedRights) GetUntilDate() (value int) {
	return c.UntilDate
}

// Decode implements bin.Decoder.
func (c *ChatBannedRights) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatBannedRights#9f120418 to nil")
	}
	if err := b.ConsumeID(ChatBannedRightsTypeID); err != nil {
		return fmt.Errorf("unable to decode chatBannedRights#9f120418: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatBannedRights#9f120418: field flags: %w", err)
		}
	}
	c.ViewMessages = c.Flags.Has(0)
	c.SendMessages = c.Flags.Has(1)
	c.SendMedia = c.Flags.Has(2)
	c.SendStickers = c.Flags.Has(3)
	c.SendGifs = c.Flags.Has(4)
	c.SendGames = c.Flags.Has(5)
	c.SendInline = c.Flags.Has(6)
	c.EmbedLinks = c.Flags.Has(7)
	c.SendPolls = c.Flags.Has(8)
	c.ChangeInfo = c.Flags.Has(10)
	c.InviteUsers = c.Flags.Has(15)
	c.PinMessages = c.Flags.Has(17)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatBannedRights#9f120418: field until_date: %w", err)
		}
		c.UntilDate = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChatBannedRights.
var (
	_ bin.Encoder = &ChatBannedRights{}
	_ bin.Decoder = &ChatBannedRights{}
)
