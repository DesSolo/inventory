## Inventory
Easy inventory client and server
## Inventory-Client
Warning!!!  
Please assemble the client yourself. You must specify secret keys and the server address.
```shell
make build-client USERNAME=user PASSWORD=passwd SERVER_URL=http://127.0.0.1:8080/upload
```
`inventory-client-linux`  
`inventory-client-windows.exe`
### keys
-s silent mode (not need write WH)

## Inventory-Server
`inventory-server-linux`

### env variables
|Kev|Default|Example|Description|
|---|---|---|---|
|LISTEN_ADDRESS|_no_|_:8080_ or _127.0.0.1:8080_|server listen address|
|INVENTORY_TOKEN|_no_|_mySuperSecretToken_|secret token must be the same for clients and on the server|
|STORAGE_TYPE|_memory_|see more TODO|data storage types|

### supported storage types
_memory_ all data store on memory. Attention: Will be destroyed after reboot!!!

_redis_ TODO

_sqlite3_ TODO