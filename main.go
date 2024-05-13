package main

import (
	"eth_approve_parse/approval"
	"eth_approve_parse/db"
	"eth_approve_parse/eth_client"
	"fmt"
)

//usdtContract := "0xdAC17F958D2ee523a2206206994597C13D831ec7"

func main() {

	err := db.InitDB()
	if err != nil {
		return
	}
	err = eth_client.InitEthClient()
	if err != nil {
		panic(err)
	}

	// 监听最新区块，解析approve数据
	go approval.WatchApproval()

	// 根据区块号查找旧区块，解析approve数据
	//go approval.FindApproval()

	contract := "0xdAC17F958D2ee523a2206206994597C13D831ec7"
	owner := "0x81d3D78FEFb33562D807686c0A0e589E5979EacD"
	spender := "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"

	// 获取USDT合约，的账户approve 剩余数量
	allowance, err := approval.GetAllowance(contract, owner, spender)
	if err != nil {
		return
	}
	fmt.Printf("合约 %s -- %s 授权给 %s, 数量：%s \n", contract, owner, spender, allowance)

	select {}

}
