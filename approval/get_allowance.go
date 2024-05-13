package approval

import (
	"eth_approve_parse/eth_client"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// GetAllowance
// 获取对应合约的授权可使用额度，为0表示无授权或已经撤销授权, 返回剩余授权数量
func GetAllowance(contract string, owner string, spender string) (*big.Int, error) {

	coin, err := eth_client.NewCoin(common.HexToAddress(contract), eth_client.EthClient)

	// 测试数据
	// owner
	// 0x81d3D78FEFb33562D807686c0A0e589E5979EacD

	// spender
	// 0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D
	// 0x8A42d311D282Bfcaa5133b2DE0a8bCDBECea3073

	allowance, err := coin.Allowance(nil, common.HexToAddress(owner), common.HexToAddress(spender))
	if err != nil {
		return nil, err
	}
	return allowance, nil
}
