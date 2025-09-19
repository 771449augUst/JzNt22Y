// 代码生成时间: 2025-09-20 07:35:29
package main

import (
    "encoding/json"
    "net"
    "os/exec"
    "strings"
    "time"

    "github.com/revel/revel"
)

// NetworkStatusChecker struct
type NetworkStatusChecker struct {

}

// CheckConnection is a Revel action method to check network connection status.
func (c *NetworkStatusChecker) CheckConnection(host string) revel.Result {
    // Ping the host to check if it's reachable
    status := pingHost(host)

    // Create the response object with status and timestamp
    response := struct {
        Status string `json:"status"`
        Time   string `json:"time"`
    }{
        Status: status,
        Time:   time.Now().Format(time.RFC3339),
    }

    // Marshal the response object to JSON
    jsonResponse, _ := json.Marshal(response)

    // Return the JSON response as a Revel result
    return c.RenderJSON(jsonResponse)
}

// pingHost pings the given host and returns the status as a string.
func pingHost(host string) string {
    // Use the system's ping command to check the connectivity
    cmd := exec.Command("ping", "-c", "1", host)
    output, err := cmd.CombinedOutput()
    if err != nil || strings.Contains(string(output), "100% packet loss") {
        return "down"
    } else {
        return "up"
    }
}

func init() {
    // Register the Controller
    revel.RegisterController((*NetworkStatusChecker)(nil), nil)
}
