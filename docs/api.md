# API Documentation

## User Management
- **`POST /api/users`**: Registers a new user. Expects `email` and `password`. Returns the user schema.
- **`POST /api/login`**: Authenticates a user securely. Expects `email` and `password`. Returns an Access Token and a Refresh Token structure.
- **`PUT /api/users`**: Allows a logged-in user to seamlessly update their email and password securely using their target JWT.

## Authorization
- **`POST /api/refresh`**: Trades a valid, untouched Refresh Token structure for a newly minted 1-Hour strict Access Token via the Authorization header.
- **`POST /api/revoke`**: Cryptographically revokes a Refresh Token internally in the database to structurally kill a session securely.

## Chirps
- **`POST /api/chirps`**: Creates a single chirp mechanically limited to 140 characters. Requires a valid JWT Authorization natively via Headers.
- **`GET /api/chirps`**: Retrieves the global timeline of all targeted chirps globally sorted chronologically ascending. Supports `?author_id=` query securely to gracefully filter exactly for one author.
- **`GET /api/chirps/{chirpID}`**: Retrieves precisely one chirp fundamentally by its universally unique UUID structure.
- **`DELETE /api/chirps/{chirpID}`**: Irreversibly deletes a specific chirp entirely. Mathematically requires an authentic JWT mathematically tied exactly to the structural string of its originally authenticated author UUID.

## External Webhooks
- **`POST /api/polka/webhooks`**: Secures automatic upgrades precisely for Chirpy Red premium accounts internally! Completely requires a structurally valid `Authorization: ApiKey ...` natively matching locally the globally hosted `POLKA_KEY`. Processes `user.upgraded` events silently.
