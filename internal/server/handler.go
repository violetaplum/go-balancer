package server

import (
	"github.com/violetaplum/go-balancer/internal/proxy"
	"io"
	"net/http"
)

type Handler struct {
	lb *proxy.LoadBalancer
}

func NewHandler(lb *proxy.LoadBalancer) *Handler {
	return &Handler{lb: lb}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	bodySize := int32(req.ContentLength)
	node := h.lb.GetNextAvailableNode(bodySize)
	if node == nil {
		http.Error(w, "There's no available nodes..", http.StatusServiceUnavailable)
		return
	}

	proxyReq, err := http.NewRequest(req.Method, node.URL, req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	proxyReq.Header = req.Header

	client := &http.Client{}
	response, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	defer response.Body.Close()

	for name, values := range response.Header {
		w.Header()[name] = values
	}

	w.WriteHeader(response.StatusCode)
	_, err = io.Copy(w, response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
