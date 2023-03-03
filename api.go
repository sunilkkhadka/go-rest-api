package main

import "net/http"

type APIServer struct {
	listenAddress string
}

func NewAPIServer(listenAddress string) *APIServer{
	server := APIServer{
		listenAddress: listenAddress,
	}

	return &server
}

func (s *APIServer) Run() {

}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error{
	return nil
}

func (s *APIServer) handleetAccount(w http.ResponseWriter, r *http.Request) error{
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error{
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error{
	return nil
}