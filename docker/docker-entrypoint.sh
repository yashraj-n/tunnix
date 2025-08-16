#!/bin/bash

if [ -n "$SSH_API_KEY" ]; then
    echo "Changing password..."
    echo "tunnix:$SSH_API_KEY" | chpasswd
fi

echo "Starting SSH daemon on port 12000"
exec /usr/sbin/sshd -D -e
