package res

// import (
// 	jwt "golang_service/auth"
// 	"golang_service/models"
// 	"net/http"

// 	"gorm.io/gorm"
// )

// type handler struct {
// 	DB *gorm.DB
// }

// func ValidateLoginState(r *http.Request, h handler) bool {
// 	userName, err := jwt.ExtractUsernameFromToken(r)
// 	if err != nil {
// 		return false
// 	}
// 	var user models.User
// 	errFindUser := h.DB.Where("user_name = ?", userName).First(&user)
// 	if errFindUser.Error != nil {
// 		return false
// 	}
// 	return true
// }
