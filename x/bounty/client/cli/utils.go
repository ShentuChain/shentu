package cli

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"

	"github.com/tendermint/tendermint/libs/tempfile"
)

// SaveKey saves the given key to a file as json and panics on error.
func SaveKey(privKey *ecies.PrivateKey, dirPath string, creator string) {
	if dirPath == "" {
		panic("cannot save private key: filePath not set")
	}
	if len(creator) < 20 {
		panic("cannot save private key: creator address is too short")
	}

	decKeyBz := crypto.FromECDSA(privKey.ExportECDSA())
	//to create a unique file name
	allBytes := bytes.Join([][]byte{[]byte(creator), decKeyBz}, nil)
	hashBytes := sha256.Sum256(allBytes)
	filename := fmt.Sprintf("dec-key-%s-%x.json", creator[6:12], hashBytes[:3])
	if err := tempfile.WriteFileAtomic(filepath.Join(dirPath, filename), decKeyBz, 0666); err != nil {
		panic(err)
	}
}

// LoadPubKey loads the key at the given location by loading the stored private key and getting the public key part.
func LoadPubKey(filePath string) []byte {
	keyBytes, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	prvK, err := crypto.ToECDSA(keyBytes)
	if err != nil {
		panic(err)
	}

	return crypto.FromECDSAPub(&prvK.PublicKey)
}

// LoadPrvKey loads the key at the given location by loading the stored private key.
func LoadPrvKey(filePath string) *ecies.PrivateKey {
	keyBytes, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	prvK, err := crypto.ToECDSA(keyBytes)
	if err != nil {
		panic(err)
	}

	return ecies.ImportECDSA(prvK)
}
