<!--- This file is autogenerated from the files in docsgenerator/templates/vm-lifecycle --->
&larr; [back to Commands](../README.md)

# `om vm-lifecycle`

commands to manage the state of the Ops Manager VM. Requires the cli of the
desired IAAS to be installed.

## Command Usage
```
Usage:
  om [OPTIONS] vm-lifecycle <command>

commands to manage the state of the Ops Manager VM. Requires the cli of the
desired IAAS to be installed.

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

Available commands:
  create-vm                   Create VM for Ops Manager to a given IaaS
  delete-vm                   Delete VM from a given IaaS
  export-opsman-config        Exports an opsman.yml for an existing Ops Manager VM
  prepare-tasks-with-secrets  Modifies task files to include config secrets as environment variables
  upgrade-opsman              Deletes the old opsman vm given an exported installation, bring up a new vm, and import that installation
```

