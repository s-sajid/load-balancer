package servers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sort"
	"sync"
)

type ServerList struct {
	ports []int
}

func (s *ServerList) Populate(numServers int) {
	if numServers > 5 {
		log.Fatal("Error: Number of Servers Exceeds Limit")
	}

	for i := 0; i < numServers; i++ {
		s.ports = append(s.ports, i)
	}
}

func (s *ServerList) Pop() int {
	port := s.ports[0]
	s.ports = s.ports[1:]
	return port
}

func RunServers(numServers int) {
	var serverList ServerList
	serverList.Populate(numServers)

	sort.Ints(serverList.ports)

	var wg sync.WaitGroup
	wg.Add(numServers)
	defer wg.Wait()

	for i := 0; i < numServers; i++ {
		go CreateServers(&serverList, &wg, i+1)
	}
}

func CreateServers(sl *ServerList, wg *sync.WaitGroup, serverNumber int) {
	r := http.NewServeMux()
	defer wg.Done()

	port := sl.Pop()

	server := http.Server{
		Addr:    fmt.Sprintf(":808%d", port),
		Handler: r,
	}

	r.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		fmt.Fprintf(rw, "Server %d", serverNumber)
	})

	r.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		w.Write([]byte("400: Server Shutdown"))
		server.Shutdown(context.Background())
	})

	fmt.Printf("Server %d is running on port:808%d\n", serverNumber, port)

	server.ListenAndServe()
}
