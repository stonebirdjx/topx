# topx
X代表扩展性和卓越性能，这是一个多用途的网关项目

# 技术点

- 依赖注入: 减少对全局变量的依赖。
- 符合软件设计七大原则：开闭原则、单一职责原则、里氏替换原则、依赖倒置原则、接口隔离原则、迪米特法则、合成复用原则。
- 统一Request-Id: X-Request-Id。
- 系统限制QPS和接口限制QPS：Reta limit。
- 可观测性：Trace、Metrics、Logs。
- 可测试：interface mock设计且代码高覆盖率。
- SDK接入: topx-go-sdk。

# 参考

- MongoDB

  ```bash
  # mongodb://admin:password@${ip}:27017/
  nerdctl run --name stonebird-mongodb -d -p 27017:27017 --restart=always -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=${password} mongo:7.0.0
  ```

- Redis

  ```bash
  # redis://<user>:<pass>@localhost:6379/<db>
  nerdctl run --name stonebird-redis -d --restart=always -e REDIS_ARGS="--requirepass ${password}" -p 6379:6379 -p 8001:8001 redis:7.2.0
  ```

  
