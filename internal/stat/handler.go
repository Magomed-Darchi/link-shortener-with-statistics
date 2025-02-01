package stat

import (
	"api-main/configs"
	"api-main/pkg/midlleware"
	"fmt"
	"net/http"
	"time"
)

const (
	FilterbyDay   = "day"
	FilterbyMonth = "month"
)

type StatkHandler struct {
	StatRepository *StatRepository
}

type StatHandlerDeps struct {
	StatRepository *StatRepository
	Config         *configs.Config
}

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	handler := &StatkHandler{
		StatRepository: deps.StatRepository,
	}

	router.Handle("GET /stat", midlleware.IsAuthed(handler.GetStat(), deps.Config))

}

func (h *StatkHandler) GetStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		from, err := time.Parse("2006-01-02", r.URL.Query().Get("from"))
		if err != nil {
			http.Error(w, "Invalid from params", http.StatusBadRequest)
			return
		}
		to, err := time.Parse("2006-01-02", r.URL.Query().Get("to"))
		if err != nil {
			http.Error(w, "Invalid to params", http.StatusBadRequest)
			return
		}
		by := r.URL.Query().Get("by")
		if by != FilterbyDay && by != FilterbyMonth {
			http.Error(w, "Invalid by params", http.StatusBadRequest)
			return
		}
		fmt.Println(from, to, by)
	}

}
