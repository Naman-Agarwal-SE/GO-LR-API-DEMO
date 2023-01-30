package request

type RegisterRequest struct {
	Email      []map[string]string `json:"Email"`
	Password   string              `validate:"required" json:"Password"`
	About      string              `json:"About"`
	FirstName  string              `validate:"required" json:"FirstName"`
	MiddleName string              `json:"MiddleName"`
	LastName   string              `json:"LastName"`
	BirthDate  string              `json:"BirthDate"`
	Company    string              `json:"Company"`
	Gender     string              `json:"Gender"`
}

type LoginRequest struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}

type ProfileUpdateRequest struct {
	Email      []map[string]string `json:"Email"`
	About      string              `json:"About"`
	FirstName  string              `json:"FirstName"`
	MiddleName string              `json:"MiddleName"`
	LastName   string              `json:"LastName"`
	BirthDate  string              `json:"BirthDate"`
	Company    string              `json:"Company"`
	Gender     string              `json:"Gender"`
}
