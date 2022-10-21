## Pretest
Pretest Project

## Run on local machine
- Clone this repository (`git clone https://github.com/yaelahan/pretest.git`)
- Install dependencies (`go mod download`)
- Copy `.env.example` to `.env`
  - Set `APP_KEY` using secret key (used for signing jwt token)
  - Set `DATABASE_URI` using mysql server uri. Can use fresh database or imported database from this repository ([db.sql](https://github.com/yaelahan/pretest/blob/main/db.sql))
- `go run main.go`

## Run using Docker
- Clone this repository (`git clone https://github.com/yaelahan/pretest.git`)
- `docker compose up -d`

## Postman Collection
https://github.com/yaelahan/pretest/blob/main/pretest.postman_collection.json