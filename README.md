# DreamCraft 

To regist and login in DreamCraft!

Info:

The password is hashed with a random salt and stored in a slice


## Github

https://git-scm.com/download/win

Git commands:

```bash
git add --all
```

```bash
git commit -m "NAME"
```

```bash
git push
```

go mod:

```bash
go mod init Backend.go
```

```bash
go mod tidy
```

## Frontend

written in html with Typescript (https://www.typescriptlang.org/docs/handbook/intro.html)

![LoginScreen](/Frontend/images/DreamCraft_LoginScreen.png)

Client (Port for VSCode LiveServer)

```bash
http://localhost:5500/Frontend/html/Login.html
```

```bash
node Frontend.ts
```

Init typescript in the project (create tsconfig.json)

```bash
npx tsc --init
```

Start ts compiler command

```bash
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned
```

Login.ts to Login.js (for html page)

```bash
tsc ts/Login.ts
```


## API

JSON:

```bash
{
  "userName": "PostmanUser",
  "userPassword": "PostmanPW"
}
```


## Backend

written in golang (https://go.dev/)

![ServerConsoleScreen](/Frontend/images/DreamCraft_Server.png)

Server

```bash
http://localhost:8080/
```

Create mod.go

```bash
go mod init DreamCraft
```

Start backend

```bash
go run Backend.go
```



