package testutil

import (
	"encoding/json"
	"io"
	"sync"
)

func init() {
	knownHandlers = make(map[string]ResponseHandler)

	RegisterResponseHandler("application/json", jsonHandler)
}

var (
	knownHandlersMu sync.Mutex
	knownHandlers   map[string]ResponseHandler
)

type ResponseHandler func(contentType string, raw io.Reader, obj interface{}) error

func RegisterResponseHandler(mime string, handler ResponseHandler) {
	knownHandlersMu.Lock()
	defer knownHandlersMu.Unlock()

	knownHandlers[mime] = handler
}

func getHandler(mime string) ResponseHandler {
	knownHandlersMu.Lock()
	defer knownHandlersMu.Unlock()

	return knownHandlers[mime]
}

// This function assumes that the response contains JSON and unmarshals it
// into the specified object.
func jsonHandler(_ string, r io.Reader, obj interface{}) error {
	return json.NewDecoder(r).Decode(obj)
}
