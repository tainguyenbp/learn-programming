# Stage 1
FROM node:14 AS builder
RUN apk add --no-cache yarn
ARG ENV
ARG DOMAIN
ARG HOST

# Stage 2
FROM alpine:3.12
COPY --from=builder /app /app
ARG ENV
ARG DOMAIN
ARG HOST
