package structions

type RoleType int

const (
	ScreenR = -1
	UserR   = iota
	Moderation
	Administration
)

type User struct {
	Login      string   `bson:"login,omitempty"`
	Password   string   `bson:"password,omitempty"`
	Role       RoleType `bson:"group,omitempty"`
	Email      string   `bson:"email,omitempty"`
	Catalogues []string `bson:"catalogues,omitempty"`
}

func NewUser(login string, password string, role RoleType, email string, catalogues []string) *User {
	return &User{
		Login:      login,
		Password:   password,
		Role:       role,
		Email:      email,
		Catalogues: catalogues,
	}
}
