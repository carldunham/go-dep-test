package coolstuff

import (
	"strings"
	"testing"

	"github.com/google/uuid"
)

func TestCoolThing(t *testing.T) {
	id := uuid.New()
	value := "hello"
	c := CoolThing{ID: id, Value: value}

	s := c.String()

	if !strings.Contains(s, id.String()) {
		t.Errorf("Not seeing id in CoolThing: id=%q, s=%q", id, s)
	}

	if !strings.Contains(s, value) {
		t.Errorf("Not seeing value in CoolThing: value=%q, s=%q", value, s)
	}
}
