// Code generated by cdpgen. DO NOT EDIT.

package applicationcache

import (
	"encoding/json"
	"log"

	shared "github.com/4ydx/cdp/protocol"
)

const (
	CommandApplicationCacheEnable                      = "ApplicationCache.enable"
	CommandApplicationCacheGetApplicationCacheForFrame = "ApplicationCache.getApplicationCacheForFrame"
	CommandApplicationCacheGetFramesWithManifests      = "ApplicationCache.getFramesWithManifests"
	CommandApplicationCacheGetManifestForFrame         = "ApplicationCache.getManifestForFrame"
)

// EnableArgs represents the arguments for Enable in the ApplicationCache domain.
type EnableArgs struct {
}

// Unmarshal the byte array into a return value for Enable in the ApplicationCache domain.
func (a *EnableArgs) UnmarshalJSON(b []byte) error {
	type Copy EnableArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = EnableArgs(*c)
	return nil
}

// Marshall the byte array into a return value for Enable in the ApplicationCache domain.
func (a *EnableArgs) MarshalJSON() ([]byte, error) {
	type Copy EnableArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// EnableReply represents the return values for Enable in the ApplicationCache domain.
type EnableReply struct {
}

// EnableReply returns whether or not the FrameID matches the reply value for Enable in the ApplicationCache domain.
func (a *EnableReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: EnableReply %s", err)
		return false, err
	}
	return true, nil
}

// EnableReply returns the FrameID value for Enable in the ApplicationCache domain.
func (a *EnableReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for Enable in the ApplicationCache domain.
func (a *EnableReply) UnmarshalJSON(b []byte) error {
	type Copy EnableReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = EnableReply(*c)
	return nil
}

// GetApplicationCacheForFrameArgs represents the arguments for GetApplicationCacheForFrame in the ApplicationCache domain.
type GetApplicationCacheForFrameArgs struct {
	FrameID shared.FrameID `json:"frameId"` // Identifier of the frame containing document whose application cache is retrieved.
}

// Unmarshal the byte array into a return value for GetApplicationCacheForFrame in the ApplicationCache domain.
func (a *GetApplicationCacheForFrameArgs) UnmarshalJSON(b []byte) error {
	type Copy GetApplicationCacheForFrameArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetApplicationCacheForFrameArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetApplicationCacheForFrame in the ApplicationCache domain.
func (a *GetApplicationCacheForFrameArgs) MarshalJSON() ([]byte, error) {
	type Copy GetApplicationCacheForFrameArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetApplicationCacheForFrameReply represents the return values for GetApplicationCacheForFrame in the ApplicationCache domain.
type GetApplicationCacheForFrameReply struct {
	ApplicationCache ApplicationCache `json:"applicationCache"` // Relevant application cache data for the document in given frame.
}

// GetApplicationCacheForFrameReply returns whether or not the FrameID matches the reply value for GetApplicationCacheForFrame in the ApplicationCache domain.
func (a *GetApplicationCacheForFrameReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: GetApplicationCacheForFrameReply %s", err)
		return false, err
	}
	return true, nil
}

// GetApplicationCacheForFrameReply returns the FrameID value for GetApplicationCacheForFrame in the ApplicationCache domain.
func (a *GetApplicationCacheForFrameReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for GetApplicationCacheForFrame in the ApplicationCache domain.
func (a *GetApplicationCacheForFrameReply) UnmarshalJSON(b []byte) error {
	type Copy GetApplicationCacheForFrameReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetApplicationCacheForFrameReply(*c)
	return nil
}

// GetFramesWithManifestsArgs represents the arguments for GetFramesWithManifests in the ApplicationCache domain.
type GetFramesWithManifestsArgs struct {
}

// Unmarshal the byte array into a return value for GetFramesWithManifests in the ApplicationCache domain.
func (a *GetFramesWithManifestsArgs) UnmarshalJSON(b []byte) error {
	type Copy GetFramesWithManifestsArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetFramesWithManifestsArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetFramesWithManifests in the ApplicationCache domain.
func (a *GetFramesWithManifestsArgs) MarshalJSON() ([]byte, error) {
	type Copy GetFramesWithManifestsArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetFramesWithManifestsReply represents the return values for GetFramesWithManifests in the ApplicationCache domain.
type GetFramesWithManifestsReply struct {
	FrameIDs []FrameWithManifest `json:"frameIds"` // Array of frame identifiers with manifest urls for each frame containing a document associated with some application cache.
}

// GetFramesWithManifestsReply returns whether or not the FrameID matches the reply value for GetFramesWithManifests in the ApplicationCache domain.
func (a *GetFramesWithManifestsReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: GetFramesWithManifestsReply %s", err)
		return false, err
	}
	return true, nil
}

// GetFramesWithManifestsReply returns the FrameID value for GetFramesWithManifests in the ApplicationCache domain.
func (a *GetFramesWithManifestsReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for GetFramesWithManifests in the ApplicationCache domain.
func (a *GetFramesWithManifestsReply) UnmarshalJSON(b []byte) error {
	type Copy GetFramesWithManifestsReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetFramesWithManifestsReply(*c)
	return nil
}

// GetManifestForFrameArgs represents the arguments for GetManifestForFrame in the ApplicationCache domain.
type GetManifestForFrameArgs struct {
	FrameID shared.FrameID `json:"frameId"` // Identifier of the frame containing document whose manifest is retrieved.
}

// Unmarshal the byte array into a return value for GetManifestForFrame in the ApplicationCache domain.
func (a *GetManifestForFrameArgs) UnmarshalJSON(b []byte) error {
	type Copy GetManifestForFrameArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetManifestForFrameArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetManifestForFrame in the ApplicationCache domain.
func (a *GetManifestForFrameArgs) MarshalJSON() ([]byte, error) {
	type Copy GetManifestForFrameArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetManifestForFrameReply represents the return values for GetManifestForFrame in the ApplicationCache domain.
type GetManifestForFrameReply struct {
	ManifestURL string `json:"manifestURL"` // Manifest URL for document in the given frame.
}

// GetManifestForFrameReply returns whether or not the FrameID matches the reply value for GetManifestForFrame in the ApplicationCache domain.
func (a *GetManifestForFrameReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: GetManifestForFrameReply %s", err)
		return false, err
	}
	return true, nil
}

// GetManifestForFrameReply returns the FrameID value for GetManifestForFrame in the ApplicationCache domain.
func (a *GetManifestForFrameReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for GetManifestForFrame in the ApplicationCache domain.
func (a *GetManifestForFrameReply) UnmarshalJSON(b []byte) error {
	type Copy GetManifestForFrameReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetManifestForFrameReply(*c)
	return nil
}
