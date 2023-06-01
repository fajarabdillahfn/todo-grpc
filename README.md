# todo-grpc
simple todo app with gRPC implementations

## How to Run
1. git clone the project
2. go to the directory folder
3. create .env file with all parameters fulfilled (see example on .env_example)
4. `docker-compose up`

## Documentation
see the proto file (./proto/task.proto)

### Methods

1. **CreateTask**<br>
example:
* Request
```
{
  "description": "et consectetur in incididunt cupidatat",
  "title": "non esse pariatur et"
}
```
* Response
```
{
  "id": "3"
}
```

2. **GetTasks**<br>
example:
* Response
```
{
    "items": [
        {
            "id": "1",
            "title": "Hello",
            "description": "Hello",
            "is_completed": true,
            "UpdatedAt": {
                "seconds": "1685627958",
                "nanos": 0
            },
            "CreatedAt": {
                "seconds": "1685627936",
                "nanos": 0
            }
        },
        {
            "id": "2",
            "title": "",
            "description": "",
            "is_completed": false,
            "UpdatedAt": {
                "seconds": "1685634775",
                "nanos": 0
            },
            "CreatedAt": {
                "seconds": "1685634775",
                "nanos": 0
            }
        }
    ]
}
```

3. **GetTaskById**<br>
example:
* Request
```
{
    "id": "3"
}
```
* Response
```
{
    "id": "3",
    "title": "non esse pariatur et",
    "description": "et consectetur in incididunt cupidatat",
    "is_completed": false,
    "UpdatedAt": {
        "seconds": "1685634809",
        "nanos": 0
    },
    "CreatedAt": {
        "seconds": "1685634809",
        "nanos": 0
    }
}
```

4. **UpdateTask**<br>
example:
* Request
```
{
    "description": "tempor labore cupidatat proident",
    "id": "3",
    "is_completed": true,
    "title": "quis occaecat ea culpa"
}
```
* Response
```
{
    "id": "3",
    "title": "quis occaecat ea culpa",
    "description": "tempor labore cupidatat proident",
    "is_completed": true,
    "UpdatedAt": {
        "seconds": "1685634872",
        "nanos": 0
    },
    "CreatedAt": {
        "seconds": "1685634809",
        "nanos": 0
    }
}
```

5. **DeleteTask**<br>
example:
* Request
```
{
    "id": "3"
}
```
* Response
```
{
    "status": "task 3 deleted"
}
```
