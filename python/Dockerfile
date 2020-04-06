FROM python:3.7 as pybuilder
WORKDIR /app
RUN python -m venv /app/venv
RUN /app/venv/bin/pip install --upgrade pip
COPY requirements.txt /app/requirements.txt
RUN cat /app/requirements.txt | /app/venv/bin/pip install -r /dev/stdin

FROM python:3.7
WORKDIR /app
COPY --from=pybuilder /app/venv/ /app/venv/
COPY app.py /app/app.py
ENV PORT=5000
CMD exec /app/venv/bin/gunicorn \
	--bind=0.0.0.0:$PORT \
	--workers=3 \
	--worker-tmp-dir=/dev/shm \
	app:app
