docker run -p 8080:80 -d --name feedback-app --rm -v "$(pwd):/app:ro" -v /app/temp -v /app/node_modules -v feedback:/app/feedback docker-tutorial-volumes
