FROM golang:1.9
RUN mkdir -p /go/src/github.com/kidsdynamic/childrenlab_admin
ADD ./build /go/src/github.com/kidsdynamic/childrenlab_admin/
WORKDIR /go/src/github.com/kidsdynamic/childrenlab_admin
CMD ["/go/src/github.com/kidsdynamic/childrenlab_admin/main"]

EXPOSE 8114