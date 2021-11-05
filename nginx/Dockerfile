FROM nginx:alpine

RUN rm /etc/nginx/nginx.conf

COPY ./nginx/proxy.conf /etc/nginx/proxy.conf
COPY ./nginx/nginx.conf /etc/nginx/nginx.conf