package main

import (
	"net/http"

	"github.com/OlaleyeJumoke/WebService/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
