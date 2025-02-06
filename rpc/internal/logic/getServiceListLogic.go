package logic

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strings"
	"time"
	"zero-online-conf/rpc/internal/svc"
	"zero-online-conf/rpc/onlineConf"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetServiceListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetServiceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServiceListLogic {
	return &GetServiceListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetServiceListLogic) GetServiceList(in *onlineConf.GetServiceListReq) (*onlineConf.GetServiceListResp, error) {

	// 创建etcd客户端配置
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   l.svcCtx.Config.Etcd.Hosts, // 修改为你的etcd服务器地址和端口
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		l.Logger.Errorf("创建etcd客户端失败: %v", err)
		return nil, err
	}
	defer cli.Close()

	// 获取所有服务信息
	serviceArray := make([]*onlineConf.GetServiceListArrays, 0)

	for i := 0; i < len(in.GetServiceName()); i++ {
		etcdResp, etcdErr := cli.Get(context.Background(), in.GetServiceName()[i], clientv3.WithPrefix())
		if etcdErr != nil {
			l.Logger.Errorf("获取服务信息失败: %v", etcdErr)
			return nil, etcdErr
		}

		serviceName := ""
		ipAddrs := make([]string, 0)
		var count int32 = 0

		if etcdResp.Kvs == nil {
			serviceArray = append(serviceArray, &onlineConf.GetServiceListArrays{
				ServiceName: in.GetServiceName()[i],
				IpAddr:      "无活跃的机器",
				Count:       count,
			})
			continue
		}
		for _, ev := range etcdResp.Kvs {
			serviceName = in.GetServiceName()[i]
			ipAddrs = append(ipAddrs, string(ev.Value))
			count = count + 1
		}
		serviceArray = append(serviceArray, &onlineConf.GetServiceListArrays{
			ServiceName: serviceName,
			IpAddr:      strings.Join(ipAddrs, " ,"),
			Count:       count,
		})

	}

	return &onlineConf.GetServiceListResp{GetServiceListArrays: serviceArray}, nil
}
