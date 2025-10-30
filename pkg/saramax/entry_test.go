package saramax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessage_Encode(t *testing.T) {
	type TestData struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	data := TestData{Name: "test", Age: 25}
	msg := NewEntry(data)

	encoded, err := msg.Encode()
	assert.NoError(t, err)
	assert.NotNil(t, encoded)

	expected := `{"name":"test","age":25}`
	assert.Equal(t, expected, string(encoded))
}

func TestMessage_Length(t *testing.T) {
	type TestData struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	data := TestData{Name: "test", Age: 25}
	msg := NewEntry(data)

	expected := `{"name":"test","age":25}`
	assert.Equal(t, len(expected), msg.Length())
}

func TestMessage_Data(t *testing.T) {
	type TestData struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	data := TestData{Name: "test", Age: 25}
	msg := NewEntry(data)

	assert.Equal(t, data, msg.Data())
}
