package api

import (
	"fmt"

	"api/internal/domain"

	"github.com/gin-gonic/gin"
)

func GetUserFromContext(c *gin.Context) (domain.User, error) {
	user, exists := c.Get("user")
	if !exists {
		return domain.User{}, fmt.Errorf("user not found in context")
	}

	userStruct, ok := user.(domain.User)
	if !ok {
		return domain.User{}, fmt.Errorf("invalid user type in context")
	}

	return userStruct, nil
}
