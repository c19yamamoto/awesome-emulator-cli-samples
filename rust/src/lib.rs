use aws_sdk_s3::{primitives::ByteStream, Client};
use std::error::Error;

pub async fn round_trip(
    s3: &Client,
    bucket: &str,
    key: &str,
    value: &str,
) -> Result<String, Box<dyn Error + Send + Sync>> {
    s3.put_object()
        .bucket(bucket)
        .key(key)
        .body(ByteStream::from(value.as_bytes().to_vec()))
        .send()
        .await?;

    let response = s3.get_object().bucket(bucket).key(key).send().await?;
    let bytes = response.body.collect().await?.into_bytes();
    Ok(String::from_utf8(bytes.to_vec())?)
}
