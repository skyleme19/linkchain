package bootnode

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"

	"github.com/lianxiangcloud/linkchain/libs/log"
)

var (
	bootNodeLocker   sync.RWMutex
	MainnetBootnodes = []string{
		"https://39.97.128.184:8087",
		"https://39.97.197.181:8087",
		"https://120.55.156.239:8087",
		"https://47.110.211.42:8087",
		"https://47.91.221.28:8087",
		"https://161.117.157.31:8087",
	}
	index = rand.Intn(len(MainnetBootnodes))
)

//UpdateBootNode update MainnetBootnodes from bootnodeAddrs,bootnodeAddrs's format are like https://ip1:port1,https://ip2:port2
func UpdateBootNode(bootnodeAddrs string, logger log.Logger) {
	var bootNodes []string
	endpoints := strings.Split(bootnodeAddrs, ",")
	for i := 0; i < len(endpoints); i++ {
		var addr string
		netinfo := strings.Split(endpoints[i], ":")
		if len(netinfo) == 2 { //maybe is ip:port, not https://ip1:port1
			addr = fmt.Sprintf("https://%s", endpoints[i])
		} else {
			addr = endpoints[i]
		}
		bootNodes = append(bootNodes, addr)
	}
	if len(bootNodes) > 0 {
		bootNodeLocker.Lock()
		index = rand.Intn(len(bootNodes))
		MainnetBootnodes = bootNodes
		bootNodeLocker.Unlock()
	}
	logger.Debug("UpdateBootNode", "index", index, "len(endpoints)", len(endpoints))
}

func GetBootNodesNum() int {
	bootNodeLocker.RLock()
	num := len(MainnetBootnodes)
	bootNodeLocker.RUnlock()
	return num
}

func GetBestBootNode() (bootNodeAddr string) {
	bootNodeLocker.RLock()
	if len(MainnetBootnodes) != 0 {
		index = index % len(MainnetBootnodes)
		bootNodeAddr = MainnetBootnodes[index]
		index++
	}
	bootNodeLocker.RUnlock()
	return
}
