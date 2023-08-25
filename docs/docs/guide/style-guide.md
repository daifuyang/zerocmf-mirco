---
nav: 指南
group:
  title: 基础
  order: 1
order: -1
---

# 代码风格指南

请阅读：https://google.github.io/styleguide/go/

## 命名

### 避免重复

在为函数或方法选择名称时，请考虑读取名称的上下文。请考虑以下建议，以避免出现过多的重复：

- 函数和方法名称中通常可以省略以下内容：

  - 输入和输出的类型（当不存在冲突时）
  - 方法接收者的类型
  - 输入或输出是否为指针

- 对于函数，请不要重复包的名称：  
  在包内定义函数时，避免在函数名称中包含包名。相反，使用简洁的名称来描述函数的目的，避免不必要的重复。这种做法使代码更具可读性，避免了不必要的冗余。

```go
// Bad:
package yamlconfig

func ParseYAMLConfig(input string) (*Config, error)
```

```go
// Good:
package yamlconfig

func Parse(input string) (*Config, error)
```

- 对于方法，请不要重复方法接收器的名称：  
  同样，在类型上定义方法时，方法名称应该简洁而富有表现力，不要重复类型的名称。方法接收器的名称已经提供了上下文，因此方法名称本身可以更加专注。

```go
// Bad:
func (c *Config) WriteConfigTo(w io.Writer) (int64, error)
```

```go
// Good:
func (c *Config) WriteTo(w io.Writer) (int64, error)
```

- 不要重复传递的变量名称：  
  在为函数命名参数时，通过使用简洁而有意义的名称来避免冗余。由于参数的类型已经传达了其目的，因此不需要在变量名中包含该信息。

```go
// Bad:
func OverrideFirstWithSecond(dest, source *Config) error
```

```go
// Good:
func Override(dest, source *Config) error
```

- 不要重复返回值的名称和类型：  
  类似于上面的观点，函数返回值应具有简洁的名称，以反映其含义，而不是重复表示它们的类型。

```go
// Bad:
func TransformYAMLToJSON(input *Config) *jsonconfig.Config
```

```go
// Good:
func Transform(input *Config) *jsonconfig.Config
```

- 当有必要区分具有类似名称的函数时，可以包含额外信息：  
  在存在多个功能类似但目的不同的函数的情况下，可以为名称添加额外信息，以提高清晰度。

```go
// Good:
func (c *Config) WriteTextTo(w io.Writer) (int64, error)
func (c *Config) WriteBinaryTo(w io.Writer) (int64, error)
```

### 命名约定

在为函数和方法选择名称时，还有一些其他常见的约定：

- 函数返回某些内容时，通常使用名词样式的名称。

```go
// Good:
func (c *Config) JobName(key string) (value string, ok bool)
```

所以函数和方法的名称应避免使用前缀 "Get"。

```go
// Bad:
func (c *Config) GetJobName(key string) (value string, ok bool)
```

- 执行某些操作的函数通常使用动词样式的名称。

```go
// Good
func (c *Config) WriteDetail(w io.Writer) (int64, error)
```

- 当相同的函数只有类型不同时，可以将类型名称添加到名称的末尾。

```go
// Good
func ParseInt(input string) (int, error)
func ParseInt64(input string) (int64, error)
func AppendInt(buf []byte, value int) []byte
func AppendInt64(buf []byte, value int64) []byte
```

如果存在一个明确的 "主要" 版本，可以省略该版本的名称中的类型：

```go
// Good:
func (c *Config) Marshal() ([]byte, error)
func (c *Config) MarshalText() (string, error)
```
