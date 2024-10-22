package router

import (
	"net/http"

	"github.com/felipeazsantos/personal-finance-go/internal/service"
	"github.com/felipeazsantos/personal-finance-go/pkg/handlers"
)

type Router struct {
	server *http.Server
}

func New(addr string) *Router {
	return &Router{
		server: &http.Server{Addr: addr},
	}
}

func (r *Router) StartServer(svce *service.Service) error {
	h := handlers.NewHandlers(svce)

	http.HandleFunc("/financial-transaction", h.GetFinancialTransactionByID)
	return r.server.ListenAndServe()
}
