// Package memo provides a concurrency-safe non-blocking memoization
// of a function.  Requests for different keys proceed in parallel.
// Concurrent requests for the same key block until the first completes.
// This implementation uses a monitor goroutine.
package memo

/*
	Extend the Func type and the (*Memo).Get method so that callers may provide
	an optional done channel through which they can cancel the operation (§8.9).
	The results of a cancelled Func call should not be cached.
*/

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

// A result is the result of calling a Func.
type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	response chan<- result   // the client wants a single result
	cancel   <-chan struct{} // cancel the request and don't cache result
}

type Memo struct{ requests chan request }

// New returns a memoization of f.  Clients must subsequently call Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, cancel chan struct{}) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response, cancel}
	select {
	case <-cancel:
		return nil, nil
	case res := <-response:
		return res.value, res.err
	}
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// This is the first request for this key.
			e = &entry{ready: make(chan struct{})}
			go e.call(f, req.key) // call f(key)
			select {
			case <-req.cancel:
				return
			case <-e.ready:
			}
			cache[req.key] = e
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	// Evaluate the function.
	e.res.value, e.res.err = f(key)
	// Broadcast the ready condition.
	close(e.ready)
}

func (e *entry) deliver(response chan<- result, cancel <-chan struct{}) {
	select {
	case <-cancel:
		return
	case <-e.ready: // Wait for the ready condition.
		// Send the result to the client.
		response <- e.res
	}
}
