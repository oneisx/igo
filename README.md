# igo
Welcome to igo, igo is a good helper, can generate UUID, beautify JSON, convert timestamp, etc.
No need to copy manually, the generated content will be placed in the clipboard.
For more functions, please see the help.

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
Igo supports two modes, interactive and non-interactive. In the first mode, you do not need to enter the root command igo.
```shell
# input
igo -i

# result
igo>
```

## 2. UUID
You can generate uuid string like this:
```shell
# Non-interactive
igo uuid
14ef893c-62d5-4eaa-83a7-ddb0a87693e7

igo uuid -n 3
b1f5298d-40d5-4ba6-90f2-0aa0e40b4393
d5a75d77-c454-4163-90ea-98336b95c8f3
3d2c87b3-988b-475f-b664-b45c12e905aa

# Interactive
igo> uuid
14ef893c-62d5-4eaa-83a7-ddb0a87693e7

igo> uuid -n 3
b1f5298d-40d5-4ba6-90f2-0aa0e40b4393
d5a75d77-c454-4163-90ea-98336b95c8f3
3d2c87b3-988b-475f-b664-b45c12e905aa
```
After this, the content already in the clipboard.

## 3. JSON
JSON is widely used. We often need to beautify JSON in our daily work so that we can clearly see the relevant information. Here also provides the ability to compress JSON
```shell
# Non-interactive:
## beautify
igo json
igo>json:pretty>

## compress
igo json -u
igo>json:ugly>

# Interactive:
## beautify
igo>json
igo>json:pretty>

## compress
igo>json -u
igo>json:ugly>
```
After this, you can input your json string, and append semicolon(;) in the end, press the Enter, wonderful thing will happen. 
```shell
# Example
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
We can't directly see the date represented by the timestamp. The date command is used to convert the timestamp. It supports second and millisecond timestamps
```shell
# Non-interactive:
igo date 1639238044
2021-12-11 23:54:04 +0800 CST
igo date 1639188919040
2021-12-11 10:15:19.04 +0800 CST

# Interactive:
igo>date 1639238044
2021-12-11 23:54:04 +0800 CST
igo>date 1639188919040
2021-12-11 10:15:19.04 +0800 CST
```

## 5. CODEC
A collection of encryption and decryption tools,
including MD5/HmacMD5, Base64, SHA1/HmacSHA1, SHA256/HmacSHA256,
SHA512/HmacSHA512, AES, etc.
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
Generate aes ciphertext, decrypt aes ciphertext, use flag (-d/--decrypt) to decrypt.
It should be noted that, length of the key should be in (16,24,32), the default value is not recommended.
```shell
# Non-interactive:
igo codec aes oneisx
igo codec aes uZERhkcVewZ7S1j1co+QSkKdvf/52DqkDXgAcJktido= -d
igo codec aes oneisx -k 52DqkDXgAcJktido (Recommend custom key)
igo codec aes OYA/OY1bj6J1wRywYYCIwMC9oW8RqoByngxsBUlGhuw= -dk 52DqkDXgAcJktido (Recommend custom key)

# Interactive:
igo>codec aes oneisx
igo>codec aes -d uZERhkcVewZ7S1j1co+QSkKdvf/52DqkDXgAcJktido= -d
igo>codec aes oneisx -k 52DqkDXgAcJktido (Recommend custom key)
igo>codec aes OYA/OY1bj6J1wRywYYCIwMC9oW8RqoByngxsBUlGhuw= -dk 52DqkDXgAcJktido (Recommend custom key)
```
### 5.2 BASE64
Generate base64 ciphertext, decrypt base64 ciphertext, use flag (-d/--decrypt) to decrypt
```shell
# Non-interactive:
igo codec base64 oneisx
igo codec base64 -d b25laXN4

# Interactive:
igo>codec base64 oneisx
igo>codec base64 -d b25laXN4
```
### 5.3 MD5
Generate MD5/HmacMD5 ciphertext, Generate HmacMD5 ciphertext with flag(-k/--key)
```shell
# Non-interactive:
## MD5
igo codec md5 oneisx
## HmacMD5
igo codec md5 oneisx -k thanks
igo codec md5 oneisx --key thanks

# Interactive:
## MD5
igo>codec md5 oneisx
## HmacMD5
igo>codec md5 oneisx -k thanks
igo>codec md5 oneisx --key thanks
```
### 5.4 SHA1
Generate SHA1/HmacSHA1 ciphertext, Generate HmacSHA1 ciphertext with flag(-k/--key)
```shell
# Non-interactive:
## SHA1
igo codec sha1 oneisx
## HmacSHA1
igo codec sha1 oneisx -k thanks
igo codec sha1 oneisx --key thanks

# Interactive:
## SHA1
igo>codec sha1 oneisx
## HmacSHA1
igo>codec sha1 oneisx -k thanks
igo>codec sha1 oneisx --key thanks
```
### 5.5 SHA256
Generate SHA256/HmacSHA256 ciphertext, Generate HmacSHA256 ciphertext with flag(-k/--key)
```shell
# Non-interactive:
## SHA256
igo codec sha256 oneisx
## HmacSHA256
igo codec sha256 oneisx -k thanks
igo codec sha256 oneisx --key thanks

# Interactive:
## SHA256
igo>codec sha256 oneisx
## HmacSHA256
igo>codec sha256 oneisx -k thanks
igo>codec sha256 oneisx --key thanks
```
### 5.6 SHA512
Generate SHA512/HmacSHA512 ciphertext, Generate HmacSHA512 ciphertext with flag(-k/--key)
```shell
# Non-interactive:
## SHA512
igo codec sha512 oneisx
## HmacSHA512
igo codec sha512 oneisx -k thanks
igo codec sha512 oneisx --key thanks

# Interactive:
## SHA512
igo>codec sha512 oneisx
## HmacSHA512
igo>codec sha512 oneisx -k thanks
igo>codec sha512 oneisx --key thanks
```