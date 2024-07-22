package handlers

import "github.com/levifleal/socialMedia/backEnd/handlers/auth"

func Init() {
	auth.InitAuthHandler()
}
