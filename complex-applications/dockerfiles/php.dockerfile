FROM php:fpm-alpine

WORKDIR /var/www/html

RUN docker-php-ext-install pdo pdo_mysql

RUN addgroup -g 1000 laravel 
RUN adduser -G laravel -g laravel -s /bin/sh -D laravel

USER laravel
