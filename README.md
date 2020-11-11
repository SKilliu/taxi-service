### TAXI SERVICE

There are three roles in the service: client, driver and operator. There are next features for these roles:

- *sing in / sing up user (all roles)*
- *edit user profile (all roles)*
- *upload user profile image (all roles)*
- *get all available orders (driver only)*
- *create a new car and user with any role (operator only)*
- *create a new order (client only)*
- *update order status to in_progress or closed (driver only)*

**Pre-requirements**

- Go v1.11.4+
- Docker version 18.09.06+

#### **How to run**

You need to execute next commands

`make compose-up`

`make migration-u`

`make run`

#### Commands list

- make docs-generate -- *generate the swagger docs*
- make compose-up -- *up minio and postgreSQL in docker containers*
- make build -- *build the project*
- make run -- *run the project*
- migration-up -- *upload migrations to db*
- migration-down -- *rollback all migrations*

#### **Swagger documentation url**

{{baseUrl}}/swagger/index.html