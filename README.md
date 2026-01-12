# OnTapAppRG


## Introduction
OnTapAppRG is a solution developed based on [OnTapApp](https://github.com/leokporto/OnTapApp) repository to manage and present information about beers and their styles, providing a modern and efficient experience for users and administrators. The project uses a modern architecture based on a modular monolith, integrating backend, frontend, and cloud/containerized infrastructure.

OnTapApp was originally made using .Net 9 and Blazor. OnTapAppRG uses Go on the backend and React on the frontend.


## Technologies Used
- **Go** - backend
- **React** - frontend
- **Docker** and **Docker Compose**
- **PgSql** as the database server


## Project Structure
```
OnTapAppRG
├── Backend/                   # Main Go API
├── Frontend/                  # React folder
├── Docker/                    # Docker configuration files
└── docker-compose.yml         # Container orchestration
```


## Future Features
- User registration and authentication using Google authenticator (OpenID)


---

## :warning: Warning

The project is not intended to be used as a production application.

The project will use docker compose, docker files and use all service defaults mannually. 
The objective of this project is to learn cloud native development and deployment, not to abstract it using tools such as aspire.


---
> Project under development. Suggestions and contributions are welcome!

