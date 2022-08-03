package comsclient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coinsummer/comsclient/rpc"
	"github.com/coinsummer/comsclient/types"
	"math/big"
	"strings"
)

type Client struct {
	chainName string
	rpc       string
	wss       string
}

func (c *Client) ChainID() (string, error) {
	chainId, version, err := c.getChainInfo()
	if err != nil {
		return "", err
	}
	fmt.Printf("terdermint version: %s", version)
	return chainId, nil
}

func (c *Client) BlockByHash(ctx context.Context, hash string) (*types.Block, error) {
	return c._blockByHash(hash)
}

// https://chihuahua-rpc.polkachu.com/block_by_hash?hash=
func (c *Client) _blockByHash(hash string) (*types.Block, error) {
	c.rpc = fmt.Sprintf("%s/block_by_hash?hash=%s", c.rpc, hash)
	request, err := c.connClient(nil).Request(rpc.GetTy)
	if err != nil {
		return nil, err
	}

	_block := &types.Block{}
	err = json.Unmarshal([]byte(request), _block)
	if err != nil {
		return nil, err
	}
	return _block, err
}

func (c *Client) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	//TODO implement me
	panic("implement me")
}

// https://chihuahua-rpc.polkachu.com/block?height=2142001
func (c *Client) _blockByNumber(number *big.Int) (*types.Block, error) {
	c.rpc = fmt.Sprintf("%s/block?height=%s", c.rpc, number.String())

	request, err := c.connClient(nil).Request(rpc.GetTy)
	if err != nil {
		return nil, err
	}

	_block := &types.Block{}
	err = json.Unmarshal([]byte(request), _block)
	if err != nil {
		return nil, err
	}
	return _block, err
}

func (c *Client) BlockNumber(ctx context.Context) (string, error) {
	return c._blockNumber()
}

func (c *Client) _blockNumber() (string, error) {
	status, err := c.checkStatus()
	if err != nil {
		return "", err
	}
	return status.Result.SyncInfo.EarliestBlockHeight, nil
}

func (c *Client) TransactionByHash(ctx context.Context, hash string) (*types.Transaction, error) {
	return c._transactionByHash(hash)
}

// /tx?hash=_&prove=_
func (c *Client) _transactionByHash(hash string) (*types.Transaction, error) {
	c.rpc = fmt.Sprintf("%s/tx?hash=%s&prove=true", c.rpc, hash)

	request, err := c.connClient(nil).Request(rpc.GetTy)
	if err != nil {
		return nil, err
	}

	txInfo := &types.Transaction{}
	err = json.Unmarshal([]byte(request), txInfo)
	if err != nil {
		return nil, err
	}

	return txInfo, nil
}

func (c *Client) SubscribeNewHead(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) SubscribeFilterLogs(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func initHeader() map[string]string {
	header := make(map[string]string)
	header["content-type"] = "application/json"
	return header
}

func NewClient(chainName, rawurl string) (*Client, error) {
	_urlPrefix := rawurl[0:4]

	r := ""
	wss := ""

	if strings.Contains(_urlPrefix, "http") || strings.Contains(_urlPrefix, "https") {
		//hc = rpc.NewNet(rawurl, initHeader(), nil)
		r = rawurl
	} else if strings.Contains(_urlPrefix, "ws") || strings.Contains(_urlPrefix, "wss") {
		//wsc = rpc.NewWebsocket(rawurl)
		wss = rawurl
	} else {
		return nil, types.ErrRawURI
	}

	if chainName == "" {
		return nil, types.ErrChainNameNull
	}

	c := &Client{
		chainName: chainName,
		rpc:       r,
		wss:       wss,
	}
	return c, nil
}

func (c *Client) connClient(params map[string]interface{}) *rpc.Net {
	return rpc.NewNet(c.rpc, initHeader(), params)
}

// autoCheckChainInfo
// atom -> check node_info
// others -> check status
func (c *Client) getChainInfo() (string, string, error) {
	if isCosmos(c.chainName) {
		// check node_info
		return c.checkNodeInfo()
	}

	// check status
	status, err := c.checkStatus()
	if err != nil {
		return "", "", err
	}
	return status.Result.NodeInfo.Network, status.Result.NodeInfo.Version, nil
}

// parser chain info

// checkNodeInfo get atom
// https://api.cosmos.network/node_info
func (c *Client) checkNodeInfo() (string, string, error) {
	c.rpc = fmt.Sprintf("%s/node_info", c.rpc)

	request, err := c.connClient(nil).Request(rpc.GetTy)
	if err != nil {
		return "", "", err
	}

	info := &types.NodeInfo{}
	err = json.Unmarshal([]byte(request), info)
	if err != nil {
		return "", "", err
	}

	return info.NodeInfo.Network, info.NodeInfo.Version, nil
}

// checkNetInfo get
func (c *Client) checkStatus() (*types.ChainStatus, error) {
	c.rpc = fmt.Sprintf("%s/status?", c.rpc)
	request, err := c.connClient(nil).Request(rpc.GetTy)
	if err != nil {
		return nil, err
	}

	status := &types.ChainStatus{}
	err = json.Unmarshal([]byte(request), status)
	if err != nil {
		return nil, err
	}

	return status, err
}

func isCosmos(chainName string) bool {
	return strings.ToUpper(chainName) == "ATOM"
}
