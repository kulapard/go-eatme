VERSION=$(TRAVIS_TAG)
NAME=eatme

all: mac64 linux64 win64

mac64:
	GOOS=darwin GOARCH=amd64 go build -o _build/$(NAME)_$(VERSION)_darwin_amd64/$(NAME)
	cd _build && tar -cvzf $(NAME)_$(VERSION)_darwin_amd64.tar.gz $(NAME)_$(VERSION)_darwin_amd64/$(NAME)
	rm -rf _build/$(NAME)_$(VERSION)_darwin_amd64

linux64:
	GOOS=linux GOARCH=amd64 go build -o _build/$(NAME)_$(VERSION)_linux_amd64/$(NAME)
	cd _build && tar -cvzf $(NAME)_$(VERSION)_linux_amd64.tar.gz $(NAME)_$(VERSION)_linux_amd64/$(NAME)
	rm -rf _build/$(NAME)_$(VERSION)_linux_amd64

win64:
	GOOS=windows GOARCH=amd64 go build -o _build/$(NAME)_$(VERSION)_windows_amd64/$(NAME).exe
	cd _build && zip $(NAME)_$(VERSION)_windows_amd64.zip $(NAME)_$(VERSION)_windows_amd64/$(NAME).exe
	rm -rf _build/$(NAME)_$(VERSION)_windows_amd64
