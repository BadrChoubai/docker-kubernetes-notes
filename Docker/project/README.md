# Deploying to Production

This directory contains project code for the lessons on deploying Docker into
production on AWS (Amazon Web Services)

## Development to Production

1. Bind Mounts shouldn't be used in production
2. Containerized apps might need to include a build step
3. Multi-Container projects might need to be split across multiple hosts / remote
   machines (regionality)
4. Trade-offs between control and responsibility

## Deployment Process and Solutions

This project covers a few deployment approaches:

1. Install Docker on a remote host [aws-ec2](./aws-ec2/README.md)
2. Using a managed service: [aws-ecs-tf](./aws-ecs-tf/README.md)

