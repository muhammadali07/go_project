FROM postgres:latest

ENV POSTGRES_USER gorestuser
ENV POSTGRES_PASSWORD gorestpass
ENV POSTGRES_DB gorestdevelopment

COPY db/init.sql /docker-entrypoint-initdb.d/
