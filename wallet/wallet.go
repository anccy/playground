package wallet

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func CreateWallet() *ecdsa.PrivateKey {
	// generate a private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("private: ", hexutil.Encode(privateKeyBytes))

	// generate a public key
	publicKey := privateKey.PublicKey
	address := crypto.PubkeyToAddress(publicKey).Hex()
	fmt.Println("public: ", address)

	return privateKey
}

func CreateKs() {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount("secret")
	if err != nil {
        log.Fatal(err)
    }
	fmt.Println("address: ", account.Address.Hex())
}

func ImportKs() {
	file := ""
    ks := keystore.NewKeyStore("./tmp1", keystore.StandardScryptN, keystore.StandardScryptP)
	fmt.Println(ks.Accounts())
    jsonBytes, err := os.ReadFile(file)
    if err != nil {
        log.Fatal(err)
    }

    password := "secret"
    account, err := ks.Import(jsonBytes, password, "yyy")
	fmt.Println("address: ", account.Address.Hex())
    if err != nil {
        log.Fatal("import error: ", err)
    }

    if err := os.Remove(file); err != nil {
        log.Fatal(err)
    }
}
