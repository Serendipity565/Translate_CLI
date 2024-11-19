# Translate_CLI

**Translate** 是一个易于使用的命令行翻译工具。它支持多种语言，能够直接在终端中快速、准确地进行翻译。

## 文档

完整文档提供多种语言版本：

- [中文 README](https://github.com/Serendipity565/Translate_CLI/README.md)
- [英文 README](https://github.com/Serendipity565/Translate_CLI/docs/README_zh.md)

## 功能特点

* 支持多种语言
* 快速可靠的翻译
* 简单直观的命令行界面
* 为常用 Shell 提供自动补全脚本
* 支持多种翻译 API

## 安装方法

1. 下载并解压 [下载链接](https://github.com/Serendipity565/Translate_CLI/releases) 中的压缩包。
2. 解压之后，将 `config` 文件夹中的 `user.yaml` 配置文件中的内容修改为你自己的配置。
3. 将 `tran.exe` 的路径加入系统环境变量。例如：将 `C:\translate` 添加到环境变量中。

## 使用方法

通过以下命令轻松在不同语言之间进行翻译：

```bash
tran [flags]
tran [command]
```

### 可用命令

| 命令         | 描述                                   |
|--------------|----------------------------------------|
| `completion` | 为指定的 Shell 生成自动补全脚本         |
| `en-zh`      | 将英文翻译为中文                       |
| `zh-en`      | 将中文翻译为英文                       |
| `help`       | 查看任意命令的帮助信息                 |

### 可用参数

| 参数            | 描述                                   |
|-----------------|----------------------------------------|
| `-h`, `--help`  | 显示 `translate` 工具的帮助信息        |
| `-v`, `--version` | 显示版本号并退出                     |

## 示例

将英文翻译成中文：

```bash
tran en-zh "Hello, world!"
```

将中文翻译成英文：

```bash
tran zh-en "你好，世界！"
```

查看特定命令的帮助信息：

```bash
tran [command] --help
tran [command] -h
```

查看当前版本：

```bash
tran [command] --version
tran [command] -v
```

## 文档

完整文档可在 [此处](https://github.com/Serendipity565/Translate_CLI) 查看。

---

欢迎贡献和反馈！如果有改进建议，请提交 Issue 或 Pull Request。
