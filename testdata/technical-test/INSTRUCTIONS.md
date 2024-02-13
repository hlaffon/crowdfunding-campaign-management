# Technical test

## Introduction

First of all, thank you for trying our Go developer challenge!

The challenge proposed here is to build a simple stateless REST API.

We are building a new frontend which will replace the project owner dashboard currently it doesn't use an API to retrieve the data and it's a perfect state of the art server-side rendering.

The next frontend will rely on [React](https://reactjs.org/) to render the data from the API you will build. You can impose the data representation and the frontend team will adapt its codebase to it.

## What you should use

Software:

* PostgreSQL
* Go
* Redis (maybe you need it maybe not 😌)

Librairies:

* [gin-gonic](https://github.com/gin-gonic/gin) or stdlib
* [pgx](https://github.com/jackc/pgx)
* [scany](https://github.com/georgysavva/scany)
* [go-redis](https://github.com/go-redis/redis)

No ORM is allowed for this test, but if you need more don't hesitate.

## What we will look at

* If your code fulfills the requirements
* How you are handling errors and panics
* How you document your code and your README
* How you are implementing the storage
* How you are avoiding duplicate SQL queries, preloading SQL queries, etc.
* How you are dealing with heavy computation and response times
* How clean are your design and implementation. How easy it is to understand and maintain your code.
* Unit tests are not a requirement but it’s a plus.

## What we need

* Amount raised each day during the campaign
* Number of new contributors each day during the campaign
* Number of new contributions each day during the campaign
* Number of visits each day during the campaign
* Conversion rate each day between visits and contributions
* When did the project reach its milestones (10%, 20%, 30%, ..., 100%) and which contribution allow to reach the milestone
* Percentage reached each day during the campaign
* Average contribution amount

Since there is no authenticated user, your API will be open but if you want to provide your own authentication system, feel free to do so.

## What we provide

We are provding three different datasets that you will need to import in a PostgreSQL database with your own schema.

* `projects.csv`: Projects currently online on the platform
* `contributions.csv`: Contributions made during the funding period of each project
* `visits.csv`: Unique visitors, visits for each project

Be careful, you will have different projects with different volume and your database schema needs to work with different topography.

# ⚠️

This test should be done in a maximum of 2 weeks if you need more time don't hesitate to contact us.

It will not be re-used in production, our system is more complex than the datasets we are providing you so it cannot be used and we have already implemented a large amount of the specification. 
