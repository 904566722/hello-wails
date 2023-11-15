package etcd

func (e *EtcdClient) Put(key, val string) error {
	ctx, cancel := etcdOpCtx()
	defer cancel()
	_, err := e.cli.Put(ctx, key, val)
	if err != nil {
		return err
	}
	return nil
}
