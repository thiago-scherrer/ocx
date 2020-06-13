# ocx

![Go](https://github.com/thiago-scherrer/ocx/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/thiago-scherrer/ocx)](https://goreportcard.com/report/github.com/thiago-scherrer/ocx)

Outro CLI x para a AWS

## Instalar baixando o binário

Baixe o binário de acordo com o seu sistema operacional, indo em [releases](https://github.com/thiago-scherrer/ocx/releases)  

Linux e macOS

```sh
tar -C /bin/ -zxvf ocx_0.1.0_linux_amd64.tar.gz ocx
```

## Instalar compilando

Usando go get (go >=1.14.4):

```sh
go get github.com/thiago-scherrer/ocx
```

## Como usar

Vá até o [security credentials](https://console.aws.amazon.com/iam/home#security_credential) do seu/sua usuário/a e gere as chaves de acesso:

![key/secret](misc/keys.png)

Com os dados gerados execute o setup, com o comando `ocx setup`:

```sh
$ ocx setup
Enter the AWS_ACCESS_KEY_ID: aaaaa
Enter the AWS_SECRET_ACCESS_KEY: bbbbb
Enter the AWS region: us-east-1

```

O AWS region é onde você está executando a infra instrutura.
Você pode pegar este dado no canto superior de sua conta.

Listando Log Groups:

```sh
$ ocx log group

ola42
```

Lista todos os Stream, ordenados por atualização:

```sh
$ ocx log stream --gname ola42

...
```

Lista o stream mais recente e que possuem *stream* no nome:

```sh
$ ocx log stream --gname ola42 --filter stream

...
```

Recebendo os eventos de logs:

```sh
$ ocx log tail --gname ola42 --sname logstream --lines 10

2020/06/08 01:48:57 Event messages for stream logstream in log group ola42:
teste1
teste2
teste3
teste4
teste5
teste6
teste7
teste8
teste9
teste10

```

Exemplo de comandos de ajuda:

```sh
ocx --help
ocx log --help
ocx log tail --help
```
