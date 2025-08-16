#!/bin/bash

if [ -n "$SSH_API_KEY" ]; then
    echo "Changing password..."
    echo "tunnix:$SSH_API_KEY" | chpasswd
fi

# MOTD
echo "" > /etc/motd
echo "--------------------------------" >> /etc/motd
echo "Welcome to Tunnix" >> /etc/motd
echo "Please start this repo if you find this project useful https://github.com/yashraj-n/tunnix" >> /etc/motd
echo "--------------------------------" >> /etc/motd
echo "Use this URL: $TUNNIX_FQDN" >> /etc/motd
echo "" >> /etc/motd

echo "Starting SSH daemon on port 12000"
exec /usr/sbin/sshd -D -e
