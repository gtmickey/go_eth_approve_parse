package eth_client

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	EthClient *ethclient.Client
)

func InitEthClient() error {
	var err error
	EthClient, err = ethclient.Dial("wss://mainnet.infura.io/ws/v3/zzz")
	if err != nil {
		fmt.Println("error:", err)
		return err
	}
	return nil
}
