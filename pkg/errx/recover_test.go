package errx

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRecover(t *testing.T) {
	t.Run("测试 recover", func(t *testing.T) {
		defer Recover(func(err any) {
			require.Error(t, err.(error))
		})()

		panic(errors.New("hello, world"))
	})
}
