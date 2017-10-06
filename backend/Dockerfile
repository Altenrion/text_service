FROM golang


#ADD . /go/src/kpi_cache_guard
ADD text_service_server /text_service_server

#RUN apt-get update
#RUN apt-get upgrade --yes
#ENV TZ=Europe/Moscow
#RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

#RUN go install kpi_cache_guard

#ENTRYPOINT /go/bin/kpi_cache_guard
ENTRYPOINT ["/text_service_server"]
#EXPOSE 8084