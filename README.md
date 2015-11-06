


## run

```sh
$ gom install
$ gom exec go build main.go
$ ssh-keygen -t rsa -b 2048 -C "mymail@example.com"
  Enter file in which to save the key (/Users/you/.ssh/id_rsa): ./key_for_token
$ DB_HOST=localhost DB_PORT=3306 MYSQL_USER=root MYSQL_PASS=password ./main
```

## for Docker

prepare mroonga's container.

```sh
$ git clone https://github.com/soundkitchen/docker-mroonga.git
$ cd docker-mroonga
$ docker build -t soundkitchen/mroonga .
```

build this container.

```sh
$ fab build_server
$ fab build_assets
$ cd ./dockerfile
$ ./init.sh
$ cd ../
```

and run

```sh
$ docker-compose up -d
```
