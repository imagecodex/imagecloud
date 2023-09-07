FROM songjiayang/govips:v0.3.0
WORKDIR /build/
COPY . ./
RUN CGO_ENABLED=1 GOOS=linux go build -mod=vendor -o imagecloud .

FROM songjiayang/govips:v0.3.0
COPY --from=0 /build/imagecloud /bin/imagecloud
COPY ./configs/config.yml /etc/imagecloud/imagecloud.yml
COPY ./assets/fonts/Font-OPPOSans /usr/share/fonts/Font-OPPOSans

WORKDIR /imagecloud
RUN chown -R nobody:nobody /imagecloud

USER nobody
EXPOSE 8080
VOLUME [ "/imagecloud" ]
ENTRYPOINT [ "/bin/imagecloud" ]
CMD [ "--config.file=/etc/imagecloud/imagecloud.yml" ]