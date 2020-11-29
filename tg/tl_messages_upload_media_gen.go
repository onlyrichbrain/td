// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// MessagesUploadMediaRequest represents TL type `messages.uploadMedia#519bc2b1`.
type MessagesUploadMediaRequest struct {
	// Peer field of MessagesUploadMediaRequest.
	Peer InputPeerClass
	// Media field of MessagesUploadMediaRequest.
	Media InputMediaClass
}

// MessagesUploadMediaRequestTypeID is TL type id of MessagesUploadMediaRequest.
const MessagesUploadMediaRequestTypeID = 0x519bc2b1

// Encode implements bin.Encoder.
func (u *MessagesUploadMediaRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode messages.uploadMedia#519bc2b1 as nil")
	}
	b.PutID(MessagesUploadMediaRequestTypeID)
	if u.Peer == nil {
		return fmt.Errorf("unable to encode messages.uploadMedia#519bc2b1: field peer is nil")
	}
	if err := u.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.uploadMedia#519bc2b1: field peer: %w", err)
	}
	if u.Media == nil {
		return fmt.Errorf("unable to encode messages.uploadMedia#519bc2b1: field media is nil")
	}
	if err := u.Media.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.uploadMedia#519bc2b1: field media: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *MessagesUploadMediaRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode messages.uploadMedia#519bc2b1 to nil")
	}
	if err := b.ConsumeID(MessagesUploadMediaRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.uploadMedia#519bc2b1: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.uploadMedia#519bc2b1: field peer: %w", err)
		}
		u.Peer = value
	}
	{
		value, err := DecodeInputMedia(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.uploadMedia#519bc2b1: field media: %w", err)
		}
		u.Media = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesUploadMediaRequest.
var (
	_ bin.Encoder = &MessagesUploadMediaRequest{}
	_ bin.Decoder = &MessagesUploadMediaRequest{}
)