# go-crud
My simple public project in order to create Restful API with Golang.
Please note that understand that this repository does not guide you for best practice in golang.

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
curl -X POST \
  http://localhost:3000/createUser \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
    "id": "1",
    "name": "Teerapong",
    "email": "st.teerapong@gmail.com"
  }'
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
