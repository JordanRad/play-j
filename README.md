# Portfolio Semester 6 - Enterprise Software
## Jordan Radushev


This is my portoflio repository for the Individual project 

### About the project
This project is a web-based music player *(such as Spotify)*. It is a Progressive Web application **(PWA)** which means it could be downloaded on your mobile device through the browser and the application supports offline mode and some native features such as accessing the file system to store the songs for offline mode.

#### Tech stack
There are three main parts of the project, regarding the tech stack of this project.

1. Frontend Project - Play-J Client
    
    1.1  **Vue JS** - frontend framework

    1.2 **PrimeVue** & **Primeflex** - component library & CSS utility library 

2. Backend Project - Play-J Microservices

    2.1 **Golang** - microservices development

    2.2 **Goa** - Golang web framework

    2.3 **PostgreSQL** - database

    2.4 **Google Cloud Storage** - file system

3. DevOps - Infrastrucutre

    3.1 **Docker** - containarization platform

    3.2 **Kubernetes** - container manager

    3.3 **Google Cloud Platform** - cloud provider




### About the repository
All subprojects, related to this project, are part of this repository. Here is a brief explanation about the strucutre:

    -play-j
        - backend           // All backend code
            - cmd           // Microservices directory 
              - accountsd   // Account microservice
              - playerd     // Player microservice
              - paymentsd   // Payment microservice
            - internal      // Shared microservice code 
        - compose           // Docker enviroment (compose files)
        - frontend          // All frontend code
        - resources         // Documents, diagrams, etc.