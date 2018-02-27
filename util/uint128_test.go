package util

import (
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	maxUint8  = ^uint8(0)
	maxUint64 = ^uint64(0)
)

func TestUint128(t *testing.T) {
	bigInt0 := big.NewInt(0)

	bigInt1 := big.NewInt(1)
	bigIntNeg1 := big.NewInt(-1)

	bigMaxUint8 := &big.Int{}
	bigMaxUint8.SetUint64(uint64(maxUint8))

	bigMaxUint64 := &big.Int{}
	bigMaxUint64.SetUint64(maxUint64)

	bigMaxUint64Add1 := &big.Int{}
	bigMaxUint64Add1.Add(bigMaxUint64, big.NewInt(1))

	bigUint128 := &big.Int{}
	bigUint128.Mul(bigMaxUint64, big.NewInt(67280421310721))

	bigMaxUint128 := &big.Int{}
	bigMaxUint128.SetString(strings.Repeat("f", 32), 16)

	bigMaxUint128Add1 := &big.Int{}
	bigMaxUint128Add1.Add(bigMaxUint128, big.NewInt(1))

	tests := []struct {
		input       *big.Int // input
		expected    [16]byte // expected Big-Endian result
		expectedErr error
	}{
		{bigInt0, [16]byte{
			0, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0}, nil},
		{bigInt1, [16]byte{
			0, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 1}, nil},
		{bigIntNeg1, [16]byte{}, ErrUint128Underflow},
		{bigMaxUint8, [16]byte{
			0, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 255}, nil},
		{bigMaxUint64, [16]byte{
			0, 0, 0, 0,
			0, 0, 0, 0,
			255, 255, 255, 255,
			255, 255, 255, 255}, nil},
		{bigMaxUint64Add1, [16]byte{
			0, 0, 0, 0,
			0, 0, 0, 1,
			0, 0, 0, 0,
			0, 0, 0, 0}, nil},
		{bigUint128, [16]byte{
			0, 0, 61, 48,
			241, 156, 209, 0,
			255, 255, 194, 207,
			14, 99, 46, 255}, nil},
		{bigMaxUint128, [16]byte{
			255, 255, 255, 255,
			255, 255, 255, 255,
			255, 255, 255, 255,
			255, 255, 255, 255}, nil},
		{bigMaxUint128Add1, [16]byte{}, ErrUint128Overflow},
	}
	for _, tt := range tests {
		u1, err := NewUint128FromBigInt(tt.input)
		if err != nil {
			assert.Equal(t, tt.expectedErr, err)
			continue
		}
		fsb, err := u1.ToFixedSizeBytes()
		fmt.Println("uint128.Int =", u1.Int, "bitlen =", u1.BitLen(), "[]bytes =", u1.Bytes(), "[16]bytes =", fsb, "err =", err)

		if tt.expectedErr != nil {
			assert.Equal(t, tt.expectedErr, err)
			continue
		}

		assert.Nil(t, u1.Validate(), "Validate doesn't pass.")
		assert.Equal(t, tt.expected, fsb, "ToFixedSizeBytes result doesn't match.")

		u2 := NewUint128FromFixedSizeBytes(fsb)
		assert.Equal(t, u1.Bytes(), u2.Bytes(), "FromFixedSizeBytes result doesn't match.")
	}
}
