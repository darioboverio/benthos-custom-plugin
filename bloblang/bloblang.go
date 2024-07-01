package bloblang

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"github.com/redpanda-data/benthos/v4/public/bloblang"
	"gitlab.com/balance-inc/commons/crypt"
)

func LoadPublicKey(key string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(key))
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		return nil, errors.New("decode RSA public key failed")
	}
	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)

	if err != nil {
		return nil, err
	}

	return publicKey, nil
}

func init() {
	encryptRSASpec := bloblang.NewPluginSpec().
		Description("This is a mothod used to encrypt 'text' using RSA mechanism").
		Param(bloblang.NewStringParam("key").Description("Public key")).
		Param(bloblang.NewStringParam("text").Description("Text to encrypt"))

	err := bloblang.RegisterFunctionV2("encrypt_rsa", encryptRSASpec, func(args *bloblang.ParsedParams) (bloblang.Function, error) {
		return func() (interface{}, error) {
			// get parameters
			key, _ := args.GetString("key")
			text, _ := args.GetString("text")
			// creates PublicKey from key param
			// this will be replaced by ssl.LoadPublicKey(key)
			// from our Rain library
			publickey, err := LoadPublicKey(key)
			if err != nil {
				return nil, err
			}
			rsaPss := crypt.NewRSAPSSEncryption(
				nil, publickey, func(config *crypt.RSAConfig) {
					config.Strategy = crypt.RSAPKCS
				},
			)
			return rsaPss.Encrypt(text)
		}, nil
	})
	if err != nil {
		panic(err)
	}
}
