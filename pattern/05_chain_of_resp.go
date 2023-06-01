package main

import "fmt"

// Handler интерфейс обработчика
type Handler interface {
	SetNext(handler Handler)
	HandleRequest(request Request)
}

// ConcreteHandlerA конкретный обработчик
type ConcreteHandlerA struct {
	next Handler
}

func (c *ConcreteHandlerA) SetNext(handler Handler) {
	c.next = handler
}

func (c *ConcreteHandlerA) HandleRequest(request Request) {
	if request.Type == "A" {
		fmt.Println("Request handled by ConcreteHandlerA")
	} else if c.next != nil {
		c.next.HandleRequest(request)
	} else {
		fmt.Println("Unable to handle the request")
	}
}

// ConcreteHandlerB конкретный обработчик
type ConcreteHandlerB struct {
	next Handler
}

func (c *ConcreteHandlerB) SetNext(handler Handler) {
	c.next = handler
}

func (c *ConcreteHandlerB) HandleRequest(request Request) {
	if request.Type == "B" {
		fmt.Println("Request handled by ConcreteHandlerB")
	} else if c.next != nil {
		c.next.HandleRequest(request)
	} else {
		fmt.Println("Unable to handle the request")
	}
}

// Request Структура запроса
type Request struct {
	Type string
}

func main() {
	// Создаем обработчики
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}

	// Устанавливаем цепочку вызовов
	handlerA.SetNext(handlerB)

	// Создаем запросы
	requestA := Request{Type: "A"}
	requestB := Request{Type: "B"}
	requestC := Request{Type: "C"}

	// Обрабатываем запросы
	handlerA.HandleRequest(requestA)
	handlerA.HandleRequest(requestB)
	handlerA.HandleRequest(requestC)
}
