# Rate Limiter

## Sobre

Rate Limiter é um projeto desenvolvido em Go destinado a limitar o número de solicitações que um cliente pode fazer a uma API em um determinado período de tempo. Isso ajuda o sistema a manter-se estável e evitar sobrecarrega causada por excesso de solicitações.

## Como funciona

O Rate Limiter mantém o controle das solicitações feitas a ele, utilizando uma chave única fornecida por cada cliente. Isso é feito através da função `CheckLimit(key)`. Quando essa função é chamada, ela verifica a contagem atual de solicitações da chave fornecida. Se o limite for excedido, a função retornará false. Caso contrário, retornará true e a contagem de solicitações para essa chave será incrementada.

## Instalação

Para usar a biblioteca Rate Limiter em um projeto Go, você pode obtê-la usando o comando abaixo:  
`import "github.com/songomes/desafiotecnicoratelimiter.git"`  
  
## Configuração

As configurações de execução estão no arquivo `config.env`.  

## Teste

Testes podem ser executados com o comando `go test -v`.