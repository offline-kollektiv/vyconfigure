# VyConfigure

Declarative YAML configuration for VyOS

__Note: this project is far from production ready, use at your own risk!__

## Installation

You will need to enable the HTTP API on your VyOS instance, [refer to the upstream documentation for how to configure it.](https://docs.vyos.io/en/latest/configuration/service/https.html)

```vbash
set service https api keys id '<name>' key '<secret>'
set service https api rest
```

You may want to enable additional things, like proper TLS certificates.

[The latest binary is available in releases](https://github.com/offline-kollektiv/vyconfigure/releases).

## Workflow
You should start by syncing your existing configuration to your local filesystem so you can begin using VyConfigure.
```bash
# This will sync your existing VyOS config to your current working directory
vyconfigure sync "<VyOS IP or Hostname>" "<VyOS HTTP API key>"
```

Once the configuration is on your local filesystem, you can preview the changes using
```bash
vyconfigure plan "<VyOS IP or Hostname>" "<VyOS HTTP API key>"
```

If you're happy with the changes, then you can apply them.
```bash
vyconfigure apply "<VyOS IP or Hostname>" "<VyOS HTTP API key>"
```

There are a few flags available:
```
Flags:
    --config-dir string   Directory where config is stored. (default ".")
-h, --help                help for vyconfigure
    --insecure            Whether to skip verifying the SSL certificate.
-v, --version             version for vyconfigure
```

## How does VyConfigure work?
VyConfigure works by using [the VyOS HTTP API](https://docs.vyos.io/en/latest/configuration/service/https.html). It translates the configuration into YAML files and then back to a set of commands when you apply.

## What's the purpose of VyConfigure?
Vyconfigure was initially created by https://github.com/charlie-haley. It was meant to offer a GitOps way of managing VyOS configuration while being more lightweight than Ansible.

## Unsupported features
Currently, configuring users with vyconfigure is explicity blocked due to complexities around encrypted passwords, for now it's recommended you configure these as usual. See https://github.com/offline-kollektiv/vyconfigure/issues/15

[Please raise an issue](https://github.com/offline-kollektiv/vyconfigure/issues) for any issues or proposed features. Contributions also welcome 😊
