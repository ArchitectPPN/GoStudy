package main

import "fmt"

type ListenParams struct {
	MsgId     string `json:"msg_id"`
	EventType string `json:"event_type"`
	EventKey  string `json:"event_key"`
}

type BaseHandle interface {
	Handle(params ListenParams)
}

type Register struct {
}

func (reg *Register) Handle(params ListenParams) {
	fmt.Println(params.MsgId)
}

func main() {
	eventList := make(map[string]BaseHandle, 1)
	eventList["register"] = &Register{}

	instance := eventList["register"]

	instance.Handle(ListenParams{MsgId: "1231231"})
}
