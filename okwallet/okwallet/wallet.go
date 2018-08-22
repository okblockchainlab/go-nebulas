package okwallet

import (
  "encoding/hex"

	"github.com/nebulasio/go-nebulas/crypto/keystore"
	"github.com/nebulasio/go-nebulas/core"
	"github.com/nebulasio/go-nebulas/crypto"
	"github.com/nebulasio/go-nebulas/account"
)

//reference: Manager.NewAccount
func GetAddressByPrivateKey(privkeyStr string) (string, error) {
  pk, err := hex.DecodeString(privkeyStr)
  if err != nil {
    return "", err
  }

  priv, err := crypto.NewPrivateKey(keystore.Algorithm(account.EccSecp256K1Value), nil)
  if err != nil {
    return "", err
  }

  if err = priv.Decode(pk); err != nil {
    return "", err
  }

  pub, err := priv.PublicKey().Encoded()
  if err != nil {
    return "", err
  }

  addr, err := core.NewAddressFromPublicKey(pub)
  if err != nil {
    return "", err
  }

  return addr.String(), nil
}
