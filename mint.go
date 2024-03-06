package kernel

import (
	"encoding/json"
)

type Consensus struct {
	Batch       int64  `json:"aggregator"`
	Pledge      string `json:"payee"`
	Pool        string `json:"signer"`
	Spaces      []int  `json:"spaces"`
	State       string `json:"state"`
	Timestamp   int64  `json:"timestamp"`
	Transaction string `json:"transaction"`
	Works       []int  `json:"works"`
}

type Graph struct {
	Consensus []Consensus `json:"consensus"`
}

type Mint struct {
	Batch  int64  `json:"batch"`
	Pledge string `json:"pledge"`
	Pool   string `json:"pool"`
}

type Info struct {
	Mint  Mint  `json:"mint"`
	Graph Graph `json:"graph"`
}

func (m *MixinNetwork) GetInfo() (*Info, error) {
	b, err := m.callRPC("getinfo", []any{})
	if err != nil || b == nil {
		return nil, err
	}
	var i Info
	err = json.Unmarshal(b, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}