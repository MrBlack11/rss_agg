package handlers

import (
	"net/http"

	"github.com/mrblack11/rss_agg/common"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	common.RespondWithJson(w, 200, struct{}{})
}
