package types

import "fmt"

// Event is used for type of events
type Event string

func (e *Event) String() string {
	return fmt.Sprint(*e)
}
