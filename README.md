# 简介

前端 - pnpm monorepo
- TypeScript
- [Vue.js](https://vuejs.org/)

文档
- [VitePress](https://vitepress.dev)

后端
- web 框架使用 [Echo](https://echo.labstack.com/)

数据库
- [PostgreSQL](https://www.postgresql.org/)


insp


# FAQ
Q: panic: http server error: listen tcp :8080: bind: Only one usage of each socket address (protocol/network address/port) is normally permitted.

A: 检查端口是否可用
```bash
netstat -aon|findstr "8080"
netsh int ipv4 show excludedportrange protocol=tcp

telnet 127.0.0.1 8080 
```

# Credits
All credits go to these open-source works and resources

- [go-blueprint](https://github.com/Melkeydev/go-blueprint) - Go-blueprint allows users to spin up a quick Go project using a popular framework
- [vitesse](https://github.com/antfu-collective/vitesse) - 🏕 Opinionated Vite + Vue Starter Template