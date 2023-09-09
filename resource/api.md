<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Actions](#actions)
  - [create actions](#create-actions)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Actions

path路径: /apis/apps/v1/actions

## create actions

- 请求方式：`POST`  /apis/apps/v1/actions 

- Content-Type：application/json

- Request Body:

  ```bash
  {
      "actions":[
          {
              "name":"TestPing",
              "service_name":"iva",
              "description":"测试连通性",
              "is_auth":false,
              "path":"/ping",
              "proxy":"http://127.0.0.1:8888",
              "timeout":30000,
              "version": "2023-09-09"
  
          }
      ]
  }
  ```

  - name：对应的action名称
  - service_name：对应服务的service名称
  - description：aciton描述
  - is_auth：是否需要鉴权后访问
  - path: 后端服务路由
  - proxy: 后端服务地址
  - timeout：超时时间ms
  - version：对应版本号

- Response body

  statusCode: `201`  -> 201Created

  ```bash
  {
      "message": "creat some new actions success"
  }
  ```

  ## List Acitons

  请求方式：`GET`  /apis/apps/v1/actions 