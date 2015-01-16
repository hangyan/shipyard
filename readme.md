# Shipyard [![Build Status](http://ci.evanhazlett.com/api/badge/github.com/shipyard/shipyard/status.svg?branch=master)](http://ci.evanhazlett.com/github.com/shipyard/shipyard)

Composable Docker Management

Shipyard enables multi-host, Docker cluster management.  It uses the [Citadel](https://github.com/citadel/citadel) toolkit for cluster resourcing and scheduling.  Shipyard has been dramatically simiplified and only requires access to the Docker Remote API and a RethinkDB instance.

# Origin  Documentation
See the origin README documentation at https://github.com/shipyard/shipyard

# Improvment

* package import replace (because of Fucking GFW)
* obtain engines'cpu number and total memory automatically
* add interactive,tty,detached options to run command
* remove usage report
* add "all" option to "containers(ps)" command.
* fix "events" panic bug and change date format
* fix scale problems
* ...

# License
Shipyard is licensed under the Apache License, Version 2.0. See LICENSE for full license text.
