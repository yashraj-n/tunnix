FROM alpine:latest

RUN apk add --no-cache openssh bash shadow && \
    ssh-keygen -A && \
    adduser -D tunnix

COPY sshd_config /etc/ssh/sshd_config

COPY docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh

EXPOSE 12000
EXPOSE 12001

CMD ["/docker-entrypoint.sh"]
