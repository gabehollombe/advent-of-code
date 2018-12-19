import { Claim, Fabric } from './Fabric';
import { getInputLines } from '../../Input';

const lines = getInputLines(__dirname, 'input.txt')
let fabric = new Fabric()


for (let line of lines) {
    const claim = Claim.fromString(line)
    fabric.addClaim(claim)
}

console.log('Part one: ', fabric.overlapCount())
console.log('Part two: ', fabric.nonOverlappingClaims().values().next().value)