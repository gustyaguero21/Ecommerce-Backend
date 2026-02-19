package user

type UserDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	DNI       string `json:"dni"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
}
