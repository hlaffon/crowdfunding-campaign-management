# crowdfunding-campaign-management

## Description

This project is a REST API providing statistical information about a user's campaign.

## Prerequisites

To run this project you need a PostGreSQL database running locally on port 5432.
Create a .env file at the root of the project and insert your DB credentials in the following environment variables:
```
DB_LOGIN=******
DB_PASSWORD=******
DB_NAME=******
```

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

You can test the endpoints on the generated swagger UI or use curl requests or any http client provider such as postman.
The projects ids can be found in the projects.csv file (in testdata/technical-test/ folder).

To regenerate the swagger doc, you just need to launch this command :

`swag init -g cmd/server/main.go --output docs/
`

(see how to install and full documentation here https://github.com/swaggo/swag)

## Features

Here is the list of provided features:
* Amount raised each day during the campaign
* Number of new contributors each day during the campaign
* Number of new contributions each day during the campaign
* Number of visits each day during the campaign
* Conversion rate each day between visits and contributions
* When did the project reach its milestones (10%, 20%, 30%, ..., 100%) and which contribution allow to reach the milestone
* Percentage reached each day during the campaign
* Average contribution amount