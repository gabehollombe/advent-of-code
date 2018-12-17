// Run me via: ./node_modules/.bin/ts-node src/day/01/main.ts

import { Frequency } from './Frequency';
import fs = require('fs');
import path = require('path');


function lines(dir: string, file: string) : string[] {
    return fs
        .readFileSync(path.join(dir, file))
        .toString()
        .split('\n')
        .filter((s => s != ''));
}

const input: string[] = lines(__dirname, 'input.txt');

function part1(): void {

    console.log('Part 1: ', new Frequency(input).part1());
}

function part2(): void {
    const input: string[] = lines(__dirname, 'input.txt');

    console.log('Part 2: ', new Frequency(input).part2());
}

part1();
part2();