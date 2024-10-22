package main

import (
	"log"

	"github.com/felipeazsantos/personal-finance-go/config"
	"github.com/felipeazsantos/personal-finance-go/internal/database"
	"github.com/felipeazsantos/personal-finance-go/internal/repository"
	"github.com/felipeazsantos/personal-finance-go/internal/router"
	"github.com/felipeazsantos/personal-finance-go/internal/service"
)

/*

### Descrição
Ferramenta para usuários gerenciarem suas finanças pessoais, despesas e receitas.

### Funcionalidades
- Registro e categorização de transações financeiras.
- Orçamentos mensais e acompanhamento de gastos.
- Relatórios e visualizações gráficas das finanças.
- Integração com contas bancárias para importação automática de dados.
- Alertas de gastos e metas financeiras.

*/

const (
	serverHost = "localhost"
	serverPort = ":4001"
)

func main() {
	appConfig := config.New()
	db := database.New(appConfig)
	repo := repository.New(db)
	svce := service.New(repo)
	router := router.New(serverHost + serverPort)
	router.SetupHandlers(svce)

	log.Printf("server is listening on http:%s%s ...\n", serverHost, serverPort)

	if err := router.StartServer(); err != nil {
		log.Fatalf("error on start server: %v", err)
	}
}
