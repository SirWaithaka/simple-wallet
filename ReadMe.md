### Database Installation

The application uses postgres as the database server. So here are instructions on how to setup postgresql
on your machine using docker.

Get the official postgres docker image.
```bash
$ docker pull postgres
``` 

Then create a container from the image with the following variables
```bash
$ docker create \
--name wallet-db \
-e POSTGRES_USER=wallet \
-e POSTGRES_PASSWORD=wallet \
-p 5432:5432 \
postgres
```

Run the following command to start the container
```bash
$ docker start wallet-db
```

