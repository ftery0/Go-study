package main

import (
	"fmt"
	"log"
	"net/http"

	"go-study/server/handlers"
)

func main() {
	// 라우팅: URL 경로와 핸들러 함수 연결
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/users", handlers.UsersHandler)
	http.HandleFunc("/users/", handlers.UsersHandler) // /users/1 같은 경로도 처리

	// 서버 시작
	port := ":8080"
	fmt.Printf("서버 시작: http://localhost%s\n", port)
	fmt.Println("  GET    /users      - 목록")
	fmt.Println("  GET    /users/:id  - 조회")
	fmt.Println("  POST   /users      - 추가")
	fmt.Println("  DELETE /users/:id  - 삭제")

	log.Fatal(http.ListenAndServe(port, nil))
}
