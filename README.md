# Overview

## Running
From the root of the monorepo, run `docker-compose up --build`. This starts the frontend, db and the backend server.

From there, enter any URL from Garage into the input field and check out the PDF that auto-downloads.

LMK If any questions, and I go into more details on the implementation below!

## Tech Stack Breakdown
### Frontend
- React TS
- Redux RTK code auto-gen
  - auto-generates the API layer for our frontend based on an openapi spec file that is created from our protobuf definitions
- CSS written out of the box for each component / page
- basic styling, due to time constraints i decided to spend less time here and more time refining the overall architecture

### Backend
- Golang
- Protobuf (for api modeling and generating an openapi file for our frontend code gen)
- godoc for generating custom HTML templates with dynamic data, which are then converted to PDFs with weasyprint
- custom CSS for PDF file, which includes a title, image, description, and price

I tried to demonstrate much of my opinions around backend structure in this project. Specifically, I separate our core logic into "services". Each service has a corresponding `models` file, and a corresponding `errors` file for custom errors.

My opinion is mainly that services should be a DAG. Dependencies can only flow one way. Additionally, I treat errors as first-class citizens, mostly because errors are so key to "good" Go programming (in my opinion).

All generated code lives under `gen`. All db tables and db connection logic live under `db`. Any connections to third parties lives under `providers`.

My code follows a few rules, the main one being:
- api calls service
- service calls repo
- repo calls db / ORM

service should not need to know about db / ORM. api should not need to know about db or repo models.

### Protobuf
As mentioned above, I use protobuf to define the API models for the backend. Then, I generate the openapi spec for these models running `make all` from the root of the monorepo. From there, I then run `pnpm run validate && pnpm run codegen` in the `app` frontend directory to generate the frontend models.

This makes it super easy to add new api clients on the frontend, and to define API models on the backend. Most of the boilerplate is auto-generated, which is a great developer experience (-:

## Improvements
Things I would do If I had more time:
1. More robust frontend, including basic auth and multiple pages. I would also probably add some more features to the invoice generator, although I think that a "dumb" UI for this project is fine (simple UX, doesn't over-complicate things)
2. Unit testing (I take this very seriously in a production environment but decided the tradeoff on no unit tests here was worth it)