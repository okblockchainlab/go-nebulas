package main

// #include <jni.h>
import "C"

import (
	"github.com/nebulasio/go-nebulas/okwallet/okwallet"
	"regexp"
)

const (
	GET_ADDRESS_BY_PRIVATE_KEY_CMD = "getaddressbyprivatekey"
	CREATE_RAW_TRANSACTION_CMD     = "createrawtransaction"
	SIGN_RAW_TRANSACTION           = "signrawtransaction"
)

func setErrorResult(env *C.JNIEnv, errMsg string) C.jobjectArray {
	result := newStringObjectArray(env, 1)
	setObjectArrayStringElement(env, result, 0, errMsg)
	return result
}

func getAddressByPrivateKeyExecute(env *C.JNIEnv, args []string) C.jobjectArray {
	if len(args) != 1 {
		return setErrorResult(env, "error: "+GET_ADDRESS_BY_PRIVATE_KEY_CMD+" wrong argument count")
	}

	addr, err := okwallet.GetAddressByPrivateKey(args[0])
	if err != nil {
		return setErrorResult(env, "error: "+err.Error())
	}

	result := newStringObjectArray(env, 1)
	setObjectArrayStringElement(env, result, 0, addr)

	return result
}

func createRawTransactionExecute(env *C.JNIEnv, args []string) C.jobjectArray {
	if !(len(args) == 7 || len(args) == 8) {
		return setErrorResult(env, "error: "+CREATE_RAW_TRANSACTION_CMD+" wrong argument count")
	}

	var binary string
	if len(args) == 8 {
		binary = args[7]
	}

	rawTx, err := okwallet.CreateRawTransaction(args[0], args[1], args[2], args[3], args[4], args[5], args[6], binary)
	if err != nil {
		return setErrorResult(env, "error: "+err.Error())
	}

	result := newStringObjectArray(env, 1)
	setObjectArrayStringElement(env, result, 0, rawTx)

	return result
}

func SignRawTransactionExecute(env *C.JNIEnv, args []string) C.jobjectArray {
	if len(args) != 2 {
		return setErrorResult(env, "error: "+SIGN_RAW_TRANSACTION+" wrong argument count")
	}

	signedTx, err := okwallet.SignRawTransaction(args[0], args[1])
	if err != nil {
		return setErrorResult(env, "error: "+err.Error())
	}

	result := newStringObjectArray(env, 1)
	setObjectArrayStringElement(env, result, 0, signedTx)

	return result
}

//export Java_com_okcoin_vault_jni_neb_Nebj_execute
func Java_com_okcoin_vault_jni_neb_Nebj_execute(env *C.JNIEnv, _ C.jclass, _ C.jstring, jcommand C.jstring) C.jobjectArray {
	command, err := jstring2string(env, jcommand)
	if err != nil {
		return setErrorResult(env, "error: "+err.Error())
	}

	sepExp, err := regexp.Compile(`\s+`)
	if err != nil {
		return setErrorResult(env, "error: "+err.Error())
	}

	args := sepExp.Split(command, -1)
	if len(args) < 2 {
		return setErrorResult(env, "error: invalid command")
	}

	switch args[0] {
	case GET_ADDRESS_BY_PRIVATE_KEY_CMD:
		return getAddressByPrivateKeyExecute(env, args[1:])
	case CREATE_RAW_TRANSACTION_CMD:
		return createRawTransactionExecute(env, args[1:])
	case SIGN_RAW_TRANSACTION:
		return SignRawTransactionExecute(env, args[1:])
	default:
		return setErrorResult(env, "error: unknown command: "+args[0])
	}
	return setErrorResult(env, "error: unknown command: "+args[0])
}

func main() {}
