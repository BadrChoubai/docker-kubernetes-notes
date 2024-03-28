# Managing Data and Volumes

In a docker container we can consider two types of data we may need to manage:

1. Temporary application data: [Stored with containers]
    - Fetched and produced in the running container
    - Stored in-memory or by temporary files
    -  Dynamic and changing, but cleared regularly
2. Permanent application data: [Stored with containers and volumes]
    - Fetched and produced in the running container
    - Stored in files or a database
    - Must not be lost if container stops or restarts

