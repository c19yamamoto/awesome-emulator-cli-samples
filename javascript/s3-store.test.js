import assert from "node:assert/strict";
import { randomUUID } from "node:crypto";
import test from "node:test";
import { CreateBucketCommand, S3Client } from "@aws-sdk/client-s3";
import { roundTrip } from "./s3-store.js";

test("writes and reads an S3 object", async () => {
  const s3 = new S3Client({ disableHostPrefix: true });
  const bucket = `emulator-js-${randomUUID().replaceAll("-", "")}`;
  try {
    await s3.send(new CreateBucketCommand({ Bucket: bucket }));
    assert.equal(await roundTrip(s3, bucket, "hello.txt", "Hello from JavaScript"), "Hello from JavaScript");
  } finally {
    s3.destroy();
  }
});
