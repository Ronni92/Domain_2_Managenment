FROM python:3.10

WORKDIR /app

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

# Documenta el puerto interno que usará el contenedor
EXPOSE 4000

# Ejecuta el servidor Uvicorn con recarga automática (solo para desarrollo)
CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "4000", "--reload"]
