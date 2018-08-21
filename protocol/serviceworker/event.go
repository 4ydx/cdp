// Code generated by cdpgen. DO NOT EDIT.

package serviceworker

import (
	"encoding/json"
	"log"
)

const (
	EventServiceWorkerWorkerErrorReported       = "ServiceWorker.workerErrorReported"
	EventServiceWorkerWorkerRegistrationUpdated = "ServiceWorker.workerRegistrationUpdated"
	EventServiceWorkerWorkerVersionUpdated      = "ServiceWorker.workerVersionUpdated"
)

type Unmarshaler func() json.Unmarshaler

var EventConstants = map[string]Unmarshaler{
	EventServiceWorkerWorkerErrorReported:       func() json.Unmarshaler { return &WorkerErrorReportedReply{} },
	EventServiceWorkerWorkerRegistrationUpdated: func() json.Unmarshaler { return &WorkerRegistrationUpdatedReply{} },
	EventServiceWorkerWorkerVersionUpdated:      func() json.Unmarshaler { return &WorkerVersionUpdatedReply{} },
}

func GetEventReply(event string) (json.Unmarshaler, bool) {
	e, ok := EventConstants[event]
	if ok {
		return e(), ok
	}
	return nil, ok
}

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

// WorkerErrorReportedReply returns whether or not the FrameID matches the reply value for WorkerErrorReported in the ServiceWorker domain.
func (a *WorkerErrorReportedReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: WorkerErrorReportedReply %s", err)
	}
	return true
}

// WorkerErrorReportedReply returns the FrameID for WorkerErrorReported in the ServiceWorker domain.
func (a *WorkerErrorReportedReply) GetFrameID() string {
	return ""
}

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

// WorkerRegistrationUpdatedReply returns whether or not the FrameID matches the reply value for WorkerRegistrationUpdated in the ServiceWorker domain.
func (a *WorkerRegistrationUpdatedReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: WorkerRegistrationUpdatedReply %s", err)
	}
	return true
}

// WorkerRegistrationUpdatedReply returns the FrameID for WorkerRegistrationUpdated in the ServiceWorker domain.
func (a *WorkerRegistrationUpdatedReply) GetFrameID() string {
	return ""
}

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

// WorkerVersionUpdatedReply returns whether or not the FrameID matches the reply value for WorkerVersionUpdated in the ServiceWorker domain.
func (a *WorkerVersionUpdatedReply) MatchFrameID(frameID string, m []byte) bool {
	err := a.UnmarshalJSON(m)
	if err != nil {
		log.Fatalf("unmarshal error: WorkerVersionUpdatedReply %s", err)
	}
	return true
}

// WorkerVersionUpdatedReply returns the FrameID for WorkerVersionUpdated in the ServiceWorker domain.
func (a *WorkerVersionUpdatedReply) GetFrameID() string {
	return ""
}
