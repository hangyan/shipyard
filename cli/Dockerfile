FROM 192.168.200.14:5001/ubuntu:14.04
RUN apt-get update && apt-get install -y ca-certificates less
ADD cli /bin/shipyard
ENV PROG shipyard
ADD .docker/autocomplete /root/.autocomplete
ADD .docker/bashrc /root/.bashrc
RUN /bin/shipyard --generate-bash-completion > /etc/bash_completion.d/shipyard
CMD ["/bin/bash"]
