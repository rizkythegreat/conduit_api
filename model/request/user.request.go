package request

type UserCreateRequest struct {
	Name  string `json:"name"`  // required: true
	Email string `json:"email"` // required: true
}
