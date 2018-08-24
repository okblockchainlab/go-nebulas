package okwallet

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"

	"github.com/gogo/protobuf/proto"
	"github.com/nebulasio/go-nebulas/core"
	"github.com/nebulasio/go-nebulas/core/pb"
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

func privateKeyFromString(keyStr string) (keystore.PrivateKey, error) {
	pkByte, err := hex.DecodeString(keyStr)
	if err != nil {
		return nil, err
	}

	key, err := crypto.NewPrivateKey(keystore.SECP256K1, nil)
	if err != nil {
		return nil, err
	}

	if err = key.Decode(pkByte); err != nil {
		return nil, err
	}

	return key, nil
}

func getAddress(privkeyStr string) (*core.Address, error) {
	priv, err := privateKeyFromString(privkeyStr)
	if err != nil {
		return nil, err
	}

	pub, err := priv.PublicKey().Encoded()
	if err != nil {
		return nil, err
	}

	addr, err := core.NewAddressFromPublicKey(pub)
	if err != nil {
		return nil, err
	}

	return addr, nil
}

//reference: Manager.NewAccount
func GetAddressByPrivateKey(privkeyStr string) (string, error) {
	addr, err := getAddress(privkeyStr)
	if err != nil {
		return "", err
	}

	return addr.String(), nil
}

//reference: AdminService.SignTransactionWithPassphrase
func CreateRawTransaction(chainIDStr, from, to, value, nonce, gasPrice, gasLimit, bin string) (string, error) {
	chainID, err := strconv.ParseUint(chainIDStr, 10, 32)
	if err != nil {
		return "", errors.New("Invalid chain id")
	}

	tx, err := parseTransaction(uint32(chainID), from, to, value, nonce, gasPrice, gasLimit, bin)
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

	//resp := rpcpb.SignTransactionPassphraseResponse{Data:data}
	//resp.String

	return hex.EncodeToString(data), nil
}

//reference: APIService.SendRawTransaction, AdminService.SignTransactionWithPassphrase
func SignRawTransaction(rawTx, prvkeyStr string) (string, error) {
	rawTxByte, err := hex.DecodeString(rawTx)
	if err != nil {
		return "", fmt.Errorf("invalid rawtx data: %v", err)
	}

	pbTx := new(corepb.Transaction)
	if err := proto.Unmarshal(rawTxByte, pbTx); err != nil {
		return "", err
	}
	tx := new(core.Transaction)
	if err := tx.FromProtoWithoutAlgCheck(pbTx); err != nil {
		return "", err
	}

	addr, err := getAddress(prvkeyStr)
	if err != nil {
		return "", err
	}

	// check sign addr is tx's from addr
	if !tx.From().Equals(addr) {
		return "", errors.New("transaction sign not use from address")
	}

	key, err := privateKeyFromString(prvkeyStr)
	if err != nil {
		return "", err
	}

	signature, err := crypto.NewSignature(keystore.SECP256K1)
	if err != nil {
		return "", err
	}
	signature.InitSign(key)
	if err = tx.Sign(signature); err != nil {
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

	return base64.StdEncoding.EncodeToString(data), nil
}
