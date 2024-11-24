# Translate_CLI

**Translate** is an easy-to-use command-line translation tool. It supports multiple languages and allows fast and accurate translations directly in the terminal.

## 文档

完整文档提供多种语言版本：

- [中文 README](https://github.com/Serendipity565/Translate_CLI/blob/main/README.md)
- [English README](https://github.com/Serendipity565/Translate_CLI/blob/main/docs/README_en.md)

## Features

* Supports multiple languages
* Fast and reliable translations
* Simple and intuitive command-line interface
* Provides auto-completion scripts for popular shells
* Supports multiple translation APIs

## Installation

1. Download and extract the zip file from the [download link](https://github.com/Serendipity565/Translate_CLI/releases).
2. After extraction, modify the contents of the `user.yaml` configuration file in the `config` folder with your own settings.
3. Add the path to `tran.exe` to your system’s environment variables. For example, add `C:\translate` to the environment variable.

## Usage

Easily translate between languages with the following commands:

```bash
tran [flags]
tran [command]
```

### Available Commands

| Command      | Description                                                |
|--------------|------------------------------------------------------------|
| `completion` | Generate the autocompletion script for the specified shell |
| `en-zh`      | Translate English to Chinese                               |
| `zh-en`      | Translate Chinese to English                               |
| `ja-en`      | Translate Japanese to English                              |  
| `zh-ja`      | Translate Chinese to Japanese                              |  
| `en-ja`      | Translate English to Japanese                              |  
| `ja-zh`      | Translate Japanese to Chinese                              |  
| `help`       | View help information for any command                      |  
| `d`          | Auto-detect language, supports only Chinese and English    |  

### Available Flags

| Flag              | Description                           |
|-------------------|---------------------------------------|
| `-h`, `--help`    | Display help for the `translate` tool |
| `-v`, `--version` | Print the version number and exit     |

## Examples

Translate English to Chinese:
```bash
tran en-zh "Hello, world!"
```

Translate Chinese to English:
```bash
tran zh-en "你好，世界！"
```

View help for a specific command:
```bash
tran [command] --help
tran [command] -h
```

View the current version:
```bash
tran [command] --version
tran [command] -v
```

## Documentation

Complete documentation is available [here](https://github.com/Serendipity565/Translate_CLI/blob/main/README.md).

---

Contributions and feedback are welcome! If you have any suggestions for improvements, please submit an Issue or Pull Request.