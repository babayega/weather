# Weather Update Server

The Weather Update Server is a Go-based HTTP server built using the Gin framework. It provides an endpoint for users to submit weather updates. The server includes two middlewares: one for verifying the signature of the request to ensure authentication and another for rate limiting the number of requests per user. The server leverages the Geth library for signature verification.

## Features

1. Accepts weather update requests from users via an HTTP server.
2. Includes two middlewares:
    - Signature Verification Middleware: Verifies the signature of the request to ensure authentication using the Geth library.
    - Rate Limiting Middleware: Limits the number of requests per user to prevent abuse.
3. Rate limiting middleware uses a mutex to handle simultaneous requests from the same user.
4. Authenticates requests by verifying the signature using the Geth library.
5. Adds the weather value from the request to a channel for real-time updates.
6. The channel updates the weather value in real time and ensures concurrency handling as it can only process one request at a time.

## Architecture

Here's a high-level block diagram of the system architecture:

     +------------------------+
     |                        |
     |    Weather Update      |
     |        Server          |
     |                        |
     +------------------------+
                 |
                 | HTTP Requests
                 |
     +------------------------+
     |                        |
     |   Signature            |
     |   Verification         |
     |   Middleware           |
     |                        |
     +------------------------+
                 |
                 | Verified Requests
                 |
     +------------------------+
     |                        |
     |   Rate Limiting         |
     |   Middleware            |
     |                        |
     +------------------------+
                 |
                 | Processed Requests
                 |
     +------------------------+
     |                        |
     |   Weather Value        |
     |   Channel              |
     |                        |
     +------------------------+
                 |
                 | Real-Time Weather Updates
                 |
     +------------------------+
     |                        |
     |   Weather Data         |
     |   Processing           |
     |                        |
     +------------------------+
                 |
                 | Updated Weather Data
                 |
     +------------------------+
     |                        |
     |   Client Response       |
     |                        |
     +------------------------+


The Weather Service follows a layered architecture, consisting of the following components:

1. **Verification Middleware**: This middleware verifies the signature and message from the request header using the geth client. It ensures the authenticity of the request before further processing.
2. **Rate Limit Middleware**: The rate limit middleware checks whether the requesting address exceeds the limit of one request per 12 seconds. Additionally, it verifies that the request was fired within 2 seconds of the 12-second interval. A mutex is utilized to handle simultaneous requests accurately.
3. **Handler**: The handler is responsible for extracting the weather value from the request body. It makes a contract call through the chain service to validate whether the requesting address is currently registered. The handler utilizes the weather service to add the weather value to the processing channel.
4. **Weather Service**: The weather service consumes the values from the channel and updates the latest weather in the system. It processes the weather values one by one, ensuring concurrency and real-time updates.

## Lifecycle of a request

1. The request is received by the server and passed through the verification middleware, which validates the request's signature and message using the geth client.
2. Next, the rate limit middleware ensures that the requesting address does not exceed the maximum limit of one request per 12 seconds. It also checks if the request was fired within 2 seconds of the 12-second interval. A mutex is employed to handle simultaneous requests correctly.
3. After passing the middleware, the request reaches the handler. The handler extracts the weather value from the request body. It verifies the requesting address's registration status through a contract call using the chain service. The handler then adds the weather value to the weather service's processing channel.
4. The weather service continuously consumes the weather values from the channel. It performs the necessary processing and updates the latest weather in the system. Concurrency is maintained by processing the values one by one, ensuring real-time updates.


## Enhancements

1. **Event-based Address Registration**: Instead of making a contract call every time to check the registration status, a separate service can listen to events emitted by the contract. This service can maintain an updated record of registered addresses, which can be queried efficiently by the Weather Service to validate address registration.
2. **Persistent Data Storage**: To keep a record of successful requests, a database can be integrated into the system. Storing the details of processed requests can provide valuable insights and enable further analysis.