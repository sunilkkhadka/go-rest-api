package main

type APIServer struct {
	listenAddress string
}

func NewAPIServer(listenAddress string) *APIServer{
	server := APIServer{
		listenAddress: listenAddress,
	}

	return &server
}