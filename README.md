# 简介

前端 - pnpm monorepo
- TypeScript
- [Vue.js](https://vuejs.org/)

文档
- [VitePress](https://vitepress.dev)

后端
- 基于 [go-blueprint](https://github.com/Melkeydev/go-blueprint) 初始化，文件目录参考 [project-layout](https://github.com/golang-standards/project-layout)
- web 框架使用 [Echo](https://echo.labstack.com/)

数据库
- [PostgreSQL](https://www.postgresql.org/)



## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```
Create DB container
```bash
make docker-run
```

Shutdown DB Container
```bash
make docker-down
```

DB Integrations Test:
```bash
make itest
```

Live reload the application:
```bash
make watch
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```
