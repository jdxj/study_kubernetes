# $ sudo docker build -t="jdxj/study_kubernetes:v0.2.0" .
# $ sudo docker push jdxj/study_kubernetes:v0.2.0
FROM alpine
COPY main.out /app/main.out
WORKDIR /app
EXPOSE 8080
ENTRYPOINT ["/app/main.out"]