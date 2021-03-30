FROM alpine
ENV WORKSPACE=/study_k8s/cmd/api
COPY api.out $WORKSPACE/api.out
WORKDIR $WORKSPACE
EXPOSE 8080
ENTRYPOINT ["./api.out"]
