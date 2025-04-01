package service_discover

import (
	"context"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
)

const (
	etcdScheme             = "etcd"
	resolverUpdateInterval = 10 // 解析器更新周期为10秒
	etcdDialTimeout        = 5  // etcd连接超时时间设置为5秒
)

type etcdResolver struct {
	target        resolver.Target
	cc            resolver.ClientConn
	closeCh       chan struct{}
	etcdEndpoints []string
	etcdClient    *clientv3.Client
	keyPrefix     string
}

func (er *etcdResolver) ResolveNow(resolver.ResolveNowOptions) {
	resp, err := er.etcdClient.Get(context.Background(), er.keyPrefix, clientv3.WithPrefix())
	if err != nil {
		log.Printf("Failed to get data from etcd: %v", err)
		return
	}

	addrs := make([]resolver.Address, 0, len(resp.Kvs))

	for _, v := range resp.Kvs {
		addr := resolver.Address{Addr: string(v.Value)}
		addrs = append(addrs, addr)
	}
	er.cc.UpdateState(resolver.State{Addresses: addrs})
	log.Printf("ResolveNow addrs: %v\n", addrs)
}

// Close closes the resolver.
func (er *etcdResolver) Close() {
	er.etcdClient.Close()
	close(er.closeCh)
}

func (er *etcdResolver) start() {
	cli, err := er.NewEtcdClient()
	if err != nil {
		return
	}
	er.etcdClient = cli
	er.keyPrefix = er.getKeyPrefix(er.target.Endpoint())
	er.ResolveNow(resolver.ResolveNowOptions{})
	go er.watch()
}

func (er *etcdResolver) NewEtcdClient() (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   er.etcdEndpoints,
		DialTimeout: etcdDialTimeout * time.Second,
	})
	if err != nil {
		log.Fatalf("Failed to connect to etcd: %v, etcdEndpoints:%v", err, er.etcdEndpoints)
		return nil, err
	}
	return cli, nil
}

func (er *etcdResolver) getKeyPrefix(endpoint string) string {
	return fmt.Sprintf("/%s", endpoint)

}

// watch 监听 etcd 中服务地址的变化
func (er *etcdResolver) watch() {
	watchChan := er.etcdClient.Watch(context.Background(), er.keyPrefix, clientv3.WithPrefix())
	for {
		select {
		case <-er.closeCh:
			return
		case wresp, ok := <-watchChan:
			if !ok {
				break
			}
			if wresp.Err() != nil {
				log.Printf("Error watching etcd: %v", wresp.Err())
				break
			}
			resp, err := er.etcdClient.Get(context.Background(), er.keyPrefix, clientv3.WithPrefix())
			if err != nil {
				log.Printf("Failed to get data from etcd: %v", err)
				continue
			}
			addrs := make([]resolver.Address, 0, len(resp.Kvs))
			for _, kv := range resp.Kvs {
				addrs = append(addrs, resolver.Address{Addr: string(kv.Value)})
			}
			er.cc.UpdateState(resolver.State{
				Addresses: addrs,
			})
			log.Printf("update addrs: %v\n", addrs)
		}
	}
}

type EtcdBuilder struct {
	etcdEndpoints []string
}

func NewEtcdBuilder(etcdEndpoints []string) *EtcdBuilder {
	return &EtcdBuilder{etcdEndpoints: etcdEndpoints}
}

func (eb *EtcdBuilder) Scheme() string {
	fmt.Println("call Scheme")
	return etcdScheme
}

func (eb *EtcdBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	fmt.Println("call Build, target:", target)
	er := &etcdResolver{
		target:        target,
		cc:            cc,
		closeCh:       make(chan struct{}),
		etcdEndpoints: eb.etcdEndpoints,
	}

	er.start()
	return er, nil
}
