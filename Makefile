# PROJECT_NAME := ""
APP_NAME := "airbnb"
PKG := "./"
CMD := "$(PKG)/cmd"

build: ## Build the binary file
	@go build -o $(APP_NAME).o -i -v $(CMD)

run: build
	@./$(APP_NAME).o