package types

import "fmt"

var (
	ErrRawURI        = fmt.Errorf("raw url error, please check it")
	ErrChainNameNull = fmt.Errorf("chain name is null")
)
