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

// MessagesRequestMainWebViewRequest represents TL type `messages.requestMainWebView#c9e01e7b`.
// Open a Main Mini App¹.
//
// Links:
//  1. https://core.telegram.org/api/bots/webapps#main-mini-apps
//
// See https://core.telegram.org/method/messages.requestMainWebView for reference.
type MessagesRequestMainWebViewRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// If set, requests to open the mini app in compact mode (as opposed to normal or
	// fullscreen mode). Must be set if the mode parameter of the Main Mini App link¹ is
	// equal to compact.
	//
	// Links:
	//  1) https://core.telegram.org/api/links#main-mini-app-links
	Compact bool
	// If set, requests to open the mini app in fullscreen mode (as opposed to compact or
	// normal mode). Must be set if the mode parameter of the Main Mini App link¹ is equal
	// to fullscreen.
	//
	// Links:
	//  1) https://core.telegram.org/api/links#main-mini-app-links
	Fullscreen bool
	// Currently open chat, may be inputPeerEmpty¹ if no chat is currently open.
	//
	// Links:
	//  1) https://core.telegram.org/constructor/inputPeerEmpty
	Peer InputPeerClass
	// Bot that owns the main mini app.
	Bot InputUserClass
	// Start parameter, if opening from a Main Mini App link »¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/links#main-mini-app-links
	//
	// Use SetStartParam and GetStartParam helpers.
	StartParam string
	// Theme parameters »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/bots/webapps#theme-parameters
	//
	// Use SetThemeParams and GetThemeParams helpers.
	ThemeParams DataJSON
	// Short name of the application; 0-64 English letters, digits, and underscores
	Platform string
}

// MessagesRequestMainWebViewRequestTypeID is TL type id of MessagesRequestMainWebViewRequest.
const MessagesRequestMainWebViewRequestTypeID = 0xc9e01e7b

// Ensuring interfaces in compile-time for MessagesRequestMainWebViewRequest.
var (
	_ bin.Encoder     = &MessagesRequestMainWebViewRequest{}
	_ bin.Decoder     = &MessagesRequestMainWebViewRequest{}
	_ bin.BareEncoder = &MessagesRequestMainWebViewRequest{}
	_ bin.BareDecoder = &MessagesRequestMainWebViewRequest{}
)

func (r *MessagesRequestMainWebViewRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.Compact == false) {
		return false
	}
	if !(r.Fullscreen == false) {
		return false
	}
	if !(r.Peer == nil) {
		return false
	}
	if !(r.Bot == nil) {
		return false
	}
	if !(r.StartParam == "") {
		return false
	}
	if !(r.ThemeParams.Zero()) {
		return false
	}
	if !(r.Platform == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesRequestMainWebViewRequest) String() string {
	if r == nil {
		return "MessagesRequestMainWebViewRequest(nil)"
	}
	type Alias MessagesRequestMainWebViewRequest
	return fmt.Sprintf("MessagesRequestMainWebViewRequest%+v", Alias(*r))
}

// FillFrom fills MessagesRequestMainWebViewRequest from given interface.
func (r *MessagesRequestMainWebViewRequest) FillFrom(from interface {
	GetCompact() (value bool)
	GetFullscreen() (value bool)
	GetPeer() (value InputPeerClass)
	GetBot() (value InputUserClass)
	GetStartParam() (value string, ok bool)
	GetThemeParams() (value DataJSON, ok bool)
	GetPlatform() (value string)
}) {
	r.Compact = from.GetCompact()
	r.Fullscreen = from.GetFullscreen()
	r.Peer = from.GetPeer()
	r.Bot = from.GetBot()
	if val, ok := from.GetStartParam(); ok {
		r.StartParam = val
	}

	if val, ok := from.GetThemeParams(); ok {
		r.ThemeParams = val
	}

	r.Platform = from.GetPlatform()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesRequestMainWebViewRequest) TypeID() uint32 {
	return MessagesRequestMainWebViewRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesRequestMainWebViewRequest) TypeName() string {
	return "messages.requestMainWebView"
}

