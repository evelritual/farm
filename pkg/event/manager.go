package event

import "fmt"

var (
	GlobalManager *Manager
)

type Manager struct {
	subscribers map[string]chan *Event
}

func NewManager() *Manager {
	if GlobalManager != nil {
		return GlobalManager
	}

	s := map[string]chan *Event{}
	GlobalManager = &Manager{
		subscribers: s,
	}
	return GlobalManager
}

func (m *Manager) Submit(event Event) {
	for _, s := range m.subscribers {
		s <- &event
	}
}

func (m *Manager) Subscribe(consumerName string) (chan *Event, error) {
	if m.subscribers[consumerName] != nil {
		return nil, fmt.Errorf("invalid name: %s", consumerName)
	}

	c := make(chan *Event, 10)
	m.subscribers[consumerName] = c

	return c, nil
}

func (m *Manager) Unsubscribe(consumerName string) error {
	c := m.subscribers[consumerName]
	if c == nil {
		return fmt.Errorf("invalid name: %s", consumerName)
	}

	close(c)
	delete(m.subscribers, consumerName)
	return nil
}
