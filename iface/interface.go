package iface

import (
	"context"
	"github.com/coinsummer/comsclient/types"
	"math/big"
)

type IClient interface {
	ChainID() (string, error)
	BlockByHash(ctx context.Context, hash string) (*types.Block, error)
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
	BlockNumber(ctx context.Context) (string, error)
	TransactionByHash(ctx context.Context, hash string) (*types.Transaction, error)
	SubscribeNewHead(ctx context.Context)
	SubscribeFilterLogs(ctx context.Context)

	//BalanceAt(ctx context.Context)
	//TransactionCount()
	//PendingTransactionCount()
	//TransactionByHash(hash string)
	//SendTransaction()
	//CancelTransaction()
	//CheckTransaction()
}
