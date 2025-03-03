package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"sync"
	"time"
)

const (
	machineIDBits  = 4  // 机器ID占用的位数
	sequenceBits   = 10 // 序列号占用的位数
	machineIDShift = sequenceBits
	timestampShift = machineIDBits + sequenceBits
	maxMachineID   = -1 ^ (-1 << machineIDBits)
	maxSequence    = -1 ^ (-1 << sequenceBits)
)

type OrderIDGenerator struct {
	machineID int64
	sequence  int64
	lastStamp int64
	mu        sync.Mutex
}

// 获取Docker容器的唯一ID
func getContainerID() string {
	containerID, _ := os.Hostname()
	return containerID
}

// 自动获取机器码
func getMachineID() int64 {
	// 获取容器的唯一ID
	containerID := getContainerID()

	// 对容器ID进行哈希处理
	hash := md5.Sum([]byte(containerID))
	hashStr := hex.EncodeToString(hash[:])

	// 取哈希值的前4个字节作为机器ID
	machineID := int64(hashStr[0])<<24 | int64(hashStr[1])<<16 | int64(hashStr[2])<<8 | int64(hashStr[3])
	machineID = machineID & maxMachineID // 确保机器ID在范围内

	return machineID
}

func NewOrderIDGenerator() *OrderIDGenerator {
	machineID := getMachineID()
	return &OrderIDGenerator{
		machineID: machineID,
	}
}

func (g *OrderIDGenerator) NextID() string {
	g.mu.Lock()
	defer g.mu.Unlock()
	var t = time.Now()
	now := t.UnixNano() / 1e6 // 获取当前时间的毫秒数
	if now == g.lastStamp {
		g.sequence = (g.sequence + 1) & maxSequence
		if g.sequence == 0 {
			// 如果序列号超出范围，等待下一毫秒
			for now <= g.lastStamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		g.sequence = 0
	}

	g.lastStamp = now

	// 生成20位订单编号
	id := (now << timestampShift) | (g.machineID << machineIDShift) | g.sequence
	var date = t.Format("200601021504")
	orderId := fmt.Sprintf("420%s%020d", date, id)
	return orderId
}

type OrderIds struct {
	Ids []string
}

func main() {
	var generator = NewOrderIDGenerator() // 自动获取机器码
	var wg errgroup.Group
	var list = &OrderIds{
		Ids: make([]string, 0),
	}
	for i := 0; i < 100; i++ {
		for k := 0; k < 50; k++ {
			id := generator.NextID()
			fmt.Println(id)
			list.Ids = append(list.Ids, id)
		}
	}
	err := wg.Wait()
	if err != nil {
		fmt.Println(err)
	}
	var ids = make([]string, 0)
	//将list 去重
	fmt.Println("over")
	for index, id := range list.Ids {
		if !InArray(id, ids) {
			ids = append(ids, id)
		}
		fmt.Println("已经处理了", index, "个订单")
	}

	return

}

func InArray(id string, ids []string) bool {
	for _, orderId := range ids {
		if orderId == id {
			return true
		}
	}
	return false

}
