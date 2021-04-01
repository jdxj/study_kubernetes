# docker build -t="jdxj/echo:0.1.0" .
FROM alpine
ENV WORKSPACE=/study_k8s/cmd/echo
COPY echo.out $WORKSPACE/echo.out
WORKDIR $WORKSPACE
ENTRYPOINT ["./echo.out"]
CMD ["--registry=consul", "--registry_address=10.103.112.102:8500"]
