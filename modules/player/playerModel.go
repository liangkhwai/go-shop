package player

type (
	PlayerProfile struct {
		Id        string `json:"id"`
		Email     string `json:"email"`
		Username  string `json:"username"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
	PlayerClaims struct {
		Id       string `json:"id"`
		RoleCode int    `json:"role_code"`
	}
	CreatePlayerReq struct {
		Email    string `json:"email" form:"email" validate:"required,email,max=255"`
		Username string `json:"username" form:"username" validate:"required,max=64"`
		Password string `json:"password" form:"password" validate:"required,max=32"`
	}

	CreatePlayerTransectionReq struct {
		PlayerId string  `json:"player_id" validate:"required,max=64"`
		Amount   float64 `json:"amount" validate:"required"`
	}
)
