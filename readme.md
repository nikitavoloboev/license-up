# License Up [![Thanks](https://img.shields.io/badge/Say%20Thanks-💗-ff69b4.svg)](https://www.patreon.com/nikitavoloboev)
> Create a license quickly for a given name

## Contents
- [Install](#install)
- [Usage](#usage)
	- [MIT License](#mit-license)
- [Demo](#demo)
- [Contributing](#contributing)

## Install
If you have [Go](https://golang.org/doc/install) installed, run:
```Bash
go get -u github.com/nikitavoloboev/license-up
```

## Usage
Use it in the directory where you want to create the `LICENSE` file.

### MIT License
To create [MIT license](https://opensource.org/licenses/MIT), run:

```Bash
license-up mit <name> <surname> <website>
```

Where `<website>` is optional. Here is a working example of it.

```Bash
license-up mit Nikita Voloboev nikitavoloboev.xyz
```

Will create a file `LICENSE` with following content:

```Markdown
MIT License

Copyright (c) 2017 Nikita Voloboev (nikitavoloboev.xyz)

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
alias mit='license-up mit Nikita Voloboev nikitavoloboev.xyz'
```

## Demo
[![asciicast](https://asciinema.org/a/153542.png)](https://asciinema.org/a/153542)

## Contributing
[Suggestions](../../issues/) and pull requests are highly encouraged!

## Related
- [OSS Ninja](https://oss.ninja) - Open Source licenses with just a link.

## Thank you 💜
You can support what I do on [Patreon](https://www.patreon.com/nikitavoloboev) or look into [other projects](https://nikitavoloboev.xyz/projects) I shared.

## License
MIT © [Nikita Voloboev](https://www.nikitavoloboev.xyz)