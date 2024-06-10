package bls

import bls "github.com/herumi/bls-eth-go-binary/bls"

func GenerateKeypair() (string, string) {

	// Init lib
	bls.Init(bls.BLS12_381)

	// Generate keys
	secretKey := bls.SecretKey{}

	secretKey.SetByCSPRNG()

	publicKey := secretKey.GetPublicKey()

	return secretKey.SerializeToHexStr(), "0x" + publicKey.SerializeToHexStr()

}

func AggregatePubKeys(arrayOfPubKeysAsHexWith0x []string) string {

	// Init lib

	bls.Init(bls.BLS12_381)

	aggregatedPubKey := bls.PublicKey{}

	for _, pubKeyAsHexWith0x := range arrayOfPubKeysAsHexWith0x {

		componentKey := bls.PublicKey{}

		componentKey.DeserializeHexStr(pubKeyAsHexWith0x[2:])

		aggregatedPubKey.Add(&componentKey)

	}

	return "0x" + aggregatedPubKey.SerializeToHexStr()

}

func AggregateSignatures(arrayOfSignaturesAsHex []string) string {

	bls.Init(bls.BLS12_381)

	aggregatedSignature := bls.Sign{}

	for _, signatureAshex := range arrayOfSignaturesAsHex {

		componentSigna := bls.Sign{}

		componentSigna.DeserializeHexStr(signatureAshex)

		aggregatedSignature.Add(&componentSigna)

	}

	return aggregatedSignature.SerializeToHexStr()

}

func GenerateSignature(privateKeyAsHex, message string) string {

	bls.Init(bls.BLS12_381)

	// Recover private key from hex

	privateKey := bls.SecretKey{}

	privateKey.DeserializeHexStr(privateKeyAsHex)

	return privateKey.Sign(message).SerializeToHexStr()

}

func VerifySignature(pubKeyAsHexWith0x, message, signatureAsHex string) bool {

	bls.Init(bls.BLS12_381)

	// Recover public key and signature from hex

	publicKey := bls.PublicKey{}

	publicKey.DeserializeHexStr(pubKeyAsHexWith0x[2:])

	signature := bls.Sign{}

	signature.DeserializeHexStr(signatureAsHex)

	return signature.Verify(&publicKey, message)

}

func VerifyThresholdSignature(aggregatedPubkeyWhoSignAsHexWith0x, aggregatedSignatureAsHex, rootPubAsHexWith0x, message string, afkPubkeys []string, reverseThreshold uint) bool {

	if len(afkPubkeys) <= int(reverseThreshold) {

		verifiedSignature := VerifySignature(

			aggregatedPubkeyWhoSignAsHexWith0x,
			message,
			aggregatedSignatureAsHex,
		)

		if verifiedSignature {

			bls.Init(bls.BLS12_381)

			// If all the previos steps are OK - do the most CPU intensive task - pubkeys aggregation

			// Aggregate AFK signers

			aggregatedPubKeyOfAfkSignersWith0x := AggregatePubKeys(afkPubkeys)

			return AggregatePubKeys([]string{aggregatedPubKeyOfAfkSignersWith0x, aggregatedPubkeyWhoSignAsHexWith0x}) == rootPubAsHexWith0x

		} else {

			return false

		}

	} else {

		return false

	}

}
