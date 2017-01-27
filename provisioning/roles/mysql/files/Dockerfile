FROM ubuntu:16.04

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update
RUN apt-get -y install apt-utils

RUN echo percona-server-server-5.6 percona-server-server/root_password password '' | debconf-set-selections
RUN echo percona-server-server-5.6 percona-server-server/root_password_again password '' | debconf-set-selections
RUN apt-get -y install percona-server-server-5.6 curl

RUN sed -Ei 's/^bind-address/#&/' /etc/mysql/percona-server.conf.d/mysqld.cnf

RUN service mysql start && \
    mysql -e 'GRANT ALL PRIVILEGES ON *.* TO "root"@"%" WITH GRANT OPTION;' && \
    mysql -e 'SET PASSWORD FOR "root"@"%" = PASSWORD("test");'

COPY start_mysql.sh /root

EXPOSE 3306

ENTRYPOINT ["/bin/bash", "-c", "/root/start_mysql.sh"]
