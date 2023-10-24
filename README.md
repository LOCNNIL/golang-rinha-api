# RINHA BACKEND (Latter Participation)

## Description

This is my simple RESTful API implementation to be tested with the [Rinha de Backend 2023](https://github.com/zanfranceschi/rinha-de-backend-2023-q3/blob/main/INSTRUCOES.md) stress test script. The [Rinha de Backend](https://github.com/zanfranceschi/rinha-de-backend-2023-q3/tree/main) it's over at this time, but I decided to make this implementation just for fun.

## Building and starting the Book Management service

It's necessary to have `docker-cli` installed on the machine.
So run:

```bash
docker compose up -d
```

After a few seconds, the application should be up and running on port `8081`, or any other port configured in the `.env` file.

---------------

### Interacting with the API

The interactions of with this API consists of 4 endpoints:

- `POST /pessoas` – To create the "person" resource.
- `GET /pessoas/[:id]` – To consult a created resource with the previous request.
- `GET /pessoas?t=[:termo da busca]` – To make a search for people.
- `GET /contagem-pessoas` – Special endpoint to the end of the stress test see how much people were created on DB.

About the schema of the person resouce:
| Attibute | description|
| --- | --- |
| **apelido** | Mandatory, unique, string of up to 32 chars. |
| **nome** | Mandatory, string of up to 100 characters. |
| **nascimento** | Mandatory, string to date in format YYYY-MM-DD (year, month, day).
| **stack** | Optional, vector of strings with each element up to 32 characters.

#### Some examples of API interactions

- Inserting a person:

```bash
curl -X 'POST' \
  'http://localhost:8081/pessoas' \
  -H 'Content-Type: application/json' \
  -d '{
    "apelido" : "josé",
    "nome" : "José Roberto",
    "nascimento" : "2000-10-01",
    "stack" : ["C#", "Node", "Oracle"]
  }'
```

The response should be something like:

```json
{
    "id":"b9060d27-8278-4acb-8944-d18057ad7217",
    "apelido" : "josé",
    "nome" : "José Roberto",
    "nascimento" : "2000-10-01",
    "stack" : ["C#", "Node", "Oracle"]
}
```

(As the Id is generated ramdomly by google uuid library it probabily will not be the same.)

---------------

### Shutting everything down

After you have finished testing, you can simply stop and clean everything with
the following command:

```bash
docker compose down --volumes
```

Please make sure to only run this command in this project's directory as it is a
destructive operation that irreversibly deletes all volumes specified in the
`docker-compose` configuration.

---------------

### Stress test with gatling

As part of the rules of the [Rinha de Backend](https://github.com/zanfranceschi/rinha-de-backend-2023-q3/tree/main) a stress test will be done on the API endpoints:

- `POST /pessoas`
- `GET /pessoas/[:id]`
- `GET /pessoas?t=[:termo da busca]`
