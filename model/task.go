package model

type Task struct {
	InfoHash   string `json:"info_hash"`
	Complete   int    `json:"complete"`
	Downloaded int    `json:"downloaded"`
	Incomplete int64  `json:"incomplete"`
	Name       string `json:"name"`
}
