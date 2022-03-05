FROM fluent/fluentd:v1.14

USER root

RUN gem install fluent-plugin-elasticsearch && \
    gem install fluent-plugin-concat && \
    gem sources --clear-all

USER fluent
