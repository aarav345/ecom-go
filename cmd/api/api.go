package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/aarav345/ecom-go/services/cart"
	"github.com/aarav345/ecom-go/services/history"
	"github.com/aarav345/ecom-go/services/order"
	"github.com/aarav345/ecom-go/services/product"
	"github.com/aarav345/ecom-go/services/user"
	"github.com/gorilla/mux"
)

type ApiServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		db:   db,
	}
}

func (s *ApiServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)

	orderStore := order.NewStore(s.db)
	historyStore := history.NewStore(s.db)

	cartHandler := cart.NewHandler(orderStore, productStore, userStore, historyStore)
	cartHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
