package middelware

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
)

//Data that you will use to obtain the token
type DateValidate struct {
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
}

type Token struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserName  string             `bson:"userName" json:"userName"`
	FirstName string             `bson:"firstName" json:"firstName"`
	Email     string             `bson:"email" json:"email"`
	Website   string             `bson:"website" json:"website"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
	//Birthday  time.Time          `bson:"birthday" json:"birthday"`
}

//The token body
type Claim struct {
	Token `json:"user"`
	//standar claim
	jwt.StandardClaims
}

//Token response
type Responsetoken struct {
	Token string `json:"token"`
}

//using method RS256
func init() {
	//the read archive in format bytes for save  the keys private anad public
	privateBytes, err := ioutil.ReadFile("middelware/private.rsa")
	if err != nil {
		log.Fatal(err)
	}

	publicBytes, err := ioutil.ReadFile("middelware/public.rsa.pub")
	if err != nil {
		log.Fatal(err)
	}

	//for load in the form of a key private and public
	PrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("could not do the parse of private")
	}
	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("could not do the parse of public")
	}
}

//using method HS256
func Keys() []byte {
	privateBytes, err := ioutil.ReadFile("middelware/private.rsa")
	if err != nil {
		log.Fatal("private key was not read")
	}
	return privateBytes
}

/*
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)
	//path := filepath.Join(dir,, "hola.txt")
*/
