package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/mainflux/mainflux/certs/pki"
)

const (
	TOKEN   = "hvs.0SERjuVSghZ8sI2ZQ2r2WfOG"
	HOST    = "http://127.0.0.1:8200"
	PATH    = "pki"
	ROLE    = "example-dot-com"
	TTL     = "3650d"
	KEYBITS = 2048
	KEYTYPE = "rsa"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {

	certPrefixName := RandStringRunes(5)

	j := 0
	for {
		j++
		Spanner(j, certPrefixName)
	}

}

func Spanner(j int, certPrefixName string) {

	var wg sync.WaitGroup
	for i := 0; i < 300; i++ {
		wg.Add(1)
		go VaultGenAndRevokeCert(fmt.Sprintf("%s%d.example.com", certPrefixName, i*j), &wg)
	}
	wg.Wait()

}

func VaultGenAndRevokeCert(cn string, wg *sync.WaitGroup) {
	fmt.Println(cn)
	vc, err := pki.NewVaultClient(TOKEN, HOST, PATH, ROLE)
	if err != nil {
		wg.Done()
		fmt.Println(err)
		return
	}

	cert, err := vc.IssueCert(cn, TTL, KEYTYPE, KEYBITS)
	if err != nil {
		wg.Done()
		fmt.Println(err)
		return
	}
	revTime, err := vc.Revoke(cert.Serial)
	if err != nil {
		wg.Done()
		fmt.Println(err)
		return
	}
	fmt.Printf("Issued ceritifcate %s , revoked time %v \n", cert.Serial, revTime)
	wg.Done()
}
