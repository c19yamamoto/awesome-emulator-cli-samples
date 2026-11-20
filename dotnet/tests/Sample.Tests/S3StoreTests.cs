using Amazon.S3;
using Amazon.S3.Model;
using Sample;
using Xunit;

namespace Sample.Tests;

public sealed class S3StoreTests
{
    [Fact]
    public async Task WritesAndReadsAnObject()
    {
        const string bucket = "emulator-dotnet-sample";
        using var s3 = new AmazonS3Client();
        await s3.PutBucketAsync(new PutBucketRequest { BucketName = bucket });

        var got = await S3Store.RoundTripAsync(s3, bucket, "hello.txt", "Hello from .NET");

        Assert.Equal("Hello from .NET", got);
    }
}
