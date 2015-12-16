# package
s/package main/package web/

# imports
s|impl/||g

# main -> Run
s/func main()/func Run(addr string)/
/func Run/ i\
// Run starts the server

# make the server graceful (and change the name)
s/goa.New("API")/goa.NewGraceful("gestalt")/

# ports
s/port 8080/given addr/
s/":8080"/addr/