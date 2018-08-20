// Code generated by cdpgen. DO NOT EDIT.

package css

import (
	"encoding/json"
	"log"

	"github.com/4ydx/cdp/protocol"
	"github.com/4ydx/cdp/protocol/dom"
)

const (
	CommandCSSAddRule                          = "CSS.addRule"
	CommandCSSCollectClassNames                = "CSS.collectClassNames"
	CommandCSSCreateStyleSheet                 = "CSS.createStyleSheet"
	CommandCSSDisable                          = "CSS.disable"
	CommandCSSEnable                           = "CSS.enable"
	CommandCSSForcePseudoState                 = "CSS.forcePseudoState"
	CommandCSSGetBackgroundColors              = "CSS.getBackgroundColors"
	CommandCSSGetComputedStyleForNode          = "CSS.getComputedStyleForNode"
	CommandCSSGetInlineStylesForNode           = "CSS.getInlineStylesForNode"
	CommandCSSGetMatchedStylesForNode          = "CSS.getMatchedStylesForNode"
	CommandCSSGetMediaQueries                  = "CSS.getMediaQueries"
	CommandCSSGetPlatformFontsForNode          = "CSS.getPlatformFontsForNode"
	CommandCSSGetStyleSheetText                = "CSS.getStyleSheetText"
	CommandCSSSetEffectivePropertyValueForNode = "CSS.setEffectivePropertyValueForNode"
	CommandCSSSetKeyframeKey                   = "CSS.setKeyframeKey"
	CommandCSSSetMediaText                     = "CSS.setMediaText"
	CommandCSSSetRuleSelector                  = "CSS.setRuleSelector"
	CommandCSSSetStyleSheetText                = "CSS.setStyleSheetText"
	CommandCSSSetStyleTexts                    = "CSS.setStyleTexts"
	CommandCSSStartRuleUsageTracking           = "CSS.startRuleUsageTracking"
	CommandCSSStopRuleUsageTracking            = "CSS.stopRuleUsageTracking"
	CommandCSSTakeCoverageDelta                = "CSS.takeCoverageDelta"
)

// AddRuleArgs represents the arguments for AddRule in the CSS domain.
type AddRuleArgs struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // The css style sheet identifier where a new rule should be inserted.
	RuleText     string       `json:"ruleText"`     // The text of a new rule.
	Location     SourceRange  `json:"location"`     // Text position of a new rule in the target style sheet.
}

// Unmarshal the byte array into a return value for AddRule in the CSS domain.
func (a *AddRuleArgs) UnmarshalJSON(b []byte) error {
	type Copy AddRuleArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = AddRuleArgs(*c)
	return nil
}

