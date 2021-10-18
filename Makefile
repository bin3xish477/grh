windows:
	GOOS=windows GOARCH=amd64 go build -o ./grh.exe grh.go
windows-release:
	GOOS=windows GOARCH=amd64 go build -o ./grh.exe grh.go
	tar -czvf grh-windows-amd64.tar.gz ./grh.exe ./README.md
	sha256sum grh-windows-amd64.tar.gz > grh-windows-amd64.tar.gz.sha256sum
	GOOS=windows GOARCH=386 go build -o ./grh.exe grh.go
	tar -czvf grh-windows-386.tar.gz ./grh.exe ./README.md
	sha256sum grh-windows-386.tar.gz > grh-windows-386.tar.gz.sha256sum

darwin:
	GOOS=darwin GOARCH=amd64 go build -o ./grh grh.go
darwin-release:
	GOOS=darwin GOARCH=amd64 go build -o ./grh grh.go
	tar -czvf grh-darwin-amd64.tar.gz ./grh ./README.md
	sha256sum grh-darwin-amd64.tar.gz > grh-darwin-amd64.tar.gz.sha256sum

linux:
	GOOS=linux GOARCH=amd64 go build -o ./grh grh.go
linux-release:
	GOOS=windows GOARCH=amd64 go build -o ./grh grh.go
	tar -czvf grh-linux-amd64.tar.gz ./grh ./README.md
	sha256sum grh-linux-amd64.tar.gz > grh-linux-amd64.tar.gz.sha256sum
	GOOS=windows GOARCH=386 go build -o ./grh grh.go
	tar -czvf grh-linux-386.tar.gz ./grh ./README.md
	sha256sum grh-linux-386.tar.gz > grh-linux-386.tar.gz.sha256sum

releases: linux-release windows-release darwin-release

clean:
	@if [ -f grh ]; then rm grh; fi
	@if [ -f grh.exe ]; then rm grh.exe; fi
