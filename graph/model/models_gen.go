// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type InputCart struct {
	Gameid string `json:"gameid"`
	Userid string `json:"userid"`
}

type InputOwnedGame struct {
	Gameid string `json:"gameid"`
	Userid string `json:"userid"`
}

type InputReview struct {
	Gameid      string `json:"gameid"`
	Userid      string `json:"userid"`
	Description string `json:"description"`
	Rating      string `json:"rating"`
}

type InputUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Country  string `json:"country"`
}

type InputWishlist struct {
	Gameid string `json:"gameid"`
	Userid string `json:"userid"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
