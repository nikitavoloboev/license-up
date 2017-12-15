# License Up [![Thanks](https://img.shields.io/badge/Say%20Thanks-💗-ff69b4.svg)](https://www.patreon.com/nikitavoloboev)
> Create a license quickly for a given name

Currently only MIT license is supported.

## Installing
If you have [Go](https://golang.org/doc/install) installed, run:

```Bash
go get -u github.com/nikitavoloboev/license-up
```

## Usage 
Use it in the directory where you want to create the `LICENSE` file. The first two arguments to the command are the full name that you want to be attached to the License file and third is the website or email that you wish was attached.

Running this:

```Bash
license-up Nikita Voloboev nikitavoloboev.xyz
```

Will create a file LICENSE with following content:

```Markdown
MIT License

Copyright (c) 2017 Nikita Voloboev, nikitavoloboev.xyz

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
```

If you wish, you can create an [alias](http://tldp.org/LDP/abs/html/aliases.html) to simplify this. For example I have this alias that I use:

```Bash
alias mit='license-up Nikita Voloboev nikitavoloboev.xyz'
```

## Demo
[![asciicast](https://asciinema.org/a/UUyyFWzZRLJFyOaF91vCtiTkR.png)](https://asciinema.org/a/UUyyFWzZRLJFyOaF91vCtiTkR)

## Contributing
[Suggestions](https://github.com/nikitavoloboev/license-up/issues) and pull requests are highly encouraged!

## Thank you 💜
You can support what I do on [Patreon](https://www.patreon.com/nikitavoloboev) or look into [other repositories](https://my.mindnode.com/ZKGETDkUaQUsL3q8q9z788CxG84oEHgDiT79GuzX#-143.5,-902.6,0) I shared. 

## License
MIT © [Nikita Voloboev](https://www.nikitavoloboev.xyz)