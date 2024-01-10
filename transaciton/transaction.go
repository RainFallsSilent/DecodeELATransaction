package transaciton

import (
	"bytes"

	"github.com/elastos/Elastos.ELA/common"
	"github.com/elastos/Elastos.ELA/core/transaction"
	"github.com/elastos/Elastos.ELA/servers"
)

func DecodeRawTransaction(rawTransaction string) (*servers.TransactionInfo, error) {
	txBytes, err := common.HexStringToBytes(rawTransaction)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(txBytes)
	txn, err := transaction.GetTransactionByBytes(r)
	if err != nil {
		return nil, err
	}
	if err := txn.Deserialize(r); err != nil {
		return nil, err
	}

	return servers.GetTransactionInfo(txn), nil
}
