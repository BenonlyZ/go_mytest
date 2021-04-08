package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var sentinelClient *redis.SentinelClient
var masterClient *redis.Client
var redisCluster *redis.ClusterClient

func main() {
	str := "10.0.0.1:26389"
	sen := RedisSentinel(str)
	GetMasterClient(sen, "mastername")
	fmt.Println("--------------------------------------------")
	sc := "10.0.0.1:6389"
	ps := "redis123"
	cl := RedisCluster(sc, ps)
	GetRedisClusterStats(cl)
	GetRedisClusterInfo(cl)
}

func RedisSentinel(st string) *redis.SentinelClient {
	/* rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "mymaster",
		SentinelAddrs: []string{"127.0.0.1:26379","127.0.0.1:26480","127.0.0.1:26481"},
	}) */
	if sentinelClient != nil {
		return sentinelClient
	}
	/**
	连接26379端口的sentinel；多个sentinel应循环获取可用的sentinel；
	这里简单实现
	*/
	sentinelClient = redis.NewSentinelClient(&redis.Options{
		Network: "tcp",
		Addr:    st,
	})

	return sentinelClient
}

func RedisCluster(st string, ps string) *redis.ClusterClient {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{st},
		Password: ps,
	})

	return rdb
}

/**
哨兵获取主节点信息
*/
func GetMasterClient(sentinelClient *redis.SentinelClient, name string) *redis.Client {
	/*if masterClient !=nil {
		return masterClient
	}*/

	masterInfo, _ := sentinelClient.GetMasterAddrByName(name).Result()

	fmt.Println("masterInfo", masterInfo)

	masterClient = redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    masterInfo[0] + ":" + masterInfo[1],
	})
	fmt.Println("master", masterInfo[0]+masterInfo[1])
	return masterClient
}

//返回集群状态
func GetRedisClusterStats(ClusterClient *redis.ClusterClient) {
	clusterStats, err := ClusterClient.ClusterInfo().Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("clusterStats", clusterStats)
}

//返回集群信息
func GetRedisClusterInfo(ClusterClient *redis.ClusterClient) {
	clusterInfo, err := ClusterClient.ClusterNodes().Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("clusterInfo", clusterInfo)
}
