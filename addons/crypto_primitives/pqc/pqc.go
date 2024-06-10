/*


                Links:

https://github.com/LoCCS/bliss/search?q=entropy
https://github.com/LoCCS/bliss


*/

package pqc

import (
	"math/rand"

	"time"

	"github.com/cloudflare/circl/sign/dilithium"

	"github.com/LoCCS/bliss/sampler"

	"github.com/LoCCS/bliss"

	"encoding/hex"
)

//______________________________ Dilithium ______________________________

var modename string = "Dilithium5" // Dilithium2-AES Dilithium3 Dilithium3-AES Dilithium5 Dilithium5-AES

var mode = dilithium.ModeByName(modename)

func GenerateDilithiumKeypair() (string, string) {

	publicKey, privateKey, _ := mode.GenerateKey(nil)

	return hex.EncodeToString(publicKey.Bytes()), hex.EncodeToString(privateKey.Bytes())

}

/*
0 - privateKey
1 - message
*/
func GenerateDilithiumSignature(privateKey, msg string) string {

	privateKeyAsBytes, _ := hex.DecodeString(privateKey)

	msgAsBytes := []byte(msg)

	return hex.EncodeToString(mode.Sign(mode.PrivateKeyFromBytes(privateKeyAsBytes), msgAsBytes))

}

/*
0 - message that was signed
1 - pubKey
2 - signature
*/
func VerifyDilithiumSignature(msg, pubKey, hexSignature string) bool {

	msgAsBytes := []byte(msg)

	publicKey, _ := hex.DecodeString(pubKey)

	signature, _ := hex.DecodeString(hexSignature)

	return mode.Verify(mode.PublicKeyFromBytes(publicKey), msgAsBytes, signature)

}

//________________________________ BLISS ________________________________

func GenerateBlissKeypair() (string, string) {

	rand.Seed(time.Now().UnixNano())

	seed := make([]byte, sampler.SHA_512_DIGEST_LENGTH)

	rand.Read(seed)

	entropy, _ := sampler.NewEntropy(seed)

	prv, _ := bliss.GeneratePrivateKey(0, entropy)

	pub := prv.PublicKey()

	return hex.EncodeToString(pub.Encode()), hex.EncodeToString(seed)

}

/*
0 - privateKey
1 - message
*/
func GenerateBlissSignature(pritaveKey, msg string) string {

	//Decode msg an seed => entropy => privateKey

	sid, _ := hex.DecodeString(pritaveKey)

	msgAsBytes := []byte(msg)

	seed := []byte(sid) // uint8/byte array

	entropy, _ := sampler.NewEntropy(seed)

	key, _ := bliss.GeneratePrivateKey(0, entropy)

	//Gen signature
	sig, _ := key.Sign(msgAsBytes, entropy)

	return hex.EncodeToString(sig.Encode())

}

/*
0 - message
1 - publicKey
2 - signature
*/
func VerifyBlissSignature(msg, hexPublicKey, hexSignature string) bool {

	//Decode msg an publicKey
	msgAsBytes := []byte(msg)

	hexEncodedPublicKey, _ := hex.DecodeString(hexPublicKey)

	publicKey, _ := bliss.DecodePublicKey(hexEncodedPublicKey)

	//Decode signature
	decodedSignature, _ := hex.DecodeString(hexSignature)

	signature, _ := bliss.DecodeSignature(decodedSignature)

	//Verification itself
	_, err := publicKey.Verify(msgAsBytes, signature)

	return err == nil

}
