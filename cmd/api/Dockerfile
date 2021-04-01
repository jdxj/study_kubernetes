# docker build -t="jdxj/api:0.5.0" .
FROM alpine
ENV WORKSPACE=/study_k8s/cmd/api
COPY api.out $WORKSPACE/api.out
WORKDIR $WORKSPACE
EXPOSE 8080
ENTRYPOINT ["./api.out"]
CMD ["--registry=consul", "--registry_address=10.103.112.102:8500"]
