// Code generated by cdpgen. DO NOT EDIT.

package css

import (
	"encoding/json"
	"log"
)

const (
	EventCSSFontsUpdated            = "CSS.fontsUpdated"
	EventCSSMediaQueryResultChanged = "CSS.mediaQueryResultChanged"
	EventCSSStyleSheetAdded         = "CSS.styleSheetAdded"
	EventCSSStyleSheetChanged       = "CSS.styleSheetChanged"
	EventCSSStyleSheetRemoved       = "CSS.styleSheetRemoved"
)

type Unmarshaler func() json.Unmarshaler

var EventConstants = map[string]Unmarshaler{
	EventCSSFontsUpdated:            func() json.Unmarshaler { return &FontsUpdatedReply{} },
	EventCSSMediaQueryResultChanged: func() json.Unmarshaler { return &MediaQueryResultChangedReply{} },
	EventCSSStyleSheetAdded:         func() json.Unmarshaler { return &StyleSheetAddedReply{} },
	EventCSSStyleSheetChanged:       func() json.Unmarshaler { return &StyleSheetChangedReply{} },
	EventCSSStyleSheetRemoved:       func() json.Unmarshaler { return &StyleSheetRemovedReply{} },
}

func GetEventReply(event string) (json.Unmarshaler, bool) {
	e, ok := EventConstants[event]
	if ok {
		return e(), ok
	}
	return nil, ok
}

// FontsUpdatedReply is the reply for FontsUpdated events.
type FontsUpdatedReply struct {
	Font *FontFace `json:"font,omitempty"` // The web font that has loaded.
}

// Unmarshal the byte array into a return value for FontsUpdated in the CSS domain.
func (a *FontsUpdatedReply) UnmarshalJSON(b []byte) error {
	type Copy FontsUpdatedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = FontsUpdatedReply(*c)
	return nil
}

// FontsUpdatedReply returns whether or not the FrameID matches the reply value for FontsUpdated in the CSS domain.
func (a *FontsUpdatedReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: FontsUpdatedReply %s", err)
	}
	return true
}

// FontsUpdatedReply returns the FrameID for FontsUpdated in the CSS domain.
func (a *FontsUpdatedReply) GetFrameID() string {
	return ""
}

// MediaQueryResultChangedReply is the reply for MediaQueryResultChanged events.
type MediaQueryResultChangedReply struct {
}

// Unmarshal the byte array into a return value for MediaQueryResultChanged in the CSS domain.
func (a *MediaQueryResultChangedReply) UnmarshalJSON(b []byte) error {
	type Copy MediaQueryResultChangedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = MediaQueryResultChangedReply(*c)
	return nil
}

// MediaQueryResultChangedReply returns whether or not the FrameID matches the reply value for MediaQueryResultChanged in the CSS domain.
func (a *MediaQueryResultChangedReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: MediaQueryResultChangedReply %s", err)
	}
	return true
}

// MediaQueryResultChangedReply returns the FrameID for MediaQueryResultChanged in the CSS domain.
func (a *MediaQueryResultChangedReply) GetFrameID() string {
	return ""
}

// StyleSheetAddedReply is the reply for StyleSheetAdded events.
type StyleSheetAddedReply struct {
	Header StyleSheetHeader `json:"header"` // Added stylesheet metainfo.
}

// Unmarshal the byte array into a return value for StyleSheetAdded in the CSS domain.
func (a *StyleSheetAddedReply) UnmarshalJSON(b []byte) error {
	type Copy StyleSheetAddedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = StyleSheetAddedReply(*c)
	return nil
}

// StyleSheetAddedReply returns whether or not the FrameID matches the reply value for StyleSheetAdded in the CSS domain.
func (a *StyleSheetAddedReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: StyleSheetAddedReply %s", err)
	}
	return true
}

// StyleSheetAddedReply returns the FrameID for StyleSheetAdded in the CSS domain.
func (a *StyleSheetAddedReply) GetFrameID() string {
	return ""
}

// StyleSheetChangedReply is the reply for StyleSheetChanged events.
type StyleSheetChangedReply struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // No description.
}

// Unmarshal the byte array into a return value for StyleSheetChanged in the CSS domain.
func (a *StyleSheetChangedReply) UnmarshalJSON(b []byte) error {
	type Copy StyleSheetChangedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = StyleSheetChangedReply(*c)
	return nil
}

// StyleSheetChangedReply returns whether or not the FrameID matches the reply value for StyleSheetChanged in the CSS domain.
func (a *StyleSheetChangedReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: StyleSheetChangedReply %s", err)
	}
	return true
}

// StyleSheetChangedReply returns the FrameID for StyleSheetChanged in the CSS domain.
func (a *StyleSheetChangedReply) GetFrameID() string {
	return ""
}

// StyleSheetRemovedReply is the reply for StyleSheetRemoved events.
type StyleSheetRemovedReply struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // Identifier of the removed stylesheet.
}

// Unmarshal the byte array into a return value for StyleSheetRemoved in the CSS domain.
func (a *StyleSheetRemovedReply) UnmarshalJSON(b []byte) error {
	type Copy StyleSheetRemovedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = StyleSheetRemovedReply(*c)
	return nil
}

// StyleSheetRemovedReply returns whether or not the FrameID matches the reply value for StyleSheetRemoved in the CSS domain.
func (a *StyleSheetRemovedReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: StyleSheetRemovedReply %s", err)
	}
	return true
}

// StyleSheetRemovedReply returns the FrameID for StyleSheetRemoved in the CSS domain.
func (a *StyleSheetRemovedReply) GetFrameID() string {
	return ""
}
