# REPARTNTERS Interview

## PackSize Calculator API

### Run locally
    1. `docker compose -f docker-compose.local.yml up`
    2. Open `http://localhost:8080` in your browser

### Live Demo
1. Open `http://repartners-interview-env-v3.eba-mmpw86mp.eu-central-1.elasticbeanstalk.com` in your favorite browser.

## Api Routes
- GET   /api/pack-size
- POST  /api/pack-size
- DEL   /api/pack-size/{id}
- GET   /api/calculate?order={n} // order query param is required
