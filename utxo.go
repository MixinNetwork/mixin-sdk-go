package kernel

import (
	"github.com/MixinNetwork/go-number"
	"github.com/MixinNetwork/mixin/common"
	"github.com/MixinNetwork/mixin/crypto"
)

type UTXO struct {
	Asset     string
	Hash      string
	Index     int64
	Amount    number.Decimal
	Threshold int64
	Receivers []string
	Keys      []*crypto.Key
	Mask      crypto.Key

	Extra string

	SendersHash      string
	Senders          []string
	SendersThreshold int64

	Deposit *common.DepositData
}

func (utxo UTXO) KeysSlice() []string {
	keys := make([]string, len(utxo.Keys))
	for i, k := range utxo.Keys {
		keys[i] = k.String()
	}
	return keys
}

func (tx *Transaction) Deposit() *common.DepositData {
	return tx.Inputs[0].Deposit
}

func (tx *Transaction) Withdrawal() *common.WithdrawalData {
	return tx.Outputs[0].Withdrawal
}
