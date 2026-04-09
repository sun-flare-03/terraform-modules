# terraform-modules

[![Build Status](https://img.shields.io/github/actions/workflow/status/user/terraform-modules/ci.yml?branch=main)]()
[![Go Version](https://img.shields.io/badge/go-1.22+-blue.svg)]()
[![License: MIT](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

**terraform-modules** reusable Terraform modules for AWS, GCP, and Azure infrastructure. Built with simplicity and performance in mind.

## Features

- Graceful Shutdown: Clean shutdown handling with configurable drain timeout
- Minimal Allocations: Designed to minimize GC pressure in hot paths
- Zero Dependencies: No external packages required for core functionality
- Context Support: Full context.Context propagation for cancellation and deadlines

## Installation

```bash
go get github.com/user/terraform-modules@latest
```

## Quick Start

```go
package main

import (
	"fmt"
	"github.com/user/terraform-modules"
)

func main() {
	client := terraformmodules.New(
		terraformmodules.WithTimeout(30 * time.Second),
	)

	result, err := client.Run(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
```

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.
