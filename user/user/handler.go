package user

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/labstack/echo"
	"github.com/orderforme/user/middelware"
	"github.com/orderforme/user/user/database"
	"github.com/orderforme/user/user/errors"
	"github.com/orderforme/user/user/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User interface {
	Done()
	GetUserByID(context echo.Context) error
	GetUsers(context echo.Context) error
	CreateUser(context echo.Context) error
	UpdateUser(context echo.Context) error
	DeleteUser(context echo.Context) error
	UploadAvatar(context echo.Context) error
	UploadFront(context echo.Context) error
	LoginUser(context echo.Context) error
	ValidateToken(context echo.Context) error
}

type Handler struct {
	userRepo database.MongoDB
}

// NewHandler allocates a new Handler
func NewHandler() (*Handler, error) {

	handler := Handler{
		userRepo: &database.Mongodb{},
	}

	err := handler.userRepo.ConnectDB()

	return &handler, err
}

// DisconnectDB all
func (handler *Handler) Done() {
	handler.userRepo.DisconnectDB()
}

// GetUserByID method
func (handler *Handler) GetUserByID(context echo.Context) error {

	ID := context.QueryParam("Id")
	user, err := handler.userRepo.GetByID(ID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultUserResponse{
		Error: false,
		User:  user,
	})
}

// GetUsers method
func (handler *Handler) GetUsers(context echo.Context) error {

	params := new(model.GetLimit)
	if err := context.Bind(params); err != nil {
		fmt.Println("Get limit")
		return context.JSON(http.StatusBadRequest, err)
	}

	users, err := handler.userRepo.GetAll(params)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	totalLocations, err := handler.userRepo.GetCantTotal()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, UserList{
		Data:         users,
		TotalRecords: totalLocations,
	})
}

// CreateUser method
func (handler *Handler) CreateUser(context echo.Context) error {

	request := new(model.User)
	if err := context.Bind(request); err != nil {
		fmt.Println("12")
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	req := CreateValidate(*request)
	if err := req.ValiCreate(context); err != nil {
		fmt.Println("13")
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	//validate email, pwd
	vali, mssg := CheckUser(handler, request.Email, request.Password)
	if vali == false {
		fmt.Println("14")
		return context.JSON(http.StatusInternalServerError, errors.New(mssg))
	}

	// Dates Mongodb
	request.ID = primitive.NewObjectID()
	request.Password, _ = HashPwd(request.Password)
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	err := handler.userRepo.CreateNew(request)
	if err != nil {
		fmt.Println("15")
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultUserResponse{
		Error: false,
		User:  *request,
	})

}

// UpdateUser method
func (handler *Handler) UpdateUser(context echo.Context) error {

	ID := context.QueryParam("Id")
	if len(ID) == 0 {
		return context.JSON(http.StatusInternalServerError, errors.New("userId queryParam is missing"))
	}

	request := new(model.UserUpdate)
	if err := context.Bind(request); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	req := UpdateValidate(*request)
	if err := req.ValiUpdate(context); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	user, err := handler.userRepo.GetByID(ID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	req.Populate(user)
	updated, err := handler.userRepo.Update(
		ID,
		model.Userupdate,
	)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultUserResponse{
		Error: false,
		User:  updated,
	})

}

// DeleteUser method
func (handler *Handler) DeleteUser(context echo.Context) error {

	ID := context.QueryParam("Id")

	if len(ID) == 0 {
		return context.JSON(http.StatusInternalServerError, errors.New("userId queryParam is missing"))
	}

	user, err := handler.userRepo.GetByID(ID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	if err := handler.userRepo.Delete(ID); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultUserResponse{
		Error: false,
		User:  user,
	})

}

// UploadAvatar method
func (handler *Handler) UploadAvatar(context echo.Context) error {

	token := context.Get("userid").(*jwt.Token)
	claims := token.Claims.(*middelware.Claim)
	userid := claims.ID.Hex()

	file, err := context.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		fmt.Println("open")
		return err
	}

	defer src.Close()

	// Destination
	var extension = strings.Split(file.Filename, ".")[1]
	var archivo string = "user/upload/avatars/" + userid + "." + extension
	fmt.Println(archivo)
	dst, err := os.Create(archivo)
	if err != nil {
		fmt.Println("create")
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		fmt.Println("copy")
		return err
	}

	var avatar model.UserUpdate
	avatar.Avatar = userid + "." + extension

	user, err := handler.userRepo.GetByID(userid)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	req := UpdateValidate(avatar)
	req.Populate(user)
	updated, err := handler.userRepo.Update(
		userid,
		model.Userupdate,
	)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultUserResponse{
		Error: false,
		User:  updated,
	})

}

