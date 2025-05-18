# REPARTNTERS Interview

## PackSize Calculator API

### Run locally
    1. `docker compose -f docker-compose.local.yml build && docker compose -f docker-compose.local.yml up`
    2. Open `http://localhost` in your browser
        - if port '80' is not available on your local machine, replace the `docker-compose.yml` port segment from "80:80" to "8080:80" and use localhost:8080

### Live Demo
1. Open `http://repartners-env-v4.eba-mmpw86mp.eu-central-1.elasticbeanstalk.com` in your favorite browser.   
Note: I've deployed this service using AWS elasticbeanstalk.

## Api Routes
- GET   /api/pack-size  // To get all available pack sizes
- POST  /api/pack-size  // To create pack-size using json body {"value": {packSize: int}}
- DEL   /api/pack-size/{id} // To delete pack size by id
- GET   /api/calculate?order={n} // To calculate pack size [order query param is required]

## Tests
There are table driven tests written for PackSize.Calulate method.   
To run the tests: `go test ./...`