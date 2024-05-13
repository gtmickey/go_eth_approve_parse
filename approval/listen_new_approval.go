package approval

import (
	"context"
	"encoding/hex"
	"eth_approve_parse/consts"
	"eth_approve_parse/db"
	"eth_approve_parse/eth_client"
	"eth_approve_parse/model"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
)

// WatchApproval
// 参考approval数据
// https://etherscan.io/tx/0x15d11edcaff5a1ee7c8a6facd8d8438ade3234ddf9935151c519b01248ed3f6b#eventlog
// 监听最新块数据
func WatchApproval() {

	logs := make(chan types.Log)

	filterLogs, err := eth_client.EthClient.SubscribeFilterLogs(context.Background(), ethereum.FilterQuery{}, logs)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	//var blockHeight uint64

	for {
		select {
		case err := <-filterLogs.Err():
			fmt.Println("error:", err)
			return
		case log := <-logs:
			//blockHeight = vLog.BlockNumber

			if strings.ToLower(consts.ApprovalTopics) == strings.ToLower(log.Topics[0].Hex()) {
				hexAmount := hex.EncodeToString(log.Data)
				//fmt.Println("blockHeight: ", blockHeight)
				amount := new(big.Int)
				amount.SetString(hexAmount, 16)
				//fmt.Println("授权USDT")
				//if amount.Int64() == 0 {
				//	fmt.Println("授权数量 : 撤销授权")
				//} else {
				//	fmt.Println("授权数量 : ", amount)
				//}
				//
				//fmt.Println("授权人 ", common.HexToAddress(vLog.Topics[1].Hex()).Hex())
				//fmt.Println("被授权人 ", common.HexToAddress(vLog.Topics[2].Hex()).Hex())
				//fmt.Println("交易hash ", vLog.TxHash)

				err := db.UpdateApproval(&model.Approval{Contract: log.Address.Hex(), Owner: common.HexToAddress(log.Topics[1].Hex()).Hex(), Spender: common.HexToAddress(log.Topics[2].Hex()).Hex(), Value: amount.String(), TxHash: log.TxHash.Hex(), BlockNumber: log.BlockNumber, LogIndex: log.Index})
				if err != nil {
					fmt.Println("UpdateApproval error ", err)
					return
				}
				fmt.Println("监听到新的approve，交易hash--------", log.TxHash)
			}
		}
	}
}
