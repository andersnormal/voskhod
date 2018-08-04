FROM scratch
ADD bin/voskhod_*_linux_amd64 /voskhod
CMD ["/voskhod"]
