package adapter

import "sync"

// this is an adapter desigan pattern case from golang net/http standard library,
// the origin code is here: https://github.com/golang/go/blob/go1.18.6/src/net/http/server.go#L2515

// ResponseWriter ...
type ResponseWriter interface{}

// Request ...
type Request struct{}

// IHandler ...
type IHandler interface {
	ServeHTTP(ResponseWriter, *Request)
}

// ServerMux ...
type ServerMux struct {
	routerGroup map[string]IHandler
	lock        sync.RWMutex
}

// Handle ...
func (mux *ServerMux) Handle(pattern string, handler IHandler) {
	mux.lock.Lock()
	defer mux.lock.Unlock()

	if _, ok := mux.routerGroup[pattern]; ok {
		panic("duplicate handler")
	}
	mux.routerGroup[pattern] = handler
}

// HandleFunc ...
func (mux *ServerMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	if handler == nil {
		panic("nil handler")
	}

	// !!!!!!!!!!!
	// this is the core logic handled by adapter pattern
	// !!!!!!!!!!!
	// we can use adapter pattern to resolve this error
	// Analysis:
	// as we know, when we handle a request from client,
	// we need to find the matching handler for url/pattern, and the call the handler.
	// so, we need an adapter which not only could warp the handler function to match the IHandler require,
	// but also call the handler itself.
	mux.Handle(pattern, HandlerFunc(handler))
}

// HandlerFunc ...
type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}

func stdLibLogci() {
	// TODO
}
