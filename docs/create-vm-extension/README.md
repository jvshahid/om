<!--- This file is autogenerated from the files in docsgenerator/templates/create-vm-extension --->
&larr; [back to Commands](../README.md)

# `om create-vm-extension`

The `create-vm-extension` command will create or update an existing vm extension.

## Command Usage
```
Usage:
  om [OPTIONS] create-vm-extension [create-vm-extension-OPTIONS]

This creates/updates a VM extension

Application Options:
      --ca-cert=               OpsManager CA certificate path or value
                               [$OM_CA_CERT]
  -c, --client-id=             Client ID for the Ops Manager VM (not required
                               for unauthenticated commands) [$OM_CLIENT_ID]
  -s, --client-secret=         Client Secret for the Ops Manager VM (not
                               required for unauthenticated commands)
                               [$OM_CLIENT_SECRET]
  -o, --connect-timeout=       timeout in seconds to make TCP connections
                               (default: 10) [$OM_CONNECT_TIMEOUT]
  -d, --decryption-passphrase= Passphrase to decrypt the installation if the
                               Ops Manager VM has been rebooted (optional for
                               most commands) [$OM_DECRYPTION_PASSPHRASE]
  -e, --env=                   env file with login credentials
  -p, --password=              admin password for the Ops Manager VM (not
                               required for unauthenticated commands)
                               [$OM_PASSWORD]
  -r, --request-timeout=       timeout in seconds for HTTP requests to Ops
                               Manager (default: 1800) [$OM_REQUEST_TIMEOUT]
  -k, --skip-ssl-validation    skip ssl certificate validation during http
                               requests [$OM_SKIP_SSL_VALIDATION]
  -t, --target=                location of the Ops Manager VM [$OM_TARGET]
      --trace                  prints HTTP requests and response payloads
                               [$OM_TRACE]
  -u, --username=              admin username for the Ops Manager VM (not
                               required for unauthenticated commands)
                               [$OM_USERNAME]
      --vars-env=              load vars from environment variables by
                               specifying a prefix (e.g.: 'MY' to load
                               MY_var=value) [$OM_VARS_ENV]
  -v, --version                prints the om release version

Help Options:
  -h, --help                   Show this help message

[create-vm-extension command options]
      -n, --name=              VM extension name
      -c, --config=            path to yml file for configuration (keys must
                               match the following command line flags)
          --vars-env=          load variables from environment variables
                               matching the provided prefix (e.g.: 'MY' to load
                               MY_var=value) [$OM_VARS_ENV]
      -l, --vars-file=         load variables from a YAML file
      -v, --var=               load variable from the command line. Format:
                               VAR=VAL
      -o, --ops-file=          YAML operations file
          --cloud-properties=  cloud properties in JSON format
```

### Configuring via file

#### Example YAML:
```yaml
vm-extension-config:
  name: some_vm_extension
  cloud_properties:
    source_dest_check: false
```

#### Variables

The `create-vm-extension` command now supports variable substitution inside the config template:

```yaml
# config.yml
vm-extension-config:
  name: some_vm_extension
  cloud_properties:
    source_dest_check: ((enable_source_dest_check))
```

Values can be provided from a separate variables yaml file (`--vars-file`) or from environment variables (`--vars-env`).

To load variables from a file use the `--vars-file` flag.

```yaml
# vars.yml
enable_source_dest_check: false
```

```
om create-vm-extension \
  --config config.yml \
  --vars-file vars.yml
```

To load variables from a set of environment variables, specify the common
environment variable prefix with the `--vars-env` flag.

```
OM_VAR_enable_source_dest_check=false om create-vm-extension \
  --config config.yml \
  --vars-env OM_VAR
```

The interpolation support is inspired by similar features in BOSH. You can
[refer to the BOSH documentation](https://bosh.io/docs/cli-int/) for details on how interpolation
is performed.