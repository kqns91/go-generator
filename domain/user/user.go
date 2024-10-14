package user

//go:generate go run ../../cmd/generator/main.go .

type User struct {
	id         string
	name       string
	age        int
	company    Company
	attributes map[string]interface{}
}

type Company struct {
	id string
}
