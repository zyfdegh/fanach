# mockserver
Fanach mock server.

# API
| Method        | Schema           | Response  |Description
| :--------- |:----------------:| :---------|:--------:
| GET | / | Fanach mock server | Homepage |
| GET | /api/ssaccount | `{ "server":"45.76.214.188","server_port":8392,  "password": "incorrect",  "method": "aes-256-cfb"} `| Get mock server config. |
