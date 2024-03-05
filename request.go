package kernel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type MixinNetwork struct {
	httpClient *http.Client
	node       string
}

func NewMixinNetwork(node string) *MixinNetwork {
	return &MixinNetwork{
		httpClient: &http.Client{Timeout: 30 * time.Second},
		node:       node,
	}
}

func (m *MixinNetwork) Request(method string, params []any) ([]byte, error) {
	return m.callRPC(method, params)
}

func (m *MixinNetwork) callRPC(method string, params []any) ([]byte, error) {
	body, err := json.Marshal(map[string]any{
		"method": method,
		"params": params,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", m.node, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	resp, err := m.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Data  any `json:"data"`
		Error any `json:"error"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, fmt.Errorf("ERROR %s", result.Error)
	}

	return json.Marshal(result.Data)
}