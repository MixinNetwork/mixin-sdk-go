package kernel

import (
	"encoding/json"
)

type Asset struct {
	AssetKey string `json:"asset_key"`
	Balance  string `json:"balance"`
	Chain    string `json:"chain"`
	ID       string `json:"id"`
}

func (m *MixinNetwork) GetAsset(id string) (*Asset, error) {
	b, err := m.callRPC("getasset", []any{id})
	if err != nil || b == nil {
		return nil, err
	}
	var a Asset
	err = json.Unmarshal(b, &a)
	return &a, err
}
