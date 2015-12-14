# Gestalt

[![Build Status](https://travis-ci.org/asteris-llc/gestalt.svg)](https://travis-ci.org/asteris-llc/gestalt)

Gestalt sits in front of your K/V stores and enforces schemas.

**Note that the instructional content in this README was written to define
consumer requirements, and is not necessarily all implemented. Please follow
[issue #1](https://github.com/asteris-llc/gestalt/issues/2) for updates.**

<!-- markdown-toc start - Don't edit this section. Run M-x markdown-toc-generate-toc again -->
**Table of Contents**

- [Gestalt](#gestalt)
    - [Deploying Gestalt](#deploying-gestalt)
        - [Running a Server](#running-a-server)
    - [Basic Workflow](#basic-workflow)
        - [Define a Schema](#define-a-schema)
        - [Submit the Schema to Gestalt](#submit-the-schema-to-gestalt)
        - [Use the Schema](#use-the-schema)
    - [License](#license)

<!-- markdown-toc end -->

## Deploying Gestalt

Gestalt is shipped as a single binary that can be both a client and a server. To
see the commands available, run `gestalt`.

### Running a Server

You'll need the source and credential information for any K/V stores you want to
control with Gestalt. In this example, we'll use the [Consul](https://consul.io)
backend by writing the following to `gestalt.toml`:

```toml
port=3000

[dev]
backend=consul
host=http://127.0.0.1:8500
prefix=/dev
```

This means that on startup, there will be a store named "dev", whose keys will
be rooted at `/dev`. Start the server with `gestalt serve gestalt.toml`.

## Basic Workflow

The basic workflow for using Gestalt should be quite simple. You need to:

1. [Define a Schema](#define-a-schema)
2. [Submit the schema to Gestalt](#submit-the-schema-to-gestalt)
3. [Use the Schema](#use-the-schema)

### Define a Schema

Gestalt uses an extension of [JSON Schema](http://json-schema.org/) to specify
schemas (specifically, arrays are not currently supported, and root objects can
have a `backend` key) Let's specify the configuration for a simple app with some
feature flags.

```json
{
    "name": "sample-app",
    "backend": "dev",
    "type": "object",
    "properties": {
        "email-host": {
            "type": "string",
            "required": true
        },
        "features": {
            "type": "object",
            "properties": {
                "an-amazing-feature": {
                    "type": "bool",
                    "default": false
                }
            }
        }
    }
}
```

(this sample is also available at [`sample-schema.json`](sample-schema.json))

In this (admittedly simplified) schema, we have two keys: `email-host`, which is
required, and `an-amazing-feature`, a boolean which defaults to `false`. These
keys will end up in the store at the following locations:

- `email-host`: `/dev/sample-app/email-host`
- `an-amazing-feature`: `/dev/sample-app/features/an-amazing-feature`

If you want to exclude the prefixes from your keys, you can prefix any name with
`/`. Otherwise, a path prefix will be constructed with the backend's prefix and
the schema name.

**Note**: you can define fields that are not required and have no default, but
that will result in the keys not being present in the store. Be aware of this
condition when reading keys back out.

### Submit the Schema to Gestalt

After you [have a server running](#running-a-server), you can use the `gestalt`
CLI tool to submit the schema:

    $ gestalt submit sample-schema.json --host=localhost:3000

This command will set any defaults set in your schema. Use the same command to
update an existing schema. If you don't have access to the `gestalt` tool, you
can also use cURL:

    $ curl -X POST -d @sample-schema.json -H "Content-Type: application/schema+json" http://localhost:3000/v1/schemas/
    { response json elided }

Once you have the schema submitted, you can also use `gestalt` to set the values:

    $ gestalt write sample-app email-host 1.2.3.4 --host=localhost:3000

Or the corresponding cURL:

    $ curl -X PUT -d '"1.2.3.4"' -H "Content-Type: applicatin/json" http://localhost:3000/v1/schemas/sample-app/values/email-host
    { response json elided }

### Use the Schema

Once you have values written to your K/V store through Gestalt, you can use them
with the tools you usually would, but with the added assurance that anytime you
set a value through Gestalt, it will be the right type in the K/V store. You can
also read values through Gestalt directly (see `gestalt read`), but your K/V
store will probably scale to read-heavy workloads much better and have built-in
tools for updates (see
[consul-template](https://github.com/hashicorp/consul-template) for Consul or
[confd](https://github.com/kelseyhightower/confd) for etcd).

## License

Gestalt is licensed under the [Apache 2.0 license](LICENSE).
