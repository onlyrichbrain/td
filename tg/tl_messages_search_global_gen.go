// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// MessagesSearchGlobalRequest represents TL type `messages.searchGlobal#4bc6589a`.
// Search for messages and peers globally
//
// See https://core.telegram.org/method/messages.searchGlobal for reference.
type MessagesSearchGlobalRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// If set, only returns results from channels (used in the global channel search tab
	// »¹).
	//
	// Links:
	//  1) https://core.telegram.org/api/search#global-search
	BroadcastsOnly bool
	// Whether to search only in groups
	GroupsOnly bool
	// Whether to search only in private chats
	UsersOnly bool
	// Peer folder ID, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders#peer-folders
	//
	// Use SetFolderID and GetFolderID helpers.
	FolderID int
	// Query
	Q string
	// Global search filter
	Filter MessagesFilterClass
	// If a positive value was specified, the method will return only messages with date
	// bigger than min_date
	MinDate int
	// If a positive value was transferred, the method will return only messages with date
	// smaller than max_date
	MaxDate int
	// Initially 0, then set to the next_rate parameter of messages.messagesSlice¹
	//
	// Links:
	//  1) https://core.telegram.org/constructor/messages.messagesSlice
	OffsetRate int
	// Offsets for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	OffsetPeer InputPeerClass
	// Offsets for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	OffsetID int
	// Offsets for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Limit int
}

// MessagesSearchGlobalRequestTypeID is TL type id of MessagesSearchGlobalRequest.
const MessagesSearchGlobalRequestTypeID = 0x4bc6589a

// Ensuring interfaces in compile-time for MessagesSearchGlobalRequest.
var (
	_ bin.Encoder     = &MessagesSearchGlobalRequest{}
	_ bin.Decoder     = &MessagesSearchGlobalRequest{}
	_ bin.BareEncoder = &MessagesSearchGlobalRequest{}
	_ bin.BareDecoder = &MessagesSearchGlobalRequest{}
)

