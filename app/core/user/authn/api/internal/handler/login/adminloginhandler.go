package login

import (
	"app/core/user/authn/api/internal/logic/login"
	"app/core/user/authn/api/internal/svc"
	"app/core/user/authn/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func AdminLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := login.NewAdminLoginLogic(r, svcCtx)
		result := l.AdminLogin(&req)

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
