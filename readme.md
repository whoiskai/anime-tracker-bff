# Anime Tracker BFF

BFF == Backend For Frontend

This is the backend for the anime tracker frontend. Currently in initial status, so check back later when there are more updates.

Referencing this repo for basic CRUD operations

https://github.com/guizot/golang-gin-mongo

## Quick start

Using docker for mongodb

```bash
docker run --name mongodb -d -p 27017:27017 mongo:4.2
source .env && go run main.go
```