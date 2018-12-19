// This is what I'm working on:
// https://adventofcode.com/2018/day/3

import { Claim, Fabric } from './Fabric';

describe('Tracking cells claimed', function() {
    it('counts the number of overlapping cells', () => {
        let fabric = new Fabric();
        fabric.addClaim(Claim.fromString('#1 @ 1,3: 4x4'))
        fabric.addClaim(Claim.fromString('#2 @ 3, 1: 4x4'))
        fabric.addClaim(Claim.fromString('#3 @ 5, 5: 2x2'))

        expect(fabric.overlapCount()).toEqual(4);
        
    });
});