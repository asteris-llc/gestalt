all: web/web.go web/app web/client web/swagger

web/app: web/design
	cd web && goagen app -d github.com/asteris-llc/gestalt/web/design

web/swagger: web/design
	cd web && goagen swagger -d github.com/asteris-llc/gestalt/web/design

web/client: web/design
	cd web && goagen client -d github.com/asteris-llc/gestalt/web/design

web/impl: web/design
	@mkdir web/impl
	cd web/impl && goagen main -d github.com/asteris-llc/gestalt/web/design

web/web.go: web/impl
	echo '// !!! automatically generated !!!\n// Use "make web/web.go" instead of editing this file.\n' | cat - web/impl/main.go > web/web.go
	sed -i '' -e 's/impl\///g' web/web.go
	sed -i '' -e 's/package main/package web/' web/web.go
	sed -i '' -e 's/func main()/func Run(addr string)/' web/web.go
	sed -i '' -e 's/":8080"/addr/' web/web.go
	sed -i '' -e 's/port 8080/given addr/' web/web.go
