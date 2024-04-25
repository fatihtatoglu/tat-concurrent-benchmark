package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

// DockerStats struct
type DockerStats struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CPU         string `json:"CPUPerc"`
	Memory      string `json:"MemUsage"`
	MemoryLimit string
	NetInput    string `json:"NetIO"`
	NetOutput   string
}

func main() {
	// Promote a handler for /metrics endpoint
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		// Get Docker container stats
		dockerStats, err := getDockerStats()
		if err != nil {
			http.Error(w, "Error getting Docker stats", http.StatusInternalServerError)
			return
		}

		// Format Docker stats as Prometheus metrics
		for _, stat := range dockerStats {
			fmt.Fprintf(w, "docker_stats{id=\"%s\",name=\"%s\"} %s\n", stat.ID, stat.Name, stat.CPU)
			fmt.Fprintf(w, "docker_stats_memory{id=\"%s\",name=\"%s\"} %s\n", stat.ID, stat.Name, stat.Memory)
			fmt.Fprintf(w, "docker_stats_memory_limit{id=\"%s\",name=\"%s\"} %s\n", stat.ID, stat.Name, stat.MemoryLimit)
			fmt.Fprintf(w, "docker_stats_net_input{id=\"%s\",name=\"%s\"} %s\n", stat.ID, stat.Name, stat.NetInput)
			fmt.Fprintf(w, "docker_stats_net_output{id=\"%s\",name=\"%s\"} %s\n", stat.ID, stat.Name, stat.NetOutput)
		}
	})

	// Start HTTP server
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}

// getDockerStats function to get Docker container stats
func getDockerStats() ([]DockerStats, error) {
	// Execute "docker stats" command
	cmd := exec.Command("docker", "stats", "--no-stream", "--format", "{{json .}}")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// Parse Docker stats JSON
	var stats []DockerStats
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		var stat DockerStats
		if err := json.Unmarshal([]byte(line), &stat); err != nil {
			return nil, err
		}

		mem := strings.Split(stat.Memory, " / ")[0]
		memLimit := strings.Split(stat.Memory, " / ")[1]

		stat.Memory = mem
		stat.MemoryLimit = memLimit

		neti := strings.Split(stat.NetInput, " / ")[0]
		neto := strings.Split(stat.NetInput, " / ")[1]

		stat.NetInput = neti
		stat.NetOutput = neto

		stats = append(stats, stat)
	}

	return stats, nil
}
