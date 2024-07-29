// Copyright 2024 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package keys

import (
	"fmt"

	"golang.org/x/crypto/sha3"

	"github.com/coinbase/rosetta-sdk-go/asserter"
	"github.com/coinbase/rosetta-sdk-go/types"

	nemcrypto "github.com/NemProject/nem/gocrypto"
)

// SignerEdwards25519Keccak is initialized from a keypair
type SignerEdwards25519Keccak struct {
	KeyPair *KeyPair
}

var _ Signer = (*SignerEdwards25519Keccak)(nil)

// PublicKey returns the PublicKey of the signer
func (s *SignerEdwards25519Keccak) PublicKey() *types.PublicKey {
	return s.KeyPair.PublicKey
}

// Sign arbitrary payloads using a KeyPair
func (s *SignerEdwards25519Keccak) Sign(
	payload *types.SigningPayload,
	sigType types.SignatureType,
) (*types.Signature, error) {
	err := s.KeyPair.IsValid()
	if err != nil {
		return nil, fmt.Errorf("key pair is invalid: %w", err)
	}

	if !(payload.SignatureType == types.Ed25519_Keccak || payload.SignatureType == "") {
		return nil, fmt.Errorf(
			"expected signing payload signature type %v but got %v: %w",
			types.Ed25519_Keccak,
			payload.SignatureType,
			ErrSignUnsupportedPayloadSignatureType,
		)
	}

	if sigType != types.Ed25519_Keccak {
		return nil, fmt.Errorf(
			"expected signature type %v but got %v: %w",
			types.Ed25519_Keccak,
			sigType,
			ErrSignUnsupportedSignatureType,
		)
	}

	hasher := sha3.NewLegacyKeccak512()
	privKeyBytes := s.KeyPair.PrivateKey
	privKey := nemcrypto.NewKeyFromSeed(privKeyBytes, hasher)
	sig := nemcrypto.Sign(privKey, payload.Bytes, hasher)

	return &types.Signature{
		SigningPayload: payload,
		PublicKey:      s.KeyPair.PublicKey,
		SignatureType:  payload.SignatureType,
		Bytes:          sig,
	}, nil
}

// Verify verifies a Signature, by checking the validity of a Signature,
// the SigningPayload, and the PublicKey of the Signature.
func (s *SignerEdwards25519Keccak) Verify(signature *types.Signature) error {
	if signature.SignatureType != types.Ed25519_Keccak {
		return fmt.Errorf(
			"expected signing payload signature type %v but got %v: %w",
			types.Ed25519_Keccak,
			signature.SignatureType,
			ErrVerifyUnsupportedPayloadSignatureType,
		)
	}

	pubKey := signature.PublicKey.Bytes
	message := signature.SigningPayload.Bytes
	sig := signature.Bytes
	err := asserter.Signatures([]*types.Signature{signature})
	if err != nil {
		return fmt.Errorf("signature is invalid: %w", err)
	}

	hasher := sha3.NewLegacyKeccak512()
	verify := nemcrypto.Verify(pubKey, message, sig, hasher)
	if !verify {
		return ErrVerifyFailed
	}

	return nil
}
