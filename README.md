# pachyderm-operator

I wanted the ability to apply my pachyderm pipelines with a simple `kubectl apply -k .` command.

## Description

This project is a Kubernetes operator for [Pachyderm](https://www.pachyderm.com/) that allows you to manage your pipelines with Kubernetes manifests.

This is in very early stages of development, will likely never be feature complete, and is not recommended for production use.

## Deploy

```bash
kubectl apply -k config/default
```

## Getting Started

All credit goes to the [Pachyderm](https://www.pachyderm.com/) team for pachyderm itself.
Also some credit to the [operator framework](https://sdk.operatorframework.io/) maintainers.

## License

Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

