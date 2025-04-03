package models

type Robot struct {
    ID       string  `json:"robot_id"`
    Status   string  `json:"status"`
    Position Position `json:"position"`
}

type Position struct {
    X float64 `json:"x"`
    Y float64 `json:"y"`
    Z float64 `json:"z"`
}

type RobotCommand struct {
    Command string                 `json:"command"`
    Params  map[string]interface{} `json:"params"`
}