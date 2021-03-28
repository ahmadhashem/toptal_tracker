FROM golang
COPY entrypoint.sh /
ENTRYPOINT ["/entrypoint.sh"]
