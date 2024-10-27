package global

import (
	"context"

	"example/be/service/api/internal/svc"
	"example/be/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTimeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取系统时间
func NewGetTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTimeLogic {
	return &GetTimeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTimeLogic) GetTime() (resp *types.TimeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
