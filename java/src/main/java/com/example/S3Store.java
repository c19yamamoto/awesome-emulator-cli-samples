package com.example;

import java.io.IOException;
import java.nio.charset.StandardCharsets;
import software.amazon.awssdk.core.sync.RequestBody;
import software.amazon.awssdk.services.s3.S3Client;
import software.amazon.awssdk.services.s3.model.GetObjectRequest;
import software.amazon.awssdk.services.s3.model.PutObjectRequest;

public final class S3Store {
    private S3Store() {}

    public static String roundTrip(S3Client s3, String bucket, String key, String value) throws IOException {
        s3.putObject(
                PutObjectRequest.builder().bucket(bucket).key(key).build(),
                RequestBody.fromString(value, StandardCharsets.UTF_8));
        return s3.getObjectAsBytes(GetObjectRequest.builder().bucket(bucket).key(key).build())
                .asUtf8String();
    }
}
