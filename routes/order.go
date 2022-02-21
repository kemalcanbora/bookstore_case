package routers

import (
	"bookstore_case/models"
	"bookstore_case/pkg"
	"encoding/json"
	"fmt"
	"net/http"
)

func OrderBuy(response http.ResponseWriter, r *http.Request) {
	var order []models.Order
	var u models.JwtUserAuth

	user := r.Context().Value("data")
	tmp, _ := json.Marshal(user)
	json.Unmarshal(tmp, &u)

	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, item := range order {
		pkg.Mongo.Order(pkg.DbClient, item)
	}
	fmt.Fprintf(response, "Hello %s", u.Email, "Order is Done!")
}
