package stream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersectSet(t *testing.T) {
	t.Run("test intersectSet", func(t *testing.T) {
		src := []byte{'a', 'b', 'c', 'd'}
		dst := []byte{'a', 'd', 'e', 'g'}

		set := IntersectSet(src, dst)
		assert.NotEmpty(t, set)
		assert.Equal(t, []byte{'a', 'd'}, set)
	})
}

func TestDiffSet(t *testing.T) {
	t.Run("test diff set", func(t *testing.T) {
		src := []byte{'a', 'b', 'c', 'd'}
		dst := []byte{'a', 'd', 'e', 'g'}

		set := DiffSet(src, dst)
		assert.NotEmpty(t, set)
		assert.Equal(t, []byte{'b', 'c'}, set)
	})
}

func TestFilter(t *testing.T) {
	t.Run("test filter", func(t *testing.T) {
		src := []byte{'a', 'b', 'c', 'd'}

		arr := Filter(src, func(value byte) bool {
			return value == 'a'
		})

		assert.NotEmpty(t, arr)
		assert.Equal(t, []byte{'a'}, arr)
	})
}

func TestObjectMap(t *testing.T) {
	type School struct {
		Name string
		Rank int
	}

	type Person struct {
		Name   string
		Phone  string
		Age    int
		Ids    []string
		School School
		Scores map[string]float64
	}

	type PersonDao struct {
		Name   string
		Phone  string
		Age    int
		School School
		Ids    []string
		Scores map[string]float64
		Ctime  string
		Utime  string
	}
	t.Run("Basic types", func(t *testing.T) {
		person := Person{Name: "person", Phone: "12345678", Age: 12,
			School: School{Name: "xxx", Rank: 1},
			Ids:    []string{"1", "2"}}
		to := MapObject[Person, PersonDao](&person)
		assert.Equal(t, to.Phone, person.Phone)
		assert.Equal(t, to.Name, person.Name)
		assert.Equal(t, to.Age, person.Age)
		assert.Equal(t, to.School.Name, person.School.Name)
		assert.Equal(t, to.School.Rank, person.School.Rank)
		assert.ElementsMatch(t, to.Ids, person.Ids)
	})

	t.Run("Map field", func(t *testing.T) {
		scores := map[string]float64{"math": 90.5, "english": 85.0}
		person := Person{Scores: scores}
		to := MapObject[Person, PersonDao](&person)
		assert.Equal(t, to.Scores["math"], scores["math"])
		assert.Equal(t, to.Scores["english"], scores["english"])
	})

	t.Run("Partial fields", func(t *testing.T) {
		type PartialPerson struct {
			Name string
			Age  int
		}
		person := Person{Name: "partial", Age: 20}
		to := MapObject[Person, PartialPerson](&person)
		assert.Equal(t, to.Name, person.Name)
		assert.Equal(t, to.Age, person.Age)
	})
}
