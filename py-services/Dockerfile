FROM python:3.9

WORKDIR /app

# Copy requirements first for better caching
COPY requirements.txt .
RUN pip install -r requirements.txt

# Then copy the rest of the application
COPY . .

CMD ["python", "main.py"]