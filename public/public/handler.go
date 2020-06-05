package public

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/orderforme/user/middelware"
	"github.com/orderforme/user/public/database"
	"github.com/orderforme/user/public/errors"
	"github.com/orderforme/user/public/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//public Public
type Public interface {
	Done()
	GetPublicByID(context echo.Context) error
	GetPublics(context echo.Context) error
	CreatePublic(context echo.Context) error
	UpdatePublic(context echo.Context) error
	DeletePublic(context echo.Context) error
	ValidateToken(context echo.Context) error
}

type Handler struct {
	publicRepo database.MongoDB
}

// NewHandler allocates a new Handler
func NewHandler() (*Handler, error) {

	handler := Handler{
		publicRepo: &database.Mongodb{},
	}

	err := handler.publicRepo.ConnectDB()

	return &handler, err
}

// DisconnectDB all
func (handler *Handler) Done() {
	handler.publicRepo.DisconnectDB()
}

// GetPublicByID method
func (handler *Handler) GetPublicByID(context echo.Context) error {

	ID := context.QueryParam("Id")
	public, err := handler.publicRepo.GetByID(ID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultPublicResponse{
		Error:  false,
		Public: public,
	})
}

// GetPublics method
func (handler *Handler) GetPublics(context echo.Context) error {

	params := new(model.GetLimit)
	if err := context.Bind(params); err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	publics, err := handler.publicRepo.GetAll(params)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	totalPublics, err := handler.publicRepo.GetCantTotal()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, PublicList{
		Data:         publics,
		TotalRecords: totalPublics,
	})
}

/*
func hola(context echo.Context) error {
	hola := context.Get("public").(*jwt.Token)
	claims := hola.Claims.(*model.Claim)
	hola11 := claims.UserName
	return context.JSON(http.StatusOK, hola11)

}
*/
// CreatePublic method
func (handler *Handler) CreatePublic(context echo.Context) error {
	request := new(model.Public)
	if err := context.Bind(request); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	token := context.Get("userid").(*jwt.Token)
	claims := token.Claims.(*middelware.Claim)

	req := CreateValidate(*request)
	if err := req.ValiCreate(context); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}
	/*
		//validate ID externt
		err := handler.publicRepo.ValidateID("merchant", request.MerchantID)
		if err != nil {
			return context.JSON(http.StatusInternalServerError, errors.NewError(err))
		}
	*/
	// Dates Mongodb
	request.ID = primitive.NewObjectID()
	request.UserID = claims.ID
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	err := handler.publicRepo.CreateNew(request)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultPublicResponse{
		Error:  false,
		Public: *request,
	})

}

// UpdatePublic method
func (handler *Handler) UpdatePublic(context echo.Context) error {

	ID := context.QueryParam("Id")
	if len(ID) == 0 {
		return context.JSON(http.StatusInternalServerError, errors.New("publicId queryParam is missing"))
	}

	request := new(model.PublicUpdate)
	if err := context.Bind(request); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	req := UpdateValidate(*request)
	if err := req.ValiUpdate(context); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	public, err := handler.publicRepo.GetByID(ID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	req.Populate(public)
	updatedPublic, err := handler.publicRepo.Update(
		ID,
		model.Publicupdate,
	)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultPublicResponse{
		Error:  false,
		Public: updatedPublic,
	})

}

// DeletePublic
func (handler *Handler) DeletePublic(context echo.Context) error {

	ID := context.QueryParam("Id")

	if len(ID) == 0 {
		return context.JSON(http.StatusInternalServerError, errors.New("publicId queryParam is missing"))
	}

	public, err := handler.publicRepo.GetByID(ID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	if err := handler.publicRepo.Delete(ID); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}
	//return context.JSON(http.StatusOK, public)
	return context.JSON(http.StatusOK, DefaultPublicResponse{
		Error:  false,
		Public: public,
	})

}

// ValidateToken
func (handler *Handler) ValidateToken(context echo.Context) error {
	hola := context.Get("userid").(*jwt.Token)
	claims := hola.Claims.(*middelware.Claim)
	hola11 := claims.UserName
	return context.JSON(http.StatusOK, hola11)
}
