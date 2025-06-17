package handlefn

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestHashBcrypt(t *testing.T) {
	b := NewBcrypt(bcrypt.DefaultCost)
	t.Run("hash", func(t *testing.T) {
		password := "testpassword123"
		result := b.HashSecret(password)
		require.NoError(t, result.Error())
	})
}

func TestAuthenticate(t *testing.T) {
	b := NewBcrypt(bcrypt.DefaultCost)
	t.Run("authenticate", func(t *testing.T) {
		password := "testpassword123"
		result := b.HashSecret(password)
		require.NoError(t, result.Error())
		hashValue := result.Val
		isSame := b.Authenticate(hashValue, password)
		require.NoError(t, isSame.Error())
		require.Equal(t, true, isSame.Val)
	})
}
