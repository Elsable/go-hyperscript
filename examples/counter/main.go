package main

import (
	"strconv"
	"syscall/js"

	h "github.com/syumai/go-hyperscript/hyperscript"
	"github.com/syumai/go-hyperscript/dom"
)

var body = js.Global().Get("document").Get("body")
var rootNode h.VNode

type State struct {
	count int
}

var state = State{}

func increment() {
	state.count++
}

func decrement() {
	state.count--
}

func action(f func()) js.Callback {
	return js.NewCallback(func([]js.Value) { f(); update() })
}

func Counter(props h.Object) h.VNode {
	return h.H("div", h.Props{"className": "counter"},
		h.H("div", nil, h.Text(strconv.Itoa(props.Int("count")))),
		h.H("div", nil,
			h.H("button", h.Props{"onclick": action(increment)}, h.Text("+")),
			h.H("button", h.Props{"onclick": action(decrement)}, h.Text("-")),
		),
	)
}

func render() h.VNode {
	return h.H("div", nil,
		h.H("h1", nil, h.Text("Counter")),
		h.H(Counter, h.Props{"count": state.count}),
		h.H("a", h.Props{"href": "https://github.com/syumai/go-hyperscript/"},
			h.Text("Show the code on GitHub"),
		),
	)
}

func update() {
	body.Set("innerHTML", "")
	dom.Render(render(), body)
}

func main() {
	update()
	select {}
}
