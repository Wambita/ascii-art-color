# ASCII Art with Colors
![image](https://www.itsupportwale.com/blog/wp-content/uploads/2020/01/figlet-toilet-featured-image.jpg)

## Authors
- [mdudi](https://github.com/Dudimath)
- hokwach
- shfana


### Overview

ASCII Art with Colors is a Go program that generates graphical representations of text using ASCII characters and allows you to color specific letters or the entire string using various color notations such as ANSI, RGB, and HSL.

## Features

- Colorize ASCII art strings with various color systems: RGB, Hex, HSL, and ANSI.

- Option to colorize specific letters or the entire string.
Clear and concise usage message for easy understanding.

## Graphic Representation Using ASCII

A graphic representation using ASCII means representing the input string using ASCII characters. Here's an example:
```bash
@@@@@@BB@@@@``^^``^^``@@BB$$@@BB$$
@@%%$$$$^^^^WW&&8888&&^^""BBBB@@@@
@@@@@@""WW8888&&WW888888WW``@@@@$$
BB$$``&&&&WWWW8888&&&&8888&&``@@@@
$$``&&WW88&&88&&&&8888&&88WW88``$$
@@""&&&&&&&&88888888&&&&&&88&&``$$
``````^^``^^^^^^````""^^``^^``^^``
""WW^^@@@@^^``````^^BB@@^^``^^&&``
^^&&^^@@````^^``&&``@@````^^^^&&``
``WW&&^^""``^^WW&&&&""``^^^^&&88``
^^8888&&&&&&WW88&&88WW&&&&88&&WW``
@@``&&88888888WW&&WW88&&88WW88^^$$
@@""88&&&&&&&&888888&&``^^&&88``$$
@@@@^^&&&&&&""``^^^^^^8888&&^^@@@@
@@@@@@^^888888&&88&&&&MM88^^BB$$$$
@@@@@@BB````&&&&&&&&88""``BB@@BB$$
$$@@$$$$$$$$``````````@@$$@@$$$$$$
```
This project can handle inputs with numbers, letters, spaces, special characters, and `\n`.

## Instructions

- The project must be written in Go.
- The code must follow good practices.
- Unit test files are recommended.
- Three banner files with specific graphical templates represented using ASCII are provided:
  - [`shadow`](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/shadow.txt)
  - [`standard`](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/standard.txt)
  - [`thinkertoy`](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/thinkertoy.txt)

## Banner Format

- Each character has a height of 8 lines.
- Characters are separated by a new line `\n`.
## Funtionality
### [`CreateMap`](https://learn.zone01kisumu.ke/git/mdudi/ascii-art/src/branch/master/functionFiles/createmap.go)

Reads ASCII art from a file specified by filename.
Returns a 2D slice representing the ASCII art.

### [`DisplayArt`](https://learn.zone01kisumu.ke/git/mdudi/ascii-art/src/branch/master/functionFiles/displaymap.go)

Prints a string using ASCII art, where words are separated by spaces and lines are separated by \n.
Uses PrintWord to print individual words.
Returns an error if encountered during printing.

### [`IsPrintable`](https://learn.zone01kisumu.ke/git/mdudi/ascii-art/src/branch/master/functionFiles/isprintable.go)

This Go function, IsPrintable, determines whether a given string contains printable ASCII characters, accounting for escape sequences such as `\t, \a, \b, \v, \f, and \r.`

### [`main()`](https://github.com/Dudimath/ascii_art/blob/main/programme/main.go)
Parses command-line flags -shadow and -thinkertoy to determine which ASCII art file to use.
Prints the ASCII art based on the input string and chosen ASCII art file.
This code uses command-line flags to select different ASCII art files (shadow.txt, thinkertoy.txt, or standard.txt) and prints the input string using ASCII art.

## Usage

To run the program, use the following commands:

```bash
go run . "" | cat -e
go run . "\n" | cat -e
go run . "Hello\n" | cat -e
go run . "hello" | cat -e
go run . "HeLlO" | cat -e
go run . "Hello There" | cat -e
go run . "1Hello 2There" | cat -e
go run . "{Hello There}" | cat -e
go run . "Hello\nThere" | cat -e
go run . "Hello\n\nThere" | cat -e
```
## Flags
- `shadow`: Use shadow.txt for ASCII art.
- `thinkertoy`: Use thinkertoy.txt for ASCII art.
```bash
go run . "Hello" -shadow
go run . "Hello" -thinkertoy
```
## Output Examples
### - Standard Text File
```bash
$ go run . "mdudi\nhokwach\nshfana" | cat -e
                 _               _   _  $
                | |             | | (_) $
 _ __ ___     __| |  _   _    __| |  _  $
| '_ ` _ \   / _` | | | | |  / _` | | | $
| | | | | | | (_| | | |_| | | (_| | | | $
|_| |_| |_|  \__,_|  \__,_|  \__,_| |_| $
                                        $
                                        $
 _                                                _      $
| |              _                               | |     $
| |__     ___   | | _  __      __   __ _    ___  | |__   $
|  _ \   / _ \  | |/ / \ \ /\ / /  / _` |  / __| |  _ \  $
| | | | | (_) | |   <   \ V  V /  | (_| | | (__  | | | | $
|_| |_|  \___/  |_|\_\   \_/\_/    \__,_|  \___| |_| |_| $
                                                         $
                                                         $
       _        __                          $
      | |      / _|                         $
 ___  | |__   | |_    __ _   _ __     __ _  $
/ __| |  _ \  |  _|  / _` | | '_ \   / _` | $
\__ \ | | | | | |   | (_| | | | | | | (_| | $
|___/ |_| |_| |_|    \__,_| |_| |_|  \__,_| $
                                            $
                                            $
```
### - Shadow Text File

```bash
$ go run . "mdudi\nshfana" -shadow | cat -e
                                             $
                     _|                _| _| $
_|_|_|  _|_|     _|_|_| _|    _|   _|_|_|    $
_|    _|    _| _|    _| _|    _| _|    _| _| $
_|    _|    _| _|    _| _|    _| _|    _| _| $
_|    _|    _|   _|_|_|   _|_|_|   _|_|_| _| $
                                             $
                                             $
                                                      $
         _|           _|_|                            $
  _|_|_| _|_|_|     _|       _|_|_| _|_|_|     _|_|_| $
_|_|     _|    _| _|_|_|_| _|    _| _|    _| _|    _| $
    _|_| _|    _|   _|     _|    _| _|    _| _|    _| $
_|_|_|   _|    _|   _|       _|_|_| _|    _|   _|_|_| $
                                                      $
                                                      $
```
### -Thinkertoy Text File

```bash
$ go run . "mdudi\nshfana\nokwach" -thinkertoy | cat -e
                       $
         o         o   $
         |         | o $
o-O-o  o-O o  o  o-O   $
| | | |  | |  | |  | | $
o o o  o-o o--o  o-o | $
                       $
                       $
                             $
    o     o-o                $
    |     |                  $
o-o O--o -O-   oo  o-o   oo  $
 \  |  |  |   | |  |  | | |  $
o-o o  o  o   o-o- o  o o-o- $
                             $
                             $
                                  $
    o                        o    $
    | /                      |    $
o-o OO   o   o   o  oo   o-o O--o $
| | | \   \ / \ /  | |  |    |  | $
o-o o  o   o   o   o-o-  o-o o  o $
                                  $
                                  $
```

## Color Code types
- **Hexadecimal**: A 6-digit hexadecimal representation of a color, with 2 digits for each of the red, green and blue components. For example, #ff0000 represents the color red.

- **RGB**: A comma-separated list of the red, green and blue components of a color, each ranging from 0 to 255. For example, rgb(255, 0, 0) represents the color red.

- **HSL**: A comma-separated list of the hue, saturation and lightness components of a color, with the hue ranging from 0 to 360, and the saturation and lightness ranging from 0% to 100%. For example, hsl(0, 100%, 50%) represents the color red.
                                                             
## Conclusion
This project helps you learn about:

- Go file system (fs) API.
- Data manipulation.

If you encounter any issues or have suggestions for improvement, feel free to submit an issue or propose a change!

To get started working on the ASCII Art project, follow these steps:

Clone the Repository: First, clone the repository to your local machine using Git. You can do this by running the following command in your terminal:
bash
```bash
git clone https://learn.zone01kisumu.ke/git/mdudi/ascii-art.git
```
Navigate to the Project Directory: Move into the project directory using the cd command:

```bash
cd ascii-art
```
Then you start ,making changes 🏃‍♂️🏃‍♂️🏃‍♂️
