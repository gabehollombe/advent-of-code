// This is what I'm working on:
// https://adventofcode.com/2018/day/3

import { Coordinate, Fabric } from './Fabric';

describe('Tracking cells claimed', function() {
    it('counts the number of overlapping cells', () => {
        // #1 @ 1,3: 4x4
        // #2 @ 3, 1: 4x4
        // #3 @ 5, 5: 2x2

        let fabric = new Fabric();

        fabric.claim(new Coordinate(1,3), [4, 4]);
        fabric.claim(new Coordinate(3,1), [4, 4]);
        fabric.claim(new Coordinate(5,5), [2, 2]);

        expect(fabric.overlapCount()).toEqual(4);
        
    });
});