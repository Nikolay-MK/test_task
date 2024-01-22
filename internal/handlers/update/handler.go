package update

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
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

// UpdateHandler обработчик для обновления данных
func (h *Handler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get(h.config.URL)
	if err != nil {
		h.sendResponse(w, false, "service unavailable", http.StatusServiceUnavailable)
		return
	}
	defer response.Body.Close()

	xmlData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		h.sendResponse(w, false, "error reading XML data", http.StatusInternalServerError)
		return
	}

	entries, err := h.parseXML(xmlData)
	if err != nil {
		h.sendResponse(w, false, "error parsing XML", http.StatusInternalServerError)
		return
	}

	err = h.repository.UpdateIndividuals(entries)
	if err != nil {
		h.sendResponse(w, false, fmt.Sprintf("error updating database: %s", err), http.StatusInternalServerError)
		return
	}

	h.sendResponse(w, true, "", http.StatusOK)
}

// ParseXML разбирает XML данные
func (h *Handler) parseXML(xmlData []byte) ([]selltech.SDNEntry, error) {
	var sdnList selltech.SDNList
	err := xml.Unmarshal(xmlData, &sdnList)
	if err != nil {
		return nil, err
	}
	return sdnList.Entries, nil
}

// sendResponse отправляет ответ согласно заданной структуре
func (h *Handler) sendResponse(w http.ResponseWriter, result bool, info string, code int) {
	responseData := map[string]any{
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
