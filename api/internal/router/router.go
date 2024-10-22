package router

import (
	"net/http"

	"github.com/felipeazsantos/personal-finance-go/internal/middlewares"
	"github.com/felipeazsantos/personal-finance-go/internal/service"
	"github.com/felipeazsantos/personal-finance-go/pkg/handlers"
)

type Route struct {
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

type Router struct {
	server *http.Server
	routes []Route
}

func New(addr string) *Router {
	return &Router{
		server: &http.Server{Addr: addr},
	}
}

func (r *Router) SetupHandlers(svce *service.Service) {
	h := handlers.NewHandlers(svce)
	r.registerRoute("GET", "/financial-transaction", h.GetFinancialTransactionByID)
	middlewares.RegisterRoutes(r.routesToMiddlewareRoutes())
}

func (r *Router) StartServer() error {
	return r.server.ListenAndServe()
}

func (r *Router) registerRoute(method, pattern string, handler http.HandlerFunc) {
	r.routes = append(r.routes, Route{Method: method, Pattern: pattern, Handler: handler})
}

func (r *Router) routesToMiddlewareRoutes() []middlewares.MiddlewareRoute {
	mwRoutes := []middlewares.MiddlewareRoute{}

	for _, route := range r.routes {
		mwRoute := middlewares.MiddlewareRoute{
			Method:  route.Method,
			Pattern: route.Pattern,
			Handler: route.Handler,
		}
		mwRoutes = append(mwRoutes, mwRoute)
	}

	return mwRoutes
}
