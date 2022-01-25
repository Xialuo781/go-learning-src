package hello

import "fmt"

type HelloService struct{}

func (h *HelloService) SayHello(username string) string {
	return fmt.Sprintf("hello, %s", username)
}

type GreeterService struct{}

func (h *GreeterService) SayHello(username string) string {
	return fmt.Sprintf("hello, %s, how are you!", username)
}
