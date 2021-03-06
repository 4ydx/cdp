// Code generated by cdpgen. DO NOT EDIT.

package webauthn

import (
	"encoding/json"
	"log"
)

const (
	CommandWebAuthnEnable                     = "WebAuthn.enable"
	CommandWebAuthnDisable                    = "WebAuthn.disable"
	CommandWebAuthnAddVirtualAuthenticator    = "WebAuthn.addVirtualAuthenticator"
	CommandWebAuthnRemoveVirtualAuthenticator = "WebAuthn.removeVirtualAuthenticator"
	CommandWebAuthnAddCredential              = "WebAuthn.addCredential"
	CommandWebAuthnGetCredential              = "WebAuthn.getCredential"
	CommandWebAuthnGetCredentials             = "WebAuthn.getCredentials"
	CommandWebAuthnRemoveCredential           = "WebAuthn.removeCredential"
	CommandWebAuthnClearCredentials           = "WebAuthn.clearCredentials"
	CommandWebAuthnSetUserVerified            = "WebAuthn.setUserVerified"
)

// EnableArgs represents the arguments for Enable in the WebAuthn domain.
type EnableArgs struct {
}

// Unmarshal the byte array into a return value for Enable in the WebAuthn domain.
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

// Marshall the byte array into a return value for Enable in the WebAuthn domain.
func (a *EnableArgs) MarshalJSON() ([]byte, error) {
	type Copy EnableArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// EnableReply represents the return values for Enable in the WebAuthn domain.
type EnableReply struct {
}

// EnableReply returns whether or not the FrameID matches the reply value for Enable in the WebAuthn domain.
func (a *EnableReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: EnableReply %s", err)
		return false, err
	}
	return true, nil
}

