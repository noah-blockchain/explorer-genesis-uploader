
[![last commit](https://img.shields.io/github/last-commit/noah-blockchain/explorer-genesis-uploader.svg)]()
[![license](https://img.shields.io/packagist/l/doctrine/orm.svg)](https://github.com/noah-blockchain/explorer-genesis-uploader/blob/master/LICENSE)
[![version](https://img.shields.io/github/tag/noah-blockchain/explorer-genesis-uploader.svg)](https://github.com/noah-blockchain/explorer-genesis-uploader/releases/latest)
[![](https://tokei.rs/b1/github/noah-blockchain/explorer-genesis-uploader?category=lines)](https://github.com/noah-blockchai/explorer-genesis-uploader)

# NOAH Explorer Genesis Uploader

The official repository of Noah Explorer Genesis Uploader service.

Noah Explorer Genesis Uploader is a service which provides to upload primary network state data to Noah Explorer database after network reset or first start.

## BUILD

- make create_vendor
- make build

## Configure Extender Service from Environment (example in .env.example)
1) Set up connect to PostgresSQL Databases.
2) Set up connect to Node which working in non-validator mode. 
3) Set up connect to Extender service. 

## RUN
./uploader

_We recommend use our official docker image._
### Important Environments
Example for all important environments you can see in file .env.example.
Its config for connect to PostgresSQL, Node API URL, Extender URL and service mode (debug, prod).
