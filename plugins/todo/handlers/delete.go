package handlers

import (
	"encoding/json"
	"net/http"
)

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	// an example API handler
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
