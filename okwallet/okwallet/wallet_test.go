package okwallet

import (
	"testing"
)

func TestGetAddressByPrivateKey(t *testing.T) {
	privkey1 := `c9312e9c279309ec9df68bd374124676f2bf9d050a53c6fab35045c62d15cc9c`
	const expect1 = `n1TcMVWCuE6vLfhugpSQdk92vHH7thGLnZt`
	privkey2 := `734daad005b7a69c42d7c25e286bc5f376ac757f4869f0d521245b46af74b7e6`
	const expect2 = `n1HSJnmDxvUWejKjKkWQvBKCK1d7AVFZckq`

	addr, err := GetAddressByPrivateKey(privkey1)
	if err != nil {
		t.Fatal(err)
	}
	if expect1 != addr {
		msg := "GetAddressByPrivateKey failed. expect " + expect1 + " but return " + addr
		t.Fatal(msg)
	}

	addr, err = GetAddressByPrivateKey(privkey2)
	if err != nil {
		t.Fatal(err)
	}
	if expect2 != addr {
		msg := "GetAddressByPrivateKey failed. expect " + expect2 + " but return " + addr
		t.Fatal(msg)
	}
}

func TestCreateRawTransaction(t *testing.T) {
	from := `n1TcMVWCuE6vLfhugpSQdk92vHH7thGLnZt`
	to := `n1HSJnmDxvUWejKjKkWQvBKCK1d7AVFZckq`
	value := `1000000000000000000`
	nonce := `1000`
	gasPrice := `1000000`
	gasLimit := `2000000`
	binary := ""
	const expect = `121a19578fa4152b27a53c34eb79718e722f5631fcd1daa90b94b66f1a1a1957200c68cba489da88c9c8887df46ce13a9024c1c31c55ca7a221000000000000000000de0b6b3a764000028e80730d4dbf8db053a080a0662696e617279400b4a10000000000000000000000000000f42405210000000000000000000000000001e8480`

	rawTx, err := CreateRawTransaction("11", from, to, value, nonce, gasPrice, gasLimit, binary)
	if err != nil {
		t.Fatal(err)
	}
	if expect != rawTx {
		t.Fatal("CreateRawTransaction failed. exptect result " + expect + " but return " + rawTx)
	}
}
