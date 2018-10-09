package service

import (
	"blockclient/src/server/entity"
	"blockclient/src/server/common/stacode"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	myCommon "blockclient/src/server/common"
	"log"
	"strings"
	"context"
	"math/big"
	"fmt"
	"time"
	"bytes"
	"blockclient/src/server/tokens"




)

type EthService struct {
}


// get balance of a account by account's address
func (es *EthService) GetBalanceByAddress(addresses, num string) *entity.Result {
	if "" == strings.TrimSpace(addresses) {
		return errcode.REQUEST_PARAM_ERR.Result(nil)
	}
	//balance, err := myCommon.EthConn.BalanceAt(context.Background(),  common.HexToAddress(address), big.NewInt(number))
	block, err := myCommon.EthConn.BlockByNumber(context.Background(), nil)
	if nil != err {
		log.Println("get blockNumber err:", err)
		return errcode.SYSTEMBUSY_ERROR.Result(nil)
	}else {
		log.Println("blockNumber:", block.Number().String())
	}
	resMap := make(map[string]interface{}, 0)
	for _, address := range strings.Split(addresses, ",") {
		balance, err := myCommon.EthConn.BalanceAt(context.Background(),  common.HexToAddress(address), nil)
		if nil != err {
			log.Println("Failed to query ether balance by address: %v", err)
			log.Println("Failed to query ether balance by address: ", err.Error())
			return errcode.SYSTEMBUSY_ERROR.Result(nil)
		}
		resMap[address] = balance
	}
	return errcode.SUCCESS.Result(resMap)
}

func (es *EthService) DeployContract(total int64, decimals uint8, name, symbol, key, passphrase string) *entity.Result {

	// 【说明】
	// 安装solidity编译器solc (为 solcjs)： sudo npm install -g solc
	// 设置链接至 /usr/bin/中 (虽然可以用solcjs已经设置在 /usr/local/bin中了，但是有些地方需要以solc 启动，如abigen就是)：sudo ln -s /usr/local/lib/node_modules/solc/solcjs /usr/bin/solc
	// 查看下是否可以以solc访问: solc --help
	// 进入到项目中的  *sol 文件所在目录： cd ~/go-workspace/src/blockclient/src/server/tokens/
	// 生成abi文件： solc GavinToken.sol -o filedir --abi
	// 生成bin文件： solc GavinToken.sol -o filedir --bin
	// 生成go文件： abigen --abi filedir/GavinToken_sol_GavinToken.abi --bin filedir/GavinToken_sol_GavinToken.bin --pkg mytoken --out gavinToken.go
	// 【注意】：网上的sol直接生成go文件是有问题的可能会报一下错误： abigen --sol GavinToken.sol --pkg mytoken --out gavinToken.go
	// Failed to build Solidity contract: solc: exit status 1
	// Invalid option selected, must specify either --bin or --abi
	// 或者
	// Failed to build Solidity contract: exit status 7


	if "" == strings.TrimSpace(passphrase) {
		return errcode.REQUEST_PARAM_ERR.ResultWithMsg("must have pwd")
	}

	auth, err := bind.NewTransactor(strings.NewReader(key), passphrase)
	if nil != err {
		log.Panic("Failed to create authorized transactor: %v", err)
	}
	// 发布合约返回 合约地址、 当前交易信息、 和当前的合约实例(操作合约abi用)
	address, tx, token, err := mytoken.DeployMytoken(auth, myCommon.EthConn, big.NewInt(total), name, decimals, symbol)
	if nil != err {
		log.Panic("Failed to deploy new token contract: %v", err)
	}
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
	startTime := time.Now()
	fmt.Printf("TX start @:%s", time.Now())

	// 等待挖矿确认
	addressAfterMined, err := bind.WaitDeployed(context.Background(), myCommon.EthConn, tx)
	if nil != err {
		log.Panic("failed to deploy contact when mining :%v", err)
	}
	fmt.Printf("tx mining take time:%s\n", time.Now().Sub(startTime))

	if bytes.Compare(address.Bytes(), addressAfterMined.Bytes()) != 0 {
		log.Panic("mined address :%s,before mined address:%s", addressAfterMined, address)
	}

	// token的Name
	nameRes, err := token.Name(&bind.CallOpts{Pending: true})
	if nil != err {
		log.Panic("Failed to retrieve pending name: %v", err)
	}
	totalRes, err := token.TotalSupply(&bind.CallOpts{Pending: true})
	if nil != err {
		log.Panic("Failed to retrieve pending total: %v", err)
	}
	symbolRes, err := token.Symbol(&bind.CallOpts{Pending: true})
	if nil != err {
		log.Panic("Failed to retrieve pending symbol: %v", err)
	}
	decimalsRes, err := token.Decimals(&bind.CallOpts{Pending: true})
	if nil != err {
		log.Panic("Failed to retrieve pending decimals: %v", err)
	}

	res := make(map[string]interface{}, 0)
	res["txHAsh"] = tx.Hash()
	res["contractAddress"] = address.String()
	res["Name"] = nameRes
	res["total"] = totalRes.String()
	res["symbol"] = symbolRes
	res["decimals"] = decimalsRes

	return errcode.SUCCESS.Result(res)
}

