#!/bin/bash

terraform -chdir="$(pwd)/backend" init
terraform -chdir="$(pwd)/backend" apply

terraform -chdir="$(pwd)/ecs-infrastructure" init
terraform -chdir="$(pwd)/ecs-infrastructure" apply