func (s *MessagesSearchGlobalRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.BroadcastsOnly == false) {
		return false
	}
	if !(s.GroupsOnly == false) {
		return false
	}
	if !(s.UsersOnly == false) {
		return false
	}
	if !(s.FolderID == 0) {
		return false
	}
	if !(s.Q == "") {
		return false
	}
	if !(s.Filter == nil) {
		return false
	}
	if !(s.MinDate == 0) {
		return false
	}
	if !(s.MaxDate == 0) {
		return false
	}
	if !(s.OffsetRate == 0) {
		return false
	}
	if !(s.OffsetPeer == nil) {
		return false
	}
	if !(s.OffsetID == 0) {
		return false
	}
	if !(s.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSearchGlobalRequest) String() string {
	if s == nil {
		return "MessagesSearchGlobalRequest(nil)"
	}
	type Alias MessagesSearchGlobalRequest
	return fmt.Sprintf("MessagesSearchGlobalRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSearchGlobalRequest from given interface.
func (s *MessagesSearchGlobalRequest) FillFrom(from interface {
	GetBroadcastsOnly() (value bool)
	GetGroupsOnly() (value bool)
	GetUsersOnly() (value bool)
	GetFolderID() (value int, ok bool)
	GetQ() (value string)
	GetFilter() (value MessagesFilterClass)
	GetMinDate() (value int)
	GetMaxDate() (value int)
	GetOffsetRate() (value int)
	GetOffsetPeer() (value InputPeerClass)
	GetOffsetID() (value int)
	GetLimit() (value int)
}) {
	s.BroadcastsOnly = from.GetBroadcastsOnly()
	s.GroupsOnly = from.GetGroupsOnly()
	s.UsersOnly = from.GetUsersOnly()
	if val, ok := from.GetFolderID(); ok {
		s.FolderID = val
	}

	s.Q = from.GetQ()
	s.Filter = from.GetFilter()
	s.MinDate = from.GetMinDate()
	s.MaxDate = from.GetMaxDate()
	s.OffsetRate = from.GetOffsetRate()
	s.OffsetPeer = from.GetOffsetPeer()
	s.OffsetID = from.GetOffsetID()
	s.Limit = from.GetLimit()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSearchGlobalRequest) TypeID() uint32 {
	return MessagesSearchGlobalRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSearchGlobalRequest) TypeName() string {
	return "messages.searchGlobal"
}

// TypeInfo returns info about TL type.
func (s *MessagesSearchGlobalRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.searchGlobal",
		ID:   MessagesSearchGlobalRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BroadcastsOnly",
			SchemaName: "broadcasts_only",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "GroupsOnly",
			SchemaName: "groups_only",
			Null:       !s.Flags.Has(2),
		},
		{
			Name:       "UsersOnly",
			SchemaName: "users_only",
			Null:       !s.Flags.Has(3),
		},
		{
			Name:       "FolderID",
			SchemaName: "folder_id",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "Q",
			SchemaName: "q",
		},
		{
			Name:       "Filter",
			SchemaName: "filter",
		},
		{
			Name:       "MinDate",
			SchemaName: "min_date",
		},
		{
			Name:       "MaxDate",
			SchemaName: "max_date",
		},
		{
			Name:       "OffsetRate",
			SchemaName: "offset_rate",
		},
		{
			Name:       "OffsetPeer",
			SchemaName: "offset_peer",
		},
		{
			Name:       "OffsetID",
			SchemaName: "offset_id",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *MessagesSearchGlobalRequest) SetFlags() {
	if !(s.BroadcastsOnly == false) {
		s.Flags.Set(1)
	}
	if !(s.GroupsOnly == false) {
		s.Flags.Set(2)
	}
	if !(s.UsersOnly == false) {
		s.Flags.Set(3)
	}
	if !(s.FolderID == 0) {
		s.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (s *MessagesSearchGlobalRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.searchGlobal#4bc6589a as nil")
	}
	b.PutID(MessagesSearchGlobalRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSearchGlobalRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.searchGlobal#4bc6589a as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.searchGlobal#4bc6589a: field flags: %w", err)
	}
	if s.Flags.Has(0) {
		b.PutInt(s.FolderID)
	}
	b.PutString(s.Q)
	if s.Filter == nil {
		return fmt.Errorf("unable to encode messages.searchGlobal#4bc6589a: field filter is nil")
	}
	if err := s.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.searchGlobal#4bc6589a: field filter: %w", err)
	}
	b.PutInt(s.MinDate)
	b.PutInt(s.MaxDate)
	b.PutInt(s.OffsetRate)
	if s.OffsetPeer == nil {
		return fmt.Errorf("unable to encode messages.searchGlobal#4bc6589a: field offset_peer is nil")
	}
	if err := s.OffsetPeer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.searchGlobal#4bc6589a: field offset_peer: %w", err)
	}
	b.PutInt(s.OffsetID)
	b.PutInt(s.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSearchGlobalRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.searchGlobal#4bc6589a to nil")
	}
	if err := b.ConsumeID(MessagesSearchGlobalRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.searchGlobal#4bc6589a: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSearchGlobalRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.searchGlobal#4bc6589a to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.searchGlobal#4bc6589a: field flags: %w", err)
		}
	}
	s.BroadcastsOnly = s.Flags.Has(1)
	s.GroupsOnly = s.Flags.Has(2)
	s.UsersOnly = s.Flags.Has(3)
	if s.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchGlobal#4bc6589a: field folder_id: %w", err)
		}
		s.FolderID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchGlobal#4bc6589a: field q: %w", err)
		}
		s.Q = value
	}
	{
		value, err := DecodeMessagesFilter(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchGlobal#4bc6589a: field filter: %w", err)
		}
		s.Filter = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchGlobal#4bc6589a: field min_date: %w", err)
		}
		s.MinDate = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchGlobal#4bc6589a: field max_date: %w", err)
		}
		s.MaxDate = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchGlobal#4bc6589a: field offset_rate: %w", err)
		}
		s.OffsetRate = value
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchGlobal#4bc6589a: field offset_peer: %w", err)
		}
		s.OffsetPeer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchGlobal#4bc6589a: field offset_id: %w", err)
		}
		s.OffsetID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchGlobal#4bc6589a: field limit: %w", err)
		}
		s.Limit = value
	}
	return nil
}

