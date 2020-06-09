package ctrl

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Router returns the HTTP routes handler
func (c *Controller) Router() http.Handler {
	r := mux.NewRouter()

	// auth
	r.Methods(http.MethodGet).Path("/login").HandlerFunc(c.loginHandler)

	// game state
	r.Methods(http.MethodGet).Path("/game/{game_id}/snapshot").Handler(c.auth.Wrap(c.snapshotHandler))
	r.Methods(http.MethodGet).Path("/game/{game_id}/watch").Handler(c.auth.Wrap(c.wsHandler))

	return cors(r) // enable cors
}
