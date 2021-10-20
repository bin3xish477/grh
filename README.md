# grh
grh (Get Response Headers) is a tool whose sole purpose is to print out HTTP response headers and follow redirects with color!

![image](https://user-images.githubusercontent.com/44281620/137790729-67d14e22-4aa1-4c65-aee3-7980e207fb16.png)

### Installation
1. Download a binary for your target operating system from the [Releases](https://github.com/binexisHATT/grh/releases) page
2. Install with go: `go get -v github.com/binexisHATT/grh`
3. From source (You must have Go installed)
  - git clone https://github.com/binexisHATT/grh
  - cd grh
  - go build -o grh grh.go
  - sudo mv grh /usr/local/bin

### Usage
> cat urls | grh

> grh --url 'http://youtube.com'
