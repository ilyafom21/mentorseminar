import os
import pandas as pd
import boto3

from sqlalchemy import create_engine


POSTGRES_URL = "postgresql://admin:admin@postgres:5432/oil_db"

MINIO_ENDPOINT = "http://minio:9000"
MINIO_ACCESS_KEY = "admin"
MINIO_SECRET_KEY = "password123"
BUCKET_NAME = "oil-data"


engine = create_engine(POSTGRES_URL)

s3 = boto3.client(
    "s3",
    endpoint_url=MINIO_ENDPOINT,
    aws_access_key_id=MINIO_ACCESS_KEY,
    aws_secret_access_key=MINIO_SECRET_KEY
)


buckets = s3.list_buckets()
bucket_names = [bucket["Name"] for bucket in buckets["Buckets"]]

if BUCKET_NAME not in bucket_names:
    s3.create_bucket(Bucket=BUCKET_NAME)


tables = [
    "wells",
    "production",
    "well_telemetry",
    "well_targets",
    "pumps",
    "pump_sensors",
    "pump_failures",
    "deliveries",
    "drivers",
    "vehicles",
    "mart_daily_production",
    "mart_well_kpi",
    "mart_pressure_debit",
    "mart_ml_predictions",
    "mart_ml_metrics",
    "mart_anomalies",
    "mart_before_failure",
    "mart_risk_score",
    "mart_delay_weather",
    "mart_cost_distance",
    "mart_driver_kpi"
]


os.makedirs("data", exist_ok=True)


for table in tables:
    print(f"Loading table: {table}")

    df = pd.read_sql(f"SELECT * FROM {table}", engine)

    file_path = f"data/{table}.parquet"
    df.to_parquet(file_path, index=False)

    s3_key = f"raw/{table}/{table}.parquet"

    s3.upload_file(
        file_path,
        BUCKET_NAME,
        s3_key
    )

    print(f"Uploaded to MinIO: s3://{BUCKET_NAME}/{s3_key}")


print("ETL finished successfully")