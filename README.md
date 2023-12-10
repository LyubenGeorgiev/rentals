# Rentals Project

This project is a Go backend json application for rentals information retrieval.

## Getting Started

To run this application locally, Docker and Docker Compose must be installed on your machine.

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Starting the Application

1. Clone this repository to your local machine:

    ```bash
    git clone https://github.com/LyubenGeorgiev/rentals.git
    ```

2. Navigate to the project directory:

    ```bash
    cd rentals
    ```

3. Start the application using Docker Compose:

    Optionally add ```--build``` flag if you have changed the code.
    ```bash
    docker compose up --build app
    ```

4. Once the containers are running, the Go application will be available at `http://localhost:8080`.

5. Access the application in your web browser or use API client tools to interact with the application endpoints.

## Run the tests

Optionally add ```--build``` flag if you have changed the code.
```bash
docker compose up tests
```

## Endpoints

### `GET /rentals/<RENTAL_ID>`
- Read one rental endpoint

### `GET /rentals`
- Read many (list) rentals endpoint
- Supported query parameters:
  - `price_min` (number)
  - `price_max` (number)
  - `limit` (number)
  - `offset` (number)
  - `ids` (comma separated list of rental ids)
  - `near` (comma separated pair [lat,lng])
  - `sort` (string)

#### Examples:

- `rentals?price_min=9000&price_max=75000`
- `rentals?limit=3&offset=6`
- `rentals?ids=3,4,5`
- `rentals?near=33.64,-117.93` (within 100 miles)
- `rentals?sort=price`
- `rentals?near=33.64,-117.93&price_min=9000&price_max=75000&limit=3&offset=6&sort=price`

## Stopping the Application

To stop the application and shut down the containers, use:

```bash
docker compose down