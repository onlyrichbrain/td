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

// AddChatMembersRequest represents TL type `addChatMembers#b4f60138`.
type AddChatMembersRequest struct {
	// Chat identifier
	ChatID int64
	// Identifiers of the users to be added to the chat. The maximum number of added users is
	// 20 for supergroups and 100 for channels
	UserIDs []int64
}

// AddChatMembersRequestTypeID is TL type id of AddChatMembersRequest.
const AddChatMembersRequestTypeID = 0xb4f60138

// Ensuring interfaces in compile-time for AddChatMembersRequest.
var (
	_ bin.Encoder     = &AddChatMembersRequest{}
	_ bin.Decoder     = &AddChatMembersRequest{}
	_ bin.BareEncoder = &AddChatMembersRequest{}
	_ bin.BareDecoder = &AddChatMembersRequest{}
)

func (a *AddChatMembersRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.ChatID == 0) {
		return false
	}
	if !(a.UserIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AddChatMembersRequest) String() string {
	if a == nil {
		return "AddChatMembersRequest(nil)"
	}
	type Alias AddChatMembersRequest
	return fmt.Sprintf("AddChatMembersRequest%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AddChatMembersRequest) TypeID() uint32 {
	return AddChatMembersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AddChatMembersRequest) TypeName() string {
	return "addChatMembers"
}

// TypeInfo returns info about TL type.
func (a *AddChatMembersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "addChatMembers",
		ID:   AddChatMembersRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "UserIDs",
			SchemaName: "user_ids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AddChatMembersRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode addChatMembers#b4f60138 as nil")
	}
	b.PutID(AddChatMembersRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AddChatMembersRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode addChatMembers#b4f60138 as nil")
	}
	b.PutInt53(a.ChatID)
	b.PutInt(len(a.UserIDs))
	for _, v := range a.UserIDs {
		b.PutInt53(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AddChatMembersRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode addChatMembers#b4f60138 to nil")
	}
	if err := b.ConsumeID(AddChatMembersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode addChatMembers#b4f60138: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AddChatMembersRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode addChatMembers#b4f60138 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode addChatMembers#b4f60138: field chat_id: %w", err)
		}
		a.ChatID = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode addChatMembers#b4f60138: field user_ids: %w", err)
		}

		if headerLen > 0 {
			a.UserIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode addChatMembers#b4f60138: field user_ids: %w", err)
			}
			a.UserIDs = append(a.UserIDs, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AddChatMembersRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode addChatMembers#b4f60138 as nil")
	}
	b.ObjStart()
	b.PutID("addChatMembers")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(a.ChatID)
	b.Comma()
	b.FieldStart("user_ids")
	b.ArrStart()
	for _, v := range a.UserIDs {
		b.PutInt53(v)
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
func (a *AddChatMembersRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode addChatMembers#b4f60138 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("addChatMembers"); err != nil {
				return fmt.Errorf("unable to decode addChatMembers#b4f60138: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode addChatMembers#b4f60138: field chat_id: %w", err)
			}
			a.ChatID = value
		case "user_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int53()
				if err != nil {
					return fmt.Errorf("unable to decode addChatMembers#b4f60138: field user_ids: %w", err)
				}
				a.UserIDs = append(a.UserIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode addChatMembers#b4f60138: field user_ids: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (a *AddChatMembersRequest) GetChatID() (value int64) {
	if a == nil {
		return
	}
	return a.ChatID
}

// GetUserIDs returns value of UserIDs field.
func (a *AddChatMembersRequest) GetUserIDs() (value []int64) {
	if a == nil {
		return
	}
	return a.UserIDs
}

// AddChatMembers invokes method addChatMembers#b4f60138 returning error if any.
func (c *Client) AddChatMembers(ctx context.Context, request *AddChatMembersRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
