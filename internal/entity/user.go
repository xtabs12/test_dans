package entity

type (
	User struct {
		ID       int64 `db`
		Name     string
		UserName string
		Password string
	}
)

func (e *User) TableName() string {
	return "user"
}
