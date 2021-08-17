package wallet

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/abc7468/roycoin/utils"
)

const (
	hashedMessage string = "1c5863cd55b5a4413fd59f054af57ba3c75c0698b3851d70f99b8de2d5c7338f"
	privateKey    string = "30770201010420876183c275b5371a4d0fa4494695a0ced7fee65db26b89efd1a857309af492a2a00a06082a8648ce3d030107a144034200043a772c969626888123924510367e9b6c51dc6a36ff39ed1317a0f9ca495631ec1f9ab73fbad4e60988663f42c552565334420d8700d46ab0c12e08391df1898e"
	signature     string = "d9a7d4033a86bc91322217b670fa285dac6cdad87cd0e196c3cdce8c34f64c39cbcf7105fa2eb747ac304f1749c04028a812a42866d549fb395cf7fa7dd638e1"
)

func Start() {
	privByte, err := hex.DecodeString(privateKey)
	utils.HandleErr(err)
	private, err := x509.ParseECPrivateKey(privByte)
	utils.HandleErr(err)

	sigBytes, err := hex.DecodeString(signature)
	rBytes := sigBytes[:len(sigBytes)/2]
	sBytes := sigBytes[len(sigBytes)/2:]

	var bigR, bigS = big.Int{}, big.Int{}
	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)

	hashBytes, err := hex.DecodeString(hashedMessage)
	utils.HandleErr(err)

	ok := ecdsa.Verify(&private.PublicKey, hashBytes, &bigR, &bigS)
	fmt.Println(ok)
	// privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	// keyAsBytes, err := x509.MarshalECPrivateKey(privateKey)
	// fmt.Printf("%x\n\n", keyAsBytes)
	// utils.HandleErr(err)

	// hashAsBytes, err := hex.DecodeString(hashedMessage)
	// utils.HandleErr(err)

	// r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsBytes)

	// signature := append(r.Bytes(), s.Bytes()...)

	// fmt.Printf("%x", signature)
	// utils.HandleErr(err)

}