// EnableReply returns the FrameID value for Enable in the WebAuthn domain.
func (a *EnableReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for Enable in the WebAuthn domain.
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

// DisableArgs represents the arguments for Disable in the WebAuthn domain.
type DisableArgs struct {
}

// Unmarshal the byte array into a return value for Disable in the WebAuthn domain.
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

// Marshall the byte array into a return value for Disable in the WebAuthn domain.
func (a *DisableArgs) MarshalJSON() ([]byte, error) {
	type Copy DisableArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// DisableReply represents the return values for Disable in the WebAuthn domain.
type DisableReply struct {
}

// DisableReply returns whether or not the FrameID matches the reply value for Disable in the WebAuthn domain.
func (a *DisableReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: DisableReply %s", err)
		return false, err
	}
	return true, nil
}

// DisableReply returns the FrameID value for Disable in the WebAuthn domain.
func (a *DisableReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for Disable in the WebAuthn domain.
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

// AddVirtualAuthenticatorArgs represents the arguments for AddVirtualAuthenticator in the WebAuthn domain.
type AddVirtualAuthenticatorArgs struct {
	Options VirtualAuthenticatorOptions `json:"options"` // No description.
}

// Unmarshal the byte array into a return value for AddVirtualAuthenticator in the WebAuthn domain.
func (a *AddVirtualAuthenticatorArgs) UnmarshalJSON(b []byte) error {
	type Copy AddVirtualAuthenticatorArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = AddVirtualAuthenticatorArgs(*c)
	return nil
}

// Marshall the byte array into a return value for AddVirtualAuthenticator in the WebAuthn domain.
func (a *AddVirtualAuthenticatorArgs) MarshalJSON() ([]byte, error) {
	type Copy AddVirtualAuthenticatorArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// AddVirtualAuthenticatorReply represents the return values for AddVirtualAuthenticator in the WebAuthn domain.
type AddVirtualAuthenticatorReply struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"` // No description.
}

// AddVirtualAuthenticatorReply returns whether or not the FrameID matches the reply value for AddVirtualAuthenticator in the WebAuthn domain.
func (a *AddVirtualAuthenticatorReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: AddVirtualAuthenticatorReply %s", err)
		return false, err
	}
	return true, nil
}

// AddVirtualAuthenticatorReply returns the FrameID value for AddVirtualAuthenticator in the WebAuthn domain.
func (a *AddVirtualAuthenticatorReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for AddVirtualAuthenticator in the WebAuthn domain.
func (a *AddVirtualAuthenticatorReply) UnmarshalJSON(b []byte) error {
	type Copy AddVirtualAuthenticatorReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = AddVirtualAuthenticatorReply(*c)
	return nil
}

// RemoveVirtualAuthenticatorArgs represents the arguments for RemoveVirtualAuthenticator in the WebAuthn domain.
type RemoveVirtualAuthenticatorArgs struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"` // No description.
}

// Unmarshal the byte array into a return value for RemoveVirtualAuthenticator in the WebAuthn domain.
func (a *RemoveVirtualAuthenticatorArgs) UnmarshalJSON(b []byte) error {
	type Copy RemoveVirtualAuthenticatorArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = RemoveVirtualAuthenticatorArgs(*c)
	return nil
}

// Marshall the byte array into a return value for RemoveVirtualAuthenticator in the WebAuthn domain.
func (a *RemoveVirtualAuthenticatorArgs) MarshalJSON() ([]byte, error) {
	type Copy RemoveVirtualAuthenticatorArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// RemoveVirtualAuthenticatorReply represents the return values for RemoveVirtualAuthenticator in the WebAuthn domain.
type RemoveVirtualAuthenticatorReply struct {
}

// RemoveVirtualAuthenticatorReply returns whether or not the FrameID matches the reply value for RemoveVirtualAuthenticator in the WebAuthn domain.
func (a *RemoveVirtualAuthenticatorReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: RemoveVirtualAuthenticatorReply %s", err)
		return false, err
	}
	return true, nil
}

// RemoveVirtualAuthenticatorReply returns the FrameID value for RemoveVirtualAuthenticator in the WebAuthn domain.
func (a *RemoveVirtualAuthenticatorReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for RemoveVirtualAuthenticator in the WebAuthn domain.
func (a *RemoveVirtualAuthenticatorReply) UnmarshalJSON(b []byte) error {
	type Copy RemoveVirtualAuthenticatorReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = RemoveVirtualAuthenticatorReply(*c)
	return nil
}

// AddCredentialArgs represents the arguments for AddCredential in the WebAuthn domain.
type AddCredentialArgs struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"` // No description.
	Credential      Credential      `json:"credential"`      // No description.
}

// Unmarshal the byte array into a return value for AddCredential in the WebAuthn domain.
func (a *AddCredentialArgs) UnmarshalJSON(b []byte) error {
	type Copy AddCredentialArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = AddCredentialArgs(*c)
	return nil
}

// Marshall the byte array into a return value for AddCredential in the WebAuthn domain.
func (a *AddCredentialArgs) MarshalJSON() ([]byte, error) {
	type Copy AddCredentialArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// AddCredentialReply represents the return values for AddCredential in the WebAuthn domain.
type AddCredentialReply struct {
}

// AddCredentialReply returns whether or not the FrameID matches the reply value for AddCredential in the WebAuthn domain.
func (a *AddCredentialReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: AddCredentialReply %s", err)
		return false, err
	}
	return true, nil
}

// AddCredentialReply returns the FrameID value for AddCredential in the WebAuthn domain.
func (a *AddCredentialReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for AddCredential in the WebAuthn domain.
func (a *AddCredentialReply) UnmarshalJSON(b []byte) error {
	type Copy AddCredentialReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = AddCredentialReply(*c)
	return nil
}

// GetCredentialArgs represents the arguments for GetCredential in the WebAuthn domain.
type GetCredentialArgs struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"` // No description.
	CredentialID    string          `json:"credentialId"`    // No description.
}

// Unmarshal the byte array into a return value for GetCredential in the WebAuthn domain.
func (a *GetCredentialArgs) UnmarshalJSON(b []byte) error {
	type Copy GetCredentialArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetCredentialArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetCredential in the WebAuthn domain.
func (a *GetCredentialArgs) MarshalJSON() ([]byte, error) {
	type Copy GetCredentialArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetCredentialReply represents the return values for GetCredential in the WebAuthn domain.
type GetCredentialReply struct {
	Credential Credential `json:"credential"` // No description.
}

// GetCredentialReply returns whether or not the FrameID matches the reply value for GetCredential in the WebAuthn domain.
func (a *GetCredentialReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: GetCredentialReply %s", err)
		return false, err
	}
	return true, nil
}

// GetCredentialReply returns the FrameID value for GetCredential in the WebAuthn domain.
func (a *GetCredentialReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for GetCredential in the WebAuthn domain.
func (a *GetCredentialReply) UnmarshalJSON(b []byte) error {
	type Copy GetCredentialReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetCredentialReply(*c)
	return nil
}

// GetCredentialsArgs represents the arguments for GetCredentials in the WebAuthn domain.
type GetCredentialsArgs struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"` // No description.
}

// Unmarshal the byte array into a return value for GetCredentials in the WebAuthn domain.
func (a *GetCredentialsArgs) UnmarshalJSON(b []byte) error {
	type Copy GetCredentialsArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetCredentialsArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetCredentials in the WebAuthn domain.
func (a *GetCredentialsArgs) MarshalJSON() ([]byte, error) {
	type Copy GetCredentialsArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetCredentialsReply represents the return values for GetCredentials in the WebAuthn domain.
type GetCredentialsReply struct {
	Credentials []Credential `json:"credentials"` // No description.
}

// GetCredentialsReply returns whether or not the FrameID matches the reply value for GetCredentials in the WebAuthn domain.
func (a *GetCredentialsReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: GetCredentialsReply %s", err)
		return false, err
	}
	return true, nil
}

// GetCredentialsReply returns the FrameID value for GetCredentials in the WebAuthn domain.
func (a *GetCredentialsReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for GetCredentials in the WebAuthn domain.
func (a *GetCredentialsReply) UnmarshalJSON(b []byte) error {
	type Copy GetCredentialsReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetCredentialsReply(*c)
	return nil
}

// RemoveCredentialArgs represents the arguments for RemoveCredential in the WebAuthn domain.
type RemoveCredentialArgs struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"` // No description.
	CredentialID    string          `json:"credentialId"`    // No description.
}

// Unmarshal the byte array into a return value for RemoveCredential in the WebAuthn domain.
func (a *RemoveCredentialArgs) UnmarshalJSON(b []byte) error {
	type Copy RemoveCredentialArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = RemoveCredentialArgs(*c)
	return nil
}

// Marshall the byte array into a return value for RemoveCredential in the WebAuthn domain.
func (a *RemoveCredentialArgs) MarshalJSON() ([]byte, error) {
	type Copy RemoveCredentialArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// RemoveCredentialReply represents the return values for RemoveCredential in the WebAuthn domain.
type RemoveCredentialReply struct {
}

// RemoveCredentialReply returns whether or not the FrameID matches the reply value for RemoveCredential in the WebAuthn domain.
func (a *RemoveCredentialReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: RemoveCredentialReply %s", err)
		return false, err
	}
	return true, nil
}

// RemoveCredentialReply returns the FrameID value for RemoveCredential in the WebAuthn domain.
func (a *RemoveCredentialReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for RemoveCredential in the WebAuthn domain.
func (a *RemoveCredentialReply) UnmarshalJSON(b []byte) error {
	type Copy RemoveCredentialReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = RemoveCredentialReply(*c)
	return nil
}

// ClearCredentialsArgs represents the arguments for ClearCredentials in the WebAuthn domain.
type ClearCredentialsArgs struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"` // No description.
}

// Unmarshal the byte array into a return value for ClearCredentials in the WebAuthn domain.
func (a *ClearCredentialsArgs) UnmarshalJSON(b []byte) error {
	type Copy ClearCredentialsArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ClearCredentialsArgs(*c)
	return nil
}

// Marshall the byte array into a return value for ClearCredentials in the WebAuthn domain.
func (a *ClearCredentialsArgs) MarshalJSON() ([]byte, error) {
	type Copy ClearCredentialsArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// ClearCredentialsReply represents the return values for ClearCredentials in the WebAuthn domain.
type ClearCredentialsReply struct {
}

// ClearCredentialsReply returns whether or not the FrameID matches the reply value for ClearCredentials in the WebAuthn domain.
func (a *ClearCredentialsReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: ClearCredentialsReply %s", err)
		return false, err
	}
	return true, nil
}

// ClearCredentialsReply returns the FrameID value for ClearCredentials in the WebAuthn domain.
func (a *ClearCredentialsReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for ClearCredentials in the WebAuthn domain.
func (a *ClearCredentialsReply) UnmarshalJSON(b []byte) error {
	type Copy ClearCredentialsReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ClearCredentialsReply(*c)
	return nil
}

// SetUserVerifiedArgs represents the arguments for SetUserVerified in the WebAuthn domain.
type SetUserVerifiedArgs struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"` // No description.
	IsUserVerified  bool            `json:"isUserVerified"`  // No description.
}

// Unmarshal the byte array into a return value for SetUserVerified in the WebAuthn domain.
func (a *SetUserVerifiedArgs) UnmarshalJSON(b []byte) error {
	type Copy SetUserVerifiedArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetUserVerifiedArgs(*c)
	return nil
}

// Marshall the byte array into a return value for SetUserVerified in the WebAuthn domain.
func (a *SetUserVerifiedArgs) MarshalJSON() ([]byte, error) {
	type Copy SetUserVerifiedArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// SetUserVerifiedReply represents the return values for SetUserVerified in the WebAuthn domain.
type SetUserVerifiedReply struct {
}

// SetUserVerifiedReply returns whether or not the FrameID matches the reply value for SetUserVerified in the WebAuthn domain.
func (a *SetUserVerifiedReply) MatchFrameID(frameID string, m []byte) (bool, error) {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Printf("unmarshal error: SetUserVerifiedReply %s", err)
		return false, err
	}
	return true, nil
}

// SetUserVerifiedReply returns the FrameID value for SetUserVerified in the WebAuthn domain.
func (a *SetUserVerifiedReply) GetFrameID() string {
	return ""
}

// Unmarshal the byte array into a return value for SetUserVerified in the WebAuthn domain.
func (a *SetUserVerifiedReply) UnmarshalJSON(b []byte) error {
	type Copy SetUserVerifiedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetUserVerifiedReply(*c)
	return nil
}
