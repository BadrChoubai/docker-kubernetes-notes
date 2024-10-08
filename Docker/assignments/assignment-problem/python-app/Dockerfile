# Build your image up from another base image
FROM python

# Subsequent commands will be run relative to WORKDIR
WORKDIR /app

# Which files from our machine should go into our image/container
COPY . /app

# Executes with the starting of our container
CMD ["python", "bmi.py"]

