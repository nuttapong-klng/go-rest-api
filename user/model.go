package user

type User struct {
	ID                string `json:"id" db:"id,type:uuid;primary_key;default;uuid_generate_v4()"`
	Email             string `json:"email" db:"email"`
	Password          string `json:"password" db:"password"`
	FirstName         string `json:"first_name" db:"first_name"`
	LastName          string `json:"last_name" db:"last_name"`
	IDCardFile        string `json:"id_card_file" db:"id_card_file"` //TODO: file upload handling
	IDCardNumber      string `json:"id_card_number" db:"id_card_number"`
	IDCardLaserNumber string `json:"id_card_laser_number" db:"id_card_laser_number"`
	DateOfBirth       string `json:"date_of_birth" db:"date_of_birth"`
	CreatedAt         string `json:"created_at" db:"created_at"`
	UpdatedAt         string `json:"updated_at" db:"updated_at"`
}

type LoginInfo struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
