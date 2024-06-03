package models

type TspContinuity struct {
	Index   int64  `json:"index"`
	Packets int    `json:"packets"`
	Pid     int    `json:"pid"`
	Type    string `json:"type"`
}
