package main

import (
	"fmt"
	"net/http"
	"php-analyzer-web/analyzer"
)

func main() {
	// Configurar rutas
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/analyze", handleAnalyze)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Iniciar servidor
	fmt.Println("Servidor iniciado en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func handleAnalyze(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	code := r.FormValue("code")
	if code == "" {
		http.Error(w, "Código PHP no proporcionado", http.StatusBadRequest)
		return
	}

	// Realizar análisis
	result := analyzer.AnalyzePHP(code)

	// Devolver resultados como JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(result.ToJSON()))
}
