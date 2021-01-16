package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	// this line is used by starport scaffolding # 1
)

const (
	MethodGet = "GET"
)

// RegisterRoutes registers vncarbon-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	r.HandleFunc("/vncarbon/vncarbons/{id}", getVncarbonHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/vncarbon/vncarbons", listVncarbonHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	r.HandleFunc("/vncarbon/vncarbons", createVncarbonHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/vncarbon/vncarbons/{id}", updateVncarbonHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/vncarbon/vncarbons/{id}", deleteVncarbonHandler(clientCtx)).Methods("POST")

}
