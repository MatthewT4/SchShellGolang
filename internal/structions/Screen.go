package structions

type Screen struct {
	ScreenId string `bson:"screen_id"`
	Name     string `bson:"name"`
	Data     string `bson:"data"`
	Code     string `bson:"code"`
}
