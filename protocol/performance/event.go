// Code generated by cdpgen. DO NOT EDIT.

package performance

import (
	"encoding/json"
)

const (
	EventPerformanceMetrics = "Performance.metrics"
)

var EventConstants = map[string]json.Unmarshaler{
	EventPerformanceMetrics: &MetricsReply{},
}

func GetEventReply(event string) (json.Unmarshaler, bool) {
	e, ok := EventConstants[event]
	return e, ok
}

// MetricsReply is the reply for Metrics events.
type MetricsReply struct {
	Metrics []Metric `json:"metrics"` // Current values of the metrics.
	Title   string   `json:"title"`   // Timestamp title.
}

// Unmarshal the byte array into a return value for Metrics in the Performance domain.
func (a *MetricsReply) UnmarshalJSON(b []byte) error {
	type Copy MetricsReply
	c := &Copy{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	*a = MetricsReply(*c)
	return nil
}
