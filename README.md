# telegram-exchange-bot

## Docker build

````
docker build . --build-arg VERSION=test --tag telegram-exchange-bot
````

## Docker run

````
docker run --env EXCHANGE_URL= --env EXCHANGE_TOKEN= --env TELEGRAM_TOKEN= telegram-exchange-bot
````

#### Show images
````
docker images
````
#### Show containers
````
docker ps -a
````
#### Stop container
````
docker stop name
````