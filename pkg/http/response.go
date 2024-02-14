package http

import (
	"encoding/json"
	"net/http"
)

// JSON yanıtı göndermek için yardımcı fonksiyon.
func SendJSON(w http.ResponseWriter, v interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		// JSON kodlama hatası durumunda, HTTP durum kodunu InternalServerError olarak güncelleyin
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "JSON kodlama hatası"})
	}
}

// Başarılı bir yanıt göndermek için yardımcı fonksiyon.
func SendSuccess(w http.ResponseWriter, v interface{}) {
	SendJSON(w, v, http.StatusOK)
}

// Hata yanıtı göndermek için yardımcı fonksiyon.
func SendError(w http.ResponseWriter, message string, statusCode int) {
	SendJSON(w, map[string]string{"error": message}, statusCode)
}

// NotFound yanıtı göndermek için yardımcı fonksiyon.
func SendNotFound(w http.ResponseWriter) {
	SendError(w, "Kaynak bulunamadı", http.StatusNotFound)
}

// BadRequest yanıtı göndermek için yardımcı fonksiyon.
func SendBadRequest(w http.ResponseWriter, message string) {
	SendError(w, message, http.StatusBadRequest)
}
