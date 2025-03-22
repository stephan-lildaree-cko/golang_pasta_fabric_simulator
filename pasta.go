package main

import (
	"fmt"
	"net/http"
	"sync"
)

type MeatballMonster struct {
	Name   string
	Energy int
	mu     sync.Mutex
}

func (m *MeatballMonster) Work() string {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.Energy > 0 {
		m.Energy -= 10
		return fmt.Sprintf("%s made some pasta!", m.Name)
	} else {
		return fmt.Sprintf("%s is too tired to work!", m.Name)
	}
}

func (m *MeatballMonster) Rest() string {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.Energy = 100
	return fmt.Sprintf("%s took a nap and is ready to work!", m.Name)
}

var monsters = []*MeatballMonster{
	{Name: "Meaty", Energy: 100},
	{Name: "Saucy", Energy: 100},
	{Name: "Spaghetter", Energy: 100},
}

func producePastaHandler(w http.ResponseWriter, r *http.Request) {
	var responses []string
	for _, monster := range monsters {
		responses = append(responses, monster.Work())
	}
	fmt.Fprintln(w, responses)
}

func restMonstersHandler(w http.ResponseWriter, r *http.Request) {
	var responses []string
	for _, monster := range monsters {
		responses = append(responses, monster.Rest())
	}
	fmt.Fprintln(w, responses)
}

func main() {
	http.HandleFunc("/produce", producePastaHandler)
	http.HandleFunc("/rest", restMonstersHandler)

	fmt.Println("Pasta Fabric Service is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
