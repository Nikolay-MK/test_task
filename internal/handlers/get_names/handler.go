package get_names

import (
	"encoding/json"
	"net/http"
	"selltech/internal"
	"selltech/internal/selltech"
)

// Handler структура для обработчиков HTTP
type Handler struct {
	config     internal.AppConfig
	repository repository
}

// New создает новый экземпляр Handler
func New(config internal.AppConfig, repository repository) *Handler {
	return &Handler{
		config:     config,
		repository: repository,
	}
}

func (h *Handler) GetNames(w http.ResponseWriter, r *http.Request) {
	// Получаем параметры из запроса
	name := r.FormValue("name")
	nameType := r.FormValue("type")

	var names []selltech.NameResult
	switch nameType {
	case "strong":
		names = h.repository.GetStrongNames(name)
	case "weak":
		names = h.repository.GetWeakNames(name)
	default:
		names = h.repository.GetAllNames(name)
	}

	sendJSONResponse(w, names, http.StatusOK)
}

// sendJSONResponse отправляет ответ в формате JSON
func sendJSONResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