// Marshall the byte array into a return value for AddRule in the CSS domain.
func (a *AddRuleArgs) MarshalJSON() ([]byte, error) {
	type Copy AddRuleArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// AddRuleReply represents the return values for AddRule in the CSS domain.
type AddRuleReply struct {
	Rule Rule `json:"rule"` // The newly created rule.
}

// AddRuleReply returns whether or not the FrameID matches the reply value for AddRule in the CSS domain.
func (a *AddRuleReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: AddRuleReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for AddRule in the CSS domain.
func (a *AddRuleReply) UnmarshalJSON(b []byte) error {
	type Copy AddRuleReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = AddRuleReply(*c)
	return nil
}

// CollectClassNamesArgs represents the arguments for CollectClassNames in the CSS domain.
type CollectClassNamesArgs struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // No description.
}

// Unmarshal the byte array into a return value for CollectClassNames in the CSS domain.
func (a *CollectClassNamesArgs) UnmarshalJSON(b []byte) error {
	type Copy CollectClassNamesArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = CollectClassNamesArgs(*c)
	return nil
}

// Marshall the byte array into a return value for CollectClassNames in the CSS domain.
func (a *CollectClassNamesArgs) MarshalJSON() ([]byte, error) {
	type Copy CollectClassNamesArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// CollectClassNamesReply represents the return values for CollectClassNames in the CSS domain.
type CollectClassNamesReply struct {
	ClassNames []string `json:"classNames"` // Class name list.
}

// CollectClassNamesReply returns whether or not the FrameID matches the reply value for CollectClassNames in the CSS domain.
func (a *CollectClassNamesReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: CollectClassNamesReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for CollectClassNames in the CSS domain.
func (a *CollectClassNamesReply) UnmarshalJSON(b []byte) error {
	type Copy CollectClassNamesReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = CollectClassNamesReply(*c)
	return nil
}

// CreateStyleSheetArgs represents the arguments for CreateStyleSheet in the CSS domain.
type CreateStyleSheetArgs struct {
	FrameID shared.FrameID `json:"frameId"` // Identifier of the frame where "via-inspector" stylesheet should be created.
}

// Unmarshal the byte array into a return value for CreateStyleSheet in the CSS domain.
func (a *CreateStyleSheetArgs) UnmarshalJSON(b []byte) error {
	type Copy CreateStyleSheetArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = CreateStyleSheetArgs(*c)
	return nil
}

// Marshall the byte array into a return value for CreateStyleSheet in the CSS domain.
func (a *CreateStyleSheetArgs) MarshalJSON() ([]byte, error) {
	type Copy CreateStyleSheetArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// CreateStyleSheetReply represents the return values for CreateStyleSheet in the CSS domain.
type CreateStyleSheetReply struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // Identifier of the created "via-inspector" stylesheet.
}

// CreateStyleSheetReply returns whether or not the FrameID matches the reply value for CreateStyleSheet in the CSS domain.
func (a *CreateStyleSheetReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: CreateStyleSheetReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for CreateStyleSheet in the CSS domain.
func (a *CreateStyleSheetReply) UnmarshalJSON(b []byte) error {
	type Copy CreateStyleSheetReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = CreateStyleSheetReply(*c)
	return nil
}

// DisableArgs represents the arguments for Disable in the CSS domain.
type DisableArgs struct {
}

// Unmarshal the byte array into a return value for Disable in the CSS domain.
func (a *DisableArgs) UnmarshalJSON(b []byte) error {
	type Copy DisableArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DisableArgs(*c)
	return nil
}

// Marshall the byte array into a return value for Disable in the CSS domain.
func (a *DisableArgs) MarshalJSON() ([]byte, error) {
	type Copy DisableArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// DisableReply represents the return values for Disable in the CSS domain.
type DisableReply struct {
}

// DisableReply returns whether or not the FrameID matches the reply value for Disable in the CSS domain.
func (a *DisableReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: DisableReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for Disable in the CSS domain.
func (a *DisableReply) UnmarshalJSON(b []byte) error {
	type Copy DisableReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = DisableReply(*c)
	return nil
}

// EnableArgs represents the arguments for Enable in the CSS domain.
type EnableArgs struct {
}

// Unmarshal the byte array into a return value for Enable in the CSS domain.
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

// Marshall the byte array into a return value for Enable in the CSS domain.
func (a *EnableArgs) MarshalJSON() ([]byte, error) {
	type Copy EnableArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// EnableReply represents the return values for Enable in the CSS domain.
type EnableReply struct {
}

// EnableReply returns whether or not the FrameID matches the reply value for Enable in the CSS domain.
func (a *EnableReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: EnableReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for Enable in the CSS domain.
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

// ForcePseudoStateArgs represents the arguments for ForcePseudoState in the CSS domain.
type ForcePseudoStateArgs struct {
	NodeID              dom.NodeID `json:"nodeId"`              // The element id for which to force the pseudo state.
	ForcedPseudoClasses []string   `json:"forcedPseudoClasses"` // Element pseudo classes to force when computing the element's style.
}

// Unmarshal the byte array into a return value for ForcePseudoState in the CSS domain.
func (a *ForcePseudoStateArgs) UnmarshalJSON(b []byte) error {
	type Copy ForcePseudoStateArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ForcePseudoStateArgs(*c)
	return nil
}

// Marshall the byte array into a return value for ForcePseudoState in the CSS domain.
func (a *ForcePseudoStateArgs) MarshalJSON() ([]byte, error) {
	type Copy ForcePseudoStateArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// ForcePseudoStateReply represents the return values for ForcePseudoState in the CSS domain.
type ForcePseudoStateReply struct {
}

// ForcePseudoStateReply returns whether or not the FrameID matches the reply value for ForcePseudoState in the CSS domain.
func (a *ForcePseudoStateReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: ForcePseudoStateReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for ForcePseudoState in the CSS domain.
func (a *ForcePseudoStateReply) UnmarshalJSON(b []byte) error {
	type Copy ForcePseudoStateReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ForcePseudoStateReply(*c)
	return nil
}

// GetBackgroundColorsArgs represents the arguments for GetBackgroundColors in the CSS domain.
type GetBackgroundColorsArgs struct {
	NodeID dom.NodeID `json:"nodeId"` // Id of the node to get background colors for.
}

// Unmarshal the byte array into a return value for GetBackgroundColors in the CSS domain.
func (a *GetBackgroundColorsArgs) UnmarshalJSON(b []byte) error {
	type Copy GetBackgroundColorsArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetBackgroundColorsArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetBackgroundColors in the CSS domain.
func (a *GetBackgroundColorsArgs) MarshalJSON() ([]byte, error) {
	type Copy GetBackgroundColorsArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetBackgroundColorsReply represents the return values for GetBackgroundColors in the CSS domain.
type GetBackgroundColorsReply struct {
	BackgroundColors     []string `json:"backgroundColors,omitempty"`     // The range of background colors behind this element, if it contains any visible text. If no visible text is present, this will be undefined. In the case of a flat background color, this will consist of simply that color. In the case of a gradient, this will consist of each of the color stops. For anything more complicated, this will be an empty array. Images will be ignored (as if the image had failed to load).
	ComputedFontSize     string   `json:"computedFontSize,omitempty"`     // The computed font size for this node, as a CSS computed value string (e.g. '12px').
	ComputedFontWeight   string   `json:"computedFontWeight,omitempty"`   // The computed font weight for this node, as a CSS computed value string (e.g. 'normal' or '100').
	ComputedBodyFontSize string   `json:"computedBodyFontSize,omitempty"` // The computed font size for the document body, as a computed CSS value string (e.g. '16px').
}

// GetBackgroundColorsReply returns whether or not the FrameID matches the reply value for GetBackgroundColors in the CSS domain.
func (a *GetBackgroundColorsReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: GetBackgroundColorsReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for GetBackgroundColors in the CSS domain.
func (a *GetBackgroundColorsReply) UnmarshalJSON(b []byte) error {
	type Copy GetBackgroundColorsReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetBackgroundColorsReply(*c)
	return nil
}

// GetComputedStyleForNodeArgs represents the arguments for GetComputedStyleForNode in the CSS domain.
type GetComputedStyleForNodeArgs struct {
	NodeID dom.NodeID `json:"nodeId"` // No description.
}

// Unmarshal the byte array into a return value for GetComputedStyleForNode in the CSS domain.
func (a *GetComputedStyleForNodeArgs) UnmarshalJSON(b []byte) error {
	type Copy GetComputedStyleForNodeArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetComputedStyleForNodeArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetComputedStyleForNode in the CSS domain.
func (a *GetComputedStyleForNodeArgs) MarshalJSON() ([]byte, error) {
	type Copy GetComputedStyleForNodeArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetComputedStyleForNodeReply represents the return values for GetComputedStyleForNode in the CSS domain.
type GetComputedStyleForNodeReply struct {
	ComputedStyle []ComputedStyleProperty `json:"computedStyle"` // Computed style for the specified DOM node.
}

// GetComputedStyleForNodeReply returns whether or not the FrameID matches the reply value for GetComputedStyleForNode in the CSS domain.
func (a *GetComputedStyleForNodeReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: GetComputedStyleForNodeReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for GetComputedStyleForNode in the CSS domain.
func (a *GetComputedStyleForNodeReply) UnmarshalJSON(b []byte) error {
	type Copy GetComputedStyleForNodeReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetComputedStyleForNodeReply(*c)
	return nil
}

// GetInlineStylesForNodeArgs represents the arguments for GetInlineStylesForNode in the CSS domain.
type GetInlineStylesForNodeArgs struct {
	NodeID dom.NodeID `json:"nodeId"` // No description.
}

// Unmarshal the byte array into a return value for GetInlineStylesForNode in the CSS domain.
func (a *GetInlineStylesForNodeArgs) UnmarshalJSON(b []byte) error {
	type Copy GetInlineStylesForNodeArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetInlineStylesForNodeArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetInlineStylesForNode in the CSS domain.
func (a *GetInlineStylesForNodeArgs) MarshalJSON() ([]byte, error) {
	type Copy GetInlineStylesForNodeArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetInlineStylesForNodeReply represents the return values for GetInlineStylesForNode in the CSS domain.
type GetInlineStylesForNodeReply struct {
	InlineStyle     Style `json:"inlineStyle,omitempty"`     // Inline style for the specified DOM node.
	AttributesStyle Style `json:"attributesStyle,omitempty"` // Attribute-defined element style (e.g. resulting from "width=20 height=100%").
}

// GetInlineStylesForNodeReply returns whether or not the FrameID matches the reply value for GetInlineStylesForNode in the CSS domain.
func (a *GetInlineStylesForNodeReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: GetInlineStylesForNodeReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for GetInlineStylesForNode in the CSS domain.
func (a *GetInlineStylesForNodeReply) UnmarshalJSON(b []byte) error {
	type Copy GetInlineStylesForNodeReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetInlineStylesForNodeReply(*c)
	return nil
}

// GetMatchedStylesForNodeArgs represents the arguments for GetMatchedStylesForNode in the CSS domain.
type GetMatchedStylesForNodeArgs struct {
	NodeID dom.NodeID `json:"nodeId"` // No description.
}

// Unmarshal the byte array into a return value for GetMatchedStylesForNode in the CSS domain.
func (a *GetMatchedStylesForNodeArgs) UnmarshalJSON(b []byte) error {
	type Copy GetMatchedStylesForNodeArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetMatchedStylesForNodeArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetMatchedStylesForNode in the CSS domain.
func (a *GetMatchedStylesForNodeArgs) MarshalJSON() ([]byte, error) {
	type Copy GetMatchedStylesForNodeArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetMatchedStylesForNodeReply represents the return values for GetMatchedStylesForNode in the CSS domain.
type GetMatchedStylesForNodeReply struct {
	InlineStyle       Style                  `json:"inlineStyle,omitempty"`       // Inline style for the specified DOM node.
	AttributesStyle   Style                  `json:"attributesStyle,omitempty"`   // Attribute-defined element style (e.g. resulting from "width=20 height=100%").
	MatchedCSSRules   []RuleMatch            `json:"matchedCSSRules,omitempty"`   // CSS rules matching this node, from all applicable stylesheets.
	PseudoElements    []PseudoElementMatches `json:"pseudoElements,omitempty"`    // Pseudo style matches for this node.
	Inherited         []InheritedStyleEntry  `json:"inherited,omitempty"`         // A chain of inherited styles (from the immediate node parent up to the DOM tree root).
	CSSKeyframesRules []KeyframesRule        `json:"cssKeyframesRules,omitempty"` // A list of CSS keyframed animations matching this node.
}

// GetMatchedStylesForNodeReply returns whether or not the FrameID matches the reply value for GetMatchedStylesForNode in the CSS domain.
func (a *GetMatchedStylesForNodeReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: GetMatchedStylesForNodeReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for GetMatchedStylesForNode in the CSS domain.
func (a *GetMatchedStylesForNodeReply) UnmarshalJSON(b []byte) error {
	type Copy GetMatchedStylesForNodeReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetMatchedStylesForNodeReply(*c)
	return nil
}

// GetMediaQueriesArgs represents the arguments for GetMediaQueries in the CSS domain.
type GetMediaQueriesArgs struct {
}

// Unmarshal the byte array into a return value for GetMediaQueries in the CSS domain.
func (a *GetMediaQueriesArgs) UnmarshalJSON(b []byte) error {
	type Copy GetMediaQueriesArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetMediaQueriesArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetMediaQueries in the CSS domain.
func (a *GetMediaQueriesArgs) MarshalJSON() ([]byte, error) {
	type Copy GetMediaQueriesArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetMediaQueriesReply represents the return values for GetMediaQueries in the CSS domain.
type GetMediaQueriesReply struct {
	Medias []Media `json:"medias"` // No description.
}

// GetMediaQueriesReply returns whether or not the FrameID matches the reply value for GetMediaQueries in the CSS domain.
func (a *GetMediaQueriesReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: GetMediaQueriesReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for GetMediaQueries in the CSS domain.
func (a *GetMediaQueriesReply) UnmarshalJSON(b []byte) error {
	type Copy GetMediaQueriesReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetMediaQueriesReply(*c)
	return nil
}

// GetPlatformFontsForNodeArgs represents the arguments for GetPlatformFontsForNode in the CSS domain.
type GetPlatformFontsForNodeArgs struct {
	NodeID dom.NodeID `json:"nodeId"` // No description.
}

// Unmarshal the byte array into a return value for GetPlatformFontsForNode in the CSS domain.
func (a *GetPlatformFontsForNodeArgs) UnmarshalJSON(b []byte) error {
	type Copy GetPlatformFontsForNodeArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetPlatformFontsForNodeArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetPlatformFontsForNode in the CSS domain.
func (a *GetPlatformFontsForNodeArgs) MarshalJSON() ([]byte, error) {
	type Copy GetPlatformFontsForNodeArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetPlatformFontsForNodeReply represents the return values for GetPlatformFontsForNode in the CSS domain.
type GetPlatformFontsForNodeReply struct {
	Fonts []PlatformFontUsage `json:"fonts"` // Usage statistics for every employed platform font.
}

// GetPlatformFontsForNodeReply returns whether or not the FrameID matches the reply value for GetPlatformFontsForNode in the CSS domain.
func (a *GetPlatformFontsForNodeReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: GetPlatformFontsForNodeReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for GetPlatformFontsForNode in the CSS domain.
func (a *GetPlatformFontsForNodeReply) UnmarshalJSON(b []byte) error {
	type Copy GetPlatformFontsForNodeReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetPlatformFontsForNodeReply(*c)
	return nil
}

// GetStyleSheetTextArgs represents the arguments for GetStyleSheetText in the CSS domain.
type GetStyleSheetTextArgs struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // No description.
}

// Unmarshal the byte array into a return value for GetStyleSheetText in the CSS domain.
func (a *GetStyleSheetTextArgs) UnmarshalJSON(b []byte) error {
	type Copy GetStyleSheetTextArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetStyleSheetTextArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetStyleSheetText in the CSS domain.
func (a *GetStyleSheetTextArgs) MarshalJSON() ([]byte, error) {
	type Copy GetStyleSheetTextArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetStyleSheetTextReply represents the return values for GetStyleSheetText in the CSS domain.
type GetStyleSheetTextReply struct {
	Text string `json:"text"` // The stylesheet text.
}

// GetStyleSheetTextReply returns whether or not the FrameID matches the reply value for GetStyleSheetText in the CSS domain.
func (a *GetStyleSheetTextReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: GetStyleSheetTextReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for GetStyleSheetText in the CSS domain.
func (a *GetStyleSheetTextReply) UnmarshalJSON(b []byte) error {
	type Copy GetStyleSheetTextReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetStyleSheetTextReply(*c)
	return nil
}

// SetEffectivePropertyValueForNodeArgs represents the arguments for SetEffectivePropertyValueForNode in the CSS domain.
type SetEffectivePropertyValueForNodeArgs struct {
	NodeID       dom.NodeID `json:"nodeId"`       // The element id for which to set property.
	PropertyName string     `json:"propertyName"` // No description.
	Value        string     `json:"value"`        // No description.
}

// Unmarshal the byte array into a return value for SetEffectivePropertyValueForNode in the CSS domain.
func (a *SetEffectivePropertyValueForNodeArgs) UnmarshalJSON(b []byte) error {
	type Copy SetEffectivePropertyValueForNodeArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetEffectivePropertyValueForNodeArgs(*c)
	return nil
}

// Marshall the byte array into a return value for SetEffectivePropertyValueForNode in the CSS domain.
func (a *SetEffectivePropertyValueForNodeArgs) MarshalJSON() ([]byte, error) {
	type Copy SetEffectivePropertyValueForNodeArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// SetEffectivePropertyValueForNodeReply represents the return values for SetEffectivePropertyValueForNode in the CSS domain.
type SetEffectivePropertyValueForNodeReply struct {
}

// SetEffectivePropertyValueForNodeReply returns whether or not the FrameID matches the reply value for SetEffectivePropertyValueForNode in the CSS domain.
func (a *SetEffectivePropertyValueForNodeReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: SetEffectivePropertyValueForNodeReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for SetEffectivePropertyValueForNode in the CSS domain.
func (a *SetEffectivePropertyValueForNodeReply) UnmarshalJSON(b []byte) error {
	type Copy SetEffectivePropertyValueForNodeReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetEffectivePropertyValueForNodeReply(*c)
	return nil
}

// SetKeyframeKeyArgs represents the arguments for SetKeyframeKey in the CSS domain.
type SetKeyframeKeyArgs struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // No description.
	Range        SourceRange  `json:"range"`        // No description.
	KeyText      string       `json:"keyText"`      // No description.
}

// Unmarshal the byte array into a return value for SetKeyframeKey in the CSS domain.
func (a *SetKeyframeKeyArgs) UnmarshalJSON(b []byte) error {
	type Copy SetKeyframeKeyArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetKeyframeKeyArgs(*c)
	return nil
}

// Marshall the byte array into a return value for SetKeyframeKey in the CSS domain.
func (a *SetKeyframeKeyArgs) MarshalJSON() ([]byte, error) {
	type Copy SetKeyframeKeyArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// SetKeyframeKeyReply represents the return values for SetKeyframeKey in the CSS domain.
type SetKeyframeKeyReply struct {
	KeyText Value `json:"keyText"` // The resulting key text after modification.
}

// SetKeyframeKeyReply returns whether or not the FrameID matches the reply value for SetKeyframeKey in the CSS domain.
func (a *SetKeyframeKeyReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: SetKeyframeKeyReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for SetKeyframeKey in the CSS domain.
func (a *SetKeyframeKeyReply) UnmarshalJSON(b []byte) error {
	type Copy SetKeyframeKeyReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetKeyframeKeyReply(*c)
	return nil
}

// SetMediaTextArgs represents the arguments for SetMediaText in the CSS domain.
type SetMediaTextArgs struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // No description.
	Range        SourceRange  `json:"range"`        // No description.
	Text         string       `json:"text"`         // No description.
}

// Unmarshal the byte array into a return value for SetMediaText in the CSS domain.
func (a *SetMediaTextArgs) UnmarshalJSON(b []byte) error {
	type Copy SetMediaTextArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetMediaTextArgs(*c)
	return nil
}

// Marshall the byte array into a return value for SetMediaText in the CSS domain.
func (a *SetMediaTextArgs) MarshalJSON() ([]byte, error) {
	type Copy SetMediaTextArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// SetMediaTextReply represents the return values for SetMediaText in the CSS domain.
type SetMediaTextReply struct {
	Media Media `json:"media"` // The resulting CSS media rule after modification.
}

// SetMediaTextReply returns whether or not the FrameID matches the reply value for SetMediaText in the CSS domain.
func (a *SetMediaTextReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: SetMediaTextReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for SetMediaText in the CSS domain.
func (a *SetMediaTextReply) UnmarshalJSON(b []byte) error {
	type Copy SetMediaTextReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetMediaTextReply(*c)
	return nil
}

// SetRuleSelectorArgs represents the arguments for SetRuleSelector in the CSS domain.
type SetRuleSelectorArgs struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // No description.
	Range        SourceRange  `json:"range"`        // No description.
	Selector     string       `json:"selector"`     // No description.
}

// Unmarshal the byte array into a return value for SetRuleSelector in the CSS domain.
func (a *SetRuleSelectorArgs) UnmarshalJSON(b []byte) error {
	type Copy SetRuleSelectorArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetRuleSelectorArgs(*c)
	return nil
}

// Marshall the byte array into a return value for SetRuleSelector in the CSS domain.
func (a *SetRuleSelectorArgs) MarshalJSON() ([]byte, error) {
	type Copy SetRuleSelectorArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// SetRuleSelectorReply represents the return values for SetRuleSelector in the CSS domain.
type SetRuleSelectorReply struct {
	SelectorList SelectorList `json:"selectorList"` // The resulting selector list after modification.
}

// SetRuleSelectorReply returns whether or not the FrameID matches the reply value for SetRuleSelector in the CSS domain.
func (a *SetRuleSelectorReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: SetRuleSelectorReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for SetRuleSelector in the CSS domain.
func (a *SetRuleSelectorReply) UnmarshalJSON(b []byte) error {
	type Copy SetRuleSelectorReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetRuleSelectorReply(*c)
	return nil
}

// SetStyleSheetTextArgs represents the arguments for SetStyleSheetText in the CSS domain.
type SetStyleSheetTextArgs struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // No description.
	Text         string       `json:"text"`         // No description.
}

// Unmarshal the byte array into a return value for SetStyleSheetText in the CSS domain.
func (a *SetStyleSheetTextArgs) UnmarshalJSON(b []byte) error {
	type Copy SetStyleSheetTextArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetStyleSheetTextArgs(*c)
	return nil
}

// Marshall the byte array into a return value for SetStyleSheetText in the CSS domain.
func (a *SetStyleSheetTextArgs) MarshalJSON() ([]byte, error) {
	type Copy SetStyleSheetTextArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// SetStyleSheetTextReply represents the return values for SetStyleSheetText in the CSS domain.
type SetStyleSheetTextReply struct {
	SourceMapURL string `json:"sourceMapURL,omitempty"` // URL of source map associated with script (if any).
}

// SetStyleSheetTextReply returns whether or not the FrameID matches the reply value for SetStyleSheetText in the CSS domain.
func (a *SetStyleSheetTextReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: SetStyleSheetTextReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for SetStyleSheetText in the CSS domain.
func (a *SetStyleSheetTextReply) UnmarshalJSON(b []byte) error {
	type Copy SetStyleSheetTextReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetStyleSheetTextReply(*c)
	return nil
}

// SetStyleTextsArgs represents the arguments for SetStyleTexts in the CSS domain.
type SetStyleTextsArgs struct {
	Edits []StyleDeclarationEdit `json:"edits"` // No description.
}

// Unmarshal the byte array into a return value for SetStyleTexts in the CSS domain.
func (a *SetStyleTextsArgs) UnmarshalJSON(b []byte) error {
	type Copy SetStyleTextsArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetStyleTextsArgs(*c)
	return nil
}

// Marshall the byte array into a return value for SetStyleTexts in the CSS domain.
func (a *SetStyleTextsArgs) MarshalJSON() ([]byte, error) {
	type Copy SetStyleTextsArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// SetStyleTextsReply represents the return values for SetStyleTexts in the CSS domain.
type SetStyleTextsReply struct {
	Styles []Style `json:"styles"` // The resulting styles after modification.
}

// SetStyleTextsReply returns whether or not the FrameID matches the reply value for SetStyleTexts in the CSS domain.
func (a *SetStyleTextsReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: SetStyleTextsReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for SetStyleTexts in the CSS domain.
func (a *SetStyleTextsReply) UnmarshalJSON(b []byte) error {
	type Copy SetStyleTextsReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetStyleTextsReply(*c)
	return nil
}

// StartRuleUsageTrackingArgs represents the arguments for StartRuleUsageTracking in the CSS domain.
type StartRuleUsageTrackingArgs struct {
}

// Unmarshal the byte array into a return value for StartRuleUsageTracking in the CSS domain.
func (a *StartRuleUsageTrackingArgs) UnmarshalJSON(b []byte) error {
	type Copy StartRuleUsageTrackingArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = StartRuleUsageTrackingArgs(*c)
	return nil
}

// Marshall the byte array into a return value for StartRuleUsageTracking in the CSS domain.
func (a *StartRuleUsageTrackingArgs) MarshalJSON() ([]byte, error) {
	type Copy StartRuleUsageTrackingArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// StartRuleUsageTrackingReply represents the return values for StartRuleUsageTracking in the CSS domain.
type StartRuleUsageTrackingReply struct {
}

// StartRuleUsageTrackingReply returns whether or not the FrameID matches the reply value for StartRuleUsageTracking in the CSS domain.
func (a *StartRuleUsageTrackingReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: StartRuleUsageTrackingReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for StartRuleUsageTracking in the CSS domain.
func (a *StartRuleUsageTrackingReply) UnmarshalJSON(b []byte) error {
	type Copy StartRuleUsageTrackingReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = StartRuleUsageTrackingReply(*c)
	return nil
}

// StopRuleUsageTrackingArgs represents the arguments for StopRuleUsageTracking in the CSS domain.
type StopRuleUsageTrackingArgs struct {
}

// Unmarshal the byte array into a return value for StopRuleUsageTracking in the CSS domain.
func (a *StopRuleUsageTrackingArgs) UnmarshalJSON(b []byte) error {
	type Copy StopRuleUsageTrackingArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = StopRuleUsageTrackingArgs(*c)
	return nil
}

// Marshall the byte array into a return value for StopRuleUsageTracking in the CSS domain.
func (a *StopRuleUsageTrackingArgs) MarshalJSON() ([]byte, error) {
	type Copy StopRuleUsageTrackingArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// StopRuleUsageTrackingReply represents the return values for StopRuleUsageTracking in the CSS domain.
type StopRuleUsageTrackingReply struct {
	RuleUsage []RuleUsage `json:"ruleUsage"` // No description.
}

// StopRuleUsageTrackingReply returns whether or not the FrameID matches the reply value for StopRuleUsageTracking in the CSS domain.
func (a *StopRuleUsageTrackingReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: StopRuleUsageTrackingReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for StopRuleUsageTracking in the CSS domain.
func (a *StopRuleUsageTrackingReply) UnmarshalJSON(b []byte) error {
	type Copy StopRuleUsageTrackingReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = StopRuleUsageTrackingReply(*c)
	return nil
}

// TakeCoverageDeltaArgs represents the arguments for TakeCoverageDelta in the CSS domain.
type TakeCoverageDeltaArgs struct {
}

// Unmarshal the byte array into a return value for TakeCoverageDelta in the CSS domain.
func (a *TakeCoverageDeltaArgs) UnmarshalJSON(b []byte) error {
	type Copy TakeCoverageDeltaArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = TakeCoverageDeltaArgs(*c)
	return nil
}

// Marshall the byte array into a return value for TakeCoverageDelta in the CSS domain.
func (a *TakeCoverageDeltaArgs) MarshalJSON() ([]byte, error) {
	type Copy TakeCoverageDeltaArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// TakeCoverageDeltaReply represents the return values for TakeCoverageDelta in the CSS domain.
type TakeCoverageDeltaReply struct {
	Coverage []RuleUsage `json:"coverage"` // No description.
}

// TakeCoverageDeltaReply returns whether or not the FrameID matches the reply value for TakeCoverageDelta in the CSS domain.
func (a *TakeCoverageDeltaReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: TakeCoverageDeltaReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for TakeCoverageDelta in the CSS domain.
func (a *TakeCoverageDeltaReply) UnmarshalJSON(b []byte) error {
	type Copy TakeCoverageDeltaReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = TakeCoverageDeltaReply(*c)
	return nil
}
