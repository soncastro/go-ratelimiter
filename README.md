# Rate Limiter

## Sobre
Rate Limiter é um projeto desenvolvido em Go destinado a limitar o número de solicitações que um cliente pode fazer a uma API em um determinado período de tempo. Isso ajuda o sistema a manter-se estável e evitar sobrecarrega causada por excesso de solicitações.

## Como o projeto pode ser executado  
Na raiz do projeto executar o comando `docker-compose up`  
Lembrete: Se houver tentativa de execução com o comando `go run main.go rate_limiter.go` o arquivo config.env não será lido.

## Request para teste manual
`curl -i -X GET http://localhost:8080/ratelimiter`

## Configuração
As configurações do Rate Limiter estão no arquivo `config.env`.  
As variáveis são:  
* BLOCK_TIME_IN_SECONDS : Valor em segundos
* RATE_LIMIT : Quantidade de solicitação
* TOKEN_RATE_LIMIT : Quantidade de solicitação

## Teste
Testes podem ser executados com o comando `go test -v`.

## Como funciona
O Rate Limiter mantém o controle das solicitações feitas a ele, utilizando uma chave única fornecida por cada cliente. Isso é feito através da função `CheckLimit(key)`. Quando essa função é chamada, ela verifica a contagem atual de solicitações da chave fornecida. Se o limite for excedido, a função retornará false. Caso contrário, retornará true e a contagem de solicitações para essa chave será incrementada.

## Instalação
Para usar a biblioteca Rate Limiter em um projeto Go, você pode obtê-la usando o comando abaixo:  
`import "github.com/songomes/desafiotecnicoratelimiter.git"`  
