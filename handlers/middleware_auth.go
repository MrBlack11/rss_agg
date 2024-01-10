package handlers

import (
	"fmt"
	"net/http"

	"github.com/mrblack11/rss_agg/common"
	"github.com/mrblack11/rss_agg/internal/auth"
	"github.com/mrblack11/rss_agg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *ApiConfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			common.RespondWithError(w, 403, fmt.Sprintf("auth error: %v", err))
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			common.RespondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
