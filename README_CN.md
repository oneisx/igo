# igo

[English](README.md) | Chinese

欢迎来到Igo，Igo是一个好帮手，它能生成UUID，美化JSON，转换时间戳，加解密等等。
生成的内容不需要手动复制，Igo会为我们把它放到剪切板中。
更多功能，请查看帮助文档。

顺便提一下，通过 [Igo Agent](https://github.com/oneisx/igoagent) 使用Igo会更方便哦。

```shell
Usage:
igo [flags]
igo [command]

Available Commands:
clear       clear screen
codec       A collection of encryption and decryption tools
date        transfer timestamp to date
help        Help about any command
json        beautify json
memo        memo [WIP]
uuid        generate uuid

Flags:
--config string   config file (default is $HOME/.igo.yaml)
-h, --help            help for igo
-i, --interactive     interactively execute commands
-q, --quit            quit interactive mode
-v, --version         print the version of igo

Use "igo [command] --help" for more information about a command.
```

## 1. Enter interactive mode
Igo提供两种使用模式，交互式和非交互式。在交互模式中，你不需要输入根命令igo。

`提示: 将igo放到系统变量$PATH中，会获得更好的体验`
```shell
# 在终端输入
igo -i

# 结果
igo>
```

## 2. UUID
你可以像这样生成UUID：
```shell
# 非交互式
igo uuid
14ef893c-62d5-4eaa-83a7-ddb0a87693e7

igo uuid -n 3
b1f5298d-40d5-4ba6-90f2-0aa0e40b4393
d5a75d77-c454-4163-90ea-98336b95c8f3
3d2c87b3-988b-475f-b664-b45c12e905aa

# 交互式
igo> uuid
14ef893c-62d5-4eaa-83a7-ddb0a87693e7

igo> uuid -n 3
b1f5298d-40d5-4ba6-90f2-0aa0e40b4393
d5a75d77-c454-4163-90ea-98336b95c8f3
3d2c87b3-988b-475f-b664-b45c12e905aa
```
命令执行完后，剪切板中就有相应的内容了。

## 3. JSON
JSON被广泛使用，我们经常需要在日常工作中美化JSON，以便我们能够清楚地看到相关信息，这里还提供了压缩JSON的功能
```shell
# 非交互式:
## 美化
igo json
igo>json:pretty>

## 压缩
igo json -u
igo>json:ugly>

# 交互式:
## 美化
igo>json
igo>json:pretty>

## 压缩
igo>json -u
igo>json:ugly>
```
在此之后，您可以输入json字符串，并在最后带上分号(;)，按下回车键，美好的事情就会发生。
```shell
# 例子
igo>json:pretty>{"qenoap":-844276330,"iokgphia":true,"fhabfiw":1681064845.6926622,"zwdskzh":"xct","xgvzsgbfo":-1079466053.3847966};
{
  "qenoap": -844276330,
  "iokgphia": true,
  "fhabfiw": 1681064845.6926622,
  "zwdskzh": "xct",
  "xgvzsgbfo": -1079466053.3847966
}

igo>json:ugly>{
"qenoap": -844276330,
"iokgphia": true,
"fhabfiw": 1681064845.6926622,
"zwdskzh": "xct",
"xgvzsgbfo": -1079466053.3847966
};
{"qenoap":-844276330,"iokgphia":true,"fhabfiw":1681064845.6926622,"zwdskzh":"xct","xgvzsgbfo":-1079466053.3847966}
```

## 4. DATE
我们无法直接看到时间戳表示的日期，date命令用于转换时间戳，它支持秒和毫秒时间戳。
```shell
# 非交互式:
igo date 1639238044
2021-12-11 23:54:04 +0800 CST
igo date 1639188919040
2021-12-11 10:15:19.04 +0800 CST

# 交互式:
igo>date 1639238044
2021-12-11 23:54:04 +0800 CST
igo>date 1639188919040
2021-12-11 10:15:19.04 +0800 CST
```

## 5. CODEC
加密和解密工具的集合，包括MD5/HmacMD5、Base64、SHA1/HmacSHA1、SHA256/HmacSHA256、SHA512/HmacSHA512、AES等。
```shell
Usage:
igo codec [command]

Available Commands:
aes         Generate aes ciphertext, decrypt aes ciphertext
base64      Generate base64 ciphertext, decrypt base64 ciphertext
md5         Generate MD5/HmacMD5 ciphertext
sha1        Generate SHA1/HmacSHA1 ciphertext
sha256      Generate SHA256/HmacSHA256 ciphertext
sha512      Generate SHA512/HmacSHA512 ciphertext

Flags:
-h, --help   help for codec

Global Flags:
--config string   config file (default is $HOME/.igo.yaml)

Use "igo codec [command] --help" for more information about a command.
```
### 5.1 AES
生成aes密文，解密aes密文，使用flag（-d/--decrypt）进行解密。
需要注意的是，键的长度应为（16,24,32），不建议使用默认值。
```shell
# 非交互式:
igo codec aes oneisx
igo codec aes uZERhkcVewZ7S1j1co+QSkKdvf/52DqkDXgAcJktido= -d
igo codec aes oneisx -k 52DqkDXgAcJktido (Recommend custom key)
igo codec aes OYA/OY1bj6J1wRywYYCIwMC9oW8RqoByngxsBUlGhuw= -dk 52DqkDXgAcJktido (Recommend custom key)

# 交互式:
igo>codec aes oneisx
igo>codec aes -d uZERhkcVewZ7S1j1co+QSkKdvf/52DqkDXgAcJktido= -d
igo>codec aes oneisx -k 52DqkDXgAcJktido (Recommend custom key)
igo>codec aes OYA/OY1bj6J1wRywYYCIwMC9oW8RqoByngxsBUlGhuw= -dk 52DqkDXgAcJktido (Recommend custom key)
```
### 5.2 BASE64
生成base64密文，解密base64密文，使用flag（-d/--decrypt）进行解密
```shell
# 非交互式:
igo codec base64 oneisx
igo codec base64 -d b25laXN4

# 交互式:
igo>codec base64 oneisx
igo>codec base64 -d b25laXN4
```
### 5.3 MD5
生成MD5/HmacMD5密文，通过flag（-k/--key）生成HmacMD5密文
```shell
# 非交互式:
## MD5
igo codec md5 oneisx
## HmacMD5
igo codec md5 oneisx -k thanks
igo codec md5 oneisx --key thanks

# 交互式:
## MD5
igo>codec md5 oneisx
## HmacMD5
igo>codec md5 oneisx -k thanks
igo>codec md5 oneisx --key thanks
```
### 5.4 SHA1
生成SHA1/HmacSHA1密文，通过flag（-k/--key）生成HmacSHA1密文
```shell
# 非交互式:
## SHA1
igo codec sha1 oneisx
## HmacSHA1
igo codec sha1 oneisx -k thanks
igo codec sha1 oneisx --key thanks

# 交互式:
## SHA1
igo>codec sha1 oneisx
## HmacSHA1
igo>codec sha1 oneisx -k thanks
igo>codec sha1 oneisx --key thanks
```
### 5.5 SHA256
生成SHA256/HmacSHA256密文，通过flag（-k/--key）生成HmacSHA256密文
```shell
# 非交互式:
## SHA256
igo codec sha256 oneisx
## HmacSHA256
igo codec sha256 oneisx -k thanks
igo codec sha256 oneisx --key thanks

# 交互式:
## SHA256
igo>codec sha256 oneisx
## HmacSHA256
igo>codec sha256 oneisx -k thanks
igo>codec sha256 oneisx --key thanks
```
### 5.6 SHA512
生成SHA512/HmacSHA512密文，通过flag（-k/--key）生成HmacSHA512密文
```shell
# 非交互式:
## SHA512
igo codec sha512 oneisx
## HmacSHA512
igo codec sha512 oneisx -k thanks
igo codec sha512 oneisx --key thanks

# 交互式:
## SHA512
igo>codec sha512 oneisx
## HmacSHA512
igo>codec sha512 oneisx -k thanks
igo>codec sha512 oneisx --key thanks
```