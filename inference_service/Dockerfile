FROM python:3.11-slim

WORKDIR /app
COPY . .

# Ensure the model is available at runtime
COPY models/loan_approval.onnx models/loan_approval.onnx

RUN pip install --no-cache-dir -r requirements.txt

ENV PYTHONUNBUFFERED=1

CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]
