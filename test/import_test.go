package test

import (
	"encoding/hex"
	"regexp"
	"testing"

	"github.com/descope/virtualwebauthn"
	_ "github.com/fxamacker/webauthn/packed"
	"github.com/stretchr/testify/require"
)

// Generated with:
//
//	openssl ecparam -name prime256v1 -genkey -out ec2.pem
//	openssl pkcs8 -in ec2.pem -outform der -topk8 -nocrypt | xxd -p -c 60
const ec2Key = `
	308187020100301306072a8648ce3d020106082a8648ce3d030107046d306b02010104201d5a3d9d61c14b69ddb95dcfe2829e54c8859d7c5bc1a352
	b48aa626ca18c53aa144034200043b9031636d2f754b5a16cf99d4b6ab288beeeecaa0a16f5f92b08bd44cb01f581c7b09189bb56263d2fe3b03cdef
	7b7e3e83ab9aa65c295ae4a19f96b397400b`

// Generated with:
//
//	openssl genrsa -out rsa.pem 2048
//	openssl pkcs8 -in rsa.pem -outform der -topk8 -nocrypt | xxd -p -c 60
const rsaKey = `
	308204be020100300d06092a864886f70d0101010500048204a8308204a40201000282010100b69ce132433632f96942381775c9e5d4c921b82301a0
	5006d5891e0db99df942e86685782ba942e3f90a69118f5f135cb8699265bba67025708dfc4ce579b33c0b42a888bfd312442cf8da49424e9270c5bf
	292b537de8dcc46ed898695099c741081ea058978b070c7ba1e1177e9ac32be5e458120160d93bdf0d07e5703057ba8a41db48fbcd2a6ad86a2b471c
	d9e6562053031fa6b45e39f260f9cb8a1e0ec5cced2a477c57f7473cc348ab55daa896a7c237e6d0f32f56bbf7e3528a7cbf416ec252cb47fb8f5d18
	df909bbc99d18930eba70e8eec68535658248c447eebbff045f0e351f6668f201a1f5ad517be6ac47387a674eaeb1f1dff157f2c58c7020301000102
	8201004a8d99d2ef65bd41de1b4ed3251f9c595714111d1998dd932cb2a236704572724277389e6b14db5e3f5a64b2ea99a24a046ff578db37842984
	de32d72208a1882e00b4d5bf9ae8a634c614383c7ddf8372f82c52a7ef8b96360c1b197e458dc2af66253ac970752e178691fc579f3487e1f0255d5f
	5c78b1b7a3c4aa289db8de26a58356cdb1c69fd54ca1722693f0f5d182752b8e1ca02d12832b98ff6b39953ee7ceb0eddc255d3b0807e84edb169363
	40e0a49283e51a30fda414d619d1d46de5724f6d72e238b67699f3df60a81f984b10b6868725fd7ed73c9620e3f3adce85c9a2a6098eb507ac47a08c
	caf3bd7e6c538a1043ec8ac9116581ab84004102818100efee7668ca84dbbbf0518a95d49ad57fd21695654cdfb1346f1f3d4523a0edba001025950f
	b79300d9f8215cff2233fef97568f313e4aea3e492542ea7ee0c8010f2fbfb80cbe6cba39ffd5844c56d14497d2c43fa3e763533f0e12e3bb34117d5
	2b2d9728108b42697d59db55d5020ef9eda4e5c0bb9a6088e763a69dc4ddf902818100c2d7b5974d340a7bf37615853846c6e17f678eb6dc26f9e29e
	8abf51326a141ecdf4adef0955ac7e671821efbdc637e15bc9a3596248d9d2b02054449fc137e3a558fbd62a0bdf12b5fefbea40a5e3e7d762d4d79f
	cc88adc6994a4932810c0ec20d772f0fc1f761eee5019673f47bb6b1147b85806641d1ae2b72082fff9cbf028181008bd23bcef5b657173f0545edea
	e810635cdb2c54cc67cfaceed515afa503b3862163478386954465caa07f50e29ddc0f4af0d12856ff7d86a53d61318f4b7a9d674332f56e29656667
	04656f7b24525cf036b2052b601b230611ea2837424f3cc44d55543154f2d2d106ebc6964e7bd49e718f17152a3edce2eb75773399f68102818100a6
	89595d095001e610224e22a0075ed63edf74cc373fd93629eccdb9c92d82251244a0a63f844afb7f82d0fee966133d3c070ce7c96a1b4449e658208f
	abc6e97cdaa1e65be9e9b1447dbd346c2d5eaf3b19ee729ed363bfa490413e6f3c7de1df5b4313a69453ae11530c185ce40b1a0c2145b2c61ca10567
	a91abe84c82661028180204648dda561ff2f741f4e47838f39a30f126efc7b60bbac3d7b631f396c0b80cac027ff454aaec88903f6373b82e25ace4e
	942ef460761b414ba21ba2dbc9a948005e49e7529e04c8be1eaf45c3bfa87da7c61031d87c4392894978ede1aedbef69a103d88f6ccff9def809af9f
	1eaa3c0638d115432c714c2b3b33c8f9f028`

func TestImportedEC2Key(t *testing.T) {
	keyString := regexp.MustCompile(`\s+`).ReplaceAllString(ec2Key, "")
	keyBytes, err := hex.DecodeString(keyString)
	require.NoError(t, err)

	cred := virtualwebauthn.NewCredentialWithImportedKey(virtualwebauthn.KeyTypeEC2, keyBytes)
	testImportedCredential(t, cred)
}

func TestImportedRSAKey(t *testing.T) {
	keyString := regexp.MustCompile(`\s+`).ReplaceAllString(rsaKey, "")
	keyBytes, err := hex.DecodeString(keyString)
	require.NoError(t, err)

	cred := virtualwebauthn.NewCredentialWithImportedKey(virtualwebauthn.KeyTypeRSA, keyBytes)
	testImportedCredential(t, cred)
}

func testImportedCredential(t *testing.T, cred virtualwebauthn.Credential) {
	rp := virtualwebauthn.RelyingParty{Name: WebauthnDisplayName, ID: WebauthnDomain, Origin: WebauthnOrigin}
	authenticator := virtualwebauthn.NewAuthenticator()

	attestation := startWebauthnRegister(t)
	attestationOptions, err := virtualwebauthn.ParseAttestationOptions(attestation.Options)
	require.NoError(t, err)

	attestationResponse := virtualwebauthn.CreateAttestationResponse(rp, authenticator, cred, *attestationOptions)
	webauthnCredential := finishWebauthnRegister(t, attestation, attestationResponse)

	authenticator.Options.UserHandle = []byte(UserID)
	authenticator.AddCredential(cred)

	assertion := startWebauthnLogin(t, webauthnCredential, cred.ID)
	assertionOptions, err := virtualwebauthn.ParseAssertionOptions(assertion.Options)
	require.NoError(t, err)

	foundCredential := authenticator.FindAllowedCredential(*assertionOptions)
	require.NotNil(t, foundCredential)
	require.Equal(t, cred, *foundCredential)

	assertionResponse := virtualwebauthn.CreateAssertionResponse(rp, authenticator, cred, *assertionOptions)
	finishWebauthnLogin(t, assertion, assertionResponse)
}