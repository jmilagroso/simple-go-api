# Simple Go API

Simple Go API

  - Uses Posgresql (persistent storage)
  - Uses Redis (caching)
  - Uses JWT
  - Heroku based config

## Tech

Simple Go API uses a number of open source projects to work properly:

* [Go](https://golang.org/) - Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. 
* [Heroku](https://devcenter.heroku.com/articles/getting-started-with-go) - Heroku is a cloud platform that lets companies build, deliver, monitor and scale apps — we're the fastest way to go from idea to URL, bypassing all those infrastructure headaches.
* [Heroku Posgresql](https://www.heroku.com/postgres) - Managed SQL Database as a Service for all developers.
* [Heroku Redis](https://www.heroku.com/redis) - The key-value data store you love, with the developer experience you deserve.
* [go-pg](https://github.com/go-pg/pg) - Golang ORM with focus on PostgreSQL features and performance.
* [redigo](https://github.com/gomodule/redigo) - Redigo is a Go client for the Redis database.
* [JWT](https://jwt.io/) - JSON Web Tokens are an open, industry standard RFC 7519 method for representing claims securely between two parties.


## Check Security Issues using GoSpec
```sh
$ cd github.com/jmilagroso
$ ./check_issues.sh

# Output
...
[gosec] 2019/02/11 14:56:25 Checking package: tests
{
        "Issues": [],
        "Stats": {
                "files": 19,
                "lines": 662,
                "nosec": 0,
                "found": 0
        }
}%
```

## Running on Local

### Postgresql Add-on
```sh
$ heroku addons:create heroku-postgresql:hobby-dev
```

### Redis Add-on
```sh
$ heroku addons:create heroku-redis:hobby-dev
```

### Using `heroku local`
```sh
$ cd github.com/jmilagroso
$ go install ./...
$ api

OR

$ heroku local -f Procfile
```

### Setup .env
```sh
$ heroku config:get DATABASE_URL -s >> .env
$ heroku config:get REDIS_URL -s >> .env
```

### Running with `.env` and different port
```sh
$ heroku local web -e .env -p 3001
```
## API Usage
```sh
# Authenticate/Generate request token
curl -X POST \
  http://jaymilagroso-goapi.herokuapp.com/auth \
  -H 'cache-control: no-cache' \
  -F username=<username> \
  -F password=<password>


# Output
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTAxMjU0NTUsImlhdCI6MTU0OTg2NjI1NSwic3ViIjoiMTAwMDgifQ.XsH7pA1O8nDlD3yllFk19_eN6DWkLhV5X2xHbRSn0Ks",
    "expiration": 1550125455
}
```


```sh
# List Users Paginated (Token required)
curl -X GET \
  https://jaymilagroso-goapi.herokuapp.com/users/1/5 \
  -H 'X-Token: <token>' \
  -H 'cache-control: no-cache'

# Output
[
    {
        "id": "10008",
        "username": "jaymilagroso2",
        "email": "jaymilagroso2@gmail.com"
    },
    {
        "id": "10007",
        "username": "jaymilagroso1",
        "email": "jaymilagroso1@gmail.com"
    },
    {
        "id": "10003",
        "username": "johndoe10000",
        "email": "jonhdoe@e747ab3b8b484d4fc95c5ad1db172f2d@gmail.com"
    },
    {
        "id": "10002",
        "username": "johndoe9999",
        "email": "jonhdoe@6bff131bf32446e97d464502db4766af@gmail.com"
    },
    {
        "id": "10001",
        "username": "johndoe9998",
        "email": "jonhdoe@8efb1b7c88a9c341bb37d2f5d0e90df4@gmail.com"
    }
]
```

```sh
# Creates new User record (Token not required)
curl -X POST \
  http://jaymilagroso-goapi.herokuapp.com/user \
  -H 'cache-control: no-cache' \
  -F email=<email> \
  -F username=<username> \
  -F password=<password>

# Output
{
    "id": "",
    "username": "johndoe2019",
    "email": "johndoe2019@gmail.com"
}
```

```sh
# Get user by id (Token required)
curl -X GET \
  https://jaymilagroso-goapi.herokuapp.com/users/<user_id> \
  -H 'X-Token: <token>' \
  -H 'cache-control: no-cache'

# Output
{
    "id": "9999",
    "username": "johndoe9996",
    "email": "jonhdoe@3125da62ec6df8727faee392248875f7@gmail.com"
}
```

## References
- https://devcenter.heroku.com/articles/getting-started-with-go
- https://devcenter.heroku.com/categories/heroku-postgres
- https://devcenter.heroku.com/articles/heroku-redis


## License
This code is distributed using the Apache license, Version 2.0.

### Author
Jay Milagroso <j.milagroso@gmail.com>