import { GetObjectCommand, PutObjectCommand } from "@aws-sdk/client-s3";

export async function roundTrip(s3, bucket, key, value) {
  await s3.send(new PutObjectCommand({ Bucket: bucket, Key: key, Body: value }));
  const response = await s3.send(new GetObjectCommand({ Bucket: bucket, Key: key }));
  return response.Body.transformToString();
}
