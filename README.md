## Pretest
Pretest Project

## Installation
- Clone this repository (`git clone https://github.com/yaelahan/pretest.git`)
- Install dependencies (`go get`)
- Copy `.env.example` to `.env`
  - Set `APP_KEY` using secret key (used for signing jwt token)
  - Set `DATABASE_URI` using mysql server uri. Can use fresh database or imported database from this repository ([db.sql](https://github.com/yaelahan/pretest/blob/main/db.sql))

## Run The Project
- `go run main.go`

## Postman Collection
https://github.com/yaelahan/pretest/blob/main/pretest.postman_collection.json