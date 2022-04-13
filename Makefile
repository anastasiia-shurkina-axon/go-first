.PHONY: go-start
go-start:
	docker run --rm -it -p 8081:8081 -v /Users/user/dev/go-first:/usr/src/myapp -w /usr/src/myapp golang:1.17 bash