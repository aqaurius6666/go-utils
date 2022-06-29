package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"

	"github.com/tendermint/tendermint/crypto/secp256k1"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/xerrors"
)

var (
	ErrInvalidPrivateKey = xerrors.New("invalid length private key")
	ErrInvalidMessage    = xerrors.New("invalid message")
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

func GenerateKeyPair(sec []byte) ([]byte, []byte) {
	if sec == nil {
		priv := secp256k1.GenPrivKey()
		pubKey := priv.PubKey().Bytes()
		privKey := priv.Bytes()
		return privKey, pubKey
	}
	if len(sec) != 32 {
		return nil, nil
	}
	priv := secp256k1.PrivKey(sec)
	privKey := priv.Bytes()
	pubKey := priv.PubKey().Bytes()
	return privKey, pubKey
}

func ConvertMessage(message interface{}) (out []byte, err error) {
	switch tmessage := message.(type) {
	case []byte:
		out = tmessage
	case string:
		out = []byte(tmessage)
	case json.RawMessage:
		out = []byte(tmessage)
	default:
		out, err = json.Marshal(message)
		if err != nil {
			return nil, xerrors.Errorf("%w", ErrInvalidMessage)
		}
	}
	return
}

func SignMessage(message interface{}, privKey []byte) ([]byte, error) {
	fmt.Printf("len(privKey): %v\n", len(privKey))
	if len(privKey) != 32 {
		return nil, ErrInvalidPrivateKey
	}
	priv := secp256k1.PrivKey(privKey)
	bz, err := ConvertMessage(message)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	return priv.Sign(bz)
}

func VerifySignature(message interface{}, sig []byte, pubKey []byte) (bool, error) {
	pub := secp256k1.PubKey(pubKey)
	bz, err := ConvertMessage(message)
	if err != nil {
		return false, xerrors.Errorf("%w", err)
	}
	return pub.VerifySignature(bz, sig), nil
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

func AES(sec []byte, message []byte) []byte {
	hash := Hash256(sec)
	block, _ := aes.NewCipher(hash)
	out := make([]byte, 32)
	block.Encrypt(out, message)
	pbkdf2.Key([]byte("password"), []byte("asd"), 4096, 32, sha256.New)
	return out
}
