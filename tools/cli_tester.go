package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    tester := NewRobotTester()
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Print("Enter command (move/rotate/status/quit): ")
        command, _ := reader.ReadString('\n')
        command = strings.TrimSpace(command)

        if command == "quit" {
            break
        }

        switch command {
        case "move":
            tester.SendCommand("move", map[string]interface{}{
                "position": map[string]interface{}{
                    "x": 10.0,
                    "y": 20.0,
                    "z": 0.0,
                },
            })
        case "status":
            // Get robot status
            resp, err := tester.Client.Get(tester.BaseURL + "/api/robot/state/robot1")
            if err != nil {
                fmt.Printf("Error: %v\n", err)
                continue
            }
            var result map[string]interface{}
            json.NewDecoder(resp.Body).Decode(&result)
            fmt.Printf("Robot Status: %+v\n", result)
        }
    }
}