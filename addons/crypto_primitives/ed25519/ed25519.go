package ed25519

import (
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"

	"github.com/btcsuite/btcutil/base58"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

type Ed25519Box struct {
	Mnemonic  string
	Bip44Path []uint32
	Pub, Prv  string
}

func GenerateKeyPair(mnemonic, mnemonicPassword string, bip44DerivePath []uint32) Ed25519Box {

	if mnemonic == "" {

		// Generate mnemonic if no pre-set

		entropy, _ := bip39.NewEntropy(256)

		mnemonic, _ = bip39.NewMnemonic(entropy)

	}

	// Now generate seed from 24-word mnemonic phrase (24 words = 256 bit security)
	// Seed has 64 bytes
	seed := bip39.NewSeed(mnemonic, mnemonicPassword) // password might be ""(empty) but it's not recommended

	// Generate master keypair from seed

	masterPrivateKey, _ := bip32.NewMasterKey(seed)

	// Now, to derive appropriate keypair - run the cycle over uint32 path-milestones and derive child keypairs

	// In case bip44Path empty - set the default one

	if len(bip44DerivePath) == 0 {

		bip44DerivePath = []uint32{44, 7331, 0, 0}

	}

	// Start derivation from master private key
	var childKey *bip32.Key = masterPrivateKey

	for pathPart := range bip44DerivePath {

		childKey, _ = childKey.NewChildKey(bip32.FirstHardenedChild + uint32(pathPart))

	}

	// Now, based on this - get the appropriate keypair

	publicKeyObject, privateKeyObject := generateKeyPairFromSeed(childKey.Key)

	// Export keypair

	pubKeyBytes, _ := x509.MarshalPKIXPublicKey(publicKeyObject)

	privKeyBytes, _ := x509.MarshalPKCS8PrivateKey(privateKeyObject)

	return Ed25519Box{Mnemonic: mnemonic, Bip44Path: bip44DerivePath, Pub: base58.Encode(pubKeyBytes[12:]), Prv: base64.StdEncoding.EncodeToString(privKeyBytes)}

}

// Returns signature in base64(to use it in transaction later)

func GenerateSignature(privateKeyAsBase64, msg string) string {

	// Decode private key from base64 to raw bytes

	privateKeyAsBytes, _ := base64.StdEncoding.DecodeString(privateKeyAsBase64)

	// Deserialize private key

	privKeyInterface, _ := x509.ParsePKCS8PrivateKey(privateKeyAsBytes)

	finalPrivateKey, _ := privKeyInterface.(ed25519.PrivateKey)

	msgAsBytes := []byte(msg)

	signature, _ := finalPrivateKey.Sign(rand.Reader, msgAsBytes, crypto.Hash(0))

	return base64.StdEncoding.EncodeToString(signature)

}

/*
0 - message that was signed
1 - pubKey
2 - signature
*/
func VerifySignature(stringMessage, base58PubKey, base64Signature string) bool {

	// Decode evrything

	msgAsBytes := []byte(stringMessage)

	publicKeyAsBytesWithNoAsnPrefix := base58.Decode(base58PubKey)

	// Add ASN.1 prefix

	pubKeyAsBytesWithAsnPrefix := append([]byte{0x30, 0x2a, 0x30, 0x05, 0x06, 0x03, 0x2b, 0x65, 0x70, 0x03, 0x21, 0x00}, publicKeyAsBytesWithNoAsnPrefix...)

	pubKeyInterface, _ := x509.ParsePKIXPublicKey(pubKeyAsBytesWithAsnPrefix)

	finalPubKey, _ := pubKeyInterface.(ed25519.PublicKey)

	signature, _ := base64.StdEncoding.DecodeString(base64Signature)

	return ed25519.Verify(finalPubKey, msgAsBytes, signature)

}

// Private inner function

func generateKeyPairFromSeed(seed []byte) (ed25519.PublicKey, ed25519.PrivateKey) {

	privateKey := ed25519.NewKeyFromSeed(seed)

	pubKey, _ := privateKey.Public().(ed25519.PublicKey)

	return pubKey, privateKey

}
