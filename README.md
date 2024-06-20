# proxy-scribe
Tool for generating simple OpenAPI specs from URL paths in captured HTTP traffic

## How to use
1. Configure your client application or automated tests to send requests to `localhost:4000`
2. Start the proxy-scribe repl with the `proxy-scribe` command
3. Begin recording requests and responses with the `record` command
4. End the recording and output the spec with the `finish` command

## Limitations
- Currently supports only a small portion of the OpenAPI specification (essentials of that paths object)
- Only supports query and body parameters
