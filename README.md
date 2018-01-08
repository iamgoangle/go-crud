# go-crud
My public Restful API with Golang

# Prerequisition
- Install Go compiler to your computer
- Pull code to your `$GOPATH/src`
- Install Glide package manager for Golang

# Glide Install
```
brew install glide
```

# Installation
```
glide install
```

# Application routes
## GET /users
```sh
curl http://localhost:3000/users
```
## GET /users/:id
```sh
curl http://localhost:3000/user/<your-user>
```
## GET /show?team=<string>&member=<string>
```sh
curl http://localhost:3000/show?team=foo&member=bar
```
## POST /createUser
```sh
curl http://localhost:3000/createUser
```
## PATCH
```sh
todo
```

## PUT
```sh
todo
```
## DELETE
```sh
todo
```
