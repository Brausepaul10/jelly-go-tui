package main

import (
	"fmt"
	"jelly-go-tui/internal/jellyfin"
)

func main() {
	loginDetails := jellyfin.GetUser()
	fmt.Println(loginDetails)
}
