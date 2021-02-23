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

// ContactsGetTopPeersRequest represents TL type `contacts.getTopPeers#d4982db5`.
// Get most used peers
//
// See https://core.telegram.org/method/contacts.getTopPeers for reference.
type ContactsGetTopPeersRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `schemaname:"flags"`
	// Users we've chatted most frequently with
	Correspondents bool `schemaname:"correspondents"`
	// Most used bots
	BotsPm bool `schemaname:"bots_pm"`
	// Most used inline bots
	BotsInline bool `schemaname:"bots_inline"`
	// Most frequently called users
	PhoneCalls bool `schemaname:"phone_calls"`
	// Users to which the users often forwards messages to
	ForwardUsers bool `schemaname:"forward_users"`
	// Chats to which the users often forwards messages to
	ForwardChats bool `schemaname:"forward_chats"`
	// Often-opened groups and supergroups
	Groups bool `schemaname:"groups"`
	// Most frequently visited channels
	Channels bool `schemaname:"channels"`
	// Offset for pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Offset int `schemaname:"offset"`
	// Maximum number of results to return, see pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Limit int `schemaname:"limit"`
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int `schemaname:"hash"`
}

// ContactsGetTopPeersRequestTypeID is TL type id of ContactsGetTopPeersRequest.
const ContactsGetTopPeersRequestTypeID = 0xd4982db5

func (g *ContactsGetTopPeersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Correspondents == false) {
		return false
	}
	if !(g.BotsPm == false) {
		return false
	}
	if !(g.BotsInline == false) {
		return false
	}
	if !(g.PhoneCalls == false) {
		return false
	}
	if !(g.ForwardUsers == false) {
		return false
	}
	if !(g.ForwardChats == false) {
		return false
	}
	if !(g.Groups == false) {
		return false
	}
	if !(g.Channels == false) {
		return false
	}
	if !(g.Offset == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *ContactsGetTopPeersRequest) String() string {
	if g == nil {
		return "ContactsGetTopPeersRequest(nil)"
	}
	type Alias ContactsGetTopPeersRequest
	return fmt.Sprintf("ContactsGetTopPeersRequest%+v", Alias(*g))
}

// FillFrom fills ContactsGetTopPeersRequest from given interface.
func (g *ContactsGetTopPeersRequest) FillFrom(from interface {
	GetCorrespondents() (value bool)
	GetBotsPm() (value bool)
	GetBotsInline() (value bool)
	GetPhoneCalls() (value bool)
	GetForwardUsers() (value bool)
	GetForwardChats() (value bool)
	GetGroups() (value bool)
	GetChannels() (value bool)
	GetOffset() (value int)
	GetLimit() (value int)
	GetHash() (value int)
}) {
	g.Correspondents = from.GetCorrespondents()
	g.BotsPm = from.GetBotsPm()
	g.BotsInline = from.GetBotsInline()
	g.PhoneCalls = from.GetPhoneCalls()
	g.ForwardUsers = from.GetForwardUsers()
	g.ForwardChats = from.GetForwardChats()
	g.Groups = from.GetGroups()
	g.Channels = from.GetChannels()
	g.Offset = from.GetOffset()
	g.Limit = from.GetLimit()
	g.Hash = from.GetHash()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *ContactsGetTopPeersRequest) TypeID() uint32 {
	return ContactsGetTopPeersRequestTypeID
}

// SchemaName returns MTProto type name.
func (g *ContactsGetTopPeersRequest) SchemaName() string {
	return "contacts.getTopPeers"
}

// Encode implements bin.Encoder.
func (g *ContactsGetTopPeersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode contacts.getTopPeers#d4982db5 as nil")
	}
	b.PutID(ContactsGetTopPeersRequestTypeID)
	if !(g.Correspondents == false) {
		g.Flags.Set(0)
	}
	if !(g.BotsPm == false) {
		g.Flags.Set(1)
	}
	if !(g.BotsInline == false) {
		g.Flags.Set(2)
	}
	if !(g.PhoneCalls == false) {
		g.Flags.Set(3)
	}
	if !(g.ForwardUsers == false) {
		g.Flags.Set(4)
	}
	if !(g.ForwardChats == false) {
		g.Flags.Set(5)
	}
	if !(g.Groups == false) {
		g.Flags.Set(10)
	}
	if !(g.Channels == false) {
		g.Flags.Set(15)
	}
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contacts.getTopPeers#d4982db5: field flags: %w", err)
	}
	b.PutInt(g.Offset)
	b.PutInt(g.Limit)
	b.PutInt(g.Hash)
	return nil
}

// SetCorrespondents sets value of Correspondents conditional field.
func (g *ContactsGetTopPeersRequest) SetCorrespondents(value bool) {
	if value {
		g.Flags.Set(0)
		g.Correspondents = true
	} else {
		g.Flags.Unset(0)
		g.Correspondents = false
	}
}

// GetCorrespondents returns value of Correspondents conditional field.
func (g *ContactsGetTopPeersRequest) GetCorrespondents() (value bool) {
	return g.Flags.Has(0)
}

// SetBotsPm sets value of BotsPm conditional field.
func (g *ContactsGetTopPeersRequest) SetBotsPm(value bool) {
	if value {
		g.Flags.Set(1)
		g.BotsPm = true
	} else {
		g.Flags.Unset(1)
		g.BotsPm = false
	}
}

