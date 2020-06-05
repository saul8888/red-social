package follow

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/orderforme/user/followers/database"
	"github.com/orderforme/user/followers/errors"
	"github.com/orderforme/user/followers/model"
	"github.com/orderforme/user/middelware"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Follow interface {
	Done()
	Follow(context echo.Context) error
	Unfollow(context echo.Context) error
	Queryfollow(context echo.Context) error
}

type Handler struct {
	followRepo database.MongoDB
}

// NewHandler allocates a new Handler
func NewHandler() (*Handler, error) {

	handler := Handler{
		followRepo: &database.Mongodb{},
	}

	err := handler.followRepo.ConnectDB()

	return &handler, err
}

// DisconnectDB all
func (handler *Handler) Done() {
	handler.followRepo.DisconnectDB()
}

// Follow method
func (handler *Handler) Follow(context echo.Context) error {

	request := new(model.Follow)

	ID := context.QueryParam("Id")

	token := context.Get("userid").(*jwt.Token)
	claims := token.Claims.(*middelware.Claim)

	objectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.New("ID invalid"))
	}

	request.ID = primitive.NewObjectID()
	request.UserID = claims.ID
	request.FollowingID = objectID
	request.CreatedAt = time.Now()

	err = handler.followRepo.CreateNew(request)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultFollowResponse{
		Error:  false,
		Follow: *request,
	})
}

// Unfollow method
func (handler *Handler) Unfollow(context echo.Context) error {

	ID := context.QueryParam("Id")

	if len(ID) == 0 {
		return context.JSON(http.StatusInternalServerError, errors.New("Id queryParam is missing"))
	}

	token := context.Get("userid").(*jwt.Token)
	claims := token.Claims.(*middelware.Claim)

	follow, err := handler.followRepo.GetFollow(claims.ID, ID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	if err := handler.followRepo.Delete(claims.ID, ID); err != nil {
		return context.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, DefaultFollowResponse{
		Error:  false,
		Follow: follow,
	})
}

func (handler *Handler) Queryfollow(context echo.Context) error {
	ID := context.QueryParam("Id")

	if len(ID) == 0 {
		return context.JSON(http.StatusInternalServerError, errors.New("Id queryParam is missing"))
	}

	token := context.Get("userid").(*jwt.Token)
	claims := token.Claims.(*middelware.Claim)

	follow, err := handler.followRepo.GetFollow(claims.ID, ID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errors.New("not follwing"))
	}

	return context.JSON(http.StatusOK, DefaultFollowResponse{
		Error:  false,
		Follow: follow,
	})
}
