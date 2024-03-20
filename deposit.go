package kernel

import "encoding/json"

func (m *MixinNetwork) GetDepositTransaction(chain, hash string, index int64) (*Transaction, error) {
	b, err := m.callRPC("getdeposittransaction", []any{chain, hash, index})
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