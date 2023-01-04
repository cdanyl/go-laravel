package handlers

import (
	"github.com/cdanyl/celeritas"
	"net/http"
)

type Handlers struct {
	App *celeritas.Celeritas
}

func (h *Handlers) HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.Page(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering: ", err)
	}
}
