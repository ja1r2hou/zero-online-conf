package logic

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"zero-online-conf/api/internal/svc"
	"zero-online-conf/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateRSAKeyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewGenerateRSAKeyLogic 获取新的公钥和私钥
func NewGenerateRSAKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateRSAKeyLogic {
	return &GenerateRSAKeyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateRSAKeyLogic) GenerateRSAKey() (*types.GenerateRSAKeyResp, error) {

	privateKey, publicKey, err := GenerateRSAKey(2048)
	if err != nil {
		return nil, err
	}

	resp := &types.GenerateRSAKeyResp{
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}

	return resp, nil
}
func GenerateRSAKey(bits int) (string, string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "", "", err
	}

	publicKey := &privateKey.PublicKey

	// 将私钥编码为 PEM 格式
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	privatePEM := pem.EncodeToMemory(privateBlock)

	// 将公钥编码为 PEM 格式
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", "", err
	}
	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	publicPEM := pem.EncodeToMemory(publicBlock)
	return string(privatePEM), string(publicPEM), nil
}
