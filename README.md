
# pls

> Impressive Linux commands cheat sheet cli. [Python 版本](https://github.com/chenjiandongx/how)

## Installation

## 1) Install it by `go install` command 

```shell
go install github.com/Abeautifulsnow/pls@latest
```

## 2) The compiled binary file here

https://github.com/Abeautifulsnow/pls/releases

## Usages

After the installation, change the `pls_<os>_amd64` binary file name to `pls` or other name you like[Let's say you name it `pls`].

And then, on `linux or macos` system, you can copy the binary file to `/usr/local/bin/` or `/usr/bin/` directory which must be under the `PATH` variable that Linux or macos will search for executables when running a command.

But on windows, you need to add it to the `PATH` variable list through windows interface.

```shell
pls --help
Impressive Linux commands cheat sheet cli.

Usage:
  pls [command]

Available Commands:
  help        Help about any command
  search      Search command by keywords
  show        Show the specified command usage.
  upgrade     Upgrade all commands from remote.
  version     Prints the version of pls

Flags:
  -h, --help   help for pls

Use "pls [command] --help" for more information about a command.
```

To initialize all commands the first time you run them is strong recommended.
```shell
pls upgrade
```

The config file is located in `~/.commands/config.json`, you can change the download directory in it.

```shell
cat ~/.commands/config.json
{"dir":"/home/runstone/.commands"}
```

You can also pass the output to the `less` pipe.
```shell
pls show curl | less
```

## ScreenShot

<img width="976" alt="image" src="https://github.com/Abeautifulsnow/pls/assets/28704977/355acc9d-cba3-4d24-b4c9-03ebdfcaa1f0">
<img width="951" alt="image" src="https://github.com/Abeautifulsnow/pls/assets/28704977/8b523a65-528a-4046-b5f8-9ed388e9e873">


## LICENSE

MIT [©Abeautifulsnow](https://github.com/Abeautifulsnow)
