# num2words

[![GoDoc](https://godoc.org/github.com/divan/num2words?status.svg)](https://godoc.org/github.com/divan/num2words)

num2words - Numbers to words converter in Go (Golang)

## Usage

First, import package num2words

`import github.com/divan/num2words`

Convert number

```go
  str := num2words.Convert(17) // outputs "seventeen"
  ...
  str := num2words.Convert(1024) // outputs "one thousand twenty four"
  ...
  str := num2words.Convert(-123) // outputs "minus one hundred twenty three"
```

Convert number with " and " between number groups:

```go
  str := num2words.ConvertAnd(514) // outputs "five hundred and fourteen"
  ...
  str := num2words.ConvertAnd(123) // outputs "one hundred and twenty three"
```

## Language Support

Ukrainian language support with gender agreement and noun forms:

```go
  str, _ := num2words.ConvertLang(42, "uk") // outputs "сорок два"
  str, _ := num2words.ConvertLang(42, "en") // outputs "forty-two"
```

Direct Ukrainian API with advanced features:

```go
  import "github.com/divan/num2words/lang/ukrainian"
  
  str := ukrainian.Convert(42) // outputs "сорок два"
  
  // Gender agreement (masculine/feminine/neuter)
  str := ukrainian.ConvertWithOptions(1, ukrainian.Options{
    Gender: ukrainian.Feminine,
  }) // outputs "одна"
  
  // With noun forms (automatically handles singular/plural/genitive)
  str := ukrainian.ConvertWithOptions(2, ukrainian.Options{
    Gender: ukrainian.Feminine,
    WithNoun: true,
    NounForms: &ukrainian.NounForms{
      Singular: "гривня",
      NominativePlural: "гривні", 
      GenitivePlural: "гривень",
    },
  }) // outputs "дві гривні"
```
