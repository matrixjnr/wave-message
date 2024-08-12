# Wave-Message Library

## Overview

The wave-message library is a data structure used by the Wave application to send data efficiently and securely.

## Features
Flexibility: Supports various data types in messages, including strings and complex structures, with easy serialization and deserialization.

## Installation

To install the library, run the following command:

```bash
go get github.com/yourusername/wave-message

```

## Usage

### Instantiating a Message

You can create a Message using the MessagingService.

```go
package main

import (
	"fmt"
	"wave-message/internal/service"
)

func main() {
	msgService := service.NewMessageService()
	msg, err := msgService.CreateMessage("channel-1", "sender-1", "Hello, World!", true)
	if err != nil {
		fmt.Println("Error creating message:", err)
		return
	}
	fmt.Println("Message created:", msg)
}
    
```

### Serializing and Deserializing a Message

The library provides methods to serialize and deserialize data:

```go
// Serialize
data := service.SerializeData(message)

// Deserialize
deserialized := service.DeserializeData(data)
```

### Testing

Run the tests using the Makefile:

```bash
make test
```

### Building

Build the project using the Makefile:

```bash
make build
```

### Cleaning

Clean up the generated files:

```bash
make clean
```

### Versioning

The current version of the library is 0.1.0. To check the version, use:

```bash
make version
```

### Contributing

see [CONTRIBUTING.md](CONTRIBUTING.md)

### License
This project is licensed under the MIT License.