package kernel

import "encoding/json"

func (m *MixinNetwork) GetWithdrawalClaim(tx string) (*Transaction, error) {
	b, err := m.callRPC("getwithdrawalclaim", []any{tx})
	if err != nil || b == nil {
		return nil, err
	}
	var t Transaction
	err = json.Unmarshal(b, &t)
	return &t, err
}