// GetBotsPm returns value of BotsPm conditional field.
func (g *ContactsGetTopPeersRequest) GetBotsPm() (value bool) {
	return g.Flags.Has(1)
}

// SetBotsInline sets value of BotsInline conditional field.
func (g *ContactsGetTopPeersRequest) SetBotsInline(value bool) {
	if value {
		g.Flags.Set(2)
		g.BotsInline = true
	} else {
		g.Flags.Unset(2)
		g.BotsInline = false
	}
}

// GetBotsInline returns value of BotsInline conditional field.
func (g *ContactsGetTopPeersRequest) GetBotsInline() (value bool) {
	return g.Flags.Has(2)
}

// SetPhoneCalls sets value of PhoneCalls conditional field.
func (g *ContactsGetTopPeersRequest) SetPhoneCalls(value bool) {
	if value {
		g.Flags.Set(3)
		g.PhoneCalls = true
	} else {
		g.Flags.Unset(3)
		g.PhoneCalls = false
	}
}

// GetPhoneCalls returns value of PhoneCalls conditional field.
func (g *ContactsGetTopPeersRequest) GetPhoneCalls() (value bool) {
	return g.Flags.Has(3)
}

// SetForwardUsers sets value of ForwardUsers conditional field.
func (g *ContactsGetTopPeersRequest) SetForwardUsers(value bool) {
	if value {
		g.Flags.Set(4)
		g.ForwardUsers = true
	} else {
		g.Flags.Unset(4)
		g.ForwardUsers = false
	}
}

// GetForwardUsers returns value of ForwardUsers conditional field.
func (g *ContactsGetTopPeersRequest) GetForwardUsers() (value bool) {
	return g.Flags.Has(4)
}

// SetForwardChats sets value of ForwardChats conditional field.
func (g *ContactsGetTopPeersRequest) SetForwardChats(value bool) {
	if value {
		g.Flags.Set(5)
		g.ForwardChats = true
	} else {
		g.Flags.Unset(5)
		g.ForwardChats = false
	}
}

// GetForwardChats returns value of ForwardChats conditional field.
func (g *ContactsGetTopPeersRequest) GetForwardChats() (value bool) {
	return g.Flags.Has(5)
}

// SetGroups sets value of Groups conditional field.
func (g *ContactsGetTopPeersRequest) SetGroups(value bool) {
	if value {
		g.Flags.Set(10)
		g.Groups = true
	} else {
		g.Flags.Unset(10)
		g.Groups = false
	}
}

// GetGroups returns value of Groups conditional field.
func (g *ContactsGetTopPeersRequest) GetGroups() (value bool) {
	return g.Flags.Has(10)
}

// SetChannels sets value of Channels conditional field.
func (g *ContactsGetTopPeersRequest) SetChannels(value bool) {
	if value {
		g.Flags.Set(15)
		g.Channels = true
	} else {
		g.Flags.Unset(15)
		g.Channels = false
	}
}

// GetChannels returns value of Channels conditional field.
func (g *ContactsGetTopPeersRequest) GetChannels() (value bool) {
	return g.Flags.Has(15)
}

// GetOffset returns value of Offset field.
func (g *ContactsGetTopPeersRequest) GetOffset() (value int) {
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *ContactsGetTopPeersRequest) GetLimit() (value int) {
	return g.Limit
}

// GetHash returns value of Hash field.
func (g *ContactsGetTopPeersRequest) GetHash() (value int) {
	return g.Hash
}

// Decode implements bin.Decoder.
func (g *ContactsGetTopPeersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode contacts.getTopPeers#d4982db5 to nil")
	}
	if err := b.ConsumeID(ContactsGetTopPeersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.getTopPeers#d4982db5: %w", err)
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode contacts.getTopPeers#d4982db5: field flags: %w", err)
		}
	}
	g.Correspondents = g.Flags.Has(0)
	g.BotsPm = g.Flags.Has(1)
	g.BotsInline = g.Flags.Has(2)
	g.PhoneCalls = g.Flags.Has(3)
	g.ForwardUsers = g.Flags.Has(4)
	g.ForwardChats = g.Flags.Has(5)
	g.Groups = g.Flags.Has(10)
	g.Channels = g.Flags.Has(15)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.getTopPeers#d4982db5: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.getTopPeers#d4982db5: field limit: %w", err)
		}
		g.Limit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.getTopPeers#d4982db5: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsGetTopPeersRequest.
var (
	_ bin.Encoder = &ContactsGetTopPeersRequest{}
	_ bin.Decoder = &ContactsGetTopPeersRequest{}
)

// ContactsGetTopPeers invokes method contacts.getTopPeers#d4982db5 returning error if any.
// Get most used peers
//
// Possible errors:
//  400 TYPES_EMPTY: No top peer type was provided
//
// See https://core.telegram.org/method/contacts.getTopPeers for reference.
func (c *Client) ContactsGetTopPeers(ctx context.Context, request *ContactsGetTopPeersRequest) (ContactsTopPeersClass, error) {
	var result ContactsTopPeersBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.TopPeers, nil
}
