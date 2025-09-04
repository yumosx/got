package handlefn

//func TestHashBcrypt(t *testing.T) {
//	b := NewBcrypt(bcrypt.DefaultCost)
//	t.Run("hash", func(t *testing.T) {
//		password := "testpassword123"
//		_, err := b.HashSecret(password)
//		require.NoError(t, err)
//	})
//}
//
//func TestAuthenticate(t *testing.T) {
//	b := NewBcrypt(bcrypt.DefaultCost)
//	t.Run("authenticate", func(t *testing.T) {
//		password := "testpassword123"
//		_, err := b.HashSecret(password)
//		require.NoError(t, result.Error())
//		hashValue := result.Val
//		isSame := b.Authenticate(hashValue, password)
//		require.NoError(t, isSame.Error())
//		require.Equal(t, true, isSame.Val)
//	})
//}
