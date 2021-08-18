# Unico dev test

### Dependências

* make
* Docker
* Docker compose

#### Execução

* Clonar este repositório
* Executar `make up`

Após execução, os seguintes serviços estarão disponíveis para acesso.

API = http://0.0.0.0:7179

Swagger = http://0.0.0.0:28082

PhpMyAdmin = http://0.0.0.0:23307

### O projeto

O projeto foi feito usando alguns conceitos de Domain-Driven Design, como a aplicação de camadas, contexto delimitado (bounded context) e repository pattern.

A camada de domínio tem 100% de cobertura de testes, sendo todos eles unitários, justamente para cobrir todos os casos de sucesso e erros possíveis. O coverage pode ser consultado no comando `make coverage-docker` quando o projeto estiver em execução, essa possiblidade foi por conta da implementação do DIP "Dependency Inversion Principle", representado pela letra D no SOLID.

Não vi necessidade de implementação de testes na camada de infra, pois neste ponto é melhor a implementação de testes de integração em alguns casos e isto estaria muito fora do escopo.

Como a camada de domínio é agnóstica a origem de entrada de dados, foi possível fazer um comando para o listen de api HTTP, mas também um comando para importação das feiras livres em batch, podendo ser implementado outras formas de entrada, como mensagerias ou filas por exemplo.

Um outro ponto interessante é ver que as entidades do domínio não estão presos a modelagem de banco de dados, sendo que o schema implementado foi proposital como forma de mostrar esta separação, o que permite melhorias, correções ou novas implementações, tanto do lado do domínio, quanto do lado da infra, nesse caso, a estrutura de dados, como tabelas específicas para distritos, sub-prefeituras e vínculos entre as tabelas, por exemplo.

Já a camada de infra contém todo o código necessário para a sustentação da camada de domínio, como o código que de fato faz a interação com o banco de dados, controllers da api e workers para importação. Alguns design patterns foram utilizados para tornar mais dinâmica a operação, como abstract factory para criação das entidades de domínio a partir dos dados de request ou o contrário, criação das entidades de response a partir das entidades de domínio.

Os logs em arquivo foi implementado conforme demanda e pode ser consultado em `make logs-arquivo` durante a execução do projeto, porém, o aconselhável é toda aplicação em container direcione seus logs para o stdout ou stderr, tendo um único ponto de log e tornando mais fácil a implementação de sidecars ou agentes de ferramentas monitoramento ou observabilidade. 

Os logs não foram aplicados a nível de domínio pois não vi necessidade para tal, porém, caso seja necessário, pode ser seguido a mesma lógica do repository pattern, implementação de interface a nível de domínio e implementação concreta a nível de infra e inclusão no domínio por meio de injeção de dependência. 

### Limitações

Como não era o objetivo do projeto, as validações de dados nas requests da API, como o cadastro ou alteração de feira livre foi feito de forma mais simples, sendo apenas validado de um dado está sendo informado ou não.

Esse projeto foi feito em um MacBook Pro M1, sendo uma arquitetura arm, e por conta de ser diferente da tradicional arquitetura amd64 ou x86_64, ainda não existe 100% de compatibilidade, ocorrendo alguns erros durante o desenvolvimento da parte de devops, como o exemplo abaixo relacionado ao qemu:

```bash
docker exec -it unico-dev-test-2k21_api.unico.dev_1 make coverage-html
go test -v ./domain/... -coverprofile coverage.out
# internal/syscall/unix
qemu: uncaught target signal 11 (Segmentation fault) - core dumped
?   	github.com/joubertredrat-tests/unico-dev-test-2k21/domain/fair/entity	[no test files]
```

Esse projeto já possui o arquivo `DEINFO_AB_FEIRASLIVRES_2014.csv`, pois foi necessário fazer uma correção na última linha por estar faltando uma vírgula para manter a consistência em relação a quantidade de colunas relacionados as feiras livres. 