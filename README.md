# Rickandmorty API
This is an API service for retrieving episodes of the **Rick and Morty** series. The following endpoints are available to interact with episode data.

## Base Path
The URL Base for this API is: http://localhost:8080/api/deuna-rickandmorty-api/v1

## Endpoints

| Método | Endpoint                     | Descripción                             | Parámetros                       | Respuesta         | Código de Estado |
|--------|------------------------------|-----------------------------------------|-----------------------------------|-------------------|------------------|
| `GET`  | `/health`                    | Check service status                    | N/A                               | `200 OK`          | 200              |
| `GET`  | `/episodes`                  | Get all episodes                        | N/A                               | Array of episodes | `200 OK`  |
| `GET`  | `/episodes/:id`              | Get episode by ID                       | `id` (Path Parameter)             | Episodes          | `200 OK`, `404 Not Found` |
| `GET`  | `/episodes/multiple?ids=1,2,3,4`     | Get multiple episodes and sorted by IDs | `ids` (Query Parameter, lista de IDs separados por comas) | Array of episodes | `200 OK`, `400 Bad Request` |


## To start the api:
```shell
docker-compose up
```

## To test the connection:
```shell
curl -X GET "http://localhost:8080/api/deuna-rickandmorty-api/v1/episodes/multiple?ids=3,5,2,1"
```

You should see
```json
[
  {
    "id": 1,
    "name": "Pilot",
    "air_date": "December 2, 2013",
    "episode": "S01E01",
    "characters": [
      "https://rickandmortyapi.com/api/character/1",
      "https://rickandmortyapi.com/api/character/2",
      "https://rickandmortyapi.com/api/character/35",
      "https://rickandmortyapi.com/api/character/38",
      "https://rickandmortyapi.com/api/character/62",
      "https://rickandmortyapi.com/api/character/92",
      "https://rickandmortyapi.com/api/character/127",
      "https://rickandmortyapi.com/api/character/144",
      "https://rickandmortyapi.com/api/character/158",
      "https://rickandmortyapi.com/api/character/175",
      "https://rickandmortyapi.com/api/character/179",
      "https://rickandmortyapi.com/api/character/181",
      "https://rickandmortyapi.com/api/character/239",
      "https://rickandmortyapi.com/api/character/249",
      "https://rickandmortyapi.com/api/character/271",
      "https://rickandmortyapi.com/api/character/338",
      "https://rickandmortyapi.com/api/character/394",
      "https://rickandmortyapi.com/api/character/395",
      "https://rickandmortyapi.com/api/character/435"
    ],
    "url": "https://rickandmortyapi.com/api/episode/1",
    "created": "2017-11-10T12:56:33.798Z"
  },
  {
    "id": 2,
    "name": "Lawnmower Dog",
    "air_date": "December 9, 2013",
    "episode": "S01E02",
    "characters": [
      "https://rickandmortyapi.com/api/character/1",
      "https://rickandmortyapi.com/api/character/2",
      "https://rickandmortyapi.com/api/character/38",
      "https://rickandmortyapi.com/api/character/46",
      "https://rickandmortyapi.com/api/character/63",
      "https://rickandmortyapi.com/api/character/80",
      "https://rickandmortyapi.com/api/character/175",
      "https://rickandmortyapi.com/api/character/221",
      "https://rickandmortyapi.com/api/character/239",
      "https://rickandmortyapi.com/api/character/246",
      "https://rickandmortyapi.com/api/character/304",
      "https://rickandmortyapi.com/api/character/305",
      "https://rickandmortyapi.com/api/character/306",
      "https://rickandmortyapi.com/api/character/329",
      "https://rickandmortyapi.com/api/character/338",
      "https://rickandmortyapi.com/api/character/396",
      "https://rickandmortyapi.com/api/character/397",
      "https://rickandmortyapi.com/api/character/398",
      "https://rickandmortyapi.com/api/character/405"
    ],
    "url": "https://rickandmortyapi.com/api/episode/2",
    "created": "2017-11-10T12:56:33.916Z"
  },
  {
    "id": 3,
    "name": "Anatomy Park",
    "air_date": "December 16, 2013",
    "episode": "S01E03",
    "characters": [
      "https://rickandmortyapi.com/api/character/1",
      "https://rickandmortyapi.com/api/character/2",
      "https://rickandmortyapi.com/api/character/12",
      "https://rickandmortyapi.com/api/character/17",
      "https://rickandmortyapi.com/api/character/38",
      "https://rickandmortyapi.com/api/character/45",
      "https://rickandmortyapi.com/api/character/96",
      "https://rickandmortyapi.com/api/character/97",
      "https://rickandmortyapi.com/api/character/98",
      "https://rickandmortyapi.com/api/character/99",
      "https://rickandmortyapi.com/api/character/100",
      "https://rickandmortyapi.com/api/character/101",
      "https://rickandmortyapi.com/api/character/108",
      "https://rickandmortyapi.com/api/character/112",
      "https://rickandmortyapi.com/api/character/114",
      "https://rickandmortyapi.com/api/character/169",
      "https://rickandmortyapi.com/api/character/175",
      "https://rickandmortyapi.com/api/character/186",
      "https://rickandmortyapi.com/api/character/201",
      "https://rickandmortyapi.com/api/character/268",
      "https://rickandmortyapi.com/api/character/300",
      "https://rickandmortyapi.com/api/character/302",
      "https://rickandmortyapi.com/api/character/338",
      "https://rickandmortyapi.com/api/character/356"
    ],
    "url": "https://rickandmortyapi.com/api/episode/3",
    "created": "2017-11-10T12:56:34.022Z"
  },
  {
    "id": 5,
    "name": "Meeseeks and Destroy",
    "air_date": "January 20, 2014",
    "episode": "S01E05",
    "characters": [
      "https://rickandmortyapi.com/api/character/1",
      "https://rickandmortyapi.com/api/character/2",
      "https://rickandmortyapi.com/api/character/38",
      "https://rickandmortyapi.com/api/character/41",
      "https://rickandmortyapi.com/api/character/89",
      "https://rickandmortyapi.com/api/character/116",
      "https://rickandmortyapi.com/api/character/117",
      "https://rickandmortyapi.com/api/character/120",
      "https://rickandmortyapi.com/api/character/175",
      "https://rickandmortyapi.com/api/character/193",
      "https://rickandmortyapi.com/api/character/238",
      "https://rickandmortyapi.com/api/character/242",
      "https://rickandmortyapi.com/api/character/271",
      "https://rickandmortyapi.com/api/character/303",
      "https://rickandmortyapi.com/api/character/326",
      "https://rickandmortyapi.com/api/character/333",
      "https://rickandmortyapi.com/api/character/338",
      "https://rickandmortyapi.com/api/character/343",
      "https://rickandmortyapi.com/api/character/399",
      "https://rickandmortyapi.com/api/character/400"
    ],
    "url": "https://rickandmortyapi.com/api/episode/5",
    "created": "2017-11-10T12:56:34.236Z"
  }
]
```

## Perform unit test:
```shell
make test.unit

```

## Perform e2e test:
```shell
make test.e2e
```

## Docs
You can try the api directly from swagger

Swagger Docs URL:  http://localhost:8080/api/deuna-rickandmorty-api/v1/docs/index.html

![Imagen local](./images/swagger.jpg)

## Traces
See Traces: http://localhost:16686/

![Imagen local](./images/jaeger_search.jpg)

![Imagen local](./images/trace_sample.jpg)
