package main

import "github.com/funny/link"

// fastbin:service = 1
type MyService struct {
}

// fastbin:message = 1
type MyMessage1 struct {
	Field1 []byte
	Field2 []int
}

func (s *MyService) HandleMessage1(session *link.Session, msg *MyMessage1) {

}

// fastbin:message = 2
type MyMessage2 struct {
	Field1 int
	Field2 string
}

func (s *MyService) HandleMessage2(session *link.Session, msg *MyMessage2) {

}
