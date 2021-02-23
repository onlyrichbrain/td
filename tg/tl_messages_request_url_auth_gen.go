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

// MessagesRequestUrlAuthRequest represents TL type `messages.requestUrlAuth#e33f5613`.
// Get more info about a Seamless Telegram Login authorization request, for more info click here »¹
//
// Links:
//  1) https://core.telegram.org/api/url-authorization
//
// See https://core.telegram.org/method/messages.requestUrlAuth for reference.
type MessagesRequestUrlAuthRequest struct {
	// Peer where the message is located
	Peer InputPeerClass `schemaname:"peer"`
	// The message
	MsgID int `schemaname:"msg_id"`
	// The ID of the button with the authorization request
	ButtonID int `schemaname:"button_id"`
}

// MessagesRequestUrlAuthRequestTypeID is TL type id of MessagesRequestUrlAuthRequest.
const MessagesRequestUrlAuthRequestTypeID = 0xe33f5613

func (r *MessagesRequestUrlAuthRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Peer == nil) {
		return false
	}
	if !(r.MsgID == 0) {
		return false
	}
	if !(r.ButtonID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesRequestUrlAuthRequest) String() string {
	if r == nil {
		return "MessagesRequestUrlAuthRequest(nil)"
	}
	type Alias MessagesRequestUrlAuthRequest
	return fmt.Sprintf("MessagesRequestUrlAuthRequest%+v", Alias(*r))
}

// FillFrom fills MessagesRequestUrlAuthRequest from given interface.
func (r *MessagesRequestUrlAuthRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetMsgID() (value int)
	GetButtonID() (value int)
}) {
	r.Peer = from.GetPeer()
	r.MsgID = from.GetMsgID()
	r.ButtonID = from.GetButtonID()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *MessagesRequestUrlAuthRequest) TypeID() uint32 {
	return MessagesRequestUrlAuthRequestTypeID
}

// SchemaName returns MTProto type name.
func (r *MessagesRequestUrlAuthRequest) SchemaName() string {
	return "messages.requestUrlAuth"
}

// Encode implements bin.Encoder.
func (r *MessagesRequestUrlAuthRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.requestUrlAuth#e33f5613 as nil")
	}
	b.PutID(MessagesRequestUrlAuthRequestTypeID)
	if r.Peer == nil {
		return fmt.Errorf("unable to encode messages.requestUrlAuth#e33f5613: field peer is nil")
	}
	if err := r.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.requestUrlAuth#e33f5613: field peer: %w", err)
	}
	b.PutInt(r.MsgID)
	b.PutInt(r.ButtonID)
	return nil
}

// GetPeer returns value of Peer field.
func (r *MessagesRequestUrlAuthRequest) GetPeer() (value InputPeerClass) {
	return r.Peer
}

// GetMsgID returns value of MsgID field.
func (r *MessagesRequestUrlAuthRequest) GetMsgID() (value int) {
	return r.MsgID
}

// GetButtonID returns value of ButtonID field.
func (r *MessagesRequestUrlAuthRequest) GetButtonID() (value int) {
	return r.ButtonID
}

// Decode implements bin.Decoder.
func (r *MessagesRequestUrlAuthRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.requestUrlAuth#e33f5613 to nil")
	}
	if err := b.ConsumeID(MessagesRequestUrlAuthRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.requestUrlAuth#e33f5613: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestUrlAuth#e33f5613: field peer: %w", err)
		}
		r.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestUrlAuth#e33f5613: field msg_id: %w", err)
		}
		r.MsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestUrlAuth#e33f5613: field button_id: %w", err)
		}
		r.ButtonID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesRequestUrlAuthRequest.
var (
	_ bin.Encoder = &MessagesRequestUrlAuthRequest{}
	_ bin.Decoder = &MessagesRequestUrlAuthRequest{}
)

// MessagesRequestUrlAuth invokes method messages.requestUrlAuth#e33f5613 returning error if any.
// Get more info about a Seamless Telegram Login authorization request, for more info click here »¹
//
// Links:
//  1) https://core.telegram.org/api/url-authorization
//
// See https://core.telegram.org/method/messages.requestUrlAuth for reference.
func (c *Client) MessagesRequestUrlAuth(ctx context.Context, request *MessagesRequestUrlAuthRequest) (UrlAuthResultClass, error) {
	var result UrlAuthResultBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.UrlAuthResult, nil
}
