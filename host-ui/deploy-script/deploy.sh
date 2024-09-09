#!/bin/bash

/bin/bash tool/s3/delete.sh "$1"
/bin/bash tool/s3/copy.sh "$1"
