package cryptography

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenKeyPair(t *testing.T) {
	priv, pub, err := GenerateKeyPair(nil)
	assert.Nil(t, err)
	assert.Equal(t, len(priv), 32, "invalid length privatekey")
	assert.Equal(t, len(pub), 33, "invalid length publickey")

	sample_priv, err := Base64ToBytes("54ZzZznyetDwgPI593tefWxll51RKtd/twbsgW3aTYU=")
	assert.Nil(t, err)
	sample_pub, err := Base64ToBytes("AptMqF53OFvU7sWCEcdCu/F9XykczUMJkM1fQQCXgkon")
	assert.Nil(t, err)

	priv, pub, err = GenerateKeyPair(sample_priv)
	assert.Nil(t, err)
	assert.Equal(t, len(priv), 32, "invalid length privatekey")
	assert.Equal(t, priv, sample_priv, "invalid privatekey")
	assert.Equal(t, len(pub), 33, "invalid length publickey")
	assert.Equal(t, pub, sample_pub, "invalid publickey")
}

func TestSignMessage(t *testing.T) {

	body := map[string]interface{}{
		"a": "1231413123",
	}

	sample_priv, err := Base64ToBytes("GmQE4ZljJ5PCXev2dRPCW2JHVefgsTM6+96CmqJjb0w=")
	assert.Nil(t, err)
	priv, _, err := GenerateKeyPair(sample_priv)
	assert.Nil(t, err)
	sig, err := SignMessage(body, priv)
	assert.Nil(t, err)
	assert.Equal(t, len(sig), 64, "invalid length signature")

	sample_sig, err := Base64ToBytes("OJ71Yo3e3Pblrgjy4EEYzac75pJCa79VqpZKisW05dMLScWneUmA5SQyeAt0nrfkxL1+bIQfgeX7+CgRvUNiTg==")
	assert.Nil(t, err)
	assert.Equal(t, sig, sample_sig, "invalid sig")
	t.Log(sig)

}

func TestVerify(t *testing.T) {
	body := `{"a":1}`
	priv, pub, err := GenerateKeyPair(nil)
	assert.Nil(t, err)
	sig, err := SignMessage(body, priv)
	assert.Nil(t, err)
	ok, err := VerifySig(body, sig, pub)
	assert.Nil(t, err)
	assert.Equal(t, ok, true, "sig fail")
}

func TestConvertMessage(t *testing.T) {
	body := json.RawMessage(`{"a":"1231413123"}`)
	hash, err := ConvertMessage(body)
	assert.Nil(t, err)
	t.Error(hash)
}
