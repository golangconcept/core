# Uppercase for Exported Fields

If you want the field to be exported (i.e., accessible from other packages), the field name must start with an uppercase letter. This is how Go differentiates between public and private fields.

- `Exported fields` (`public`) must start with a capital letter.
- `Unexported fields` (`private`) must start with a lowercase letter.

## Struct Tags
When working with struct tags (e.g., for JSON serialization or database mapping), the convention is to use lowercase letters for tag names. Struct tags are typically enclosed in backticks **(`)**.

##  Note: Field names must match JSON keys exactly.