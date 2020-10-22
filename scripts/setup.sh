#!/bin/sh

docker-compose -f ./scripts/docker-compose.yml up -d
echo 'Waiting for database to be up'
sleep 10
psql gokit_demo -h localhost -p 5434 -U postgres -c "create table videos
(
    id BIGSERIAL primary key,
    title text,
    duration int
);"
echo 'Setup Done!'
