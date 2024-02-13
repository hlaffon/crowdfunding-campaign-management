# crowdfunding-campaign-management

## Description

This project is a REST API providing statistical information about a user's campaign.

## Prerequisites

To run this project you need a PostGreSQL database running locally on port 5432.
Create a .env file at the root of the project and insert your DB credentials in the following environment variables:
DB_LOGIN=******
DB_PASSWORD=******
DB_NAME=******

### Data loading

Inside this DB you can run the initialization script
`cmd/init/init.sql`
to create the data tables.

Then to populate the database with the provided data set you need to run the following command:

`go run cmd/init/main.go
`

## Usage

To launch the API you just have to run the following command:

`go run cmd/server/main.go
`

## Documentation

Once the server is running, you can find all the available endpoints at this address:

http://localhost:3000/swagger/index.html

To regenerate the swagger doc, you just need to launch this command :

`swag init -g cmd/server/main.go --output docs/
`

(see how to install and full documentation here https://github.com/swaggo/swag)
