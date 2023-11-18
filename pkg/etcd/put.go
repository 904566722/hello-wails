package etcd

import (
	"bytes"
	"encoding/json"
)

func (e *EtcdClient) Put(key, val string) error {
	ctx, cancel := e.etcdOpCtx()
	defer cancel()
	_, err := e.cli.Put(ctx, key, val)
	if err != nil {
		return err
	}
	return nil
}

func (e *EtcdClient) PutCompact(key, val string) error {
	compactJson, err := compact([]byte(val))
	if err != nil {
		return err
	}
	return e.Put(key, compactJson)
}

func compact(s []byte) (string, error) {
	buffer := bytes.Buffer{}
	err := json.Compact(&buffer, s)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}
