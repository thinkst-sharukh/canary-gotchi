# Go-React-Postgres Boilerplate

Welcome to the Go-React-Postgres boilerplate! This repository provides a ready-to-use template for building a fullstack application with Go for the backend, React for the frontend, and Postgres as the database. The entire setup is Dockerized for easy development and deployment.

## To get started please visit [here](https://github.com/sharukhkhanajm/fullstack-boilerplates?tab=readme-ov-file#getting-started)

## View pgAdmin GUI

- Visiting this URL http://localhost:5050 should open a pgAdmin login screen.
- Enter `admin@admin.com` in email
- Enter `123` in password

You can change the credentials in `docker-compose.dev.yml`

docker compose up -d --build
docker ps
sudo docker logs --since=1h d73a57b07172
docker ps
