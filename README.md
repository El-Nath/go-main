
# go-main

Simple stack implementation


## Installation
  ex : branch - go-main-simple-elas
  
  $ docker build --build-arg REPO=https://github.com/El-Nath/go-main.git --build-arg BRANCH=go-main-simple-elas -t elas:01 .

## Usage

  $ docker run -it -d --name sample_elas1 --restart=always -h elashost -p 80:1338 elas:01 
