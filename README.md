# Alias Gen

Go based API for generating usernames, emails, and passwords for test data. Soon
to include other information for fulfilling tests.

## Requirements

`Go 1.18`

## Free API

At some point, you can expect https://exhausted.club/aliasgen to be a working
API to generate test data, if you don't want to spin up your own version.

## Usage

TBA

## License

```
MIT License

Copyright (c) 2023 Exhausted Club

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

## Original Work

The original work was forked from
https://github.com/edwardr/php-random-name-generator, but this was for the seed
data to generate names. The code aspect was not copied over, and was written
from the ground up using Go. IMHO the license should still be copied over, as
the data from the JSON could be considered falling under the license, and
therefore this project falls in line with MIT, however, IANAL so feel free to
open an issue to correct this if need be, I suppose.

## TODO:

- Potentially remove spaces and `\n` from the resource.json files to save on
  bytes as well as read time. While it might not be significant, might as well.
  Also maybe store them in a format more optimal for reading into memory.
- Swagger gen
- Rate Limiting based off of IP (Must be done before API opening)
- API Keys (Must be done before API opening)
