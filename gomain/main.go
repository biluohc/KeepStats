package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	abis "abis"
)

// geth支持直接根据abi生成go代码 https://www.coder.work/article/195746
// go build main.go
func main() {
	client, err := ethclient.Dial("https://ropsten.infura.io/v3/")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x60535a59b4e71f908f3feb0116f450703fb35ed8")
	instance, err := abis.NewKeepBonding(address, client)
	if err != nil {
		log.Fatal(err)
	}

	session := abis.KeepBondingSession{
		Contract:     instance,
		CallOpts:     bind.CallOpts{},
		TransactOpts: bind.TransactOpts{},
	}

	value, err := session.UnbondedValue(common.HexToAddress("0xA1Dc51A02070ccAFA8882D378644419eED530eb9"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(value.Div(value, big.NewInt(1_000_000_000_000_000)))
}
