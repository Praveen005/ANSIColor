**What Are ANSI Escape Codes?**

Imagine we're writing a message in a text editor. Normally, we type letters, numbers, and punctuation, and they appear on the screen as we expect. However, what if we wanted to do something *more* with that text?

- Change the color to red or blue?
- Make it bold or underlined?
- Move the cursor to a different position on the screen?

This is where ANSI escape codes come in. They're special sequences of characters that tell our terminal (the program we use to interact with our computer's command line) to do something other than simply display text.

**How Do They Work?**

ANSI escape codes start with the Escape character (represented by `\e` or the ASCII code 27), followed by a bracket `[`, and then a series of numbers and letters that specify the command. Here's the basic structure:

```
\e[<command parameters>;<command parameters>m
```

> Note: Some control escape sequences, like `\e` for `ESC`, are not guaranteed to work in all languages and compilers. It is recommended to use the decimal, octal or hex representation as escape code. 
>   
> **Decimal:** `\033` (This is the decimal equivalent of the ASCII value 27)
>
> **Octal:** `\x1B` (This is the hexadecimal equivalent of the ASCII value 27)

**Common Code Sequences:**

Here are a few we can use:

- **Text Styling:**
    - `\e[0m` - Reset all text styles (default)
    - `\e[1m` - Bold
    - `\e[4m` - Underline
    - `\e[7m` - swap foreground and background colors

- **Text Color:**
    - `\e[30m` - Black
    - `\e[31m` - Red
    - `\e[32m` - Green
    - `\e[33m` - Yellow
    - `\e[34m` - Blue
    - `\e[35m` - Magenta
    - `\e[36m` - Cyan
    - `\e[37m` - White

- **Background Color:** (Similar to text color but with `40m` to `47m`)

- **Cursor Movement:**
    - `\e[nA` - Move cursor up *n* lines
    - `\e[nB` - Move cursor down *n* lines
    - `\e[nC` - Move cursor forward *n* columns
    - `\e[nD` - Move cursor backward *n* columns

**Example:**


```
echo -e "\x1b[32mRam Ram Ji\x1b[0m"
```

or

```
echo -e "\033[32mRam Ram Ji\033[0m"
```

We should see "Ram Ram Ji" printed in green!

