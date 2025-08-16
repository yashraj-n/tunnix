FROM alpine:latest

RUN apk add --no-cache openssh bash shadow && \
    ssh-keygen -A && \
    adduser -D tunnix

COPY sshd_config /etc/ssh/sshd_config

COPY docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh

EXPOSE 12000
EXPOSE 12001

HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
  CMD pgrep sshd >/dev/null || exit 1


CMD ["/docker-entrypoint.sh"]
