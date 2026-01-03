# MacOS Keychain CLI

[![Build Status](https://github.com/adreasnow/keychain-cli/actions/workflows/publish.yaml/badge.svg?branch=main)](https://github.com/adreasnow/keychain-cli/actions/workflows/publish.yaml) ![Go Coverage](https://github.com/adreasnow/keychain-cli/wiki/coverage.svg)

This project is a very simple command-line interface for managing secrets with the MacOS Keychain.

It's not intended for public consumption but can be used if so desired

## Installation

Either install using the provided `package.nix` or `go install github.com/adreasnow/keychain-cli`

### Tab completions

Tab completions are supported for bash, zsh, and fish shells. To enable them add the following to your shell configuration file, replacing <shell> with your shell of choice:

```bash
source <(keychain-cli completion <shell>)
```

If installing via Nix, completions will be auto-enabled

## Usage

```
NAME:
   keychain-cli - A new cli application

USAGE:
   keychain-cli [global options] [command [command options]]

COMMANDS:
   list     List all secrets
   get      Get a specific secret by key
   set      Create/update a secret
   delete   Delete a secret
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```
