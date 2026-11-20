use aws_config::BehaviorVersion;
use aws_emulator_sample::round_trip;
use aws_sdk_s3::{config::Region, Client};

#[tokio::test]
async fn writes_and_reads_an_object() -> Result<(), Box<dyn std::error::Error + Send + Sync>> {
    let config = aws_config::defaults(BehaviorVersion::latest())
        .region(Region::new("us-east-1"))
        .load()
        .await;
    let s3 = Client::new(&config);
    let bucket = "emulator-rust-sample";
    s3.create_bucket().bucket(bucket).send().await?;

    let got = round_trip(&s3, bucket, "hello.txt", "Hello from Rust").await?;
    assert_eq!(got, "Hello from Rust");
    Ok(())
}
