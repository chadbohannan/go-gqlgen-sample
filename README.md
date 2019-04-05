
This is a crash course in GraphQL using a schema-first approach (generated code).

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
When it was created it was filled with function stubs that contained panic statements.
The panic statements have been replaced with some basic server functions.
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

Run the server with:
```
go run server/server.go
```

Then browse to:
````
http://localhost:8080/graphql
````

If everything went well you should see a dark themed developer playground.
The server is running from memory, so a fresh instance has no data.

This sample project has defined [Articles](./gql.article) and [Users](./gql.user).

_from gql.article_:
```
...
input NewArticle {
  text: String! # content is required
  userID: Int!  # userID is required
}
...
```
To create an Article we will need to use the NewArticle object, which requires 
a userID attribute. This means we need to first create a user.


Copy this snippet of GQL into the playground:
```
mutation {
  newUser(input: {
    name:"Code Smith"
  }){
    id
  }
}
```

This is parsed by the code in `generated.go` and calls `NewUser(ctx context.Context, input NewUser)` 
in `resolver.go` which returns a User, but we have only requested the id attribute to be returned.

_expected response_:
```
{
  "data": {
    "newUser": {
      "id": 0
    }
  }
}
```

Now we have a userID can create a new Article:
```
mutation {
  newArticle(input: {
    text:"When you have a hammer everything is a nail. When you have a smith everything is anything you want it to be."
    userID: 0
  }){
    id
    text
    user {
      name
    }
  }
}
```

You should now have output that looks like:
```
{
  "data": {
    "newArticle": {
      "id": 0,
      "text": "When you have a hammer everything is a nail. When you have a smith everything is anything you want it to be.",
      "user": {
        "name": "Code Smith"
      }
    }
  }
}
```

You have now been introduced to everything you need to create new mutation query.
Try changing the text of the article text to anything you want it to be.
