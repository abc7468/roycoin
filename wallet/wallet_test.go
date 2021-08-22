package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"io/fs"
	"reflect"
	"testing"

	"github.com/abc7468/roycoin/utils"
)

const (
	testKey     string = "307702010104201ca6df5f18f8e1f41491f40b1be7aecf71adf93b503a7fe1fa51ac565e0d9153a00a06082a8648ce3d030107a144034200046a626f7bceb3840b5e6699ca51bb0e7c0c4ee2352f0fa75968cbe806f5a4a3fe9cc120ce1d69062d467ba43c09274f455efe212f224a5d16f1ff3a6eae6316d8"
	testPayload string = "0005e9c1cccae0e1c904a29174e72a00646d762045dbbb94f55c126b7ca41063"
	testSig     string = "b3972c349697fcede7cdc3ef96c2bfc9d25e6824df8d9c6011e5f90a29448089bd87914953f9f80cb173013c4eac038c1ee540f3d39b72bf56da69e37e8969f4"
)

type fakeLayer struct {
	fakeHasWalletFile func() bool
}

func TestWallet(t *testing.T) {
	t.Run("Wallet is created", func(t *testing.T) {
		files = fakeLayer{
			fakeHasWalletFile: func() bool { return false },
		}
		tw := Wallet()
		if reflect.TypeOf(tw) != reflect.TypeOf(&wallet{}) {
			t.Error("New Wallet should return a new wallet instance")
		}
	})
	t.Run("Wallet is restore", func(t *testing.T) {
		files = fakeLayer{
			fakeHasWalletFile: func() bool { return false },
		}
		w = nil
		tw := Wallet()
		if reflect.TypeOf(tw) != reflect.TypeOf(&wallet{}) {
			t.Error("New Wallet should return a new wallet instance")
		}
	})
}
func (f fakeLayer) hasWalletFile() bool {
	return f.fakeHasWalletFile()
}

func (fakeLayer) writeFile(name string, data []byte, perm fs.FileMode) error {

	return nil
}

func (fakeLayer) readFile(name string) ([]byte, error) {
	return utils.ToBytes(makeTestWallet().privateKey), nil
}

func makeTestWallet() *wallet {
	w := &wallet{}
	b, _ := hex.DecodeString(testKey)
	key, _ := x509.ParseECPrivateKey(b)
	w.privateKey = key
	w.Address = aFromK(key)
	return w
}

// func TestVerify(t *testing.T) {
// 	privKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
// 	b, _ := x509.MarshalECPrivateKey(privKey)
// 	t.Logf("%x", b)
// }

func TestSign(t *testing.T) {
	s := Sign(testPayload, makeTestWallet())
	_, err := hex.DecodeString(s)
	if err != nil {
		t.Errorf("Sign() should return a hex encoded string, got %s", s)
	}
}

func TestVerify(t *testing.T) {
	type test struct {
		input string
		ok    bool
	}
	tests := []test{
		{testPayload, true},
		{"0105e9c1cccae0e1c904a29174e72a00646d762045dbbb94f55c126b7ca41063", false},
	}
	for _, tc := range tests {
		w := makeTestWallet()
		ok := Verify(testSig, tc.input, w.Address)
		if ok != tc.ok {
			t.Error("Verify() could not verify testSignature and testPayload")
		}
	}
}

func TestRestoreBigInts(t *testing.T) {
	_, _, err := restoreBigInt("xx")
	if err == nil {
		t.Error("restoreBigInts should return error when payload is not hex.")
	}
}
