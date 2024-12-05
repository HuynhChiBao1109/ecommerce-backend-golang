package main

import "ecommerce-backend-golang/internal/routers"

func main() {
	r := routers.InitRouter()

	r.Run(":8002") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
