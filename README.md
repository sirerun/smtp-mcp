# smtp-mcp

Give MCP-compatible agents a simple way to preview and send email through your
SMTP server. `smtp-mcp` runs locally over stdio and works with an existing SMTP
relay; it does not require [smtpd](https://github.com/sirerun/smtpd).

Sending is **off by default**. You configure the sender, SMTP endpoint and
credentials when starting the server. Agents supply the recipient and message.

## Install

Download an archive for your platform from
[Releases](https://github.com/sirerun/smtp-mcp/releases), verify it against the
published checksums, and extract the `smtp-mcp` binary. Binaries are provided for
macOS, Linux and Windows on ARM64 and AMD64.

Or install from source with Go 1.25 or later:

```sh
go install github.com/sirerun/smtp-mcp/cmd/smtp-mcp@v0.1.0
smtp-mcp --version
```

## Connect your agent

Add this entry to a client that supports stdio MCP servers. Replace the executable
path and example sender with your own values:

```json
{
  "mcpServers": {
    "smtp": {
      "command": "/absolute/path/to/smtp-mcp",
      "args": ["--from", "you@example.com"]
    }
  }
}
```

This starts in preview-only mode. No SMTP connection is made by preview or status.
Client configuration formats and environment handling vary; use your client's
secure mechanism to supply credentials to the subprocess.

### Enable sending

Configure an SMTP submission endpoint, then explicitly add `--enable-send`:

```json
{
  "mcpServers": {
    "smtp": {
      "command": "/absolute/path/to/smtp-mcp",
      "args": [
        "--smtp-host", "smtp.example.com:587",
        "--from", "you@example.com",
        "--smtp-user", "you@example.com",
        "--enable-send"
      ]
    }
  }
}
```

Supply the password in the subprocess environment as `SMTP_MCP_PASSWORD`. Do not
put it in tool arguments, prompts or source control. The server supports SMTP
PLAIN authentication over verified TLS; OAuth authentication is not implemented.
Your provider may require a dedicated SMTP credential or app password.

STARTTLS is the default. For an implicit-TLS endpoint (commonly port 465), add
`--smtp-tls-mode implicit`. Use your provider's documented hostname, port and TLS
mode; the port number alone does not select the mode.

## Tools

| Tool | What it does |
| --- | --- |
| `smtp_preview_email` | Validates and returns message fields without sending. |
| `smtp_send_email` | Submits one plain-text email to one recipient when enabled. |
| `smtp_status` | Reports configuration flags without exposing credentials. |

Preview and send accept the same arguments:

```json
{
  "to": "person@example.net",
  "subject": "Workshop details",
  "text": "Hello, here are the workshop details you requested.",
  "reply_to": "you@example.com"
}
```

`reply_to` is optional. Addresses must be bare ASCII email addresses, such as
`person@example.net`. UTF-8 subjects and body text are supported; SMTPUTF8
addresses, HTML, attachments, CC and BCC are not. The body is limited to 64 KiB,
with additional header and total-payload limits.

A preview does not authorize or reserve a send. Access to a send-enabled process
allows an agent to send as the configured identity to any valid recipient, so
restrict which clients can launch it and apply your approval policy in the client.

## SMTP security and delivery behavior

- TLS certificates and hostnames are verified. Custom trust roots can be supplied
  with `--tls-ca-file`, and `--tls-server-name` overrides the verification name.
- `--allow-plaintext-loopback` is only for explicit local test sinks. It requires
  a literal loopback IP and does not permit credentials over plaintext.
- Sender identity and SMTP endpoint are fixed at startup. Tool inputs cannot
  change them, and header injection is rejected.
- `--timeout` bounds each submission (default `30s`). Cancellation closes the
  connection. The server does not retry automatically.
- `accepted_by_smtp` means the SMTP endpoint accepted the message. It does **not**
  prove delivery or inbox placement. If acknowledgement is lost, the outcome can
  be unknown; inspect your SMTP server's logs before retrying to avoid duplicates.

Your SMTP service owns queue durability, delivery, SPF/DKIM/DMARC configuration,
bounces and reputation. This project does not sign messages, run a mail server,
warm IP addresses, manage campaigns, track opt-outs, rate-limit across agents or
provide deduplication. A separate process is launched for each MCP client.

Use `smtp-mcp --help` for every flag. Protocol messages use stdout and diagnostics
use stderr. The initial release supports stdio only. Tool registration is separate
from the transport so a future authenticated HTTP server can reuse the handlers;
there is no HTTP listener in this release.

## Development

```sh
go test ./...
go test -race ./...
go vet ./...
go build -o smtp-mcp ./cmd/smtp-mcp
```

Tests use in-memory MCP transports and local fake SMTP servers, including verified
STARTTLS and implicit TLS. They do not send external email. See
[release instructions](docs/release.md) for linting, vulnerability scanning and
cross-platform packaging.

## License

[Apache License 2.0](LICENSE). Extracted from the SMTP MCP adapter developed in
[sirerun/smtpd](https://github.com/sirerun/smtpd/pull/7).
