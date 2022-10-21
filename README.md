# Portfolio Semester 6 - Enterprise Software

This is my portoflio repository for the Individual project 

## About the project
This project is a web-based music player API *(such as Spotify)*. It is built upon the concept of microservices, makes use of Kubernetes. The main requirements for the project were:
- Good system design, resulting in scalable (cloud) architecture.
- Security which covers at least the OWASP top 10 security vulnerabilities (when applicable)
- Automated CI/CD pipeline
- Integration with other cloud services 

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
        - backend           // Main folder
            - cmd           // Microservices directory (starting point for each microservice)
              - accountsd   // Account microservice
              - playerd     // Player microservice
              - paymentsd   // Payment microservice
            - internal      // Shared code (internal for all microservices)
        - compose           // Docker enviroment (compose files)
        - k8                // Kubernetes manifests
        - resources         // Documents, diagrams, etc.
