import fs = require('fs');
import path = require('path');


export function getInputLines(dir: string, file: string) : string[] {
    return fs
        .readFileSync(path.join(dir, file))
        .toString()
        .split('\n')
        .filter((s => s != ''));
}