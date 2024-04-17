# PasswordGenerator
Create strong passwords with PasswordGenerator  
PasswordGenerator generates strong passwords of specific lengths and with specific characters. Written with the Golang Cobra library.
# Usage
PasswordGenerator generate -l 12 -s -d  
PasswordGenerator generate --length 12 --special --digit
```
Generate strong password passwords with customized options

Usage:
  Password [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  generate    Generate random passwords
  help        Help about any command

Flags:
  -h, --help     help for Password
  -t, --toggle   Help message for toggle

Use "Password [command] --help" for more information about a command.
```
