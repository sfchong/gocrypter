# gocrypter
gocrypter is a simple encrypt/decrypt cli tool using AES encryption written in Go. It is built with [Cobra](https://github.com/spf13/cobra).

## Overview
```
$ gocrypter help

Gocrypter is a simple go cli tool that encrypt and decrypt an input string using AES encryption.

Usage:
  gocrypter [command]

Available Commands:
  decrypt     Decrypt a string with a key
  encrypt     Encrypt a string
  help        Help about any command

Flags:
      --config string   config file (default is $HOME/.gocrypter.yaml)
  -h, --help            help for gocrypter
  -t, --toggle          Help message for toggle

Use "gocrypter [command] --help" for more information about a command.
```

## Usage
### Encrypt
```
gocrypter encrypt <input>
```

Example:
```
$ gocrypter encrypt mysecret

Encrypted Text:
a3455c28fbad4c98755df98608627c9a3ff29cfa2517d2b255fdd17e10014f80dd9f8644
Key to decrypt:
41d38fb61ed911be8510a9de1c111157456597dcfb832363092daa9390ba4c9a
```

### Decrypt
```
gocrypter decrypt <input> --key <key>
```

Example:
```
$ gocrypter decrypt a3455c28fbad4c98755df98608627c9a3ff29cfa2517d2b255fdd17e10014f80dd9f8644 -k 41d38fb61ed911be8510a9de1c111157456597dcfb832363092daa9390ba4c9a

mysecret
```