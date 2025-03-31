package service_discover

import (
	"context"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func RegisterService(endpoints []string, serviceName, serviceAddr string, ttl int64) {
	log.Printf("RegisterService endpoints: %v", endpoints)
	// 创建 etcd 客户端
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: etcdDialTimeout * time.Second,
	})
	if err != nil {
		log.Fatalf("Failed to connect to etcd: %v", err)
		return
	}

	// 创建租约
	resp, err := cli.Grant(context.TODO(), ttl)
	if err != nil {
		log.Fatalf("Failed to create lease: %v", err)
		return
	}
	leaseID := resp.ID

	// 注册服务
	key := "/" + serviceName + "/" + serviceAddr
	_, err = cli.Put(context.TODO(), key, serviceAddr, clientv3.WithLease(leaseID))
	if err != nil {
		log.Fatalf("Failed to register service: %v", err)
		return
	}
	log.Printf("Service %s registered with address %s", serviceName, serviceAddr)

	// 保持租约
	keepAliveChan, err := cli.KeepAlive(context.TODO(), leaseID)
	if err != nil {
		log.Fatalf("Failed to keep lease alive: %v", err)
		return
	}
	go func() {
		for {
			select {
			case _, ok := <-keepAliveChan:
				if !ok {
					log.Println("Lease keep alive channel closed")
					return
				}
			}
		}
	}()
}
