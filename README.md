# InjectEnv
Wraps and executes commands with additional environmental variables.
<br/><br/>


## Manpage
```
InjectEnv
Wraps and executes commands with additional environmental variables.

Usage: injectenv <command> [<flags> ...]
Example: injectenv exec -profile profile1 -- printenv | grep key1

Commands:
  help
      Show help.

  list [<flags>]
    List profiles.

	Flags:
	 --profile			If specified the program will list all variables under the target profile.

  exec [<flags>] [<cmd>] [<args>...]
    Executes a command with additional profile vars in the environment

	Flags:
	 --profile			Name of the profile to use (Required).
```


## Demo
Example config
```
 ~/.injectenv.yaml
---
profile1:
    key1: value
    key2: value2
profile2:
    key1: value
    key2: vlaue2

```

List Profiles
```
$ injectenv list
Available Profiles
profile1
profile2
```
Display profile variables
```
$ injectenv list -profile profile1
profile1
key1: value
key2: value2
```
Execute command with injected variables
```
$ injectenv exec -profile profile1 -- printenv | grep key1
key1=value1
```