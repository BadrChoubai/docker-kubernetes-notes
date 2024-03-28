docker run -p 8081:80 -d --name feedback-app-prod --rm -v "$(pwd):/app:ro" -v /app/temp -v /app/node_modules -v feedback:/app/feedback docker-tutorial-volumes:prod
docker run -p 8080:8080 -d --name feedback-app-dev --rm -v "$(pwd):/app:ro" -v /app/temp -v /app/node_modules -v feedback:/app/feedback docker-tutorial-volumes:dev
