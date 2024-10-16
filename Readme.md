# ets Experimental typesetting system

This software repository contains a JavaScript frontend for the typesetting library [“Boxes and Glue”](https://github.com/boxeandglue/boxesandglue) which is an algorithmic typesetting machine in the spirits of TeX.

## Installation

You can get a precompiled binary from [the releases page](https://github.com/boxesandglue/ets/releases) or build the software yourself (see below).


## Documentation

The documentation is available at <https://boxesandglue.dev/ets/>.

## Build

You just need Go installed on your system, clone the repository and run

    go build -o bin/ets github.com/boxesandglue/ets/cmd

If you have [Rake](https://ruby.github.io/rake/) installed, you can type `rake build`

## Status

This software is more or less a demo of the architecture and not usable for any serious purpose yet. Once all the basic examples work fine, the setting will stabilize.

Feedback welcome!


## Sample code

```js
bag = require("bag:backend/bag")
fe = require("bag:frontend")


const f = fe.new("out.pdf")

const str = `In olden times when wishing still helped one, there lived a king whose
daughters were all beautiful; and the youngest was so beautiful that the sun itself,
which has seen so much, was astonished whenever it shone in her face.
Close by the king’s castle lay a great dark forest, and under an old lime-tree in the
forest was a well, and when the day was very warm, the king’s child went out into the
forest and sat down by the side of the cool fountain; and when she was bored she
took a golden ball, and threw it up on high and caught it; and this ball was her
favorite plaything.`.replace(/\s+/g, ' ').trim()

ff = f.newFontFamily("text")
ff.addMember( new fe.fontSource( {location: "CrimsonPro-Bold.ttf"}), fe.fontWeight700, fe.fontStyleNormal)
ff.addMember( new fe.fontSource( {location: "CrimsonPro-Regular.ttf"}), fe.fontWeight400, fe.fontStyleNormal)
const para = fe.newText()

para.settings[fe.settingSize] = 12 * bag.factor
para.items.push(str);

ret = f.formatParagraph(para, bag.mustSP("125pt"), fe.leading(14 * bag.factor), fe.family(ff))
p = f.doc.newPage()
p.outputAt(bag.mustSP("1cm"), bag.mustSP("26cm"), ret[0])
p.shipout()
f.doc.finish()
```

Now run (on the command line) `ets myfile.js` which creates a PDF (`out.pdf`) and a log file (`myfile-protocol.xml`).


## Examples

There is a [separate section with examples on GitHub](https://github.com/boxesandglue/boxesandglue-examples/tree/main/ets).

## Contact and information

License: 3 clause BSD<br>
Contact: Patrick Gundlach <gundlach@speedata.de>
