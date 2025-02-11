package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

func rsaEncrypt(plaintext, publicKeyStr string) (string, error) {
	publicKey := []byte(publicKeyStr)
	block, _ := pem.Decode(publicKey)
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	pub := pubInterface.(*rsa.PublicKey)
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(plaintext))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func rsaDecrypt(ciphertext, privateKeyStr string) (string, error) {
	privateKey := []byte(privateKeyStr)
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	block, _ := pem.Decode(privateKey)
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertextBytes)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

func main() {
	encrypt, _ := rsaEncrypt("njsnjnjdnjndndn", "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAriP29/ufdgJOGqcDuax0\nT08PXRsiFVDO5DXmlJ3CmXHID6R3hRzTah4+ou4XiB01A+yzxkZZsAsKyQArC/WF\nBqAfkuq+HZ+B1gxIetMMzm7cc+Sq+N+fX918oXDIANFOvku+q9PUV6BbGoTzrwkH\nX/TfSezNYqZE0z8+2k8bZm00CVUQh0OYeqLPu42+hWrt5lR36WCEsVJgl3iB2WnO\nhFxhrx+Hf59NadswC3CyTPWfRiLiVH9to5P9Me2vj6rI4H4AtY6mv+tu/h6R+j4v\nlucV/DW6gc7I1F08KbRJJMi4XOZ5NTqna8BunrI1q2MqohnoV30+cBmYWCwE4raY\nIQIDAQAB\n-----END PUBLIC KEY-----")
	fmt.Println("加密")
	fmt.Println(encrypt)
	decrypt, _ := rsaDecrypt(encrypt, "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAriP29/ufdgJOGqcDuax0T08PXRsiFVDO5DXmlJ3CmXHID6R3\nhRzTah4+ou4XiB01A+yzxkZZsAsKyQArC/WFBqAfkuq+HZ+B1gxIetMMzm7cc+Sq\n+N+fX918oXDIANFOvku+q9PUV6BbGoTzrwkHX/TfSezNYqZE0z8+2k8bZm00CVUQ\nh0OYeqLPu42+hWrt5lR36WCEsVJgl3iB2WnOhFxhrx+Hf59NadswC3CyTPWfRiLi\nVH9to5P9Me2vj6rI4H4AtY6mv+tu/h6R+j4vlucV/DW6gc7I1F08KbRJJMi4XOZ5\nNTqna8BunrI1q2MqohnoV30+cBmYWCwE4raYIQIDAQABAoIBAELqIPlzX/f5tMd7\nC+xp/xuGlrHBPlyQe5+nsp3C7UcHOMgB+8dTp01sp8b11GcCSh/i8cWrvMTvyUop\nXWEwC9ja6KcMutcpNrvoZmWD+bTNVrrV0HjnfRdaRYzHiKL/ytFGy9K9CChPvxI2\nqmjFAGd38eE3NEjG+UVvVRiQM0gs5gB00viHQXP1w4A09Iku3upb2M0siC+HG1zl\nLgqtqNANdkD5HyZzUNlPTVsGsY1iwhooj80Mkzi3eR/cqYE0QCK+cfaACfmmEuFD\nxWuOgaCF5G5IW1UgcTnv7ga2J8MMJjCWmkIN+sVBqN8xE4vFjsaY2PBRJ5jcmTKC\nWiAa9UECgYEA1B6riND15OR9AywIMTTlONCqSfdwT747Hl7igqjTpU1AKxCSbXdm\n63ndjdqjSmVY1Bez9xkbXAUYCpHelZ4z89VzIV2FXZetxCWD63tUF1aqO2eK9juf\nd4gAz9wBfd6fNK9fRG5wvSbZTdhIUiUkbk+/hoL/yuP+XpfqItk0JNkCgYEA0ioB\nenutIEZYeEctwSMmLiokAy2gk30qxH429ZrmvRY7XiYdmrEwGtAJ7VoF+LPbNzEM\nm557MVpjVAx/Wy/0XRQ/IGTgGaLgx22lNJjgypXl2Ncz9yNJ5aF0oLDAA8CEHiq9\nY+r3NF+3HVnLeyflrfTFwXyOh8lsVhI8u5N54IkCgYEAtDt/GpgaZqsyAXD3YQAF\nyDmch46kMk17o/mRUt2qP2kdHbpOAI8Utv3sH2znk/369PNS6aC+m3iEje6VhuFO\nWV9DHNV+zSBk+CW8Kmi28cGkkScQbtoITCMWNYdFCTMQaU6djuKcDkwlFbyw7oM6\nqjQ+k0dZmoYQi0VfHs8ZPDkCgYBLSg0ZifDtnQXYoPc9N5BX/XTFRrU94RtPJUAf\nII0EtFA0XEdCwbNQB7NmuldlT1l7HjE8FrxfY9gtpgSB8F4EclpjCoBV7snD5/3F\nJ/dv299pnT7ajGPOxdZ4Mpm3PmKWRA1xHB3PP3LkKuUAi7x2YzftJugQsRIEi6gC\nMsT4sQKBgQCdJy6GIyQkK2BcqU5jFawQLvjB4kkDj+r7QqZibQmaY6NtqzQFJWqc\nyc5tn7ZPNMV37jCGVF61nut0ufg8DV18oCSL+EkkEBKAsCEibaG3DlJ1N6tPMjAg\n+eOfO1SxxKlEK1MT+4iZhssKA7gG34kzeqHO0lVDoH9FieyawYMkOg==\n-----END RSA PRIVATE KEY-----")

	fmt.Println("解密")

	fmt.Println(decrypt)
}
