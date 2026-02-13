package models

// User 구조체 - 사용자 정보
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// 메모리에 저장하는 사용자 목록 =
var Users = []User{
	{ID: 1, Name: "김철수", Email: "kim@example.com", Age: 25},
	{ID: 2, Name: "이영희", Email: "lee@example.com", Age: 30},
	{ID: 3, Name: "박민수", Email: "park@example.com", Age: 28},
}

// 새 사용자 추가시 사용할 ID
var NextID = 4
