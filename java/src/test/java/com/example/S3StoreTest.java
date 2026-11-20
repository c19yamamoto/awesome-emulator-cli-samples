package com.example;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.UUID;
import org.junit.jupiter.api.Test;
import software.amazon.awssdk.services.s3.S3Client;
import software.amazon.awssdk.services.s3.model.CreateBucketRequest;

class S3StoreTest {
    @Test
    void writesAndReadsAnObject() throws Exception {
        String bucket = "emulator-java-" + UUID.randomUUID().toString().replace("-", "");
        try (S3Client s3 = S3Client.builder().build()) {
            s3.createBucket(CreateBucketRequest.builder().bucket(bucket).build());
            assertEquals("Hello from Java", S3Store.roundTrip(s3, bucket, "hello.txt", "Hello from Java"));
        }
    }
}
