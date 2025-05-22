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

func TestIntersectSet(t *testing.T) {
	src := []byte{'a', 'b', 'c', 'd'}
	dst := []byte{'a', 'd', 'e', 'g'}

	set := IntersectSet(src, dst)
	assert.NotEmpty(t, set)
	assert.Equal(t, []byte{'a', 'd'}, set)
}

func TestDiffSet(t *testing.T) {
	src := []byte{'a', 'b', 'c', 'd'}
	dst := []byte{'a', 'd', 'e', 'g'}

	set := DiffSet(src, dst)
	assert.NotEmpty(t, set)
	assert.Equal(t, []byte{'b', 'c'}, set)
}

func TestFilter(t *testing.T) {
	src := []byte{'a', 'b', 'c', 'd'}

	arr := Filter(src, func(value byte) bool {
		return value == 'a'
	})

	assert.NotEmpty(t, arr)
	assert.Equal(t, []byte{'a'}, arr)
}
