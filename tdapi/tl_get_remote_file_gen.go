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

// GetRemoteFileRequest represents TL type `getRemoteFile#7f632732`.
type GetRemoteFileRequest struct {
	// Remote identifier of the file to get
	RemoteFileID string
	// File type; pass null if unknown
	FileType FileTypeClass
}

// GetRemoteFileRequestTypeID is TL type id of GetRemoteFileRequest.
const GetRemoteFileRequestTypeID = 0x7f632732

// Ensuring interfaces in compile-time for GetRemoteFileRequest.
var (
	_ bin.Encoder     = &GetRemoteFileRequest{}
	_ bin.Decoder     = &GetRemoteFileRequest{}
	_ bin.BareEncoder = &GetRemoteFileRequest{}
	_ bin.BareDecoder = &GetRemoteFileRequest{}
)

func (g *GetRemoteFileRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.RemoteFileID == "") {
		return false
	}
	if !(g.FileType == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetRemoteFileRequest) String() string {
	if g == nil {
		return "GetRemoteFileRequest(nil)"
	}
	type Alias GetRemoteFileRequest
	return fmt.Sprintf("GetRemoteFileRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetRemoteFileRequest) TypeID() uint32 {
	return GetRemoteFileRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetRemoteFileRequest) TypeName() string {
	return "getRemoteFile"
}

// TypeInfo returns info about TL type.
func (g *GetRemoteFileRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getRemoteFile",
		ID:   GetRemoteFileRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "RemoteFileID",
			SchemaName: "remote_file_id",
		},
		{
			Name:       "FileType",
			SchemaName: "file_type",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetRemoteFileRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getRemoteFile#7f632732 as nil")
	}
	b.PutID(GetRemoteFileRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetRemoteFileRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getRemoteFile#7f632732 as nil")
	}
	b.PutString(g.RemoteFileID)
	if g.FileType == nil {
		return fmt.Errorf("unable to encode getRemoteFile#7f632732: field file_type is nil")
	}
	if err := g.FileType.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getRemoteFile#7f632732: field file_type: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetRemoteFileRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getRemoteFile#7f632732 to nil")
	}
	if err := b.ConsumeID(GetRemoteFileRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getRemoteFile#7f632732: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetRemoteFileRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getRemoteFile#7f632732 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getRemoteFile#7f632732: field remote_file_id: %w", err)
		}
		g.RemoteFileID = value
	}
	{
		value, err := DecodeFileType(b)
		if err != nil {
			return fmt.Errorf("unable to decode getRemoteFile#7f632732: field file_type: %w", err)
		}
		g.FileType = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetRemoteFileRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getRemoteFile#7f632732 as nil")
	}
	b.ObjStart()
	b.PutID("getRemoteFile")
	b.Comma()
	b.FieldStart("remote_file_id")
	b.PutString(g.RemoteFileID)
	b.Comma()
	b.FieldStart("file_type")
	if g.FileType == nil {
		return fmt.Errorf("unable to encode getRemoteFile#7f632732: field file_type is nil")
	}
	if err := g.FileType.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getRemoteFile#7f632732: field file_type: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetRemoteFileRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getRemoteFile#7f632732 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getRemoteFile"); err != nil {
				return fmt.Errorf("unable to decode getRemoteFile#7f632732: %w", err)
			}
		case "remote_file_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getRemoteFile#7f632732: field remote_file_id: %w", err)
			}
			g.RemoteFileID = value
		case "file_type":
			value, err := DecodeTDLibJSONFileType(b)
			if err != nil {
				return fmt.Errorf("unable to decode getRemoteFile#7f632732: field file_type: %w", err)
			}
			g.FileType = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetRemoteFileID returns value of RemoteFileID field.
func (g *GetRemoteFileRequest) GetRemoteFileID() (value string) {
	if g == nil {
		return
	}
	return g.RemoteFileID
}

// GetFileType returns value of FileType field.
func (g *GetRemoteFileRequest) GetFileType() (value FileTypeClass) {
	if g == nil {
		return
	}
	return g.FileType
}

// GetRemoteFile invokes method getRemoteFile#7f632732 returning error if any.
func (c *Client) GetRemoteFile(ctx context.Context, request *GetRemoteFileRequest) (*File, error) {
	var result File

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
