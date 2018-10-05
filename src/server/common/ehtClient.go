package common

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"fmt"
)

var (
	EthConn  *ethclient.Client
)

func InitEthClient (){
	//conn, err := ethclient.Dial("http://0.0.0.0:10070")
	//conn, err := ethclient.Dial("http://127.0.0.1:10070")
	conn, err := ethclient.Dial("http://127.0.0.1:8545")
	//conn, err := ethclient.Dial("\\\\.\\pipe\\geth.ipc")
	if nil != err {
		log.Panic(err)
	}

	EthConn = conn
	fmt.Println("conn: %v", conn)
}
