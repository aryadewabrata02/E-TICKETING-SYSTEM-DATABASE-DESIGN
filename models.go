package main

type User struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Terminal struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
