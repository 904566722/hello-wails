package etcd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEtcdClient_Get(t *testing.T) {
	client, err := NewEtcdClient()
	if err != nil {
		t.Error(err)
		return
	}
	kvs := []struct {
		k string
		v string
	}{
		{
			k: "key1",
			v: "val1",
		},
	}
	if err := client.Put(kvs[0].k, kvs[0].v); err != nil {
		t.Error(err)
		return
	}
	res, err := client.Get(kvs[0].k)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, kvs[0].v, res.Value)
}
