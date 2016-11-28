# mockserver
Fanach mock server.

# API
| Method        | Schema           | Response  |Description
| :--------- |:----------------:| :---------|:--------:
| GET | / | Fanach mock server | Homepage |
| POST | /api/ssaccount | `{ "server":"45.76.214.188","server_port":8392,  "password": "mock123",  "method": "aes-256-cfb"} `| Get mock server config. |
