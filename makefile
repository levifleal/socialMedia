.PHONY: default run build test docs clean

APP_NAME=MyFlix
BACKEND_DIR=./backEnd
MAIN_FILE=main.go

default: run-backend

run-backend:
	@clear
	@echo "Starting $(APP_NAME) API..."
	@cd $(BACKEND_DIR) && go run $(MAIN_FILE)