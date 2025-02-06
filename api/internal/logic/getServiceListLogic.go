package logic

import (
	"context"
	"zero-online-conf/rpc/onlineConf"

	"zero-online-conf/api/internal/svc"
	"zero-online-conf/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetServiceListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewGetServiceListLogic 获取服务注册列表
func NewGetServiceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServiceListLogic {
	return &GetServiceListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetServiceListLogic) GetServiceList(req *types.GetServiceListReq) (*types.GetServiceListResp, error) {
	getServiceList, err := l.svcCtx.OnlineConf.GetServiceList(l.ctx, &onlineConf.GetServiceListReq{ServiceName: req.ServiceName})

	if err != nil {
		return nil, err
	}
	arrays := make([]*types.GetServiceListArrays, 0)
	for i := 0; i < len(getServiceList.GetServiceListArrays); i++ {
		getServiceListArray := getServiceList.GetServiceListArrays[i]

		arrays = append(arrays, &types.GetServiceListArrays{
			ServiceName: getServiceListArray.ServiceName,
			IpAddr:      getServiceListArray.IpAddr,
			Count:       int64(getServiceListArray.Count),
		})

	}

	return &types.GetServiceListResp{GetServiceListArrays: arrays}, nil
}
