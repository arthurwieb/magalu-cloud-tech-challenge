
# Magalu Cloud - Desafio técnico - Sistema de Bilhetagem - Ingestor
# Parte 1: https://docs.google.com/document/d/e/2PACX-1vQGCw6WuWHVMuBiK-ZKUmalws00GGwXZdynklOj4xwexQRznN6GNlGkt0qcmR9aWYB6JwGmt0DqJZwZ/pub <br>
Sistema de monitoramento de consumo de recursos em nuvem que coleta e agrega pulsos de utilização (armazenamento e rede) para fins de cobrança.

# Instalação
```bash
Instale a versão 1.24.1 ou superior do Golang - https://go.dev/doc/install
```

Após clonar o projeto, digite o comando na raiz do projeto
```bash
go run main.go
```

Você deve obter um resultado dos pulsos gerados pelo programa que vão estar agregados por cliente e produto com uma lista de consumo diário
```
Ex: 
[
 {
  "tenant": "tenant_5",
  "product_sku": "storage_1gb",
  "usage_unit": "GB x seg",
  "daily_usage": {
   "1": 5929.662,
   "10": 5936.264,
   "11": 517.2528,
   "12": 4382.1704,
   "13": 8388.8,
   "14": 2128.4072,
   "16": 7054.875,
   ...
  }
 },
```
Como esse ingestor funciona?
O sistema simula o monitoramento de recursos em nuvem através de pulsos de consumo, geramos 20 pulsos por dia durante 30 dias. Estrutura do pulso:
```
type Pulse struct {
    Tenant     string    // "tenant_"
    ProductSKU string    // "storage_1gb ou network_egress"
    UsedAmount float64   // 60 (GB × seg) ou mb transferido
    UsageUnit  string    // "GB × seg" ou "MB"
}
```
Enquanto nosso loop executa, nosso Aggregator estará lendo e acumulando os dados conforme cliente(tenant) e serviço (product_sku).<br>
Após terminar o loop, podemos exibir o que foi agregado em um "relatório mensal".


Algumas anotações que fiz durante a construção do projeto, não estão muito estruturadas pois era meu "dump de ideias" <br> 
https://www.notion.so/Desafio-1c9e6c8ece2b808d9dcffd09199ceb22

Há também testes unitários na geração de pulso, para executar o teste: 
```
cd pulsegenerator/
go test
```
