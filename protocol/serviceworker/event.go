// Code generated by cdpgen. DO NOT EDIT.

package serviceworker

import (
	"encoding/json"
)

const EventServiceWorkerWorkerErrorReported = "ServiceWorker.workerErrorReported"

// WorkerErrorReportedReply is the reply for WorkerErrorReported events.
type WorkerErrorReportedReply struct {
	ErrorMessage ErrorMessage `json:"errorMessage"` // No description.
}

// Unmarshal the byte array into a return value for WorkerErrorReported in the ServiceWorker domain.
func (a *WorkerErrorReportedReply) UnmarshalJSON(b []byte) error {
	type Copy WorkerErrorReportedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = WorkerErrorReportedReply(*c)
	return nil
}

const EventServiceWorkerWorkerRegistrationUpdated = "ServiceWorker.workerRegistrationUpdated"

// WorkerRegistrationUpdatedReply is the reply for WorkerRegistrationUpdated events.
type WorkerRegistrationUpdatedReply struct {
	Registrations []Registration `json:"registrations"` // No description.
}

// Unmarshal the byte array into a return value for WorkerRegistrationUpdated in the ServiceWorker domain.
func (a *WorkerRegistrationUpdatedReply) UnmarshalJSON(b []byte) error {
	type Copy WorkerRegistrationUpdatedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = WorkerRegistrationUpdatedReply(*c)
	return nil
}

const EventServiceWorkerWorkerVersionUpdated = "ServiceWorker.workerVersionUpdated"

// WorkerVersionUpdatedReply is the reply for WorkerVersionUpdated events.
type WorkerVersionUpdatedReply struct {
	Versions []Version `json:"versions"` // No description.
}

// Unmarshal the byte array into a return value for WorkerVersionUpdated in the ServiceWorker domain.
func (a *WorkerVersionUpdatedReply) UnmarshalJSON(b []byte) error {
	type Copy WorkerVersionUpdatedReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = WorkerVersionUpdatedReply(*c)
	return nil
}