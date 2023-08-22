package oauth

import (
	"app/core/user/authn/api/internal/logic/oauth"
	"app/core/user/authn/api/internal/svc"
	"net/http"
)

func TokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := oauth.NewTokenLogic(w, r, svcCtx)
		l.Token()
	}
}
