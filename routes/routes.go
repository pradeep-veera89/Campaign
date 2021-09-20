package routes

import (
	"net/http"

	"github.com/pradeep-veera89/campaign/internal/handler"
)

func Routes() {
	p := handler.Page{}
	http.HandleFunc("/", p.StartHandler)
	http.HandleFunc("/success", p.SuccessHandler)
	http.HandleFunc("/error", p.ErrorHandler)
}
