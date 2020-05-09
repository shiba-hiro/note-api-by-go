# note-api-by-go

note-api-by-go is a sample web api project implemented by Go.

## Requirements

Run database  
You can refer [here](https://github.com/shiba-hiro/note-mysql).

## Usage

```
$ http POST localhost:1323/api/v1/notes title="My First Note" content="I started to take note."
HTTP/1.1 201 Created
Content-Length: 208
Content-Type: application/json; charset=UTF-8
Date: Sat, 09 May 2020 14:42:16 GMT

{
    "content": "I started to take note.", 
    "created_at": "2020-05-09T23:42:16.804679836+09:00", 
    "id": "00a2e8c6-ccc8-4413-a6ae-521099180dcd", 
    "title": "My First Note", 
    "updated_at": "2020-05-09T23:42:16.80468306+09:00"
}
```

```
$ http GET localhost:1323/api/v1/notes/00a2e8c6-ccc8-4413-a6ae-521099180dcd
HTTP/1.1 200 OK
Content-Length: 186
Content-Type: application/json; charset=UTF-8
Date: Sat, 09 May 2020 14:49:39 GMT

{
    "content": "I started to take note.", 
    "created_at": "2020-05-09T14:42:16.805Z", 
    "id": "00a2e8c6-ccc8-4413-a6ae-521099180dcd", 
    "title": "My First Note", 
    "updated_at": "2020-05-09T14:42:16.805Z"
}
```

```
$ http PUT localhost:1323/api/v1/notes/00a2e8c6-ccc8-4413-a6ae-521099180dcd title="My First Note (edited)" content="I started to take note. But it's difficult to continue everyday..."
HTTP/1.1 200 OK
Content-Length: 249
Content-Type: application/json; charset=UTF-8
Date: Sat, 09 May 2020 15:05:40 GMT

{
    "content": "I started to take note. But it's difficult to continue everyday...", 
    "created_at": "2020-05-09T14:42:16.805Z", 
    "id": "00a2e8c6-ccc8-4413-a6ae-521099180dcd", 
    "title": "My First Note (edited)", 
    "updated_at": "2020-05-10T00:05:40.333431457+09:00"
}
```

```
$ http DELETE localhost:1323/api/v1/notes/00a2e8c6-ccc8-4413-a6ae-521099180dcd
HTTP/1.1 204 No Content
Date: Sat, 09 May 2020 15:09:24 GMT
```

## Testing

```
$ go test -v github.com/shiba-hiro/note-api-by-go/...
```