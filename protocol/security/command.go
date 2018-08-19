// Code generated by cdpgen. DO NOT EDIT.

package security

import (
	"encoding/json"
)

const CommandSecurityDisable = "Security.disable"

// DisableArgs represents the arguments for Disable in the Security domain.
type DisableArgs struct {
}

// Unmarshal the byte array into a return value for Disable in the Security domain.
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

// Marshall the byte array into a return value for Disable in the Security domain.
func (a *DisableArgs) MarshalJSON() ([]byte, error) {
	type Copy DisableArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// DisableReply represents the return values for Disable in the Security domain.
type DisableReply struct {
}

// Unmarshal the byte array into a return value for Disable in the Security domain.
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

const CommandSecurityEnable = "Security.enable"

// EnableArgs represents the arguments for Enable in the Security domain.
type EnableArgs struct {
}

// Unmarshal the byte array into a return value for Enable in the Security domain.
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

// Marshall the byte array into a return value for Enable in the Security domain.
func (a *EnableArgs) MarshalJSON() ([]byte, error) {
	type Copy EnableArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// EnableReply represents the return values for Enable in the Security domain.
type EnableReply struct {
}

// Unmarshal the byte array into a return value for Enable in the Security domain.
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

const CommandSecuritySetIgnoreCertificateErrors = "Security.setIgnoreCertificateErrors"

// SetIgnoreCertificateErrorsArgs represents the arguments for SetIgnoreCertificateErrors in the Security domain.
type SetIgnoreCertificateErrorsArgs struct {
	Ignore bool `json:"ignore"` // If true, all certificate errors will be ignored.
}

// Unmarshal the byte array into a return value for SetIgnoreCertificateErrors in the Security domain.
func (a *SetIgnoreCertificateErrorsArgs) UnmarshalJSON(b []byte) error {
	type Copy SetIgnoreCertificateErrorsArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetIgnoreCertificateErrorsArgs(*c)
	return nil
}

// Marshall the byte array into a return value for SetIgnoreCertificateErrors in the Security domain.
func (a *SetIgnoreCertificateErrorsArgs) MarshalJSON() ([]byte, error) {
	type Copy SetIgnoreCertificateErrorsArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// SetIgnoreCertificateErrorsReply represents the return values for SetIgnoreCertificateErrors in the Security domain.
type SetIgnoreCertificateErrorsReply struct {
}

// Unmarshal the byte array into a return value for SetIgnoreCertificateErrors in the Security domain.
func (a *SetIgnoreCertificateErrorsReply) UnmarshalJSON(b []byte) error {
	type Copy SetIgnoreCertificateErrorsReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetIgnoreCertificateErrorsReply(*c)
	return nil
}

const CommandSecurityHandleCertificateError = "Security.handleCertificateError"

// HandleCertificateErrorArgs represents the arguments for HandleCertificateError in the Security domain.
type HandleCertificateErrorArgs struct {
	EventID int                    `json:"eventId"` // The ID of the event.
	Action  CertificateErrorAction `json:"action"`  // The action to take on the certificate error.
}

// Unmarshal the byte array into a return value for HandleCertificateError in the Security domain.
func (a *HandleCertificateErrorArgs) UnmarshalJSON(b []byte) error {
	type Copy HandleCertificateErrorArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = HandleCertificateErrorArgs(*c)
	return nil
}

// Marshall the byte array into a return value for HandleCertificateError in the Security domain.
func (a *HandleCertificateErrorArgs) MarshalJSON() ([]byte, error) {
	type Copy HandleCertificateErrorArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// HandleCertificateErrorReply represents the return values for HandleCertificateError in the Security domain.
type HandleCertificateErrorReply struct {
}

// Unmarshal the byte array into a return value for HandleCertificateError in the Security domain.
func (a *HandleCertificateErrorReply) UnmarshalJSON(b []byte) error {
	type Copy HandleCertificateErrorReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = HandleCertificateErrorReply(*c)
	return nil
}

const CommandSecuritySetOverrideCertificateErrors = "Security.setOverrideCertificateErrors"

// SetOverrideCertificateErrorsArgs represents the arguments for SetOverrideCertificateErrors in the Security domain.
type SetOverrideCertificateErrorsArgs struct {
	Override bool `json:"override"` // If true, certificate errors will be overridden.
}

// Unmarshal the byte array into a return value for SetOverrideCertificateErrors in the Security domain.
func (a *SetOverrideCertificateErrorsArgs) UnmarshalJSON(b []byte) error {
	type Copy SetOverrideCertificateErrorsArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetOverrideCertificateErrorsArgs(*c)
	return nil
}

// Marshall the byte array into a return value for SetOverrideCertificateErrors in the Security domain.
func (a *SetOverrideCertificateErrorsArgs) MarshalJSON() ([]byte, error) {
	type Copy SetOverrideCertificateErrorsArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// SetOverrideCertificateErrorsReply represents the return values for SetOverrideCertificateErrors in the Security domain.
type SetOverrideCertificateErrorsReply struct {
}

// Unmarshal the byte array into a return value for SetOverrideCertificateErrors in the Security domain.
func (a *SetOverrideCertificateErrorsReply) UnmarshalJSON(b []byte) error {
	type Copy SetOverrideCertificateErrorsReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = SetOverrideCertificateErrorsReply(*c)
	return nil
}
