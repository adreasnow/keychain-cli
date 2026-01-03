package main

import (
	"testing"

	"github.com/adreasnow/keychain-cli/keys"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zalando/go-keyring"
)

func TestObfuscate(t *testing.T) {
	for _, s := range []string{
		"testString1",
		"short1",
		"veryLongStringThatIsProbablyAComplexAPIKeyOrSomethingLikeThat...1",
	} {
		t.Run(s, func(t *testing.T) {
			key := obfuscate(s)
			assert.Len(t, key, len(s))
			assert.NotContains(t, key, "1")
		})
	}
}

func TestSet(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		keyring.MockInit()
		dict := keys.NewDict()

		err := set("testKey", "testSecret")
		assert.NoError(t, err)

		keys, err := dict.GetAllKeys()
		assert.NoError(t, err)
		assert.Len(t, keys, 1)
		assert.Contains(t, keys, "testKey")
	})

	t.Run("missing key", func(t *testing.T) {
		keyring.MockInit()

		err := set("", "testSecret")
		assert.ErrorIs(t, err, ErrMissingKey)
	})

	t.Run("missing secret", func(t *testing.T) {
		keyring.MockInit()

		err := set("testKey", "")
		assert.ErrorIs(t, err, ErrMissingSecret)
	})
}

func TestGet(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		keyring.MockInit()
		dict := keys.NewDict()

		key := "testKey"
		secret := "testSecret"
		err := dict.SetSecret(key, secret)
		require.NoError(t, err)

		err = get(key)
		assert.NoError(t, err)

		keys, err := dict.GetAllKeys()
		require.NoError(t, err)
		require.Len(t, keys, 1)
	})

	t.Run("missing key", func(t *testing.T) {
		keyring.MockInit()

		err := get("")
		assert.ErrorIs(t, err, ErrMissingKey)
	})
}

func TestDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		keyring.MockInit()
		dict := keys.NewDict()

		key := "testKey"
		secret := "testSecret"
		err := dict.SetSecret(key, secret)
		require.NoError(t, err)

		err = delete(key)
		assert.NoError(t, err)

		dict = keys.NewDict()
		keys, err := dict.GetAllKeys()
		assert.NoError(t, err)
		assert.Len(t, keys, 0)
	})

	t.Run("missing key", func(t *testing.T) {
		keyring.MockInit()

		err := delete("")
		assert.ErrorIs(t, err, ErrMissingKey)
	})
}
