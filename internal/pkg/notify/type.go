package notify

import (
	"context"
	"sync"
)

type Publisher[T any] interface {
	Subscribe(event *int, subscriber Subscriber[T])
	UnSubscribe(subscriber Subscriber[T])
}

type Subscriber[T any] interface {
	Notify(ctx context.Context, event uint32, payload T) Response[T]
}

type Notify[T any] struct { // nolint:decorder
	subscriberMap map[uint32][]Subscriber[T]
	sync.RWMutex
}

type Response[T any] struct { // nolint:decorder
	Name    string
	Payload T
	Error   error
}

type Callback struct { // nolint:decorder
	CB             chan<- interface{}
	ResponseFilter string
}
