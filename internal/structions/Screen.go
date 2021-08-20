package structions

type Screen struct {
	Id       string `bson:"id"`
	Name     string `bson:"name"`
	Image    string `bson:"image"`
	Position string `bson:"position"`
}
