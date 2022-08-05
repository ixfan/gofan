package connectors

type SubscribeInterface interface {
	Handle(value string)
}
