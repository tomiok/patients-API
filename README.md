# Backend - Patients

## Objective

Please demonstrate technical proficiency by creating a backend server application along with documentation and tests.

## Submission

To submit your work, send a tarball of your local repository.

## Requirements

The backend stack must use technologies we use in-house:

* Golang
* PostgreSQL
* docker-compose

Implement the following REST APIs:

| Method | URL                             | Description                       |
|--------|---------------------------------|-----------------------------------|
| GET    | /api/v1/patients                | Get all patients                  |
| GET    | /api/v1/patients/:id            | Get one patient                   |
| POST   | /api/v1/patients                | Add one patient                   |

* The APIs must be auth-protected.
* The request and response bodies must be in JSON format.
