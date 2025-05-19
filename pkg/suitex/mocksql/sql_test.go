package mocksql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSQLAdapter(t *testing.T) {
	adapter, err := NewSQLAdapter()
	assert.NoError(t, err)
	assert.NotNil(t, adapter)
	assert.NotNil(t, adapter.db)
	assert.NotNil(t, adapter.mock)
}
