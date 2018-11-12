FROM golang:1.10.2

RUN mv /etc/apt/sources.list /etc/apt/sources.list.bak && \
    echo "deb http://mirrors.163.com/debian/ stretch main non-free contrib" >/etc/apt/sources.list && \
    echo "deb http://mirrors.163.com/debian/ stretch-proposed-updates main non-free contrib" >>/etc/apt/sources.list && \
    echo "deb-src http://mirrors.163.com/debian/ stretch main non-free contrib" >>/etc/apt/sources.list && \
    echo "deb-src http://mirrors.163.com/debian/ stretch-proposed-updates main non-free contrib" >>/etc/apt/sources.list

RUN apt-get update

RUN apt-get install -y build-essential git autoconf automake libtool unzip

RUN git clone -b v3.6.1 --depth=1  https://github.com/google/protobuf &&\
    cd protobuf &&\
    ./autogen.sh &&\
    ./configure --prefix=/usr &&\
    make -j4 &&\
    make install &&\
    cd ../ &&\
    rm -rf protobuf

RUN go get github.com/mwitkow/go-proto-validators

RUN go get github.com/golang/protobuf/protoc-gen-go

RUN go get github.com/mwitkow/go-proto-validators/protoc-gen-govalidators

RUN mkdir -p src/golang.org/x &&\
    cd src/golang.org/x &&\
    git clone --depth=1 https://github.com/golang/sys &&\
    git clone --depth=1 https://github.com/golang/net && \
    git clone --depth=1 https://github.com/golang/text &&\
    cd ../../ &&\
    mkdir -p google.golang.org/ &&\
    cd google.golang.org/ &&\
    git clone --depth=1 https://github.com/google/go-genproto &&\
    mv go-genproto genproto &&\
    cd ../ &&\
    git clone -b v1.16.0 --depth=1 https://github.com/grpc/grpc-go &&\
    mv grpc-go google.golang.org/grpc &&\
    go install google.golang.org/grpc &&\
    cd ../../../

CMD mkdir -p src/github.com/nayotta/metathings &&\
    cd src/github.com/nayotta/metathings &&\
    make protos
