FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY announcerd /app/announcerd

EXPOSE 6000 

USER nonroot:nonroot

ENTRYPOINT ["/announcerd"]
