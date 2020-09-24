/*
 * @Descripttion:
 * @version:
 * @Author: WangShuaibing
 * @Date: 2020-09-23 17:49:16
 * @LastEditors: WangShuaibing
 * @LastEditTime: 2020-09-23 18:07:09
 */
package consistent_hash

import (
	"fmt"
	"log"
	"testing"

	consistent "stathat.com/c/consistent"
)

func TestNewConsistentHashNodesRing(t *testing.T) {
	// JudgeNodeRing := NewConsistentHashNodesRing(500, []string{"node1", "node2", "node3"})
	// fmt.Printf("%+v", JudgeNodeRing)

	c := consistent.New()
	c.Add("node1")
	c.Add("node2")
	users := []string{"key1", "key2", "key3", "key4", "key5", "key6", "key7"}
	for _, u := range users {
		server, err := c.Get(u)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s => %s\n", u, server)
	}
}
