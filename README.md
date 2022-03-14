# Log Reader with Golang, MongoDB and Docker

## Requisitos:

- Docker
- Docker-Compose

**Log Parser** desenvolvido em Golang para ler um arquivo de Logs registrados como objetos JSON e armazenar os dados utilizando MongoDB.

Após processar o arquivo o sistema gera 3 relatórios em formato CSV:
- Relatório com informações dos logs organizados por ID do Consumidor.
- Relatório com informações dos logs organizados por ID do Serviço.
- Relatório com informações de Tempo Médio de request , proxy e gateway por serviço.

## Executando o Projeto

Antes de executar o projeto é necessário descompactar o arquivo ``logs.txt`` utilizado para teste, na pasta ``src``:

```bash
unzip logs.zip
```

Na raíz do projeto execute:
```bash
docker-compose build

docker-compose up
```
O sistema irá exibir um log de execução.

Após finalizar a execução, os arquivos dos relatórios estarão disponíveis na pasta ``src`` com os nomes:
- **``ReportByConsumer.csv``**
- **``ReportByService.csv``**
- **``ReportByServiceTime.csv``**