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
- make create_vendor

- make build

- ./builds/explorer-genesis-uploader -config=/etc/noah/config.json (This place is not important because we using ENV configuration)

### Important Environments

Example for all important environments you can see in file .env.example.
Its config for connect to PostgresSQL, Node API URL, Extender URL and service mode (debug, prod).

### Config file

Support JSON and YAML formats 

Example:

```
{
  "name": "Noah Explorer Genesis Uploader",
  "app": {
    "debug": true,
    "baseCoin": "NOAH",
    "txChunkSize": 200,
    "addrChunkSize": 30,
    "eventsChunkSize": 200
  },
  "workers": {
    "saveTxs": ME_WRK_SAVE_TXS,
    "saveTxsOutput": ME_WRK_SAVE_OUTPUT_TXS,
    "saveInvalidTxs": ME_WRK_SAVE_INVALID_TXS,
    "saveRewards": ME_WRK_SAVE_REWARDS,
    "saveSlashes": ME_WRK_SAVE_SLASHES,
    "saveAddresses": ME_WRK_SAVE_ADDR,
    "saveTxValidator": ME_WRK_SAVE_TX_VAL,
    "updateBalance": ME_WRK_UPD_BALANCE,
    "balancesFromNode": ME_WRK_BALANCE_NODE
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
  },
  "extenderApi": {
    "host": "localhost",
    "port": "9000"
  },
  "wsServer":{
    "isSecure" : false,
    "link" : "ME_WS_LINK",
    "port" : "ME_WS_PORT",
    "key"  : "ME_WS_KEY"
  }
}
```