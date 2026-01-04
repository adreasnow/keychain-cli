# Keychain CLI

[![Build Status](https://github.com/adreasnow/keys/actions/workflows/publish.yaml/badge.svg?branch=main)](https://github.com/adreasnow/keys/actions/workflows/publish.yaml) ![Go Coverage](https://github.com/adreasnow/keys/wiki/coverage.svg)

This project is a very simple, cross-platform command-line interface for managing secrets with your OS's built-in keychain manager.

It's not intended for public consumption but can be used if so desired

## Installation

Either install using the provided `package.nix` or `go install github.com/adreasnow/keys/cmd/keys@latest`

### Tab completions

Tab completions are supported for bash, zsh, and fish shells. To enable them add the following to your shell configuration file, replacing <shell> with your shell of choice:

```bash
source <(keys completion <shell>)
```

If installing via Nix, completions will be auto-enabled

## Usage

```
NAME:
   keys - A new cli application

USAGE:
   keys [global options] [command [command options]]

COMMANDS:
   list     List all secrets
   get      Get a specific secret by key
   set      Create/update a secret
   delete   Delete a secret
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```
