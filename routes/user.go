package routers

import (
	"bookstore_case/models"
	"bookstore_case/pkg"
	"bookstore_case/pkg/helper"
	"encoding/json"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"time"
)

var user models.User

func SignUp(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(request.Body).Decode(&user)
	validate := validator.New()
	errVal := validate.Struct(user)
	if errVal != nil {
		helper.HTTPErrorHandler(response, "Email or Password field is empty!", http.StatusUnauthorized)
		return
	}

	userCheck, _ := pkg.Mongo.FindUserWithEmail(pkg.DbClient, user.Email)
	if userCheck.Email == user.Email {
		helper.HTTPErrorHandler(response, "This email is already registered!", http.StatusNotAcceptable)
		return
	}

	if err != nil {
		helper.HTTPErrorHandler(response, `{"message":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	user.Password = pkg.GetHash([]byte(user.Password))
	user.CreatedTime = time.Now().Unix()

	result, err := pkg.Mongo.Insert(pkg.DbClient, user, "users")
	if err != nil {
		helper.HTTPErrorHandler(response, `{"message":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}
	json.NewEncoder(response).Encode(result)
}

func UserLogin(response http.ResponseWriter, request *http.Request) {
	json.NewDecoder(request.Body).Decode(&user)
	response.Header().Set("Content-Type", "application/json")
	result, err := pkg.Mongo.FindUserWithEmail(pkg.DbClient, user.Email)
	if err != nil {
		helper.HTTPErrorHandler(response, `{"message":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}
	passErr := pkg.CheckPasswordHash(user.Password, result.Password)

	if passErr != true {
		log.Println(passErr)
		response.Write([]byte(`{"response":"Wrong Password!"}`))
		return
	}

	jwtToken, err := pkg.GenerateJWT(result)
	if err != nil {
		helper.HTTPErrorHandler(response, `{"message":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}
	response.Write([]byte(`{"token":"` + jwtToken + `"}`))
}
