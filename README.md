# License Up [![Thanks](http://bit.ly/saythankss)](https://patreon.com/nikitavoloboev)

> Create a license quickly

##### Contents

-   [Install](#install)
-   [Usage](#usage)
    -   [MIT License](#mit-license)
    -   [BSD License](#bsd-license)
-   [Contributing](#contributing)

## Install

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

Copyright (c) 2017-present, Nikita Voloboev (nikitavoloboev.xyz)

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

### BSD License

To create [BSD license](https://opensource.org/licenses/BSD-3-Clause), run:

```Bash
license-up bsd <name> <surname>
```

Where `<website>` is optional. Here is a working example of it.

```Bash
license-up bsd Nikita Voloboev
```

Will create a file `LICENSE` with following content:

```Markdown
BSD 3-Clause License

Copyright (c) 2019, Nikita Voloboev
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

* Neither the name of the copyright holder nor the names of its
  contributors may be used to endorse or promote products derived from
  this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
```

If you wish, you can create an [alias](http://tldp.org/LDP/abs/html/aliases.html) to simplify this. For example I have this alias that I use:

```Bash
alias bsd='license-up bsd Nikita Voloboev'
```

## Contributing

See [contribution guidelines](CONTRIBUTING.md#readme).

## Related

-   [OSS Ninja](https://oss.ninja) - Open Source licenses with just a link.

## Thank you

You can support me on [Patreon](https://www.patreon.com/nikitavoloboev) or look into [other projects](https://nikitavoloboev.xyz/projects) I shared.

[![MIT](https://img.shields.io/badge/license-MIT-0a0a0a.svg?style=flat&colorA=0a0a0a)](LICENSE) [![Twitter](http://bit.ly/nikitweet)](https://twitter.com/nikitavoloboev)
