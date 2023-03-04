package cor

type Next func(string)

type MessageHandler struct {
	next    *MessageHandler
	handler func(message string)
}

type Middleware struct {
	head *MessageHandler
	tail *MessageHandler
}

func (m *Middleware) Use(handle func(msg string, next Next)) {

	messageHandler := &MessageHandler{
		next: &MessageHandler{handler: func(message string) {}},
	}

	messageHandler.handler = func(msg string) { handle(msg, messageHandler.next.handler) }

	if m.head == nil {
		m.head = messageHandler
		m.tail = messageHandler
	} else {
		m.tail.next = messageHandler
		m.tail = m.tail.next
	}

}

func (m *Middleware) Run(msg string) {
	if m.head == nil {
		return
	}
	m.head.handler(msg)
}
func NewMiddleware() *Middleware {
	return &Middleware{}
}
