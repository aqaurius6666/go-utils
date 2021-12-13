package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func copyBytes(b [32]byte) []byte {
	h := make([]byte, 0)
	for _, e := range b {
		h = append(h, e)
	}
	return h
}

func Hash256(in []byte) []byte {
	out32 := sha256.Sum256([]byte(in))
	out := copyBytes(out32)
	return out
}

func BytesToBase64(in []byte) string {
	return base64.StdEncoding.EncodeToString(in)
}

func Base64ToBytes(in string) ([]byte, error) {
	b, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func GenerateKeyPair(privateKey []byte) (privkey []byte, pubkey []byte, err error) {
	if privateKey == nil {
		key, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
		if err != nil {
			return nil, nil, err
		}
		pubkey = elliptic.MarshalCompressed(secp256k1.S256(), key.X, key.Y)
		privkey = make([]byte, 32)
		blob := key.D.Bytes()
		copy(privkey[32-len(blob):], blob)

		return privkey, pubkey, nil
	}
	var e ecdsa.PrivateKey
	e.D = new(big.Int).SetBytes(privateKey)
	e.PublicKey.Curve = secp256k1.S256()
	e.PublicKey.X, e.PublicKey.Y = e.PublicKey.Curve.ScalarBaseMult(e.D.Bytes())
	return e.D.Bytes(), elliptic.MarshalCompressed(secp256k1.S256(), e.X, e.Y), nil
}

func SignMessage(msg interface{}, secKey []byte) ([]byte, error) {
	hash, err := ConvertMessage(msg)
	if err != nil {
		return nil, err
	}
	sig, err := secp256k1.Sign(hash, secKey)
	if err != nil {
		log.Fatal((err))
	}
	return sig[0:64], nil
}

func VerifySig(msg interface{}, sig []byte, pub []byte) (bool, error) {
	hash, err := ConvertMessage(msg)
	if err != nil {
		return false, err
	}
	ok := secp256k1.VerifySignature(pub, hash, sig)
	if !ok {
		return false, errors.New("verify signature fail")
	}

	return ok, nil

}

func ConvertMessage(message interface{}) ([]byte, error) {
	var bmsg []byte
	var err error
	switch message := message.(type) {
	case json.RawMessage:
		bmsg = message

	case string:
		bmsg = []byte(message)
	default:
		bmsg, err = json.Marshal(message)
		if err != nil {
			return nil, err
		}
	}
	hash := Hash256(bmsg)
	return hash, nil
}

func EncryptMessage(message []byte, secKey []byte) ([]byte, error) {
	block, _ := aes.NewCipher(Hash256(secKey))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, message, nil)
	return ciphertext, nil
}
func DecryptCipher(data []byte, secKey []byte) ([]byte, error) {
	key := Hash256(secKey)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
