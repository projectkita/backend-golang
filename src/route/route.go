package route

import (
	"fmt"

	"github.com/julienschmidt/httprouter"
	"github.com/projectkita/project-harapan-backend-golang/src/handler"
)

func banner() {
	//Cari tau cara untuk buat huruf di cmd
	fmt.Println(`APP HARAPAN KITA`)
}

func InitRoutes(port string) {
	router := httprouter.New()

	router.GET("/hello", MiddlewareAccessLog(handler.GetHelloHandler))
	router.GET("/hello/:id", MiddlewareAccessLog(handler.GetHelloByIDHandler))

	banner()
}
