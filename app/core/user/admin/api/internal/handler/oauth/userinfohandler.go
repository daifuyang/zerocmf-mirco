package oauth

import (
	"app/core/user/admin/api/internal/logic/oauth"
	"app/core/user/admin/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := oauth.NewUserInfoLogic(r, svcCtx)
		result := l.UserInfo()

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
