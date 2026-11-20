import uuid

import boto3

from emulator_sample.store import round_trip


def test_s3_round_trip() -> None:
    s3 = boto3.client("s3")
    bucket = f"emulator-python-{uuid.uuid4().hex}"
    s3.create_bucket(Bucket=bucket)

    assert round_trip(s3, bucket, "hello.txt", "Hello from Python") == "Hello from Python"
