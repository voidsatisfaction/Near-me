# Near me server

## Tech Stack

1. Revel(Go)
2. Docker
3. Heroku

## Deployment

Docker + Heroku

### How to Deploy

1. Commentize `EXPOSE` on Dockerfile
2. `heroku container:push web`

### Refer this

[Container Registry & Runtime - heroku](https://devcenter.heroku.com/articles/container-registry-and-runtime#dockerfile-commands-and-runtime)

## How to start on local environment

### Prerequisite

- Docker
- .env(For YAHOO API and DOORKEEPER API KEYS setting)

### Dev Start

1. Download near_me_server and near_me_client on `$GOPATH/src`
2. Cd to near_me_server
3. `docker-compose up`
4. Enter localhost:13000
5. Fix codes on local environment(Codes are connected to dockerized app)

- React hot reloading is supported
- Revel hot reloading is supported as well

## Used api

- HTML5 Geolocation api
- Yahoo Geolocation api
- Google Map api
- Connpass api
- Doorkeeper api
