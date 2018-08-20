// Code generated by cdpgen. DO NOT EDIT.

package lib

import (
	"encoding/json"

	"github.com/4ydx/cdp/protocol/animation"
	"github.com/4ydx/cdp/protocol/applicationcache"
	"github.com/4ydx/cdp/protocol/console"
	"github.com/4ydx/cdp/protocol/css"
	"github.com/4ydx/cdp/protocol/database"
	"github.com/4ydx/cdp/protocol/debugger"
	"github.com/4ydx/cdp/protocol/dom"
	"github.com/4ydx/cdp/protocol/domstorage"
	"github.com/4ydx/cdp/protocol/emulation"
	"github.com/4ydx/cdp/protocol/headlessexperimental"
	"github.com/4ydx/cdp/protocol/heapprofiler"
	"github.com/4ydx/cdp/protocol/inspector"
	"github.com/4ydx/cdp/protocol/layertree"
	"github.com/4ydx/cdp/protocol/log"
	"github.com/4ydx/cdp/protocol/network"
	"github.com/4ydx/cdp/protocol/overlay"
	"github.com/4ydx/cdp/protocol/page"
	"github.com/4ydx/cdp/protocol/performance"
	"github.com/4ydx/cdp/protocol/profiler"
	"github.com/4ydx/cdp/protocol/runtime"
	"github.com/4ydx/cdp/protocol/security"
	"github.com/4ydx/cdp/protocol/serviceworker"
	"github.com/4ydx/cdp/protocol/storage"
	"github.com/4ydx/cdp/protocol/target"
	"github.com/4ydx/cdp/protocol/tethering"
	"github.com/4ydx/cdp/protocol/tracing"
)

// GetEventUnmarshaler returns an object that can receive and unmarshal event data.
func GetEventUnmarshaler(event string) (json.Unmarshaler, bool) {
	var o json.Unmarshaler
	var ok bool
	o, ok = animation.GetEventReply(event)
	if ok {
		return o, true
	}
	o, ok = applicationcache.GetEventReply(event)
	if ok {
		return o, true
	}
	o, ok = css.GetEventReply(event)
	if ok {
		return o, true
	}
	o, ok = console.GetEventReply(event)
	if ok {
		return o, true
	}
	o, ok = dom.GetEventReply(event)
	if ok {
		return o, true
	}
	o, ok = domstorage.GetEventReply(event)
	if ok {
		return o, true
	}
	o, ok = database.GetEventReply(event)
	if ok {
		return o, true
	}
	o, ok = debugger.GetEventReply(event)
	if ok {
		return o, true
	}
	o, ok = emulation.GetEventReply(event)
	if ok {
		return o, true
	}
	o, ok = headlessexperimental.GetEventReply(event)
	if ok {
		return o, true
	}
	o, ok = heapprofiler.GetEventReply(event)
	if ok {
		return o, true
	}
	o, ok = inspector.GetEventReply(event)
	if ok {
		return o, true
	}
	o, ok = layertree.GetEventReply(event)
	if ok {
		return o, true
	}
	o, ok = log.GetEventReply(event)
	if ok {
		return o, true
	}
	o, ok = network.GetEventReply(event)
	if ok {
		return o, true
	}
	o, ok = overlay.GetEventReply(event)
	if ok {
		return o, true
	}
	o, ok = page.GetEventReply(event)
	if ok {
		return o, true
	}
	o, ok = performance.GetEventReply(event)
	if ok {
		return o, true
	}
	o, ok = profiler.GetEventReply(event)
	if ok {
		return o, true
	}
	o, ok = runtime.GetEventReply(event)
	if ok {
		return o, true
	}
	o, ok = security.GetEventReply(event)
	if ok {
		return o, true
	}
	o, ok = serviceworker.GetEventReply(event)
	if ok {
		return o, true
	}
	o, ok = storage.GetEventReply(event)
	if ok {
		return o, true
	}
	o, ok = target.GetEventReply(event)
	if ok {
		return o, true
	}
	o, ok = tethering.GetEventReply(event)
	if ok {
		return o, true
	}
	o, ok = tracing.GetEventReply(event)
	if ok {
		return o, true
	}
	return nil, false
}
