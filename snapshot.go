package kernel

import "encoding/json"

type SnapshotWithTransaction struct {
	Hash         string         `json:"hash"`
	Timestamp    uint64         `json:"timestamp"`
	Topology     uint64         `json:"topology"`
	Transactions []*Transaction `json:"transactions"`
	Version      int            `json:"version"`
}

func (m *MixinNetwork) GetSnapshot(hash string) (*SnapshotWithTransaction, error) {
	b, err := m.callRPC("getsnapshot", []any{hash})
	if err != nil || b == nil {
		return nil, err
	}
	var s SnapshotWithTransaction
	err = json.Unmarshal(b, &s)
	if err != nil {
		return nil, err
	}
	return &s, err
}

func (m *MixinNetwork) ListSnapshotsSince(since, count, sig, tx uint64) ([]*SnapshotWithTransaction, error) {
	body, err := m.callRPC("listsnapshots", []any{since, count, sig, tx})
	if err != nil {
		return nil, err
	}
	var snapshots []*SnapshotWithTransaction
	err = json.Unmarshal(body, &snapshots)
	return snapshots, err
}
