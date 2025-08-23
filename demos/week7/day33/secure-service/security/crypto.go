package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io"
)

var (
	ErrInvalidKey  = errors.New("invalid key")
	ErrInvalidData = errors.New("invalid data")
	ErrEncryption  = errors.New("encryption failed")
	ErrDecryption  = errors.New("decryption failed")
)

// Crypto handles cryptographic operations
type Crypto struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

// NewCrypto creates a new cryptography handler
func NewCrypto() (*Crypto, error) {
	// Generate RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	return &Crypto{
		privateKey: privateKey,
		publicKey:  &privateKey.PublicKey,
	}, nil
}

// EncryptSymmetric performs AES-GCM encryption
func (c *Crypto) EncryptSymmetric(key, plaintext []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, ErrInvalidKey
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

// DecryptSymmetric performs AES-GCM decryption
func (c *Crypto) DecryptSymmetric(key, ciphertext []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, ErrInvalidKey
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, ErrInvalidData
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, ErrDecryption
	}

	return plaintext, nil
}

// EncryptAsymmetric performs RSA encryption
func (c *Crypto) EncryptAsymmetric(plaintext []byte) ([]byte, error) {
	ciphertext, err := rsa.EncryptPKCS1v15(
		rand.Reader,
		c.publicKey,
		plaintext,
	)
	if err != nil {
		return nil, ErrEncryption
	}
	return ciphertext, nil
}

// DecryptAsymmetric performs RSA decryption
func (c *Crypto) DecryptAsymmetric(ciphertext []byte) ([]byte, error) {
	plaintext, err := rsa.DecryptPKCS1v15(
		rand.Reader,
		c.privateKey,
		ciphertext,
	)
	if err != nil {
		return nil, ErrDecryption
	}
	return plaintext, nil
}

// ExportPublicKey exports the public key in PEM format
func (c *Crypto) ExportPublicKey() string {
	pubKeyBytes := x509.MarshalPKCS1PublicKey(c.publicKey)
	pubKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubKeyBytes,
	})
	return string(pubKeyPEM)
}

// ExportPrivateKey exports the private key in PEM format
func (c *Crypto) ExportPrivateKey() string {
	privKeyBytes := x509.MarshalPKCS1PrivateKey(c.privateKey)
	privKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privKeyBytes,
	})
	return string(privKeyPEM)
}

// GenerateRandomKey generates a random symmetric key
func GenerateRandomKey() ([]byte, error) {
	key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return nil, err
	}
	return key, nil
}

// EncodeBase64 encodes data to base64
func EncodeBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// DecodeBase64 decodes base64 data
func DecodeBase64(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}
