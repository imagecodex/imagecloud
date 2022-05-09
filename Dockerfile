FROM songjiayang/govips:v0.0.1
WORKDIR /build/
COPY . ./
RUN CGO_ENABLED=1 GOOS=linux go build -mod=vendor -o imagecloud .

FROM songjiayang/govips:v0.0.1
COPY --from=0 /build/imagecloud /bin/imagecloud
COPY --from=0 /build/configs/config.yml /etc/imagecloud/imagecloud.yml

WORKDIR /imagecloud
RUN chown -R nobody:nobody /imagecloud

USER nobody
EXPOSE 8080
VOLUME [ "/imagecloud" ]
ENTRYPOINT [ "/bin/imagecloud" ]
CMD [ "--config.file=/etc/imagecloud/imagecloud.yml" ]