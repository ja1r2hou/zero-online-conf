package logic

import (
	"context"
	"zero-online-conf/api/internal/svc"
	"zero-online-conf/api/internal/types"
	"zero-online-conf/common/util"

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

	privateKey, publicKey, err := util.GenerateRSAKey(2048)
	if err != nil {
		return nil, err
	}

	resp := &types.GenerateRSAKeyResp{
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}

	return resp, nil
}
