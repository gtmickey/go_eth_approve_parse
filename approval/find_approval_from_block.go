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
	"math/big"
	"strings"
)

// https://etherscan.io/tx/0x14aade514c866f31a2f43c4f4c01aad24985b1233aef0e34a02bddc008f03635#eventlog
// 在 19859649 区块中， 这一条是特殊交易，本次交易中， approve方法被调用两次，先 approve 一定数量后，使用后 revoke

func FindApproval() {

	filterLogs, err := eth_client.EthClient.FilterLogs(context.Background(), ethereum.FilterQuery{

		FromBlock: big.NewInt(19859649),
		ToBlock:   big.NewInt(19859650),
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	for _, log := range filterLogs {
		if strings.ToLower(consts.ApprovalTopics) == strings.ToLower(log.Topics[0].Hex()) {
			hexAmount := hex.EncodeToString(log.Data)
			//fmt.Println("blockHeight: ", blockHeight)
			amount := new(big.Int)
			amount.SetString(hexAmount, 16)

			err := db.UpdateApproval(&model.Approval{Contract: log.Address.Hex(), Owner: common.HexToAddress(log.Topics[1].Hex()).Hex(), Spender: common.HexToAddress(log.Topics[2].Hex()).Hex(), Value: amount.String(), TxHash: log.TxHash.Hex(), BlockNumber: log.BlockNumber, LogIndex: log.Index})
			if err != nil {
				fmt.Println("UpdateApproval error ", err)
				return
			}

			fmt.Println("旧区块中找到approve, 交易hash--------", log.TxHash)
		}

	}

}
