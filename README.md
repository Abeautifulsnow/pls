
# pls

> Impressive Linux commands cheat sheet cli. [Python 版本](https://github.com/chenjiandongx/how)

`pls` 是 [pls](https://github.com/chenjiandongx/pls) 的fork版本，原仓库作者已经停止维护。 
该命令主要用于解决我们平常遗忘一些 `linux` 命令的使用，而不得不去浏览器上找寻相关命令用法的麻烦，故而可以通过该命令查阅你需要知道的命令，并且提供了将命令用法以markdown形式渲染到终端上的能力。

## 安装

## 1) 通过 `go install` 命令安装 

```shell
go install github.com/Abeautifulsnow/pls@latest
```

## 2) 下面是打包好的二进制文件

https://github.com/Abeautifulsnow/pls/releases

## 使用

在安装之后，更改 `pls_<os>_amd64` 的二进制文件名称为 `pls` 或者其他你喜欢的名字 【下面我们通称为 `pls`】。

然后，在 `linux or macos` 系统下，你可以拷贝这个改过名字的二进制文件到 `/usr/local/bin/` 或者 `/usr/bin/` 目录下，这些目录必须在 `PATH` 变量内，这样的话，linux或者macos才会在你运行 `pls` 
命令的时候自动在这个变量内去找寻该二进制文件。

但是在Windows上，你需要通过windows的 `环境变量设置` 去将这个二进制文件所在的目录添加到 `PATH` 变量列表下。

```shell
pls --help
Impressive Linux commands cheat sheet cli.

Usage:
  pls [command]

Available Commands:
  help        Help about any command
  list        List all commands that are available
  search      Search command by keywords
  show        Show the specified command usage.
  upgrade     Upgrade all commands from remote.
  version     Prints the version of pls

Flags:
  -h, --help   help for pls

Use "pls [command] --help" for more information about a command.
```

！！！为了在首次初始化所有的命令，**推荐**先执行下面的命令。

```shell
pls upgrade
```

配置文件位于 `~/.commands/config.json`，你可以进入修改命令的 `<command>.md` 文件所存储的位置，如下：

```shell
cat ~/.commands/config.json
{"dir":"/home/runstone/.commands"}
```

你还可以将linux命令的输出内容输入到 `less` 管道中去查看：

```shell
pls show curl | less
```

## 截图

<img width="976" alt="image" src="https://github.com/Abeautifulsnow/pls/assets/28704977/355acc9d-cba3-4d24-b4c9-03ebdfcaa1f0">
<img width="951" alt="image" src="https://github.com/Abeautifulsnow/pls/assets/28704977/8b523a65-528a-4046-b5f8-9ed388e9e873">


## LICENSE

MIT [©Abeautifulsnow](https://github.com/Abeautifulsnow)
