using Amazon.S3;
using Amazon.S3.Model;

namespace Sample;

public static class S3Store
{
    public static async Task<string> RoundTripAsync(
        IAmazonS3 s3, string bucket, string key, string value, CancellationToken cancellationToken = default)
    {
        await s3.PutObjectAsync(new PutObjectRequest
        {
            BucketName = bucket,
            Key = key,
            ContentBody = value,
        }, cancellationToken);

        using var response = await s3.GetObjectAsync(bucket, key, cancellationToken);
        using var reader = new StreamReader(response.ResponseStream);
        return await reader.ReadToEndAsync(cancellationToken);
    }
}
