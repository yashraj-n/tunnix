FROM alpine:latest

RUN apk add --no-cache openssh bash shadow && \
    ssh-keygen -A && \
    adduser -D tunnix && \
    sed -i 's/#\?PermitRootLogin.*/PermitRootLogin no/' /etc/ssh/sshd_config && \
    sed -i 's/#\?PasswordAuthentication.*/PasswordAuthentication yes/' /etc/ssh/sshd_config && \
    sed -i 's/#\?UsePAM.*/UsePAM no/' /etc/ssh/sshd_config && \
    sed -i 's/#\?GatewayPorts.*/GatewayPorts yes/' /etc/ssh/sshd_config && \
    echo "Port 12000" >> /etc/ssh/sshd_config

RUN echo '#!/bin/bash' > /start.sh && \
    echo 'if [ -n "$SSH_API_KEY" ]; then' >> /start.sh && \
    echo '    echo "tunnix:$SSH_API_KEY" | chpasswd' >> /start.sh && \
    echo 'fi' >> /start.sh && \
    echo 'exec /usr/sbin/sshd -D -e' >> /start.sh && \
    chmod +x /start.sh

EXPOSE 12000
EXPOSE 12001

CMD ["/start.sh"]
