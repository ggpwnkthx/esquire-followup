#!/bin/sh
if [ -z "$DOMAIN_NAME" ]; then
  echo "No root domain set."
  exit 1
fi
if [ ! -f /config/configuration.yml ]; then
    cat <<EOF > /config/configuration.yml
###############################################################
#                   Authelia configuration                    #
###############################################################

host: 0.0.0.0
port: 9091
log_level: debug
jwt_secret: $(base64 /dev/urandom | tr -d "/+" | head -c32)
default_redirection_url: https://login.$DOMAIN_NAME

totp:
  issuer: authelia.com

authentication_backend:
  file:
    path: /config/users.yml
    password:
      algorithm: argon2id
      iterations: 1
      salt_length: 16
      parallelism: 8
      memory: 128

access_control:
  default_policy: deny
  rules:
    - domain: "*.$DOMAIN_NAME"
      policy: one_factor

session:
  name: authelia_session
  secret: $(base64 /dev/urandom | tr -d "/+" | head -c32)
  expiration: 3600 # 1 hour
  inactivity: 300 # 5 minutes
  domain: $DOMAIN_NAME

regulation:
  max_retries: 3
  find_time: 120
  ban_time: 300

storage:
  local:
    path: /config/preferences.sqlite3

EOF
    if [ -f /config/smtp.yml ]; then
        cat /config/smtp.yml >> /config/configuration.yml
    else
        cat <<EOF >> /config/configuration.yml
notifier:
  filesystem:
    filename: /config/notifications
EOF
    fi
fi

if [ ! -f /config/users.yml ]; then
    ADMIN_PASSWORD=${ADMIN_PASSWORD:-CheckThisOut}
    ADMIN_PASSWORD=$(authelia hash-password $ADMIN_PASSWORD -i 1 -m 128 -p 8 -k 32 -l 16  | awk '{print $3}')
    cat <<EOF > /config/users.yml
users:
  admin:
    password: $ADMIN_PASSWORD
    displayname: Administrator
    email: "$ADMIN_EMAIL"
    groups:
    - admin
EOF
fi

if [ ! -f /config/notifications ]; then
    touch /config/notifications
fi

/usr/local/bin/entrypoint.sh --config /config/configuration.yml $@