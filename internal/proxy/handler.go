package proxy

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/google/uuid"
)

type ProxyRequest struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

type ProxyResponse struct {
	ID      string            `json:"id"`
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers"`
	Length  int               `json:"length"`
}

var requestsMap sync.Map


// ProxyHandler handles the proxy requests
// @Summary Proxy a request
// @Description Proxies a request to the specified URL and returns the response
// @Accept  json
// @Produce  json
// @Param   request body ProxyRequest true "Proxy Request"
// @Success 200 {object} ProxyResponse
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /proxy [post]
func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	var reqData ProxyRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &reqData)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := validateRequest(reqData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


	if reqData.Method == "" || reqData.URL == "" || reqData.Headers == nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(reqData.Method, reqData.URL, nil)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	for key, value := range reqData.Headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error proxying request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	requestID := uuid.New().String()

	requestsMap.Store(requestID, ProxyResponse{
		ID:      requestID,
		Status:  resp.StatusCode,
		Headers: mapToJSONMap(resp.Header),
		Length:  len(respBody),
	})

	jsonResp, err := json.Marshal(ProxyResponse{
		ID:      requestID,
		Status:  resp.StatusCode,
		Headers: mapToJSONMap(resp.Header),
		Length:  len(respBody),
	})
	if err != nil {
		http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func validateRequest(req ProxyRequest) error {
	if req.Method == "" {
		return fmt.Errorf("missing field: method")
	}
	if req.URL == "" {
		return fmt.Errorf("missing field: url")
	}
	if req.Headers == nil {
		return fmt.Errorf("missing field: headers")
	}
	return nil
}

func mapToJSONMap(headers map[string][]string) map[string]string {
	jsonMap := make(map[string]string)
	for key, values := range headers {
		jsonMap[key] = values[0]
	}
	return jsonMap
}