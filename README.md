# Ascii-Art-Color

![image](/images/pic.png)


## Overview

The ASCII-Art program converts input text into a graphic representation using ASCII characters and colors the art using the color specified by the user. It supports various fonts and special characters, including spaces, newlines, and punctuation. This tool is ideal for creating stylized text for use in command line applications, documentation, or simply for artistic purposes.

## Features

- Multiple Color codes: you can use hexadecimal color, rgb , hsl and regular color format methods.
- Multiple font : you can use a specific font / banner file 
- Option to color certain characters in the string a specific color 
- Special Character Handling: Processes and displays numbers, letters, spaces, special characters, and \n for newlines.
- Download missing banner files i.e standard.txt, shadow.txt, thinkertoy.txt


## Funtionality
### `OpenFile`

Reads ASCII art from a file specified by filename.
Returns a string representing the ASCII art.

### `DisplayArt`

Prints a string using ASCII art, where words are separated by spaces and lines are separated by \n.
Uses PrintWord to print individual words.
Returns an error if encountered during printing.

### `MapFile`

This go function maps the rune of printable ascii characters from `' '` to `~` that is 95 printable ascii characters to their corresponding art character lines from the art file that is represented as a string
It returns a 2D array of `[rune][]string`

### `AnsiColorMap`
This function maps the a string representation of a color to its equivalent code representation in an ANSI color. It returns a 2D array of `[string]string`

### `ConvertHexToAnsi`
This function converts a hexadecimal string that is a color code and converts it to the nearest ANSI color representation. It parses the red, green and  blue values from the input string, converts them RGB and utilises another function `ConvertRGBToAnsi` to convert the RGB values to their corresponding ANSI color code.

### `ConvertHSLToRGB`
This function converts HSL (Hue, Saturation, Lightness) color values to their corresponding RGB (Red, Green, Blue) values. This allows users to specify colors in the HSL format and translate them into the RGB format. The function uses mathematical calculations based on the HSL color model to determine the RGB components accurately. Additionally, a helper function called abs is utilized to compute the absolute value, facilitating the conversion process. 

### `ConvertRGBToAnsi`
 This function maps normalized RGB values to the nearest ANSI 256 color code, facilitating the representation of colored ASCII art or text in compatible terminals. It normalizes the RGB values and calculates the corresponding ANSI color code, returning the ANSI escape sequence for the specified color. This function enhances the customization of ASCII art by enabling precise color selection using RGB values.

 ### `ExtractRGBValues`
 This function parses an RGB string to extract the red, green, and blue values, ensuring they fall within the valid range of 0 to 255. It returns the extracted RGB values or an error if the format is invalid or values are out of range.

 ### `ExtractHSLValues`
 This function extracts the hue, saturation, and lightness values from an HSL string, ensuring they are within the valid range (hue: 0-360, saturation and lightness: 0-100). It returns the extracted HSL values or an error if the format is invalid or values are out of range.


### `main()`
This function generates ASCII art with partial or full coloring based on command-line inputs. It accepts various argument configurations to customize the output:

- With no arguments, it displays a usage message.
- If provided with only one argument (the input string), it generates ASCII art using the entire input string and resets the color.
- If provided with two arguments (--color=<color> and the input string), it generates ASCII art using the entire input string with the specified color.
- If provided with two arguments (input string and banner), it generates ASCII art using the entire input string with the specified banner/font.
- If provided with three arguments (--color=<color>, letters to be colored, and the input string), it generates ASCII art with the specified letters colored.
- If provided with three arguments (--color=<color>,  input string and banner), it generates ASCII art with the specified  color and banner/font.
- If provided with four arguments (--color=<color>,letters to be colored, input string and the banner), it generates ASCII art with the specified letters colored and in the specified banner.
  
The function extracts the color and banner information from the command-line arguments, reads a template file containing ASCII art, and displays the art with the specified color configuration.


### Flags
- `--color=<color>`: Specifies the color to use for highlighting. Accepts color names from the available colors.

- `"--color=<rgb>"`: Specifies the color in RGB format. Takes a string with rgb values in the format `rgb(0, 0, 0)`between 0 and 255 representing red, green, and blue components. Due to "()" being bash commands ensure the color flag is included in quotation marks to enable the project to work.


- `"--color=<hsl>"`: Specifies the color in HSL format. Takes a string with hsl values in the format `hsl(120, 50%, 50%)` representing hue, saturation, and lightness, respectively. Hue is in the range of 0 to 360, and saturation and lightness are percentages. Due to "()" being bash commands ensure the color flag is included in quotation marks to enable the project to work.

- `--color=<hex>`: Specifies the color in hexadecimal format. Takes a string representing a 6-digit hexadecimal color code.

- `banner`  : Specifies the font/banner file to use. Available fonts are shadow, thinkertoy, standard. Standard is the default font used. When including the banner option write it as the last argument and in all lowercase.

##### Flag usage

```bash
go run . --color=red "Hello"
go run . "--color=rgb(255, 0, 0)" "Hello"
go run . "--color=hsl(0, 100, 50)" "Hello"
go run . "--color=#FF0000" "Hello"
go run . --color=red "hello" 
go run . --color=blue "h" "hello"

```

### Colors supported
When using a `--color=<color>` you can specify the color from the available colors in this list. You can use any of these color names as options for the `--color` flag when running the program.

**They are case sensitive**

- cyan
- grey
- yellow
- black
- red
- green
- blue
- magenta
- white
- pink
- brown
- orange
- lime
- brightBlack
- brightRed
- brightGreen
- brightYellow
- brightBlue
- brightMagenta
- brightCyan
- brightWhite



## Installation

Ensure that Go is installed on your system. You can install Go from the official Go website. Once Go is installed, you can clone and run the program directly from the source code.

 Clone the repository:

