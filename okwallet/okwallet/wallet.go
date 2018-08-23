package okwallet

import (
	"encoding/hex"
	"errors"
	"strconv"

	"github.com/gogo/protobuf/proto"
	"github.com/nebulasio/go-nebulas/account"
	"github.com/nebulasio/go-nebulas/core"
	"github.com/nebulasio/go-nebulas/crypto"
	"github.com/nebulasio/go-nebulas/crypto/keystore"
	"github.com/nebulasio/go-nebulas/util"
)

func getBinaryPayload(binaryStr string) (string, []byte, error) {
	var err error
	var binary []byte
	if len(binaryStr) > 0 {
		binary, err = hex.DecodeString(binaryStr)
		if err != nil {
			return "", nil, err
		}
	}

	payload, err := core.NewBinaryPayload(binary).ToBytes()
	if err != nil {
		return "", nil, err
	}

	return core.TxPayloadBinaryType, payload, nil
}

func parseTransaction(chainID uint32, fromStr, toStr, valueStr, nonceStr, gasPriceStr, gasLimitStr, binaryStr string) (*core.Transaction, error) {
	fromAddr, err := core.AddressParse(fromStr)
	if err != nil {
		return nil, err
	}
	toAddr, err := core.AddressParse(toStr)
	if err != nil {
		return nil, err
	}

	value, err := util.NewUint128FromString(valueStr)
	if err != nil {
		return nil, errors.New("invalid value")
	}
	nonce, err := strconv.ParseUint(nonceStr, 10, 64)
	if err != nil {
		return nil, errors.New("invalid nonce number")
	}
	gasPrice, err := util.NewUint128FromString(gasPriceStr)
	if err != nil {
		return nil, errors.New("invalid gasPrice")
	}
	gasLimit, err := util.NewUint128FromString(gasLimitStr)
	if err != nil {
		return nil, errors.New("invalid gasLimit")
	}

	//we just care about binary payload
	payloadType, payload, err := getBinaryPayload(binaryStr)
	if err != nil {
		return nil, err
	}

	tx, err := core.NewTransaction(chainID, fromAddr, toAddr, value, nonce, payloadType, payload, gasPrice, gasLimit)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

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

//reference: AdminService.SignTransactionWithPassphrase
func CreateRawTransaction(chainIDStr, from, to, value, nonce, gasPrice, gasLimit, binary string) (string, error) {
	chainID, err := strconv.ParseUint(chainIDStr, 10, 32)
	if err != nil {
		return "", errors.New("Invalid chain id")
	}

	tx, err := parseTransaction(uint32(chainID), from, to, value, nonce, gasPrice, gasLimit, binary)
	if err != nil {
		return "", err
	}

	pbMsg, err := tx.ToProto()
	if err != nil {
		return "", err
	}
	data, err := proto.Marshal(pbMsg)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(data), nil
}
