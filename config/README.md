# Configuration Module Documentation

This Go module provides a singleton configuration object `C()` that encapsulates various configurations essential for your application. The configurations are filled during the package initialization (`init()`) by loading from `config.yml`, `Default()`, and environment variables.

## Usage

### Installation

To use this, import the package in your code:

```go
import "your_module_path/config"
```

### Configuration Options

The package allows setting up configurations in three ways:

1. **Environment Variables**
2. **Default Settings via `Default()`**
3. **Initialization from `config.yml`**

### 1. Using Environment Variables

💡 Use this way for storing `important secrets` that should not be hard coded in repo or in `config.yml` file. For example, `DB_PASSWORD`, or `JWT_SECRET`.

Environment variables are used to configure the application at runtime. The variable names must start with `ORMUS_`, and nested variables should use `__` as a separator (`ORMUS_DB__HOST`).

Example setting environment variables:
```go
os.Setenv("ORMUS_DEBUG", "true")
os.Setenv("ORMUS_MULTI_WORD_VAR", "this is a multi-word var")
// ... set other environment variables ...
```

### 2. Initialization from `config.yml`

💡 Store variables which are `dependant to the environment` that code is running or the area, the variables that `change more frequent`.

The package supports loading configurations from a YAML file named `config.yml`. Ensure the YAML file is in the correct format.

Example `config.yml` file:
```yaml
debug: true
multi_word_var: "this is a multi-word var in YAML"
db: 
  username: ormus
  password: youcannotbreakme
# ... other configurations ...
```

### 3. Default Settings via `Default()`

💡 Store variables which they have `the least probability of change`.
The `Default()` function in the package allows defining default configurations that act as fallbacks. This function should return a `Config` struct.

Example of defining default configurations:
```go
// Define default configurations
func Default() config.Config {
    return config.Config{
        // Define your default configurations here
        Debug:        true,
        MultiWordVar: "default value",
		Manager: manager.Config{
            JWTConfig: auth.JwtConfig{
                SecretKey: "the_most_secure_secret_of_all_time",
            },
        },
        // ... other default configurations ...
    }
}
```

### Accessing Configuration

#### Accessing the Configuration Singleton

To access the configuration object:

```go
// Get the configuration object
cfg := config.C()
```

### Adding New Configurations

For adding new configurations, update the `Config` struct in the package and ensure it is filled during the initialization process in the `init()` function. Following this, access the updated configuration using the `C()` function.

### Example

Here's an example demonstrating how to access the configuration object and add new configurations:

```go
package main

import (
    "fmt"
    "your_module_path/config"
)

func main() {
    // Access the configuration object
    loadedConfig := config.C()

    // Access existing configurations
    fmt.Println("Debug mode:", loadedConfig.Debug)
    fmt.Println("Multi-word var:", loadedConfig.MultiWordVar)
    // ... access other configurations ...

    // Add new configurations (modify the Config struct in the package)
    loadedConfig.NewConfig = "new value"

    // Access the newly added configuration
    fmt.Println("New Config:", loadedConfig.NewConfig)
}
```

### Important Notes

- The `Config` object, accessed via `C()`, encapsulates all configurations set during the package initialization.
- To add or modify configurations, update the `Config` struct and ensure it is filled during the initialization process (`init()`).
- The `C()` function returns a singleton instance of the configuration object filled by the `init()` function, consolidating settings from `config.yml`, `Default()`, and environment variables.

## Conclusion

This configuration module provides a convenient singleton object encapsulating configurations filled from `config.yml`, `Default()`, and environment variables during initialization. Use `C()` to access the configurations and extend the `Config` struct to incorporate additional settings as needed.
