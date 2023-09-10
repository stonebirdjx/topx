<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Actions](#actions)
  - [CreateActions](#createactions)
  - [ListActions](#listactions)
  - [DeleteActions](#deleteactions)
  - [GetAction](#getaction)
  - [UpdateAction](#updateaction)
  - [DeleteAction](#deleteaction)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Actions

path路径: /apis/apps/v1/actions


## CreateActions

- 请求方式：`POST`  /apis/apps/v1/actions 

- Content-Type：application/json

- Request Body:

  ```json
  {
      "actions":[
          {
              "name":"TestPing",
              "service_name":"iva",
              "description":"测试连通性",
              "rate_limit":3,
              "is_auth":false,
              "path":"/ping",
              "proxy":"http://127.0.0.1:8888",
              "timeout":30000,
              "version": "2023-09-09"
  
          }
      ]
  }
  ```

  | 名称         | 类型    | 描述                         |
  | ------------ | ------- | ---------------------------- |
  | name         | string  | action名称                   |
  | service_name | string  | action对应service的服务名称  |
  | rate_limit   | float64 | 限制速率，>= 1               |
  | description  | string  | action描述                   |
  | is_auth      | bool    | action是否需要鉴权           |
  | path         | string  | 后端服务路由                 |
  | proxy        | string  | 后端服务地址                 |
  | timeout      | int     | 超时时间,d单位ms,默认30000ms |
  | version      | string  | 对应版本号                   |

- Response body

  statusCode: `201`  -> 201Created

  ```bash
  {
      "message": "creat actions success"
  }
  ```

## ListActions

- 请求方式：`GET`  /apis/apps/v1/actions?page_size=1&page_num=1

  - query非必要参数

- Response body

  statusCode: `202`  -> 200 OK

  ```json
  {
      "actions": [
          {
              "id": "64fd53c8c66f4c7eda431af7",
              "name": "TestPing",
              "service_name": "iva",
              "description": "测试连通性1",
              "rate_limit": 1,
              "is_auth": false,
              "path": "/ping",
              "proxy": "http://127.0.0.1:8888",
              "timeout": 30000,
              "version": "2023-09-09"
          }
      ],
      "totals": 2
  }
  ```

  - totals : 数据库总数

##  DeleteActions

  - 请求方式：`POST`  /apis/apps/v1/actions 
  
  - Content-Type：application/json
  
  - Request Body:
  
    ```json
    {
        "ids":["64fc4afde9f3ad7ec777bf39","64fc4afde9f3ad7ec777bf3b"]
    }
    ```
  
   - Response body
  
     - statusCode: `200`  -> 200 OK
  
     ```json
     {
         "message": "delete success"
     }
     ```
     

## GetAction

- 请求方式：`GET`  /apis/apps/v1/actions/:actionid

- Response body

  - statusCode: `200`  -> 200 OK

  ```json
  {
      "id": "64fd53c8c66f4c7eda431af7",
      "name": "TestPing",
      "service_name": "iva",
      "description": "测试连通性1",
      "rate_limit": 1,
      "is_auth": false,
      "path": "/ping",
      "proxy": "http://127.0.0.1:8888",
      "timeout": 30000,
      "version": "2023-09-09"
  }
  ```

## UpdateAction

- 请求方式：`PUT |PATCH  `  /apis/apps/v1/actions/:actionid

  - Content-Type：application/json

  - Request Body:

    ```json
    {
        "name": "TestPing2",
        "service_name": "iva",
        "description": "测试连通性1",
        "rate_limit": 1,
        "is_auth": false,
        "path": "/ping",
        "proxy": "http://127.0.0.1:8888",
        "timeout": 30000,
        "version": "2023-09-09"
    }
    ```

- Response body

  - statusCode: `200`  -> 200 OK

  ```json
  {
      "message": "update success"
  }
  ```

## DeleteAction

- 请求方式：`delete` /apis/apps/v1/actions/:actionid

- Response body

  - statusCode: `200`  -> 200 OK

  ```json
  {
      "message": "delete success"
  }
  ```

  
