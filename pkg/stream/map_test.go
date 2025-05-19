package stream

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStructMap(t *testing.T) {
	type Person struct {
		Name  string
		Phone string
		Age   int
	}

	type PersonDao struct {
		Name  string
		Phone string
		test  string
	}

	person := &Person{Name: "person", Phone: "12345678", Age: 12}
	personDao := &PersonDao{Name: "person"}

	err := StructMap(person, personDao)
	require.NoError(t, err)
	assert.Equal(t, personDao.Phone, person.Phone)
	assert.Empty(t, personDao.test)
}
