package api

type (
	CreateUserParams struct {
		Name      string
		FirstName string
		LastName  string
		Email     string
		Phone     string
		Age       int32
	}

	User struct {
		ID        int64  `db:"id"`
		Name      string `db:"name"`
		FirstName string `db:"first_name"`
		LastName  string `db:"last_name"`
		Email     string `db:"email"`
		Phone     string `db:"phone"`
		Age       int32  `db:"age"`
	}
)
