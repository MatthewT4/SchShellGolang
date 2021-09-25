package service

import (
	"context"
	"fmt"
	"github.com/MatthewT4/SchShellGolang/internal/auth"
	"github.com/MatthewT4/SchShellGolang/internal/db/repository"
	"github.com/MatthewT4/SchShellGolang/internal/structions"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"math/rand"
	"strings"
	"time"
)

type SScreens interface {
	SGetImage(id string) (int, string)
	NewScreen() (ret_code int, timeout time.Time, screen_token, screen_code string, er error)
}

type ScreenService struct {
	Scr      repository.Screens
	TokenSer auth.TokenManager
}

func NewScreenService(db *mongo.Database) *ScreenService {
	return &ScreenService{
		Scr: repository.NewScreenRepo(db),
		TokenSer: func(SignKey string) *auth.Manager {
			res, err := auth.NewManager(SignKey)
			if err != nil {
				log.Fatal(err)
			}
			return res
		}(SignKey),
	}
}

//SGetImage return Result Code and result (or what error if error there is result)
func (s *ScreenService) SGetImage(id string) (int, string) {
	res, err := s.Scr.GetImageScreen(context.TODO(), id)
	if err != nil {
		return 400, err.Error()
	}
	return 200, res
}

func (s *ScreenService) NewScreen() (ret_code int, timeout time.Time, screen_token, screen_code string, er error) {
	token := GenerateToken(23)
	code := GenerateToken(8)
	screen := structions.Screen{
		Name:     "Screen",
		ScreenId: token,
		Code:     code,
		Data:     "",
	}
	fmt.Println(screen)
	_, err := s.Scr.AddScreen(context.TODO(), screen)
	if err != nil {
		return 400, time.Time{}, "", "", err
	}
	timer := time.Now().Add(screenJWT)
	tok, e := s.TokenSer.NewJWT(token, structions.ScreenR, screenJWT)
	if e != nil {
		return 400, timer, "", "", e
	}
	return 200, timer, tok, code, nil
}

func GenerateToken(passwordLength int) string {
	lowerCharSet := "abcdedfghijklmnopqrst"
	upperCharSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet := "!@#$%&*"
	numberSet := "0123456789"
	allCharSet := lowerCharSet + upperCharSet + specialCharSet + numberSet
	rand.Seed(time.Now().Unix())
	minSpecialChar := 1
	minNum := 1
	minUpperCase := 1

	var password strings.Builder

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}
