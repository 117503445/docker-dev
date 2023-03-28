FROM 117503445/dev-base
LABEL maintainer="117503445"
COPY ./public /root/public
WORKDIR /root
RUN /root/public/kde_init.sh
EXPOSE 5900 6080
CMD [ "/root/public/kde_start.sh" ]