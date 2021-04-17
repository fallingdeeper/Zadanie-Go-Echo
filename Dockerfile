FROM kprzystalski/projobj3:latest

USER root

RUN mkdir /home/kprzystalski/data

EXPOSE 1323

USER kprzystalski
