package main

import (
	"api-main/configs"
	"api-main/internal/auth"
	"api-main/internal/link"
	"api-main/internal/stat"
	"api-main/internal/user"
	"api-main/pkg/db"
	"api-main/pkg/midlleware"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	//Repositoryes
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)
	// Servises
	authServices := auth.NewAuthService(userRepository)

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authServices,
	})

	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		Config:         conf,
		StatRepository: statRepository,
	})

	// Middlewares
	stack := midlleware.Chain(
		midlleware.CORS,
		midlleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()

}

// func ping(url string, respCh chan int, errCh chan error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		errCh <- err
// 		return
// 	}
// 	respCh <- resp.StatusCode
// }

// func main() {
// 	path := flag.String("file", "url.txt", "path to url")
// 	flag.Parse()
// 	file, err := os.ReadFile(*path)
// 	if err != nil {
// 		panic(err.Error())

// 	}
// 	urlSlice := strings.Split(string(file), "\n")
// 	respCh := make(chan int)
// 	errCh := make(chan error)
// 	for _, url := range urlSlice {
// 		go ping(url, respCh, errCh)
// 	}
// 	for range urlSlice {
// 		select {
// 		case respErr := <-errCh:
// 			fmt.Println(respErr)
// 		case res := <-respCh:
// 			fmt.Println(res)

// 		}

// 	}

// }
