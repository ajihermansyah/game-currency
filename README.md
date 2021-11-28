# game-currency
a set of APIs to be used by front-end engineers to develop an application for conversion game currency, create currency, create currency rate, show list currency, show list currency rate and convert between currencies, Tech stack : Go, Docker-Compose, MySQL, HTML, CSS

The API serve 5 endpoint routes.
at cur Collection [/order/{store_id}], Cart Collection [/cart], Tax Code Collection [/tax_code] .
- [api/game-currency/currencies] You can perform fetch List of currency [GET] from currency Collection
- [api/game-currency/currencies] You can add data currency [POST] to currency Collection
- [api/game-currency/conversion] You can perform fetch List of currency rate [GET] from conversion Collection
- [api/game-currency/conversion] You can add data currency rate [POST] to conversion Collection
- [api/game-currency/convert] You can convert between currency [POST]


See [API Documentation](https://github.com/ajiehermansyah/game-currency/blob/master/apiary.apib) on how to use it.

## Directory Structure
```
game-currency
    |--config                   - to initialize template and database
        |--db.go                - for initialize mysql database connection
        |--tpl.go               - for template view configuration
    |--controllers              - to store package controllers
        |--conversion.go        - to handle Conversions Collection [/api/game-currency/conversion]
        |--currency.go          - to handle Currency Collection [/api/game-currency/currencies]
        |--web.go               - to handle Web/Frontend for GUI game currency
    |--docker                   - Dockerfile for Golang at folder web, Dockerfile for MySQL at folder db
        |--db
            |--docker_file
        |--web
            |--docker_file
    |--models                    - to store package models for object and mysql query
        |--conversion.go         - for table conversions
        |--convert.go            - for define struct convert currency
        |--currency.go           - for table currency
        |--response_api.go       - for define struct API response
    |--mysql_init                - init Database game_currency
        |--game_currency.sql
    |--templates                 - to store view html file for golang *gohtml
    |--apiary.apib               - json file docs from APIARY for API DOCS
    |--screenshot                - screenshot game currency UI and database design
    |--docker-compose.yml        - for docker-compose service and config
    |--main.go                   

  
```

## Setup

**Steps**
1. git clone [https://github.com/ajihermansyah/game-currency.git](https://github.com/ajihermansyah/game-currency.git)
2. install docker and docker-compose 
3. open terminal and run docker-compose build (service are build), docker-compose up(builds, recreates, attaches to container for service), docker-compose down (stop containers and remove containers etc) See [Docker Documentation](https://docs.docker.com/compose/reference/build/) on how to use it.
4. now your server ready for http:localhost:1000/
