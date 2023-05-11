package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	http.HandleFunc("/ocr", handleOCR)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleOCR(w http.ResponseWriter, r *http.Request) {
	// Get the image from the multipart form
	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error retrieving image from form data", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Save the image to a temporary file
	tempFile, err := createTempFile(handler.Filename)
	if err != nil {
		http.Error(w, "Error creating temporary file for image", http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()
	_, err = io.Copy(tempFile, file)
	if err != nil {
		http.Error(w, "Error saving image to temporary file", http.StatusInternalServerError)
		return
	}
	defer os.Remove(tempFile.Name())
	fmt.Println(tempFile.Name())

	// Run Tesseract on the temporary file
	cmd := exec.Command("tesseract", tempFile.Name(), "stdout")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		http.Error(w, "Error running Tesseract OCR", http.StatusInternalServerError)
		return
	}

	// Write the OCR text to the response
	type OCRResponse struct {
		Text string `json:"text"`
	}
	response := OCRResponse{Text: out.String()}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
	}
}

func createTempFile(filename string) (*os.File, error) {
	ext := filepath.Ext(filename)
	prefix := filename[:len(filename)-len(ext)]
	tempDir := filepath.Join(".", "screenshots")
	if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
		return nil, err
	}
	return ioutil.TempFile(tempDir, prefix+"_*"+ext)
}
