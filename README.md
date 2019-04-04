
First things first:
```
go get github.com/99designs/gqlgen
```

This repo was started with gql and yml files so we don't run `init`.
Instead we run
```
go run github.com/99designs/gqlgen
```
This looks for the yml file (`gqlgen.yml` in our case) and generates models
and the code that maps the GQL parser to the resolver functions that we
need to implement.

The yaml file has resolver.go set so that's where code first went.
Go structs allow interface methods to be defined in other files,
but a nice convention is to leave the struct declarations in there.

This Go directive is at the top of resolver.go
```
//go:generate go run github.com/99designs/gqlgen
```
This directive  simplifies command line usage to:
```
$ go generate
```

A mod file was defined and with the generated output in place we
can fill a vendor directory with our dependencies: 
```
$ go mod vendor
```
Which creates the vendor subdirectory and fills it with delicious code units.
