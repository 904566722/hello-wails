package etcd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEtcdClient_ListKeyByKeyword(t *testing.T) {
	client, err := NewEtcdClient()
	assert.NoError(t, err)
	list, err := client.ListKeyByKeyword("nic")
	assert.NoError(t, err)
	t.Log(list)
}
