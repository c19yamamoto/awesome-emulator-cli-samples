# awesome-emulator-cli samples

Small integration tests showing how to use AWS SDKs against AWS Emulator without an AWS account or deployment. The Python, Go, Java, Rust, and .NET applications and tests run in their native runtimes; Node.js is only needed for [`awesome-emulator-cli`](https://www.npmjs.com/package/awesome-emulator-cli). The JavaScript example also uses Node.js for its SDK test.

AWS Emulator exposes the same local API endpoint to every language, rather than providing a separate runtime for each language. These examples cover the six SDK clients for which the AWS Emulator guide gives explicit certificate-trust instructions: Python (Boto3), Go v2, JavaScript v3, Java 2.x, Rust, and .NET. Other SDKs can also target AWS-compatible endpoints, but their certificate configuration is not covered by this sample set.

Every test stores a short string in an ephemeral S3 bucket and reads it back through the language's AWS SDK. The bucket and object live only in the emulator's in-memory state.

## Run an example

Requirements: Node.js 22 or newer, Docker, and the runtime for the example you choose. The CLI downloads and verifies the upstream AWS Emulator Linux package, then runs it in Docker. Install the CLI once:

```sh
npm ci
```

Install the selected example's dependencies, then run its test through the CLI:

| SDK | Install dependencies | Run the integration test |
| --- | --- | --- |
| Python / Boto3 | `python -m pip install -e 'python[dev]'` | `npx awesome-emulator run --telemetry off -- python -m pytest -q python/tests` |
| Go v2 | *(resolved by Go Modules)* | `npx awesome-emulator run --telemetry off -- go -C go test ./...` |
| JavaScript v3 | `npm ci --prefix javascript` | `npx awesome-emulator run --telemetry off -- node --test javascript/s3-store.test.js` |
| Java 2.x | *(Maven resolves dependencies)* | `npx awesome-emulator run --telemetry off -- mvn -B -f java/pom.xml test` |
| Rust | *(resolved by Cargo)* | `npx awesome-emulator run --telemetry off -- cargo test --manifest-path rust/Cargo.toml` |
| .NET | *(NuGet resolves dependencies)* | `npx awesome-emulator run --telemetry off -- dotnet test dotnet/tests/Sample.Tests.csproj` |

The CLI starts and stops the emulator around the child test process and provides the endpoint, credentials, region, and certificate settings. For Java and Rust, this includes the JVM trust store and `SSL_CERT_FILE` settings needed for the emulator's HTTPS certificate.

## GitHub Actions

Each SDK has its own workflow under [`.github/workflows/`](.github/workflows/). A push or pull request that changes a language's folder runs only that language's workflow; changes to the shared CLI dependency run all of them. Each workflow can also be started manually with `workflow_dispatch`. No AWS credentials or account are used.

## References

- [AWS Emulator: Pointing your AWS SDK at AWS Emulator](https://docs.aws.amazon.com/emulator/latest/userguide/point-sdk-to-emulator.html)
- [AWS Emulator: Getting started with embedded mode for Rust](https://docs.aws.amazon.com/emulator/latest/userguide/getting-started-embedded.html)
- [awesome-emulator-cli on npm](https://www.npmjs.com/package/awesome-emulator-cli)
