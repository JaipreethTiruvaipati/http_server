# Chirpy API

## What is Chirpy?
Chirpy is a highly performant, fully-functional backend HTTP API server written purely in Go. It serves as the underlying web infrastructure for a microblogging social network ("Chirpy"), providing a full suite of API endpoints to handle creating users, securing encrypted login sessions via Dual-Token architectures, managing dynamic chirp timelines, and intercepting verified third-party payment webhooks from Polka.

## Why Should You Care?
If you are looking for an incredibly clean, aggressively structured Go web server template, you've found it! This project is a masterclass in raw standard-library HTTP handling. It intentionally avoids messy external frameworks (like Gin or Fiber) to functionally demonstrate how to natively securely hash passwords with Argon2, negotiate JWT/Refresh-Token database states globally, and securely branch URL logic manually.

## How to Install & Run
This project relies natively on a local PostgreSQL database mapping strictly into Go.

1. **Clone the Repository** and make sure you have Go installed on your machine.
2. **Configure Your Environment**:
   Create a local `.env` file in the root of the project with the following cryptographic and routing keys:
   ```env
   DB_URL="postgres://username:password@localhost:5432/chirpy"
   PLATFORM="dev"
   JWT_SECRET="YourSuperSecretJWTSigningKey64BytesLongHere"
   POLKA_KEY="YourOfficialPolkaAPIKeyForWebhooks"
   ```
3. **Execute standard database migrations**:
   If you have Goose installed (`go install github.com/pressly/goose/v3/cmd/goose@latest`), seamlessly structure your PSQL mappings:
   ```bash
   cd sql/schema
   goose postgres <YOUR_DB_URL> up
   cd ../../
   ```
4. **Boot it up!**
   Execute your locally compiled executable directly in the root directory:
   ```bash
   go run .
   ```
   The API will natively deploy at `http://localhost:8080`.

## API Documentation
For highly detailed API endpoint diagrams, schemas, and payload examples, structurally read our localized technical implementation page here: [API Reference](docs/api.md)
