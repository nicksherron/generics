# Generics
Official tutorial on getting started with generics can be found  [here](https://go.dev/doc/tutorial/generics).


##  Installing go1.18beta1

```sh
# To install run
go install golang.org/dl/go1.18beta1@latest

# To download updates run 
go1.18beta1 download

# Then alias it to the go command
alias go='go1.18beta1'
```



## Installing gopls with generics support
[see tools](https://github.com/golang/tools/blob/master/gopls/doc/advanced.md) for more detailed instructions


```sh
# Create an empty go.mod file, only for tracking requirements. 
cd $(mktemp -d)
go mod init gopls-unstable

# Make sure its using go 1.18
cat go.mod
##	   module gopls-unstable
##	   
##	   go 1.18
##	   

# Use 'go get' to add requirements and to ensure they work together.
go get golang.org/x/tools/gopls@master golang.org/x/tools@master

# For go1.17 or older, the above `go get` command will build and
# install `gopls`. For go1.18+ or tip, run the following command to install
# using selected versions in go.mod.
go install golang.org/x/tools/gopls
```
