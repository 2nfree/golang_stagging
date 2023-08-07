FROM golang:1.20-alpine as Builder
ARG EXECUTABLE_APP
COPY ./ /app/
WORKDIR /app/
RUN CGO_ENABLED=0 GOOS=linux go build -a -o ${EXECUTABLE_APP} ./main.go

FROM alpine:3
ARG EXECUTABLE_APP
ENV OPTIONS="" \
    EXECUTABLE_APP=${EXECUTABLE_APP}
WORKDIR /opt/app
COPY --from=0 /app/$EXECUTABLE_APP /opt/app/
RUN chmod +x /opt/app/${EXECUTABLE_APP}
EXPOSE 7000
ENTRYPOINT exec /opt/app/${EXECUTABLE_APP} ${OPTIONS}
