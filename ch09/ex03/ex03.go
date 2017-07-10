package memo

import "fmt"

type Cancel chan struct{}

type Func func(key string, cancel Cancel) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res    result
	ready  chan struct{}
	cancel Cancel
}

type request struct {
	key      string
	response chan<- result
	cancel   Cancel
}

type Memo struct{ requests chan request }

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, cancel chan struct{}) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response, cancel}
	var res result
	for {
		select {
		case res = <-response:
			return res.value, res.err
		case <-cancel:
			return nil, fmt.Errorf("canceled")

		}
	}
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key, e.cancel)
	if !e.Cancelled() {
		close(e.ready)
	}
}

func (e *entry) deliver(response chan<- result) {
	select {
	case <-e.ready:
		response <- e.res
	case <-e.cancel:
	}
}

func (e *entry) Cancelled() bool {
	select {
	case <-e.cancel:
		return true
	default:
		return false
	}
}
