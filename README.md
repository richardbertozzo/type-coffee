<h1 align="center">
  <a href="https://github.com/richardbertozzo/type-coffee" title="type coffee doc">
    <img alt="logo" src="https://github.com/richardbertozzo/type-coffee/blob/master/data/logo.png" width="200px" height="200px" />
  </a>
  <br />
  The Best Type of Coffee
</h1>


[![ci](https://github.com/richardbertozzo/type-coffee/actions/workflows/ci.yml/badge.svg)](https://github.com/richardbertozzo/type-coffee/actions/workflows/ci.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/richardbertozzo/type-coffee)](https://goreportcard.com/report/github.com/richardbertozzo/type-coffee) [![codecov](https://codecov.io/gh/richardbertozzo/type-coffee/branch/master/graph/badge.svg?token=G667W1L2O5)](https://codecov.io/gh/richardbertozzo/type-coffee)

Get your right type of coffee ☕

Go application that provides information on the best type of coffee based on user input characteristics. 
It uses two data sources: the **Google Gemini API** and a **database** populated by ETL from a dataset. 
This dataset was gathered from [Kaggle](https://www.kaggle.com/datasets/volpatto/coffee-quality-database-from-cqi) (kudos for the author), see there for further details.

## Usage

To run the application, you need to have:
- Go installed
- Google Gemini API key, you can create one [here](https://aistudio.google.com/app/apikey)
- Postgres URL (_optional_)
  - You can run a local docker instance by running the command `$ make pg-up` or pointing to any other instance, like a cloud instance

The application has tree command entry points:

### ETL

This entrypoint has the responsibility to get the [Coffee CSV dataset](./data), convert the rows and 
then insert them into the Postgres database.

It will create and populate the second data source to retrieve the best type of coffee for you, 
which you can get using the other 2 commands, **API** and **CLI**. 

This is optimal if only want the Gemini API result, so skip it. 

You can run the ETL by the command:

```shell
go run cmd/etl/main.go --DATABASE_URL <database_url>
```

### API

Set up your local `.env` file using the file `.env.example`. Then you can start the HTTP API, run the following command:

```shell
make run
```

This will start the API server on port `3000`.
You can then find the full OpenAPI/Swagger spec definition [here](./api/openapi.yaml), or you can send a request to the following endpoint:

`GET /v1/best-coffees`: Returns the best type of coffee based on user given characteristics.

### CLI
To start the command line interface, run the following command:

```shell
go run cmd/cli/main.go --CHAT_GPT_KEY <your_chat_gpt_key> --DATABASE_URL <database_url>
```

This will run the CLI interface, which will output the best coffees for you from the selected characteristics.

## Data Sources

It uses two data sources:

### Gemini API
The Gemini API is used to generate responses to user feedback. The application sends user feedback to the API and receives a response based on the user's input.

### Database
The application also uses a database that is populated by ETL (Extract, Transform, Load) dataset data.

## License
The Best Type of Coffee is released under the Apache License. See [LICENSE](./LICENSE) for details.
