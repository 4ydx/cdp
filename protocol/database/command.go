// Code generated by cdpgen. DO NOT EDIT.

package database

import (
	"encoding/json"
	"log"
)

const (
	CommandDatabaseDisable               = "Database.disable"
	CommandDatabaseEnable                = "Database.enable"
	CommandDatabaseExecuteSQL            = "Database.executeSQL"
	CommandDatabaseGetDatabaseTableNames = "Database.getDatabaseTableNames"
)

// DisableArgs represents the arguments for Disable in the Database domain.
type DisableArgs struct {
}

// Unmarshal the byte array into a return value for Disable in the Database domain.
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

// Marshall the byte array into a return value for Disable in the Database domain.
func (a *DisableArgs) MarshalJSON() ([]byte, error) {
	type Copy DisableArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// DisableReply represents the return values for Disable in the Database domain.
type DisableReply struct {
}

// DisableReply returns whether or not the FrameID matches the reply value for Disable in the Database domain.
func (a *DisableReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: DisableReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for Disable in the Database domain.
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

// EnableArgs represents the arguments for Enable in the Database domain.
type EnableArgs struct {
}

// Unmarshal the byte array into a return value for Enable in the Database domain.
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

// Marshall the byte array into a return value for Enable in the Database domain.
func (a *EnableArgs) MarshalJSON() ([]byte, error) {
	type Copy EnableArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// EnableReply represents the return values for Enable in the Database domain.
type EnableReply struct {
}

// EnableReply returns whether or not the FrameID matches the reply value for Enable in the Database domain.
func (a *EnableReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: EnableReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for Enable in the Database domain.
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

// ExecuteSQLArgs represents the arguments for ExecuteSQL in the Database domain.
type ExecuteSQLArgs struct {
	DatabaseID ID     `json:"databaseId"` // No description.
	Query      string `json:"query"`      // No description.
}

// Unmarshal the byte array into a return value for ExecuteSQL in the Database domain.
func (a *ExecuteSQLArgs) UnmarshalJSON(b []byte) error {
	type Copy ExecuteSQLArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ExecuteSQLArgs(*c)
	return nil
}

// Marshall the byte array into a return value for ExecuteSQL in the Database domain.
func (a *ExecuteSQLArgs) MarshalJSON() ([]byte, error) {
	type Copy ExecuteSQLArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// ExecuteSQLReply represents the return values for ExecuteSQL in the Database domain.
type ExecuteSQLReply struct {
	ColumnNames []string          `json:"columnNames,omitempty"` // No description.
	Values      []json.RawMessage `json:"values,omitempty"`      // No description.
	SQLError    Error             `json:"sqlError,omitempty"`    // No description.
}

// ExecuteSQLReply returns whether or not the FrameID matches the reply value for ExecuteSQL in the Database domain.
func (a *ExecuteSQLReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: ExecuteSQLReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for ExecuteSQL in the Database domain.
func (a *ExecuteSQLReply) UnmarshalJSON(b []byte) error {
	type Copy ExecuteSQLReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = ExecuteSQLReply(*c)
	return nil
}

// GetDatabaseTableNamesArgs represents the arguments for GetDatabaseTableNames in the Database domain.
type GetDatabaseTableNamesArgs struct {
	DatabaseID ID `json:"databaseId"` // No description.
}

// Unmarshal the byte array into a return value for GetDatabaseTableNames in the Database domain.
func (a *GetDatabaseTableNamesArgs) UnmarshalJSON(b []byte) error {
	type Copy GetDatabaseTableNamesArgs
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetDatabaseTableNamesArgs(*c)
	return nil
}

// Marshall the byte array into a return value for GetDatabaseTableNames in the Database domain.
func (a *GetDatabaseTableNamesArgs) MarshalJSON() ([]byte, error) {
	type Copy GetDatabaseTableNamesArgs
	c := &Copy{}
	*c = Copy(*a)
	return json.Marshal(&c)
}

// GetDatabaseTableNamesReply represents the return values for GetDatabaseTableNames in the Database domain.
type GetDatabaseTableNamesReply struct {
	TableNames []string `json:"tableNames"` // No description.
}

// GetDatabaseTableNamesReply returns whether or not the FrameID matches the reply value for GetDatabaseTableNames in the Database domain.
func (a *GetDatabaseTableNamesReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: GetDatabaseTableNamesReply", err)
	}
	return true
}

// Unmarshal the byte array into a return value for GetDatabaseTableNames in the Database domain.
func (a *GetDatabaseTableNamesReply) UnmarshalJSON(b []byte) error {
	type Copy GetDatabaseTableNamesReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = GetDatabaseTableNamesReply(*c)
	return nil
}
