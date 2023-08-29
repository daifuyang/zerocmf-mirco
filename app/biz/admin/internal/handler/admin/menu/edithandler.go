package menu

import (
	"app/biz/admin/internal/logic/admin/menu"
	"app/biz/admin/internal/svc"
	"app/biz/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func EditHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MenuSaveReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := menu.NewEditLogic(r, svcCtx)
		result := l.Edit(&req)

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