// SetBroadcastsOnly sets value of BroadcastsOnly conditional field.
func (s *MessagesSearchGlobalRequest) SetBroadcastsOnly(value bool) {
	if value {
		s.Flags.Set(1)
		s.BroadcastsOnly = true
	} else {
		s.Flags.Unset(1)
		s.BroadcastsOnly = false
	}
}

// GetBroadcastsOnly returns value of BroadcastsOnly conditional field.
func (s *MessagesSearchGlobalRequest) GetBroadcastsOnly() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(1)
}

// SetGroupsOnly sets value of GroupsOnly conditional field.
func (s *MessagesSearchGlobalRequest) SetGroupsOnly(value bool) {
	if value {
		s.Flags.Set(2)
		s.GroupsOnly = true
	} else {
		s.Flags.Unset(2)
		s.GroupsOnly = false
	}
}

// GetGroupsOnly returns value of GroupsOnly conditional field.
func (s *MessagesSearchGlobalRequest) GetGroupsOnly() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(2)
}

// SetUsersOnly sets value of UsersOnly conditional field.
func (s *MessagesSearchGlobalRequest) SetUsersOnly(value bool) {
	if value {
		s.Flags.Set(3)
		s.UsersOnly = true
	} else {
		s.Flags.Unset(3)
		s.UsersOnly = false
	}
}

// GetUsersOnly returns value of UsersOnly conditional field.
func (s *MessagesSearchGlobalRequest) GetUsersOnly() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(3)
}

// SetFolderID sets value of FolderID conditional field.
func (s *MessagesSearchGlobalRequest) SetFolderID(value int) {
	s.Flags.Set(0)
	s.FolderID = value
}

// GetFolderID returns value of FolderID conditional field and
// boolean which is true if field was set.
func (s *MessagesSearchGlobalRequest) GetFolderID() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.FolderID, true
}

// GetQ returns value of Q field.
func (s *MessagesSearchGlobalRequest) GetQ() (value string) {
	if s == nil {
		return
	}
	return s.Q
}

// GetFilter returns value of Filter field.
func (s *MessagesSearchGlobalRequest) GetFilter() (value MessagesFilterClass) {
	if s == nil {
		return
	}
	return s.Filter
}

// GetMinDate returns value of MinDate field.
func (s *MessagesSearchGlobalRequest) GetMinDate() (value int) {
	if s == nil {
		return
	}
	return s.MinDate
}

// GetMaxDate returns value of MaxDate field.
func (s *MessagesSearchGlobalRequest) GetMaxDate() (value int) {
	if s == nil {
		return
	}
	return s.MaxDate
}

// GetOffsetRate returns value of OffsetRate field.
func (s *MessagesSearchGlobalRequest) GetOffsetRate() (value int) {
	if s == nil {
		return
	}
	return s.OffsetRate
}

// GetOffsetPeer returns value of OffsetPeer field.
func (s *MessagesSearchGlobalRequest) GetOffsetPeer() (value InputPeerClass) {
	if s == nil {
		return
	}
	return s.OffsetPeer
}

// GetOffsetID returns value of OffsetID field.
func (s *MessagesSearchGlobalRequest) GetOffsetID() (value int) {
	if s == nil {
		return
	}
	return s.OffsetID
}

// GetLimit returns value of Limit field.
func (s *MessagesSearchGlobalRequest) GetLimit() (value int) {
	if s == nil {
		return
	}
	return s.Limit
}

// MessagesSearchGlobal invokes method messages.searchGlobal#4bc6589a returning error if any.
// Search for messages and peers globally
//
// Possible errors:
//
//	400 FOLDER_ID_INVALID: Invalid folder ID.
//	400 INPUT_FILTER_INVALID: The specified filter is invalid.
//	400 SEARCH_QUERY_EMPTY: The search query is empty.
//
// See https://core.telegram.org/method/messages.searchGlobal for reference.
func (c *Client) MessagesSearchGlobal(ctx context.Context, request *MessagesSearchGlobalRequest) (MessagesMessagesClass, error) {
	var result MessagesMessagesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Messages, nil
}
