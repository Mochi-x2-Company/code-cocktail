package models

import ()

type Answer struct {
	Id     int    `json:id sql:AUTO_INCREMENT`
	Input  string `json: input`
	Output string `json: output`
}

// func AnsSingle() Answer {
// 	return Answer{}
// }
//
// func AnsAll() []Answer {
// 	return []Answer{}
// }
//
// func AnsMigrate() *Answer {
// 	return &Answer{}
// }
