package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Mandatory features. Can't remove unless there are replacements.
	_ "github.com/xtls/xray-core/app/dispatcher"
	_ "github.com/xtls/xray-core/app/proxyman/inbound"
	_ "github.com/xtls/xray-core/app/proxyman/outbound"

	// Other optional features.
	_ "github.com/xtls/xray-core/app/dns"
	_ "github.com/xtls/xray-core/app/log"
	_ "github.com/xtls/xray-core/app/policy"
	_ "github.com/xtls/xray-core/app/router"
	_ "github.com/xtls/xray-core/app/stats"

	// Fix dependency cycle caused by core import in internet package
	_ "github.com/xtls/xray-core/transport/internet/tagged/taggedimpl"

	// Inbound and outbound proxies (trimmed: only socks, http, vless, freedom, blackhole)
	_ "github.com/xtls/xray-core/proxy/blackhole"
	_ "github.com/xtls/xray-core/proxy/freedom"
	_ "github.com/xtls/xray-core/proxy/http"
	_ "github.com/xtls/xray-core/proxy/socks"
	_ "github.com/xtls/xray-core/proxy/vless/inbound"
	_ "github.com/xtls/xray-core/proxy/vless/outbound"

	// Transports (trimmed: only tcp, tls, reality, udp)
	_ "github.com/xtls/xray-core/transport/internet/reality"
	_ "github.com/xtls/xray-core/transport/internet/tcp"
	_ "github.com/xtls/xray-core/transport/internet/tls"
	_ "github.com/xtls/xray-core/transport/internet/udp"

	// JSON only (no TOML, no YAML)
	_ "github.com/xtls/xray-core/main/json"

	// Load config from file or http(s)
	_ "github.com/xtls/xray-core/main/confloader/external"

	// Commands
	_ "github.com/xtls/xray-core/main/commands/all"
)
