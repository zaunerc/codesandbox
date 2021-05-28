FROM fedora:34

# Installing https also installs systemd.
RUN dnf -y install httpd procps less; systemctl enable httpd
RUN mkdir /appl/

#COPY cprinter /appl/
#COPY cprinter.service /etc/systemd/system/

#COPY goprinter/goprinter /appl/
#COPY goprinter.service /etc/systemd/system/

STOPSIGNAL SIGRTMIN+3
EXPOSE 80
CMD [ "/sbin/init" ]
