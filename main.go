package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/adreasnow/keychain-cli/keys"
	"github.com/urfave/cli/v3"
)

var (
	ErrMissingKey    = errors.New("missing key")
	ErrMissingSecret = errors.New("missing secret")
)

func main() {
	var dict = keys.NewDict()
	var key string
	var secret string

	cmd := &cli.Command{
		Name:                  "keychain-cli",
		EnableShellCompletion: true,
		Commands: []*cli.Command{
			{
				Name:  "set",
				Usage: "set a secret",
				Arguments: []cli.Argument{
					&cli.StringArg{Name: "key", Destination: &key},
					&cli.StringArg{Name: "secret", Destination: &secret},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					return set(key, secret, dict)
				},
			},
			{
				Name:  "get",
				Usage: "get a secret",
				Arguments: []cli.Argument{
					&cli.StringArg{Name: "key", Destination: &key},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					return get(key, dict)
				},
				ShellComplete: func(ctx context.Context, cmd *cli.Command) { completion(dict) },
			},

			{
				Name:  "delete",
				Usage: "delete a secret",
				Arguments: []cli.Argument{
					&cli.StringArg{Name: "key", Destination: &key},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					return delete(key, dict)
				},
				ShellComplete: func(ctx context.Context, cmd *cli.Command) { completion(dict) },
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func obfuscate(secret string) string {
	out := strings.Builder{}
	for range secret {
		out.WriteString("*")
	}
	return out.String()
}

func set(key, secret string, dict *keys.Dict) error {
	if key == "" {
		return ErrMissingKey
	}

	if secret == "" {
		return ErrMissingSecret
	}

	err := dict.SetSecret(key, secret)
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, "Set secret %s=%s\n", key, obfuscate(secret)) //nolint:errcheck
	return nil
}

func get(key string, dict *keys.Dict) error {
	if key == "" {
		return ErrMissingKey
	}

	secret, err := dict.GetSecret(key)
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stdout, secret) //nolint:errcheck
	return nil
}

func delete(key string, dict *keys.Dict) error {
	if key == "" {
		return ErrMissingKey
	}

	err := dict.DeleteSecret(key)
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, "Deleted secret %s\n", key) //nolint:errcheck
	return nil
}

func completion(dict *keys.Dict) {
	keys, err := dict.GetAllKeys()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get keys: %v\n", err)
		os.Exit(1)
	}

	for _, key := range keys {
		fmt.Println(key)
	}
}
