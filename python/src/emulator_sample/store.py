from botocore.client import BaseClient


def round_trip(s3: BaseClient, bucket: str, key: str, value: str) -> str:
    """Write a value to S3 and return it by reading the object back."""
    s3.put_object(Bucket=bucket, Key=key, Body=value.encode("utf-8"))
    response = s3.get_object(Bucket=bucket, Key=key)
    return response["Body"].read().decode("utf-8")
