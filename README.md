# Ascii-Art-Color

![image](https://www.itsupportwale.com/blog/wp-content/uploads/2020/01/figlet-toilet-featured-image.jpg)


### Overview

The ASCII-Art program converts input text into a graphic representation using ASCII characters and colors the art using the color specified by the user. It supports various fonts and special characters, including spaces, newlines, and punctuation. This tool is ideal for creating stylized text for use in command line applications, documentation, or simply for artistic purposes.

### Features

- Multiple Color codes: you can use hexadecimal color, rgb , hsl and regular color format methods.
- Option to color certain characters in the string a specific color 
- Special Character Handling: Processes and displays numbers, letters, spaces, special characters, and \n for newlines.


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
This function converts a hexadecimal string that is a color code and converts it to the nearest ANSI color representation. It parses the red, green and  blue values from the input string and converts them RGB and utilises anothe function to convert the RGB values to their corresponding ANSI color code.

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
- If provided with three arguments (--color=<color>, letters to be colored, and the input string), it generates ASCII art with the specified letters colored.
  
The function extracts the color information from the command-line arguments, reads a standard template file containing ASCII art, and displays the art with the specified color configuration.


## Usage

To run the program, use the following commands:

```bash
go run . "--color=#ff0000" "" 
go run . "--color=rgb(255, 0, 0)" "\n" 
go run . "--color=hsl(0, 100%, 50%)" "Hello\n" 
go run . --color=blue "hello" 
go run . --color=green "HeLlO" 
go run . --color=purple "Hello There" 
go run . --color=yellow "1Hello 2There" 
go run . "--color=#00ffff" "{Hello There}" 
go run . --color=orange "Hello\nThere" 
go run . --color=gray "Hello\n\nThere"
go run . --color=red "my" "my code"


```
### Flags
- `--color=<color>`: Specifies the color to use for highlighting. Accepts color names from the available colors.

- `--color=<rgb>`: Specifies the color in RGB format. Takes a string with rgb values in the format `rgb(0, 0, 0)`between 0 and 255 representing red, green, and blue components.

- `--color=<hsl>`: Specifies the color in HSL format. Takes a string with hsl values in the format `hsl(120, 50, 50)` representing hue, saturation, and lightness, respectively. Hue is in the range of 0 to 360, and saturation and lightness are percentages.

- `--color=<hex>`: Specifies the color in hexadecimal format. Takes a string representing a 6-digit hexadecimal color code.

```bash
go run . --color=red "Hello"
go run . "--color=rgb(255, 0, 0)" "Hello"
go run . "--color=hsl(0, 100, 50)" "Hello"
go run . "--color=#FF0000" "Hello"

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
- magneta
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



### Installation

Ensure that Go is installed on your system. You can install Go from the official Go website. Once Go is installed, you can clone and run the program directly from the source code.

    Clone the repository:

   ``` bash

    git clone https://learn.zone01kisumu.ke/git/shfana/ascii-art-color
    cd ascii-art-color
```
### Usage

Run the program by navigating to ascciart directory and passing the text string as an argument along with the --color option to specify the color. Here's an example:

```bash
cd ascciart/
go run . --color=red  "Hello"
go run . --color=green "Hello" 

```                                                             
Then you start ,making changes 🏃‍♂️🏃‍♂️🏃‍♂️

Replace "Hello" with your desired text string and "red" and "green" with your preferred color choices.

**Basic Command: No Color Selection**

``` bash
cd ascciart/
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

![basic command](/bannerfiles/images/basicuse.png)

**Command with Text and Color  Selection:**
``` bash
go run . --color=green  "my" "my code" 

```
Output:

![basic command](/bannerfiles/images/lettersandcolor.png)

**Command with  Hexadecimal Color Selection**


``` bash
go run . "--color=00FF00" "Hello" 

```
Output:

![basic command](/bannerfiles/images/basicuse.png)

**Command with  HSL Color Selection**


``` bash
go run . "--color=hsl(120, 100, 50)" "Hello" 

```
Output:

![basic command](/bannerfiles/images/basicuse.png)

**Command with  RGB Color Selection**


``` bash
go run . "--color=hsl(0, 128, 0)" "Hello" 

```
Output:

![basic command](/bannerfiles/images/basicuse.png)

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

### Development Notes

### Testing

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

### Allowed Packages

Only standard Go packages are allowed in this project. This restriction is to ensure that the utility remains lightweight and dependency-free.

### Contribution

Contributions are welcome. Please adhere to the existing coding standards and include unit tests for any new features or changes. Ensure to thoroughly test the code before pushing any updates.
If you encounter any issues or have suggestions for improvement, feel free to submit an issue or propose a change!

### Limitations
This project can only use the standard font, in this case other fonts are not supported. This will be added in the future. If you are interested in supporting this feature feel free to submit a pull request.

### License

This project is licensed under the MIT License - see the LICENSE file for details.

## Resources
### Color codes
- [Hexadecimal](https://www.w3schools.com/colors/colors_hexadecimal.asp): A 6-digit hexadecimal representation of a color, with 2 digits for each of the red, green and blue components. For example, #ff0000 represents the color red.

- [RGB](https://www.w3schools.com/colors/colors_rgb.asp): A comma-separated list of the red, green and blue components of a color, each ranging from 0 to 255. For example, rgb(255, 0, 0) represents the color red.

- [HSL](https://www.w3schools.com/colors/colors_hsl.asp): A comma-separated list of the hue, saturation and lightness components of a color, with the hue ranging from 0 to 360, and the saturation and lightness ranging from 0% to 100%. For example, hsl(0, 100%, 50%) represents the color red.
- [Color Conversion Algorithm](https://en.wikipedia.org/wiki/HSL_and_HSV#Color_conversion_formulae)


## Authors
This project was build and maintained by:
- [mdudi](https://github.com/Dudimath)
- [hokwach](https://learn.zone01kisumu.ke/git/hokwach/)
- [shfana](https://learn.zone01kisumu.ke/git/shfana/)

## License 
This file is licensed under the MIT License. See LICENSE for details