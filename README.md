<p align="center" style="text-align: center;">
    <a href="https://github.com/noah-blockchain/explorer-genesis-uploader/blob/master/LICENSE">
        <img src="https://img.shields.io/packagist/l/doctrine/orm.svg" alt="License">
    </a>
    <img alt="undefined" src="https://img.shields.io/github/last-commit/noah-blockchain/explorer-genesis-uploader.svg">
</p>

# NOAH Explorer Genesis Uploader

The official repository of Noah Explorer Genesis Uploader service.

Noah Explorer Genesis Uploader is a service which provides to upload primary network state data to Noah Explorer database after network reset or first start.

## RUN

- make build

- ./builds/explorer-genesis-uploader -config=/etc/noah/config.json 

### Config file

Support JSON and YAML formats 

Example:

```
{
  "name": "Noah Explorer Genesis Uploader",
  "app": {
    "debug": true,
    "baseCoin": "MNT",
    "txChunkSize": 200,
    "addrChunkSize": 30,
    "eventsChunkSize": 200
  },
  "database": {
    "host": "localhost",
    "name": "explorer",
    "user": "noah",
    "password": "password",
    "minIdleConns": 10,
    "poolSize": 20
  },
  "noahApi": {
    "isSecure": false,
    "link": "localhost",
    "port": 8841
  }
}
```