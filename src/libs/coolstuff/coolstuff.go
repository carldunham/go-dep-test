package coolstuff

import (
	"fmt"

	"github.com/google/uuid"
)

// CoolThing does what is says
//
type CoolThing struct {
	ID    uuid.UUID
	Value string
}

// ToString is pretty cool too
//
func (c *CoolThing) String() string {
	return fmt.Sprintf("%v: %q", c.ID, c.Value)
}
