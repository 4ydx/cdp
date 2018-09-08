// Code generated by cdpgen. DO NOT EDIT.

package io

import (
	"encoding/json"
	"log"

	"github.com/4ydx/cdp/protocol"
)

const (
	CommandIOClose       = "IO.close"
	CommandIORead        = "IO.read"
	CommandIOResolveBlob = "IO.resolveBlob"
)

// CloseArgs represents the arguments for Close in the IO domain.
type CloseArgs struct {
	Handle StreamHandle `json:"handle"` // Handle of the stream to close.
}

// Unmarshal the byte array into a return value for Close in the IO domain.
func (a *CloseArgs) UnmarshalJSON(b []byte) error {
	type Copy CloseArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = CloseArgs(*c)
	return nil
}

// Marshall the byte array into a return value for Close in the IO domain.
func (a *CloseArgs) MarshalJSON() ([]byte, error) {
	type Copy CloseArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// CloseReply represents the return values for Close in the IO domain.
type CloseReply struct {
}

// CloseReply returns whether or not the FrameID matches the reply value for Close in the IO domain.
func (a *CloseReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: CloseReply %s", err)
		return false, err
	}
	return true, nil
}

// CloseReply returns the FrameID value for Close in the IO domain.
func (a *CloseReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for Close in the IO domain.
func (a *CloseReply) UnmarshalJSON(b []byte) error {
	type Copy CloseReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = CloseReply(*c)
	return nil
}

// ReadArgs represents the arguments for Read in the IO domain.
type ReadArgs struct {
	Handle StreamHandle `json:"handle"`           // Handle of the stream to read.
	Offset int          `json:"offset,omitempty"` // Seek to the specified offset before reading (if not specificed, proceed with offset following the last read). Some types of streams may only support sequential reads.
	Size   int          `json:"size,omitempty"`   // Maximum number of bytes to read (left upon the agent discretion if not specified).
}

// Unmarshal the byte array into a return value for Read in the IO domain.
func (a *ReadArgs) UnmarshalJSON(b []byte) error {
	type Copy ReadArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ReadArgs(*c)
	return nil
}

// Marshall the byte array into a return value for Read in the IO domain.
func (a *ReadArgs) MarshalJSON() ([]byte, error) {
	type Copy ReadArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// ReadReply represents the return values for Read in the IO domain.
type ReadReply struct {
	Base64Encoded bool   `json:"base64Encoded,omitempty"` // Set if the data is base64-encoded
	Data          string `json:"data"`                    // Data that were read.
	EOF           bool   `json:"eof"`                     // Set if the end-of-file condition occurred while reading.
}

// ReadReply returns whether or not the FrameID matches the reply value for Read in the IO domain.
func (a *ReadReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: ReadReply %s", err)
		return false, err
	}
	return true, nil
}

// ReadReply returns the FrameID value for Read in the IO domain.
func (a *ReadReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for Read in the IO domain.
func (a *ReadReply) UnmarshalJSON(b []byte) error {
	type Copy ReadReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ReadReply(*c)
	return nil
}

// ResolveBlobArgs represents the arguments for ResolveBlob in the IO domain.
type ResolveBlobArgs struct {
	ObjectID shared.RemoteObjectID `json:"objectId"` // Object id of a Blob object wrapper.
}

// Unmarshal the byte array into a return value for ResolveBlob in the IO domain.
func (a *ResolveBlobArgs) UnmarshalJSON(b []byte) error {
	type Copy ResolveBlobArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ResolveBlobArgs(*c)
	return nil
}

// Marshall the byte array into a return value for ResolveBlob in the IO domain.
func (a *ResolveBlobArgs) MarshalJSON() ([]byte, error) {
	type Copy ResolveBlobArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// ResolveBlobReply represents the return values for ResolveBlob in the IO domain.
type ResolveBlobReply struct {
	UUID string `json:"uuid"` // UUID of the specified Blob.
}

// ResolveBlobReply returns whether or not the FrameID matches the reply value for ResolveBlob in the IO domain.
func (a *ResolveBlobReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: ResolveBlobReply %s", err)
		return false, err
	}
	return true, nil
}

// ResolveBlobReply returns the FrameID value for ResolveBlob in the IO domain.
func (a *ResolveBlobReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for ResolveBlob in the IO domain.
func (a *ResolveBlobReply) UnmarshalJSON(b []byte) error {
	type Copy ResolveBlobReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ResolveBlobReply(*c)
	return nil
}
