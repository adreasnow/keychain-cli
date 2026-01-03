package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/adreasnow/keychain-cli/keys"
	"github.com/urfave/cli/v3"
)

var (
	ErrMissingKey    = errors.New("missing key")
	ErrMissingSecret = errors.New("missing secret")
)

func main() {
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
				Action: func(_ context.Context, _ *cli.Command) error {
					return set(key, secret)
				},
			},
			{
				Name:  "get",
				Usage: "get a secret",
				Arguments: []cli.Argument{
					&cli.StringArg{Name: "key", Destination: &key},
				},
				Action: func(_ context.Context, _ *cli.Command) error {
					return get(key)
				},
				ShellComplete: func(_ context.Context, cmd *cli.Command) { completion(cmd) },
			},

			{
				Name:  "delete",
				Usage: "delete a secret",
				Arguments: []cli.Argument{
					&cli.StringArg{Name: "key", Destination: &key},
				},
				Action: func(_ context.Context, _ *cli.Command) error {
					return delete(key)
				},
				ShellComplete: func(_ context.Context, cmd *cli.Command) { completion(cmd) },
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

func set(key, secret string) error {
	var dict = keys.NewDict()

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

func get(key string) error {
	var dict = keys.NewDict()

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

func delete(key string) error {
	var dict = keys.NewDict()

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

func completion(cmd *cli.Command) {
	var dict = keys.NewDict()

	keys, err := dict.GetAllKeys()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get keys: %v\n", err)
		os.Exit(1)
	}

	args := cmd.Args().Slice()
	for _, key := range keys {
		if slices.Contains(args, key) {
			return
		}
	}

	for _, key := range keys {
		fmt.Println(key)
	}
}
