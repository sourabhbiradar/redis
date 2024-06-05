# redis

`redis` : in-memory data-structure store used as a database, cache, and message broker.

redis creates `temporary DB in memory` allowing for `very fast read and write operations`.

no need of mongoDB,MySql etc.

used for `caching`

can be configured to priodically dump data to disk

Supports : 

`Data-structures` : strings, hashes, lists, sets, sorted sets, bitmaps, hyperloglogs, geospatial indexes, and streams

`Pub/sub messaging` : publish / subscribe messsaging ; which can be used for event notifications and messaging systems.

`Lua scripting` : providing a way to run complex operations atomically.

`Clustering` : Redis installation where data is automatically sharded across multiple Redis nodes.

`Modules` : Redis can be extended with modules to add new capabilities or integrate with other systems.


`cmd` : brew install redis
`cmd` : brew services start redis

// linux
`cmd` :

sudo apt-get update

sudo apt-get install redis-server

sudo systemctl enable redis-server.service

sudo systemctl start redis-server.service

// redis cmds

`redis-cli` : runs redis cli

`INFO Keyspace` : information about keyspaces (DBs) & number of keys in each DB.

`SELECT db_number` : switch to specific DB (pass integer)

`DBSIZE` : number of keys in DB

NOTE : to see changes in redis DB use redis GUI















