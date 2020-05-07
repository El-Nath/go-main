
# go-main

Simple stack implementation


## Installation
  ex : branch - go-main-simple-pqsql

  $ docker build --build-arg REPO=https://github.com/El-Nath/go-main.git --build-arg BRANCH=go-main-simple-pgsql -t pgsql:01 .

## Usage
usual :
  $ docker run -it -d --name sample_pqsql1 --restart=always -h psqlhost -p 80:1441 pgsql:01

or use your local network :
  $ docker run -it -d --name sample_pqsql1 --network="host" --restart=always -h psqlhost -p 80:1441 pgsql:01
