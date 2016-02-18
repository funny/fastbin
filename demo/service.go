// fb:service = 1
package main

import "github.com/funny/link"

// fb:handler
type MyService struct {
}

// fb:message = 1
type MyMessage1 struct {
	Field1 []byte
	Field2 []int
}

func (s *MyService) HandleMessage1(session link.FbSession, msg *MyMessage1) {

}

// fb:message = 2
type MyMessage2 struct {
	Field1 int
	Field2 string
}

func (s *MyService) HandleMessage2(session link.FbSession, msg *MyMessage2) {

}

// fb:message = 3
type MyMessage3 struct {
	Field1 int
	Field2 string
}
