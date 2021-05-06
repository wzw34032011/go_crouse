package server

import "log"

type s1 struct {
}

func (s *s1) Start() error {
	log.Println("")
	return nil
}

func (s *s1) Stop() error {
	return nil
}
