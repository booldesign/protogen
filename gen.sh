#!/bin/bash
module=example.com/cloud/protogen
rm -rf kitex_gen
IDL="../protos"
ls $IDL/*/*/service.proto | xargs -I param sh -c "kitex -module $module -type protobuf -I $IDL param"