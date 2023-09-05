# topx
X代表扩展性和卓越性能，这是一个多用途的网关项目

# 技术点

- request id: X-Request-Id
- Reta limit
- Metrics
- Trance

# 参考

- MongoDB

  ```bash
  # mongodb://admin:password@${ip}:27017/
  nerdctl run --name stonebird-mongodb -d -p 27017:27017 --restart=always -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=${password} mongo:7.0.0
  ```

- Redis

  ```bash
  nerdctl run --name stonebird-redis -d --restart=always -e REDIS_ARGS="--requirepass ${password}" -p 6379:6379 -p 8001:8001 redis:7.2.0
  ```

  
