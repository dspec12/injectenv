# InjectEnv
Injectenv adds profile defined variables to your environment for a single command.
<br/><br/>


## Manpage
```
Injectenv adds profile defined variables to your environment for a single command

Example:
  injectenv exec profile1 -- printenv | grep key1

Usage:
  injectenv [command]

Available Commands:
  exec        Executes a command with specified profile variables added to the current environment
  help        Help about any command
  list        Lists profiles and optionally their variables

Flags:
  -c, --config string   config file (default is $HOME/.injectenv.yaml)
  -h, --help            help for injectenv
  -v, --version         version for injectenv

Use "injectenv [command] --help" for more information about a command.
```
