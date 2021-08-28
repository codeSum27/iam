package api

import (
	apiError "github.com/codeSum27/iam/pkg/api/error"
	"github.com/codeSum27/iam/pkg/common"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func (i *IamServer) GetUserById(ctx echo.Context, id string, params GetUserByIdParams) error {
	i.Lock.Lock()
	defer i.Lock.Unlock()

	err := verifyToken(common.Cnf.Token.AccessSecret,params.AccessToken,id)
	if err != nil {
		return sendCreateTokenError(ctx, http.StatusBadRequest, err.Error())
	}

	user, err := getUser(id)
	if err != nil {
		return sendCreateTokenError(ctx, http.StatusBadRequest, err.Error())
	}

	err = ctx.JSON(http.StatusOK, user)
	if err != nil {
		// Something really bad happened, tell Echo that our handler failed
		return err
	}
	return nil
}

func (i *IamServer)	UpdateUser(ctx echo.Context, id string, params UpdateUserParams) error {

	err := verifyToken(common.Cnf.Token.AccessSecret,params.AccessToken,id)
	if err != nil {
		return sendCreateTokenError(ctx, http.StatusForbidden, "Invalid Access token for Update user")
	}

	user , err := getUser(id)
	if user == nil {
		return sendCreateTokenError(ctx, http.StatusNotFound, "Not found user")
	}

	var target User
	err = ctx.Bind(&target)
	if err != nil {
		return sendCreateTokenError(ctx, http.StatusBadRequest, "Invalid format for update User")
	}
	err = updateUser(user,&target)
	if err != nil {
		return sendCreateTokenError(ctx, http.StatusBadRequest, "Cannot update user")
	}

	err = ctx.JSON(http.StatusOK, user)
	if err != nil {
		return sendCreateTokenError(ctx, http.StatusBadRequest, err.Error())
	}
	return nil
}

// (DELETE /users)
func (i *IamServer) DeleteUserById(ctx echo.Context, id string, params DeleteUserByIdParams) error{

	err := verifyToken(common.Cnf.Token.AccessSecret,params.AccessToken,id)
	if err != nil {
		return sendCreateTokenError(ctx, http.StatusForbidden, "Invalid Access token for Delete User")
	}

	err = deleteUser(id)
	if err != nil {
		return sendCreateTokenError(ctx, http.StatusNotFound, "Invalid format for Delete User")
	}

	err = ctx.JSON(http.StatusOK, id)
	if err != nil {
		// Something really bad happened, tell Echo that our handler failed
		return sendCreateTokenError(ctx, http.StatusBadRequest, "Bad request for Delete User")
	}
	return nil
}

// (POST /users)
func (i *IamServer) PostUsers(ctx echo.Context) error {

	i.Lock.Lock()
	defer i.Lock.Unlock()

	var newUser User
	err := ctx.Bind(&newUser)
	if err != nil {
		return sendCreateTokenError(ctx, http.StatusBadRequest, "Invalid format for Create User")
	}

	userId, err := createUser(&newUser)
	if err != nil {
		return sendCreateTokenError(ctx, http.StatusBadRequest, err.Error() )
	}

	err = ctx.JSON(http.StatusCreated, userId)
	if err != nil {
		// Something really bad happened, tell Echo that our handler failed
		return err
	}
	return nil
}

func getUser(userId string)  (*User,error) {
	db := GetDB()
	user := &User{}
	return user,db.Transaction(func(tx *gorm.DB) error {
		if err := db.First(&user, "id = ?", userId).Error; err != nil {
			return err
		}
		if user.Id == nil {
			return apiError.NewUserNotFound("User not found")
		}
		user.Password = ""
		return nil
	})
}

func verifyUser(user *User)  error {
	db := GetDB()
	dbUser := &User{}
	err := db.Transaction(func(tx *gorm.DB) error {
		log.Println(user.Email)
		if err := db.First(&dbUser, "email = ?", user.Email).Error; err != nil {
			return err
		}
		if dbUser.Id == nil {
			return apiError.NewUserNotFound("User not found")
		}
		return nil
	})
	if err != nil {
		return  err
	}
	if !checkPasswordHash(user.Password,dbUser.Password) {
		return apiError.NewInvalidUserPassword("Invalid user password")
	}
	user.Id = dbUser.Id
	return nil
}

func createUser(user *User)  (string,error) {

	db := GetDB()
	hashedPassword , err := createHashedPassword(user.Password)
	if err != nil {
		return "", err
	}
	user.Password = hashedPassword
	var newUUID = new(string)
	*newUUID = uuid.New().String()
	err = db.Transaction(func(tx *gorm.DB) error {
		targetUser := &User{}
		db.Where("email = ?", user.Email).First(&targetUser)
		if targetUser.Id != nil{
			return apiError.NewUserAlreadyExist("User already exist")
		}
		user.Id = newUUID
		if db.Create(&user).Error != nil {
			return db.Create(&user).Error
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return *newUUID,nil
}

func updateUser(user *User, targetUser *User) error {

	err := updateUserData(user,targetUser)
	if err != nil {
		return err
	}
	db := GetDB()
	return db.Transaction(func(tx *gorm.DB) error {
		if err := db.Save(user).Error; err != nil{
			return err
		}
		return nil
	})
}


func updateUserData(user *User, targetUser *User) error {

	if targetUser.Email != "" {
		user.Email = targetUser.Email
	}
	if targetUser.Password != "" {
		hashedPassword, err := createHashedPassword(targetUser.Password)
		if err!= nil{
			return err
		}
		user.Password = hashedPassword
	}
	if targetUser.NickName != nil {
		user.NickName = targetUser.NickName
	}
	return nil
}

func deleteUser(userId string) error {
	db := GetDB()
	return db.Transaction(func(tx *gorm.DB) error {
		if err := db.Where("id = ?",userId).Delete(&User{}).Error; err != nil{
			return err
		}
		return nil
	})
}

func createHashedPassword(plainPassword string)  (string,error){
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 14)
	if err != nil {
		return "", err
	}
	return string(hashedPassword),nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}