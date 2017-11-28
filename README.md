# text-slideshow
This was created at the request of https://twitch.tv/typeForced for his stream in order to swap between text files automatically on stream.  

## Usage
Follow the config file [here](etc/config-default.yaml) as an example.  This file can either be placed in your home directory as ".text-slideshow" or you can pass in the config file path via the ```--config``` option on the command line if you don't want to keep it there.

### Linux
Run the command in a terminal via "./text-slideshow" or add it to your $PATH and execute it via "text-slideshow".  You are free to compile it from source using the make file in the repo.  If you want to use a config file from a custom location include the ```--config``` option.

### Windows
Run the command in command prompt via either ".\text-slideshow" or add it to your %PATH% and execute it via "text-slideshow".  If you want to use a config file from a custom location include the ```--config``` option.

## License: MIT
MIT License

Copyright (c) 2017 Sean Smith

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
