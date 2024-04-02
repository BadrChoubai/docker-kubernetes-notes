# AWS EC2 Deployment

AWS EC2 is a service that allows you to spin up and manage your own remote machines

For our project there are a few steps we need to take:

1. Create and launch EC2 instance, VPC (Virtual Private Cloud) and Security Group
2. Configure security group to expose all required ports to WWW
3. Connect to instance (SSH), install Docker and run our container

## Bind Mounts, Volumes, and Copy

- In Development:
    - Containers should encapsulate the runtime environment but not necessarily
    the code.
    - Use "Bind Mounts" to provide our local host project files to the running
    container
    - Allows for instant updates without restarting the container

- In Production: [Image/Container is the single-source of truth]
    - A container should work standalone and should not have source code on the
    remote machine
    - Use `COPY` to copy code snapshot into our deployed image
    - Ensures that every image runs without any extra configuration or code
