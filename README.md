# Inventory
## Run Locally
Clone the project

```bash
  git clone https://github.com/androsyz/inventory
```

Go to the project directory

```bash
  cd inventory
```

## Setup
Before you run the compose file make sure you already change the ```.env``` file.
```bash
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=inventory
DB_PORT=5432
DB_HOST=host.docker.internal
```

## Start the server
Before running the command, make sure you already install docker on you computer.

You can refer to this link for detail about the installation : [Docker Installation](https://docs.docker.com/engine/install/)

You can run the server through this command :
```bash
make build
```

After you run the command above, all container will run in your local computer,
included :
- Inventory API
- PostgresSQL

After that you can access all endpoint via **http://localhost:{*your_port*}**


## Endpoints
All endpoints available in postman collection file. You can see  ```docs``` folder. For open the file, you can use [postman](https://www.postman.com/). <br>
To open the collection file, you can follow this step :
```bash
1. open postman
2. open menu file
3. choose the collection file in docs folder
4. run the server and you can try the endpoint via postman.
```

## Test
To run unit testing, you can run it via this command : 
```bash 
make test
```