func (es *EthService) TokenTransfer(contractAddress, to, key, pwd string, value *big.Int) *entity.Result {
	//go es.SubcribeEvents(contractAddress, to)
	token, err := mytoken.NewMytoken(common.HexToAddress(contractAddress), myCommon.EthConn)
	if err != nil {
		log.Panic("Failed to instantiate a Token contract: %v", err)
	}
	toAddress := common.HexToAddress(to)
	toPreval, _ := token.BalanceOf(nil, toAddress)

	// Create an authorized
	auth, err := bind.NewTransactor(strings.NewReader(key), pwd)
	if err != nil {
		log.Panic("Failed to create authorized transactor: %v", err)
	}

	// Create a transfer
	tx, err := token.Transfer(auth, toAddress, value)
	if err != nil {
		log.Panic("Failed to request token transfer: %v", err)
	}

	//var keyStorage entity.KeyStorage
	//json.Unmarshal([]byte(key), &keyStorage)
	//from := keyStorage.Address
	// start-up transfer event subcribe


	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
	startTime := time.Now()
	fmt.Printf("TX start @:%s", time.Now())

	// waiting mining
	receipt, err := bind.WaitMined(context.Background(), myCommon.EthConn, tx)
	if err != nil {
		log.Panic("tx mining error:%v\n", err)
	}
	fmt.Printf("tx mining take time:%s\n", time.Now().Sub(startTime))

	val, _ := token.BalanceOf(nil, toAddress)

	res := make(map[string]interface{}, 0)
	res["toPreval"] = toPreval
	res["tx"] = tx.Hash()
	res["toVal"] = val
	res["receipt.TxHash"] = receipt.TxHash
	recp, _ := receipt.MarshalJSON()
	res["receiptJson"] = string(recp)
	return errcode.SUCCESS.Result(res)
}

func (es *EthService) QueryTokenBalance(contractAddress, eoas string) *entity.Result {
	if "" == strings.TrimSpace(contractAddress) || "" == strings.TrimSpace(eoas) {
		return errcode.REQUEST_PARAM_ERR.ResultWithMsg("contract and from  must is not empty")
	}
	token, err := mytoken.NewMytoken(common.HexToAddress(contractAddress), myCommon.EthConn)
	if err != nil {
		log.Panic("Failed to instantiate a Token contract: %v", err)
	}
	resMap := make(map[string]interface{}, 0)
	for _, from := range strings.Split(eoas, ",") {
		val, _ := token.BalanceOf(nil, common.HexToAddress(from))
		resMap[from] = val.String()
	}
	return errcode.SUCCESS.Result(resMap)
}


func (es *EthService) SubcribeEvents(contractAddress, to string) {
	filter, err := mytoken.NewMytokenFilterer(common.HexToAddress(contractAddress), myCommon.EthConn)
	ch := make(chan *mytoken.MytokenTransfer, 10)
	//sub, err := filter.WatchTransfer(&bind.WatchOpts{}, ch, []common.Address{common.HexToAddress(from)}, []common.Address{common.HexToAddress(to)})
	sub, err := filter.WatchTransfer(nil, ch, nil, []common.Address{common.HexToAddress(to)})
	if err != nil {
		log.Panic("watch transfer err %s", err)
	}
	go func() {
		for {
			select {
			case <-sub.Err():
				return
			case e := <-ch:
				log.Printf("new transfer event from %s to %s value=%s,at %d",
					e.From.String(), e.To.String(), e.Value, e.Raw.BlockNumber)
			}
		}
	}()
}



//func (es *EthService) CreateAccount (passphrase string) *entity.Result {
//	keydir, err := filepath.Abs("keystore")
//	if nil != err {
//		log.Panic("Failed to read configuration: %v", err)
//	}
//	myKeystore := geth.NewKeyStore(keydir, keystore.StandardScryptN, keystore.StandardScryptP)
//	myAccount, err := myKeystore.NewAccount(passphrase)
//	if nil != err {
//		log.Panic("Failed to Create Account: %v", err)
//	}
//	address := myAccount.GetAddress().GetHex()
//	return errcode.SUCCESS.Result(address)
//}
