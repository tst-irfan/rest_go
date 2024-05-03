package main

import (
	"rest_go/initializer"
)

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	initializer.RunApplication()
}
