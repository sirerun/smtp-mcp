// Command smtp-mcp runs the SMTP MCP server over stdio.
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/sirerun/smtp-mcp/internal/mcpserver"
)

var version = "dev"

func main() {
	fs := flag.NewFlagSet("smtp-mcp", flag.ExitOnError)
	showVersion := fs.Bool("version", false, "print version and exit")
	cfg := mcpserver.Config{Version: version}
	fs.BoolVar(&cfg.EnableSend, "enable-send", false, "enable smtp_send_email (required to send)")
	fs.StringVar(&cfg.Host, "smtp-host", "", "owner-configured SMTP submission relay host:port")
	fs.StringVar(&cfg.From, "from", "", "fixed owner-configured sender email address")
	fs.StringVar(&cfg.Username, "smtp-user", "", "SMTP authentication username")
	fs.BoolVar(&cfg.AllowPlaintext, "allow-plaintext-loopback", false, "allow unencrypted SMTP only to an explicit loopback IP")
	fs.DurationVar(&cfg.Timeout, "timeout", 30*time.Second, "per-message SMTP timeout")
	fs.StringVar(&cfg.TLSCAFile, "tls-ca-file", "", "additional CA certificate PEM file for relay TLS")
	fs.StringVar(&cfg.TLSServerName, "tls-server-name", "", "TLS certificate server name (defaults to smtp-host name)")
	fs.StringVar(&cfg.TLSMode, "smtp-tls-mode", "starttls", "SMTP TLS mode: starttls (required by default) or implicit")
	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Print(err)
		os.Exit(2)
	}
	if fs.NArg() != 0 {
		log.Printf("unexpected positional arguments: %v", fs.Args())
		os.Exit(2)
	}
	if *showVersion {
		fmt.Println(version)
		return
	}
	cfg.Password = os.Getenv("SMTP_MCP_PASSWORD")
	server, err := mcpserver.NewServer(cfg, nil)
	if err != nil {
		log.Printf("configure smtp-mcp server: %v", err)
		os.Exit(2)
	}
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		if err != io.EOF {
			log.Printf("run smtp-mcp server: %v", err)
			os.Exit(1)
		}
	}
}
