FROM debian

RUN apt-get update && apt-get install -y curl && apt-get clean

ENV HANDS_ON_SERVER docker.devel.iit.cnr.it
ENV HANDS_ON_PORT 18273

COPY book /

CMD /book
