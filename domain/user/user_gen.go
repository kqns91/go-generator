// Code generated by generator; DO NOT EDIT.
package user

func NewUser(id string, name string, age int, company Company, attributes map[string]interface{}) *User {
	return &User{
		id:         id,
		name:       name,
		age:        age,
		company:    company,
		attributes: attributes,
	}
}

func (s *User) ID() string {
	return s.id
}

func (s *User) Name() string {
	return s.name
}

func (s *User) Age() int {
	return s.age
}

func (s *User) Company() Company {
	return s.company
}

func (s *User) Attributes() map[string]interface{} {
	return s.attributes
}
