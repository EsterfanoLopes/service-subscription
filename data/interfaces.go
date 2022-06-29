package data

type UserInterface interface {
	GetAll() ([]*User, error)
	GetByEmail(email string) (*User, error)
	GetOne(ID int) (*User, error)
	Update() error
	Delete() error
	DeleteByID(ID int) error
	Insert(user User) (int, error)
	ResetPassword(password string) error
	PasswordMatches(plainText string) (bool, error)
}

type PlanInterface interface {
	GetAll() ([]Plan, error)
	GetOne(ID int) (*Plan, error)
	SubscribeUserToPlan(user User, plan Plan) error
	AmountForDisplay() string
}
