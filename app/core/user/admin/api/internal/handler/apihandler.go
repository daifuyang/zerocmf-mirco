package handler

import (
	"app/core/user/admin/api/internal/logic"
	"app/core/user/admin/api/internal/svc"
	"app/core/user/admin/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func ApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewApiLogic(r, svcCtx)
		result := l.Api(&req)

		response := result.Response()

		if response.StatusCode != nil {
			w.WriteHeader(*response.StatusCode)
		}

		if response.IsString {
			w.Write([]byte(response.Msg))
			httpx.Ok(w)
			return
		}

		httpx.OkJsonCtx(r.Context(), w, response)
	}
}
