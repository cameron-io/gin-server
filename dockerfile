FROM python:alpine
WORKDIR /var/lib/gin
ADD requirements.txt .
RUN apk update && \
    apk upgrade && \
    apk add --no-cache --virtual .build-deps \
        gcc \
        python3-dev \
        musl-dev \
        postgresql-dev \
    && pip install --no-cache-dir psycopg2 \
    && apk del --no-cache .build-deps
RUN apk --no-cache add libpq
RUN pip3 install -r requirements.txt
ADD . .
RUN source .env
CMD python3 -m flask run
