package global

import (
	"net/http"

	"example/be/service/api/internal/logic/global"
	"example/be/service/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取系统时间
func GetTimeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := global.NewGetTimeLogic(r.Context(), svcCtx)
		resp, err := l.GetTime()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
