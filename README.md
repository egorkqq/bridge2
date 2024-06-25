## Environment Variables Reference

This document outlines the environment variables used in the application along with their default values and descriptions.

### Verbosity
- **Description**: Specifies the verbosity level of logging.
- **Type**: Integer
- **Environment Variable**: `VERBOSITY`
- **Default Value**: `2`
  - `3` = debug
  - `2` = info
  - `1` = warn
  - `0` = error

### Addr
- **Description**: TCP address to listen to.
- **Type**: String
- **Environment Variable**: `LISTEN_ADDR`
- **Default Value**: `:8080`

### MetricsAddr
- **Description**: Metrics TCP address to listen to.
- **Type**: String
- **Environment Variable**: `METRICS_ADDR`
- **Default Value**: `:8081`

### TLS
- **Description**: Enable self-signed TLS.
- **Type**: Boolean
- **Environment Variable**: `TLS`
- **Default Value**: `false`

### CORS
- **Description**: Enable Cross-Origin Resource Sharing (CORS).
- **Type**: Boolean
- **Environment Variable**: `CORS`
- **Default Value**: `false`

### JsonLogs
- **Description**: Enable JSON logs output.
- **Type**: Boolean
- **Environment Variable**: `JSON_LOGS`
- **Default Value**: `false`

### HeartbeatSeconds
- **Description**: Heartbeat interval in seconds.
- **Type**: Unsigned Integer
- **Environment Variable**: `HEARTBEAT_SECONDS`
- **Default Value**: `10`

### MaxMessageTTL
- **Description**: Maximum message time-to-live (TTL) in seconds.
- **Type**: Unsigned Integer
- **Environment Variable**: `MAX_MESSAGE_TTL`
- **Default Value**: `300`

### HeartbeatGroups
- **Description**: Number of heartbeat groups (shards).
- **Type**: Unsigned Integer
- **Environment Variable**: `HEARTBEAT_GROUPS`
- **Default Value**: `10`

### PushRPS
- **Description**: Push Requests Per Second (RPS) limit.
- **Type**: Unsigned Integer
- **Environment Variable**: `PUSH_RPS_LIMIT`
- **Default Value**: `5`

### MaxSubscribersPerIP
- **Description**: Maximum parallel subscriptions per IP limit.
- **Type**: Unsigned Integer
- **Environment Variable**: `MAX_SUBSCRIBERS_PER_IP`
- **Default Value**: `100`

### MaxClientsPerSubscription
- **Description**: Maximum clients limit per subscription.
- **Type**: Unsigned Integer
- **Environment Variable**: `MAX_CLIENTS_PER_SUBSCRIPTION`
- **Default Value**: `100`

### WebhookURL
- **Description**: URL for webhook.
- **Type**: String
- **Environment Variable**: `WEBHOOK_URL`
- **Default Value**: Not set

### WebhookAuth
- **Description**: Bearer token to be sent in the Authorization header of webhook requests.
- **Type**: String
- **Environment Variable**: `WEBHOOK_AUTH`
- **Default Value**: Not set

### RedisAddr
- **Description**: Address of the Redis server.
- **Type**: String
- **Environment Variable**: `REDIS_ADDR`
- **Default Value**: `localhost:6379`

### RedisPassword
- **Description**: Password for the Redis server.
- **Type**: String
- **Environment Variable**: `REDIS_PASSWORD`
- **Default Value**: Not set

### RedisDB
- **Description**: Redis database number.
- **Type**: Integer
- **Environment Variable**: `REDIS_DB`
- **Default Value**: `0`

### Unlimited Tokens
- **Description**: You can set tokens to be unlimited. To do this, you need to fill out the file [unlimited_tokens.json](unlimited_tokens.json) in the format:

### Using the `.env` file
You can set the environment variables by creating a `.env` file in the root directory of your project with the following content:

```plaintext
REDIS_ADDR=localhost:6379
REDIS_PASSWORD=your-password
REDIS_DB=0