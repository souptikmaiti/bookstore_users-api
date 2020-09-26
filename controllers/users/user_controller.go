package users

import (
	"github.com/gin-gonic/gin"
	"github.com/souptikmaiti/bookstore_users-api/domain/users"
	"github.com/souptikmaiti/bookstore_users-api/services"
	"github.com/souptikmaiti/bookstore_users-api/utils/errors"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	//json.NewDecoder(c.Request.Body).Decode(&user)
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid Json Body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {

}

/*func SearchUser(c *gin.Context) {

}*/
