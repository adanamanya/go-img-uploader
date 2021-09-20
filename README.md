# Image uploader made with love, Firebase, and GoLang(gin)
### This project was created by purpose of learning GoLang

[![N|Solid](https://i.imgur.com/CZgZT75.png)](https://nodesource.com/products/nsolid)

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

## How to run with Docker (Ez way)

- Clone this Repo.
- Fill the **.env** and **key.json** based on **.envExample** and **keyExample.json**
- run this command 
```sh
docker-compose up
```
- Voila!! your services is ready !!.

## How to run with Gin (live-reload)
what is **Gin**?
`gin` is a simple command line utility for live-reloading Go web applications.
Just run `gin` in your app directory and your web app will be served with
`gin` as a proxy. `gin` will automatically recompile your code when it
detects a change. Your app will be restarted the next time it receives an
HTTP request.

for reference : https://github.com/codegangsta/gin
step to run with gin
- Clone this Repo.
- Fill the **.env** and **key.json** based on **.envExample** and **keyExample.json**
- run this command 
```sh
gin main.go --appPort 8080
```
- Voila!! your services is ready !!.


## How to run with Go run (Classic way)

- Clone this Repo.
- Fill the **.env** and **key.json** based on **.envExample** and **keyExample.json**
- run this command 
```sh
go run main.go  
```
- Voila!! your services is ready !!.




