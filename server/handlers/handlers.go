package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"example.com/go-study/server/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go 웹 서버 실행 중")
}

// UsersHandler - 사용자 CRUD 처리
// w: 응답 보내는 통로, r: 받은 요청 정보
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 요청 메서드별로 처리 (GET/POST/DELETE)
	switch r.Method {

	case "GET": // 조회
		path := r.URL.Path

		// "/users" -> 전체 목록
		if path == "/users" || path == "/users/" {
			json.NewEncoder(w).Encode(models.Users)

		} else {
			// "/users/1" -> 특정 사용자 조회
			parts := strings.Split(path, "/") // ["", "users", "1"]
			if len(parts) >= 3 && parts[2] != "" {
				id, err := strconv.Atoi(parts[2]) // "1" -> 1
				if err != nil {
					http.Error(w, "잘못된 ID", http.StatusBadRequest)
					return
				}

				// ID로 사용자 찾기
				for _, user := range models.Users {
					if user.ID == id {
						json.NewEncoder(w).Encode(user)
						return
					}
				}
				http.Error(w, "사용자 없음", http.StatusNotFound)
			}
		}

	case "POST": // 생성
		var newUser models.User
		// JSON 요청 -> User 구조체로 변환
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			http.Error(w, "잘못된 데이터", http.StatusBadRequest)
			return
		}

		// ID 할당 후 슬라이스에 추가
		newUser.ID = models.NextID
		models.NextID++
		models.Users = append(models.Users, newUser)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newUser)

	case "DELETE": // 삭제
		path := r.URL.Path
		parts := strings.Split(path, "/")

		if len(parts) >= 3 && parts[2] != "" {
			id, err := strconv.Atoi(parts[2])
			if err != nil {
				http.Error(w, "잘못된 ID", http.StatusBadRequest)
				return
			}

			// ID로 찾아서 슬라이스에서 제거
			for i, user := range models.Users {
				if user.ID == id {
					// [:i] + [i+1:] = i번째 요소 제거
					models.Users = append(models.Users[:i], models.Users[i+1:]...)
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(map[string]string{
						"message": fmt.Sprintf("ID %d 삭제됨", id),
					})
					return
				}
			}
			http.Error(w, "사용자 없음", http.StatusNotFound)
		} else {
			http.Error(w, "ID 필요", http.StatusBadRequest)
		}

	default:
		http.Error(w, "지원하지 않는 메서드", http.StatusMethodNotAllowed)
	}
}
