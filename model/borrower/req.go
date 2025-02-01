package borrower

type (
	CreateReqBody struct {
		Name        string `json:"name"`
		Address     string `json:"address"`
		PhoneNumber string `json:"phone_number" binding:"required,len=12"`
		Email       string `json:"email" binding:"required,email"`
	}

	DeleteReqPath struct {
		Id int `binding:"required" uri:"id"`
	}

	DetailReqPath struct {
		Id int `binding:"required" uri:"id"`
	}

	ListReqQuery struct {
		Limit  int `form:"limit"`
		Offset int
		Page   int `form:"page"`
	}
)
