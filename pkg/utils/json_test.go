package utils

import "testing"

func TestJQ(t *testing.T) {
	jsonCompact := "{\"actions\":[\"InitClusterInfo\",\"AddHost\",\"AggNicCreate\",\"ConfigStorageNetwork\",\"ConfigStorage\"],\"currentStep\":1,\"status\":\"failed\",\"installReport\":{\"0\":{\"no\":0,\"params\":{\"clusterName\":\"local-hci\"},\"status\":\"succeed\",\"installReport\":null},\"1\":{\"no\":1,\"params\":{\"10.16.0.1\":\"CDAC5F0C\",\"10.16.0.2\":\"ACE08313\"},\"status\":\"failed\",\"installReport\":{\"10.16.0.1\":\"status: failed msg: 添加节点任务执行结束，但数据库中没有host记录\",\"10.16.0.2\":\"status: failed msg: 添加节点任务执行结束，但数据库中没有host记录\"}}},\"hostNum\":0}"
	jsonExpand, err := JQ(jsonCompact)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(jsonExpand)
}
