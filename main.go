package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func callback(master *redis.Client) error {
	it := master.Scan(0, "", 1).Iterator()
	for it.Next() {

		fmt.Println(it.Val())
	}
	return nil
}
func main() {

	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"127.0.0.1:30006", "127.0.0.1:30001", "127.0.0.1:30002", "127.0.0.1:30003", "127.0.0.1:30004", "127.0.0.1:30005"},
	})
	err := client.ForEachMaster(callback)
	if err != nil {
		fmt.Println(err.Error())
	}

}
