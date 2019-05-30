#!/usr/bin/env sh

# Copyright 2019 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the License);
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an AS IS BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -x

go test --race -v ./...

ret=$?
set +x
if [ $ret -eq 0 ]; then
  echo "\e[32mSuccess\e[0m" 1>&2
  exit 0
else
  echo "\e[31mFailure\e[0m: $ret" 1>&2
  exit $ret
fi