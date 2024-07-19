package realtime

import (
	"context"
	"slices"
	"sync"

	"github.com/pkg/errors"
)

var (
	ErrInvalidTopic        = errors.New("topic not registered")
	ErrTopicRegistered     = errors.New("topic already registered")
	ErrTopicHasNoListeners = errors.New("topic has no listeners registered")
)

type routeTo byte

const (
	RouteToAll routeTo = iota
	RouteToFirst
)

type Client interface {
	Cleanup() error
	NewTopic(topic string) error

	Send(ctx context.Context, data Message) error
	Listen(topic string) (chan Message, func(), error)
}

type listener struct {
	ch chan Message
	id int
}

type PubSubClient struct {
	lastID      int
	listeners   map[string][]listener
	listenersMx *sync.Mutex
}

func NewClient() Client {
	return &PubSubClient{
		listenersMx: &sync.Mutex{},
		listeners:   make(map[string][]listener),
	}
}

type Message struct {
	Strategy routeTo
	Topic    string
	Data     any
}

func (c *PubSubClient) NewTopic(topic string) error {
	c.listenersMx.Lock()
	defer c.listenersMx.Unlock()

	if c.listeners[topic] != nil {
		return ErrTopicRegistered
	}

	c.listeners[topic] = make([]listener, 0)
	return nil
}

func (c *PubSubClient) Send(ctx context.Context, data Message) error {
	c.listenersMx.Lock()
	listeners, ok := c.listeners[data.Topic]
	c.listenersMx.Unlock()

	if ok {
		if len(listeners) < 1 {
			return ErrTopicHasNoListeners
		}

		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		for _, listener := range listeners {
			if listener.ch == nil {
				continue
			}
			select {
			case listener.ch <- data:
				if data.Strategy == RouteToFirst {
					return nil
				}
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}
	return nil
}

func (c *PubSubClient) Listen(topic string) (chan Message, func(), error) {
	c.listenersMx.Lock()
	defer c.listenersMx.Unlock()

	if c.listeners[topic] == nil {
		return nil, nil, ErrInvalidTopic
	}

	ch := make(chan Message)

	id := c.lastID + 1
	c.lastID = id
	c.listeners[topic] = append(c.listeners[topic], listener{ch, id})

	cancel := func() {
		c.listenersMx.Lock()
		defer c.listenersMx.Unlock()

		index := slices.IndexFunc(c.listeners[topic], func(l listener) bool { return l.id == id })
		if index < 0 {
			return
		}
		c.listeners[topic] = append(c.listeners[topic][:index], c.listeners[topic][index+1:]...)
		close(ch)
	}

	return ch, cancel, nil
}

/*
Shrink the underlying map, removing all topic without a listener
*/
func (c *PubSubClient) Cleanup() error {
	c.listenersMx.Lock()
	defer c.listenersMx.Unlock()

	newMap := make(map[string][]listener)
	for topic, listeners := range c.listeners {
		if len(listeners) < 1 {
			continue
		}
		newMap[topic] = listeners
	}

	c.listeners = newMap
	return nil
}
