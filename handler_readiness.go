package main

import (
	"net/http"

	"github.com/mrblack11/rss_agg/common"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	common.RespondWithJson(w, 200, struct{}{})
}
