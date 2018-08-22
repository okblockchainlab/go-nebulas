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
