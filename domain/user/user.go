package user

//go:generate go run ../../cmd/generator/main.go .

type User struct {
	id         string
	name       string
	age        int
	attributes map[string]interface{}
}
