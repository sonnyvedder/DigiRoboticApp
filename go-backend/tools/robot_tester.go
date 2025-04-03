package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

type RobotTester struct {
    BaseURL string
    Client  *http.Client
}

func NewRobotTester() *RobotTester {
    return &RobotTester{
        BaseURL: "http://localhost:8080",
        Client:  &http.Client{Timeout: 10 * time.Second},
    }

}

func (rt *RobotTester) SendCommand(command string, params map[string]interface{}) {
    payload := map[string]interface{}{
        "command": command,
        "params":  params,
    }

    jsonData, _ := json.Marshal(payload)
    resp, err := rt.Client.Post(
        rt.BaseURL+"/api/robot/command",
        "application/json",
        bytes.NewBuffer(jsonData),
    )

    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    var result map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&result)
    fmt.Printf("Response: %+v\n", result)
}

func main() {
    tester := NewRobotTester()

    // Test move command
    tester.SendCommand("move", map[string]interface{}{
        "position": map[string]interface{}{
            "x": 10.0,
            "y": 20.0,
            "z": 0.0,
        },
    })

    // Test rotation command
    tester.SendCommand("rotate", map[string]interface{}{
        "rotation": map[string]interface{}{
            "x": 0.0,
            "y": 0.0,
            "z": 45.0,
        },
    })
}