package main

import (
	"encoding/json"

	"github.com/gopherjs/eventsource"
	"github.com/gopherjs/gopherjs/js"
	"github.com/jraedisch/go_sse_example/events"
	"honnef.co/go/js/dom"
	"honnef.co/go/js/xhr"
)

func main() {
	println("initializing")
	es := eventsource.New("/events")
	doc := dom.GetWindow().Document()
	es.AddEventListener("message", false, generateUListListener(doc, "root"))
	es.AddEventListener("error", false, generateErrorListener(es))

	buttonID := "button"
	btn, ok := doc.GetElementByID(buttonID).(*dom.HTMLButtonElement)
	if !ok {
		println("%s is not a button", buttonID)
	}
	btn.AddEventListener("click", false, sayHi)
}

func sayHi(cmd dom.Event) {
	go func() {
		// js/xhr bindings are much smaller than including net/http. Compare by uncommenting!
		// resp, err := http.Post("/command", "application/json", strings.NewReader(`{"msg":"hi!"}`))
		// defer resp.Body.Close()
		msg := &events.Message{Text: "hi!"}
		jsn, err := json.Marshal(msg)
		if err != nil {
			println(err)
		}
		_, err = xhr.Send("POST", "/command", jsn)
		if err != nil {
			println(err)
		}
	}()
}

func generateErrorListener(es *eventsource.EventSource) func(*js.Object) {
	return func(obj *js.Object) {
		// TODO: Retry connection by creating a new event source after a while here.
		println(obj)
	}
}

// note usage of first class functions here
func generateUListListener(doc dom.Document, elementID string) func(*js.Object) {
	el := doc.GetElementByID(elementID)
	ul, ok := el.(*dom.HTMLUListElement)
	if !ok {
		println("casting error:")
		println(el)
	}

	return func(obj *js.Object) {
		msg := &events.Message{}
		json.Unmarshal([]byte(obj.Get("data").String()), msg)
		ul.AppendChild(listItem(doc, msg.Text))
		// including package json hugely increases js filesize. Use this for comparison:
		// ul.AppendChild(listItem(doc, obj.Get("data").String()))
	}
}

func listItem(doc dom.Document, text string) *dom.HTMLLIElement {
	li := doc.CreateElement("li").(*dom.HTMLLIElement)
	li.SetTextContent(text)
	return li
}
