package transaciton

import (
	"bytes"
	"encoding/hex"
	"github.com/elastos/Elastos.ELA/common"
	"github.com/elastos/Elastos.ELA/core/contract"
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

func PublicKeyToAddress(pubKeyHex string) (string, error) {
	pubKey, err := common.HexStringToBytes(pubKeyHex)
	if err != nil {
		return "", err
	}
	programHash, err := contract.PublicKeyToStandardProgramHash(pubKey)
	if err != nil {
		return "", err
	}
	addr, err := programHash.ToAddress()
	if err != nil {
		return "", err
	}

	return addr, nil
}

func CodeToAddress(codeStr string) (string, error) {
	code, err := hex.DecodeString(codeStr)
	if err != nil {
		return "", err
	}
	codeContract, err := contract.CreateStandardContractByCode(code)
	if err != nil {
		return "", err
	}
	codeProgramHash := codeContract.ToProgramHash()
	codeAddr, err := codeProgramHash.ToAddress()
	if err != nil {
		return "", err
	}

	return codeAddr, nil
}
