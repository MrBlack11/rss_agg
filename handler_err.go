package main

import (
	"net/http"

	"github.com/mrblack11/rss_agg/common"
)

func handlerErr(w http.ResponseWriter, r *http.Request) {
	common.RespondWithError(w, 400, "Something went wrong")
}
