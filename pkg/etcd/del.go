package etcd

func (e *EtcdClient) Del(key string) error {
	ctx, cancel := e.etcdOpCtx()
	defer cancel()
	if _, err := e.cli.Delete(ctx, key); err != nil {
		return err
	}
	return nil
}
