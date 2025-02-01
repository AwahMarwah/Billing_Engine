package borrower

type (
	DetailResData struct {
		Name        string `json:"name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		CreatedAt   string `json:"created_at"`
		Address     string `json:"address"`
	}
	ListResData struct {
		Id          uint32 `json:"id"`
		Name        string `json:"name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		CreatedAt   string `json:"created_at"`
		Address     string `json:"address"`
	}
)
