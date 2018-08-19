// Code generated by cdpgen. DO NOT EDIT.

package schema

import (
	"encoding/json"
	"log"
)

const CommandSchemaGetDomains = "Schema.getDomains"

// GetDomainsArgs represents the arguments for GetDomains in the Schema domain.
type GetDomainsArgs struct {
}

// Unmarshal the byte array into a return value for GetDomains in the Schema domain.
func (a *GetDomainsArgs) UnmarshalJSON(b []byte) error {
	type Copy GetDomainsArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetDomainsArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetDomains in the Schema domain.
func (a *GetDomainsArgs) MarshalJSON() ([]byte, error) {
	type Copy GetDomainsArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetDomainsReply represents the return values for GetDomains in the Schema domain.
type GetDomainsReply struct {
	Domains []Domain `json:"domains"` // List of supported domains.
}

// GetDomainsReply returns whether or not the FrameID matches the reply value for GetDomains in the Schema domain.
func (a *GetDomainsReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: GetDomainsReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for GetDomains in the Schema domain.
func (a *GetDomainsReply) UnmarshalJSON(b []byte) error {
	type Copy GetDomainsReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetDomainsReply(*c)
	return nil
}
