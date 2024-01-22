package state

import (
	"encoding/json"
	"net/http"
	"selltech/internal"
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

// GetState обрабатывает запрос на получение текущего состояния данных
func (h *Handler) GetState(w http.ResponseWriter, r *http.Request) {
	if internal.IsUpdating {
		h.SendResponse(w, false, "updating", http.StatusOK)
		return
	}

	if h.repository.IsDataReady() {
		h.SendResponse(w, true, "ok", http.StatusOK)
		return
	}
	h.SendResponse(w, false, "empty", http.StatusOK)

}
func (h *Handler) IsDataReady() bool {
	return h.repository.IsDataReady()
}

// SendResponse отправляет ответ
func (h *Handler) SendResponse(w http.ResponseWriter, result bool, info string, code int) {
	responseData := map[string]interface{}{
		"result": result,
		"info":   info,
		"code":   code,
	}
	sendJSONResponse(w, responseData, code)
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
