# Portfolio Semester 6 - Enterprise Software

This is my portoflio repository for the Individual project 

## About the project
This project is a web-based music player *(such as Spotify)*. It is a Progressive Web application **(PWA)** which means it could be downloaded on your mobile device through the browser and the application supports offline mode and some native features such as accessing the file system to store the songs for offline mode.

### Tech stack

1. Backend Project - Play-J Microservices

    1.1 **Golang** - microservices development

    1.2 **Goa** - Golang web framework

    1.3 **PostgreSQL** - database

    1.4 **Google Cloud Storage** - file system

2. DevOps - Infrastrucutre

    2.1 **Docker** - containarization platform

    2.2 **Kubernetes** - container manager

    2.3 **Google Cloud Platform** - cloud provider

## About the repository
All subprojects, related to this project, are part of this repository. Here is a brief explanation about the strucutre:

    -play-j
        - backend           // All backend code
            - cmd           // Microservices directory 
              - accountsd   // Account microservice
              - playerd     // Player microservice
              - paymentsd   // Payment microservice
            - internal      // Shared code (internal for all microservices)
        - compose           // Docker enviroment (compose files)
        - k8                // Kubernetes manifests
        - resources         // Documents, diagrams, etc.
