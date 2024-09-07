# bmonitor

`bmonitor` is a serverless component designed to collect simple metrics for web pages. It is built with [Fermyon](https://www.fermyon.com/). I use it in my [blog](https://gabefiori.vercel.app/).

## Running
To get started, make sure to check the [Fermyon docs](https://developer.fermyon.com/spin/v2/quickstart) as their CLI is required.

```sh
# Build the project
spin build

# Export the required environment variables
export SPIN_VARIABLE_PUBLIC_API_KEY="PUBLIC_KEY"
export SPIN_VARIABLE_PRIVATE_API_KEY="PRIVATE_KEY" 
export SPIN_VARIABLE_PUBLIC_CORS_ALLOWED_ORIGINS="*"

# Run the component with SQLite
spin up --sqlite @database/migration.sql
```

## Usage
### Insert a new metric.
```
curl -H "X-API-Key: PUBLIC_KEY" -X POST http://127.0.0.1:3000?url=someurl
```

### Retrieve metrics.
```
curl -H "X-API-Key: PRIVATE_KEY" http://127.0.0.1:3000
```

#### Output
```json
[
  {
    "id": 1,
    "url": "someurl",
    "access_count": 1,
    "last_accessed_at": "2024-09-07 12:19:31"
  }
]
```

## Tests
As of now, the Fermyon ecosystem does not handle tests very well, particularly with the [go-sdk](https://github.com/fermyon/spin-go-sdk). However, this limitation is manageable for a simple serverless component.
