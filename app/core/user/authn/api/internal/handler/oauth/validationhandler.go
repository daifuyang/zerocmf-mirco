package oauth

import (
	"app/core/user/authn/api/internal/logic/oauth"
	"app/core/user/authn/api/internal/svc"
	"app/core/user/authn/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func ValidationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ValidationReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := oauth.NewValidationLogic(r, svcCtx)
		result := l.Validation(&req)

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
