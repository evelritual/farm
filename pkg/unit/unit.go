package unit

import "time"

type Unit interface {
	Close() error
	Draw() error
	FixedUpdate(time.Duration) error
}