``` 
    git clone https://learn.zone01kisumu.ke/git/shfana/ascii-art-color
    cd ascii-art-color
```
## Usage

Run the program by passing the text string as an argument along with the --color option to specify the color. Here's are examples:

```bash
go run . "Hello" 
go run . --color=red  "Hello" 
go run .  "hello" shadow
go run . --color=red "my" "my code"
go run . --color=green  "Hello" shadow
go run . --color=#ff0000  "h" "hello" thinkertoy 
go run . --color=#ff0000 "Hello" 
go run . "--color=rgb(255, 0, 0)" "\n" | cat -E
go run . "--color=hsl(0, 100%, 50%)" "Hello\n" | cat -E
```                                                             
Then you start ,making changes 🏃‍♂️🏃‍♂️🏃‍♂️

Replace "Hello" with your desired text string and "red" and "green" with your preferred color choices.

**Basic Command: No Color Selection**

``` bash
go run .  "Hello" 

```
Output:
```
 _    _          _   _
| |  | |        | | | |
| |__| |   ___  | | | |   ___
|  __  |  / _ \ | | | |  / _ \
| |  | | |  __/ | | | | | (_) |
|_|  |_|  \___| |_| |_|  \___/

```

**Command with Color Selection**


``` bash
go run . --color=green "Hello" 

```
Output:

![command with color selection](/images/basicuse.png)

**Command with Text and Color  Selection:**
``` bash

go run . --color=green  "my" "my code" 

```
Output:

![command with text and color selection](/images/lettersandcolor.png)

**Command with  Hexadecimal Color Selection**


``` bash

go run . --color=#00FF00 "Hello" 

```
Output:

![command with hexadecimal color selection](/images/basicuse.png)

**Command with  HSL Color Selection**


``` bash

go run . "--color=hsl(120, 100%, 50%)" "Hello" 

```
Output:

![command with HSL color selection](/images/basicuse.png)

**Command with  RGB Color Selection**


``` bash

go run . "--color=rgb(0, 255, 0)" "Hello" 

```
Output:

![command with RGB color selection](/images/basicuse.png)

**Command with  banner file and inputstring**


``` bash

go run .  "Hello"  shadow

```
Output:

![command with input string](/images/Screenshot%20from%202024-05-27%2019-06-34.png)

**Command with color, input string and banner**


``` bash

go run . --color=red  "Hello" shadow

```
Output:

![Command with color, input string and banner](/images/Screenshot%20from%202024-05-27%2019-08-02.png)

**Command with color, letters to color, input string and banner**


``` bash

go run . --color=red "H" "Hello" thinkertoy

```
Output:

![Command with color, letters to color, input string and banner](/images/Screenshot%20from%202024-05-27%2019-10-08.png)



**Command with wrong input(wrong number of arguments)**
``` bash
go run . 

```
Output:

```bash
Usage: go run . [OPTION] [STRING] 

EX: go run . --color=<color> <letters to be colored> "something"
```

### Fonts / Banner Format
This project comes with three predefined ASCII fonts, located in banner files:

  - [`shadow`](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/shadow.txt)
  - [`standard`](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/standard.txt)
  - [`thinkertoy`](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/thinkertoy.txt)


**- Each character is represented over 8 lines, providing a clear and sizable output.*
**- Characters are separated by a new line `\n`.*
**-Can only specify font in the code **

## Testing

Unit tests are highly recommended to ensure the integrity and functionality of the code. Test cases can be executed with:

```bash
cd functionfiles 

go test -v
 ```

 or 

 ```bash
cd colorconversions

go test -v
 ```

 ## Learning Outcomes

This project will help you  learn about:

  +  The Go file system(fs) API
  +  Color converters
  +  Data manipulation
  +  Terminal display


## Allowed Packages

Only standard Go packages are allowed in this project. This restriction is to ensure that the utility remains lightweight and dependency-free.

## Resources
### Color codes
- [Hexadecimal](https://www.w3schools.com/colors/colors_hexadecimal.asp): A 6-digit hexadecimal representation of a color, with 2 digits for each of the red, green and blue components. For example, #ff0000 represents the color red.

- [RGB](https://www.w3schools.com/colors/colors_rgb.asp): A comma-separated list of the red, green and blue components of a color, each ranging from 0 to 255. For example, rgb(255, 0, 0) represents the color red.

- [HSL](https://www.w3schools.com/colors/colors_hsl.asp): A comma-separated list of the hue, saturation and lightness components of a color, with the hue ranging from 0 to 360, and the saturation and lightness ranging from 0% to 100%. For example, hsl(0, 100%, 50%) represents the color red.

- [Color Conversion Algorithm](https://en.wikipedia.org/wiki/HSL_and_HSV#Color_conversion_formulae) colors

- [HSL to RGB conversion](https://www.rapidtables.com/convert/color/hsl-to-rgb.html)

- [Color conversion](https://www.hackitu.de/termcolor256/) ANSI colors and conversion from rgb to ansi

- [ANSI colors](https://talyian.github.io/ansicolors/) A list of  ansi colors and their codes 

## Limitations
This project can only color using one specified color only . Multiple color selections  will be added in the future. If you are interested in supporting this feature feel free to submit a pull request.

## Contribution

Contributions are welcome. Please adhere to the existing coding standards and include unit tests for any new features or changes. Ensure to thoroughly test the code before pushing any updates.
If you encounter any issues or have suggestions for improvement, feel free to submit an issue, pull request or propose a change!

## Authors
This project was build and maintained by:

- [shfana](https://learn.zone01kisumu.ke/git/shfana/)
- [mdudi](https://github.com/Dudimath)
- [hokwach](https://learn.zone01kisumu.ke/git/hokwach/)



## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
