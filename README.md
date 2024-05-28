# Rate Limiter

## Sobre
Rate Limiter é um projeto desenvolvido em Go destinado a limitar o número de solicitações que um cliente pode fazer a uma API em um determinado período de tempo. Isso ajuda o sistema a manter-se estável e evitar sobrecarrega causada por excesso de solicitações.

## Como o projeto pode ser executado  
Na raiz do projeto executar o comando `docker-compose up`  
Lembrete: Se houver tentativa de execução com o comando `go run main.go rate_limiter.go` o arquivo `config.env` não será lido.

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

## Enunciado
Objetivo: Desenvolver um rate limiter em Go que possa ser configurado para limitar o número máximo de requisições por segundo com base em um endereço IP específico ou em um token de acesso.

Descrição: O objetivo deste desafio é criar um rate limiter em Go que possa ser utilizado para controlar o tráfego de requisições para um serviço web. O rate limiter deve ser capaz de limitar o número de requisições com base em dois critérios:

1. Endereço IP: O rate limiter deve restringir o número de requisições recebidas de um único endereço IP dentro de um intervalo de tempo definido.
2. Token de Acesso: O rate limiter deve também poderá limitar as requisições baseadas em um token de acesso único, permitindo diferentes limites de tempo de expiração para diferentes tokens. O Token deve ser informado no header no seguinte formato:
    1. API_KEY: <TOKEN>
3. As configurações de limite do token de acesso devem se sobrepor as do IP. Ex: Se o limite por IP é de 10 req/s e a de um determinado token é de 100 req/s, o rate limiter deve utilizar as informações do token.

Requisitos:

* O rate limiter deve poder trabalhar como um middleware que é injetado ao servidor web
* O rate limiter deve permitir a configuração do número máximo de requisições permitidas por segundo.
* O rate limiter deve ter ter a opção de escolher o tempo de bloqueio do IP ou do Token caso a quantidade de requisições tenha sido excedida.
* As configurações de limite devem ser realizadas via variáveis de ambiente ou em um arquivo “.env” na pasta raiz.
* Deve ser possível configurar o rate limiter tanto para limitação por IP quanto por token de acesso.
* O sistema deve responder adequadamente quando o limite é excedido:
  * Código HTTP: 429
  * Mensagem: you have reached the maximum number of requests or actions allowed within a certain time frame
* Todas as informações de "limiter” devem ser armazenadas e consultadas de um banco de dados Redis. Você pode utilizar docker-compose para subir o Redis.
* Crie uma “strategy” que permita trocar facilmente o Redis por outro mecanismo de persistência.
* A lógica do limiter deve estar separada do middleware.
* 
Exemplos:

1. Limitação por IP: Suponha que o rate limiter esteja configurado para permitir no máximo 5 requisições por segundo por IP. Se o IP 192.168.1.1 enviar 6 requisições em um segundo, a sexta requisição deve ser bloqueada.
2. Limitação por Token: Se um token abc123 tiver um limite configurado de 10 requisições por segundo e enviar 11 requisições nesse intervalo, a décima primeira deve ser bloqueada.
3. Nos dois casos acima, as próximas requisições poderão ser realizadas somente quando o tempo total de expiração ocorrer. Ex: Se o tempo de expiração é de 5 minutos, determinado IP poderá realizar novas requisições somente após os 5 minutos.

Dicas:

* Teste seu rate limiter sob diferentes condições de carga para garantir que ele funcione conforme esperado em situações de alto tráfego.

Entrega:

* O código-fonte completo da implementação.
* Documentação explicando como o rate limiter funciona e como ele pode ser configurado.
* Testes automatizados demonstrando a eficácia e a robustez do rate limiter.
* Utilize docker/docker-compose para que possamos realizar os testes de sua aplicação.
* O servidor web deve responder na porta 8080.