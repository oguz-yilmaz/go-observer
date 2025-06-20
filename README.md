## Event Observer

### Run the example

```bash
$ go run examples/httpserver/*.go
# Output:
# PLUGIN A handles ServerStarting with data:  [ServerInfo:] Starting server at http://localhost:8080
# PLUGIN B handles ServerStarting with data:  [ServerInfo:] Starting server at http://localhost:8080
# PLUGIN A handles ServerStarted with message: [ServerMessage:] Server Started successfully on port 8080
# Only Plugin A handles ServerStarted event
```

