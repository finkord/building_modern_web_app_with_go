package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/finkord/building_modern_web_app_with_go/internal/config"
)

var app *config.AppConfig

// NewHelpers sets up the helper package
func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Printf("client error with status %d", status)
	http.Error(w, http.StatusText(status), status)

}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