// TypeInfo returns info about TL type.
func (r *MessagesRequestMainWebViewRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.requestMainWebView",
		ID:   MessagesRequestMainWebViewRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Compact",
			SchemaName: "compact",
			Null:       !r.Flags.Has(7),
		},
		{
			Name:       "Fullscreen",
			SchemaName: "fullscreen",
			Null:       !r.Flags.Has(8),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Bot",
			SchemaName: "bot",
		},
		{
			Name:       "StartParam",
			SchemaName: "start_param",
			Null:       !r.Flags.Has(1),
		},
		{
			Name:       "ThemeParams",
			SchemaName: "theme_params",
			Null:       !r.Flags.Has(0),
		},
		{
			Name:       "Platform",
			SchemaName: "platform",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (r *MessagesRequestMainWebViewRequest) SetFlags() {
	if !(r.Compact == false) {
		r.Flags.Set(7)
	}
	if !(r.Fullscreen == false) {
		r.Flags.Set(8)
	}
	if !(r.StartParam == "") {
		r.Flags.Set(1)
	}
	if !(r.ThemeParams.Zero()) {
		r.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (r *MessagesRequestMainWebViewRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.requestMainWebView#c9e01e7b as nil")
	}
	b.PutID(MessagesRequestMainWebViewRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *MessagesRequestMainWebViewRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.requestMainWebView#c9e01e7b as nil")
	}
	r.SetFlags()
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.requestMainWebView#c9e01e7b: field flags: %w", err)
	}
	if r.Peer == nil {
		return fmt.Errorf("unable to encode messages.requestMainWebView#c9e01e7b: field peer is nil")
	}
	if err := r.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.requestMainWebView#c9e01e7b: field peer: %w", err)
	}
	if r.Bot == nil {
		return fmt.Errorf("unable to encode messages.requestMainWebView#c9e01e7b: field bot is nil")
	}
	if err := r.Bot.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.requestMainWebView#c9e01e7b: field bot: %w", err)
	}
	if r.Flags.Has(1) {
		b.PutString(r.StartParam)
	}
	if r.Flags.Has(0) {
		if err := r.ThemeParams.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.requestMainWebView#c9e01e7b: field theme_params: %w", err)
		}
	}
	b.PutString(r.Platform)
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesRequestMainWebViewRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.requestMainWebView#c9e01e7b to nil")
	}
	if err := b.ConsumeID(MessagesRequestMainWebViewRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.requestMainWebView#c9e01e7b: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *MessagesRequestMainWebViewRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.requestMainWebView#c9e01e7b to nil")
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.requestMainWebView#c9e01e7b: field flags: %w", err)
		}
	}
	r.Compact = r.Flags.Has(7)
	r.Fullscreen = r.Flags.Has(8)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestMainWebView#c9e01e7b: field peer: %w", err)
		}
		r.Peer = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestMainWebView#c9e01e7b: field bot: %w", err)
		}
		r.Bot = value
	}
	if r.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestMainWebView#c9e01e7b: field start_param: %w", err)
		}
		r.StartParam = value
	}
	if r.Flags.Has(0) {
		if err := r.ThemeParams.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.requestMainWebView#c9e01e7b: field theme_params: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestMainWebView#c9e01e7b: field platform: %w", err)
		}
		r.Platform = value
	}
	return nil
}

// SetCompact sets value of Compact conditional field.
func (r *MessagesRequestMainWebViewRequest) SetCompact(value bool) {
	if value {
		r.Flags.Set(7)
		r.Compact = true
	} else {
		r.Flags.Unset(7)
		r.Compact = false
	}
}

// GetCompact returns value of Compact conditional field.
func (r *MessagesRequestMainWebViewRequest) GetCompact() (value bool) {
	if r == nil {
		return
	}
	return r.Flags.Has(7)
}

// SetFullscreen sets value of Fullscreen conditional field.
func (r *MessagesRequestMainWebViewRequest) SetFullscreen(value bool) {
	if value {
		r.Flags.Set(8)
		r.Fullscreen = true
	} else {
		r.Flags.Unset(8)
		r.Fullscreen = false
	}
}

// GetFullscreen returns value of Fullscreen conditional field.
func (r *MessagesRequestMainWebViewRequest) GetFullscreen() (value bool) {
	if r == nil {
		return
	}
	return r.Flags.Has(8)
}

// GetPeer returns value of Peer field.
func (r *MessagesRequestMainWebViewRequest) GetPeer() (value InputPeerClass) {
	if r == nil {
		return
	}
	return r.Peer
}

// GetBot returns value of Bot field.
func (r *MessagesRequestMainWebViewRequest) GetBot() (value InputUserClass) {
	if r == nil {
		return
	}
	return r.Bot
}

// SetStartParam sets value of StartParam conditional field.
func (r *MessagesRequestMainWebViewRequest) SetStartParam(value string) {
	r.Flags.Set(1)
	r.StartParam = value
}

// GetStartParam returns value of StartParam conditional field and
// boolean which is true if field was set.
func (r *MessagesRequestMainWebViewRequest) GetStartParam() (value string, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(1) {
		return value, false
	}
	return r.StartParam, true
}

// SetThemeParams sets value of ThemeParams conditional field.
func (r *MessagesRequestMainWebViewRequest) SetThemeParams(value DataJSON) {
	r.Flags.Set(0)
	r.ThemeParams = value
}

// GetThemeParams returns value of ThemeParams conditional field and
// boolean which is true if field was set.
func (r *MessagesRequestMainWebViewRequest) GetThemeParams() (value DataJSON, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(0) {
		return value, false
	}
	return r.ThemeParams, true
}

// GetPlatform returns value of Platform field.
func (r *MessagesRequestMainWebViewRequest) GetPlatform() (value string) {
	if r == nil {
		return
	}
	return r.Platform
}

// MessagesRequestMainWebView invokes method messages.requestMainWebView#c9e01e7b returning error if any.
// Open a Main Mini App¹.
//
// Links:
//  1. https://core.telegram.org/api/bots/webapps#main-mini-apps
//
// Possible errors:
//
//	400 BOT_INVALID: This is not a valid bot.
//
// See https://core.telegram.org/method/messages.requestMainWebView for reference.
func (c *Client) MessagesRequestMainWebView(ctx context.Context, request *MessagesRequestMainWebViewRequest) (*WebViewResultURL, error) {
	var result WebViewResultURL

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
