package sync

import (
	"bytes"
	"fmt"
	"math/big"
	"strconv"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"strings"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/ethereum/go-ethereum/accounts/abi"
	"io"
	"encoding/json"
)

func hexoToString(input string) string {
	var buf bytes.Buffer
	x := new(big.Int)
	buf.Reset()
	buf.WriteString(input)
	if _, err := fmt.Fscanf(&buf, "%v", x); err != nil {
		fmt.Println(err)
		return "0"
	}
	return fmt.Sprintf(x.String())
}

func HexoToInt(input string) int64 {
	if input[:2] == "0x" {
		num, _ := strconv.ParseInt(input[2:], 16, 64)
		return num
	} else {
		num, _ := strconv.ParseInt(input, 16, 64)
		return num
	}

}

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	return hexutil.EncodeBig(number)
}

func marchalAddr(str string) (string, error) {
	const definition = `[
	{ "name" : "int", "constant" : false, "outputs": [ { "type": "uint256" } ] },
	{ "name" : "bool", "constant" : false, "outputs": [ { "type": "bool" } ] },
	{ "name" : "bytes", "constant" : false, "outputs": [ { "type": "bytes" } ] },
	{ "name" : "fixed", "constant" : false, "outputs": [ { "type": "bytes32" } ] },
	{ "name" : "multi", "constant" : false, "outputs": [ { "type": "bytes" }, { "type": "bytes" } ] },
	{ "name" : "intArraySingle", "constant" : false, "outputs": [ { "type": "uint256[3]" } ] },
	{ "name" : "addressSliceSingle", "constant" : false, "outputs": [ { "type": "address" } ] },
	{ "name" : "addressSliceDouble", "constant" : false, "outputs": [ { "name": "a", "type": "address[]" }, { "name": "b", "type": "address[]" } ] },
	{ "name" : "mixedBytes", "constant" : true, "outputs": [ { "name": "a", "type": "bytes" }, { "name": "b", "type": "bytes32" } ] }]`
	abi, err := jSONX(strings.NewReader(definition))
	if err != nil {
		return "", errors.New(fmt.Sprintf("[marchalAddr] jSONX failed str=%s,err=%s\n", str, err))
	}
	buff := new(bytes.Buffer)
	buff.Write(common.Hex2Bytes(str))
	var outAddr common.Address
	err = abi.Unpack(&outAddr, "addressSliceSingle", buff.Bytes())
	if err != nil {
		return "", errors.New(fmt.Sprintf("abi Unpack failed str=%s,err=%s\n", str, err))
	}
	return outAddr.String(), nil
}
func marchalInt(str string) (string, error) {
	const definition = `[
	{ "name" : "int", "constant" : false, "outputs": [ { "type": "uint256" } ] },
	{ "name" : "bool", "constant" : false, "outputs": [ { "type": "bool" } ] },
	{ "name" : "bytes", "constant" : false, "outputs": [ { "type": "bytes" } ] },
	{ "name" : "fixed", "constant" : false, "outputs": [ { "type": "bytes32" } ] },
	{ "name" : "multi", "constant" : false, "outputs": [ { "type": "bytes" }, { "type": "bytes" } ] },
	{ "name" : "intArraySingle", "constant" : false, "outputs": [ { "type": "uint256[3]" } ] },
	{ "name" : "addressSliceSingle", "constant" : false, "outputs": [ { "type": "address[]" } ] },
	{ "name" : "addressSliceDouble", "constant" : false, "outputs": [ { "name": "a", "type": "address[]" }, { "name": "b", "type": "address[]" } ] },
	{ "name" : "mixedBytes", "constant" : true, "outputs": [ { "name": "a", "type": "bytes" }, { "name": "b", "type": "bytes32" } ] }]`
	abi, err := jSONX(strings.NewReader(definition))
	if err != nil {
		return "", errors.New(fmt.Sprintf("[marchalInt] jSONX failed with str=%s,err=%s\n", str, err))
	}
	buff := new(bytes.Buffer)
	buff.Write(common.Hex2Bytes(str))
	var Int *big.Int
	err = abi.Unpack(&Int, "int", buff.Bytes())
	if err != nil {
		return "", errors.New(fmt.Sprintf("abi Unpack failed str=%s,err=%s\n", str, err))
	}
	return Int.String(), nil
}

func jSONX(reader io.Reader) (ABI, error) {
	dec := json.NewDecoder(reader)
	var abi ABI
	if err := dec.Decode(&abi); err != nil {
		return ABI{}, err
	}
	return abi, nil
}