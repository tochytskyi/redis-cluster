# redis-cluster
Make redis cluster with Go probabilistic cache implementation

1. Start all containers
`docker-compose up -d`

2. Try basic replication with `RDB` persistence mode
    - Set value to master node `docker exec -it go-redis-cluster_master redis-cli set testKey testValue`
    - Check that value `docker exec -it go-redis-cluster_master redis-cli get testKey`
    
        ```
        testValue
        ```
    - Check that slave has the same data `docker exec -it go-redis-cluster_slave redis-cli get testKey`

        ```
        testValue
        ```

3. To apply AOF persistence mode use `./conf/master_aof.conf` in `docker-compose.yml`

4. To change eviction policy use `maxmemory-policy` param in `./conf/*.conf` config files

