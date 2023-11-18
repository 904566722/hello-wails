package etcd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEtcdClient_List(t *testing.T) {
	client, err := NewEtcdClient()
	assert.NoError(t, err)
	list, err := client.ListByPrefix("hci/phyNetworkPort")
	assert.NoError(t, err)
	t.Log(list)

}
