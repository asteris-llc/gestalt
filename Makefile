all: web/web.go web/app web/client web/swagger

web/app: web/design/*.go
	cd web && goagen app -d github.com/asteris-llc/gestalt/web/design

web/swagger: web/design/*.go
	cd web && goagen swagger -d github.com/asteris-llc/gestalt/web/design

web/client: web/design/*.go
	cd web && goagen client -d github.com/asteris-llc/gestalt/web/design

web/impl: web/design/*.go
	@test -d web/impl || mkdir web/impl
	cd web/impl && goagen main -d github.com/asteris-llc/gestalt/web/design

web/web.go: web/impl web.go.sed
	echo '// !!! automatically generated !!!\n// Use "make web/web.go" instead of editing this file.\n' | cat - web/impl/main.go > web/web.go
	sed -i '' -f web.go.sed web/web.go
