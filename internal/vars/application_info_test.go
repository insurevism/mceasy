package vars

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplicationInfo(t *testing.T) {
	name := ApplicationName()
	assert.NotNil(t, name)
	assert.Equal(t, name, "Mceasy") //default value
}
