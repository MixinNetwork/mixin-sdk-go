package kernel

import (
	"encoding/json"

	"github.com/MixinNetwork/go-number"
	"github.com/MixinNetwork/mixin/common"
	"github.com/MixinNetwork/mixin/crypto"
)

type Input struct {
	Hash    string              `json:"hash"`
	Index   int                 `json:"index"`
	Genesis string              `json:"genesis,omitempty"`
	Deposit *common.DepositData `json:"deposit,omitempty"`
	Mint    *common.MintData    `json:"mint,omitempty"`
}

type Output struct {
	Type       uint8                  `json:"type"`
	Amount     number.Decimal         `json:"amount"`
	Keys       []*crypto.Key          `json:"keys,omitempty"`
	Script     common.Script          `json:"script,omitempty"`
	Mask       crypto.Key             `json:"mask,omitempty"`
	Withdrawal *common.WithdrawalData `json:"withdrawal,omitempty"`
}

type Transaction struct {
	Version    uint8     `json:"version"`
	Asset      string    `json:"asset"`
	Inputs     []*Input  `json:"inputs"`
	Outputs    []*Output `json:"outputs"`
	Extra      string    `json:"extra"`
	Hash       string    `json:"hash"`
	Raw        string    `json:"hex"`
	Snapshot   string    `json:"snapshot"`
	References []string  `json:"references"`
}

func (m *MixinNetwork) SendRawTransaction(raw string) (*Transaction, error) {
	body, err := m.callRPC("sendrawtransaction", []any{raw})
	if err != nil {
		return nil, err
	}
	var tx Transaction
	err = json.Unmarshal(body, &tx)
	return &tx, err
}

func (m *MixinNetwork) GetTransaction(hash string) (*Transaction, error) {
	b, err := m.callRPC("gettransaction", []any{hash})
	if err != nil || b == nil {
		return nil, err
	}
	var tx Transaction
	err = json.Unmarshal(b, &tx)
	if err != nil || tx.Hash == "" {
		return nil, err
	}
	return &tx, err
}
