// Copyright 2019 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json2test

import (
	"encoding/json"
	"io"
	"time"

	"github.com/u-root/u-root/pkg/uio"
)

// Action is a TestEvent action.
type Action string

// These are the possible Actions generated by `go test -json` or `test2json`.
const (
	Run       Action = "run"
	Pause     Action = "pause"
	Continue  Action = "cont"
	Pass      Action = "pass"
	Fail      Action = "fail"
	Skip      Action = "skip"
	Output    Action = "output"
	Benchmark Action = "benchmark"
)

// TestEvent is an event generated by `test2json`.
type TestEvent struct {
	Time    time.Time `json:",omitempty"`
	Action  Action
	Package string  `json:",omitempty"`
	Test    string  `json:",omitempty"`
	Elapsed float64 `json:",omitempty"`
	Output  string  `json:",omitempty"`
}

// Handler is a callback for TestEvent events.
type Handler interface {
	Handle(e TestEvent)
}

// EventParser returns a Writer that parses events and hands each event to all
// handlers h.
func EventParser(h ...Handler) io.WriteCloser {
	return uio.FullLineWriter(&eventParser{h})
}

type eventParser struct {
	hs []Handler
}

func (ep *eventParser) OneLine(line []byte) {
	// Poor man's JSON detector.
	if len(line) == 0 || line[0] != '{' {
		return
	}

	var e TestEvent
	if err := json.Unmarshal(line, &e); err != nil {
		return
	}
	for _, h := range ep.hs {
		h.Handle(e)
	}
}
