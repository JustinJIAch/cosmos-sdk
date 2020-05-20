package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

const (
	RestChannelID = "channel-id"
	RestPortID    = "port-id"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	registerQueryRoutes(cliCtx, r)
	registerTxRoutes(cliCtx, r)
}

// TransferTxReq defines the properties of a transfer tx request's body.
type TransferTxReq struct {
	BaseReq          rest.BaseReq `json:"base_req" yaml:"base_req"`
	Amount           sdk.Coins    `json:"amount" yaml:"amount"`
	Receiver         string       `json:"receiver" yaml:"receiver"`
	TimeoutHeight    uint64       `json:"timeout_height" yaml:"timeout_height"`
	TimeoutTimestamp uint64       `json:"timeout_timestamp" yaml:"timeout_timestamp"`
}
