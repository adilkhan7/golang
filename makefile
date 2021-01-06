SHELL := /bin/zsh

tidy:
	go mod tidy
	go mod vendor