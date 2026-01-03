package keys

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/zalando/go-keyring"
)

var (
	user       = os.Getenv("USER")
	keysSecret = "keychain-cli-keys"
)

type Dict struct {
	keys map[string]bool
}

func NewDict() (d *Dict) {
	d = &Dict{
		keys: map[string]bool{},
	}

	err := d.loadKeys()
	if err != nil {
		if !errors.Is(err, keyring.ErrNotFound) {
			fmt.Fprintf(os.Stderr, "Failed to load keys: %v\n", err)
			os.Exit(1)
		}

		if err := d.saveKeys(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to initialise keys: %v\n", err)
			os.Exit(1)
		}
	}

	return
}

func (d *Dict) loadKeys() (err error) {
	keyString, err := keyring.Get(keysSecret, user)
	if err != nil {
		return fmt.Errorf("failed to get keychain: %w", err)
	}

	err = json.Unmarshal([]byte(keyString), &d.keys)
	if err != nil {
		return fmt.Errorf("failed to unmarshal keys %s: %w", keyString, err)
	}

	return nil
}

func (d *Dict) saveKeys() (err error) {
	keyString, err := json.Marshal(d.keys)
	if err != nil {
		return fmt.Errorf("failed to marshal keys: %w", err)
	}

	err = keyring.Set(keysSecret, user, string(keyString))
	if err != nil {
		return fmt.Errorf("failed to set keychain: %w", err)
	}

	return nil
}

func (d *Dict) GetAllKeys() (out []string, err error) {
	if err := d.loadKeys(); err != nil {
		return nil, err
	}

	for key := range d.keys {
		out = append(out, key)
	}

	return
}

func (d *Dict) AddKey(key string) (err error) {
	if err := d.loadKeys(); err != nil {
		return err
	}

	d.keys[key] = true

	if err := d.saveKeys(); err != nil {
		return err
	}

	return nil
}

func (d *Dict) DeleteKey(key string) (err error) {
	if err := d.loadKeys(); err != nil {
		return err
	}

	delete(d.keys, key)

	if err := d.saveKeys(); err != nil {
		return err
	}

	return nil
}
