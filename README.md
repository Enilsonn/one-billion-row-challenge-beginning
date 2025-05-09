# Weather Measurements Processor

Este repositório contém um projeto que processa dados meteorológicos a partir de um arquivo CSV, converte-os para um formato de entrada otimizado e executa um programa em Go que analisa e resume as medições.

## Estrutura do Projeto

- `create_measurements.py`: Script em Python responsável por ler o arquivo CSV (`weather_stations.csv`) e gerar um arquivo chamado `measurements.txt` contendo até **N_ROWS** linhas, no formato esperado pelo programa Go.
- Arquivos `.go`: Implementam o sistema de leitura, processamento e formatação dos dados de medição.
  - `main.go`: Ponto de entrada da aplicação.
  - `measurements.go`: Lê o arquivo `measurements.txt`, processa os dados e calcula valores mínimos, máximos e médias por local.
  - `formaterMeasurements.go`: Formata a saída final dos dados processados.
  - `sortMeasurements.go`: Ordena as medições alfabeticamente por local.

## Uso

### 1. Gerar o arquivo de entrada (`measurements.txt`)

Execute o script Python para gerar o arquivo de medições com um número específico de linhas do CSV:

```bash
python3 create_measurements.py N_ROWS
```

Substitua `N_ROWS` pela quantidade desejada de linhas que serão lidas do `weather_stations.csv`.

### 2. Executar o programa principal em Go

Depois de gerar o arquivo `measurements.txt`, execute o código Go com o comando:

```bash
go run .
```

O programa irá:

- Ler o arquivo `measurements.txt`.
- Agrupar as medições por local.
- Calcular o valor mínimo, máximo e a média por local.
- Exibir os resultados no formato:

```text
Location1=min/avg/max, Location2=min/avg/max, ...
Execution time: 10.123ms (exemplo)
```

## Requisitos

- Python 3.x
- Go 1.16+
