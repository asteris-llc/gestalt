# package
s/package main/package web/

# imports
s|impl/||g
/import \(/ a\
	"github.com/asteris-llc/gestalt/store"

# main -> Run
s/func main\(\)/func Run(addr string, store *store.Store)/
/func Run/ i\
// Run starts the server

# make the server graceful (and change the name)
s/goa.New("API")/goa.NewGraceful("gestalt")/

# thread in the store
s/(New.+Controller)\((.*)\)/\1(\2, store)/g

# ports
s/port 8080/given addr/
s/":8080"/addr/