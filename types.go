package apirest

type CreateUserParams struct {
	Name      string
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Age       int32
}

type Users []*User
type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone,omitempty"`
	Age       int32  `json:"age"`
}
