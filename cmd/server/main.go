package main

import "github.com/BWCbewchan/GO-ECOMMERCE-BACKEND-API/internal/router"


func main() {
	// curl http://localhost:8080/v1/2024/ping/bewchan?uid=1234
	r := router.NewRouter()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// api or controller
