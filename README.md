# ocx

Outro CLI x

## Como usar

Vá até o [security credentials](https://console.aws.amazon.com/iam/home#security_credential) do seu/sua usuário/a e gere as chaves de acesso:

![key/secret](misc/keys.png)

Com os dados gerados execute o setup, com o comando `ocx setup`:

```sh
ocx setup
Enter the AWS_ACCESS_KEY_ID: aaaaa
Enter the AWS_SECRET_ACCESS_KEY: bbbbb
Enter the AWS region: us-east-1

```

O AWS region é onde você está executando sua infra instrutura. Você pode pegar este dado no canto superior de sua conta.
