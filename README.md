## Backend Assessment - Task 1
## Answer
Bad reviews during our 12.12 event occurred because many customers were shopping, checking out, and making payments; however, the inventory quantity was not accurately matched or reported, which led to negative inventory quantities and order cancellations due to stock unavailability. To prevent this, several steps should be taken:
* The system should be able to provide information on safety stock. Safety stock can prevent inventory shortages when market demand is uncertain. Safety stock can be calculated by considering maximum daily sales, maximum lead time, and average sales.
* The system should ensure ACID compliance—Atomicity, Consistency, Isolation, and Durability. When many customers make transactions simultaneously, multiple requests can occur at the same time, which may lead to race conditions when accessing product data. To prevent this, implementing database transactions can help maintain data integrity.

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

