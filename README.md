# Backend challenge: Shorten-URL

![GO](https://img.shields.io/badge/GO-blue)
![SQLite3](https://img.shields.io/badge/SQLite-003B57?style=for-the-badge&logo=sqlite&logoColor=white)

This project it's an API using GO and SQLite3

## Table of Content

## Instalation

1. Clone the repository

```bash
git clone repo
```

2. Run the command

```bash
go mod tidy
```

## Usage

1. Start the application

```bash
make start
```

The API will be accessible at http://localhost:`8080`

## API Endpoints

The API provides the following endpoint:

### Create Shorten URL

```markdown
POST /shorten-url - Create and save the shorten url
```

**Body Request**

```json
{
  "url": "https://www.exemple.com"
}
```

**Body Response**

```json
{
  "url": "https://www.xxx.com/something"
}
```