// UploadFront method
func (handler *Handler) UploadFront(context echo.Context) error {

	token := context.Get("userid").(*jwt.Token)
	claims := token.Claims.(*middelware.Claim)
	userid := claims.ID.Hex()

	file, err := context.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		fmt.Println("open")
		return err
	}

	defer src.Close()

	// Destination
	var extension = strings.Split(file.Filename, ".")[1]
	var archivo string = "user/upload/fronts/" + userid + "." + extension
	fmt.Println(archivo)
	dst, err := os.Create(archivo)
	if err != nil {
		fmt.Println("create")
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		fmt.Println("copy")
		return err
	}

	var front model.UserUpdate
	front.Front = userid + "." + extension

	user, err := handler.userRepo.GetByID(userid)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	req := UpdateValidate(front)
	req.Populate(user)
	updated, err := handler.userRepo.Update(
		userid,
		model.Userupdate,
	)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultUserResponse{
		Error: false,
		User:  updated,
	})

}

// LoginUser method
func (handler *Handler) LoginUser(context echo.Context) error {
	jsonResult := new(middelware.Responsetoken)
	var answer = http.StatusOK
	date := new(middelware.DateValidate)
	if err := context.Bind(date); err != nil {
		answer = http.StatusForbidden
		return context.JSON(answer, err)
	}

	user, err := handler.userRepo.ValidateUser(date.Email, date.Password)
	if err != nil {
		answer = http.StatusForbidden
		return context.JSON(answer, err)
	}

	if user.UserName != "" && user.Email != "" {
		//create a struct of my Claim
		claims := middelware.Claim{
			Token: user,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
				Issuer:    "token test", //object of token
			},
		}

		//--------------------encode to base64-----------------//
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokeS, err := token.SignedString(middelware.Keys())
		//token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
		//tokeS, err := token.SignedString(PrivateKey)
		if err != nil {
			tokeS = "could not sign private token"
		}
		answer = http.StatusOK
		jsonResult.Token = tokeS
	} else {
		answer = http.StatusForbidden
		jsonResult.Token = "usser or password invalid"
	}
	return context.JSON(answer, jsonResult)
}

// ValidateToken
func (handler *Handler) ValidateToken(context echo.Context) error {
	jsonResult := new(middelware.Responsetoken)
	var answer = http.StatusOK
	token, err := request.ParseFromRequestWithClaims(context.Request(), request.OAuth2Extractor, &middelware.Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return middelware.PublicKey, nil
		})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				answer = http.StatusUnauthorized
				jsonResult.Token = "your token expired"
			case jwt.ValidationErrorSignatureInvalid:
				answer = http.StatusUnauthorized
				jsonResult.Token = "the signature does not match"
			default:
				answer = http.StatusUnauthorized
				jsonResult.Token = "the signature does not match"
			}
		default:
			answer = http.StatusUnauthorized
			jsonResult.Token = "your token is not valid"
		}
	}
	if token.Valid {
		answer = http.StatusAccepted
		jsonResult.Token = "welcome to the system"
	} else {
		answer = http.StatusUnauthorized
		jsonResult.Token = "your token is not valid"
	}
	return context.JSON(answer, jsonResult)
}
